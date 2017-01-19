package testHelper

import (
	"github.com/FactomProject/factomd/common/adminBlock"
	"github.com/FactomProject/factomd/common/directoryBlock"
	"github.com/FactomProject/factomd/common/primitives"
)

func CreateTestAdminBlock(prev *adminBlock.AdminBlock, prevDBlock *directoryBlock.DirectoryBlock) *adminBlock.AdminBlock {
	block := new(adminBlock.AdminBlock)
	block.SetHeader(CreateTestAdminHeader(prev))

	if prevDBlock != nil {
		priv, err := primitives.NewPrivateKeyFromHex("cc1985cdfae4e32b5a454dfda8ce5e1361558482684f3367649c3ad852c8e31a")
		if err != nil {
			panic(err)
		}
		bin, err := prevDBlock.GetHeader().MarshalBinary()
		if err != nil {
			panic(err)
		}
		sig := priv.Sign(bin)

		identity, err := primitives.NewShaHashFromStr("38bab1455b7bd7e5efd15c53c777c79d0c988e9210f1da49a99d95b3a6417be9")
		if err != nil {
			panic(err)
		}
		err = block.AddDBSig(identity, sig)
		if err != nil {
			panic(err)
		}
	}

	block.GetHeader().SetMessageCount(uint32(len(block.GetABEntries())))
	return block
}

func CreateTestAdminHeader(prev *adminBlock.AdminBlock) *adminBlock.ABlockHeader {
	header := new(adminBlock.ABlockHeader)

	if prev == nil {
		header.PrevBackRefHash = primitives.NewZeroHash()
		header.DBHeight = 0
	} else {
		keyMR, err := prev.GetKeyMR()
		if err != nil {
			panic(err)
		}
		header.PrevBackRefHash = keyMR
		header.DBHeight = prev.Header.GetDBHeight() + 1
	}

	header.HeaderExpansionSize = 5
	header.HeaderExpansionArea = []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	header.MessageCount = 0
	header.BodySize = 0

	return header
}
