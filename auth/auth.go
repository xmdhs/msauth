package auth

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/go-rod/bypass"
	"github.com/go-rod/rod"
)

func newPage(b *rod.Browser, url string) (*rod.Page, error) {
	page, err := bypass.Page(b)
	if err != nil {
		return nil, fmt.Errorf("newPage: %w", err)
	}
	err = page.Navigate(url)
	if err != nil {
		return nil, fmt.Errorf("newPage: %w", err)
	}
	return page, nil
}

func getCode(page *rod.Page) (string, error) {
	for {
		info, err := page.Info()
		if err != nil {
			return "", fmt.Errorf("getCode: %w", err)
		}
		u, err := url.Parse(info.URL)
		if err != nil {
			return "", fmt.Errorf("getCode: %w", err)
		}
		pass := false
		for _, v := range whitelist {
			if v == u.Hostname() {
				pass = true
			}
		}
		if !pass {
			return "", ErrHostname
		}
		code := u.Query().Get("code")
		if code == "" {
			time.Sleep(1 * time.Second)
			continue
		}
		return code, nil
	}
}

var whitelist = []string{"login.live.com", "github.com", "login.microsoft.com"}

var (
	ErrHostname = errors.New("ErrHostname")
)

func Getcode() (string, error) {
	b, err := newBrowser(false)
	defer b.Close()
	if err != nil {
		return "", fmt.Errorf("Getcode: %w", err)
	}
	page, err := newPage(b, oauthURL)
	defer page.Close()
	if err != nil {
		return "", fmt.Errorf("Getcode: %w", err)
	}
	code, err := getCode(page)
	if err != nil {
		return "", fmt.Errorf("Getcode: %w", err)
	}
	return code, nil
}

const oauthURL = `https://login.live.com/oauth20_authorize.srf?client_id=00000000402b5328&response_type=code&scope=service%3A%3Auser.auth.xboxlive.com%3A%3AMBI_SSL&redirect_uri=https%3A%2F%2Flogin.live.com%2Foauth20_desktop.srf`
