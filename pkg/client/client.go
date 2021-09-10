package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nexclipper/curlbee/pkg/config"
)

func Request(p *config.BeePolicy, name *string, body *string) error {
	method := p.Request.Method
	method = strings.ToUpper(method)
	//if !(method == "GET" || method == "POST" || method == "DELETE" || method == "PUT") {
	if !(method == "GET") {
		return fmt.Errorf("Method type is not valid")
	}
	req, err := http.NewRequest(method, p.Request.URL, nil)
	if err != nil {
		return err
	}

	if len(p.Request.Headers) > 0 {
		for _, h := range p.Request.Headers {
			req.Header.Add(h.Key, h.Value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)

	*name = p.Name
	*body = str

	return nil
}
