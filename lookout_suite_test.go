package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLookout(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lookout Suite")
}
