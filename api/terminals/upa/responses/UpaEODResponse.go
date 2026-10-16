package responses

import (
	"github.com/globalpayments/go-sdk/api/terminals/upa/entities/enums/upamessageid"
	"github.com/globalpayments/go-sdk/api/utils"
)

type UpaEODResponse struct {
	BatchId *int
	UpaDeviceResponse
}

func NewUpaEODResponse(responseObj utils.JsonDoc) *UpaEODResponse {

	eodResponse := &UpaEODResponse{
		UpaDeviceResponse: *NewUpaDeviceResponse(responseObj, upamessageid.EODProcessing),
	}

	outerData := responseObj.Get("data")

	if outerData != nil {
		innerData := outerData.Get("data")
		if innerData != nil {
			host := innerData.Get("host")
			if host != nil {
				eodResponse.BatchId = host.GetInt("batchId")
			}
		}
	}

	return eodResponse
}

func (r *UpaEODResponse) GetBatchId() *int {
	return r.BatchId
}
