package provider

import "net/rpc"

type Provider struct {
	Client *rpc.Client
	URL    string
}

var Providers map[string]*Provider
