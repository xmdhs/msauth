package auth

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/xmdhs/lorca"
)

func Getcode(dir string) (string, error) {
	if path := lorca.LocateChrome(); path == "" {
		return "", ErrNotInstallChrome
	}
	ui, err := newUI(dir)
	if err != nil {
		return "", fmt.Errorf("Getcode: %w", err)
	}
	defer ui.Close()
	code, err := getCode(ui)
	if err != nil {
		return "", fmt.Errorf("Getcode: %w", err)
	}
	return code, nil
}

func newUI(dir string) (lorca.UI, error) {
	ui, err := lorca.New(oauthURL, dir, 800, 600)
	if err != nil {
		return nil, fmt.Errorf("newUI: %w", err)
	}
	return ui, nil
}

var ErrNotInstallChrome = errors.New("ErrNotInstallChrome")

var whitelist = []string{"login.live.com", "github.com", "login.microsoft.com", ""}

var (
	ErrHostname       = errors.New("ErrHostname")
	ErrChromeWasClose = errors.New("ErrChromeWasClose")
)

func getCode(ui lorca.UI) (string, error) {
	for {
		aurl := ui.Eval(`window.location.href`).String()
		u, err := url.Parse(aurl)
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
			select {
			case <-ui.Done():
				return "", ErrChromeWasClose
			default:
				continue
			}
		}
		return code, nil
	}
}

const oauthURL = `https://login.live.com/oauth20_authorize.srf?client_id=00000000402b5328&response_type=code&scope=service%3A%3Auser.auth.xboxlive.com%3A%3AMBI_SSL&redirect_uri=https%3A%2F%2Flogin.live.com%2Foauth20_desktop.srf`
