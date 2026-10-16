package constants

const (
	// Validation error messages
	VALIDATION_NOT_NUMERIC                string = "cannot be alphabet nor contains alphanumeric for this transaction type"
	VALIDATION_NULL_MSG                   string = "cannot be null for this transaction type"
	VALIDATION_NOT_NULL_MSG               string = "should be null for this transaction type"
	VALIDATION_EQUAL_MSG                  string = "was not the expected value"
	VALIDATION_NOT_EQUAL_MSG              string = "cannot be the value"
	VALIDATION_NOT_GREATER_THAN_MSG       string = "cannot be greater than"
	VALIDATION_NOT_LESS_THAN_OR_EQUAL_MSG string = "cannot be less than or equal"
	VALIDATION_NOT_LESS_THAN              string = "cannot be less than"

	// Error throw messages
	NO_SUCH_FIELD string = "No such field named: "

	// JSON API response message fields
	ACK_MESSAGE     string = "ACK"
	NAK_MESSAGE     string = "NAK"
	BUSY_MESSAGE    string = "BUSY"
	TIMEOUT_MESSAGE string = "TO"
	READY_MESSAGE   string = "READY"
	DATA_MESSAGE    string = "MSG"

	// JSON API data response fields
	COMMAND_USED    string = "response"
	COMMAND_MESSAGE string = "message"
	COMMAND_DATA    string = "data"
	COMMAND_RESULTS string = "cmdResult"
	COMMAND_STATUS  string = "result"
	STATUS_SUCCESS  string = "success"
	STATUS_FAILED   string = "failed"
	ERROR_CODE      string = "errorCode"
	ERROR_MESSAGE   string = "errorMessage"

	GET_PARAM string = "GetParam"
	EMV       string = "emv"
)

// JSON Request parameter fields
// Note: First index of array of string should be
// the parameter field name of each object
var PARAMS = [25]string{
	"params",
	"_clerkId", "_tokenRequest", "_tokenValue", "_batch",
	"_displayOption", "_configuration", "_timeZone", "_downloadType",
	"_fileType", "_slotNum", "_file", "_configType", "_lineItemLeft", "_lineItemRight",
	"_content", "_fileName", "_header", "_prompt1", "_prompt2", "_reportOutput", "_reportType",
	"_line1", "_line2", "_timeOut",
}

var TRANSACTION = [13]string{
	"transaction",
	"_amount",
	"_referenceNumber", "_authorizedAmount",
	"_baseAmount", "_taxAmount", "_tipAmount", "_taxIndicator",
	"_cashBackAmount", "_invoiceNbr", "_allowPartialAuth", "_tranNo", "_totalAmount",
}
