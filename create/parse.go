package create

import (
	"errors"
	"net/url"
	"strings"
)

var parseFuncMap parseFuncManage

type parseFuncManage map[string]func(u *url.URL) (string, error)

func (m *parseFuncManage) registerFunc(h string, f func(u *url.URL) (string, error)) {
	(*m)[h] = f
}

func (m *parseFuncManage) parse(u *url.URL) (string, error) {
	pf, ok := (*m)[u.Hostname()]
	if !ok {
		return "", errors.New("target host is not supported")
	}

	res, err := pf(u)
	if err != nil {
		return "", errors.New("invalid url")
	}
	return res, nil
}

func init() {
	parseFuncMap = parseFuncManage{}

	parseFuncMap.registerFunc(
		"atcoder.jp",
		func(u *url.URL) (string, error) {
			p := strings.Split(u.Path, "/")
			if p[len(p)-2] != "tasks" {
				return "", errors.New("invalid url")
			}
			return p[len(p)-1], nil
		},
	)

	parseFuncMap.registerFunc(
		"judge.u-aizu.ac.jp",
		func(u *url.URL) (string, error) {
			id := u.Query().Get("id")
			if id == "" {
				return "", errors.New("invalid id")
			}
			return id, nil
		},
	)
}

func ParseIDFromURL(target string) (string, error) {
	u, err := url.Parse(target)
	if err != nil {
		return "", err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return "", errors.New("invalid URL")
	}

	problemID, err := parseFuncMap.parse(u)
	if err != nil {
		return "", err
	}

	return problemID, nil
}
