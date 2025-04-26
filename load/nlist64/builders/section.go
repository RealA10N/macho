package nlist64_builders

import (
	"alon.kr/x/macho/load/nlist64"
	"alon.kr/x/macho/load/symtab/symbol"
)

type SectionNlist64Builder struct {
	Name        string
	Type        nlist64.SymbolType
	Section     uint8
	Offset      uint64
	Description nlist64.SymbolDescription
}

func (symbol SectionNlist64Builder) GenString() string {
	return symbol.Name
}

func (builder SectionNlist64Builder) GenEntryList(
	ctx *symbol.EntryListContext,
) (nlist64.Nlist64, error) {
	return nlist64.Nlist64{
		StringTableOffset: ctx.StringTableOffset,
		SymbolType:        builder.Type,
		Section:           builder.Section,
		Description:       builder.Description,
		Value:             builder.Offset,
	}, nil
}
