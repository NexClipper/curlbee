package main

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/nexclipper/curlbee/cmd"
	"github.com/nexclipper/curlbee/pkg/config"
)

func main() {
	fileName := flag.String("filename", "", "yaml file name")
	execType := flag.String("type", "cli", "specifies the type of execution process[cli(default)|http]")
	param := flag.String("param", "", "parameters can be transferred")
	flag.Parse()

	//*execType = "http"
	//*fileName = "./example/policy4.yml"
	//fmt.Println(*fileName)
	//fmt.Println(*param)
	yamlBuf, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	cfg := &config.BeeConfig{}
	err = yaml.Unmarshal(yamlBuf, cfg)
	if err != nil {
		panic(err)
	}

	bee := cmd.NewBee(*execType, *param)
	if bee != nil {
		err := bee.Run(cfg)
		if err != nil {
			panic(err)
		}
	} else {
		log.Println("type is not valid")
	}
}
