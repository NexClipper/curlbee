package cmd

import (
	"strings"

	"github.com/nexclipper/CurlBee/pkg/policy"
)

type CurlBee interface {
	Run(cfg []policy.BeePolicy) error
}

func NewBee(execType string) CurlBee {
	execType = strings.ToUpper(execType)

	if execType == "HTTP" {
		return &HttpBee{handler: &BeeHandler{}}
	} else if execType == "CLI" {
		return &CLIBee{}
	}

	return nil
}
