// GenX509Cert
// A program to generate a self-signed X.509 certificate for my web site and store it in a .cer file
//
// You can execute the GenX509Cert like this:
// 		GenX509Cert
//
// after execution, program will create new files:
//		mujib.programmer.name.cer
// 		mujib.programmer.name.pem
//
//	If you don't have a private.key file for your private key, you can create one by executing another program:
//		GenRSAKeys


package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader

	var key rsa.PrivateKey

	loadKey("private.key", &key)

	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000) // one year

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:
			"mujib.programmer.name",
			Organization: []string{"Mujib Programmer"},
		},
		// NotBefore: time.Unix(now, 0).UTC(),
		// NotAfter: time.Unix(now+60*60*24*365, 0).UTC(),
		NotBefore: now,
		NotAfter: then,
		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage:
		x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:
		true,
		DNSNames:
		[]string{"mujib.programmer.name", "localhost"},
	}

	derBytes, err := x509.CreateCertificate(random, &template,
		&template, &key.PublicKey, &key)
	checkError(err)

	certCerFile, err := os.Create("mujib.programmer.name.cer")
	checkError(err)

	certCerFile.Write(derBytes)
	certCerFile.Close()

	certPEMFile, err := os.Create("mujib.programmer.name.pem")
	checkError(err)

	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPEMFile.Close()

	keyPEMFile, err := os.Create("private.pem")
	checkError(err)

	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key)})

	keyPEMFile.Close()
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