// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

//A placeholder structure for messages
type FactoidTransaction struct {
	MessageBase
	Transaction interfaces.ITransaction

	//No signature!

	//Not marshalled
	hash      interfaces.IHash
	processed bool
}

var _ interfaces.IMsg = (*FactoidTransaction)(nil)

func (a *FactoidTransaction) IsSameAs(b *FactoidTransaction) bool {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionIsSameAs.Observe(float64(time.Now().UnixNano() - callTime))
	if b == nil {
		return false
	}

	ok, err := primitives.AreBinaryMarshallablesEqual(a.Transaction, b.Transaction)
	if err != nil || ok == false {
		return false
	}

	return true
}

func (m *FactoidTransaction) GetRepeatHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionGetRepeatHash.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Transaction.GetSigHash()
}

func (m *FactoidTransaction) GetHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionGetHash.Observe(float64(time.Now().UnixNano() - callTime))

	if m.hash == nil {
		m.SetFullMsgHash(m.Transaction.GetFullHash())

		data, err := m.Transaction.MarshalBinarySig()
		if err != nil {
			panic(fmt.Sprintf("Error in CommitChain.GetHash(): %s", err.Error()))
		}
		m.hash = primitives.Sha(data)
	}

	return m.hash
}

func (m *FactoidTransaction) GetMsgHash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionGetMsgHash.Observe(float64(time.Now().UnixNano() - callTime))
	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *FactoidTransaction) GetTimestamp() interfaces.Timestamp {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Transaction.GetTimestamp()
}

func (m *FactoidTransaction) GetTransaction() interfaces.ITransaction {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionGetTransaction.Observe(float64(time.Now().UnixNano() - callTime))
	return m.Transaction
}

func (m *FactoidTransaction) SetTransaction(transaction interfaces.ITransaction) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionSetTransaction.Observe(float64(time.Now().UnixNano() - callTime))
	m.Transaction = transaction
}

func (m *FactoidTransaction) Type() byte {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransaction.Observe(float64(time.Now().UnixNano() - callTime))
	return constants.FACTOID_TRANSACTION_MSG
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *FactoidTransaction) Validate(state interfaces.IState) int {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionValidate.Observe(float64(time.Now().UnixNano() - callTime))
	// Is the transaction well formed?
	err := m.Transaction.Validate(1)
	if err != nil {
		return -1 // No, object!
	}

	// Is the transaction properly signed?
	err = m.Transaction.ValidateSignatures()
	if err != nil {
		return -1 // No, object!
	}

	// Is the transaction valid at this point in time?
	err = state.GetFactoidState().Validate(1, m.Transaction)
	if err != nil {
		return 0 // Well, mumble.  Might be out of order.
	}
	return 1
}

func (m *FactoidTransaction) ComputeVMIndex(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionComputeVMIndex.Observe(float64(time.Now().UnixNano() - callTime))
	m.VMIndex = state.ComputeVMIndex(constants.FACTOID_CHAINID)
}

// Execute the leader functions of the given message
func (m *FactoidTransaction) LeaderExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionLeaderExecute.Observe(float64(time.Now().UnixNano() - callTime))
	state.LeaderExecute(m)
}

func (m *FactoidTransaction) FollowerExecute(state interfaces.IState) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionFollowerExecute.Observe(float64(time.Now().UnixNano() - callTime))
	state.FollowerExecuteMsg(m)
}

func (m *FactoidTransaction) Process(dbheight uint32, state interfaces.IState) bool {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionProcess.Observe(float64(time.Now().UnixNano() - callTime))
	if m.processed {
		return true
	}
	m.processed = true
	err := state.GetFactoidState().AddTransaction(1, m.Transaction)
	if err != nil {
		fmt.Println(err)
		return false
	}

	state.IncFactoidTrans()

	return true

}

func (m *FactoidTransaction) UnmarshalTransData(datax []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionUnmarshalTransData.Observe(float64(time.Now().UnixNano() - callTime))
	newData = datax
	defer func() {
		return
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Transaction Factoid: %v", r)
		}
	}()

	m.Transaction = new(factoid.Transaction)
	newData, err = m.Transaction.UnmarshalBinaryData(newData)

	return newData, err
}

func (m *FactoidTransaction) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))

	newData = data

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Factoid: %v", r)
		}
	}()
	if newData[0] != m.Type() {
		return nil, fmt.Errorf("Invalid Message type")
	}
	newData = newData[1:]

	m.Transaction = new(factoid.Transaction)
	newData, err = m.Transaction.UnmarshalBinaryData(newData)
	return newData, err
}

func (m *FactoidTransaction) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *FactoidTransaction) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransaction.Observe(float64(time.Now().UnixNano() - callTime))
	var buf primitives.Buffer
	buf.Write([]byte{m.Type()})

	if d, err := m.Transaction.MarshalBinary(); err != nil {
		return nil, err
	} else {
		buf.Write(d)
	}

	return buf.DeepCopyBytes(), nil
}

func (m *FactoidTransaction) String() string {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionString.Observe(float64(time.Now().UnixNano() - callTime))
	return fmt.Sprintf("Factoid VM %d Leader %x GetHash %x",
		m.VMIndex,
		m.GetLeaderChainID().Bytes()[:3],
		m.GetHash().Bytes()[:3])
}

func (e *FactoidTransaction) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSON(e)
}

func (e *FactoidTransaction) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer messagesFactoidTransactionJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	return primitives.EncodeJSONString(e)
}
