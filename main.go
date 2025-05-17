package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sales_analysis_service/db"
	"sales_analysis_service/logger"
	"sales_analysis_service/readtoml"
	refreshdata "sales_analysis_service/refreshData"
	service "sales_analysis_service/salesAnalysisService"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server Started...")
	var lFile *os.File
	var lErr error
	// Logging
	if lFile, lErr = GetLogFile(); lErr != nil {
		log.Fatalf("(ERROR) On Opening Logfile: %v", lErr)
	}
	defer lFile.Close()
	log.SetOutput(lFile)

	if lErr = db.DataBaseInit(db.GetDbConfig()); lErr != nil {
		log.Fatalf("(ERROR) While Initializing DB Connection: %v", lErr)
	}
	defer db.CloseDbConnection()

	// Periodic Data Refresh Mechanism
	go refreshdata.AutoRefreshData()

	lSalesDataService := &http.Server{
		Handler: GetRouters(),
		Addr:    GetPortNo(),
	}
	log.Fatalf("Server Terminated..!!! : %v", lSalesDataService.ListenAndServe())
}

func GetRouters() (lRouter *mux.Router) {
	logger.Info("GetRouters (+)")
	lRouter = mux.NewRouter()
	// Sales Revenue API
	lRouter.HandleFunc("/getSalesRevenue", service.GetRevenueApi).Methods(http.MethodPost)
	// On-Demand Data Refresh Mechanism
	lRouter.HandleFunc("/refreshData", refreshdata.AutoRefreshDataApi).Methods(http.MethodGet)
	logger.Info("GetRouters (-)")
	return
}

func GetPortNo() string {
	return ":" + fmt.Sprintf("%v", readtoml.ReadToml("./toml/serviceconfig.toml").(map[string]any)["PORT"])
}

func GetLogFile() (lFile *os.File, lErr error) {
	logger.Info("GetLogFile (+)")
	lFile, lErr = os.OpenFile("./log/logfile"+time.Now().Format(time.StampMilli)+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		logger.Err(lErr)
		return
	}
	logger.Info("GetLogFile (-)")
	return
}
