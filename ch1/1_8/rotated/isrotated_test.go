package rotated_test

import (
	"fmt"
	"github.com/smudugal/gocrackcode/ch1/1_8/rotated"
	"testing"
)

func TestTrue(t *testing.T) {
	fmt.Println("Testing the true condition")

	var original_string string = "waterbottle"
	var rotated_string string = "erbottlewat"

	if rotated.IsRotated(original_string, rotated_string) != true {
		t.Error("Expected TRUE got FALSE")
	}
}

func TestFalse(t *testing.T) {
	fmt.Println("Testing the false condition")

	var original_string string = "waterbottle"
	var rotated_string string = "erhelelewat"

	if rotated.IsRotated(original_string, rotated_string) != false {
		t.Error("Expected FALSE got TRUE")
	}
}

func TestDifferentLength(t *testing.T) {
	fmt.Println("Testing 2 differnt length strings")

	var original_string string = "waterbottle"
	var rotated_string string = "erbottle"

	if rotated.IsRotated(original_string, rotated_string) != false {
		t.Error("Expected FALSE got TRUE")
	}
}

func TestZeroLength(t *testing.T) {
	fmt.Println("Testing 0 length strings")

	var original_string string = "waterbottle"
	var rotated_string string = ""

	if rotated.IsRotated(original_string, rotated_string) != false {
		t.Error("Expected FALSE got TRUE")
	}
}
