package lexical

import (
	"fmt"
	"testing"
)

func init() {

}

func TestNew(t *testing.T) {
	var a Analysis
	a.New("tHhes5e are 1some words")
	fmt.Printf("%#v", a)
}
