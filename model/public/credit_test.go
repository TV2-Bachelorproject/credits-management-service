package public

import (
	"fmt"
	"testing"
)

func TestCredit(t *testing.T) {
	credits := Credits{}.ForProgram(25)

	fmt.Println(credits[0].CreditGroup)

}
