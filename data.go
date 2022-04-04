package main

type MarketStatus uint

const (
	OpenBefore MarketStatus = 2
	Open       MarketStatus = 5
	CloseAfter MarketStatus = 6
	Closed     MarketStatus = 7
)

type Market struct {
	Region              string       `json:"region"`
	Status              MarketStatus `json:"status_id"`
	ZhStatusDescription string       `json:"status"`
}
type Quote struct {
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Symbol   string  `json:"symbol"`
	Current  float64 `json:"current"`
	Change   float64 `json:"chg"`
	Percent  float64 `json:"percent"`
	Exchange string  `json:"exchange"`
}
type Data struct {
	Market `json:"market"`
	Quote  `json:"quote"`
}

type TargetData struct {
	Code        string  `json:"code"`
	TargetPrice float64 `json:"target_price"`
}

func GetEnStatusDescription(status MarketStatus) string {
	switch status {
	case OpenBefore:
		return "Premarket"
	case Open:
		return "In transaction"
	case CloseAfter:
		return "After-hours trading"
	case Closed:
		return "Closed"
	}
	return ""
}
