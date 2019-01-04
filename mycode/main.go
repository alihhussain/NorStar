package main

import (
	"fmt"

	"./dog"
)

func main() {
	humanYears := 5
	dogYears := dog.Years(humanYears)

	fmt.Println(dogYears)

}
