package main

import (
	"fmt"
	"strconv"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

func (cc *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	var arg int
	var err error

	arg, err = strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Expecting an integer value for the init function during chaincode deploy")
	}

	fmt.Printf("arguments: %a", arg)

	return nil, nil
}

func (cc *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke function called")
	return nil, nil
}

func (cc *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Query function called")
	return nil, nil
}

func main() {
	var err error
	err = shim.Start(new (SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %e", err)
	}
}