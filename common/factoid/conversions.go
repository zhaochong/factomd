package factoid

import (
	"encoding/hex"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

/******************************************************************************/
/****************************To addresses**************************************/
/******************************************************************************/

func PublicKeyStringToECAddressString(public string) (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidPublicKeyStringToECAddressString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidPublicKeyStringToECAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	pubHex, err := hex.DecodeString(public)
	if err != nil {
		return nil, err
	}
	return PublicKeyToECAddress(pubHex)
}

func PublicKeyToECAddress(public []byte) (interfaces.IAddress, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidPublicKeyToECAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return NewAddress(public), nil
}

func PublicKeyStringToFactoidAddressString(public string) (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidPublicKeyStringToFactoidAddressString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidPublicKeyToFactoidAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	rcd := NewRCD_1(public)
	add, err := rcd.GetAddress()
	if err != nil {
		return nil, err
	}
	return add, nil
}

func PublicKeyStringToFactoidAddress(public string) (interfaces.IAddress, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidPublicKeyStringToFactoidAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidPublicKeyStringToFactoidRCDAddress.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidHumanReadiblePrivateKeyStringToEverythingString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	priv, err := primitives.HumanReadableFactoidPrivateKeyToPrivateKeyString(private)
	if err != nil {
		return "", "", "", err
	}
	return PrivateKeyStringToEverythingString(priv)
}

func PrivateKeyStringToEverythingString(private string) (string, string, string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomdfactoidPrivateKeyStringToEverythingString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

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
