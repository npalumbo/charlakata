package charlakata_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCharlakata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Charlakata Suite")
}
