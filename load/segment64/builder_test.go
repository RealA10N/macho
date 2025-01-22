package segment64_test

import (
	"bytes"
	"testing"

	"alon.kr/x/macho/builder/context"
	"alon.kr/x/macho/load/section64"
	"alon.kr/x/macho/load/segment64"
	"github.com/stretchr/testify/assert"
)

func TestSegment64BuilderExpectedBinary(t *testing.T) {
	expectedHeader := []byte{
		0x19, 0x00, 0x00, 0x00, 0x98, 0x00, 0x00, 0x00,
		0x5F, 0x5F, 0x54, 0x45, 0x58, 0x54, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x39, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x5F, 0x5F, 0x74, 0x65, 0x78, 0x74, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x5F, 0x5F, 0x54, 0x45, 0x58, 0x54, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x39, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	data := []byte{0x00, 0x00, 0x01, 0x8B, 0xC0, 0x03, 0x5F, 0xD6}

	sectionBuilder := section64.Section64Builder{
		SectionName: [16]byte{'_', '_', 't', 'e', 'x', 't', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		SegmentName: [16]byte{'_', '_', 'T', 'E', 'X', 'T', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Data:        data,
		Flags:       section64.AttrPureInstructions | section64.AttrSomeInstructions,
	}

	segmentBuilder := segment64.Segment64Builder{
		SegmentName:        [16]byte{'_', '_', 'T', 'E', 'X', 'T', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Sections:           []section64.Section64Builder{sectionBuilder},
		VirtualMemorySize:  8,
		MaxProtections:     segment64.AllowAllProtection,
		InitialProtections: segment64.AllowAllProtection,
	}

	// HeaderLen
	assert.EqualValues(t, len(expectedHeader), segmentBuilder.HeaderLen())

	// DataLen
	assert.EqualValues(t, len(data), segmentBuilder.DataLen())

	{
		// HeaderWriteTo
		buffer := bytes.Buffer{}
		ctx := context.CommandContext{DataOffset: 1337}
		n, err := segmentBuilder.HeaderWriteTo(&buffer, &ctx)
		assert.NoError(t, err)
		assert.EqualValues(t, len(expectedHeader), n)
		assert.Equal(t, expectedHeader, buffer.Bytes())
	}

	{
		// DataWriteTo
		buffer := bytes.Buffer{}
		n, err := segmentBuilder.DataWriteTo(&buffer)
		assert.NoError(t, err)
		assert.EqualValues(t, len(data), n)
	}
}
