package helpers

import (
	"time"

	"github.com/RegiAdi/pos-mobile-backend/config"
)

func GetCurrentTime() time.Time {
	loc, _ := time.LoadLocation(config.GetAppTimezone())

	return time.Now().In(loc)
}