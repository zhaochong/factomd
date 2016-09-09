// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package p2p

import (
	"net"
	"time"
	"encoding/binary"
	"fmt"
)

type middle struct {
	conn net.Conn
	incoming chan *Parcel
	outgoing chan *Parcel
	closed  bool
}

func (m *middle) Init() {
	if m.incoming == nil {
		fmt.Println("Make Channels")
		m.incoming = make(chan *Parcel, 10000)
		m.outgoing = make(chan *Parcel, 10000)
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

func (m *middle) Send(p *Parcel) (err error) {
	if len(m.outgoing) > 9900 {
		for len(m.outgoing) > 9000 {
			<- m.outgoing
		}
	}
	m.outgoing <- p
	fmt.Println("Outgoing messages", len(m.outgoing))
	fmt.Println("Incoming Messages",len(m.incoming))
	return nil
}

func (m *middle) Receive() (p *Parcel, err error) {
	select {
	case p := <-m.incoming:
		fmt.Println("Outgoing messages", len(m.outgoing))
		fmt.Println("Incoming Messages",len(m.incoming)+1)
		return p, nil
	default:
		return nil, nil
	}
}

// goWrite pulls from the outgoing queue, and marshals and sends parcels across the wire.
//
func (m *middle) goWrite() {
	fmt.Println("Write Opening!")
	for !m.closed {
		select {
		case p := <-m.outgoing:
			data, err := p.MarshalBinary()
			if err != nil {
				fmt.Println("Marshal Error")
				continue
			}
			fmt.Printf("Putting %x\n", data)
			m.conn.SetWriteDeadline(time.Now().Add(1 * time.Millisecond))
			m.conn.Write(magic)
			m.conn.Write(data)
			fmt.Println("Write Data")
		default:
			fmt.Print("Write Queue empty\r")
			time.Sleep(time.Millisecond * 100)
		}
	}
	fmt.Println("Write Closing")
}

func (m *middle) FullRead(buff []byte) error {
	sum := 0
	for sum < len(buff) {
		i, _ := m.conn.Read(buff[sum:])
		if i == 0 {
			time.Sleep(100*time.Millisecond)
		}
		sum += i
	}
	return nil
}

func (m *middle) goRead() {
	fmt.Println("Read Opening")
	for !m.closed {

		m.conn.SetWriteDeadline(time.Now().Add(1 * time.Millisecond))

		m.Sync()

		var lengthb [4]byte
		err := m.FullRead(lengthb[:])
		if err != nil {
			fmt.Println("Read error ",err.Error())
			time.Sleep(time.Millisecond*100)
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

		m.incoming <- p
	}
	fmt.Println("Read Closing")
}

var magic = []byte{0x7e, 0xa3, 0x9d, 0xA6}

func (m *middle) Sync() {
	fmt.Println("Syncing")
	var b = []byte{0}
	loop:
	for {
		for b[0] != magic[0] {
			m.conn.Read(b)
		}
		for i := 1; i < len(magic); i++ {
			r := 0
			for r == 0 {
				r, _ = m.conn.Read(b)
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

