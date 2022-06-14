package os

import (
	"path"

	"github.com/c2fo/vfs/v6"
	"github.com/c2fo/vfs/v6/backend"
	"github.com/c2fo/vfs/v6/backend/options/fileattrsopts"
	"github.com/c2fo/vfs/v6/utils"
)

// Scheme defines the file system type.
const Scheme = "file"
const name = "os"

// FileSystem implements vfs.Filesystem for the OS file system.
type FileSystem struct{}

// Retry will return a retriever provided via options, or a no-op if none is provided.
func (fs *FileSystem) Retry() vfs.Retry {
	return vfs.DefaultRetryer()
}

// NewFile function returns the os implementation of vfs.File.
func (fs *FileSystem) NewFile(volume, name string, opts ...vfs.FileAttrsOption) (vfs.File, error) {
	err := utils.ValidateAbsoluteFilePath(name)
	if err != nil {
		return nil, err
	}

	file := &File{name: name, filesystem: fs}
	for _, o := range opts {
		switch o.FileAttrsOptionName() {
		case fileattrsopts.OptionNamePermissions:
			file.permissions = o.(fileattrsopts.Permissions).GetPermissions()
		}
	}
	return file, nil
}

// NewLocation function returns the os implementation of vfs.Location.
func (fs *FileSystem) NewLocation(volume, name string) (vfs.Location, error) {
	err := utils.ValidateAbsoluteLocationPath(name)
	if err != nil {
		return nil, err
	}

	return &Location{
		fileSystem: fs,
		name:       utils.EnsureTrailingSlash(path.Clean(name)),
	}, nil
}

// Name returns "os"
func (fs *FileSystem) Name() string {
	return name
}

// Scheme return "file" as the initial part of a file URI ie: file://
func (fs *FileSystem) Scheme() string {
	return Scheme
}

func init() {
	backend.Register(Scheme, &FileSystem{})
}
