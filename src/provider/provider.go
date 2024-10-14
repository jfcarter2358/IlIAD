package provider

type Provider struct {
	Dependencies []string `json:"dependencies" yaml:"dependencies"`
	Language     string   `json:"language" yaml:"language"`
}
