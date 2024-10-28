package xfile

import (
	"fmt"
	"strconv"
	"strings"
)

// Size Convert byte count to a string with units
func Size(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// ToBytes converts strings such as "1MB", "1M", "1GB", "1G" to int64 bytes
func ToBytes(sizeStr string) (int64, error) {
	// Remove spaces and convert them to uppercase while maintaining consistency
	sizeStr = strings.TrimSpace(strings.ToUpper(sizeStr))

	var multiplier int64 = 1
	if len(sizeStr) > 1 {
		switch sizeStr[len(sizeStr)-2:] {
		case "KB":
			multiplier = 1 << 10 // 1KB = 2^10 bytes
			sizeStr = sizeStr[:len(sizeStr)-2]
		case "MB":
			multiplier = 1 << 20 // 1MB = 2^20 bytes
			sizeStr = sizeStr[:len(sizeStr)-2]
		case "GB":
			multiplier = 1 << 30 // 1GB = 2^30 bytes
			sizeStr = sizeStr[:len(sizeStr)-2]
		case "TB":
			multiplier = 1 << 40 // 1TB = 2^40 bytes
			sizeStr = sizeStr[:len(sizeStr)-2]
		case "PB":
			multiplier = 1 << 50 // 1PB = 2^50 bytes
			sizeStr = sizeStr[:len(sizeStr)-2]
		case "EB":
			multiplier = 1 << 60 // 1EB = 2^60 bytes
			sizeStr = sizeStr[:len(sizeStr)-2]
		default:
			// Check for single character suffix
			switch sizeStr[len(sizeStr)-1:] {
			case "K":
				multiplier = 1 << 10 // 1K = 2^10 bytes
				sizeStr = sizeStr[:len(sizeStr)-1]
			case "M":
				multiplier = 1 << 20 // 1M = 2^20 bytes
				sizeStr = sizeStr[:len(sizeStr)-1]
			case "G":
				multiplier = 1 << 30 // 1G = 2^30 bytes
				sizeStr = sizeStr[:len(sizeStr)-1]
			case "T":
				multiplier = 1 << 40 // 1T = 2^40 bytes
				sizeStr = sizeStr[:len(sizeStr)-1]
			case "P":
				multiplier = 1 << 50 // 1P = 2^50 bytes
				sizeStr = sizeStr[:len(sizeStr)-1]
			case "E":
				multiplier = 1 << 60 // 1E = 2^60 bytes
				sizeStr = sizeStr[:len(sizeStr)-1]
			}
		}
	}

	// Convert the remaining string to an integer
	value, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		return 0, err
	}

	// Calculate the total number of bytes
	return value * multiplier, nil
}
