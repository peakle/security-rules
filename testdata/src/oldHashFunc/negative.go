package oldHashFunc

import (
	"crypto/md5"
	sha1 "crypto/sha256"
)

func _() {
	_ = sha1.New()

	print(md5.Size, md5.BlockSize)
}
