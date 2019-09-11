package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDirectoryServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DirectoryServer Suite")
}
