package main

import (
	"fmt"
	"strconv"
	"errors"
	"net/http"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

type APIResponse struct {
	Result int
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

	resp, err := http.Get("http://192.168.99.1:3000")
	if err != nil {
		return nil, err
	}

	r := APIResponse{}

	defer resp.Body.Close()

    err = json.NewDecoder(resp.Body).Decode(&r)
    if err != nil {
    	return nil, err
    }

    fmt.Println("================\n\n\n\n")
    fmt.Printf("Result: %v\n\n\n\n", r.Result)
    fmt.Println("================")

    err = stub.PutState("value", []byte(strconv.Itoa(r.Result)))

    return nil, err
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