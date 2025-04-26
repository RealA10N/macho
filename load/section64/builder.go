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
	Relocations RelocationsBuilder
	// TODO: support custom section size (for BSS, etc.)
}

func (builder Section64Builder) Build(
	ctx *context.CommandContext,
) Section64Header {
	numberOfRelocations := uint32(builder.Relocations.NumberOfRelocations())
	relocationsOffset := uint32(0)
	if numberOfRelocations != 0 {
		relocationsOffset = uint32(ctx.DataOffset + uint64(len(builder.Data)))
	}

	return Section64Header{
		SectionName:         builder.SectionName,
		SegmentName:         builder.SegmentName,
		Address:             builder.Address,
		Size:                uint64(len(builder.Data)),
		Offset:              uint32(ctx.DataOffset),
		Align:               builder.Align,
		RelocationOffset:    relocationsOffset,
		NumberOfRelocations: numberOfRelocations,
		Flags:               builder.Flags,
	}
}

// CommandBuilder Implementation

func (builder Section64Builder) HeaderLen() uint64 {
	return Section64HeaderSize
}

func (builder Section64Builder) DataLen() uint64 {
	dataLen := uint64(len(builder.Data))
	relocationsLen := builder.Relocations.Len()
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

func (builder Section64Builder) DataWriteTo(writer io.Writer) (int64, error) {
	return writertoutils.MultiWriterTo(
		writertoutils.BufferWriterTo(builder.Data),
		builder.Relocations,
	).WriteTo(writer)
}
