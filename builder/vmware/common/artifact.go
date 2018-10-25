package common

import (
	"fmt"
	"os"
	"path/filepath"

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
	dir       OutputDir
	f         []string
	config    map[string]string
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
	return a.dir.RemoveAll()
}

// NewLocalArtifact wraps NewArtifact and finds the files in the given directory.
func NewLocalArtifact(vmname string, dir OutputDir, files []string, config map[string]string) (packer.Artifact, error) {
	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	}
	if err := filepath.Walk(dir.String(), visit); err != nil {
		return nil, err
	}
	return NewArtifact(vmname, dir, files, config, false)
}

func NewArtifact(vmname string, dir OutputDir, files []string, config map[string]string, esxi bool) (packer.Artifact, error) {
	builderID := BuilderId
	if esxi {
		builderID = BuilderIdESX
	}

	return &artifact{
		builderId: builderID,
		id:        vmname,
		dir:       dir,
		f:         files,
	}, nil
}
