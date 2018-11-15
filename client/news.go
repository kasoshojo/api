// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "samclick": news Resource Client
//
// Command:
// $ goagen
// --design=github.com/kasoshojo/api/design
// --out=$(GOPATH)src/github.com/kasoshojo/api
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListNewsPath computes a request path to the list action of news.
func ListNewsPath() string {

	return fmt.Sprintf("/news/")
}

// Get news list
func (c *Client) ListNews(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListNewsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListNewsRequest create the request corresponding to the list action endpoint of the news resource.
func (c *Client) NewListNewsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
