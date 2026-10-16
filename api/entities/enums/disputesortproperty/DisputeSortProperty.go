package disputesortproperty

import "github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"

type DisputeSortProperty string

const (
	Id                        DisputeSortProperty = "id"
	ARN                       DisputeSortProperty = "arn"
	Brand                     DisputeSortProperty = "brand"
	Status                    DisputeSortProperty = "status"
	Stage                     DisputeSortProperty = "stage"
	FromStageTimeCreated      DisputeSortProperty = "from_stage_time_created"
	ToStageTimeCreated        DisputeSortProperty = "to_stage_time_created"
	AdjustmentFunding         DisputeSortProperty = "adjustment_funding"
	FromAdjustmentTimeCreated DisputeSortProperty = "from_adjustment_time_created"
	ToAdjustmentTimeCreated   DisputeSortProperty = "to_adjustment_time_created"
)

func (d DisputeSortProperty) GetBytes() []byte {
	return []byte(d)
}

func (d DisputeSortProperty) GetValue() string {
	return string(d)
}

func (d DisputeSortProperty) StringConstants() []istringconstant.IStringConstant {
	return []istringconstant.IStringConstant{
		Id,
		ARN,
		Brand,
		Status,
		Stage,
		FromStageTimeCreated,
		ToStageTimeCreated,
		AdjustmentFunding,
		FromAdjustmentTimeCreated,
		ToAdjustmentTimeCreated,
	}
}
