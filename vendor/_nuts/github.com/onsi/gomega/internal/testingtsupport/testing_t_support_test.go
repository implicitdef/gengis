package testingtsupport_test

import (
	. "github.com/mtailor/gengis/vendor/_nuts/github.com/onsi/gomega"

	"testing"
)

func TestTestingT(t *testing.T) {
	RegisterTestingT(t)
	Ω(true).Should(BeTrue())
}
