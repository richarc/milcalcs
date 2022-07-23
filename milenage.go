package main

import (
	"encoding/hex"
	"fmt"

	"github.com/magma/milenage"
)

//clean this up when we do the variable milenage
func calculateOPC(k_string, op_string string) string {
	//These Ki and OP values are from 35 207, implementers test data
	//k_string := "465b5ce8b199b49faa5f0a2ee238a6bc"
	//op_string := "cdc202d5123e20f62b6d676ac72cb318"

	k_bytes, err := hex.DecodeString(k_string)
	if err != nil {
		return "Ki: Not a valid Hex String"
	}
	if len(k_bytes) != milenage.ExpectedKeyBytes {
		return fmt.Sprintf("Ki: incorrect key size. Expected %v bytes, but got %v bytes", milenage.ExpectedKeyBytes, len(k_bytes))
	}

	op_bytes, err := hex.DecodeString(op_string)
	if err != nil {
		return "OP: Not a valid Hex String"
	}

	opc_bytes, err := milenage.GenerateOpc(k_bytes, op_bytes)
	if err != nil {
		return fmt.Sprintf("Error %v ", err)
	}
	opc_string := fmt.Sprintf("%x", opc_bytes)
	return opc_string
}
