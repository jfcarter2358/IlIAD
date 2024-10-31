package s3

import (
	"bytes"
	_ "bytes"
	_ "embed"
	"encoding/json"
	"text/template"

	"github.com/jfcarter2358/gitdb"
	"github.com/jfcarter2358/go-logger"
)

//go:embed s3.tf
var bucketTemplate string

type Bucket struct {
	Name string `json:"name" gitdb_meta:"true"`
	TF   string `json:"tf"`
}

func Post(obj map[string]interface{}, repo gitdb.Repo, path string) error {
	if err := repo.Init(); err != nil {
		return err
	}
	repo.Path = path
	tmpl, err := template.New("s3").Parse(bucketTemplate)
	if err != nil {
		return err
	}

	logger.Infof("", "Starting repo get")
	logger.Debugf("", "%v", obj)
	logger.Debugf("", path)

	var buckets []Bucket
	dat, err := repo.Get(path)
	if err != nil {
		logger.Warnf("", "Error on get: %s", err.Error())
	} else {
		var b Bucket
		if err := gitdb.Unmarshal(dat, &b); err != nil {
			if err := gitdb.Unmarshal(dat, &buckets); err != nil {
				logger.Warnf("", "Unable to get documents: %s", err.Error())
			}
		} else {
			buckets = []Bucket{b}
		}
	}

	logger.Debugf("", "Marshalling bucket")

	var b Bucket
	bs, err := json.Marshal(obj)
	if err != nil {
		logger.Errorf("", "Unable to marshal %v", obj)
		return err
	}
	if err := json.Unmarshal(bs, &b); err != nil {
		logger.Errorf("", "Unable to unmarshal %s", string(bs))
		return err
	}

	logger.Debugf("", "Marshalled")

	logger.Debugf("", "Rendering template")

	var tpl bytes.Buffer
	if err = tmpl.Execute(&tpl, b); err != nil {
		return err
	}
	b.TF = tpl.String()
	buckets = append(buckets, b)

	logger.Debugf("", "Rendered")

	logger.Debugf("", "%s", repo.LocalDir)

	out := ""
	for _, bucket := range buckets {
		bytes, err := gitdb.Marshal(bucket)
		if err != nil {
			return err
		}
		out += string(bytes)
	}
	if err := repo.Post([]byte(out), path); err != nil {
		return err
	}
	return nil
}

// func Post(args *contract.Comm, resp *contract.Comm) error {
// 	tmpl, err := template.New("s3").Parse(bucketTemplate)
// 	if err != nil {
// 		return err
// 	}

// 	logger.Infof("", "Starting S3 post")

// 	repo := gitdb.Repo{
// 		URL:    config.Config.GitURL,
// 		Ref:    config.Config.GitRef,
// 		Branch: config.Config.GitBranch,
// 		Path:   config.Config.S3Path,
// 	}
// 	if err := repo.Init(); err != nil {
// 		return err
// 	}

// 	logger.Infof("", "Starting repo get")

// 	var buckets []Bucket
// 	dat, err := repo.Get()
// 	if err != nil {
// 		logger.Warnf("", "Error on get: %s", err.Error())
// 	} else {
// 		var b Bucket
// 		if err := gitdb.Unmarshal(dat, &b); err != nil {
// 			if err := gitdb.Unmarshal(dat, &buckets); err != nil {
// 				logger.Warnf("", "Unable to get documents: %s", err.Error())
// 			}
// 		} else {
// 			buckets = []Bucket{b}
// 		}
// 	}

// 	logger.Infof("", "%s", string(args.Body))

// 	var newBuckets []Bucket
// 	if err := json.Unmarshal(args.Body, &newBuckets); err != nil {
// 		return err
// 	}
// 	for _, b := range newBuckets {
// 		var tpl bytes.Buffer
// 		if err = tmpl.Execute(&tpl, b); err != nil {
// 			return err
// 		}
// 		b.TF = tpl.String()
// 		buckets = append(buckets, b)
// 	}

// 	out := ""
// 	for _, bucket := range buckets {
// 		bytes, err := gitdb.Marshal(bucket)
// 		if err != nil {
// 			return err
// 		}
// 		out += string(bytes)
// 	}
// 	logger.Debugf("", out)
// 	if err := repo.Post([]byte(out)); err != nil {
// 		return err
// 	}
// 	if err := repo.Push("Adding S3 bucket(s)"); err != nil {
// 		panic(err)
// 	}
// 	resp.StatusCode = http.StatusOK
// 	return nil
// }
