package data

import (
	"github.com/jaideepr97/parser/pkg/devfile/data/common"
)

// DevfileData is an interface that defines functions for Devfile data operations
type DevfileData interface {
	GetMetadata() common.DevfileMetadata
	GetParent() common.DevfileParent
	GetEvents() common.DevfileEvents
	GetComponents() []common.DevfileComponent
	GetAliasedComponents() []common.DevfileComponent
	GetProjects() []common.DevfileProject
	GetCommands() []common.DevfileCommand
}
