package entities

import (
	"github.com/globalpayments/go-sdk/api/entities/enums/host"
)

type PriorMessageInformation struct {
	ResponseTime                string
	CardType                    string
	FunctionCode                string
	ProcessingCode              string
	MessageReasonCode           string
	MessageTransactionIndicator string
	SystemTraceAuditNumber      string
	ProcessingHost              host.Host
}

func (pmi *PriorMessageInformation) GetResponseTime() string {
	return pmi.ResponseTime
}

func (pmi *PriorMessageInformation) SetResponseTime(responseTime string) {
	pmi.ResponseTime = responseTime
}

func (pmi *PriorMessageInformation) GetCardType() string {
	return pmi.CardType
}

func (pmi *PriorMessageInformation) SetCardType(cardType string) {
	pmi.CardType = cardType
}

func (pmi *PriorMessageInformation) GetFunctionCode() string {
	return pmi.FunctionCode
}

func (pmi *PriorMessageInformation) SetFunctionCode(functionCode string) {
	pmi.FunctionCode = functionCode
}

func (pmi *PriorMessageInformation) GetMessageReasonCode() string {
	return pmi.MessageReasonCode
}

func (pmi *PriorMessageInformation) SetMessageReasonCode(messageReasonCode string) {
	pmi.MessageReasonCode = messageReasonCode
}

func (pmi *PriorMessageInformation) GetMessageTransactionIndicator() string {
	return pmi.MessageTransactionIndicator
}

func (pmi *PriorMessageInformation) SetMessageTransactionIndicator(messageTransactionIndicator string) {
	pmi.MessageTransactionIndicator = messageTransactionIndicator
}

func (pmi *PriorMessageInformation) GetProcessingCode() string {
	return pmi.ProcessingCode
}

func (pmi *PriorMessageInformation) SetProcessingCode(processingCode string) {
	pmi.ProcessingCode = processingCode
}

func (pmi *PriorMessageInformation) GetSystemTraceAuditNumber() string {
	return pmi.SystemTraceAuditNumber
}

func (pmi *PriorMessageInformation) SetSystemTraceAuditNumber(systemTraceAuditNumber string) {
	pmi.SystemTraceAuditNumber = systemTraceAuditNumber
}

func (pmi *PriorMessageInformation) GetProcessingHost() host.Host {
	return pmi.ProcessingHost
}

func (pmi *PriorMessageInformation) SetProcessingHost(processingHost host.Host) {
	pmi.ProcessingHost = processingHost
}
