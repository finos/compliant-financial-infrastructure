package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var funcMap = map[string]interface{}{
	"GCS_CCC_OS_C1_TR02": GCS_CCC_OS_C1_TR02,
	"GCS_CCC_OS_C1_TR03": GCS_CCC_OS_C1_TR03,
	"GCS_CCC_OS_C1_TR01": GCS_CCC_OS_C1_TR01,
	"GCS_CCC_OS_C2_TR01": GCS_CCC_OS_C2_TR01,
	"GCS_CCC_OS_C2_TR02": GCS_CCC_OS_C2_TR02,
	"GCS_CCC_OS_C2_TR03": GCS_CCC_OS_C2_TR03,
	"GCS_CCC_OS_C3_TR01": GCS_CCC_OS_C3_TR01,
	"GCS_CCC_OS_C3_TR02": GCS_CCC_OS_C3_TR02,
	"GCS_CCC_OS_C3_TR03": GCS_CCC_OS_C3_TR03,
	"GCS_CCC_OS_C4_TR01": GCS_CCC_OS_C4_TR01,
	"GCS_CCC_OS_C4_TR02": GCS_CCC_OS_C4_TR02,
	"GCS_CCC_OS_C4_TR03": GCS_CCC_OS_C4_TR03,
	"GCS_CCC_OS_C5_TR01": GCS_CCC_OS_C5_TR01,
	"GCS_CCC_OS_C5_TR02": GCS_CCC_OS_C5_TR02,
	"GCS_CCC_OS_C5_TR03": GCS_CCC_OS_C5_TR03,
	
}

// Pass takes any number of strings and prints them as a single string
func CFIPass(values ...string) {
	cfiPrint(values, "PASS")
}

// FAIL takes any number of strings and prints them as a single string
func CFIFAIL(values ...string) {
	cfiPrint(values, "FAIL")
}

// Error takes any number of strings and prints them as a single string
func CFIError(values ...string) {
	cfiPrint(values, "ERROR")
}

// cfiPrint prints the provided values as a single string with a prefix
func cfiPrint(values []string, prefix string) {
	str := strings.Join(values, " ")
	fmt.Printf("[%s]: %s", prefix, str)
}

func main() {
	if len(os.Args) == 1 {
		CFIError("Please provide a function name to call")
		return
	}

	funcName := os.Args[1]
	if function, exists := funcMap[funcName]; exists {
		reflect.ValueOf(function).Call(nil)
	} else {
		CFIError("Function does not exist")
		// panic and print function name if it does not exist
		CFIError("Function %s does not exist", funcName)
	}
}
