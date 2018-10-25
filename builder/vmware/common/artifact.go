package common

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer/packer"
)

const (
	// BuilderId for the local artifacts
	BuilderId    = "mitchellh.vmware"
	BuilderIdESX = "mitchellh.vmware-esx"

	ArtifactConfFormat         = "artifact.conf.format"
	ArtifactConfKeepRegistered = "artifact.conf.keep_registered"
	ArtifactConfSkipExport     = "artifact.conf.skip_export"
)

// Artifact is the result of running the VMware builder, namely a set
// of files associated with the resulting machine.
type artifact struct {
	builderId string
	id        string
	dir       string
	f         []string
	config    map[string]string
}

func NewArtifact(vmname string, dir OutputDir, files []string, config map[string]string, esxi bool) (packer.Artifact, error) {
	builderID := BuilderId
	if esxi {
		builderID = BuilderIdESX
	}

	return &artifact{
		builderId: builderID,
		id:        vmname,
		dir:       dir.String(),
		f:         files,
	}, nil
}

func (a *artifact) BuilderId() string {
	return a.builderId
}

func (a *artifact) Files() []string {
	return a.f
}

func (a *artifact) Id() string {
	return a.id
}

func (a *artifact) String() string {
	return fmt.Sprintf("VM files in directory: %s", a.dir)
}

func (a *artifact) State(name string) interface{} {
	return a.config[name]
}

func (a *artifact) Destroy() error {
	return os.RemoveAll(a.dir)
}
