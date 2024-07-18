package RinGo

import (
	"github.com/imroc/req/v3"
)

const RINGO_URL = "https://www.ringodoor.com"

type sessionImpl struct {
	RingoHash string
	Client    *req.Client
}

type Session interface {
	DoorUnlock(lockId int, relayId *int, pin *int) (err error)
}
