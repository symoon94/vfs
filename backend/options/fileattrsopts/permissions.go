package fileattrsopts

import (
	"github.com/c2fo/vfs/v6"
)

const OptionNamePermissions = "permissions"

func WithPermissions(perms uint32) vfs.FileAttrsOption {
	if perms == 0 {
		perms = 0644
	}
	return Permissions{
		permissions: perms,
	}
}

type Permissions struct {
	permissions uint32
}

func (p Permissions) FileAttrsOptionName() string {
	return OptionNamePermissions
}

func (p Permissions) GetPermissions() uint32 {
	return p.permissions
}
