package main

import (
	"fmt"
	"log"

	"github.com/henryangulo/ts-go-rust/pkg/projector"
)

func main(){
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatalf("unale to get options %v", err)
	}
	fmt.Printf("opts: %+v\n", opts)
}