package main

import "errors"

// ------------------------------
// Test Requirements
// ------------------------------

// GCS_CCC_OS_C1_TR01: All supported network data protocols must be running on secure channels
func GCS_CCC_OS_C1_TR01() {
	err := GCS_CCC_OS_C1_TR01_T01()
	if err != nil {
		CFIError(err.Error())
	}
	err = GCS_CCC_OS_C1_TR01_T02()
	if err != nil {
		CFIError(err.Error())
	}
	err = GCS_CCC_OS_C1_TR01_T03()
	if err != nil {
		CFIError(err.Error())
	}
	CFIPass("All supported network data protocols are running on secure channels")
}

// GCS_CCC_OS_C1_TR02: All clear text channels should be disabled
func GCS_CCC_OS_C1_TR02() {
	CFIError("Not Yet Implemented")
}

// GCS_CCC_OS_C1_TR03: The cipher suite implemented should conform with the latest suggested cipher suites
func GCS_CCC_OS_C1_TR03() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C2_TR01() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C2_TR02() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C2_TR03() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C3_TR01() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C3_TR02() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C3_TR03() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C4_TR01() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C4_TR02() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C4_TR03() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C5_TR01() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C5_TR02() {
	CFIError("Not Yet Implemented")
}

func GCS_CCC_OS_C5_TR03() {
	CFIError("Not Yet Implemented")
}

// ------------------------------
// Tests
// ------------------------------

// GCS_CCC_OS_C1_TR01_T01: Ensure HTTPS succeeds
func GCS_CCC_OS_C1_TR01_T01() error {
	return errors.New("not yet implemented")
}

// GCS_CCC_OS_C1_TR01_T02: Ensure SFTP succeeds
func GCS_CCC_OS_C1_TR01_T02() error {
	return errors.New("not yet implemented")
}

// GCS_CCC_OS_C1_TR01_T03: Ensure gRPC over TLS succeeds
func GCS_CCC_OS_C1_TR01_T03() error {
	return errors.New("not yet implemented")
}

// GCS_CCC_OS_C1_TR01_T01: Ensure HTTP fails
func GCS_CCC_OS_C1_TR02_T01() error {
	return errors.New("not yet implemented")
}

// GCS_CCC_OS_C1_TR01_T02: Ensure FTP fails
func GCS_CCC_OS_C1_TR02_T02() error {
	return errors.New("not yet implemented")
}

// GCS_CCC_OS_C1_TR01_T03: Ensure unencrypted gRPC fails
func GCS_CCC_OS_C1_TR02_T03() error {
	return errors.New("not yet implemented")
}
