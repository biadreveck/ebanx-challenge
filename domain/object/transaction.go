package object

type Transaction struct {
	AccountId string
	Amount    float64
}

type TransferTransaction struct {
	OriginAccountId string
	Transaction
}
