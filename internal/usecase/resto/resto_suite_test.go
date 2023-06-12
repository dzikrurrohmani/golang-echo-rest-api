package resto_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestResto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resto Suite")
}
