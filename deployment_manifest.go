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

// UAA represents UAA admin client credentials
type UAA struct {
	URI               string
	AdminClientID     string
	AdminClientSecret string
}

// CloudController represents the CC domains, admin client credentials
type CloudController struct {
	RootDomain        string
	SystemDomain      string
	AppDomains        []string
	APIDomain         string
	AdminUser         string
	AdminPassword     string
	SSLSkipCertVerify bool
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

// UAA discovers the admin client credentials for the UAA
func (manifest *CFDeploymentManifest) UAA() (uaa UAA) {
	if manifest.Properties == nil {
		return
	}
	properties := *manifest.Properties
	if properties["uaa"] == nil {
		return
	}
	uaaProperties := properties["uaa"].(map[string]interface{})

	if uaaProperties["admin"] != nil {
		admin := uaaProperties["admin"].(map[string]interface{})
		if admin["client_secret"] != nil {
			uaa.AdminClientSecret = admin["client_secret"].(string)
		}
		if admin["client_id"] != nil {
			uaa.AdminClientID = admin["client_id"].(string)
		} else {
			uaa.AdminClientID = "admin"
		}
	}
	if uaaProperties["url"] != nil {
		uaa.URI = uaaProperties["url"].(string)
	}

	return
}

// CloudController discovers the client credentials & domains for the Cloud Controller
func (manifest *CFDeploymentManifest) CloudController() (cc CloudController) {
	if manifest.Properties == nil {
		return
	}
	properties := *manifest.Properties
	if properties["domain"] != nil {
		cc.RootDomain = properties["domain"].(string)
	}
	if properties["system_domain"] != nil {
		cc.SystemDomain = properties["system_domain"].(string)
	}
	if properties["api_domain"] != nil {
		cc.APIDomain = properties["api_domain"].(string)
	}
	if properties["app_domains"] != nil {
		cc.AppDomains = properties["app_domains"].([]string)
	}

	if properties["acceptance_tests"] != nil {
		acceptanceTests := properties["acceptance_tests"].(map[string]interface{})
		if acceptanceTests["admin_user"] != nil {
			cc.AdminUser = acceptanceTests["admin_user"].(string)
			cc.AdminPassword = acceptanceTests["admin_password"].(string)
		}
	}

	if properties["ssl"] != nil {
		ssl := properties["ssl"].(map[string]interface{})
		if ssl["skip_cert_verify"] != nil {
			cc.SSLSkipCertVerify = ssl["skip_cert_verify"].(bool)
		}
	}
	return
}
