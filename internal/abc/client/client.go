package client

import (
	"github.com/zhenisduissekov/dummy_service/internal/tools"
	"net/http"
)

const (
	CurrencyType   = 000
	LanguageType   = ""
	PageView       = "DESKTOP" // DESKTOP/MOBILE
	WriteOffAmount = 1000
)

type Client struct {
	Addr string
	User string
	Pass string
	Http *http.Client
}

func New(addr, user, pass string, proxy string) *Client {
	pr := tools.NewProxy(proxy)
	return &Client{
		Addr: addr,
		User: user,
		Pass: pass,
		Http: pr,
	}
}
