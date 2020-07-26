package parser

import (
	devfileCtx "github.com/jaideepr97/parser/pkg/devfile/context"
	"github.com/jaideepr97/parser/pkg/devfile/data"
)

// Default filenames for create devfile
const (
	OutputDevfileJsonPath = "odo-devfile.json"
	OutputDevfileYamlPath = "odo-devfile.yaml"
)

// DevfileObj is the runtime devfile object
type DevfileObj struct {

	// Ctx has devfile context info
	Ctx devfileCtx.DevfileCtx

	// Data has the devfile data
	Data data.DevfileData
}
