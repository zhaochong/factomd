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

type middle struct {
	conn     net.Conn
	incoming chan *Parcel
	outgoing chan *Parcel
	closed   bool
}

var _ = fmt.Print

func (m *middle) Init() {
	if m.incoming == nil {
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
			<-m.outgoing
		}
	}
	m.outgoing <- p
	return nil
}

func (m *middle) Receive() (p *Parcel, err error) {
	select {
	case p := <-m.incoming:
		return p, nil
	default:
		time.Sleep(10 * time.Millisecond)
		return nil, nil
	}
}

// goWrite pulls from the outgoing queue, and marshals and sends parcels across the wire.
//
func (m *middle) goWrite() {
	for !m.closed {
		select {
		case p := <-m.outgoing:
			data, err := p.MarshalBinary()
			if err != nil {
				continue
			}
			m.conn.SetWriteDeadline(time.Now().Add(1 * time.Millisecond))
			m.conn.Write(magic)
			m.conn.Write(data)
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func (m *middle) FullRead(buff []byte) error {
	sum := 0
	for sum < len(buff) {
		//m.conn.SetReadDeadline(time.Now().Add(1 * time.Millisecond))
		i, _ := m.conn.Read(buff[sum:])
		if i == 0 {
			time.Sleep(10 * time.Millisecond)
		}
		sum += i
	}
	return nil
}

func (m *middle) goRead() {
	for !m.closed {

		m.Sync()

		var lengthb [4]byte
		err := m.FullRead(lengthb[:])
		if err != nil {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		length := int(binary.BigEndian.Uint32(lengthb[:]))
		data := make([]byte, length)
		copy(data[:4], lengthb[:])
		m.FullRead(data[4:])

		p := new(Parcel)
		err = p.UnmarshalBinary(data)

		m.incoming <- p
	}
}

var magic = []byte{0x7e, 0xa3, 0x9d, 0xA6}

func (m *middle) Sync() {
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
				continue loop
			}
		}
		break loop
	}
}
