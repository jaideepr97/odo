package project

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/openshift/odo/pkg/util"
	"github.com/openshift/odo/tests/helper"
)

func TestProject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Project Suite")
	// Keep CustomReporters commented till https://github.com/onsi/ginkgo/issues/628 is fixed
	// RunSpecsWithDefaultAndCustomReporters(t, "Project Suite", []Reporter{reporter.JunitReport(t, "../../../reports")})
}

var tempdir string

var _ = BeforeSuite(func() {
	originalKubeconfig := os.Getenv("KUBECONFIG")
	if len(originalKubeconfig) > 0 {
		tempdir = helper.CreateNewContext()
		info, err := os.Stat(originalKubeconfig)
		Expect(err).NotTo(HaveOccurred())

		err = util.CopyFile(originalKubeconfig, filepath.Join(tempdir, "kubeconfig"), info)
		Expect(err).NotTo(HaveOccurred())
	}
})

var _ = AfterSuite(func() {
	helper.DeleteDir(tempdir)
})
