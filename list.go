package main

import (
	"io/ioutil"
)

func GetDisposableDomainList() []byte {
	data, err := ioutil.ReadFile("./list.txt")

	if err != nil {
		panic(err)
	}

	return data
}
