package en_GB_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLocale(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "en_GB Locale Suite")
}
