package p2p_test

import (
	"testing"

	. "github.com/FactomProject/factomd/p2p"
	"fmt"
)

func TestParcelMarshal(t *testing.T) {
	p := new(Parcel)
	var _ = p
	p.Header.Network = 1
	p.Header.Version = 2
	p.Header.Type = ParcelCommandType(3)
	p.Header.TargetPeer = "123"
	p.Header.NodeID = 4
	p.Header.PeerAddress = "456"
	p.Header.PeerPort = "789"
	p.Payload = []byte("This is a test")

	data,err := p.MarshalBinary()
	if err != nil {
		t.Fail()
	}

	fmt.Printf("Data: %x\n",data)

	p2 := new(Parcel)
	err = p2.UnmarshalBinary(data)
	if err != nil {
		t.Fail()
	}

	if len(data) != int(p2.Length) {
		t.Fail()
	}
}