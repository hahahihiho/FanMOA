# FanMOA

![FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled.png](FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled.png)

- 2021. 07. 08 발표자료 수정본(다운가능)
    
    [팬MOA.pptx](FanMOA%2053b80be1e4004f0da279d86cddc1c625/MOA.pptx)
    

## 개요

팬MOA는 팬들이 직접 원하는 공연을 추진할 수 있도록 돕는 **공연 공동구매 플랫폼**입니다. 

일방적으로 주최사의 통보를 바탕으로 참여해야했던 기존 공연의 틀을 깨고 **팬들이 주체적으로 나서서 공연을 요청**할 수 있습니다. 

![FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled.jpeg](FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled.jpeg)

## 1. 추진배경

### 1-1 프로젝트 배경

[≪한국NGO신문≫ 300명 울리고 '먹튀' 최악 공동구매 '우자매맘'](http://www.ngonews.kr/115194)

카페에서 공동구매를 주도했던 우자매맘 사건

[유명 SNS 공구 마켓서 피해자 속출...'제2의 우자매맘' 사건 되나](http://www.ntoday.co.kr/news/articleView.html?idxno=76775)

'엣지베베'라고 하는 공동구매 플랫폼의 운영자의 먹튀사건

[YB 공연취소, 첫 온라인 공연 무산..."주최 측의 일방적 통보"](https://www.etoday.co.kr/news/view/2018156)

주최측의 일방적인 통보로 무산되어버린 공연

기존의 공동구매는 **'사기' 문제**가 주요한 이슈였습니다. 개인간의 공동구매 뿐만 아니라 공동구매 플랫폼에서도 거액의 사기가 발생하며 큰 플랫폼일지라도 **신뢰하기 어렵다는 문제점**이 있습니다.

공연 같은 경우 대형 소속사가 있어 직접 공연을 추진하는 경우도 있지만 한 주최사에서 섭외하여 진행하는 경우 주최사의 **일방적인 통보로 공연 준비가 물거품이 되는 경우도 발생**합니다.

---

### 1-2 기존 문제점

- 일반적인 공동구매의 경우 돈이 바로 주최자의 통장에 입금된다.
- 공동구매 사기가 발생하는 경우 환불에 어려움을 겪는다.
- 행사는 주최측의 일방적 통보로 진행된다.

---

### 1.3 해결 방안

- 거래 내역을 전부 **블록체인**에 기록한다.
- 중앙화된 곳에 돈을 송금하는 것이 아닌, **스마트 컨트랙트를 통해 조건이 달성되지 않으면 모금 금액이 넘어가지 않도록** 한다.
- 공연 제안은 팬이 하고 공동구매 모금액은 **공연주체에게 정산**될 수 있도록 한다.
- 팬의 아이디어로 시작해 가수/행사 주체의 주도로 진행되기 때문에 **손해없이 공연**을 열 수 있다.

---

### 1.4 비즈니스 모델

![FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%201.png](FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%201.png)

- 평소 보고/듣고싶었던 강의나 공연이 있었던 사용자를 타겟팅, **직접 공연을 제안하고 추진할 수 있다는 점**을 내세워 홍보한다.
- 팬MOA을 통해 원하는 **공연/강의를 제안**하고 **참여자들을 모을 수 있는 플랫폼**을 제공하여 수익을 창출한다.

---

### 1.5 운영 시나리오

![FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%201.jpeg](FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%201.jpeg)

---

## 2. 개발 목표 및 내용

### 2.1 개발계획

![FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%202.png](FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%202.png)

- **주요기능**
    1. 이용자(팬)가 직접 공연 **등록**
    2. 다른 이용자들이 공연 세부사항 **조회 및 결제**
    3. 공연 입장시 contract transfer 작동
    4. 이용자가 마이페이지에서 환불, 참여내역 확인
    

prototype page : [https://ovenapp.io/project/ylWPXaxiO97vpE9X7yUsk4Tom9PEqnnF#mlgMK](https://ovenapp.io/project/ylWPXaxiO97vpE9X7yUsk4Tom9PEqnnF#mlgMK)

---

### 2.2 소프트웨어 구성도

![FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%202.jpeg](FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%202.jpeg)

User(팬, 행사주체) ↔ front-end+server ↔ blockchain-network(smartcontract)

유저는 웹 브라우저를 통하여 “팬MOA" 웹 플랫폼에 접속한다.

- F**ront-end**
    
    ▶ Html를 이용하여 공연(Event) 공동구매 플랫폼을 웹 형태로 구축함
    
- **Server**
    
    ▶ Node.js를 이용하여 User와 Event 관리를 위한 서버를 구축함
    
- **Block Chain**
    
    ▶ Hyperledger fabric으로 블록체인 네트워크 구성
    
    ▶ User가 제안한 Event 내용과 거래 내용을 블록에 기록.
    
- **Smart Contract**
    
    ▶ Hyper-ledger fabric v1.7로 Chaincode를 작성함.
    

## 3. 설계서

### 3-1 App 설계서

[App](FanMOA%2053b80be1e4004f0da279d86cddc1c625/App%2030c1a713d53744ac8d0516ef8755e3fd.md)

### 3-2 Chain Code 설계서

[Chain Code](FanMOA%2053b80be1e4004f0da279d86cddc1c625/Chain%20Code%20ce4f43e26244417ba623cebfb914993e.md)

[DApp url ](FanMOA%2053b80be1e4004f0da279d86cddc1c625/DApp%20url%20713171c8d7444380a7a9d3c5eba3d190.csv)

## 4. 주요 특징 및 핵심기술

![FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%203.jpeg](FanMOA%2053b80be1e4004f0da279d86cddc1c625/Untitled%203.jpeg)

▪️ 펀딩의 내용과 거래 과정을 전부 블록체인에 기록하여 데이터 변조가 불가능함 ➡️ 신뢰성 확보

▪️ 조건이 달성되면 자동으로 계약이 실행됨 ➡️ 편의성 확보

▪️ 펀딩 모금액이 블록체인에 기록되어있기 때문에 모금액 사기 위험을 방지 ➡️ 안전성 확보

▪️  자금 조달이 어려웠던 행사를 미리 모금하여 수요조사 및 자금 확보 후 진행 가능. ➡️ 기회 제공

[color](FanMOA%2053b80be1e4004f0da279d86cddc1c625/color%20a4642da9a5334ce4b7ff3967b24a9b7a.md)

[link](FanMOA%2053b80be1e4004f0da279d86cddc1c625/link%20d27d9c3685c24d048ee80b4246af0d7a.md)