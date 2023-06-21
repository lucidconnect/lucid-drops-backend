package utils

import (
	"net/http"
	"net/url"

	"github.com/machinebox/graphql"
)

func GetDebugClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse("http://192.168.100.167:9090") //this sshould be dynamic based on the proxyman url
			},
		},
	}
}

func GetGraphQLDebugClient(endpoint string) *graphql.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse("http://192.168.100.167:9090") //this sshould be dynamic based on the proxyman url
			},
		},
	}
	client := graphql.NewClient(endpoint, graphql.WithHTTPClient(httpClient))
	
	return client
}
