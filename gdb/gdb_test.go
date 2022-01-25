package gdb

import "testing"

func TestOpen(t *testing.T) {

	Open("root:HongluoJdjhDB@1234@tcp(8.141.57.107:3306)/jdjh_gcm")

	QueryCount()

}
