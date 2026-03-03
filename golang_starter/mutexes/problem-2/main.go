package main

import (
	"fmt"
	"sync"
	"time"
)

type APIRequests struct {
	Limit int64
	Requests int64
	//because of enscapsulation
	mutex sync.Mutex
	rMutex sync.RWMutex
}

func (b *APIRequests) Write(wg *sync.WaitGroup) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	defer wg.Done()
	if b.Requests == b.Limit {
		fmt.Println("Limit Exceeded")
		return
	}
	b.Requests ++
	fmt.Println("Requests Created", b.Requests)
	time.Sleep(3 * time.Second)
}

func (b *APIRequests) Read() {
	b.rMutex.RLock()
	defer b.rMutex.RUnlock()

	fmt.Println("APIRequests", b.Requests)
}

func main(){
	api_requests := APIRequests{
		Limit: 5,
		Requests: 0,
	}

	var wg sync.WaitGroup

	//make 3 requests
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go api_requests.Write(&wg)
	}

	wg.Wait()

	api_requests.Read()

	fmt.Println("Done")
}
