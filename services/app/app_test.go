package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	a := New()
	a.Bootstrap("../../config.yaml")
	err := a.DownloadWF()
	assert.Nil(t, err)
}

func TestOutput(t *testing.T) {
	a := New()
	a.Bootstrap("../../config.yaml")
	a.OutputWFs()
}
