package databaseOverlay

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"time"
)

func (db *Overlay) FetchFactoidTransaction(hash interfaces.IHash) (interfaces.ITransaction, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchFactoidTransaction.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	in, err := db.FetchIncludedIn(hash)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, nil
	}
	block, err := db.FetchFBlock(in)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("Block not found, should not happen")
	}
	txs := block.GetTransactions()
	for _, tx := range txs {
		if tx.GetHash().IsSameAs(hash) {
			return tx, nil
		}
		if tx.GetSigHash().IsSameAs(hash) {
			return tx, nil
		}
	}
	return nil, fmt.Errorf("Transaction not found in block, should not happen")
}

func (db *Overlay) FetchECTransaction(hash interfaces.IHash) (interfaces.IECBlockEntry, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddatabaseOverlayOverlayFetchECTransaction.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	in, err := db.FetchIncludedIn(hash)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, nil
	}
	block, err := db.FetchECBlock(in)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("Block not found, should not happen")
	}
	tx := block.GetEntryByHash(hash)
	if tx != nil {
		return tx, nil
	}
	return nil, fmt.Errorf("Transaction not found in block, should not happen")
}
