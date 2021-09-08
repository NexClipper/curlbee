package cmd

import (
	"fmt"

	"github.com/nexclipper/curlbee/pkg/client"
	"github.com/nexclipper/curlbee/pkg/policy"
)

type CLIBee struct {
}

func (c *CLIBee) Run(cfg []policy.BeePolicy) error {
	respBuf := make(map[string]string)

	for _, p := range cfg {
		var name, respBody string
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
