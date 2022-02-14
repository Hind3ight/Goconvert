package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

func main() {
	path := "/home/hind3ight/Projects/GoProject/src/github.com/hind3ight/Goconvert/demo/1_sample_demo.json"
	ret, err := ioutil.ReadFile(path)
	y, err := yaml.JSONToYAML(ret)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))

	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j2))
}
