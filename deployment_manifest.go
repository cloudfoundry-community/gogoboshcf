package gogoboshcf

import "github.com/cloudfoundry-community/gogobosh/models"

// CFDeploymentManifest is a deployment manifest for a Cloud Foundry deployment
type CFDeploymentManifest models.DeploymentManifest

// FindNATSMachines discovers the hostnames/static IPs for the NATS servers
func (manifest *CFDeploymentManifest) FindNATSMachines() (hostnames []string) {
	if manifest.Properties == nil {
		return
	}
	properties := *manifest.Properties
	if properties["nats"] == nil {
		return
	}
	nats := properties["nats"].(map[string]interface{})
	if nats["machines"] == nil {
		return
	}

	return nats["machines"].([]string)
}
