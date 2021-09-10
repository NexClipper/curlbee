package cmd

import (
	"fmt"

	"github.com/nexclipper/curlbee/pkg/client"
	"github.com/nexclipper/curlbee/pkg/config"
	"github.com/nexclipper/curlbee/pkg/util"
)

type CLIBee struct {
	params string
}

func (c *CLIBee) Run(cfg *config.BeeConfig) error {
	respBuf := make(map[string]string)

	parameters := util.SplitParameter(c.params)

	for _, p := range cfg.Policies {
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
