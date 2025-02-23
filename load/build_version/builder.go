package build_version

import (
	"io"

	"alon.kr/x/macho/builder/context"
	"alon.kr/x/macho/load"
	"alon.kr/x/writertoutils"
)

type BuildVersionBuilder struct {
	Platform Platform
	MinOs    Version
	Sdk      Version
	Tools    []BuildTool
}

func (builder BuildVersionBuilder) HeaderLen() uint64 {
	tools := uint64(len(builder.Tools))
	toolsSize := tools * BuildToolSize
	return BuildVersionHeaderSize + toolsSize
}

func (BuildVersionBuilder) DataLen() uint64 {
	return 0
}

func (builder BuildVersionBuilder) HeaderWriteTo(
	writer io.Writer,
	ctx *context.CommandContext,
) (int64, error) {
	header := BuildVersionHeader{
		CommandType: load.BuildVersion,
		CommandSize: uint32(builder.HeaderLen()),
		Platform:    builder.Platform,
		MinOs:       builder.MinOs,
		Sdk:         builder.Sdk,
		NumOfTools:  uint32(len(builder.Tools)),
	}

	headerWriterTo := writertoutils.BinaryMarshalerAdapter(header)
	writerTos := []io.WriterTo{headerWriterTo}

	for _, tool := range builder.Tools {
		toolWriterTo := writertoutils.BinaryMarshalerAdapter(tool)
		writerTos = append(writerTos, toolWriterTo)
	}

	return writertoutils.MultiWriterTo(writerTos...).WriteTo(writer)
}
