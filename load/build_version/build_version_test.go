package build_version_test

import (
	"bytes"
	"testing"

	"alon.kr/x/macho/builder/context"
	"alon.kr/x/macho/load/build_version"
	"github.com/stretchr/testify/assert"
)

func TestBuildVersionHeaderExpectedBinary(t *testing.T) {
	expected := []byte{
		0x32, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x0a, 0x00,
		0x00, 0x01, 0x0c, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x02,
	}

	builder := build_version.BuildVersionBuilder{
		Platform: build_version.PlatformMacOS,
		MinOs:    build_version.Version{Major: 10, Minor: 15},
		Sdk:      build_version.Version{Major: 12, Minor: 1},
		Tools: []build_version.BuildTool{
			{
				ToolType: build_version.ToolTypeLd,
				Version:  0x02c70000,
			},
		},
	}

	// We pass a nil context here since the LC_BUILD_VERSION command does not
	// actually require a context.
	var ctx *context.CommandContext
	buffer := new(bytes.Buffer)

	n, err := builder.HeaderWriteTo(buffer, ctx)
	assert.NoError(t, err)
	assert.Equal(t, int64(len(expected)), n)
	assert.Equal(t, expected, buffer.Bytes())
}
