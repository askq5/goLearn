package main

import (
	"fmt"
	"time"

	"code.byted.org/kv/goredis"
)

// exampleCmd: simple demo for proxy supported commands
func exampleCmd() {
	// init option
	options := goredis.NewOption()
	// set consul discovery, default false
	options.SetServiceDiscoveryWithConsul()
	// connection pool initial size, default 10
	options.SetPoolInitSize(4)
	// set auto load interval, default 30s
	options.SetAutoLoadInterval(time.Second * 30)
	// Circuit Breader: max failure rate, min sample, time window, default 0.6, 50, 10s
	options.SetCircuitBreakerParam(0.6, 50, time.Millisecond*10000)
	// timeout parameter
	options.DialTimeout = 50 * time.Millisecond
	options.ReadTimeout = 50 * time.Millisecond
	options.WriteTimeout = 50 * time.Millisecond
	options.PoolTimeout = options.ReadTimeout + time.Second
	options.IdleTimeout = 5 * time.Minute
	options.LiveTimeout = time.Hour
	// for more avaiabile option, highly recommend to read option.go

	// init connect
	cli, err := goredis.NewClientWithOption("toutiao.redis.liguancheng_db", options)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[CMD] PING: %s\n", cli.Ping())

	ret, err := cli.Set("askq_0", 0, 1*time.Second).Result()

	fmt.Printf("ret:%+v\n", ret)
	msetVals := make([]interface{}, 0)
	msetVals = append(msetVals, "askq_1")
	msetVals = append(msetVals, 1)
	msetVals = append(msetVals, "askq_2", 2)
	fmt.Println(msetVals)
	fmt.Printf("%s\n", cli.MSet(msetVals...))
	keys := []string{"askq_0", "askq_1", "askq_2"}
	fmt.Printf("%s\n", cli.MGet(keys...))

	//ret := cli.Set("hello_1", "world", 1*time.Second)
	//
	//res, err := ret.Result()
	//if err != nil {
	//
	//}
	//fmt.Printf("[CMD] SET: %s\n", res)
	//// command test
	//fmt.Printf("[CMD] PING: %s\n", cli.Ping())

	// string
	// tip: string length < 10k
	//fmt.Printf("[CMD] SET: %s\n", cli.Set("hello_1", "world", 0))
	//fmt.Printf("[CMD] SETEX: %s\n", cli.Set("hello_2", "world", 3600*time.Second))
	//fmt.Printf("[CMD] SET: %s\n", cli.Set("hello_3", "world", 1*time.Hour))
	//fmt.Printf("[CMD] DEL: %s\n", cli.Del("hello_4"))
	//fmt.Printf("[CMD] SETNX: %s\n", cli.SetNX("hello_4", "100", 0))
	//fmt.Printf("[CMD] STRLEN: %s\n", cli.StrLen("hello_4"))
	//fmt.Printf("[CMD] SETRANGE: %s\n", cli.SetRange("hello_2", 2, "NEW"))
	//fmt.Printf("[CMD] GET: %s\n", cli.Get("hello_1"))
	//fmt.Printf("[CMD] GETSET: %s\n", cli.GetSet("hello_1", "new world"))
	////fmt.Printf("[CMD] EXISTS: %d, error:%+v\n", cli.Exists("hello_3").Result())
	//fmt.Printf("[CMD] EXPIRE: %s\n", cli.Expire("hello_3", time.Hour))
	//fmt.Printf("[CMD] EXPIREAT: %s\n", cli.ExpireAt("hello_3", time.Now().Add(+time.Minute*10)))
	//fmt.Printf("[CMD] PERSIST: %s\n", cli.Persist("hello_3"))
	//fmt.Printf("[CMD] PEXPIRE: %s\n", cli.PExpire("hello_3", time.Duration(5)*time.Second))
	//fmt.Printf("[CMD] PEXPIREAT: %s\n", cli.PExpireAt("hello_3", time.Now().Add(+time.Hour*10)))
	//fmt.Printf("[CMD] TTL: %s\n", cli.TTL("hello_3"))
	//fmt.Printf("[CMD] PTTL: %s\n", cli.PTTL("hello_3"))
	//fmt.Printf("[CMD] TYPE: %s\n", cli.Type("hello_3"))
	//fmt.Printf("[CMD] APPEND: %s\n", cli.Append("hello_1", "new world"))
	//fmt.Printf("[CMD] SETBIT: %s\n", cli.SetBit("bits", 0, 1))
	//fmt.Printf("[CMD] BITCOUNT: %s\n", cli.BitCount("bits", &redis.BitCount{Start: 1, End: 1}))
	//fmt.Printf("[CMD] DECR: %s\n", cli.Decr("hello_4"))
	//fmt.Printf("[CMD] DECRBY: %s\n", cli.DecrBy("hello_4", 1))
	//fmt.Printf("[CMD] INCR: %s\n", cli.Incr("hello_4"))
	//fmt.Printf("[CMD] INCRBY: %s\n", cli.IncrBy("hello_4", 10))
	//fmt.Printf("[CMD] INCRBYFLOAT: %s\n", cli.IncrByFloat("hello_4", 10.5))

	// mget
	// tip: commands in one mget < 100
	//fmt.Printf("[CMD] MGET: %s\n", cli.MGet("hello_1", "hello_2", "hello_3"))

	// hash
	// tip: elements number < 5000
	/*fmt.Printf("[CMD] HSET: %s\n", cli.HSet("h_hello_1", "f_1", "world_1"))
	fmt.Printf("[CMD] HSET: %s\n", cli.HSet("h_hello_1", "f_2", 10))
	fmt.Printf("[CMD] HSETNX: %s\n", cli.HSetNX("h_hello_1", "f_3", "world_3"))
	fmt.Printf("[CMD] HMSET: %s\n", cli.HMSet("h_hello_1", map[string]interface{}{"f_3": "world_3", "f_4": "world_4"}))
	fmt.Printf("[CMD] HGETALL: %s\n", cli.HGetAll("h_hello_1")) // hgetall not recommanded
	fmt.Printf("[CMD] HDEL: %s\n", cli.HDel("h_hello_1", "f_1"))
	fmt.Printf("[CMD] HEXISTS: %s\n", cli.HExists("h_hello_1", "f_1"))
	fmt.Printf("[CMD] HINCRBY: %s\n", cli.HIncrBy("h_hello_1", "f_2", 1))
	fmt.Printf("[CMD] HINCRBY: %s\n", cli.HIncrBy("h_hello_1", "f_2", -1))
	fmt.Printf("[CMD] HINCRBYFLOAT: %s\n", cli.HIncrByFloat("h_hello_1", "f_2", 0.1))
	fmt.Printf("[CMD] HINCRBYFLOAT: %s\n", cli.HIncrByFloat("h_hello_1", "f_2", -0.1))
	fmt.Printf("[CMD] HKEYS: %s\n", cli.HKeys("h_hello_1"))
	fmt.Printf("[CMD] HVALS: %s\n", cli.HVals("h_hello_1"))
	fmt.Printf("[CMD] HLEN: %s\n", cli.HLen("h_hello_1"))*/

	// list
	// tip: elements number < 5000
	//fmt.Printf("[CMD] LPUSH: %s\n", cli.LPush("l_hello_1", "world"))
	//fmt.Printf("[CMD] LPUSHX: %s\n", cli.LPushX("l_hello_2", "world"))
	//fmt.Printf("[CMD] LRANGE: %s\n", cli.LRange("l_hello_1", 0, -1))
	//fmt.Printf("[CMD] LLEN: %s\n", cli.LLen("l_hello_1"))
	//fmt.Printf("[CMD] LSET: %s\n", cli.LSet("l_hello_1", 3, "new world"))
	//fmt.Printf("[CMD] LREM: %s\n", cli.LRem("l_hello_2", 2, "world"))
	//fmt.Printf("[CMD] LTRIM: %s\n", cli.LTrim("l_hello_2", 2, 3))
	//fmt.Printf("[CMD] LPOP: %s\n", cli.LPop("l_hello_1"))
	//fmt.Printf("[CMD] LINDEX: %s\n", cli.LIndex("l_hello_1", 2))
	//fmt.Printf("[CMD] LINSERT: %s\n", cli.LInsert("l_hello_1", "BEFORE", "world", "new"))
	//fmt.Printf("[CMD] RPUSH: %s\n", cli.RPush("l_hello_1", "world"))
	//fmt.Printf("[CMD] RPUSHX: %s\n", cli.RPushX("l_hello_1", 3))
	//fmt.Printf("[CMD] RPOP: %s\n", cli.RPop("l_hello_1"))

	// set
	// tip: elements number < 5000
	/*fmt.Printf("[CMD] SADD: %s\n", cli.SAdd("s_hello_1", "world_1", "world_2", "world_3"))
	fmt.Printf("[CMD] SCARD: %s\n", cli.SCard("s_hello_1"))
	fmt.Printf("[CMD] SISMEMBER: %s\n", cli.SIsMember("s_hello_1", "world_3"))
	fmt.Printf("[CMD] SMEMBERS: %s\n", cli.SMembers("s_hello_1"))
	fmt.Printf("[CMD] SRANDMEMBER: %s\n", cli.SRandMember("s_hello_1"))
	fmt.Printf("[CMD] SPOP: %s\n", cli.SPop("s_hello_1"))
	fmt.Printf("[CMD] SREM: %s\n", cli.SRem("s_hello_1", "world_3"))*/

	// sorted set
	// tip: elements number < 5000
	//for i := 0; i < 30; i++ {
	//	fmt.Printf("[CMD] ZADD: %s\n", cli.ZAdd("z_hello_1", redis.Z{Score: float64(i), Member: fmt.Sprintf("world_%d", i)}))
	//}
	//fmt.Printf("[CMD] ZCARD: %s\n", cli.ZCard("z_hello_1"))
	//fmt.Printf("[CMD] ZINCRBY: %s\n", cli.ZIncrBy("z_hello_1", 1.0, "world_1"))
	//fmt.Printf("[CMD] ZCOUNT: %s\n", cli.ZCount("z_hello_1", "-inf", "+inf"))
	//fmt.Printf("[CMD] ZRANGE: %s\n", cli.ZRange("z_hello_1", 3, 6))
	//fmt.Printf("[CMD] ZSCORE: %s\n", cli.ZScore("z_hello_1", "world_3"))
	//fmt.Printf("[CMD] ZREVRANGE: %s\n", cli.ZRevRangeWithScores("z_hello_1", 3, 6))
	//fmt.Printf("[CMD] ZRANGEBYSCORE: %s\n", cli.ZRangeByScore("z_hello_1", redis.ZRangeBy{Min: "0", Max: "0"}))
	//fmt.Printf("[CMD] ZRANK: %s\n", cli.ZRank("z_hello_1", "world_3"))
	//fmt.Printf("[CMD] ZREVRANK: %s\n", cli.ZRevRank("z_hello_1", "world_4"))
	//fmt.Printf("[CMD] ZREM: %s\n", cli.ZRem("z_hello_1", "world_4"))
	//fmt.Printf("[CMD] ZREMRANGEBYRANK: %s\n", cli.ZRemRangeByRank("z_hello_1", 0, 2))
	//fmt.Printf("[CMD] ZRANGEBYSCORE: %s\n", cli.ZRangeByScore("z_hello_1", redis.ZRangeBy{Min: "0", Max: "20"}))
	//fmt.Printf("[CMD] ZRANGEBYSCORE: %s\n", cli.ZRange("z_hello_1", -31,-1))
	//fmt.Printf("[CMD] ZREMRANGEBYSCORE: %s\n", cli.ZRemRangeByScore("z_hello_1", "0", "200"))

	// hyperloglog
	/*	fmt.Printf("[CMD] PFADD: %s\n", cli.PFAdd("databases", "Redis", "MongoDB", "MySQL"))
		fmt.Printf("[CMD] PFCOUNT: %s\n", cli.PFCount("databases"))*/

	// pipe
	// tip: put < 50 commands in one pipe
	/*pipe := cli.Pipeline()
	fmt.Printf("[CMD] PIPE: %s\n", pipe.Set("pipe_test", 0, 0))
	fmt.Printf("[CMD] PIPE: %s\n", pipe.Incr("pipe_test"))
	fmt.Printf("[CMD] PIPE: %s\n", pipe.IncrBy("pipe_test", 1))
	fmt.Printf("[CMD] PIPE: %s\n", pipe.IncrBy("pipe_test", 3))
	for i := 0; i < 50; i++ {
		pipe.Set(fmt.Sprintf("pipe_k_%d", i), fmt.Sprintf("pipe_v_%d", i), 0)
	}
	_, err = pipe.Exec()
	if err != nil {
		fmt.Printf("---pipe error:%s\n", err)
	}
	*/
	//pipe.Close()
	return

}

func main() {
	exampleCmd()
}
