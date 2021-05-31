package main

import (
	"fmt"

	"github.com/melodie-r/shukujitsu-go"
)

// import 省略...
func main() {
	entries, err := shukujitsu.AllEntries()
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		fmt.Printf("%d/%d/%d = %s\n", e.Year, e.Month, e.Day, e.Name)
	}
}
