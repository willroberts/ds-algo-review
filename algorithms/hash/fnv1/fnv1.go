package fnv1

// FNV_OFFSET_BASIS is the base value of the FNV1 and FNV1a hash values.
const FNV_OFFSET_BASIS uint64 = 14695981039346656037

// FNV_PRIME is the multiplier applied to the hash value in each iteration.
const FNV_PRIME uint64 = 1099511628211

// FNV1 implements the Fowler-Noll-Vo hash function.
// FNV1 operates on the FNV_OFFSET_BASIS, manipulating the value based on each incoming byte
// of data, which is multiplied by the FNV_PRIME before toggling all hash bits against the byte.
// See https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function for details.
func FNV1(data []byte) uint64 {
	hash := FNV_OFFSET_BASIS
	for _, b := range data {
		hash *= FNV_PRIME
		hash ^= uint64(b)
	}
	return hash
}

// FNV1a implements the 1a version of the Fowler-Noll-Vo hash function.
// FNV1a differs from FNV1 only byy the order in which the multiply and XOR is performed.
// It offers better performance than FNV1, with a slightly higher chance of collision.
func FNV1a(data []byte) uint64 {
	hash := FNV_OFFSET_BASIS
	for _, b := range data {
		hash ^= uint64(b)
		hash *= FNV_PRIME
	}
	return hash
}
