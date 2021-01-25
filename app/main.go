package main

import (
	 "../core/actions"
	"../models"
	"fmt"
)

func main() {

	place := models.Place{

	}

	p := &place
	actions.ActionGet("b2690eaf-0b73-4100-8a14-ebac27ae027c", p)
	fmt.Println(*p)
}

