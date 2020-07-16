package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	myUUID, _ := uuid.NewRandom()
	fmt.Println(quote.Hello(), myUUID)
}
