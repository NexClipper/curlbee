package config

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type BeeConfig struct {
	Title    string      `yaml:"title"`
	Port     uint32      `yaml:"port"`
	Policies []BeePolicy `yaml:"policies`
}

type BeePolicy struct {
	Name    string     `yaml:"name"`
	Request BeeRequest `yaml:"request"`
}

type BeeRequest struct {
	Method  string      `yaml:"method"`
	URL     string      `yaml:"url"`
	Headers []BeeHeader `yaml:"headers"`
}

type BeeHeader struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

func (b *BeePolicy) VariableMatching(params map[string]string) error {
	// match type: {{param.GROUP_ID}} or {{env.AUTHKEY}}
	r := regexp.MustCompile("{{2}(param|env).[a-zA-Z0-9_]*}{2}")
	m := r.MatchString(b.Name)
	if m == true {
		// replace
		matches := r.FindAll([]byte(b.Name), -1)
		for _, m := range matches {
			name := string(m)
			res, err := b.matching(name, params)
			if err != nil {
				return err
			}

			// there may be no parameter value
			if len(res) > 0 {
				b.Name = strings.Replace(b.Name, name, res, -1)
			}
		}
	}

	m = r.MatchString(b.Request.Method)
	if m == true {
		// replace
		matches := r.FindAll([]byte(b.Request.Method), -1)
		for _, m := range matches {
			method := string(m)
			res, err := b.matching(method, params)
			if err != nil {
				return err
			}

			// there may be no parameter value
			if len(res) > 0 {
				b.Name = strings.Replace(b.Request.Method, method, res, -1)
			}
		}
	}

	m = r.MatchString(b.Request.URL)
	if m == true {
		// replace
		matches := r.FindAll([]byte(b.Request.URL), -1)
		for _, m := range matches {
			url := string(m)
			res, err := b.matching(url, params)
			if err != nil {
				return err
			}

			// there may be no parameter value
			if len(res) > 0 {
				b.Request.URL = strings.Replace(b.Request.URL, url, res, -1)
			}
		}
	}

	for i, h := range b.Request.Headers {
		m = r.MatchString(h.Key)
		if m == true {
			// replace
			matches := r.FindAll([]byte(b.Request.Headers[i].Key), -1)
			for _, m := range matches {
				key := string(m)
				res, err := b.matching(key, params)
				if err != nil {
					return err
				}

				// there may be no parameter value
				if len(res) > 0 {
					b.Request.Headers[i].Key = strings.Replace(b.Request.Headers[i].Key, key, res, -1)
				}
			}
		}
		m = r.MatchString(h.Value)
		if m == true {
			// replace
			matches := r.FindAll([]byte(b.Request.Headers[i].Value), -1)
			for _, m := range matches {
				value := string(m)
				res, err := b.matching(value, params)
				if err != nil {
					return err
				}

				// there may be no parameter value
				if len(res) > 0 {
					b.Request.Headers[i].Value = strings.Replace(b.Request.Headers[i].Value, value, res, -1)
				}
			}
		}
	}

	log.Printf("%+v", b)

	return nil
}

func (b *BeePolicy) matching(v string, p map[string]string) (string, error) {
	tmplUnit := strings.TrimRight(strings.TrimLeft(v, "{{"), "}}")
	splits := strings.Split(tmplUnit, ".")
	var paramValue string

	if len(splits) == 2 {
		paramType := strings.ToUpper(splits[0])
		paramKey := strings.ToUpper(splits[1])
		if paramType == "ENV" {
			paramValue = os.Getenv(paramKey)
		} else if paramType == "PARAM" {
			paramValue = p[paramKey]
		} else {
			return "", fmt.Errorf("check the paramter type")
		}
	} else {
		return "", fmt.Errorf("check the parameter format")
	}

	return paramValue, nil
}
