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
})
