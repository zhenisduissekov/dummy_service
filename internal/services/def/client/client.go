package client

import "net/http"

type Client struct {
	Url    string
	User   string
	Pass   string
	Client *http.Client
}

func New(url, user, pass string) *Client {
	client := &http.Client{}

	return &Client{
		Url:    url,
		User:   user,
		Pass:   pass,
		Client: client,
	}
}
