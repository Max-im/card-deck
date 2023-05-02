package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	value int
	name  string
}

func (c card) toString() string {
	return c.name + ";" + strconv.Itoa(c.value)
}

func toCard(s string) card {
	splittedStr := strings.Split(s, ";")
	val, err := strconv.Atoi(splittedStr[1])

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	return card{name: splittedStr[0], value: val}
}
