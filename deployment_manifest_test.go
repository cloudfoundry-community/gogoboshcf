package gogoboshcf_test

import (
	"github.com/cloudfoundry-community/gogoboshcf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFDeploymentManifest", func() {
	It("NATS in global properties", func() {
		manifest := &gogoboshcf.CFDeploymentManifest{
			Properties: &map[string]interface{}{
				"nats": map[string]interface{}{
					"machines": []string{"server1", "server2"},
					"port":     4444,
					"username": "nats",
					"password": "password",
				},
				"another": "FIXME",
			},
		}
		nats := manifest.NATS()
		Expect(nats.MachinesHostnames).To(Equal([]string{"server1", "server2"}))
		Expect(nats.Port).To(Equal(4444))
		Expect(nats.Username).To(Equal("nats"))
		Expect(nats.Password).To(Equal("password"))
	})

	It("UAA in global properties", func() {
		manifest := &gogoboshcf.CFDeploymentManifest{
			Properties: &map[string]interface{}{
				"uaa": map[string]interface{}{
					"admin": map[string]interface{}{
						"client_secret": "admin-secret",
					},
					"url": "https://uaa.10.244.0.34.xip.io",
				},
			},
		}
		uaa := manifest.UAA()
		Expect(uaa.AdminClientSecret).To(Equal("admin-secret"))
		Expect(uaa.AdminClientID).To(Equal("admin"))
		Expect(uaa.URI).To(Equal("https://uaa.10.244.0.34.xip.io"))
	})

	It("CloudController in global properties", func() {
		manifest := &gogoboshcf.CFDeploymentManifest{
			Properties: &map[string]interface{}{
				"domain":        "10.244.0.34.xip.io",
				"system_domain": "10.244.0.34.xip.io",
				"api_domain":    "api.10.244.0.34.xip.io",
				"app_domains":   []string{"10.244.0.34.xip.io"},
				"acceptance_tests": map[string]interface{}{
					"admin_user":     "user",
					"admin_password": "password",
				},
				"ssl": map[string]interface{}{"skip_cert_verify": true},
			},
		}
		cc := manifest.CloudController()
		Expect(cc.RootDomain).To(Equal("10.244.0.34.xip.io"))
		Expect(cc.SystemDomain).To(Equal("10.244.0.34.xip.io"))
		Expect(cc.AppDomains).To(Equal([]string{"10.244.0.34.xip.io"}))
		Expect(cc.APIDomain).To(Equal("api.10.244.0.34.xip.io"))
		Expect(cc.AdminUser).To(Equal("user"))
		Expect(cc.AdminPassword).To(Equal("password"))

		Expect(cc.SSLSkipCertVerify).To(Equal(true))
	})
})
