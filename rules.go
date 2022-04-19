package rules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

var Bundle = dsl.Bundle{}

func httpWithoutSSL(m dsl.Matcher) {
	m.Match(`$_ = $str`,
		`$_ := $str`,
		`var $_ $*_= $str`,
		`const $_ $*_ = $str`,
		`const $_ = $str`,
		`$_($*_, $str, $*_)`,
	).
		Where(m["str"].Type.Underlying().Is("string") && m["str"].Const && m["str"].Text.Matches(`http://`)).
		Report(`HTTP links are not secure`)
}

func insecureSkipVerify(m dsl.Matcher) {
	m.Match(`tls.Config{$*_, InsecureSkipVerify: true, $*_}`).
		Report(`In this mode, TLS is susceptible to MITM attacks unless custom verification is used`)
}

func oldTLS(m dsl.Matcher)                {}
func passwordInCode(m dsl.Matcher)        {}
func oldHashFunc(m dsl.Matcher)           {}
func oldCryptoAlgorithms(m dsl.Matcher)   {}
func swaggerBodyValidation(m dsl.Matcher) {}
