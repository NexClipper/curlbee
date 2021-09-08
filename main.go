package main

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/nexclipper/curlbee/cmd"
	"github.com/nexclipper/curlbee/pkg/policy"
)

func main() {
	fileName := flag.String("filename", "", "yaml file name")
	execType := flag.String("type", "cli", "specifies the type of execution process[cli(default)|http]")
	flag.Parse()

	//*fileName = "./example/policy2.yml"
	//fmt.Println(*fileName)
	yamlBuf, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	cfg := make([]policy.BeePolicy, 0)
	err = yaml.Unmarshal(yamlBuf, &cfg)
	if err != nil {
		panic(err)
	}

	bee := cmd.NewBee(*execType)
	if bee != nil {
		err := bee.Run(cfg)
		if err != nil {
			panic(err)
		}
	} else {
		log.Println("type is not valid")
	}
}
