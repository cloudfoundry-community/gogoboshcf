package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"launchpad.net/goyaml"

	"github.com/cloudfoundry-community/gogobosh/models"
	"github.com/cloudfoundry-community/gogoboshcf"
	"github.com/kr/pretty"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Println("USAGE: go run extractmanifest.go MANIFEST_PATH")
		return
	}
	manifestPath := flag.Arg(0)

	rawManifest, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	manifest := &models.DeploymentManifest{}
	goyaml.Unmarshal([]byte(rawManifest), manifest)

	cfManifest := gogoboshcf.CFDeploymentManifest(*manifest)
	pm, err := cfManifest.GlobalProperties()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%# v\n", pretty.Formatter(*pm))
}
