package cards

import "fmt"

func ValidateBatch(prefix string, count int) error {
	if prefix == "" {
		return fmt.Errorf("empty prefix")
	}
	if count < 1 || count > 1000 {
		return fmt.Errorf("count out of range")
	}
	return nil
}
func ParseExpiry(v int64) int64 {
	if v < 0 {
		return 0
	}
	return v
}
