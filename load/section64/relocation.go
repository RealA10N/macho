package section64

import (
	"fmt"

	"alon.kr/x/macho/utils"
)

type RelocationLength uint32

const (
	RelocationLengthByte RelocationLength = iota // 1 byte
	RelocationLengthWord                         // 2 bytes
	RelocationLengthLong                         // 4 bytes
	RelocationLengthQuad                         // 8 bytes
)

// The r_type:4 field contains the relocation type, which dictates where
// exactly and in what format the relocation is applied.
//
// The zero value is a generic "vanilla" relocation which just writes the
// relocation value of the provided size (according to the r_length field).
// Other (non-zero) values are specific to the architecture.
//
// The following resource contains a list of relocation types for x86,
// x86_64, and powerpc architectures:
// https://github.com/aidansteele/osx-abi-macho-file-format-reference
//
// More relocation types can be found here:
// https://github.com/gimli-rs/object/blob/master/src/macho.rs
type RelocationType uint32

const (
	RelocationSymbolNumMask uint32 = 0xFFFFFF
	RelocationPcRelative    uint32 = 1 << 24
	RelocationExtern        uint32 = 1 << 27
)

const (
	// for pointers
	RelocationTypeArm64Unsigned RelocationType = iota

	// must be followed by a ARM64_RELOC_UNSIGNED
	RelocationTypeArm64Subtractor

	// a B/BL instruction with 26-bit displacement
	RelocationTypeArm64Branch26

	// pc-rel distance to page of target
	RelocationTypeArm64Page21

	// offset within page, scaled by r_length
	RelocationTypeArm64PageOff12

	// pc-rel distance to page of GOT slot
	RelocationTypeArm64GotLoadPage21

	// offset within page of GOT slot, scaled by r_length
	RelocationTypeArm64GotLoadPageOff12

	// for pointers to GOT slots
	RelocationTypeArm64PointerToGot

	// pc-rel distance to page of TLVP slot
	RelocationTypeArm64TlvpLoadPage21

	// offset within page of TLVP slot, scaled by r_length
	RelocationTypeArm64TlvpLoadPageOff12

	// must be followed by PAGE21 or PAGEOFF12
	RelocationTypeArm64Addend
)

// Source: https://alexdremov.me/mystery-of-mach-o-object-file-builders/
//
//	struct relocation_info {
//		int32_t  r_address;			/* offset in the section to */
//									/* what is being relocated */
//		uint32_t r_symbolnum:24,	/* symbol index if r_extern == 1 or */
//									/* section ordinal if r_extern == 0 */
//		r_pcrel:1,					/* was relocated pc relative already */
//		r_length:2,					/* 0=byte, 1=word, 2=long, 3=quad */
//		r_extern:1,					/* does not include value of sym referenced */
//		r_type:4;					/* if not 0, machine specific relocation type */
//	};
type RelocationInfo struct {
	Address uint32
	Details uint32
}

const RelocationInfoSize int = 8

func NewRelocationInfo(
	address uint32,
	symbolIndex uint32,
	isRelocationPcRelative bool,
	length RelocationLength,
	isRelocationExtern bool,
	typ RelocationType,
) (RelocationInfo, error) {
	if symbolIndex > RelocationSymbolNumMask {
		return RelocationInfo{}, fmt.Errorf("symbol index too large")
	}

	details := symbolIndex | (uint32(length) << 25) | (uint32(typ) << 28)

	if isRelocationPcRelative {
		details |= RelocationPcRelative
	}

	if isRelocationExtern {
		details |= RelocationExtern
	}

	return RelocationInfo{Address: address, Details: details}, nil
}

func (r RelocationInfo) MarshalBinary() ([]byte, error) {
	return utils.GenericMarshalBinary(r)
}
