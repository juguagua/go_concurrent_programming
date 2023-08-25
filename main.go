package main

import "go_concurrent_programming/example/sync_channel"

func main() {
	//ConcurrentAdd()
	//Try()
	//GetState()
	//waitGroupCounter()
	//cond.ReadyRun()
	//sync_pool.ExecWithPool()
	//sync_pool.ExecWithoutPool()

	//var c uint32 = 5
	//x := c - 1
	//println(x)
	//println(^uint32(x) + c - 1)
	//atomic.AddUint32(&x, ^uint32(0))
	//println(^uint32(0))
	//
	//println(x)

	//sync_atomic.ChangeConfig()
	//sync_channel.FourChanSchedule()
	sync_channel.ChanSchedule()
}
