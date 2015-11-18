// GenRSAKeys
// A program to generate RSA private and public keys
//
// You can execute the GenRSAKeys like this:
// 		GenRSAKeys
//
// Which will output something like:
//		Private key primes 108741439227668620461933533279790215465881923993363355200514950814177371533627 96978980597253607411035157882933085297507263638607277015670930870910646489267
//		Private key exponent 7000922239466582142639774668899318549729578179863210567078948279312924818567530131023770351855400223512440555393187695397565330251689164346432074518400017
//		Public key modulus 10545633924977507444198375298236063247991803005739064791179852748490648934321714918657041929434532995544495827070261918039113055690250404043221331185081409
//		Public key exponent 65537
//
// after execution, program will create new files:
//		private.key
// 		private.pem
// 		public.key


package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {

	reader := rand.Reader

	bitSize := 512

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	publicKey := key.PublicKey

	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

	saveGobKey("private.key", key)

	saveGobKey("public.key", publicKey)

	savePEMKey("private.pem", key)
}

func saveGobKey(fileName string, key interface{}) {

	outFile, err := os.Create(fileName)
	checkError(err)

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)

	checkError(err)

	outFile.Close()
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {

	outFile, err := os.Create(fileName)
	checkError(err)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)

	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}