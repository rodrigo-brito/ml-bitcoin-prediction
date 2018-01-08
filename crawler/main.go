package main

import (
	"log"
)

func main() {
	err := RunCripto()
	if err != nil {
		log.Println(err)
	}
	err = RunMarket()
	if err != nil {
		log.Println(err)
	}
}
