package gogoboshcf

import (
	"github.com/cloudfoundry-community/gogobosh/models"
	"launchpad.net/goyaml"
)

// CFDeploymentManifest is a deployment manifest for a Cloud Foundry deployment
type CFDeploymentManifest models.DeploymentManifest

// PropertiesManifest represents the "properties" section of CFDeploymentManifest
type PropertiesManifest struct {
	NATS             NATS
	UAA              UAA
	RootDomain       string   `yaml:"domain"`
	SystemDomain     string   `yaml:"system_domain"`
	AppDomains       []string `yaml:"app_domains"`
	SSL              SSL
	SyslogAggregator map[string]interface{} `yaml:"syslog_aggregator"`
}

// NATS represents the NATS client credentials
type NATS struct {
	MachinesHostnames []string `yaml:"machines"`
	Port              int
	Username          string `yaml:"user"`
	Password          string
}

// UAA represents UAA admin client credentials
type UAA struct {
	URI   string `yaml:"url"`
	Admin ClientIDSecret
}

// ClientIDSecret describes an ID/Secret pair
type ClientIDSecret struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// SSL describes if components should skip SSL certificate verification (due to self-signed certs)
type SSL struct {
	SkipCertificateVerify bool `yaml:"skip_cert_verify"`
}

// GlobalProperties returns the properties
// TODO: marshal Properties to YAML; then unmarshal to a CF struct
func (manifest *CFDeploymentManifest) GlobalProperties() (pm *PropertiesManifest, err error) {
	yaml, err := goyaml.Marshal(manifest.Properties)
	if err != nil {
		return
	}
	pm = &PropertiesManifest{}
	goyaml.Unmarshal([]byte(yaml), pm)
	return
}
