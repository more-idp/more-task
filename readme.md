# 개요

분산작업 엔진
Redis 를 이용한 작업의 요청/ 작업의 처리 / 작업결과의 응답에 관련된 기본 구조를 설계한다.

# 구현
Redis 의 put , lpop( Locked pop )을 이용하여 여러개의 Worker 가 동작할수 있도록 한다.

하나의 Redis 에 여러 작업군들이 동작할수 있기에 초기에 관련 정보들을 기입하여 상태를 관리할수 있도록 한다.

- 세가지의 구현체를 구성한다.
 * golang
 * nodejs(typescript)
 * python
  
## 흐름
두개의 채널을 만들어서 이용한다.
* 요청채널
* 응답채널
### 요청자 (Publisher)

* pub은 지정된 작업에 정보를 입력하여 요청한다.
* 응답을 받아야 한다면, 응답채널을 poll 한다.
* 
### 소비자 (Consumer, Worker)
* 하나의 아이템을 가지고 온다.
* 응답을 받아야 되는 작업이면, 가지고 온후 작업의 상태에 대해서 Push 한다.
* 작업을 수행한다.
* 요청 작업에 hook 이 존재한다면 hook 을 콜한다.


## 저장키
/ucodkr/task/worker  : 초기에 작업자를 등록한다. 
Pub Lib 은 초기에 자신의 키값들을 등록하고, Sub(Worker) 는 존재하는 작업에 풀을 할수 있도록 한다.
Worker 는 항상 자신의 ID (Host 정보) 를 주기적으로 등록하여야 한다. (TTL 을 이용하여 Worker 의 생존 여부를 확인할수 있어야 한다.)

베이스 키
* ucodkr/task/worker
    id: worker 이름
    host: 동작 Host 이름
    desc: 설명
    last: 최종 Active

## Publisher


## Worker 
### Put

## golan


## nodejs


## python
