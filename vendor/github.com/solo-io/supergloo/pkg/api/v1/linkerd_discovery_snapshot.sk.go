// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type LinkerdDiscoverySnapshot struct {
	Meshes   MeshesByNamespace
	Installs InstallsByNamespace
}

func (s LinkerdDiscoverySnapshot) Clone() LinkerdDiscoverySnapshot {
	return LinkerdDiscoverySnapshot{
		Meshes:   s.Meshes.Clone(),
		Installs: s.Installs.Clone(),
	}
}

func (s LinkerdDiscoverySnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashMeshes(),
		s.hashInstalls(),
	)
}

func (s LinkerdDiscoverySnapshot) hashMeshes() uint64 {
	return hashutils.HashAll(s.Meshes.List().AsInterfaces()...)
}

func (s LinkerdDiscoverySnapshot) hashInstalls() uint64 {
	return hashutils.HashAll(s.Installs.List().AsInterfaces()...)
}

func (s LinkerdDiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("meshes", s.hashMeshes()))
	fields = append(fields, zap.Uint64("installs", s.hashInstalls()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type LinkerdDiscoverySnapshotStringer struct {
	Version  uint64
	Meshes   []string
	Installs []string
}

func (ss LinkerdDiscoverySnapshotStringer) String() string {
	s := fmt.Sprintf("LinkerdDiscoverySnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Meshes %v\n", len(ss.Meshes))
	for _, name := range ss.Meshes {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Installs %v\n", len(ss.Installs))
	for _, name := range ss.Installs {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s LinkerdDiscoverySnapshot) Stringer() LinkerdDiscoverySnapshotStringer {
	return LinkerdDiscoverySnapshotStringer{
		Version:  s.Hash(),
		Meshes:   s.Meshes.List().NamespacesDotNames(),
		Installs: s.Installs.List().NamespacesDotNames(),
	}
}