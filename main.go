package main

import (
	"fmt"

	"github.com/edmontongo/darksky/darksky"
)

func main() {
	const secretKey = "0123456789abcdef9876543210fedcba"
	ds := darksky.New(secretKey)

	edmonton := darksky.Location{Lat: 53.5458874, Long: -113.5034304}
	w, err := ds.Forecast(edmonton)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", w)
}
