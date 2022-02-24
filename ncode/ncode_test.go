package ncode_test

import (
	"fmt"
	"testing"

	"github.com/devarchi33/common/ncode"
	"github.com/devarchi33/goutils/test"
)

func Test(t *testing.T) {
	n := int64(60466175)

	code, err := ncode.Encode(n)
	test.Ok(t, err)
	fmt.Println("code:", code)

	number, err := ncode.Decode(code)
	test.Ok(t, err)
	fmt.Println("number:", number)

	test.Equals(t, n, number)
}
