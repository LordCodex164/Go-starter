package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err error
}

func main () {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "foo", "bar")
	userId := 34
	val, err := FetchUserData(ctx, userId)
	if err != nil {
	log.Println(err)
	}
	log.Println(val, "value")
	log.Println("this time took us", time.Since(start))
}

func FetchUserData(ctx context.Context, userId int) (int, error) {
	ctxValue := ctx.Value("foo")
	log.Printf("here is the context value: %v", ctxValue)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 900)
	defer cancel()
	respch := make(chan Response)
	go func () {
	  val, err := fetchThirdPartyStuffSlow()
	  respch <- Response{
	  value: val,
	  err: err,
	 }
	}()

	for {
		select {
		case <- ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took to long")
		case resp := <- respch:
			return resp.value, resp.err
		}
	}

	// if err != nil {
	// 	return 0, err
	// }
	// return val, nil
}

func fetchThirdPartyStuffSlow() (int, error) {
	time.Sleep(time.Millisecond * 700)
	
	return 666, nil
}