package utils

import "fmt"

func CalcContentRange(name string, from, to, total int) string {
	return fmt.Sprintf("%s %d-%d/%d", name, from, to, total)
}
