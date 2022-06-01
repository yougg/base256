package base256

var (
	uint2byte = make(map[uint64]byte)
	byte2uint = make(map[byte]uint64)
	baseLen   uint64
)

func init() {
	var k uint64
	var v byte
	for range [256]struct{}{} {
		uint2byte[k] = v
		byte2uint[v] = k
		k++
		v++
	}
	baseLen = uint64(len(uint2byte))
}

// DecimalToBase256 convert decimal to base256
func DecimalToBase256(num uint64) (chars []byte) {
	for remainder, target := uint64(0), byte(0); num != 0; {
		remainder = num % baseLen
		target = uint2byte[remainder]
		chars = append(chars, target)
		num = num / baseLen
	}
	// reverse
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return
}

// Base256ToDecimal convert base256 to decimal
func Base256ToDecimal(chars []byte) uint64 {
	length := len(chars)
	if length == 0 || length > 8 {
		return 0
	}
	var num uint64
	exponential := length - 1
	for _, value := range chars {
		dec := byte2uint[value]
		num = num + dec*powUint(baseLen, uint64(exponential))
		exponential = exponential - 1
	}
	return num
}

// powUint calculate n to the mth power n^m
func powUint(n, m uint64) (v uint64) {
	if m == 0 {
		return 1
	}
	v = n
	for i := uint64(2); i <= m; i++ {
		v *= n
	}
	return
}
