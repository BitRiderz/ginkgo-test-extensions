package ginkgo_test_extensions

import (
	. "github.com/onsi/ginkgo/v2"
)

func ItShould(str string, args ...interface{}) bool {
	return It("\n\tShould "+str, args)
}

func Given(str string, args ...interface{}) bool {
	return Describe("\n\tGiven "+str, args)
}

func When_(str string, args ...interface{}) bool {
	return Context("\n\tWhen "+str, args)
}

func And_(str string, args ...interface{}) bool {
	return Context("\n\tAnd "+str, args)
}
