package build_version

import (
	"alon.kr/x/macho/load"
	"alon.kr/x/macho/utils"
)

type BuildVersionHeader struct {
	// Source: https://github.com/apple-oss-distributions/xnu/blob/8d741a5de7ff4191bf97d57b9f54c2f6d4a15585/EXTERNAL_HEADERS/mach-o/loader.h#L1260

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
