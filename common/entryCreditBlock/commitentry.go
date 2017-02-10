// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package entryCreditBlock

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"

	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

const (
	// CommitEntrySize = 1 + 6 + 32 + 1 + 32 + 64
	CommitEntrySize int = 136
)

type CommitEntry struct {
	Version   uint8
	MilliTime *primitives.ByteSlice6
	EntryHash interfaces.IHash
	Credits   uint8
	ECPubKey  *primitives.ByteSlice32
	Sig       *primitives.ByteSlice64
	SigHash   interfaces.IHash
}

var _ interfaces.Printable = (*CommitEntry)(nil)
var _ interfaces.BinaryMarshallable = (*CommitEntry)(nil)
var _ interfaces.ShortInterpretable = (*CommitEntry)(nil)
var _ interfaces.IECBlockEntry = (*CommitEntry)(nil)
var _ interfaces.ISignable = (*CommitEntry)(nil)

func (e *CommitEntry) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	var out primitives.Buffer
	out.WriteString(fmt.Sprintf(" %-20s\n", "CommitEntry"))
	out.WriteString(fmt.Sprintf("   %-20s %d\n", "Version", e.Version))
	out.WriteString(fmt.Sprintf("   %-20s %x\n", "MilliTime", e.MilliTime))
	out.WriteString(fmt.Sprintf("   %-20s %x\n", "EntryHash", e.EntryHash.Bytes()[:3]))
	out.WriteString(fmt.Sprintf("   %-20s %x\n", "Credits", e.Credits))
	out.WriteString(fmt.Sprintf("   %-20s %x\n", "ECPubKey", e.ECPubKey[:3]))
	out.WriteString(fmt.Sprintf("   %-20s %d\n", "Sig", e.Sig[:3]))

	return (string)(out.DeepCopyBytes())
}

func (a *CommitEntry) GetEntryHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryGetEntryHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return a.EntryHash
}

func (a *CommitEntry) IsSameAs(b *CommitEntry) bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if b == nil {
		return false
	}
	bin1, err := a.MarshalBinary()
	if err != nil {
		return false
	}
	bin2, err := b.MarshalBinary()
	if err != nil {
		return false
	}
	return primitives.AreBytesEqual(bin1, bin2)
}

func NewCommitEntry() *CommitEntry {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockNewCommitEntry.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c := new(CommitEntry)
	c.Version = 0
	c.MilliTime = new(primitives.ByteSlice6)
	c.EntryHash = primitives.NewZeroHash()
	c.Credits = 0
	c.ECPubKey = new(primitives.ByteSlice32)
	c.Sig = new(primitives.ByteSlice64)
	return c
}

func (e *CommitEntry) Hash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}

func (b *CommitEntry) IsInterpretable() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryIsInterpretable.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return false
}

func (b *CommitEntry) Interpret() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryInterpret.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ""
}

// CommitMsg returns the binary marshaled message section of the CommitEntry
// that is covered by the CommitEntry.Sig.
func (c *CommitEntry) CommitMsg() []byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryCommitMsg.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	p, err := c.MarshalBinary()
	if err != nil {
		return []byte{byte(0)}
	}
	return p[:len(p)-64-32]
}

// Return the timestamp
func (c *CommitEntry) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	a := make([]byte, 2, 8)
	a = append(a, c.MilliTime[:]...)
	milli := uint64(binary.BigEndian.Uint64(a))
	return primitives.NewTimestampFromMilliseconds(milli)
}

// InTime checks the CommitEntry.MilliTime and returns true if the timestamp is
// whitin +/- 12 hours of the current time.
func (c *CommitEntry) InTime() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryInTime.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	now := time.Now()
	sec := c.GetTimestamp().GetTimeSeconds()
	t := time.Unix(sec, 0)

	return t.After(now.Add(-constants.COMMIT_TIME_WINDOW*time.Hour)) && t.Before(now.Add(constants.COMMIT_TIME_WINDOW*time.Hour))
}

func (c *CommitEntry) IsValid() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryIsValid.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	//double check the credits in the commit
	if c.Credits < 1 || c.Version != 0 {
		return false
	}
	return ed.VerifyCanonical((*[32]byte)(c.ECPubKey), c.CommitMsg(), (*[64]byte)(c.Sig))
}

func (c *CommitEntry) GetHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryGetHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	h, _ := c.MarshalBinary()
	return primitives.Sha(h)
}

func (c *CommitEntry) GetSigHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryGetSigHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if c.SigHash == nil {
		data := c.CommitMsg()
		c.SigHash = primitives.Sha(data)
	}
	return c.SigHash
}

func (c *CommitEntry) MarshalBinarySig() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryMarshalBinarySig.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	buf := new(primitives.Buffer)

	// 1 byte Version
	if err := binary.Write(buf, binary.BigEndian, c.Version); err != nil {
		return nil, err
	}

	// 6 byte MilliTime
	buf.Write(c.MilliTime[:])

	// 32 byte Entry Hash
	buf.Write(c.EntryHash.Bytes())

	// 1 byte number of Entry Credits
	if err := binary.Write(buf, binary.BigEndian, c.Credits); err != nil {
		return nil, err
	}

	return buf.DeepCopyBytes(), nil

}

// Transaction hash of entry commit. (version through pub key hashed)
func (c *CommitEntry) MarshalBinaryTransaction() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryMarshalBinaryTransaction.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	buf := new(primitives.Buffer)

	b, err := c.MarshalBinarySig()
	if err != nil {
		return nil, err
	}

	buf.Write(b)

	// 32 byte Public Key
	buf.Write(c.ECPubKey[:])

	return buf.DeepCopyBytes(), nil
}

func (c *CommitEntry) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	buf := new(primitives.Buffer)

	b, err := c.MarshalBinaryTransaction()
	if err != nil {
		return nil, err
	}

	buf.Write(b)

	// 32 byte Public Key
	//buf.Write(c.ECPubKey[:])

	// 64 byte Signature
	buf.Write(c.Sig[:])

	return buf.DeepCopyBytes(), nil
}

func (c *CommitEntry) Sign(privateKey []byte) error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntrySign.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	sig, err := primitives.SignSignable(privateKey, c)
	if err != nil {
		return err
	}
	if c.Sig == nil {
		c.Sig = new(primitives.ByteSlice64)
	}
	err = c.Sig.UnmarshalBinary(sig)
	if err != nil {
		return err
	}
	pub, err := primitives.PrivateKeyToPublicKey(privateKey)
	if err != nil {
		return err
	}
	if c.ECPubKey == nil {
		c.ECPubKey = new(primitives.ByteSlice32)
	}
	err = c.ECPubKey.UnmarshalBinary(pub)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommitEntry) ValidateSignatures() error {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryValidateSignatures.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if c.ECPubKey == nil {
		return fmt.Errorf("No public key present")
	}
	if c.Sig == nil {
		return fmt.Errorf("No signature present")
	}
	data, err := c.MarshalBinarySig()
	if err != nil {
		return err
	}
	return primitives.VerifySignature(data, c.ECPubKey[:], c.Sig[:])
}

func (c *CommitEntry) ECID() byte {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryECID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return ECIDEntryCommit
}

func (c *CommitEntry) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling CommitEntry: %v", r)
		}
	}()

	buf := primitives.NewBuffer(data)
	hash := make([]byte, 32)

	var b byte
	var p []byte
	// 1 byte Version
	if b, err = buf.ReadByte(); err != nil {
		return
	} else {
		c.Version = uint8(b)
	}

	if buf.Len() < 6 {
		err = io.EOF
		return
	}

	// 6 byte MilliTime
	if p = buf.Next(6); p == nil {
		err = fmt.Errorf("Could not read MilliTime")
		return
	} else {
		c.MilliTime = new(primitives.ByteSlice6)
		err = c.MilliTime.UnmarshalBinary(p)
		if err != nil {
			return
		}
	}

	// 32 byte Entry Hash
	if _, err = buf.Read(hash); err != nil {
		return
	}
	c.EntryHash = primitives.NewHash(hash)

	// 1 byte number of Entry Credits
	if b, err = buf.ReadByte(); err != nil {
		return
	} else {
		c.Credits = uint8(b)
	}

	if buf.Len() < 32 {
		err = io.EOF
		return
	}

	// 32 byte Public Key
	if p = buf.Next(32); p == nil {
		err = fmt.Errorf("Could not read ECPubKey")
		return
	} else {
		c.ECPubKey = new(primitives.ByteSlice32)
		err = c.ECPubKey.UnmarshalBinary(p)
		if err != nil {
			return
		}
	}

	if buf.Len() < 64 {
		err = io.EOF
		return
	}

	// 64 byte Signature
	if p = buf.Next(64); p == nil {
		err = fmt.Errorf("Could not read Sig")
		return
	} else {
		c.Sig = new(primitives.ByteSlice64)
		err = c.Sig.UnmarshalBinary(p)
		if err != nil {
			return
		}
	}

	newData = buf.DeepCopyBytes()

	return
}

func (c *CommitEntry) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = c.UnmarshalBinaryData(data)
	return
}

func (e *CommitEntry) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *CommitEntry) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdentryCreditBlockCommitEntryJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}
