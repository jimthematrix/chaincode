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

var KEY = "balance"

func (cc *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	var arg int
	var err error

	arg, err = strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Expecting an integer value for the init function during chaincode deploy")
	}

    err = stub.PutState(KEY, []byte(strconv.Itoa(arg)))

	return nil, err
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

    fmt.Println("================\n\n")
    fmt.Printf("Result: %v\n\n", r.Result)
    fmt.Println("================")

	valbytes, err := stub.GetState(KEY)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}

	val, _ := strconv.Atoi(string(valbytes))

    err = stub.PutState(KEY, []byte(strconv.Itoa(val - r.Result)))

    return nil, err
}

func (cc *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	valbytes, err := stub.GetState(KEY)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"value\",\"Amount\":\"" + string(valbytes) + "\"}"

    fmt.Println("================\n\n")
	fmt.Printf("Query Response:%s\n\n", jsonResp)
    fmt.Println("================")

	return valbytes, nil
}

func main() {
	var err error
	err = shim.Start(new (SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %e", err)
	}
}