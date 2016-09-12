// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package p2p

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

var _ = fmt.Print

type middle struct {
	conn    net.Conn
	encoder *gob.Encoder // Wire format is gobs in this version, may switch to binary
	decoder *gob.Decoder // Wire format is gobs in this version, may switch to binary

	output chan *Parcel
	input  chan *Parcel

	isNew bool
}

var Writes int
var Reads int
var WritesErr int
var ReadsErr int

var Deadline time.Duration = time.Duration(100)

type ParcelPack struct {
	Payload []byte
}

func (m *middle) Init() {
	m.encoder = gob.NewEncoder(m)
	m.decoder = gob.NewDecoder(m)

	m.input = make(chan *Parcel, 10000)

	go m.goInput()
}

func (m *middle) Close() {
	m.conn.Close()
	m.decoder = nil
	m.encoder = nil
}

func (m *middle) goOutput(p *Parcel) {

	defer func() {
		if r := recover(); r != nil {
			time.Sleep(100 * time.Millisecond)
			return
		}
	}()

}

func (m *middle) goInput() {

	defer func() {
		if r := recover(); r != nil {
			time.Sleep(100 * time.Millisecond)
			go m.goInput()
		}
	}()

	for m.decoder != nil {
		var pack ParcelPack
		p := new(Parcel)
		err := m.decoder.Decode(&pack)
		if len(pack.Payload) > 0 {
			err = p.UnmarshalBinary(pack.Payload)
			if err != nil {
				continue
			}
			m.input <- p
		}
	}
}

func (m *middle) Send(p Parcel) (err error) {
	if m.encoder != nil {
		pack := new(ParcelPack)
		var err error
		pack.Payload, err = p.MarshalBinary()
		if err != nil && len(pack.Payload) == 0 {
			return err
		}
		m.encoder.Encode(pack)
	}
	return err
}

func (m *middle) Receive() (p *Parcel, err error) {
	select {
	case p = <-m.input:
	default:
	}
	return
}

func (m *middle) Write(b []byte) (int, error) {

	m.conn.SetWriteDeadline(time.Now().Add(Deadline))

	i, e := m.conn.Write(b)

	Writes += i

	if i > 0 {
		e = nil
	}

	if e != nil {
		WritesErr++
	}
	//fmt.Println("Write Done",time.Now().String())
	return i, e
}

func (m *middle) Read(b []byte) (int, error) {

	m.conn.SetReadDeadline(time.Now().Add(Deadline))

	i, e := m.conn.Read(b)

	if i > 0 {
		e = nil
	}

	//end := 10
	//if end > len(b) {
	//	end = len(b)
	//}
	//if e == nil {
	//	fmt.Printf("bbbb Read  %s %d bytes, Data: %x\n", time.Now().String(), len(b), b[:end])
	//}
	Reads += i

	if e != nil {
		ReadsErr++
	}
	//fmt.Println("Read Done",time.Now().String())
	return i, e
}
