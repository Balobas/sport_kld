package main

import (
	"../app/controllers/place"
	"fmt"
)

func main() {

	p, err := place.GetPlaceByUID("35cc8cde5f0111eb960b0c9d92446328")

	fmt.Println(p, err)
}

