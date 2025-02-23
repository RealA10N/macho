package build_version

type Version struct {
	// Source: https://github.com/apple-oss-distributions/xnu/blob/8d741a5de7ff4191bf97d57b9f54c2f6d4a15585/EXTERNAL_HEADERS/mach-o/loader.h#L1265

	Revision uint8
	Minor    uint8
	Major    uint16
}
