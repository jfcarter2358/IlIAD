package main

import (
	_ "embed"
)

//go:embed s3.tf
var BucketTemplate string

type Bucket struct {
	Name string
	Tags map[string]string
}

func main() {

}
