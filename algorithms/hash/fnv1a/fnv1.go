package fnv1a

const FNV_OFFSET_BASIS uint64 = 14695981039346656037
const FNV_PRIME uint64 = 1099511628211

func FNV1(data []byte) uint64 {
	hash := FNV_OFFSET_BASIS
	for _, b := range data {
		hash *= FNV_PRIME
		hash ^= uint64(b)
	}
	return hash
}

func FNV1a(data []byte) uint64 {
	hash := FNV_OFFSET_BASIS
	for _, b := range data {
		hash ^= uint64(b)
		hash *= FNV_PRIME
	}
	return hash
}
