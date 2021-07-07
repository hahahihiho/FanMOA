package main

//외부모듈
import (
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//클래스 정의
type SmartContract struct {
}

type Event struct {
	Name         string `json: "name"`
	Id           string `json: "eventid"`
	Registrant   string `json: "reg"`
	Celebrity    string `json: "celebrity"`
	CloseTime    string `json: "closetime"`
	EventEndTime string `json: "endtime"`
	LimitP       int    `json: "limitp"`
	Fee          int    `json: "fee"`
	Balance		 int 	`json: "balance"`
	CurrentP     int    `json: "currentp"`
	State        int 	`json: "state"`
}

type RegisteredUser struct {
	Id  string   `json: "userid"`
	Balance   int   `json:"balance"`
	PlannedEvents []string `json: "plannede"`
	PaidEvents  []string   `json: "paide"`
}

const (
	REGISTERED int = iota
	FINISHED
	CANCELED
)
const (
	NEW int = iota
	USED
	OVER
)

//init 함수
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

//invoke 함수
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "registerEvent" {
		return s.registerEvent(stub, args)
	} else if fn == "registerUser" {
		return s.registerUser(stub, args)
	} else if fn == "putMoney" {
		return s.putMoney(stub, args)
	} else if fn == "completeEvent" {
		return s.completeEvent(stub, args)
	} else if fn == "refundAll" {
		return s.refundAll(stub, args)
	}
	return shim.Error("Not supported smartcontract function name")
}

func (s *SmartContract) registerEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 8 {	// id, name, registrant, celebrity, closetime, eventendtime, limitp, fee, | balance, currentp, state
		return shim.Error("register Event function needs a parameter")
	}
	limitp,_ := strconv.Atoi(args[6])
	fee,_ := strconv.Atoi(args[7])
	event := Event{Id: args[0], Name: args[1], Registrant: args[2], Celebrity: args[3], CloseTime: args[4], EventEndTime: args[5], LimitP: limitp, Fee: fee, Balance:0, CurrentP: 0, State:REGISTERED}
	eventAsBytes, _ := json.Marshal(event)
	stub.PutState(args[0], eventAsBytes)

	return shim.Success([]byte("tx is submitted"))
}

func (s *SmartContract) registerUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {	// id | balance, plannedevent, paidevent
		return shim.Error("register Event function needs a parameter")
	}
	user := RegisteredUser{Id: args[0], Balance: 0, PlannedEvents: make([]string, 0), PaidEvents:make([]string, 0)}
	dataAsBytes, _ := json.Marshal(user)
	stub.PutState(args[0], dataAsBytes)

	return shim.Success([]byte("tx is submitted"))
}

func (s *SmartContract) putMoney(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {	// id, event_id, fee
		return shim.Error("input function needs 3 parameters")
	}
	// registerUser
	usr_id:=args[0]
	evt_id:=args[1]
	strfee:=args[2]
	userAsBytes, uerr :=stub.GetState(usr_id)
	eventAsBytes, eerr := stub.GetState(evt_id)
	if uerr == nil && eerr == nil {
		user := RegisteredUser{}
		_ = json.Unmarshal(userAsBytes, &user)
		user.PaidEvents = append(user.PaidEvents, evt_id)
		userAsBytes, _ = json.Marshal(user)

		evt := Event{}
		_ = json.Unmarshal(eventAsBytes, &evt)
		fee,_ := strconv.Atoi(strfee)
		evt.Balance += fee
		eventAsBytes, _ = json.Marshal(evt)

		stub.PutState(usr_id, userAsBytes)
		stub.PutState(evt_id, eventAsBytes)
		return shim.Success([]byte("put money"))
	} else {
		return shim.Error("err")
	}	
}

func (s *SmartContract) completeEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {	//event_id
		return shim.Error("register Event function needs a parameter")
	}
	// event state update
	// event balance -> (if exist) celebrity balance
	evt_id := args[0]
	eventAsBytes, eerr := stub.GetState(evt_id)
	if eerr == nil {
		evt := Event{}
		_ = json.Unmarshal(eventAsBytes, &evt)

		usr_id:=evt.Celebrity
		userAsBytes, uerr := stub.GetState(usr_id)
		if uerr == nil && userAsBytes != nil {
			user := RegisteredUser{}
			_ = json.Unmarshal(userAsBytes, &user)
			user.Balance = evt.Balance
			evt.Balance = 0
			evt.State = FINISHED

			userAsBytes, _ = json.Marshal(user)
			eventAsBytes, _ = json.Marshal(evt)

			stub.PutState(evt_id, eventAsBytes)
			stub.PutState(usr_id, userAsBytes)
		} else {
			return shim.Error("err")
		}
	}

	return shim.Success([]byte("tx is submitted"))
}

func (s *SmartContract) refundAll(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//future
	return shim.Error("Not implemented yet")
}

func main() {
	err := shim.Start((new(SmartContract)))
	if err != nil {
		fmt.Printf("Error starting chaincode : %s", err)
	}
}
