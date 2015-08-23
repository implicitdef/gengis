package containernode_test

import (
	. "github.com/mtailor/gengis/vendor/_nuts/github.com/onsi/ginkgo"
	. "github.com/mtailor/gengis/vendor/_nuts/github.com/onsi/gomega"

	"testing"
)

func TestContainernode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containernode Suite")
}
