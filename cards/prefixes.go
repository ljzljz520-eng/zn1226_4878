package cards

func PrefixDescription(prefix string) string {
	switch prefix {
	case "TRN":
		return "general training"
	case "SAFE":
		return "security training"
	case "OPS":
		return "operations training"
	default:
		return "unknown"
	}
}
func PrefixQuota(prefix string) int {
	switch prefix {
	case "TRN":
		return 1000
	case "SAFE":
		return 500
	case "OPS":
		return 750
	default:
		return 0
	}
}
func PrefixColor(prefix string) string {
	switch prefix {
	case "TRN":
		return "blue"
	case "SAFE":
		return "green"
	case "OPS":
		return "orange"
	default:
		return "gray"
	}
}
func PrefixRequiresReview(prefix string) bool { return prefix == "SAFE" || prefix == "OPS" }
