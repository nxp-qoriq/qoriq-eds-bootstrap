/*************************
*                        *
*   Copyright 2019 NXP   *
*                        *
*************************/

package main

import (
	"encoding/base64"
	"fmt"
	"github.com/laurentluce/est-client-go"
	"io/ioutil"
	"os/exec"
)

func dec(in []byte) string {
	out := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i] ^ byte(195-i)
	}
	return string(out)
}

func main() {
	fmt.Printf("starting Phase1\n")

	var (
		rootCA []byte
		err    error
	)
	if RootCA != "" {
		rootCA, err = base64.RawStdEncoding.DecodeString(RootCA)
		fmt.Println(err)
	} else {
		fmt.Println("Download EST Server RootCA")
		cmd := "curl -f https://image.edgescale.org/CA/int.b-est.edgescale.org.rootCA.pem -o /tmp/rootCA.pem"
		err = exec.Command("bash", "-c", cmd).Run()
		if err != nil {
			rootCA, _ = ioutil.ReadFile("/etc/ssl/certs/ca-certificates.crt")
		} else {
			rootCA, _ = ioutil.ReadFile("/tmp/rootCA.pem")
		}
	}

	client := est.Client{
		URLPrefix:  URLPrefix,
		Username:   dec(Name),
		Password:   dec(Password),
		ServerCert: rootCA}

	commonName := "Bootstrap"
	country := "CN"
	state := "China"
	city := "Beijing"
	organization := ""
	organizationalUnit := ""
	emailAddress := "admin@localhost.ltd"

	fmt.Println("create PKCS10 request")
	priv, csr, _ := est.CreateCsr(commonName, country, state, city, organization, organizationalUnit, emailAddress)
	fmt.Printf("Starting B-EST certificate Enrollment\n")
	cert, err := client.SimpleEnroll(csr)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("/tmp/bootstrap.pem", cert, 0644)
	ioutil.WriteFile("/tmp/bootstrap.key", priv, 0644)
}
