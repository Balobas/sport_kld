package main

import (
	 "../core/actions"
	"../data"
	"fmt"
)

func main() {

	place := data.Place{

	}

	p := &place
	actions.ActionGet("b2690eaf-0b73-4100-8a14-ebac27ae027c", p)
	fmt.Println(*p)
}

