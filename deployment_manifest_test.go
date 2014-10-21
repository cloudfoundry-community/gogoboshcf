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
					"user":     "nats",
					"password": "password",
				},
				"another": "FIXME",
			},
		}
		gp, err := manifest.GlobalProperties()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(gp.NATS.MachinesHostnames).To(Equal([]string{"server1", "server2"}))
		Expect(gp.NATS.Port).To(Equal(4444))
		Expect(gp.NATS.Username).To(Equal("nats"))
		Expect(gp.NATS.Password).To(Equal("password"))
	})

	It("UAA in global properties", func() {
		manifest := &gogoboshcf.CFDeploymentManifest{
			Properties: &map[string]interface{}{
				"uaa": map[string]interface{}{
					"admin": map[string]interface{}{
						"client_secret": "admin-secret",
					},
					"url": "https://uaa.10.244.0.34.xip.io",
					"scim": map[string]interface{}{
						"users": []string{"admin|password|scim.write,scim.read,openid,cloud_controller.admin,clients.read,clients.write,doppler.firehose"},
					},
				},
			},
		}
		gp, err := manifest.GlobalProperties()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(gp.UAA.Admin.ClientSecret).To(Equal("admin-secret"))
		Expect(gp.UAA.URI).To(Equal("https://uaa.10.244.0.34.xip.io"))
		Expect(gp.UAA.Scim.Users).To(Equal([]string{"admin|password|scim.write,scim.read,openid,cloud_controller.admin,clients.read,clients.write,doppler.firehose"}))
		Expect(len(gp.UAA.ScimUsers)).To(Equal(1))
		Expect(gp.UAA.ScimUsers[0].Username).To(Equal("admin"))
		Expect(gp.UAA.ScimUsers[0].Password).To(Equal("password"))
		Expect(gp.UAA.ScimUsers[0].Scopes).To(Equal([]string{"scim.write", "scim.read", "openid", "cloud_controller.admin", "clients.read", "clients.write", "doppler.firehose"}))
	})

	It("CloudController in global properties", func() {
		manifest := &gogoboshcf.CFDeploymentManifest{
			Properties: &map[string]interface{}{
				"domain":        "10.244.0.34.xip.io",
				"system_domain": "10.244.0.34.xip.io",
				"app_domains":   []string{"10.244.0.34.xip.io"},
				"acceptance_tests": map[string]interface{}{
					"admin_user":     "user",
					"admin_password": "password",
				},
				"ssl": map[string]interface{}{"skip_cert_verify": true},
			},
		}
		gp, err := manifest.GlobalProperties()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(gp.RootDomain).To(Equal("10.244.0.34.xip.io"))
		Expect(gp.SystemDomain).To(Equal("10.244.0.34.xip.io"))
		Expect(gp.AppDomains).To(Equal([]string{"10.244.0.34.xip.io"}))

		Expect(gp.SSL.SkipCertificateVerify).To(Equal(true))
	})
})
