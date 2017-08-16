package main 
import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer" 
) 

type SmartContract struct {
}

type Block struct {
	IM_ID string `json:"Im_Id"`
	BR_Ref_Id string `json:"Br_Ref_Id"`
	QTY string `json:"Qty"`
	STOCK_CODE string `json:"Stock_Code"`
	COMM string `json:"Comm"`
	Trade_Date string `json:"Trade_Date"`
	Settle_Date string `json:"Settle_Date"`
	Buy string `json:"Buy"`
	Avrg_Price string `json:"Avrg_Price"`
	Total_Amount string `json:"Total_Amount"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "createBlock" {
		return s.createBlock(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	blocks := []Block{
		Block{IM_ID: "IM001", BR_Ref_Id: "Br-Ref-002", QTY: "10", STOCK_CODE: "123", COMM: "123", Trade_Date: "20170816", Settle_Date: "20170818", Buy: 
"B", Avrg_Price: "500", Total_Amount: "1000"},
		Block{IM_ID: "IM001", BR_Ref_Id: "Br-Ref-003", QTY: "10", STOCK_CODE: "123", COMM: "123", Trade_Date: "20170816", Settle_Date: "20170818", Buy: 
"B", Avrg_Price: "500", Total_Amount: "1000"},
		Block{IM_ID: "IM001", BR_Ref_Id: "Br-Ref-004", QTY: "10", STOCK_CODE: "123", COMM: "123", Trade_Date: "20170816", Settle_Date: "20170818", Buy: 
"B", Avrg_Price: "500", Total_Amount: "1000"},
		Block{IM_ID: "IM001", BR_Ref_Id: "Br-Ref-005", QTY: "10", STOCK_CODE: "123", COMM: "123", Trade_Date: "20170816", Settle_Date: "20170818", Buy: 
"B", Avrg_Price: "500", Total_Amount: "1000"},
	}
	i := 0
	for i < len(blocks) {
		fmt.Println("i is ", i)
		blockAsBytes, _ := json.Marshal(blocks[i])
		APIstub.PutState("BLOCK"+strconv.Itoa(i), blockAsBytes)
		fmt.Println("Added", blocks[i])
		i = i + 1
	}
	return shim.Success(nil)
}

func (s *SmartContract) createBlock(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 11")
	}
	var block = Block{
		IM_ID: args[1],
		BR_Ref_Id: args[2],
		QTY: args[3],
		STOCK_CODE: args[4],
		COMM: args[5],
		Trade_Date: args[6],
		Settle_Date: args[7],
		Buy: args[8],
		Avrg_Price: args[9],
		Total_Amount: args[10],
	}
	blockAsBytes, _ := json.Marshal(block)
	APIstub.PutState(args[0], blockAsBytes)
	return shim.Success(nil)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
