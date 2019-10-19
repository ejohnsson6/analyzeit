package verloader

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const verNr = 0
const date = 1
const account = 3
const accountDescription = 4
const costplace = 5
const verText = 7
const debit = 9
const credit = 10

// LoadVerificationsFromFile loads verifications from a file
func LoadVerificationsFromFile(file string) []Verification {
	b, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	f, err := charmap.ISO8859_1.NewDecoder().String(str)

	rows := strings.Split(f, "\n")

	verList := []Verification{}
	_vernr := ""

	var verification Verification

	for _, v := range rows {

		if v == "" {
			break
		}

		columns := strings.Split(v, "\t")

		if _vernr != columns[verNr] {
			if _vernr != "" {
				verList = append(verList, verification)
			}
			_vernr = columns[verNr]

			verification = Verification{
				Vernr:       columns[verNr],
				Date:        columns[date],
				Description: columns[verText],
				Verlines:    []VerificationLine{},
			}
		}

		_account, _ := strconv.Atoi(columns[account])
		_costplace, _ := strconv.Atoi(columns[costplace])
		_credit, _ := strconv.Atoi(strings.Replace(columns[credit], ",", "", -1))
		_debit, _ := strconv.Atoi(strings.Replace(columns[debit], ",", "", -1))

		verLine := VerificationLine{
			Account:            _account,
			AccountDescription: columns[accountDescription],
			Costplace:          _costplace,
			Credit:             _credit,
			Debit:              _debit,
		}

		verification.Verlines = append(verification.Verlines, verLine)

	}

	return (verList)

}
