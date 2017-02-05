// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package p2p

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Wire interface {
	Send(p *Parcel) (err error)
	Receive() (p *Parcel, err error)
	Init(*Connection, net.Conn)
	Close()
}

type WireSerializer struct {
	Connection *Connection
	Conn       net.Conn
	Incoming   chan *Parcel
	Outgoing   chan *Parcel
	closed     bool
}

var _ Wire = (*WireSerializer)(nil)

func (m *WireSerializer) Close() {
	m.Conn.Close()
	m.closed = true
}

func (m *WireSerializer) Init(connection *Connection, conn net.Conn) {
	m.Connection = connection
	m.Conn = conn
	if m.Incoming == nil {
		fmt.Println("Make Channels")
		m.Incoming = make(chan *Parcel, 10000)
		m.Outgoing = make(chan *Parcel, 10000)
	}
	m.closed = false
	go m.goRead()
	go m.goWrite()
}

var Writes int
var Reads int
var WritesErr int
var ReadsErr int

var Deadline time.Duration = time.Duration(100)

func (m *WireSerializer) Send(p *Parcel) (err error) {
	if len(m.Outgoing) > 9900 {
		for len(m.Outgoing) > 9000 {
			<-m.Outgoing
		}
	}
	m.Outgoing <- p
	fmt.Println("Outgoing messages", len(m.Outgoing))
	fmt.Println("Incoming Messages", len(m.Incoming))
	return nil
}

func (m *WireSerializer) Receive() (p *Parcel, err error) {
	select {
	case p := <-m.Incoming:
		fmt.Println("Outgoing messages", len(m.Outgoing))
		fmt.Println("Incoming Messages", len(m.Incoming)+1)
		return p, nil
	default:
		return nil, nil
	}
}

// goWrite pulls from the outgoing queue, and marshals and sends parcels across the wire.
//
func (m *WireSerializer) goWrite() {
	fmt.Println("Write Opening!")
	for !m.closed {
		select {
		case p := <-m.Outgoing:
			data, err := p.MarshalBinary()
			if err != nil {
				fmt.Println("Marshal Error")
				continue
			}
			fmt.Printf("Putting %x\n", data)
			m.Conn.SetWriteDeadline(time.Now().Add(1 * time.Millisecond))
			m.Conn.Write(magic)
			m.Conn.Write(data)
			fmt.Println("Write Data")
		default:
			fmt.Print("Write Queue empty\r")
			time.Sleep(time.Millisecond * 100)
		}
	}
	fmt.Println("Write Closing")
}

func (m *WireSerializer) FullRead(buff []byte) error {
	sum := 0
	for sum < len(buff) {
		i, _ := m.Conn.Read(buff[sum:])
		if i == 0 {
			time.Sleep(100 * time.Millisecond)
		}
		sum += i
	}
	return nil
}

func (m *WireSerializer) goRead() {
	fmt.Println("Read Opening")
	for !m.closed {

		m.Conn.SetWriteDeadline(time.Now().Add(1 * time.Millisecond))

		m.Sync()

		var lengthb [4]byte
		err := m.FullRead(lengthb[:])
		if err != nil {
			fmt.Println("Read error ", err.Error())
			time.Sleep(time.Millisecond * 100)
			continue
		}

		length := int(binary.BigEndian.Uint32(lengthb[:]))
		data := make([]byte, length)
		fmt.Println("Length:", length, len(data))
		copy(data[:4], lengthb[:])
		m.FullRead(data[4:])

		p := new(Parcel)
		fmt.Printf("Getting %x\n", data)
		err = p.UnmarshalBinary(data)

		m.Incoming <- p
	}
	fmt.Println("Read Closing")
}

var magic = []byte{0x7e, 0xa3, 0x9d, 0xA6}

func (m *WireSerializer) Sync() {
	fmt.Println("Syncing")
	var b = []byte{0}
loop:
	for {
		for b[0] != magic[0] {
			m.Conn.Read(b)
		}
		for i := 1; i < len(magic); i++ {
			r := 0
			for r == 0 {
				r, _ = m.Conn.Read(b)
			}
			if b[0] != magic[i] {
				fmt.Print("Start over...\n")
				continue loop
			}
		}
		break loop
	}
	fmt.Print("Synced\n")
}
