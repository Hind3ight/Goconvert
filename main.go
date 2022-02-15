package main

import (
	"flag"
	"fmt"
	"github.com/hind3ight/Goconvert/consts"
	fileUtils "github.com/hind3ight/Goconvert/pkg/lib/file"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/ghodss/yaml"
)

var (
	oldPath    string
	flagSet    *flag.FlagSet
	outPutPath string
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet("go-convert", flag.ContinueOnError)
	flagSet.StringVar(&outPutPath, "o", "", "")
	files := fileUtils.GetFilesFromParams(os.Args[2:])
	flagSet.Parse(os.Args[len(files)+2:])

	ret, err := ioutil.ReadFile(files[0])
	y, err := yaml.JSONToYAML(ret)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// create file
	resultDir := consts.WorkDir + outPutPath + consts.PthSep + consts.DefaultOutPutFile
	fileUtils.WriteFile(resultDir, string(y))

	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j2))
}

func init() {
	consts.WorkDir = fileUtils.GetWorkDir()
}
