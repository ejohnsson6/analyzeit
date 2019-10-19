package verloader

type Verification struct {
	Vernr string
	Date string
	Description string
	Verlines []VerificationLine
}

type VerificationLine struct {
	Account int
	AccountDescription string
	Costplace int
	Credit int
	Debit int
}