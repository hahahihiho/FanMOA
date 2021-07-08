package main

import (
	"fmt"
	"encoding/json"
	"strconv"
	"bytes"
	"time"
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

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "registerUser" {
		return s.registerUser(stub, args)
	} else if fn == "registerEvent" {
		return s.registerEvent(stub, args)
	} else if fn == "putMoney" {
		return s.putMoney(stub, args)
	} else if fn == "completeEvent" {
		return s.completeEvent(stub, args)
	} else if fn == "refundAll" {
		return s.refundAll(stub, args)
	} else if fn == "getUser"{
		return s.getUser(stub,args)
	} else if fn == "getEvent"{
		return s.getEvent(stub,args)
	} else if fn == "getUserHistory"{
		return s.getUserHistory(stub,args)
	}
	return shim.Error("Not supported smartcontract function name")
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

func (s *SmartContract) registerEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 8 {	// id, name, registrant, celebrity, closetime, eventendtime, limitp, fee, | balance, currentp, state
		return shim.Error("register Event function needs a parameter")
	}
	usr_id := args[2]
	cel_id := args[3]
	usrAsBytes, uerr := stub.GetState(usr_id)
	celAsBytes, cerr := stub.GetState(cel_id)
	if uerr == nil && cerr == nil && usrAsBytes != nil && celAsBytes != nil {
		limitp,_ := strconv.Atoi(args[6])
		fee,_ := strconv.Atoi(args[7])
		event := Event{Id: args[0], Name: args[1], Registrant: args[2], Celebrity: args[3], CloseTime: args[4], EventEndTime: args[5], LimitP: limitp, Fee: fee, Balance:0, CurrentP: 0, State:REGISTERED}
		eventAsBytes, _ := json.Marshal(event)

		usr := RegisteredUser{}
		_ = json.Unmarshal(usrAsBytes,&usr)
		usr.PlannedEvents = append(usr.PlannedEvents,event.Id)
		usrAsBytes,_ = json.Marshal(usr)

		stub.PutState(event.Id, eventAsBytes)
		stub.PutState(usr.Id,usrAsBytes)
		return shim.Success([]byte("tx is submitted"))
	} else {
		return shim.Error("No user")
	}
}

func (s *SmartContract) putMoney(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {	// id, event_id | fee
		return shim.Error("input function needs 3 parameters")
	}
	// registerUser
	usr_id:=args[0]
	evt_id:=args[1]
	userAsBytes, uerr :=stub.GetState(usr_id)
	eventAsBytes, eerr := stub.GetState(evt_id)
	if uerr == nil && eerr == nil {
		user := RegisteredUser{}
		_ = json.Unmarshal(userAsBytes, &user)
		user.PaidEvents = append(user.PaidEvents, evt_id)
		userAsBytes, _ = json.Marshal(user)

		evt := Event{}
		_ = json.Unmarshal(eventAsBytes, &evt)
		evt.CurrentP += 1
		evt.Balance += evt.Fee
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

// query
func (s *SmartContract) getUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {	//user_id
		return shim.Error("register Event function needs a parameter")
	}
	usrAsBytes,_:=stub.GetState(args[0])
	return shim.Success(usrAsBytes)
}
func (s *SmartContract) getEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {	//event_id
		return shim.Error("register Event function needs a parameter")
	}
	evtAsBytes,_:=stub.GetState(args[0])
	return shim.Success(evtAsBytes)
}
func (s *SmartContract) getUserHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {	//usr_id
		return shim.Error("register Event function needs a parameter")
	}

	keyName := args[0]
	// 로그 남기기
	fmt.Println("readTxHistory:" + keyName)

	resultsIterator, err := stub.GetHistoryForKey(keyName)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	// 로그 남기기
	fmt.Println("readTxHistory returning:\n" + buffer.String() + "\n")

	return shim.Success(buffer.Bytes())
}
func main() {
	err := shim.Start((new(SmartContract)))
	if err != nil {
		fmt.Printf("Error starting chaincode : %s", err)
	}
}
