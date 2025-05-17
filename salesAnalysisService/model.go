package service

import (
	"errors"
)

type ReqType string

const (
	TOTAL       ReqType = "REV01"
	BY_PRODUCT  ReqType = "REV02"
	BY_CATEGORY ReqType = "REV03"
	BY_REGION   ReqType = "REV04"
)

type RequestStruct struct {
	ReqType   string `json:"reqType"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type Revenvue struct {
	TotalRevenue       float64           `gorm:"column:Revenue" json:"totalRevenue,omitempty"`
	ProductRevenueArr  []ProductRevenue  `json:"productRevenueArr,omitempty"`
	CategoryRevenueArr []CategoryRevenue `json:"categoryRevenueArr,omitempty"`
	RegionRevenueArr   []RegionRevenue   `json:"regionRevenueArr,omitempty"`
	Status             string            `json:"status,omitempty"`
}

type ProductRevenue struct {
	ProductName    string  `gorm:"column:ProductName" json:"productName"`
	ProductRevenue float64 `gorm:"column:Revenue" json:"productRevenue"`
}

type CategoryRevenue struct {
	CategoryName    string  `gorm:"column:CategoryName" json:"categoryName"`
	CategoryRevenue float64 `gorm:"column:Revenue" json:"categoryRevenue"`
}

type RegionRevenue struct {
	RegionName      string  `gorm:"column:RegionName" json:"regionName"`
	RegionalRevenue float64 `gorm:"column:Revenue" json:"regionalRevenue"`
}

func (r *RequestStruct) Validate() (lErr error) {
	if r.StartDate == "" {
		lErr = errors.New("(Error) sartDate is required")
		return
	}
	if r.EndDate == "" {
		lErr = errors.New("(Error) endDate is required")
		return
	}
	if r.ReqType == "" {
		lErr = errors.New("(Error) reqType is required")
		return
	}
	return
}
