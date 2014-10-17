package main

import (
	"fmt"

	"github.com/cloudfoundry-community/gogobosh"
	"github.com/cloudfoundry-community/gogobosh/api"
	"github.com/cloudfoundry-community/gogobosh/local"
	"github.com/cloudfoundry-community/gogobosh/net"
	"github.com/cloudfoundry-community/gogobosh/utils"
)

func main() {
	utils.Logger = utils.NewLogger()

	configPath, err := local.DefaultBoshConfigPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	config, err := local.LoadBoshConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	target, username, password, err := config.CurrentBoshTarget()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Targeting %s with user %s...\n", target, username)

	director := gogobosh.NewDirector(target, username, password)
	repo := api.NewBoshDirectorRepository(&director, net.NewDirectorGateway())

	fmt.Println("Finding deployments using 'cf' release...")
	deployments, apiResponse := repo.GetDeployments()
	if apiResponse.IsNotSuccessful() {
		fmt.Println(apiResponse)
		return
	}
	cfDeployments := deployments.FindByRelease("cf")
	if len(cfDeployments) < 1 {
		fmt.Println("No deployments include 'cf' release.")
		return
	}
	deployment := cfDeployments[0]
	fmt.Printf("Deployments found %v, selecting %s\n", cfDeployments, deployment)

	fmt.Println("Fetching deployment manifest...")
	manifest, apiResponse := repo.GetDeploymentManifest(deployment.Name)
	if apiResponse.IsNotSuccessful() {
		fmt.Println(apiResponse)
		return
	}
	fmt.Printf("Manifest: %#v\n", manifest)
	// cfManifest := (*manifest).(gogoboshcf.CFDeploymentManifest)
}
