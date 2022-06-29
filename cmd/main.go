package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ziouf/kubeconfig-merger/api/v1"

	"gopkg.in/yaml.v3"
)

func main() {
	flag.Parse()

	output := v1.Config{
		Kind:       "Config",
		ApiVersion: "v1",
	}

	for _, file := range flag.Args() {
		dat, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		var config v1.Config
		if err := yaml.Unmarshal(dat, &config); err != nil {
			panic(err)
		}

		output = output.Merge(config)
	}

	outputYaml, err := yaml.Marshal(output)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(outputYaml))
}
