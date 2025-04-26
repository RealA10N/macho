package section64

import (
	"io"

	"alon.kr/x/writertoutils"
)

type RelocationsBuilder struct {
	Relocations []RelocationInfo
}

func (builder RelocationsBuilder) Build() RelocationInfo {
	return RelocationInfo{}
}

// Returns the number of relocations in the relocations array.
func (builder RelocationsBuilder) NumberOfRelocations() int {
	return len(builder.Relocations)
}

func (builder RelocationsBuilder) IsEmpty() bool {
	return builder.NumberOfRelocations() == 0
}

// Returns the data length of the relocations array, in bytes.
func (builder RelocationsBuilder) Len() uint64 {
	return uint64(builder.NumberOfRelocations() * RelocationInfoSize)
}

func (builder RelocationsBuilder) WriteTo(writer io.Writer) (int64, error) {
	writerTos := make([]io.WriterTo, len(builder.Relocations))
	for i, relocation := range builder.Relocations {
		writerTos[i] = writertoutils.BinaryMarshalerAdapter(relocation)
	}

	return writertoutils.MultiWriterTo(writerTos...).WriteTo(writer)
}
