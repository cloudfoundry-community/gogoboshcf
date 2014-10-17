package gogoboshcf_test

import (
	"github.com/cloudfoundry-community/gogoboshcf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFDeploymentManifest", func() {
	It("FindNATSMachines in global properties", func() {
		manifest := &gogoboshcf.CFDeploymentManifest{
			Properties: &map[string]interface{}{
				"nats": map[string]interface{}{
					"machines": []string{"server1", "server2"},
				},
				"another": "FIXME",
			},
		}
		machines := manifest.FindNATSMachines()
		Expect(machines).To(Equal([]string{"server1", "server2"}))
	})
})
