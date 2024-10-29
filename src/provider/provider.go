package provider

import (
	"bytes"
	_ "embed"
	"fmt"
	"iad/config"
	"os"
	"os/exec"
	"strings"
)

//go:embed structure.sql
var INIT_SQL string

type Provider struct {
	Dependencies []string `json:"dependencies" yaml:"dependencies"`
	Language     string   `json:"language" yaml:"language"`
	Version      string   `json:"version" yaml:"version"`
}

func (p *Provider) Get(source string) error {
	source = strings.Trim(source, " ")
	if strings.HasPrefix(source, "github.com") {
		gitParts := strings.Split(source, "@")
		if len(gitParts) != 2 {
			return fmt.Errorf("github URL is of invalid format, should be 'github.com/<org>/<repo>/<path>@<ref>'")
		}
		urlParts := strings.Split(gitParts[0], "/")
		if len(urlParts) < 3 {
			return fmt.Errorf("github URL is of invalid format, should be 'github.com/<org>/<repo>/<path>@<ref>'")
		}

		ref := gitParts[1]
		org := urlParts[1]
		repo := urlParts[2]
		name := repo
		path := ""
		if len(urlParts) > 3 {
			path = strings.Join(urlParts[3:len(urlParts)-1], "/")
			name = urlParts[len(urlParts)-1]
		}

		providerPath := fmt.Sprintf("%s/%s", config.Config.ProviderPath, name)

		tmpDir, err := os.MkdirTemp("", "illiad_provider")
		if err != nil {
			return err
		}

		cmd := exec.Command("git", "clone", "-b", ref, fmt.Sprintf("git@github.com:%s/%s", org, repo), tmpDir)
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		if err := cmd.Run(); err != nil {
			return err
		}
	}
}
