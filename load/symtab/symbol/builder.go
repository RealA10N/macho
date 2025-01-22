package symbol

import "alon.kr/x/macho/load/nlist64"

type EntryListContext struct {
	StringTableOffset uint32
}

func (ctx *EntryListContext) UpdateAfterProcessing(sym SymbolBuilder) {
	ctx.StringTableOffset += uint32(len(sym.GenString())) + 1
}

type SymbolBuilder interface {
	GenString() string
	GenEntryList(*EntryListContext) (nlist64.Nlist64, error)
}
