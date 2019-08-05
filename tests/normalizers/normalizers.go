package normalizers

import "github.com/revel/revel/testing"

type TestNormalizers struct {
	testing.TestSuite
}

func (t *TestNormalizers) Before() {
	println("Set up")
}

func (t *TestNormalizers) After() {
	println("Tear down")
}
