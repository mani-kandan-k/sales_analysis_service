package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sales_analysis_service/common"
	"sales_analysis_service/db"
	"sales_analysis_service/logger"
)

func GetRevenueApi(w http.ResponseWriter, r *http.Request) {
	logger.Info("GetRevenueApi (+) ")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	if r.Method == http.MethodPost {

		var lRequest RequestStruct
		var lRespone Revenvue
		lRespone.Status = string(common.ERROR)

		lErr := json.NewDecoder(r.Body).Decode(&lRequest)
		if lErr != nil {
			logger.Err(lErr)
			fmt.Fprint(w, logger.ErrResponse(lErr, "RGRA01"))
			return
		}
		if lErr := lRequest.Validate(); lErr != nil {
			logger.Err(lErr)
			fmt.Fprint(w, logger.ErrResponse(lErr, "RGRA02"))
			return
		}

		if lRespone, lErr = GetRevenue(lRequest); lErr != nil {
			logger.Err(lErr)
			fmt.Fprint(w, logger.ErrResponse(lErr, "RGRA03"))
			return
		}

		lData, lErr := json.Marshal(lRespone)
		if lErr != nil {
			logger.Err(lErr)
			fmt.Fprint(w, logger.ErrResponse(lErr, "RGRA04"))
			return
		}
		fmt.Fprint(w, string(lData))
	}
	logger.Info("GetRevenueApi (-) ")
}

func GetRevenue(pRequest RequestStruct) (lRespone Revenvue, lErr error) {
	logger.Info("GetRevenue (+) ")
	lGormDB := db.GMysqlGormDB.Debug().Table("OrderDetails o").
		Joins("JOIN ProductDetails p ON o.product_id = p.id").
		Where("DATE(o.date_of_sale) BETWEEN ? AND ?", pRequest.StartDate, pRequest.EndDate)
	lSelect := `SUM((o.quantity_sold * p.unit_price * (1 - p.discount)) - o.shipping_cost) AS Revenue`
	switch ReqType(pRequest.ReqType) {
	case TOTAL:
		lErr = lGormDB.Select(lSelect).Scan(&lRespone.TotalRevenue).Error
	case BY_PRODUCT:
		lErr = lGormDB.Select(lSelect, `p.product_name AS ProductName`).Group("o.product_id, p.product_name").
			Scan(&lRespone.ProductRevenueArr).Error
	case BY_CATEGORY:
		lErr = lGormDB.Select(lSelect, `o.category AS CategoryName`).Group("o.category").
			Scan(&lRespone.CategoryRevenueArr).Error
	case BY_REGION:
		lErr = lGormDB.Select(lSelect, `o.region AS RegionName`).Group("o.region").
			Scan(&lRespone.RegionRevenueArr).Error
	}
	if lErr != nil {
		logger.Err(lErr)
		return
	}
	logger.Info("GetRevenue (-) ")
	return
}
