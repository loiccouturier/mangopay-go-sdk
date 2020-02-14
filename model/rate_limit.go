package model

import "fmt"

type Rate uint8

const (
	RateLimit15Minutes Rate = iota + 1
	RateLimit30Minutes
	RateLimit60Minutes
	RateLimit24Mours
)

// RateLimit is in the header response of the Mangopay service ??
// it uses some duplicated datas fiedds that are in the header
// so we have to hijac the response before the map[string][]string get the Header values.
// ?? end ??
type RateLimit struct {
	Limit1          int `json:"X-RateLimit-1"`
	Limit2          int `json:"X-RateLimit-2"`
	Limit3          int `json:"X-RateLimit-3"`
	Limit4          int `json:"X-RateLimit-4"`
	LimitRemaining1 int `json:"X-RateLimit-Remaining-1"`
	LimitRemaining2 int `json:"X-RateLimit-Remaining-2"`
	LimitRemaining3 int `json:"X-RateLimit-Remaining-3"`
	LimitRemaining4 int `json:"X-RateLimit-Remaining-4"`
	LimitReset1     int `json:"X-RateLimit-Reset-1"`
	LimitReset2     int `json:"X-RateLimit-Reset-2"`
	LimitReset3     int `json:"X-RateLimit-Reset-3"`
	LimitReset4     int `json:"X-RateLimit-Reset-4"`
}

const (
	rateFormatData = "rateLimit: %v rateRemaining: %v rateReset: %v"
)

func (rl *RateLimit) GetData(r Rate) string {
	switch r {
	case RateLimit15Minutes:
		return fmt.Sprintf(rateFormatData, rl.Limit1, rl.LimitRemaining1, rl.LimitReset1)
	case RateLimit30Minutes:
		return fmt.Sprintf(rateFormatData, rl.Limit2, rl.LimitRemaining2, rl.LimitReset2)
	case RateLimit60Minutes:
		return fmt.Sprintf(rateFormatData, rl.Limit3, rl.LimitRemaining3, rl.LimitReset3)
	case RateLimit24Mours:
		return fmt.Sprintf(rateFormatData, rl.Limit4, rl.LimitRemaining4, rl.LimitReset4)
	default:
		return fmt.Sprintf(rateFormatData, rl.Limit1, rl.LimitRemaining1, rl.LimitReset1)
	}
}
