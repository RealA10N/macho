package build_version

import (
	"alon.kr/x/macho/load"
	"alon.kr/x/macho/utils"
)

type BuildVersionHeader struct {
	CommandType load.CommandType
	CommandSize uint32
	Platform    Platform
	MinOs       Version
	Sdk         Version
	NumOfTools  uint32
}

const BuildVersionHeaderSize uint64 = 0x18

func (bv BuildVersionHeader) MarshalBinary() ([]byte, error) {
	return utils.GenericMarshalBinary(bv)
}
