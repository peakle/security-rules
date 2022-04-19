package insecureSkipVerify

import (
	"crypto/tls"
)

func warn() {
	print(tls.Config{MaxVersion: 1, InsecureSkipVerify: true}) // want `In this mode, TLS is susceptible to MITM attacks unless custom verification is used`

	_ = tls.Config{MaxVersion: 1, InsecureSkipVerify: true, MinVersion: 1} // want `In this mode, TLS is susceptible to MITM attacks unless custom verification is used`
	_ = tls.Config{InsecureSkipVerify: true, MinVersion: 1}                // want `In this mode, TLS is susceptible to MITM attacks unless custom verification is used`
	_ = tls.Config{InsecureSkipVerify: true}                               // want `In this mode, TLS is susceptible to MITM attacks unless custom verification is used`
}
