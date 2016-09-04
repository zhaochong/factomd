package p2p_test

import (
	"testing"

	. "github.com/FactomProject/factomd/p2p"
)

func TestParcelMarshal(t *testing.T) {
	p := new(Parcel)
	var _ = p
/*	Network     =NetworkID(1)
	Version     =5
	Type        =7
	Crc32       uint32            // 4 bytes - data integrity hash (of the payload itself.)
	NodeID      uint64
	TargetPeer  string            // ? bytes - "" or nil for broadcast, otherwise the destination peer's hash.
	PeerAddress string 						// address of the peer set by connection to know who sent message (for tracking source of other peers)
	PeerPort    string 						// port of the peer , or we are listening on
	Length      uint32            // 4 bytes - length of the payload (that follows this header) in bytes

	// ParcelHeaderSize is the number of bytes in a parcel header
	const ParcelHeaderSize = 28

	type ParcelHeader struct {
	}
*/
}
