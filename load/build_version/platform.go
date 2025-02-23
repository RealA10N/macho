package build_version

type Platform uint32

const (
	// Source: https://github.com/apple-oss-distributions/xnu/blob/8d741a5de7ff4191bf97d57b9f54c2f6d4a15585/EXTERNAL_HEADERS/mach-o/loader.h#L1275C1-L1275C49

	PlatformMacOS            Platform = 1
	PlatformIOS              Platform = 2
	PlatformTvOS             Platform = 3
	PlatformWatchOS          Platform = 4
	PlatformBridgeOS         Platform = 5
	PlatformMacCatalyst      Platform = 6
	PlatformIOSSimulator     Platform = 7
	PlatformTvOSSimulator    Platform = 8
	PlatformWatchOSSimulator Platform = 9
	PlatformDriverKit        Platform = 10
)
