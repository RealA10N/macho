package build_version

import "alon.kr/x/macho/utils"

type ToolType uint32

const (
	// Source: https://github.com/apple-oss-distributions/xnu/blob/8d741a5de7ff4191bf97d57b9f54c2f6d4a15585/EXTERNAL_HEADERS/mach-o/loader.h#L1290

	ToolTypeClang ToolType = 1
	ToolTypeSwift ToolType = 2
	ToolTypeLd    ToolType = 3
)

type BuildTool struct {
	ToolType ToolType

	// Version of the tool.
	// Seems like each tool defines its own versioning scheme, and there is no
	// unified way to format those 32 bits.
	Version uint32
}

const BuildToolSize uint64 = 0x8

func (tool BuildTool) MarshalBinary() ([]byte, error) {
	return utils.GenericMarshalBinary(tool)
}
