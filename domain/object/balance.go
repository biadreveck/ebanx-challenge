package object

type Balance struct {
	AccountId string  `json:"id"`
	Balance   float64 `json:"balance"`
}

type DepositBalance struct {
	Destination *Balance `json:"destination"`
}

type WithdrawBalance struct {
	Origin *Balance `json:"origin"`
}

type TransferBalance struct {
	WithdrawBalance
	DepositBalance
}
