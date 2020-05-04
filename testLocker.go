package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg sync.WaitGroup
func main() {
	rwMutex = new(sync.RWMutex)
	wg.Add(2)

	go readData1()
	//go readData2()
	time.Sleep(1*time.Second)
	go writeData()
	wg.Wait()
}

func readData1() {
	defer wg.Done()
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	fmt.Println("读锁------")
	time.Sleep(3*time.Second)
}

func readData2(){
	defer wg.Done()
	Lock := rwMutex.RLocker()
	Lock.Lock()
	defer Lock.Unlock()
	fmt.Println("RLocker------")
	time.Sleep(3*time.Second)
}

func writeData(){
	defer wg.Done()
	rwMutex.Lock()
	defer rwMutex.Unlock()
	fmt.Println("写锁------")

}