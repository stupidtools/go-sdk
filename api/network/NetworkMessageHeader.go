package network

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/ibyteconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
	"github.com/globalpayments/go-sdk/api/network/enums/characterset"
	"github.com/globalpayments/go-sdk/api/network/enums/connectiontype"
	"github.com/globalpayments/go-sdk/api/network/enums/messagetype"
	"github.com/globalpayments/go-sdk/api/network/enums/networkprocessingflag"
	"github.com/globalpayments/go-sdk/api/network/enums/networkresponsecode"
	"github.com/globalpayments/go-sdk/api/network/enums/networkresponsecodeorigin"
	"github.com/globalpayments/go-sdk/api/network/enums/networktransactiontype"
	"github.com/globalpayments/go-sdk/api/network/enums/protocoltype"
	"github.com/globalpayments/go-sdk/api/utils/stringutils"
)

type NetworkMessageHeader struct {
	networkTransactionType networktransactiontype.NetworkTransactionType
	messageType            messagetype.MessageType
	characterSet           characterset.CharacterSet
	responseCode           networkresponsecode.NetworkResponseCode
	responseCodeOrigin     networkresponsecodeorigin.NetworkResponseCodeOrigin
	processingFlag         networkprocessingflag.NetworkProcessingFlag
	protocolType           protocoltype.ProtocolType
	connectionType         connectiontype.ConnectionType
	nodeIdentification     string
	originCorrelation1     []byte
	companyId              string
	originCorrelation2     []byte
	version                byte
}

func (header *NetworkMessageHeader) GetNetworkTransactionType() networktransactiontype.NetworkTransactionType {
	return header.networkTransactionType
}

func (header *NetworkMessageHeader) SetNetworkTransactionType(networkTransactionType networktransactiontype.NetworkTransactionType) {
	header.networkTransactionType = networkTransactionType
}

func (header *NetworkMessageHeader) GetMessageType() messagetype.MessageType {
	return header.messageType
}

func (header *NetworkMessageHeader) SetMessageType(messageType messagetype.MessageType) {
	header.messageType = messageType
}

func (header *NetworkMessageHeader) GetCharacterSet() characterset.CharacterSet {
	return header.characterSet
}

func (header *NetworkMessageHeader) SetCharacterSet(characterSet characterset.CharacterSet) {
	header.characterSet = characterSet
}

func (header *NetworkMessageHeader) GetResponseCode() networkresponsecode.NetworkResponseCode {
	return header.responseCode
}

func (header *NetworkMessageHeader) SetResponseCode(responseCode networkresponsecode.NetworkResponseCode) {
	header.responseCode = responseCode
}

func (header *NetworkMessageHeader) GetResponseCodeOrigin() networkresponsecodeorigin.NetworkResponseCodeOrigin {
	return header.responseCodeOrigin
}

func (header *NetworkMessageHeader) SetResponseCodeOrigin(responseCodeOrigin networkresponsecodeorigin.NetworkResponseCodeOrigin) {
	header.responseCodeOrigin = responseCodeOrigin
}

func (header *NetworkMessageHeader) GetProcessingFlag() networkprocessingflag.NetworkProcessingFlag {
	return header.processingFlag
}

func (header *NetworkMessageHeader) SetProcessingFlag(processingFlag networkprocessingflag.NetworkProcessingFlag) {
	header.processingFlag = processingFlag
}

func (header *NetworkMessageHeader) GetProtocolType() protocoltype.ProtocolType {
	return header.protocolType
}

func (header *NetworkMessageHeader) SetProtocolType(protocolType protocoltype.ProtocolType) {
	header.protocolType = protocolType
}

func (header *NetworkMessageHeader) GetConnectionType() connectiontype.ConnectionType {
	return header.connectionType
}

func (header *NetworkMessageHeader) SetConnectionType(connectionType connectiontype.ConnectionType) {
	header.connectionType = connectionType
}

func (header *NetworkMessageHeader) GetNodeIdentification() string {
	return header.nodeIdentification
}

func (header *NetworkMessageHeader) SetNodeIdentification(nodeIdentification string) {
	header.nodeIdentification = nodeIdentification
}

func (header *NetworkMessageHeader) GetOriginCorrelation1() []byte {
	return header.originCorrelation1
}

func (header *NetworkMessageHeader) SetOriginCorrelation1(originCorrelation1 []byte) {
	header.originCorrelation1 = originCorrelation1
}

func (header *NetworkMessageHeader) GetCompanyId() string {
	return header.companyId
}

func (header *NetworkMessageHeader) SetCompanyId(companyId string) {
	header.companyId = companyId
}

func (header *NetworkMessageHeader) GetOriginCorrelation2() []byte {
	return header.originCorrelation2
}

func (header *NetworkMessageHeader) SetOriginCorrelation2(originCorrelation2 []byte) {
	header.originCorrelation2 = originCorrelation2
}

func (header *NetworkMessageHeader) GetVersion() byte {
	return header.version
}

func (header *NetworkMessageHeader) SetVersion(version byte) {
	header.version = version
}

func Parse(buffer []byte) (*NetworkMessageHeader, error) {
	sp := stringutils.NewStringParserFromBytes(buffer)
	header := &NetworkMessageHeader{}
	var sc istringconstant.IStringConstant
	var bc ibyteconstant.IByteConstant

	sc = sp.ReadStringConstant(2, networktransactiontype.KeepAlive)
	if sc != nil {
		header.networkTransactionType = sc.(networktransactiontype.NetworkTransactionType)
	}

	bc = sp.ReadByteConstant(messagetype.WEX_Processing)
	if bc != nil {
		header.messageType = bc.(messagetype.MessageType)
	}

	bc = sp.ReadByteConstant(characterset.ASCII)
	if bc != nil {
		header.characterSet = bc.(characterset.CharacterSet)
	}

	bc = sp.ReadByteConstant(networkresponsecode.FailedConnection)
	if bc != nil {
		header.responseCode = bc.(networkresponsecode.NetworkResponseCode)
	}

	bc = sp.ReadByteConstant(networkresponsecodeorigin.AuthorizationHost)
	if bc != nil {
		header.responseCodeOrigin = bc.(networkresponsecodeorigin.NetworkResponseCodeOrigin)
	}

	bc = sp.ReadByteConstant(networkprocessingflag.NonPersistentConnection)
	if bc != nil {
		header.processingFlag = bc.(networkprocessingflag.NetworkProcessingFlag)
	}

	protocolPointer := sp.ReadSingleByte()
	if protocolPointer != nil {
		protocolBytes := *protocolPointer
		if protocolBytes == 0x07 {
			protocolBytes = 0x05
		}
		protocolType := stringutils.ReverseByteEnumMapParse(protocolBytes, protocoltype.Async)
		if protocolType != nil {
			header.protocolType = protocolType.(protocoltype.ProtocolType)
		}
	}

	bc = sp.ReadByteConstant(connectiontype.Accel)
	if bc != nil {
		header.connectionType = bc.(connectiontype.ConnectionType)
	}

	header.nodeIdentification = sp.ReadString(4)
	header.originCorrelation1 = sp.ReadBytes(2)
	header.companyId = sp.ReadString(4)
	header.originCorrelation2 = sp.ReadBytes(8)
	header.version = *sp.ReadSingleByte()

	return header, nil
}
