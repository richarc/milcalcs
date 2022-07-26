package main

import (
	"encoding/hex"
	"fmt"

	"github.com/forgoer/openssl"
	"github.com/magma/milenage"
)

//struct to represent the defualt Milenage constants
//while 35 206 defines them as constant they are actually variable

type c16 [16]byte
type milconst struct {
	r1, r2, r3, r4, r5 int8
	c1, c2, c3, c4, c5 c16
}

var milvals = milconst{
	r1: 64,
	r2: 0,
	r3: 32,
	r4: 64,
	r5: 96,
	c1: [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	c2: [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	c3: [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
	c4: [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8},
	c5: [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16},
}

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

// calulate an eKi or an eOPc given a K4 key and the selected algorthm
//this also needs a big clean up
func calculateK4out(k4_string, alg_option, kiopc_string string) string {

	switch alg_option {
	case "DES ECB":
		//convert string to []bytes and check thathe key length is 8 bytes
		key := []byte(k4_string)
		if len(key) != 8 {
			return "Error: Key length must be 8"
		}
		src := []byte(kiopc_string)
		dst, err := openssl.DesECBEncrypt(src, key, openssl.PKCS7_PADDING)
		if err != nil {
			return fmt.Sprintf("Error %v ", err)
		}
		return fmt.Sprintf("%x", dst)

	case "3DES ECB":
		//convert string to []bytes and check thathe key length is 8 bytes
		key := []byte(k4_string)
		if len(key) != 24 {
			return "Error: Key length must be 24"
		}
		src := []byte(kiopc_string)
		dst, err := openssl.Des3ECBEncrypt(src, key, openssl.PKCS7_PADDING)
		if err != nil {
			return fmt.Sprintf("Error %v ", err)
		}
		return fmt.Sprintf("%x", dst)

	case "AES 128 ECB":
		//convert string to []bytes and check thathe key length is 8 bytes
		key := []byte(k4_string)
		if len(key) != 16 {
			return "Error: Key length must be 16"
		}
		src := []byte(kiopc_string)
		dst, err := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
		if err != nil {
			return fmt.Sprintf("Error %v ", err)
		}
		return fmt.Sprintf("%x", dst)
	}

	return "Invalid Algo Selected"
}
