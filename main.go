package main

import (
	"analyzeit/verloader"
	"fmt"
	"strings"
)

func main() {

	verifications := verloader.LoadVerificationsFromFile("./rep.xls")

	// O(N^2) FTW

	notpaid := []verloader.Verification{}
	for _, element := range verifications {
		if strings.HasPrefix(element.Description, "MU") {

			if !findMatchingMB(element, verifications) {
				notpaid = append(notpaid, element)
			}

		}
	}

	for _, element := range notpaid {
		fmt.Println(element.Description)
	}

}

func findMatchingMB(verification verloader.Verification, verifications []verloader.Verification) bool {

	noPrefixDesc1 := removePrefix(verification.Description)
	matchingExists := false

	for _, element := range verifications {
		if element.Vernr == verification.Vernr {
			continue
		}

		noPrefixDesc2 := removePrefix(element.Description)
		if noPrefixDesc2 != noPrefixDesc1 {
			continue
		}

		matchingExists = true
		fmt.Printf("%s matches with %s \n", verification.Vernr, element.Vernr)
	}

	return (matchingExists)
}

func removePrefix(verification string) string {

	noPrefix := strings.Replace(verification, "FM", "", 1)
	noPrefix = strings.Replace(verification, "FB", "", 1)
	noPrefix = strings.Replace(verification, "MU", "", 1)
	noPrefix = strings.Replace(verification, "MB", "", 1)

	return noPrefix
}
