package processState

import (
	"github.com/FactomProject/factomd/btcd/wire"
	"github.com/FactomProject/factomd/consensus"
	"github.com/FactomProject/factomd/database"

	. "github.com/FactomProject/factomd/common"
	. "github.com/FactomProject/factomd/common/AdminBlock"
	. "github.com/FactomProject/factomd/common/DirectoryBlock"
	. "github.com/FactomProject/factomd/common/EntryBlock"
	. "github.com/FactomProject/factomd/common/EntryCreditBlock"

	. "github.com/FactomProject/factomd/common/primitives"
)

type ProcessState struct {
	DB       database.Db // database
	DChain   *DChain     //Directory Block Chain
	ECChain  *ECChain    //Entry Credit Chain
	AChain   *AdminChain //Admin Chain
	FChain   *FctChain   // factoid Chain
	FChainID *Hash

	InMsgQueue  chan wire.FtmInternalMsg //incoming message queue for factom application messages
	OutMsgQueue chan wire.FtmInternalMsg //outgoing message queue for factom application messages

	InCtlMsgQueue  chan wire.FtmInternalMsg //incoming message queue for factom control messages
	OutCtlMsgQueue chan wire.FtmInternalMsg //outgoing message queue for factom control messages

	//TODO: To be moved to ftmMemPool??
	ChainIDMap     map[string]*EChain // ChainIDMap with chainID string([32]byte) as key
	CommitChainMap map[string]*CommitChain
	CommitEntryMap map[string]*CommitEntry
	ECreditMap     map[string]int32 // eCreditMap with public key string([32]byte) as key, credit balance as value

	ChainIDMapBackup map[string]*EChain //previous block bakcup - ChainIDMap with chainID string([32]byte) as key
	ECreditMapBackup map[string]int32   // backup from previous block - eCreditMap with public key string([32]byte) as key, credit balance as value

	FMemPool *FtmMemPool
	PlMgr    *consensus.ProcessListMgr

	//Server Private key and Public key for milestone 1
	ServerPrivKey PrivateKey
	ServerPubKey  PublicKey

	FactoshisPerCredit uint64 // .001 / .15 * 100000000 (assuming a Factoid is .15 cents, entry credit = .1 cents

	FactomdUser string
	FactomdPass string

	DirectoryBlockInSeconds int
	DataStorePath           string
	Ldbpath                 string
	NodeMode                string
	DevNet                  bool
	ServerPrivKeyHex        string
	ServerIndex             *ServerIndexNumber
}
