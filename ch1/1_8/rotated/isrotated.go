package rotated

import (
	"strings"
)

func IsRotated(orig_str string, rotated_str string) bool {

	root_string := orig_str + orig_str

	if len(orig_str) != len(rotated_str) {
        return false
	}

    return strings.Contains(root_string, rotated_str)
}
