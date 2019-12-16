package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/bregydoc/shh"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	x, p, err := shh.NewWizard().GeneratePair()
	if err != nil {
		panic(err)
	}

	spew.Dump(x)
	fmt.Println()
	spew.Dump(p)



	fmt.Println(string(memPriv))
	fmt.Println("=======----=======")
	fmt.Println(string(memPub))
}
