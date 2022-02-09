package app

import (
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func testSkipNonWasm(t *testing.T) {
	if goarch := runtime.GOARCH; goarch != "wasm" {
		t.Skip("not wasm, but is", goarch)
	}
}

func testSkipWasm(t *testing.T) {
	if goarch := runtime.GOARCH; goarch == "wasm" {
		t.Skip("not wasm, but is", goarch)
	}
}

func testCreateDir(t *testing.T, path string) func() {
	err := os.MkdirAll(path, 0755)
	require.NoError(t, err)

	return func() {
		os.RemoveAll(path)
	}
}

func testCreateFile(t *testing.T, path, content string) {
	err := ioutil.WriteFile(path, stob(content), 0666)
	require.NoError(t, err)
}
