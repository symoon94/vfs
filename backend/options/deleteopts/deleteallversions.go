package deleteopts

import "github.com/c2fo/vfs/v6"

const OptionNameDeleteAllVersions = "deleteAllVersions"

func WithDeleteAllVersion() vfs.DeleteOption {
	return DeleteAllVersions{}
}

type DeleteAllVersions struct{}

func (w DeleteAllVersions) DeleteOptionName() string {
	return OptionNameDeleteAllVersions
}
