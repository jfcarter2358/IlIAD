package main

import (
	"aws/config"
	"aws/contract"
	"aws/s3"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/jfcarter2358/gitdb"
	"github.com/jfcarter2358/go-logger"
)

type Provider int64

var MainRepo gitdb.Repo

func (p *Provider) Post(args *contract.Comm, resp *contract.Comm) error {
	var objs []map[string]interface{}
	if err := json.Unmarshal(args.Body, &objs); err != nil {
		return err
	}

	for _, obj := range objs {
		kind := obj["kind"].(string)
		switch kind {
		case "s3":
			if err := s3.Post(obj, MainRepo, config.Config.S3Path); err != nil {
				return err
			}
		}
	}

	if err := MainRepo.Push("Add infrastructure"); err != nil {
		return err
	}

	if err := MainRepo.PR("Infrastructure Addition", "Added some infrastructure"); err != nil {
		return err
	}

	resp.StatusCode = 200
	return nil
}

// func (p *Provider) Delete(args *contract.Comm, reply *contract.Comm) error {

// }

func main() {
	config.LoadConfig()
	logger.SetLevel(config.Config.LogLevel)
	logger.SetFormat(config.Config.LogFormat)

	MainRepo = gitdb.Repo{
		URL:    "git@github.com:jfcarter2358/iliad",
		Path:   "aws/aws.tf",
		Ref:    "main",
		Branch: "iliad_test",
	}

	// Create a new RPC server
	provider := new(Provider)

	// Register RPC server
	rpc.Register(provider)
	rpc.HandleHTTP()

	// Listen for requests on port 1234
	l, e := net.Listen("tcp", fmt.Sprintf(":%d", config.Config.Port))
	if e != nil {
		logger.Fatalf("", "listen error: %s", e.Error())
	}
	http.Serve(l, nil)
}
