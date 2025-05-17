package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"sales_analysis_service/common"

	"runtime"
)

type ErrResponseStruct struct {
	Status       common.ResponseStatus
	ErrorCode    string
	ErrorMessage string
}

func ErrResponse(pErr error, pErrCode string) (lRespData string) {
	var lErrResponse ErrResponseStruct
	lErrResponse.Status = common.SUCCESS
	lErrResponse.ErrorCode = pErrCode
	lErrResponse.ErrorMessage = pErr.Error()

	lData, lErr := json.Marshal(lErrResponse)
	if lErr != nil {
		log.Fatalf("(ERROR) %v: %v", "LGES01", lErr.Error())
	}
	return string(lData)
}

func Err(pErr error) {
	if pErr == nil {
		return
	}
	var lFuncName string
	lPc, _, lLine, lOk := runtime.Caller(1)
	lDetails := runtime.FuncForPC(lPc)
	if lOk && lDetails != nil {
		lFuncName = lDetails.Name()
	}
	log.Printf("(ERROR) %s:%d - %v\n", lFuncName, lLine, pErr)
	fmt.Printf("(ERROR) %s:%d - %v\n", lFuncName, lLine, pErr)
}

func Info(pMsg ...any) {
	log.Println(pMsg...)
}
