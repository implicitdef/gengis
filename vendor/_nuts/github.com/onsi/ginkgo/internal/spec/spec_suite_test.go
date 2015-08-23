package spec_test

import (
	. "github.com/mtailor/gengis/vendor/_nuts/github.com/onsi/ginkgo"
	. "github.com/mtailor/gengis/vendor/_nuts/github.com/onsi/gomega"

	"testing"
)

func TestSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Suite")
}
