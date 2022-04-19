package insecureSkipVerify

import (
	"crypto/tls"
)

func _() {
	print(tls.Config{MaxVersion: 1, InsecureSkipVerify: false})

	_ = tls.Config{MaxVersion: 1, MinVersion: 1}
	_ = tls.Config{}
	_ = tls.Config{InsecureSkipVerify: false, MinVersion: 1}
}
