package section64

import (
	"io"

	"alon.kr/x/macho/builder/context"
	"alon.kr/x/writertoutils"
)

type Section64Builder struct {
	SectionName [16]byte
	SegmentName [16]byte
	Data        []byte
	Address     uint64
	Align       uint32
	Flags       Section64Flags
	Relocations []RelocationBuilder
	// TODO: support custom section size (for BSS, etc.)
}

func (builder Section64Builder) relocationsOffset(
	ctx *context.CommandContext,
) uint64 {
	if builder.relocationsLen() == 0 {
		// Although calculating the offset even if the allocation array is of
		// size zero is technically valid, it is not useful, and can be
		// confusing, and even cause bugs in the future.
		// We mimic other compilers and tools by returning 0 in this case.
		return 0
	}

	dataStart := ctx.DataOffset
	dataEnd := dataStart + uint64(len(builder.Data))
	relocationsStart := dataEnd
	return uint64(relocationsStart)
}

func (builder Section64Builder) relocationsLen() uint64 {
	return uint64(len(builder.Relocations)) * RelocationInfoSize
}

func (builder Section64Builder) Build(
	ctx *context.CommandContext,
) Section64Header {
	return Section64Header{
		SectionName:         builder.SectionName,
		SegmentName:         builder.SegmentName,
		Address:             builder.Address,
		Size:                uint64(len(builder.Data)),
		Offset:              uint32(ctx.DataOffset),
		Align:               builder.Align,
		RelocationOffset:    uint32(builder.relocationsOffset(ctx)),
		NumberOfRelocations: uint32(len(builder.Relocations)),
		Flags:               builder.Flags,
	}
}

// CommandBuilder Implementation

func (builder Section64Builder) HeaderLen() uint64 {
	return Section64HeaderSize
}

func (builder Section64Builder) DataLen() uint64 {
	dataLen := uint64(len(builder.Data))
	relocationsLen := builder.relocationsLen()
	return dataLen + relocationsLen
}

func (builder Section64Builder) HeaderWriteTo(
	writer io.Writer,
	ctx *context.CommandContext,
) (int64, error) {
	section := builder.Build(ctx)
	writerTo := writertoutils.BinaryMarshalerAdapter(section)
	return writerTo.WriteTo(writer)
}

func (builder Section64Builder) relocationWriterTos() []io.WriterTo {
	writers := make([]io.WriterTo, len(builder.Relocations))
	for i, relocation := range builder.Relocations {
		info := relocation.Build()
		writers[i] = writertoutils.BinaryMarshalerAdapter(info)
	}
	return writers
}

func (builder Section64Builder) DataWriteTo(writer io.Writer) (int64, error) {
	writerTos := append(
		[]io.WriterTo{writertoutils.BufferWriterTo(builder.Data)},
		builder.relocationWriterTos()...,
	)
	return writertoutils.MultiWriterTo(writerTos...).WriteTo(writer)
}
