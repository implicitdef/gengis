package specrunner_test

import (
	. "github.com/mtailor/gengis/vendor/_nuts/github.com/onsi/ginkgo"
	. "github.com/mtailor/gengis/vendor/_nuts/github.com/onsi/gomega"

	"testing"
)

func TestSpecRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Runner Suite")
}
