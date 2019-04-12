package othello

// MostSignificantBit returns the most significant bit in a bitset
func MostSignificantBit(x uint64) uint64 {

	shift := uint(0)

	if x >= 1<<32 {
		x >>= 32
		shift += 32
	}
	if x >= 1<<16 {
		x >>= 16
		shift += 16
	}
	if x >= 1<<8 {
		x >>= 8
		shift += 8
	}
	return msb8tab[x] << shift
}

var msb8tab = [256]uint64{
	0, 1, 2, 2, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 8, 8,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128}

//
var doMoveToLowerLookup = [64]uint64{
	// no bits are flipped if player to move has no discs on the inspected line
	^(uint64(0)),
	(uint64(1) << 1) - 1,
	(uint64(1) << 2) - 1,
	(uint64(1) << 3) - 1,
	(uint64(1) << 4) - 1,
	(uint64(1) << 5) - 1,
	(uint64(1) << 6) - 1,
	(uint64(1) << 7) - 1,
	(uint64(1) << 8) - 1,
	(uint64(1) << 9) - 1,
	(uint64(1) << 10) - 1,
	(uint64(1) << 11) - 1,
	(uint64(1) << 12) - 1,
	(uint64(1) << 13) - 1,
	(uint64(1) << 14) - 1,
	(uint64(1) << 15) - 1,
	(uint64(1) << 16) - 1,
	(uint64(1) << 17) - 1,
	(uint64(1) << 18) - 1,
	(uint64(1) << 19) - 1,
	(uint64(1) << 20) - 1,
	(uint64(1) << 21) - 1,
	(uint64(1) << 22) - 1,
	(uint64(1) << 23) - 1,
	(uint64(1) << 24) - 1,
	(uint64(1) << 25) - 1,
	(uint64(1) << 26) - 1,
	(uint64(1) << 27) - 1,
	(uint64(1) << 28) - 1,
	(uint64(1) << 29) - 1,
	(uint64(1) << 30) - 1,
	(uint64(1) << 31) - 1,
	(uint64(1) << 32) - 1,
	(uint64(1) << 33) - 1,
	(uint64(1) << 34) - 1,
	(uint64(1) << 35) - 1,
	(uint64(1) << 36) - 1,
	(uint64(1) << 37) - 1,
	(uint64(1) << 38) - 1,
	(uint64(1) << 39) - 1,
	(uint64(1) << 40) - 1,
	(uint64(1) << 41) - 1,
	(uint64(1) << 42) - 1,
	(uint64(1) << 43) - 1,
	(uint64(1) << 44) - 1,
	(uint64(1) << 45) - 1,
	(uint64(1) << 46) - 1,
	(uint64(1) << 47) - 1,
	(uint64(1) << 48) - 1,
	(uint64(1) << 49) - 1,
	(uint64(1) << 50) - 1,
	(uint64(1) << 51) - 1,
	(uint64(1) << 52) - 1,
	(uint64(1) << 53) - 1,
	(uint64(1) << 54) - 1,
	(uint64(1) << 55) - 1,
	(uint64(1) << 56) - 1,
	(uint64(1) << 57) - 1,
	(uint64(1) << 58) - 1,
	(uint64(1) << 59) - 1,
	(uint64(1) << 60) - 1,
	(uint64(1) << 61) - 1,
	(uint64(1) << 62) - 1,
	(uint64(1) << 63) - 1}