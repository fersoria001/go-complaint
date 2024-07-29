package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
)

type ComplaintsInfo struct {
	Received []*ComplaintData `json:"received"`
	Resolved []*ComplaintData `json:"resolved"`
	Reviewed []*ComplaintData `json:"reviewed"`
	Sent     []*ComplaintData `json:"sent"`
}

type ComplaintData struct {
	Id          string `json:"id"`
	OwnerId     string `json:"ownerId"`
	ComplaintId string `json:"complaintId"`
	OccurredOn  string `json:"occurredOn"`
	DataType    string `json:"dataType"`
}

func NewComplaintData(obj complaint.ComplaintData) *ComplaintData {

	return &ComplaintData{
		Id:          obj.Id().String(),
		OwnerId:     obj.OwnerId().String(),
		ComplaintId: obj.ComplaintId().String(),
		OccurredOn:  common.StringDate(obj.OccurredOn()),
		DataType:    obj.DataType().String(),
	}
}
