package auth

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func newBrowser(headless bool) (*rod.Browser, error) {
	l := launcher.New().Headless(headless)
	url, err := l.Launch()
	if err != nil {
		return nil, fmt.Errorf("newBrowser: %w", err)
	}
	b := rod.New().ControlURL(url)
	err = b.Connect()
	if err != nil {
		return nil, fmt.Errorf("newBrowser: %w", err)
	}
	return b, nil
}
