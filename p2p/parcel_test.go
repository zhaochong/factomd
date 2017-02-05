package p2p_test

import (
	"testing"
	. "github.com/FactomProject/factomd/p2p"
)

func TestMarshalParcel (t *testing.T) {
	p := NewParcel(187, []byte("this is a test"))

	bts,err := p.MarshalBinary()
	if err != nil {
		t.Error("Failed to marshal",err)
	}
	p2 := new(Parcel)
	err = p2.UnmarshalBinary(bts)
	if err != nil {
		t.Error("Failed to Unmarshal",err)
	}

	if !p.SameAs(p2) {
		t.Error("Failed to Unmarshal Correctly",err)
	}
	bts,err = p2.MarshalBinary()
	if err != nil {
		t.Error("Failed to marshal second time",err)
	}

	err = p.UnmarshalBinary(bts)
	if err != nil {
		t.Error("Failed to Unmarshal",err)
	}

	if !p.SameAs(p2) {
		t.Error("Failed to Unmarshal Correctly second time",err)
	}
}