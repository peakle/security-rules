package rules

import (
    "github.com/quasilyte/go-ruleguard/dsl"
)

var Bundle = dsl.Bundle{}

func httpWithoutSSL(m dsl.Matcher)        {}
func insecureSkipVerify(m dsl.Matcher)    {}
func oldTLS(m dsl.Matcher)                {}
func passwordInCode(m dsl.Matcher)        {}
func oldHashFunc(m dsl.Matcher)           {}
func oldCryptoAlgorithms(m dsl.Matcher)   {}
func swaggerBodyValidation(m dsl.Matcher) {}
