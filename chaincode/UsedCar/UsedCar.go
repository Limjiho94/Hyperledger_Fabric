/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi" //contractapli ==> SDK
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car

type Government struct {
	OwnerName    string `json:"ownername"`
	OwnerAddress string `json:"ownerAdd"`
	OwnerNum     string `json:"ownerNumber"`
}

type Manufacture struct {
	PIN          string `json:"PIN"`
	CCVol        string `json:"CCvol"`
	Model        string `json:"model"`
	ReleasePrice string `json:"relprice"`
	Color        string `json:"color"`
}

type Repair struct {
	RepairHistory string `json:"repairhistory"`
	RepairDate    string `json:"repairdate"`
	RepairPlace   string `json:"repairplace"`
	MechanicName  string `json:"machanicname"`
}

type Insurance struct {
	SubNumber   string `json:"subnumber"`
	ProductInfo string `json:"productinfo"`
	Subscriber  string `json:"subscriber"`
	ManagerName string `json:"managername"`
	CompanyName string `json:"companyname"`
}

type Total struct {
	OwnerName     string `json:"ownername"`
	OwnerAddress  string `json:"ownerAdd"`
	OwnerNum      string `json:"ownerNumber"`
	PIN           string `json:"PIN"`
	CCVol         string `json:"CCvol"`
	Model         string `json:"model"`
	ReleasePrice  string `json:"relprice"`
	Color         string `json:"color"`
	RepairHistory string `json:"repairhistory"`
	RepairDate    string `json:"repairdate"`
	RepairPlace   string `json:"repairplace"`
	MechanicName  string `json:"machanicname"`
	SubNumber     string `json:"subnumber"`
	ProductInfo   string `json:"productinfo"`
	Subscriber    string `json:"subscriber"`
	ManagerName   string `json:"managername"`
	CompanyName   string `json:"companyname"`
}

// QueryResult structure used for handling result of query

type QueryResult struct {
	Key string `json:"Key"`

	Record  *Total
	ARecord *Government
	BRecord *Manufacture
	CRecord *Repair
	DRecord *Insurance
}

//CreateCar adds a new car to the world state with given details

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	govs := []Government{
		Government{OwnerName: "sung", OwnerAddress: "S-class", OwnerNum: "010-2424"},
	}
	mans := []Manufacture{
		Manufacture{PIN: "1234", CCVol: "1999", Model: "M2", ReleasePrice: "9000", Color: "Red"},
	}
	reps := []Repair{
		Repair{RepairHistory: "Bumper_damage", RepairDate: "21-12-08", RepairPlace: "Samsung_RepaireCenter", MechanicName: "JihoLim"},
	}
	inss := []Insurance{
		Insurance{SubNumber: "1234-567-891", ProductInfo: "Good_Life_Product", Subscriber: "JihoLim", ManagerName: "David", CompanyName: "Hyundai"},
	}

	for i, gov := range govs {
		govAsBytes, _ := json.Marshal(gov)
		err := ctx.GetStub().PutState("GOV"+strconv.Itoa(i), govAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	for i, man := range mans {
		manAsBytes, _ := json.Marshal(man)
		err := ctx.GetStub().PutState("MAN"+strconv.Itoa(i), manAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	for i, rep := range reps {
		repAsBytes, _ := json.Marshal(rep)
		err := ctx.GetStub().PutState("REP"+strconv.Itoa(i), repAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	for i, ins := range inss {
		insAsBytes, _ := json.Marshal(ins)
		err := ctx.GetStub().PutState("INS"+strconv.Itoa(i), insAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

func (s *SmartContract) SetManufactureInfo(ctx contractapi.TransactionContextInterface, carNumber string, PIN string, CCvol string, model string, relprice string, color string) error {

	//results := []QueryResult{}

	man := Manufacture{
		PIN:          PIN,
		CCVol:        CCvol,
		Model:        model,
		ReleasePrice: relprice,
		Color:        color,
	}

	manAsBytes, _ := json.Marshal(man)

	//var GovCarNum string = "GOV" + carNumber
	//Car_Num = ctx.GetStub().PutState(carNumber, manAsBytes)
	//results = append(results, Car_Num)

	return ctx.GetStub().PutState(carNumber, manAsBytes)

}

func (s *SmartContract) SetGovernmentInfo(ctx contractapi.TransactionContextInterface, carNumber string, ownername string, owneradd string, ownernumber string) error {

	gov := Government{
		OwnerName:    ownername,
		OwnerAddress: owneradd,
		OwnerNum:     ownernumber,
	}

	govAsBytes, _ := json.Marshal(gov)
	//var GovCarNum string = "GOV" + carNumber
	return ctx.GetStub().PutState(carNumber, govAsBytes)
}

func (s *SmartContract) SetRepairInfo(ctx contractapi.TransactionContextInterface, carNumber string, repairhistory string, repairdate string, repairplace string, machanicname string) error {
	rep := Repair{
		RepairHistory: repairhistory,
		RepairDate:    repairdate,
		RepairPlace:   repairplace,
		MechanicName:  machanicname,
	}

	repAsBytes, _ := json.Marshal(rep)
	//var RepCarNum string = "REP" + carNumber
	return ctx.GetStub().PutState(carNumber, repAsBytes)
}

func (s *SmartContract) SetInsuranceInfo(ctx contractapi.TransactionContextInterface, carNumber string, subnumber string, productinfo string, subscriber string, managername string, companyname string) error {
	ins := Insurance{
		SubNumber:   subnumber,
		ProductInfo: productinfo,
		Subscriber:  subscriber,
		ManagerName: managername,
		CompanyName: companyname,
	}

	insAsBytes, _ := json.Marshal(ins)
	//var InsCarNum string = "INS" + carNumber
	return ctx.GetStub().PutState(carNumber, insAsBytes)
}

func (s *SmartContract) QueryTotalInfo(ctx contractapi.TransactionContextInterface, carNumber string) (*Total, error) {

	tolAsBytes, err := ctx.GetStub().GetState(carNumber)

	tol := new(Total)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if tolAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", carNumber)
	}

	_ = json.Unmarshal(tolAsBytes, tol)

	return tol, nil
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAll(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}
		tol := new(Total)
		_ = json.Unmarshal(queryResponse.Value, tol)
		//man := new(Manufacture)
		//_ = json.Unmarshal(queryResponse.Value, man)
		//gov := new(Government)
		//_ = json.Unmarshal(queryResponse.Value, gov)
		//rep := new(Repair)
		//_ = json.Unmarshal(queryResponse.Value, rep)
		//ins := new(Insurance)
		//_ = json.Unmarshal(queryResponse.Value, ins)

		queryResult := QueryResult{Key: queryResponse.Key, Record: tol}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
	tol, err := s.QueryTotalInfo(ctx, carNumber) //asdl;kgjasd;lkgjasdl;kgjasd;lkgjasdk;l~!!!!!!!!!!!!!!!!!!

	if err != nil {
		return err
	}

	tol.OwnerName = newOwner

	tolAsBytes, _ := json.Marshal(tol)

	return ctx.GetStub().PutState(carNumber, tolAsBytes)

}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
