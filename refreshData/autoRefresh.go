package refreshdata

import (
	"fmt"
	"net/http"
	"sales_analysis_service/logger"
	"sales_analysis_service/readtoml"
	"strconv"
	"time"
)

func AutoRefreshDataApi(w http.ResponseWriter, r *http.Request) {
	logger.Info("AutoRefreshDataApi (+) ")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")
	var lErr error
	if r.Method == http.MethodGet {
		if lErr = GetCsvData("./salesData.csv"); lErr != nil {
			logger.Err(lErr)
			fmt.Fprint(w, logger.ErrResponse(lErr, "RARDA01"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	logger.Info("AutoRefreshDataApi (-) ")
}

func AutoRefreshData() {
	var lErr error
	var lExecutionFrequency int
	for {
		logger.Info("AutoRefreshData (+)")
		// time.Sleep(1 * time.Minute)
		lTimeConfig := readtoml.ReadToml("./toml/serviceconfig.toml")
		lFrequency := readtoml.GetConfigValue(lTimeConfig, "FREQUENCY")
		//
		if lExecutionFrequency, lErr = strconv.Atoi(lFrequency); lErr != nil {
			logger.Err(lErr)
			lExecutionFrequency = 6 // DEFAULT - RUNS EVERY 6 HOURS
		}
		if lErr = GetCsvData("./salesData.csv"); lErr != nil {
			logger.Err(lErr)
		}
		time.Sleep(time.Duration(lExecutionFrequency) * time.Hour)
		logger.Info("AutoRefreshData (-)")
	}
}
