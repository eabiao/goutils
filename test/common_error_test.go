package test

import (
	"fmt"
	. "github.com/eabiao/goutils/common"
	"testing"
)

func TestCommonError(t *testing.T) {
	defer Catch(func(err error) {
		fmt.Println("catch error", err)
	})

	doSomeThing()
}

func doSomeThing() {
	fmt.Println("do start")
	Throw(Error("duang~"))
	fmt.Println("do finish")
}
