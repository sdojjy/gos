package ssls

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"time"
)

func GetCertEndTime(filepath string) (startTime time.Time, endTime time.Time, err error) {
	certPEMBlock, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		return
	}
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return
	}

	startTime, endTime = x509Cert.NotBefore, x509Cert.NotAfter

	return
}
