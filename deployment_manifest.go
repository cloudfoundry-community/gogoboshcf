package gogoboshcf

import (
	"strings"

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

// SSL describes if components should skip SSL certificate verification (due to self-signed certs)
type SSL struct {
	SkipCertificateVerify bool `yaml:"skip_cert_verify"`
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
	URI       string `yaml:"url"`
	Admin     ClientIDSecret
	Scim      Scim
	ScimUsers []ScimUser
}

// ClientIDSecret describes an ID/Secret pair
type ClientIDSecret struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// Scim lists the built-in UAA users
type Scim struct {
	Users []string
}

// ScimUser summarizes nicely the built-in UAA users
type ScimUser struct {
	Username string
	Password string
	Scopes   []string
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
	pm.setupScimUsers()
	return
}

func (pm *PropertiesManifest) setupScimUsers() {
	for _, userString := range pm.UAA.Scim.Users {
		tokens := strings.Split(userString, "|")
		scimUser := ScimUser{
			Username: tokens[0],
			Password: tokens[1],
			Scopes:   strings.Split(tokens[2], ","),
		}
		pm.UAA.ScimUsers = append(pm.UAA.ScimUsers, scimUser)
	}
}
