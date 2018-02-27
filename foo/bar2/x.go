package bar2

import (
	"fmt"

	_ "github.com/rogpeppe/test/v2/foo/bar1"
)

func init() {
	fmt.Println("bar2")
}
