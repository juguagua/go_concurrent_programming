package sync_map

import (
	"hash/crc32"
	"sync"
)

var SHARD_COUNT = 32 // 默认使用 32 个分片

// ConcurrentMapShared ：通过RWMutex保护的线程安全的分片，包含一个 map
type ConcurrentMapShared struct {
	items        map[string]interface{}
	sync.RWMutex // guard access to internal map
}

// ConcurrentMap ：分成SHARD_COUNT个分片的map
type ConcurrentMap []*ConcurrentMapShared

// New : 创建并发map
func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}

// GetShard : 根据key计算分片索引
func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	// 根据key计算出相应的分片
	shard := m.GetShard(key)
	shard.Lock() // 对这个分片加锁执行业务操作
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	// 根据key计算出相应的分片
	shard := m.GetShard(key)
	shard.RLock() // 加读锁
	// 从这个分片读取key的值
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

func fnv32(key string) uint32 {
	res := crc32.ChecksumIEEE([]byte(key))
	return res
}
