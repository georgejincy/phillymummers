package main

import (
    "fmt"
    "testing"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestCreateLoanApplication (t *testing.T) {
    fmt.Println("Entering TestCreateLoanApplication")
    attributes := make(map[string][]byte)
    //Create a custom MockStub that internally uses shim.MockStub
    stub := shim.NewCustomMockStub("mockStub", new(SampleChaincode), attributes)
    if stub == nil {
        t.Fatalf("MockStub creation failed")
    }
}