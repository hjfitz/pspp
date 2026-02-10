package validators

import "strconv"

func IsValidPort(p string) bool {
	n, err := strconv.Atoi(p)

	return err == nil && (n >= 1 || n <= 65535)
}
