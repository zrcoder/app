//go:build wasm
// +build wasm

package app

import (
	"runtime"

	"github.com/zrcoder/app/errors"
)

const (
	wasmExecJS   = ""
	appJS        = ""
	appWorkerJS  = ""
	manifestJSON = ""
	appCSS       = ""
)

var (
	errBadInstruction = errors.New("unsupported instruction").
		Tag("architecture", runtime.GOARCH)
)

func GenerateStaticWebsite(dir string, h *Handler, pages ...string) error {
	panic(errBadInstruction)
}
