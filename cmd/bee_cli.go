package cmd

import (
	"fmt"
	"strings"

	"github.com/nexclipper/curlbee/pkg/client"
	"github.com/nexclipper/curlbee/pkg/policy"
)

type CLIBee struct {
	params string
}

func (c *CLIBee) Run(cfg []policy.BeePolicy) error {
	respBuf := make(map[string]string)

	parameters := c.splitParameter(c.params)

	for _, p := range cfg {
		var name, respBody string
		p.VariableMatching(parameters)
		err := client.Request(&p, &name, &respBody)
		if err != nil {
			return err
		} else {
			respBuf[name] = respBody
		}

	}

	for k, v := range respBuf {
		fmt.Printf("%s\n\n%s\n\n", k, v)
	}

	return nil
}

func (c *CLIBee) splitParameter(p string) map[string]string {
	parameters := make(map[string]string)
	if len(c.params) > 0 {
		splits := strings.Split(c.params, ",")
		for _, s := range splits {
			kv := strings.Split(s, "=")
			if len(kv) == 2 {
				parameters[kv[0]] = kv[1]
			}
		}

		return parameters
	}

	return nil
}
