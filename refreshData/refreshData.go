package refreshdata

import (
	"encoding/csv"
	"os"
	"sales_analysis_service/common"
	"sales_analysis_service/db"
	"sales_analysis_service/logger"
	"strconv"
	"strings"
)

func GetCsvData(pFileName string) (lErr error) {
	logger.Info("GetCsvData (+)")
	var lFileDataArr [][]string
	//
	var lCustomerDataMap = make(map[string]CustomerDetailsStruct)
	var lProductDataMap = make(map[string]ProductDetailsStruct)
	var lOrderDataMap = make(map[string]OrderDetailsStruct)
	//
	var lExistingCustomerMap = make(map[string]CustomerDetailsStruct)
	var lExistingProductMap = make(map[string]ProductDetailsStruct)
	var lExistingOrderMap = make(map[string]OrderDetailsStruct)

	// READ CSV FILE
	if lFileDataArr, lErr = GetCsvFileData(pFileName); lErr != nil {
		logger.Err(lErr)
		return
	}
	// GET CURRENT FILE DATA FOR EACH TABLE
	if lCustomerDataMap, lProductDataMap, lOrderDataMap, lErr = MapFileDataToStruct(lFileDataArr); lErr != nil {
		logger.Err(lErr)
		return
	}
	// GET EXISTING DATA FROM RESPECTIVE DATABASE TABLES
	if lExistingCustomerMap, lErr = GetExistingCustomers(); lErr != nil {
		logger.Err(lErr)
	}
	if lExistingProductMap, lErr = GetExistingProducts(); lErr != nil {
		logger.Err(lErr)
	}
	if lExistingOrderMap, lErr = GetExistingOrders(); lErr != nil {
		logger.Err(lErr)
	}
	// CHECK FOR DUPLICATE DATA & INSERT THE NEW DATA
	UpdateNewCustomers(lExistingCustomerMap, lCustomerDataMap)
	UpdateNewProducts(lExistingProductMap, lProductDataMap)
	UpdateNewOrders(lExistingOrderMap, lOrderDataMap)
	logger.Info("GetCsvData (-)")
	return
}

func GetCsvFileData(pFileName string) (lFileDataArr [][]string, lErr error) {
	logger.Info("GetCsvFileData (+)")
	var lFile *os.File
	var lCsvReader *csv.Reader

	lFile, lErr = os.Open(pFileName)
	if lErr != nil {
		logger.Err(lErr)
		return
	}
	defer lFile.Close()

	lCsvReader = csv.NewReader(lFile)
	lFileDataArr, lErr = lCsvReader.ReadAll()
	if lErr != nil {
		logger.Err(lErr)
		return
	}
	logger.Info("GetCsvFileData (-)")
	return
}

func MapFileDataToStruct(lFileDataArr [][]string) (lCustomerDataMap map[string]CustomerDetailsStruct, lProductDataMap map[string]ProductDetailsStruct, lOrderDataMap map[string]OrderDetailsStruct, lErr error) {
	logger.Info("MapFileDataToStruct (+)")
	lCustomerDataMap = make(map[string]CustomerDetailsStruct)
	lProductDataMap = make(map[string]ProductDetailsStruct)
	lOrderDataMap = make(map[string]OrderDetailsStruct)

	for lIndex, lRecordArr := range lFileDataArr {
		if lIndex == 0 || len(lRecordArr) < 14 {
			continue
		}

		var lOrderData OrderDetailsStruct
		var lCustomerData CustomerDetailsStruct
		var lProductData ProductDetailsStruct

		lCustomerData = GetCustomerDetails(lRecordArr)
		lCustomerDataMap[lCustomerData.CustomerID] = lCustomerData
		if lOrderData, lErr = GetOrderDetails(lRecordArr); lErr != nil {
			logger.Err(lErr)
			return
		}
		lOrderDataMap[lOrderData.OrderID] = lOrderData
		if lProductData, lErr = GetProductDetails(lRecordArr); lErr != nil {
			logger.Err(lErr)
			return
		}
		lProductDataMap[lProductData.ProductID] = lProductData
	}
	logger.Info("MapFileDataToStruct (-)")
	return
}

func GetOrderDetails(lRecordArr []string) (lOrderData OrderDetailsStruct, lErr error) {
	logger.Info("GetOrderDetails (+)")
	lOrderData.OrderID = lRecordArr[0]
	lOrderData.ProductId = lRecordArr[1]
	lOrderData.CustomerId = lRecordArr[2]
	lOrderData.Category = lRecordArr[4]
	lOrderData.Region = lRecordArr[5]
	lOrderData.DateOfSale = lRecordArr[6]
	if lOrderData.QuantitySold, lErr = strconv.Atoi(strings.TrimSpace(lRecordArr[7])); lErr != nil {
		logger.Err(lErr)
		return
	}
	if lOrderData.ShippingCost, lErr = strconv.ParseFloat(strings.TrimSpace(lRecordArr[10]), 64); lErr != nil {
		logger.Err(lErr)
		return
	}
	lOrderData.PaymentMethod = lRecordArr[11]
	logger.Info("GetOrderDetails (-)")
	return
}

func GetCustomerDetails(lRecordArr []string) (lCustomerData CustomerDetailsStruct) {
	logger.Info("GetCustomerDetails (+)")
	lCustomerData.CustomerID = lRecordArr[2]
	lCustomerData.CustomerName = lRecordArr[12]
	lCustomerData.CustomerEmail = lRecordArr[13]
	lCustomerData.CustomerAddress = lRecordArr[14]
	logger.Info("GetCustomerDetails (-)")
	return
}

func GetProductDetails(lRecordArr []string) (lProductData ProductDetailsStruct, lErr error) {
	lProductData.ProductID = lRecordArr[1]
	lProductData.ProductName = lRecordArr[3]
	if lProductData.UnitPrice, lErr = strconv.ParseFloat(strings.TrimSpace(lRecordArr[8]), 64); lErr != nil {
		logger.Err(lErr)
		return
	}
	if lProductData.Discount, lErr = strconv.ParseFloat(strings.TrimSpace(lRecordArr[9]), 64); lErr != nil {
		logger.Err(lErr)
		return
	}
	return
}

func GetExistingCustomers() (lCustomerMap map[string]CustomerDetailsStruct, lErr error) {
	var lCustomerArr []CustomerDetailsStruct
	lCustomerMap = make(map[string]CustomerDetailsStruct)
	if lErr = db.GMysqlGormDB.Table("CustomerDetails").Find(&lCustomerArr).Error; lErr != nil {
		logger.Err(lErr)
		return
	}
	for _, lCustomer := range lCustomerArr {
		lCustomerMap[lCustomer.CustomerID] = lCustomer
	}
	return
}

func GetExistingProducts() (lProductMap map[string]ProductDetailsStruct, lErr error) {
	var lProductArr []ProductDetailsStruct
	lProductMap = make(map[string]ProductDetailsStruct)
	if lErr = db.GMysqlGormDB.Table("ProductDetails").Find(&lProductArr).Error; lErr != nil {
		logger.Err(lErr)
		return
	}
	for _, lProduct := range lProductArr {
		lProductMap[lProduct.ProductID] = lProduct
	}
	return
}

func GetExistingOrders() (lOrderMap map[string]OrderDetailsStruct, lErr error) {
	var lOrderArr []OrderDetailsStruct
	lOrderMap = make(map[string]OrderDetailsStruct)
	if lErr = db.GMysqlGormDB.Table("OrderDetails").Find(&lOrderArr).Error; lErr != nil {
		logger.Err(lErr)
		return
	}
	for _, lOrder := range lOrderArr {
		lOrderMap[lOrder.OrderID] = lOrder
	}
	return
}

func UpdateNewCustomers(pExistingData, pNewData map[string]CustomerDetailsStruct) {
	var lFilteredData []CustomerDetailsStruct
	for _, lCustomer := range pNewData {
		if _, lExist := pExistingData[lCustomer.CustomerID]; !lExist {
			lCustomer.CreatedDate = common.GetCurrentDateTime()
			lCustomer.UpdatedDate = common.GetCurrentDateTime()
			lFilteredData = append(lFilteredData, lCustomer)
		}
	}
	if lErr := db.GMysqlGormDB.Table("CustomerDetails").CreateInBatches(&lFilteredData, 1000).Error; lErr != nil {
		logger.Err(lErr)
		return
	}
}

func UpdateNewProducts(pExistingData, pNewData map[string]ProductDetailsStruct) {
	var lFilteredData []ProductDetailsStruct
	for _, lProduct := range pNewData {
		if _, lExist := pExistingData[lProduct.ProductID]; !lExist {
			lProduct.CreatedDate = common.GetCurrentDateTime()
			lProduct.UpdatedDate = common.GetCurrentDateTime()
			lFilteredData = append(lFilteredData, lProduct)
		}
	}
	if lErr := db.GMysqlGormDB.Table("ProductDetails").CreateInBatches(&lFilteredData, 1000).Error; lErr != nil {
		logger.Err(lErr)
		return
	}
}

func UpdateNewOrders(pExistingData, pNewData map[string]OrderDetailsStruct) {
	var lFilteredData []OrderDetailsStruct
	for _, lOrder := range pNewData {
		if _, lExist := pExistingData[lOrder.OrderID]; !lExist {
			lOrder.CreatedDate = common.GetCurrentDateTime()
			lOrder.UpdatedDate = common.GetCurrentDateTime()
			lOrder.CustomerKeyId = GetCustomerKeyId(lOrder.CustomerId)
			lOrder.ProductKeyId = GetProductKeyId(lOrder.ProductId)
			lFilteredData = append(lFilteredData, lOrder)
		}
	}
	if lErr := db.GMysqlGormDB.Table("OrderDetails").CreateInBatches(&lFilteredData, 1000).Error; lErr != nil {
		logger.Err(lErr)
		return
	}
}

func GetCustomerKeyId(pCustomerId string) (lId int) {
	if lErr := db.GMysqlGormDB.Table("CustomerDetails").Select("id").
		Where("customer_id = ?", pCustomerId).Scan(&lId).Error; lErr != nil {
		logger.Err(lErr)
		return
	}
	return
}

func GetProductKeyId(pProductId string) (lId int) {
	if lErr := db.GMysqlGormDB.Table("ProductDetails").Select("id").
		Where("product_id = ?", pProductId).Scan(&lId).Error; lErr != nil {
		logger.Err(lErr)
		return
	}
	return
}
