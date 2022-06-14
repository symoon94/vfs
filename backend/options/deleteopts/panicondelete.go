package deleteopts

import "github.com/c2fo/vfs/v6"

const OptionNamePanicOnDelete = "panicOnDelete"

func WithPanicOnDelete() vfs.DeleteOption {
	return PanicOnDelete{}
}

type PanicOnDelete struct{}

func (p PanicOnDelete) DeleteOptionName() string {
	return OptionNamePanicOnDelete
}
