package common

import (
	"github.com/hashicorp/packer/packer"
	"testing"
)

func TestLocalArtifact_impl(t *testing.T) {
	var _ packer.Artifact = new(artifact)
}
