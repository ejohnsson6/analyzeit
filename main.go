package main

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	b, err := ioutil.ReadFile("rep.xls")

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	f, err := charmap.ISO8859_1.NewDecoder().String(str)


}
