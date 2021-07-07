package main

//외부모듈
import (
	"encoding/json"

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
	State        string `json: "state"` //go는 enum을 지원하지 않음.
	CloseTime    string `json: "close time"`
	EventEndTime string `json: "endtime"`
	LimitP       int    `json: "limitp"`
	Cost         int    `json: "cost"`
	NowP         int    `json: "nowp"`
}

type RegisteredUser struct {
	UserId  string   `json: "userid"`
	Money   string   `json:"user"`
	EventId []string `json: "eventid"`
	Ticket  string   `json: "ticket"`
}

const (
	REGISTERED int = iota
	CANCELED
	FINISHED
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
	} else if fn == "checkTime" {
		return s.checkTime(stub, args)
	} else if fn == "input" {
		return s.input(stub, args)
	} else if fn == "success" {
		return s.success(stub, args)
	} else if fn == "fail" {
		return s.fail(stub, args)
	}
	return shim.Error("Not supported smartcontract function name")
}

// registerEvent 함수
func (s *SmartContract) registerEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//args 갯수 확인
	if len(args) != 8 {
		return shim.Error("register Event function needs a parameter")
	}
	var event = Event{Name: args[0], Id: args[0], Registrant: args[0], Celebrity: args[0], State: args[0], CloseTime: args[0], EventEndTime: args[0], MaxP: 0, MinP: 0, Cost: 0, CostN: 0}
	eventAsBytes, _ := json.Marshal(event)
	stub.PutState(args[0], eventAsBytes)

	return shim.Success([]byte("tx is submitted"))
}

// checkTime 함수
func (s *SmartContract) checkTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

// input 함수
func (s *SmartContract) input(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//본인 id, eventid, 금액
	if len(args) != 3 {
		return shim.Error("input function needs 3 parameters")
	}
	//getState 먼저 해봐서 에러체크 ->본인 id가 있으면? error
	eventAsBytes, err := stub.GetState(args[0])
	if err != nil { //GetState가 수행중 오류를 가져오면
		return shim.Error("GetState function occured a error")
	}
	if eventAsBytes == nil {
		return shim.Error("id is not registered")
	}
	// var event = Event{}
	// _ = json.Unmarshal(eventAsBytes, &event)
	// event.CostN, _ = strconv.Atoi(args[2])

}

// fail 함수
func (s *SmartContract) fail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//limitp와 nowp, UserId, money
	if args[0] != args[1] {

		return shim.Success([]byte("인원이 미달되어 환불되었습니다."))
	}
}

// success 함수
func (s *SmartContract) success(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//minp 도달& 입장신호 -> celebrity에 input 전송
}
