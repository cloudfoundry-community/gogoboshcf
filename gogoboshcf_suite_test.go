package gogoboshcf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoGoBoshForCloudFoundry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoGoBOSH for Cloud Foundry suite")
}
