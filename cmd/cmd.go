package cmd

import (
	"strings"

	"github.com/nexclipper/curlbee/pkg/config"
)

type curlbee interface {
	Run(cfg *config.BeeConfig) error
}

func NewBee(execType, param string) curlbee {
	execType = strings.ToUpper(execType)

	if execType == "HTTP" {
		return &HttpBee{handler: &BeeHandler{}}
	} else if execType == "CLI" {
		return &CLIBee{params: param}
	}

	return nil
}
