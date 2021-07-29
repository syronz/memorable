package main

import (
	"fmt"
	"log"

	"github.com/syronz/memorable"
)

func main() {
	// mem, err := memorable.New().Chars("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789").Length(2).Build()
	// mem, err := memorable.New().Chars("ABCDEFGHJKLMNPQRSTUVWXYZ123456789").Length(6).Build()
	mem, err := memorable.New().Chars("AB").Length(2).Build()
	if err != nil {
		log.Fatal(err.Error())
	}

	variations := mem.Variation()

	for _, v := range variations {
		fmt.Printf("%4d - %v\n", v.Mark, v.Code)
	}

}
