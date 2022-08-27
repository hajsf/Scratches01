package structs

import (
	"time"

	"go.mau.fi/whatsmeow"
)

type LoginResponse struct {
	ImagePath string           `json:"image_path"`
	Duration  time.Duration    `json:"duration"`
	Code      string           `json:"code"`
	Client    whatsmeow.Client `json:"client"`
}
