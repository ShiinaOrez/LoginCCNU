package spoc

import (
	"time"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

func NewClient() (*http.Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
		Jar:     jar,
	}
	return &client, nil
}
