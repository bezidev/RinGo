package RinGo

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"net/http"
	"regexp"
)

func NewSession(url string) (Session, error) {
	r := regexp.MustCompile("https://www.ringodoor.com/door/\\?hash=(?P<Hash>.*)")
	submatches := r.FindStringSubmatch(url)
	if len(submatches) < 2 {
		return nil, errors.New(fmt.Sprintf("Could not parse URL: %s", url))
	}
	client := req.C()
	client.Headers = make(http.Header)
	client.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	//client.DevMode()
	hash := submatches[1]
	return &sessionImpl{
		RingoHash: hash,
		Client:    client,
	}, nil
}
