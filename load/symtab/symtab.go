package symtab

import (
	"alon.kr/x/macho/load"
	"alon.kr/x/macho/utils"
)

type SymtabHeader struct {
	CommandType       load.CommandType
	CommandSize       uint32
	SymbolTableOffset uint32
	NumOfSymbols      uint32
	StringTableOffset uint32
	StringTableSize   uint32
}

const SymTabHeaderSize uint64 = 0x18

func (symtab SymtabHeader) MarshalBinary() ([]byte, error) {
	return utils.GenericMarshalBinary(symtab)
}
