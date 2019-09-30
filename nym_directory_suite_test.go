package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNymDirectory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NymDirectory Suite")
}
