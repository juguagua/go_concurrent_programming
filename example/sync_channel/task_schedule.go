package sync_channel

import (
	"fmt"
	"time"
)

func FourChanSchedule() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	ch4 := make(chan struct{})

	go func() {
		for {
			<-ch1
			print("1", " ")
			time.Sleep(time.Second)
			ch2 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch2
			print("2", " ")
			time.Sleep(time.Second)
			ch3 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch3
			print("3", " ")
			time.Sleep(time.Second)
			ch4 <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch4
			print("4")
			println()
			time.Sleep(time.Second)
			ch1 <- struct{}{}
		}

	}()

	ch1 <- struct{}{}

	time.Sleep(time.Second * 20)
	//select {}  		// 阻塞 main
}

func ChanSchedule() {
	ch := make(chan int, 1)

	var worker = func(id int) {
		for {
			select {
			case n := <-ch:
				if n == id {
					fmt.Print(id, " ")
					time.Sleep(time.Second)
					if id == 4 {
						println()
						ch <- 1
					} else {
						ch <- id + 1
					}
				} else {
					ch <- n
				}
			}
		}
	}
	for i := 1; i <= 4; i++ {
		go worker(i)
	}
	ch <- 1

	time.Sleep(time.Second * 20)

}
