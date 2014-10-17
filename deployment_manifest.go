package gogoboshcf

import "github.com/cloudfoundry-community/gogobosh/models"

// CFDeploymentManifest is a deployment manifest for a Cloud Foundry deployment
type CFDeploymentManifest models.DeploymentManifest

// NATS represents the NATS client credentials
type NATS struct {
	MachinesHostnames []string
	Port              int
	Username          string
	Password          string
}

// NATS discovers the hostnames/static IPs for the NATS servers
func (manifest *CFDeploymentManifest) NATS() (nats NATS) {
	if manifest.Properties == nil {
		return
	}
	properties := *manifest.Properties
	if properties["nats"] == nil {
		return
	}
	natsProperties := properties["nats"].(map[string]interface{})

	if natsProperties["machines"] != nil {
		nats.MachinesHostnames = natsProperties["machines"].([]string)
	}
	if natsProperties["machines"] != nil {
		nats.Port = natsProperties["port"].(int)
	} else {
		nats.Port = 4222
	}
	if natsProperties["username"] != nil {
		nats.Username = natsProperties["username"].(string)
	}
	if natsProperties["password"] != nil {
		nats.Password = natsProperties["password"].(string)
	}

	return
}
