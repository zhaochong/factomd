package factoid

import (
	"encoding/hex"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

/******************************************************************************/
/****************************To addresses**************************************/
/******************************************************************************/

func PublicKeyStringToECAddressString(public string) (string, error) {
	callTime := time.Now().UnixNano()
	defer factoidConversionsPublicKeyStringToECAddressString.Observe(float64(time.Now().UnixNano() - callTime))	
		pubHex, err := hex.DecodeString(public)
	if err != nil {
		return "", err
	}

	add, err := PublicKeyToECAddress(pubHex)
	if err != nil {
		return "", err
	}

	return primitives.ConvertECAddressToUserStr(add), nil
}

func PublicKeyStringToECAddress(public string) (interfaces.IAddress, error) {
	pubHex, err := hex.DecodeString(public)
	if err != nil {
		return nil, err
	}
	return PublicKeyToECAddress(pubHex)
}

func PublicKeyToECAddress(public []byte) (interfaces.IAddress, error) {
	callTime := time.Now().UnixNano()
	defer factoidConversionsPublicKeyToECAddress.Observe(float64(time.Now().UnixNano() - callTime))	
	return NewAddress(public), nil
}

func PublicKeyStringToFactoidAddressString(public string) (string, error) {
	callTime := time.Now().UnixNano()
	defer factoidConversionsPublicKeyStringToFactoidAddressString.Observe(float64(time.Now().UnixNano() - callTime))	
	pubHex, err := hex.DecodeString(public)
	if err != nil {
		return "", err
	}
	add, err := PublicKeyToFactoidAddress(pubHex)
	if err != nil {
		return "", err
	}

	return primitives.ConvertFctAddressToUserStr(add), nil
}

func PublicKeyToFactoidAddress(public []byte) (interfaces.IAddress, error) {
	callTime := time.Now().UnixNano()
	defer factoidConversionsPublicKeyToFactoidAddress.Observe(float64(time.Now().UnixNano() - callTime))	
	rcd := NewRCD_1(public)
	add, err := rcd.GetAddress()
	if err != nil {
		return nil, err
	}
	return add, nil
}

func PublicKeyStringToFactoidAddress(public string) (interfaces.IAddress, error) {
	callTime := time.Now().UnixNano()
	defer factoidConversionsPublicKeyStringToFactoidAddress.Observe(float64(time.Now().UnixNano() - callTime))	
	pubHex, err := hex.DecodeString(public)
	if err != nil {
		return nil, err
	}
	rcd := NewRCD_1(pubHex)
	add, err := rcd.GetAddress()
	if err != nil {
		return nil, err
	}
	return add, nil
}

func PublicKeyStringToFactoidRCDAddress(public string) (interfaces.IRCD, error) {
	callTime := time.Now().UnixNano()
	defer factoidConversionsPublicKeyStringToFactoidRCDAddress.Observe(float64(time.Now().UnixNano() - callTime))	
	pubHex, err := hex.DecodeString(public)
	if err != nil {
		return nil, err
	}
	rcd := NewRCD_1(pubHex)
	return rcd, nil
}

/******************************************************************************/
/******************************Combined****************************************/
/******************************************************************************/

func HumanReadiblePrivateKeyStringToEverythingString(private string) (string, string, string, error) {
	callTime := time.Now().UnixNano()
	defer factoidConversionsHumanReadiblePrivateKeyStringToEverythingString.Observe(float64(time.Now().UnixNano() - callTime))	
	priv, err := primitives.HumanReadableFactoidPrivateKeyToPrivateKeyString(private)
	if err != nil {
		return "", "", "", err
	}
	return PrivateKeyStringToEverythingString(priv)
}

func PrivateKeyStringToEverythingString(private string) (string, string, string, error) {
	callTime := time.Now().UnixNano()
	defer factoidConversionsPrivateKeyStringToEverythingString.Observe(float64(time.Now().UnixNano() - callTime))	
	pub, err := primitives.PrivateKeyStringToPublicKeyString(private)
	if err != nil {
		return "", "", "", err
	}
	add, err := PublicKeyStringToFactoidAddressString(pub)
	if err != nil {
		return "", "", "", err
	}
	return private, pub, add, nil
}
