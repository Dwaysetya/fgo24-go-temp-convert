package calculate

import "fmt"

func CkeK(c int) string {
	convert := c + 273
	return fmt.Sprintf("%d°C = %dK", c, convert)
}

func CkeF(c int) string {
	convert := (9 * c / 5) + 32
	return fmt.Sprintf("%d°C = %d°F", c, convert)
}

func CkeR(c int) string {
	convert := 4 * c / 5
	return fmt.Sprintf("%d°C = %d°R", c, convert)
}
