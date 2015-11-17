// LoadRSAKeys
// Load rsa keys from private.key and public.key files
//
// You can execute the FTPServer like this:
// 		LoadRSAKeys
//
// Which will output something like:
//		Private key primes 108741439227668620461933533279790215465881923993363355200514950814177371533627 96978980597253607411035157882933085297507263638607277015670930870910646489267
//		Private key exponent 7000922239466582142639774668899318549729578179863210567078948279312924818567530131023770351855400223512440555393187695397565330251689164346432074518400017
//		Public key modulus 10545633924977507444198375298236063247991803005739064791179852748490648934321714918657041929434532995544495827070261918039113055690250404043221331185081409
//		Public key exponent 65537
//


package main

import (
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	var key rsa.PrivateKey

	loadKey("private.key", &key)

	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	var publicKey rsa.PublicKey

	loadKey("public.key", &publicKey)

	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)
}

func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)

	checkError(err)

	inFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}