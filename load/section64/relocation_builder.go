package section64

type RelocationBuilder struct {
	Address                uint32
	SymbolIndex            uint32
	IsRelocationPcRelative bool
	Length                 RelocationLength
	IsRelocationExtern     bool
	Type                   RelocationType
}

func (builder RelocationBuilder) Build() RelocationInfo {
	if builder.SymbolIndex > RelocationSymbolNumMask {
		return RelocationInfo{} // TODO: handle error
	}

	details := (builder.SymbolIndex |
		(uint32(builder.Length) << 25) |
		(uint32(builder.Type) << 28))

	if builder.IsRelocationPcRelative {
		details |= RelocationPcRelative
	}

	if builder.IsRelocationExtern {
		details |= RelocationExtern
	}

	return RelocationInfo{
		Address: builder.Address,
		Details: details,
	}
}
