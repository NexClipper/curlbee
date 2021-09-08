package cmd

import (
	"strings"

	"github.com/nexclipper/curlbee/pkg/policy"
)

type curlbee interface {
	Run(cfg []policy.BeePolicy) error
}

func NewBee(execType string) curlbee {
	execType = strings.ToUpper(execType)

	if execType == "HTTP" {
		return &HttpBee{handler: &BeeHandler{}}
	} else if execType == "CLI" {
		return &CLIBee{}
	}

	return nil
}
