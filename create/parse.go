package create

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var parseFuncMap parseFuncManage

var (
	ErrNotSupportedHost = errors.New("not supported host")
	ErrInvalidURL       = errors.New("invalid url")
)

type parseFuncManage map[string]func(u *url.URL) (string, error)

func (m *parseFuncManage) registerFunc(host string, f func(u *url.URL) (string, error)) {
	(*m)[host] = f
}

func (m *parseFuncManage) parse(target *url.URL) (string, string, error) {
	host := target.Hostname()

	prseFunc, ok := (*m)[host]
	if !ok {
		return "", "", fmt.Errorf("%w: %s", ErrNotSupportedHost, target.Hostname())
	}

	res, err := prseFunc(target)
	if err != nil {
		return "", "", fmt.Errorf("%w: %s", err, target.String())
	}

	return res, host, nil
}

func init() {
	parseFuncMap = parseFuncManage{}

	parseFuncMap.registerFunc(
		"atcoder.jp",
		func(u *url.URL) (string, error) {
			p := strings.Split(u.Path, "/")
			if p[len(p)-2] != "tasks" {
				return "", fmt.Errorf("%w: can't find tasks", ErrInvalidURL)
			}

			return p[len(p)-1], nil
		},
	)

	parseFuncMap.registerFunc(
		"judge.u-aizu.ac.jp",
		func(u *url.URL) (string, error) {
			id := u.Query().Get("id")
			if id == "" {
				return "", fmt.Errorf("%w: can't find id", ErrInvalidURL)
			}

			return id, nil
		},
	)
}

func ParseIDFromURL(target string) (string, string, error) {
	parsed, err := url.Parse(target)
	if err != nil {
		return "", "", fmt.Errorf("%w: %w", ErrInvalidURL, err)
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", "", fmt.Errorf("%w: scheme must be 'http' or 'https'", ErrInvalidURL)
	}

	problemID, host, err := parseFuncMap.parse(parsed)
	if err != nil {
		return "", "", err
	}

	return problemID, host, nil
}
