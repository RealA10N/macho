package build_version

type Version uint32

// Source: https://github.com/apple-oss-distributions/xnu/blob/8d741a5de7ff4191bf97d57b9f54c2f6d4a15585/EXTERNAL_HEADERS/mach-o/loader.h#L1265

func (v Version) Major() uint16 {
	return uint16(v >> 16)
}

func (v Version) Minor() uint8 {
	return uint8((v >> 8) & 0xFF)
}

func (v Version) Revision() uint8 {
	return uint8(v & 0xFF)
}
