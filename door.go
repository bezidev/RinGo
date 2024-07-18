package RinGo

import (
	"fmt"
	"net/url"
)

func (session *sessionImpl) DoorUnlock(lockId int, relayId *int, pin *int) (err error) {
	rid := 0
	if relayId != nil {
		rid = *relayId
	}
	p := ""
	if pin != nil {
		p = fmt.Sprint(*pin)
	}
	r := session.Client.R()
	r.FormData = url.Values{
		"action":   []string{"open_lock"},
		"hash":     []string{session.RingoHash},
		"pin":      []string{p},
		"relay_id": []string{fmt.Sprint(rid)},
		"lock_id":  []string{fmt.Sprint(lockId)},
	}
	_, err = r.Post(fmt.Sprintf("%s/ajax", RINGO_URL))
	return err
}
