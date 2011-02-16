package client

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/zhenisduissekov/dummy_service/internal/tools"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	CurrencyType   = 000
	LanguageType   = ""
	PageView       = "DESKTOP" // DESKTOP/MOBILE
	WriteOffAmount = 1000
)

type Client struct {
	Url  string
	User string
	Pass string
	Http *http.Client
}

func New(url, user, pass string, proxy string) *Client {
	pr := tools.NewProxy(proxy)
	return &Client{
		Url:  url,
		User: user,
		Pass: pass,
		Http: pr,
	}
}

type ResponseItems struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
}

func (c *Client) SendRequest() ([]ResponseItems, error) {
	var respItems []ResponseItems
	resp, err := c.Http.Get(c.Url)
	if err != nil {
		log.Err(err).Msg("error at Get")
		return []ResponseItems{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Err(err).Msg("error at body close")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("error at read body")
		return []ResponseItems{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []ResponseItems{}, errors.New("response status not Ok")
	}

	err = json.Unmarshal(body, &respItems)
	return respItems, err
}
