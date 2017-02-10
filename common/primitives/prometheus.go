package primitives

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdprimitivesBufferDeepCopyBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Buffer_DeepCopyBytes_Summary",
		Help: "Summary of calls to  factomd_primitives_Buffer_DeepCopyBytes",
	})
	factomdprimitivesNewBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewBuffer_Summary",
		Help: "Summary of calls to  factomd_primitives_NewBuffer",
	})
	factomdprimitivesAreBytesEqual = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_AreBytesEqual_Summary",
		Help: "Summary of calls to  factomd_primitives_AreBytesEqual",
	})
	factomdprimitivesAreBinaryMarshallablesEqual = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_AreBinaryMarshallablesEqual_Summary",
		Help: "Summary of calls to  factomd_primitives_AreBinaryMarshallablesEqual",
	})
	factomdprimitivesEncodeBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_EncodeBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_EncodeBinary",
	})
	factomdprimitivesDecodeBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_DecodeBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_DecodeBinary",
	})
	factomdprimitivesStringToByteSlice32 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_StringToByteSlice32_Summary",
		Help: "Summary of calls to  factomd_primitives_StringToByteSlice32",
	})
	factomdprimitivesByte32ToByteSlice32 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Byte32ToByteSlice32_Summary",
		Help: "Summary of calls to  factomd_primitives_Byte32ToByteSlice32",
	})
	factomdprimitivesByteSlice32MarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice32_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice32_MarshalBinary",
	})
	factomdprimitivesByteSlice32Fixed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice32_Fixed_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice32_Fixed",
	})
	factomdprimitivesByteSlice32UnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice32_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice32_UnmarshalBinaryData",
	})
	factomdprimitivesByteSlice32UnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice32_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice32_UnmarshalBinary",
	})
	factomdprimitivesByteSlice32JSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice32_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice32_JSONByte",
	})
	factomdprimitivesByteSlice32JSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice32_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice32_JSONString",
	})
	factomdprimitivesByteSlice32String = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice32_String_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice32_String",
	})
	factomdprimitivesByteSlice32MarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice32_MarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice32_MarshalText",
	})
	factomdprimitivesByteSlice64MarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice64_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice64_MarshalBinary",
	})
	factomdprimitivesByteSlice64UnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice64_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice64_UnmarshalBinaryData",
	})
	factomdprimitivesByteSlice64UnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice64_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice64_UnmarshalBinary",
	})
	factomdprimitivesByteSlice64JSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice64_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice64_JSONByte",
	})
	factomdprimitivesByteSlice64JSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice64_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice64_JSONString",
	})
	factomdprimitivesByteSlice64String = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice64_String_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice64_String",
	})
	factomdprimitivesByteSlice64MarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice64_MarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice64_MarshalText",
	})
	factomdprimitivesByteSlice6MarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice6_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice6_MarshalBinary",
	})
	factomdprimitivesByteSlice6UnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice6_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice6_UnmarshalBinaryData",
	})
	factomdprimitivesByteSlice6UnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice6_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice6_UnmarshalBinary",
	})
	factomdprimitivesByteSlice6JSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice6_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice6_JSONByte",
	})
	factomdprimitivesByteSlice6JSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice6_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice6_JSONString",
	})
	factomdprimitivesByteSlice6String = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice6_String_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice6_String",
	})
	factomdprimitivesByteSlice6MarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice6_MarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice6_MarshalText",
	})
	factomdprimitivesByteSliceSigMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_MarshalBinary",
	})
	factomdprimitivesByteSliceSigGetFixed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_GetFixed_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_GetFixed",
	})
	factomdprimitivesByteSliceSigUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_UnmarshalBinaryData",
	})
	factomdprimitivesByteSliceSigUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_UnmarshalBinary",
	})
	factomdprimitivesByteSliceSigJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_JSONByte",
	})
	factomdprimitivesByteSliceSigJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_JSONString",
	})
	factomdprimitivesByteSliceSigString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_String_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_String",
	})
	factomdprimitivesByteSliceSigMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_MarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_MarshalText",
	})
	factomdprimitivesByteSliceSigUnmarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSliceSig_UnmarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSliceSig_UnmarshalText",
	})
	factomdprimitivesByteSlice20MarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice20_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice20_MarshalBinary",
	})
	factomdprimitivesByteSlice20GetFixed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice20_GetFixed_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice20_GetFixed",
	})
	factomdprimitivesByteSlice20UnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice20_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice20_UnmarshalBinaryData",
	})
	factomdprimitivesByteSlice20UnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice20_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice20_UnmarshalBinary",
	})
	factomdprimitivesByteSlice20JSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice20_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice20_JSONByte",
	})
	factomdprimitivesByteSlice20JSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice20_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice20_JSONString",
	})
	factomdprimitivesByteSlice20String = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice20_String_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice20_String",
	})
	factomdprimitivesByteSlice20MarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice20_MarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice20_MarshalText",
	})
	factomdprimitivesStringToByteSlice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_StringToByteSlice_Summary",
		Help: "Summary of calls to  factomd_primitives_StringToByteSlice",
	})
	factomdprimitivesByteSliceNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice_New_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice_New",
	})
	factomdprimitivesByteSliceMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice_MarshalBinary",
	})
	factomdprimitivesByteSliceUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice_UnmarshalBinaryData",
	})
	factomdprimitivesByteSliceUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice_UnmarshalBinary",
	})
	factomdprimitivesByteSliceJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice_JSONByte",
	})
	factomdprimitivesByteSliceJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice_JSONString",
	})
	factomdprimitivesByteSliceString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice_String_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice_String",
	})
	factomdprimitivesByteSliceMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ByteSlice_MarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_ByteSlice_MarshalText",
	})
	factomdprimitivesMnemonicStringToPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_MnemonicStringToPrivateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_MnemonicStringToPrivateKey",
	})
	factomdprimitivesMnemonicStringToPrivateKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_MnemonicStringToPrivateKeyString_Summary",
		Help: "Summary of calls to  factomd_primitives_MnemonicStringToPrivateKeyString",
	})
	factomdprimitivesHumanReadableFactoidPrivateKeyToPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_HumanReadableFactoidPrivateKeyToPrivateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_HumanReadableFactoidPrivateKeyToPrivateKey",
	})
	factomdprimitivesHumanReadableFactoidPrivateKeyToPrivateKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_HumanReadableFactoidPrivateKeyToPrivateKeyString_Summary",
		Help: "Summary of calls to  factomd_primitives_HumanReadableFactoidPrivateKeyToPrivateKeyString",
	})
	factomdprimitivesHumanReadableECPrivateKeyToPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_HumanReadableECPrivateKeyToPrivateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_HumanReadableECPrivateKeyToPrivateKey",
	})
	factomdprimitivesHumanReadableECPrivateKeyToPrivateKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_HumanReadableECPrivateKeyToPrivateKeyString_Summary",
		Help: "Summary of calls to  factomd_primitives_HumanReadableECPrivateKeyToPrivateKeyString",
	})
	factomdprimitivesPrivateKeyStringToHumanReadableFactoidPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKeyStringToHumanReadableFactoidPrivateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKeyStringToHumanReadableFactoidPrivateKey",
	})
	factomdprimitivesPrivateKeyStringToHumanReadableECPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKeyStringToHumanReadableECPrivateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKeyStringToHumanReadableECPrivateKey",
	})
	factomdprimitivesPrivateKeyStringToHumanReadablePrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKeyStringToHumanReadablePrivateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKeyStringToHumanReadablePrivateKey",
	})
	factomdprimitivesPrivateKeyStringToPublicKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKeyStringToPublicKeyString_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKeyStringToPublicKeyString",
	})
	factomdprimitivesPrivateKeyStringToPublicKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKeyStringToPublicKey_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKeyStringToPublicKey",
	})
	factomdprimitivesPrivateKeyToPublicKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKeyToPublicKey_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKeyToPublicKey",
	})
	factomdprimitivesGenerateKeyFromPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_GenerateKeyFromPrivateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_GenerateKeyFromPrivateKey",
	})
	factomdprimitivesErrorError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Error_Error_Summary",
		Help: "Summary of calls to  factomd_primitives_Error_Error",
	})
	factomdprimitivesCreateError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_CreateError_Summary",
		Help: "Summary of calls to  factomd_primitives_CreateError",
	})
	factomdprimitivesretreiveErrorParameters = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_retreiveErrorParameters_Summary",
		Help: "Summary of calls to  factomd_primitives_retreiveErrorParameters",
	})
	factomdprimitivesRandomHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_RandomHash_Summary",
		Help: "Summary of calls to  factomd_primitives_RandomHash",
	})
	factomdprimitivesHashCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_Copy_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_Copy",
	})
	factomdprimitivesHashNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_New_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_New",
	})
	factomdprimitivesHashMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_MarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_MarshalText",
	})
	factomdprimitivesHashIsZero = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_IsZero_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_IsZero",
	})
	factomdprimitivesNewShaHashFromStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewShaHashFromStr_Summary",
		Help: "Summary of calls to  factomd_primitives_NewShaHashFromStr",
	})
	factomdprimitivesHashUnmarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_UnmarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_UnmarshalText",
	})
	factomdprimitivesHashFixed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_Fixed_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_Fixed",
	})
	factomdprimitivesHashBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_Bytes_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_Bytes",
	})
	factomdprimitivesGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_GetHash(__Summary",
		Help: "Summary of calls to  factomd_primitives_GetHash(_",
	})
	factomdprimitivesCreateHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_CreateHash_Summary",
		Help: "Summary of calls to  factomd_primitives_CreateHash",
	})
	factomdprimitivesHashMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_MarshalBinary",
	})
	factomdprimitivesHashUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_UnmarshalBinaryData",
	})
	factomdprimitivesHashUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_UnmarshalBinary",
	})
	factomdprimitivesHashGetBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_GetBytes_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_GetBytes",
	})
	factomdprimitivesHashSetBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_SetBytes_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_SetBytes",
	})
	factomdprimitivesNewShaHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewShaHash_Summary",
		Help: "Summary of calls to  factomd_primitives_NewShaHash",
	})
	factomdprimitivesSha512Half = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Sha512Half_Summary",
		Help: "Summary of calls to  factomd_primitives_Sha512Half",
	})
	factomdprimitivesHashString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_String_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_String",
	})
	factomdprimitivesHashByteString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_ByteString_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_ByteString",
	})
	factomdprimitivesHexToHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_HexToHash_Summary",
		Help: "Summary of calls to  factomd_primitives_HexToHash",
	})
	factomdprimitivesHashIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_IsSameAs",
	})
	factomdprimitivesHashIsMinuteMarker = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_IsMinuteMarker_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_IsMinuteMarker",
	})
	factomdprimitivesHashJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_JSONByte",
	})
	factomdprimitivesHashJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Hash_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_Hash_JSONString",
	})
	factomdprimitivesSha = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Sha_Summary",
		Help: "Summary of calls to  factomd_primitives_Sha",
	})
	factomdprimitivesShad = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Shad_Summary",
		Help: "Summary of calls to  factomd_primitives_Shad",
	})
	factomdprimitivesNewZeroHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewZeroHash_Summary",
		Help: "Summary of calls to  factomd_primitives_NewZeroHash",
	})
	factomdprimitivesNewHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewHash_Summary",
		Help: "Summary of calls to  factomd_primitives_NewHash",
	})
	factomdprimitivesDoubleSha = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_DoubleSha_Summary",
		Help: "Summary of calls to  factomd_primitives_DoubleSha",
	})
	factomdprimitivesNewShaHashFromStruct = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewShaHashFromStruct_Summary",
		Help: "Summary of calls to  factomd_primitives_NewShaHashFromStruct",
	})
	factomdprimitivesJSON2RequestJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_JSON2Request_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_JSON2Request_JSONByte",
	})
	factomdprimitivesJSON2RequestJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_JSON2Request_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_JSON2Request_JSONString",
	})
	factomdprimitivesJSON2RequestString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_JSON2Request_String_Summary",
		Help: "Summary of calls to  factomd_primitives_JSON2Request_String",
	})
	factomdprimitivesNewJSON2RequestBlank = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewJSON2RequestBlank_Summary",
		Help: "Summary of calls to  factomd_primitives_NewJSON2RequestBlank",
	})
	factomdprimitivesNewJSON2Request = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewJSON2Request_Summary",
		Help: "Summary of calls to  factomd_primitives_NewJSON2Request",
	})
	factomdprimitivesParseJSON2Request = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ParseJSON2Request_Summary",
		Help: "Summary of calls to  factomd_primitives_ParseJSON2Request",
	})
	factomdprimitivesJSON2ResponseJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_JSON2Response_JSONByte_Summary",
		Help: "Summary of calls to  factomd_primitives_JSON2Response_JSONByte",
	})
	factomdprimitivesJSON2ResponseJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_JSON2Response_JSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_JSON2Response_JSONString",
	})
	factomdprimitivesJSON2ResponseString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_JSON2Response_String_Summary",
		Help: "Summary of calls to  factomd_primitives_JSON2Response_String",
	})
	factomdprimitivesNewJSON2Response = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewJSON2Response_Summary",
		Help: "Summary of calls to  factomd_primitives_NewJSON2Response",
	})
	factomdprimitivesJSON2ResponseAddError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_JSON2Response_AddError_Summary",
		Help: "Summary of calls to  factomd_primitives_JSON2Response_AddError",
	})
	factomdprimitivesNewJSONError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewJSONError_Summary",
		Help: "Summary of calls to  factomd_primitives_NewJSONError",
	})
	factomdprimitivesJSONErrorError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_JSONError_Error_Summary",
		Help: "Summary of calls to  factomd_primitives_JSONError_Error",
	})
	factomdprimitivesPrivateKeyInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_Init_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_Init",
	})
	factomdprimitivesRandomPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_RandomPrivateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_RandomPrivateKey",
	})
	factomdprimitivesPrivateKeyCustomMarshalText2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_CustomMarshalText2_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_CustomMarshalText2",
	})
	factomdprimitivesPrivateKeyPublic = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_Public_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_Public",
	})
	factomdprimitivesPrivateKeyAllocateNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_AllocateNew_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_AllocateNew",
	})
	factomdprimitivesNewPrivateKeyFromHex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewPrivateKeyFromHex_Summary",
		Help: "Summary of calls to  factomd_primitives_NewPrivateKeyFromHex",
	})
	factomdprimitivesNewPrivateKeyFromHexBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewPrivateKeyFromHexBytes_Summary",
		Help: "Summary of calls to  factomd_primitives_NewPrivateKeyFromHexBytes",
	})
	factomdprimitivesPrivateKeySign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_Sign_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_Sign",
	})
	factomdprimitivesPrivateKeyMarshalSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_MarshalSign_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_MarshalSign",
	})
	factomdprimitivesPrivateKeyGenerateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_GenerateKey_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_GenerateKey",
	})
	factomdprimitivesPrivateKeyPrivateKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_PrivateKeyString_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_PrivateKeyString",
	})
	factomdprimitivesPrivateKeyPublicKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PrivateKey_PublicKeyString_Summary",
		Help: "Summary of calls to  factomd_primitives_PrivateKey_PublicKeyString",
	})
	factomdprimitivesPublicKeyCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_Copy_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_Copy",
	})
	factomdprimitivesPublicKeyFixed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_Fixed_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_Fixed",
	})
	factomdprimitivesPublicKeyIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_IsSameAs",
	})
	factomdprimitivesPublicKeyMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_MarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_MarshalText",
	})
	factomdprimitivesPublicKeyUnmarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_UnmarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_UnmarshalText",
	})
	factomdprimitivesPublicKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_String_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_String",
	})
	factomdprimitivesPubKeyFromString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PubKeyFromString_Summary",
		Help: "Summary of calls to  factomd_primitives_PubKeyFromString",
	})
	factomdprimitivesPublicKeyVerify = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_Verify_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_Verify",
	})
	factomdprimitivesPublicKeyMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_MarshalBinary",
	})
	factomdprimitivesPublicKeyUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_UnmarshalBinaryData",
	})
	factomdprimitivesPublicKeyUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_PublicKey_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_PublicKey_UnmarshalBinary",
	})
	factomdprimitivesVerify = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Verify_Summary",
		Help: "Summary of calls to  factomd_primitives_Verify",
	})
	factomdprimitivesVerifySlice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_VerifySlice_Summary",
		Help: "Summary of calls to  factomd_primitives_VerifySlice",
	})
	factomdprimitivesLog = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Log_Summary",
		Help: "Summary of calls to  factomd_primitives_Log",
	})
	factomdprimitivesLogJSONs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_LogJSONs_Summary",
		Help: "Summary of calls to  factomd_primitives_LogJSONs",
	})
	factomdprimitivesDecodeJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_DecodeJSON_Summary",
		Help: "Summary of calls to  factomd_primitives_DecodeJSON",
	})
	factomdprimitivesEncodeJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_EncodeJSON_Summary",
		Help: "Summary of calls to  factomd_primitives_EncodeJSON",
	})
	factomdprimitivesEncodeJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_EncodeJSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_EncodeJSONString",
	})
	factomdprimitivesDecodeJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_DecodeJSONString_Summary",
		Help: "Summary of calls to  factomd_primitives_DecodeJSONString",
	})
	factomdprimitivesEncodeJSONToBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_EncodeJSONToBuffer_Summary",
		Help: "Summary of calls to  factomd_primitives_EncodeJSONToBuffer",
	})
	factomdprimitivesNextPowerOfTwo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NextPowerOfTwo_Summary",
		Help: "Summary of calls to  factomd_primitives_NextPowerOfTwo",
	})
	factomdprimitivesHashMerkleBranches = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_HashMerkleBranches_Summary",
		Help: "Summary of calls to  factomd_primitives_HashMerkleBranches",
	})
	factomdprimitivesComputeMerkleRoot = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ComputeMerkleRoot_Summary",
		Help: "Summary of calls to  factomd_primitives_ComputeMerkleRoot",
	})
	factomdprimitivesBuildMerkleTreeStore = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_BuildMerkleTreeStore_Summary",
		Help: "Summary of calls to  factomd_primitives_BuildMerkleTreeStore",
	})
	factomdprimitivesBuildMerkleBranchForEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_BuildMerkleBranchForEntryHash_Summary",
		Help: "Summary of calls to  factomd_primitives_BuildMerkleBranchForEntryHash",
	})
	factomdprimitivesBuildMerkleBranch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_BuildMerkleBranch_Summary",
		Help: "Summary of calls to  factomd_primitives_BuildMerkleBranch",
	})
	factomdprimitivesSignatureInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_Init_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_Init",
	})
	factomdprimitivesSignatureGetPubBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_GetPubBytes_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_GetPubBytes",
	})
	factomdprimitivesSignatureGetSigBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_GetSigBytes_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_GetSigBytes",
	})
	factomdprimitivesRandomSignatureSet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_RandomSignatureSet_Summary",
		Help: "Summary of calls to  factomd_primitives_RandomSignatureSet",
	})
	factomdprimitivesSignatureIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_IsSameAs",
	})
	factomdprimitivesSignatureCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_CustomMarshalText_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_CustomMarshalText",
	})
	factomdprimitivesSignatureBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_Bytes_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_Bytes",
	})
	factomdprimitivesSignatureSetPub = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_SetPub_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_SetPub",
	})
	factomdprimitivesSignatureGetKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_GetKey_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_GetKey",
	})
	factomdprimitivesSignatureSetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_SetSignature_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_SetSignature",
	})
	factomdprimitivesSignatureGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_GetSignature_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_GetSignature",
	})
	factomdprimitivesSignatureMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_MarshalBinary",
	})
	factomdprimitivesSignatureUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_UnmarshalBinaryData",
	})
	factomdprimitivesSignatureUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_UnmarshalBinary",
	})
	factomdprimitivesSignatureDetachSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_DetachSig_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_DetachSig",
	})
	factomdprimitivesDetachedSignatureString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_DetachedSignature_String_Summary",
		Help: "Summary of calls to  factomd_primitives_DetachedSignature_String",
	})
	factomdprimitivesSignatureVerify = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Signature_Verify_Summary",
		Help: "Summary of calls to  factomd_primitives_Signature_Verify",
	})
	factomdprimitivesSignSignable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_SignSignable_Summary",
		Help: "Summary of calls to  factomd_primitives_SignSignable",
	})
	factomdprimitivesSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Sign_Summary",
		Help: "Summary of calls to  factomd_primitives_Sign",
	})
	factomdprimitivesVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_primitives_VerifySignature",
	})
	factomdprimitivesGetTimeMilli = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_GetTimeMilli_Summary",
		Help: "Summary of calls to  factomd_primitives_GetTimeMilli",
	})
	factomdprimitivesGetTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_GetTime_Summary",
		Help: "Summary of calls to  factomd_primitives_GetTime",
	})
	factomdprimitivesNewTimestampNow = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewTimestampNow_Summary",
		Help: "Summary of calls to  factomd_primitives_NewTimestampNow",
	})
	factomdprimitivesNewTimestampFromSeconds = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewTimestampFromSeconds_Summary",
		Help: "Summary of calls to  factomd_primitives_NewTimestampFromSeconds",
	})
	factomdprimitivesNewTimestampFromMinutes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewTimestampFromMinutes_Summary",
		Help: "Summary of calls to  factomd_primitives_NewTimestampFromMinutes",
	})
	factomdprimitivesNewTimestampFromMilliseconds = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_NewTimestampFromMilliseconds_Summary",
		Help: "Summary of calls to  factomd_primitives_NewTimestampFromMilliseconds",
	})
	factomdprimitivesTimestampSetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_SetTimestamp_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_SetTimestamp",
	})
	factomdprimitivesTimestampSetTimeNow = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_SetTimeNow_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_SetTimeNow",
	})
	factomdprimitivesTimestampSetTimeMilli = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_SetTimeMilli_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_SetTimeMilli",
	})
	factomdprimitivesTimestampSetTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_SetTime_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_SetTime",
	})
	factomdprimitivesTimestampSetTimeSeconds = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_SetTimeSeconds_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_SetTimeSeconds",
	})
	factomdprimitivesTimestampGetTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_GetTime_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_GetTime",
	})
	factomdprimitivesTimestampUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_UnmarshalBinaryData",
	})
	factomdprimitivesTimestampUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_UnmarshalBinary",
	})
	factomdprimitivesTimestampGetTimeSeconds = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_GetTimeSeconds_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_GetTimeSeconds",
	})
	factomdprimitivesTimestampGetTimeMinutesUInt32 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_GetTimeMinutesUInt32_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_GetTimeMinutesUInt32",
	})
	factomdprimitivesTimestampGetTimeMilli = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_GetTimeMilli_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_GetTimeMilli",
	})
	factomdprimitivesTimestampGetTimeMilliUInt64 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_GetTimeMilliUInt64_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_GetTimeMilliUInt64",
	})
	factomdprimitivesTimestampGetTimeSecondsUInt32 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_GetTimeSecondsUInt32_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_GetTimeSecondsUInt32",
	})
	factomdprimitivesTimestampMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_MarshalBinary",
	})
	factomdprimitivesTimestampString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_Timestamp_String_Summary",
		Help: "Summary of calls to  factomd_primitives_Timestamp_String",
	})
	factomdprimitivesAddCommas = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_AddCommas_Summary",
		Help: "Summary of calls to  factomd_primitives_AddCommas",
	})
	factomdprimitivesWriteNumber64 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_WriteNumber64_Summary",
		Help: "Summary of calls to  factomd_primitives_WriteNumber64",
	})
	factomdprimitivesWriteNumber32 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_WriteNumber32_Summary",
		Help: "Summary of calls to  factomd_primitives_WriteNumber32",
	})
	factomdprimitivesWriteNumber16 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_WriteNumber16_Summary",
		Help: "Summary of calls to  factomd_primitives_WriteNumber16",
	})
	factomdprimitivesWriteNumber8 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_WriteNumber8_Summary",
		Help: "Summary of calls to  factomd_primitives_WriteNumber8",
	})
	factomdprimitivesConvertDecimalToFloat = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertDecimalToFloat_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertDecimalToFloat",
	})
	factomdprimitivesConvertDecimalToString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertDecimalToString_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertDecimalToString",
	})
	factomdprimitivesConvertDecimalToPaddedString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertDecimalToPaddedString_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertDecimalToPaddedString",
	})
	factomdprimitivesConvertFixedPoint = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertFixedPoint_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertFixedPoint",
	})
	factomdprimitivesConvertAddressToUser = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertAddressToUser_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertAddressToUser",
	})
	factomdprimitivesConvertFctAddressToUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertFctAddressToUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertFctAddressToUserStr",
	})
	factomdprimitivesConvertFctPrivateToUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertFctPrivateToUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertFctPrivateToUserStr",
	})
	factomdprimitivesConvertECAddressToUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertECAddressToUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertECAddressToUserStr",
	})
	factomdprimitivesConvertECPrivateToUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertECPrivateToUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertECPrivateToUserStr",
	})
	factomdprimitivesvalidateUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_validateUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_validateUserStr",
	})
	factomdprimitivesValidateFUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ValidateFUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_ValidateFUserStr",
	})
	factomdprimitivesValidateFPrivateUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ValidateFPrivateUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_ValidateFPrivateUserStr",
	})
	factomdprimitivesValidateECUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ValidateECUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_ValidateECUserStr",
	})
	factomdprimitivesValidateECPrivateUserStr = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ValidateECPrivateUserStr_Summary",
		Help: "Summary of calls to  factomd_primitives_ValidateECPrivateUserStr",
	})
	factomdprimitivesConvertUserStrToAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_ConvertUserStrToAddress_Summary",
		Help: "Summary of calls to  factomd_primitives_ConvertUserStrToAddress",
	})
	factomdprimitivesRandomVarInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_RandomVarInt_Summary",
		Help: "Summary of calls to  factomd_primitives_RandomVarInt",
	})
	factomdprimitivesVarIntLength = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_VarIntLength_Summary",
		Help: "Summary of calls to  factomd_primitives_VarIntLength",
	})
	factomdprimitivesDecodeVarInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_DecodeVarInt_Summary",
		Help: "Summary of calls to  factomd_primitives_DecodeVarInt",
	})
	factomdprimitivesEncodeVarInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_EncodeVarInt_Summary",
		Help: "Summary of calls to  factomd_primitives_EncodeVarInt",
	})
	factomdprimitivesDecodeVarIntGo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_DecodeVarIntGo_Summary",
		Help: "Summary of calls to  factomd_primitives_DecodeVarIntGo",
	})
	factomdprimitivesEncodeVarIntGo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_primitives_EncodeVarIntGo_Summary",
		Help: "Summary of calls to  factomd_primitives_EncodeVarIntGo",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdprimitivesBufferDeepCopyBytes)
	prometheus.MustRegister(factomdprimitivesNewBuffer)
	prometheus.MustRegister(factomdprimitivesAreBytesEqual)
	prometheus.MustRegister(factomdprimitivesAreBinaryMarshallablesEqual)
	prometheus.MustRegister(factomdprimitivesEncodeBinary)
	prometheus.MustRegister(factomdprimitivesDecodeBinary)
	prometheus.MustRegister(factomdprimitivesStringToByteSlice32)
	prometheus.MustRegister(factomdprimitivesByte32ToByteSlice32)
	prometheus.MustRegister(factomdprimitivesByteSlice32MarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSlice32Fixed)
	prometheus.MustRegister(factomdprimitivesByteSlice32UnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesByteSlice32UnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSlice32JSONByte)
	prometheus.MustRegister(factomdprimitivesByteSlice32JSONString)
	prometheus.MustRegister(factomdprimitivesByteSlice32String)
	prometheus.MustRegister(factomdprimitivesByteSlice32MarshalText)
	prometheus.MustRegister(factomdprimitivesByteSlice64MarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSlice64UnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesByteSlice64UnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSlice64JSONByte)
	prometheus.MustRegister(factomdprimitivesByteSlice64JSONString)
	prometheus.MustRegister(factomdprimitivesByteSlice64String)
	prometheus.MustRegister(factomdprimitivesByteSlice64MarshalText)
	prometheus.MustRegister(factomdprimitivesByteSlice6MarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSlice6UnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesByteSlice6UnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSlice6JSONByte)
	prometheus.MustRegister(factomdprimitivesByteSlice6JSONString)
	prometheus.MustRegister(factomdprimitivesByteSlice6String)
	prometheus.MustRegister(factomdprimitivesByteSlice6MarshalText)
	prometheus.MustRegister(factomdprimitivesByteSliceSigMarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSliceSigGetFixed)
	prometheus.MustRegister(factomdprimitivesByteSliceSigUnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesByteSliceSigUnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSliceSigJSONByte)
	prometheus.MustRegister(factomdprimitivesByteSliceSigJSONString)
	prometheus.MustRegister(factomdprimitivesByteSliceSigString)
	prometheus.MustRegister(factomdprimitivesByteSliceSigMarshalText)
	prometheus.MustRegister(factomdprimitivesByteSliceSigUnmarshalText)
	prometheus.MustRegister(factomdprimitivesByteSlice20MarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSlice20GetFixed)
	prometheus.MustRegister(factomdprimitivesByteSlice20UnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesByteSlice20UnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSlice20JSONByte)
	prometheus.MustRegister(factomdprimitivesByteSlice20JSONString)
	prometheus.MustRegister(factomdprimitivesByteSlice20String)
	prometheus.MustRegister(factomdprimitivesByteSlice20MarshalText)
	prometheus.MustRegister(factomdprimitivesStringToByteSlice)
	prometheus.MustRegister(factomdprimitivesByteSliceNew)
	prometheus.MustRegister(factomdprimitivesByteSliceMarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSliceUnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesByteSliceUnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesByteSliceJSONByte)
	prometheus.MustRegister(factomdprimitivesByteSliceJSONString)
	prometheus.MustRegister(factomdprimitivesByteSliceString)
	prometheus.MustRegister(factomdprimitivesByteSliceMarshalText)
	prometheus.MustRegister(factomdprimitivesMnemonicStringToPrivateKey)
	prometheus.MustRegister(factomdprimitivesMnemonicStringToPrivateKeyString)
	prometheus.MustRegister(factomdprimitivesHumanReadableFactoidPrivateKeyToPrivateKey)
	prometheus.MustRegister(factomdprimitivesHumanReadableFactoidPrivateKeyToPrivateKeyString)
	prometheus.MustRegister(factomdprimitivesHumanReadableECPrivateKeyToPrivateKey)
	prometheus.MustRegister(factomdprimitivesHumanReadableECPrivateKeyToPrivateKeyString)
	prometheus.MustRegister(factomdprimitivesPrivateKeyStringToHumanReadableFactoidPrivateKey)
	prometheus.MustRegister(factomdprimitivesPrivateKeyStringToHumanReadableECPrivateKey)
	prometheus.MustRegister(factomdprimitivesPrivateKeyStringToHumanReadablePrivateKey)
	prometheus.MustRegister(factomdprimitivesPrivateKeyStringToPublicKeyString)
	prometheus.MustRegister(factomdprimitivesPrivateKeyStringToPublicKey)
	prometheus.MustRegister(factomdprimitivesPrivateKeyToPublicKey)
	prometheus.MustRegister(factomdprimitivesGenerateKeyFromPrivateKey)
	prometheus.MustRegister(factomdprimitivesErrorError)
	prometheus.MustRegister(factomdprimitivesCreateError)
	prometheus.MustRegister(factomdprimitivesretreiveErrorParameters)
	prometheus.MustRegister(factomdprimitivesRandomHash)
	prometheus.MustRegister(factomdprimitivesHashCopy)
	prometheus.MustRegister(factomdprimitivesHashNew)
	prometheus.MustRegister(factomdprimitivesHashMarshalText)
	prometheus.MustRegister(factomdprimitivesHashIsZero)
	prometheus.MustRegister(factomdprimitivesNewShaHashFromStr)
	prometheus.MustRegister(factomdprimitivesHashUnmarshalText)
	prometheus.MustRegister(factomdprimitivesHashFixed)
	prometheus.MustRegister(factomdprimitivesHashBytes)
	prometheus.MustRegister(factomdprimitivesGetHash)
	prometheus.MustRegister(factomdprimitivesCreateHash)
	prometheus.MustRegister(factomdprimitivesHashMarshalBinary)
	prometheus.MustRegister(factomdprimitivesHashUnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesHashUnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesHashGetBytes)
	prometheus.MustRegister(factomdprimitivesHashSetBytes)
	prometheus.MustRegister(factomdprimitivesNewShaHash)
	prometheus.MustRegister(factomdprimitivesSha512Half)
	prometheus.MustRegister(factomdprimitivesHashString)
	prometheus.MustRegister(factomdprimitivesHashByteString)
	prometheus.MustRegister(factomdprimitivesHexToHash)
	prometheus.MustRegister(factomdprimitivesHashIsSameAs)
	prometheus.MustRegister(factomdprimitivesHashIsMinuteMarker)
	prometheus.MustRegister(factomdprimitivesHashJSONByte)
	prometheus.MustRegister(factomdprimitivesHashJSONString)
	prometheus.MustRegister(factomdprimitivesSha)
	prometheus.MustRegister(factomdprimitivesShad)
	prometheus.MustRegister(factomdprimitivesNewZeroHash)
	prometheus.MustRegister(factomdprimitivesNewHash)
	prometheus.MustRegister(factomdprimitivesDoubleSha)
	prometheus.MustRegister(factomdprimitivesNewShaHashFromStruct)
	prometheus.MustRegister(factomdprimitivesJSON2RequestJSONByte)
	prometheus.MustRegister(factomdprimitivesJSON2RequestJSONString)
	prometheus.MustRegister(factomdprimitivesJSON2RequestString)
	prometheus.MustRegister(factomdprimitivesNewJSON2RequestBlank)
	prometheus.MustRegister(factomdprimitivesNewJSON2Request)
	prometheus.MustRegister(factomdprimitivesParseJSON2Request)
	prometheus.MustRegister(factomdprimitivesJSON2ResponseJSONByte)
	prometheus.MustRegister(factomdprimitivesJSON2ResponseJSONString)
	prometheus.MustRegister(factomdprimitivesJSON2ResponseString)
	prometheus.MustRegister(factomdprimitivesNewJSON2Response)
	prometheus.MustRegister(factomdprimitivesJSON2ResponseAddError)
	prometheus.MustRegister(factomdprimitivesNewJSONError)
	prometheus.MustRegister(factomdprimitivesJSONErrorError)
	prometheus.MustRegister(factomdprimitivesPrivateKeyInit)
	prometheus.MustRegister(factomdprimitivesRandomPrivateKey)
	prometheus.MustRegister(factomdprimitivesPrivateKeyCustomMarshalText2)
	prometheus.MustRegister(factomdprimitivesPrivateKeyPublic)
	prometheus.MustRegister(factomdprimitivesPrivateKeyAllocateNew)
	prometheus.MustRegister(factomdprimitivesNewPrivateKeyFromHex)
	prometheus.MustRegister(factomdprimitivesNewPrivateKeyFromHexBytes)
	prometheus.MustRegister(factomdprimitivesPrivateKeySign)
	prometheus.MustRegister(factomdprimitivesPrivateKeyMarshalSign)
	prometheus.MustRegister(factomdprimitivesPrivateKeyGenerateKey)
	prometheus.MustRegister(factomdprimitivesPrivateKeyPrivateKeyString)
	prometheus.MustRegister(factomdprimitivesPrivateKeyPublicKeyString)
	prometheus.MustRegister(factomdprimitivesPublicKeyCopy)
	prometheus.MustRegister(factomdprimitivesPublicKeyFixed)
	prometheus.MustRegister(factomdprimitivesPublicKeyIsSameAs)
	prometheus.MustRegister(factomdprimitivesPublicKeyMarshalText)
	prometheus.MustRegister(factomdprimitivesPublicKeyUnmarshalText)
	prometheus.MustRegister(factomdprimitivesPublicKeyString)
	prometheus.MustRegister(factomdprimitivesPubKeyFromString)
	prometheus.MustRegister(factomdprimitivesPublicKeyVerify)
	prometheus.MustRegister(factomdprimitivesPublicKeyMarshalBinary)
	prometheus.MustRegister(factomdprimitivesPublicKeyUnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesPublicKeyUnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesVerify)
	prometheus.MustRegister(factomdprimitivesVerifySlice)
	prometheus.MustRegister(factomdprimitivesLog)
	prometheus.MustRegister(factomdprimitivesLogJSONs)
	prometheus.MustRegister(factomdprimitivesDecodeJSON)
	prometheus.MustRegister(factomdprimitivesEncodeJSON)
	prometheus.MustRegister(factomdprimitivesEncodeJSONString)
	prometheus.MustRegister(factomdprimitivesDecodeJSONString)
	prometheus.MustRegister(factomdprimitivesEncodeJSONToBuffer)
	prometheus.MustRegister(factomdprimitivesNextPowerOfTwo)
	prometheus.MustRegister(factomdprimitivesHashMerkleBranches)
	prometheus.MustRegister(factomdprimitivesComputeMerkleRoot)
	prometheus.MustRegister(factomdprimitivesBuildMerkleTreeStore)
	prometheus.MustRegister(factomdprimitivesBuildMerkleBranchForEntryHash)
	prometheus.MustRegister(factomdprimitivesBuildMerkleBranch)
	prometheus.MustRegister(factomdprimitivesSignatureInit)
	prometheus.MustRegister(factomdprimitivesSignatureGetPubBytes)
	prometheus.MustRegister(factomdprimitivesSignatureGetSigBytes)
	prometheus.MustRegister(factomdprimitivesRandomSignatureSet)
	prometheus.MustRegister(factomdprimitivesSignatureIsSameAs)
	prometheus.MustRegister(factomdprimitivesSignatureCustomMarshalText)
	prometheus.MustRegister(factomdprimitivesSignatureBytes)
	prometheus.MustRegister(factomdprimitivesSignatureSetPub)
	prometheus.MustRegister(factomdprimitivesSignatureGetKey)
	prometheus.MustRegister(factomdprimitivesSignatureSetSignature)
	prometheus.MustRegister(factomdprimitivesSignatureGetSignature)
	prometheus.MustRegister(factomdprimitivesSignatureMarshalBinary)
	prometheus.MustRegister(factomdprimitivesSignatureUnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesSignatureUnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesSignatureDetachSig)
	prometheus.MustRegister(factomdprimitivesDetachedSignatureString)
	prometheus.MustRegister(factomdprimitivesSignatureVerify)
	prometheus.MustRegister(factomdprimitivesSignSignable)
	prometheus.MustRegister(factomdprimitivesSign)
	prometheus.MustRegister(factomdprimitivesVerifySignature)
	prometheus.MustRegister(factomdprimitivesGetTimeMilli)
	prometheus.MustRegister(factomdprimitivesGetTime)
	prometheus.MustRegister(factomdprimitivesNewTimestampNow)
	prometheus.MustRegister(factomdprimitivesNewTimestampFromSeconds)
	prometheus.MustRegister(factomdprimitivesNewTimestampFromMinutes)
	prometheus.MustRegister(factomdprimitivesNewTimestampFromMilliseconds)
	prometheus.MustRegister(factomdprimitivesTimestampSetTimestamp)
	prometheus.MustRegister(factomdprimitivesTimestampSetTimeNow)
	prometheus.MustRegister(factomdprimitivesTimestampSetTimeMilli)
	prometheus.MustRegister(factomdprimitivesTimestampSetTime)
	prometheus.MustRegister(factomdprimitivesTimestampSetTimeSeconds)
	prometheus.MustRegister(factomdprimitivesTimestampGetTime)
	prometheus.MustRegister(factomdprimitivesTimestampUnmarshalBinaryData)
	prometheus.MustRegister(factomdprimitivesTimestampUnmarshalBinary)
	prometheus.MustRegister(factomdprimitivesTimestampGetTimeSeconds)
	prometheus.MustRegister(factomdprimitivesTimestampGetTimeMinutesUInt32)
	prometheus.MustRegister(factomdprimitivesTimestampGetTimeMilli)
	prometheus.MustRegister(factomdprimitivesTimestampGetTimeMilliUInt64)
	prometheus.MustRegister(factomdprimitivesTimestampGetTimeSecondsUInt32)
	prometheus.MustRegister(factomdprimitivesTimestampMarshalBinary)
	prometheus.MustRegister(factomdprimitivesTimestampString)
	prometheus.MustRegister(factomdprimitivesAddCommas)
	prometheus.MustRegister(factomdprimitivesWriteNumber64)
	prometheus.MustRegister(factomdprimitivesWriteNumber32)
	prometheus.MustRegister(factomdprimitivesWriteNumber16)
	prometheus.MustRegister(factomdprimitivesWriteNumber8)
	prometheus.MustRegister(factomdprimitivesConvertDecimalToFloat)
	prometheus.MustRegister(factomdprimitivesConvertDecimalToString)
	prometheus.MustRegister(factomdprimitivesConvertDecimalToPaddedString)
	prometheus.MustRegister(factomdprimitivesConvertFixedPoint)
	prometheus.MustRegister(factomdprimitivesConvertAddressToUser)
	prometheus.MustRegister(factomdprimitivesConvertFctAddressToUserStr)
	prometheus.MustRegister(factomdprimitivesConvertFctPrivateToUserStr)
	prometheus.MustRegister(factomdprimitivesConvertECAddressToUserStr)
	prometheus.MustRegister(factomdprimitivesConvertECPrivateToUserStr)
	prometheus.MustRegister(factomdprimitivesvalidateUserStr)
	prometheus.MustRegister(factomdprimitivesValidateFUserStr)
	prometheus.MustRegister(factomdprimitivesValidateFPrivateUserStr)
	prometheus.MustRegister(factomdprimitivesValidateECUserStr)
	prometheus.MustRegister(factomdprimitivesValidateECPrivateUserStr)
	prometheus.MustRegister(factomdprimitivesConvertUserStrToAddress)
	prometheus.MustRegister(factomdprimitivesRandomVarInt)
	prometheus.MustRegister(factomdprimitivesVarIntLength)
	prometheus.MustRegister(factomdprimitivesDecodeVarInt)
	prometheus.MustRegister(factomdprimitivesEncodeVarInt)
	prometheus.MustRegister(factomdprimitivesDecodeVarIntGo)
	prometheus.MustRegister(factomdprimitivesEncodeVarIntGo)
}
