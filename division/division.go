package division

import (
	"fmt"
)

func Division(a, b int) (int, error) {
	if b == 0 {
		return -1, fmt.Errorf("на ноль делить нельзя")
	}
	return a / b, nil
}
