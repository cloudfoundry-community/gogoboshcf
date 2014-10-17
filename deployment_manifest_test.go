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
})
