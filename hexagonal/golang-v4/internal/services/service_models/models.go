package service_models

import (
	"time"
)

type Account struct {
	IMSI      string
	CreatedAt time.Time
}
