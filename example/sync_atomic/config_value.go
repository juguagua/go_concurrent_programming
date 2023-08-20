package sync_atomic

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func loadNewConfig() Config { // 创造一个 new config
	return Config{
		NodeName: "beijing",
		Addr:     "10.58.144.186",
		Count:    rand.Int31(),
	}
}

func ChangeConfig() {
	var config atomic.Value                // Value 类型存储配置
	config.Store(loadNewConfig())          // Value 类型的 config 存储新的配置
	var cond = sync.NewCond(&sync.Mutex{}) // new 一个 cond 用于通知

	// 设置新的 config
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(5+rand.Int63n(5))) // 随机休眠 5-10 秒
			c := loadNewConfig()
			config.Store(c)                       // 更新配置
			println("Write New Config:", c.Count) // 打印配置
			cond.Broadcast()                      // 通知等待者 配置已经变更
		}
	}()

	go func() {
		for {
			cond.L.Lock()                        // 加锁
			cond.Wait()                          // 等待配置变更的通知
			c := config.Load().(Config)          // 读取新的配置
			cond.L.Unlock()                      // 解锁
			println("Read New Config:", c.Count) // 打印配置
		}
	}()

	select {} // 阻塞主 goroutine
}
