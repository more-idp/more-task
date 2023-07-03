package more_task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/more-idp/more-task/utils"
	"github.com/redis/go-redis/v9"
)

const (
	KeyBase        = "ucodkr:task"
	KeyRedis       = "redis_host"
	KeyRedisPasswd = "redis_passwd"
	KeyRedisPort   = "redis_port"
	KeyHost        = "host"
	KeyName        = "name"
	KeyType        = "type"  // pub / sub
	KeyTopic       = "topic" //

)
const (
	DefaultMax = 4
)

type Request struct {
	Type  int    // 요청 종류 . job : 응답이 없는, task: 응답이 있는
	UUID  string //
	Param map[string]interface{}
	Topic string
}
type Response struct {
	Status  int      // 작업상태 , Working, Error, Done , job은 요청에 성공하면 바로 응답을 준다.
	ID      string   // 작업 ID
	Message []string // 응답메시지 , 메시지는 계속 채워진다.
	Req     *Request
}
type ClientConfig struct {
	RedisServer  string
	RedisUser    string
	RedisPasswd  string
	DefaultTopic string
	Name         string
	MaxRun       int
}
type Client struct {
	Connected bool
	config    *ClientConfig
	rdb       *redis.Client
}

func (c *Client) OnConnect(ctx context.Context, cn *redis.Conn) error {
	log.Println("Connected", cn)
	c.Connected = true
	info := utils.GetSysInfo()

	a, _ := json.Marshal(info)
	log.Println(string(a))
	err := cn.Set(ctx, KeyBase+":client", string(a), 0)
	log.Println("err", err.Err())
	// err := c.rdb.Set(ctx, KeyBase+":client", string(a), 0)
	// log.Println(err)
	return nil
}

func NewPubSub(config *ClientConfig) Client {
	if len(config.DefaultTopic) < 1 {
		config.DefaultTopic = "more"
	}
	cli := Client{config: config}
	opt := redis.Options{
		Addr:     config.RedisServer,
		Password: config.RedisPasswd, // no password set
		Username: config.RedisUser,   // no password set
		DB:       0,                  // use default DB

	}
	opt.OnConnect = cli.OnConnect

	rdb := redis.NewClient(&opt)

	ctx := context.Background()
	log.Println(ctx)
	cmd := rdb.Get(ctx, "XX")
	log.Println("cmd", cmd.Err(), cli)
	// log.Println(rdb, config, cmd)
	cli.rdb = rdb

	return cli
}

// 초기정보를 적재한다.
func (c Client) init() {

}
func (c Client) String() string {
	return fmt.Sprintf("srv:%s,name:%s,topic:%s", c.config.RedisServer, c.config.Name, c.config.DefaultTopic)
}
func (c *Client) uuid() string {
	return uuid.NewString()
}

// 데이터를 열어본다.
func (c *Client) Pop() (*Request, error) {
	ctx := context.Background()
	a := KeyBase + ":" + c.config.DefaultTopic
	b := c.rdb.BRPop(ctx, 0, a)
	if b == nil {
		return nil, nil
	}
	r := Request{}
	err := json.Unmarshal([]byte(b.Val()[1]), &r)
	if err != nil {
		log.Println("ERRO", b.Val())
		return nil, err
	}
	log.Println("pop b", r)
	return &r, b.Err()
}

// 작업을 구동한다.
// Req 에 응답을 요구하면
func (c *Client) RunTask(req *Request) (*Response, error) {
	ctx := context.Background()
	req.UUID = c.uuid()
	a := KeyBase + ":" + c.config.DefaultTopic
	b, _ := json.Marshal(req)
	log.Println(req, "marshal", string(b))
	cmd := c.rdb.LPush(ctx, a, string(b))

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	log.Println(cmd.Err())
	return nil, nil
}
