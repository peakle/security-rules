package oldHashFunc

import (
	"crypto/md5"
	"crypto/sha1"
)

func warn() {
	_ = md5.New()  // want `Don't use old hash functions with short digest size`
	_ = sha1.New() // want `Don't use old hash functions with short digest size`

	print(sha1.New()) // want `Don't use old hash functions with short digest size`
	print(md5.New())  // want `Don't use old hash functions with short digest size`
}
