package main

import (
	"log"
)

type PreviewResponse struct {
	message string
	status string
	code int
	file []byte //large??
}

/*
the point of using pointers is to have a memory address or reference of the variable
so that it can be useful when modifying or updating it in the entire memory and not just
in the scope of a function.
when
1.when we want to update a state
2. pointer -> 8bytes (without pointers(*) it is going to be the number of bytes)
when we want to opimitize the memory for large objects that are getting called a lot
3. 
*/

/*  -> 8bytes */
func (p *PreviewResponse) Message() string {
	p.message = "your request was successful"
	return p.message
}

func (p *PreviewResponse) GetResponse () (*PreviewResponse, error) {
	return &PreviewResponse{
		message: "here is your response",
	}, nil
}

func main () {
	pR := PreviewResponse{
		message: "waiting for a request",
	}
	log.Printf("c %v", pR.Message())
	log.Println(pR)
	p, err := pR.GetResponse()
	if(err != nil) {
		log.Println(err)
		return
	}
	log.Printf("PreviewResponse message - %v", p.Message())
}