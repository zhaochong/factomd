package p2p

import (
	"encoding/gob"
	"net"
	"time"
)

type WireGob struct {
	Connection *Connection
	Conn       net.Conn
	encoder    *gob.Encoder // Wire format is gobs in this version, may switch to binary
	decoder    *gob.Decoder // Wire format is gobs in this version, may switch to binaryclosed   bool
	closed     bool
}

var _ Wire = (*WireGob)(nil)

func (w *WireGob) Init(connection *Connection, conn net.Conn) {
	w.Connection = connection
	w.Conn = conn
	w.encoder = gob.NewEncoder(conn)
	w.decoder = gob.NewDecoder(conn)
}

func (w *WireGob) Close() {
	w.decoder = nil
	w.encoder = nil
}

func (w *WireGob) Send(parcel *Parcel) (err error) {
	w.Conn.SetWriteDeadline(time.Now().Add(NetworkDeadline))
	parcel.Trace("Connection.sendParcel().encoder.Encode(parcel)", "f")
	err = w.encoder.Encode(parcel)
	return
}

func (w *WireGob) Receive() (p *Parcel, err error) {
	var message Parcel
	w.Conn.SetReadDeadline(time.Now().Add(NetworkDeadline))
	err = w.decoder.Decode(&message)
	return &message, err
}
