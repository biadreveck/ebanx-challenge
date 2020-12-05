package payload

const (
	EVENT_TYPE_DEPOSIT  = "deposit"
	EVENT_TYPE_WITHDRAW = "withdraw"
	EVENT_TYPE_TRANSFER = "transfer"
)

type EventPayload struct {
	Type                 string  `json:"type"`
	Amount               float64 `json:"amount"`
	OriginAccountId      string  `json:"origin"`
	DestinationAccountId string  `json:"destination"`
}
