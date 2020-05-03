package public

import (
	"fmt"
	"testing"
)

func TestCredit(t *testing.T) {
	programs := Programs{}.Find()

	fmt.Println(programs)

}
