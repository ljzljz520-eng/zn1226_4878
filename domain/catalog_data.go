package domain

type CatalogEntry struct {
	Code, Title, Category string
	Minutes               int
}

func SeedCatalog() []CatalogEntry {
	return []CatalogEntry{
		{"L001", "Welcome to internal learning", "orientation", 8}, {"L002", "Security foundations", "security", 14}, {"L003", "Password hygiene", "security", 11}, {"L004", "Device handling", "security", 12}, {"L005", "Incident reporting", "security", 15}, {"L006", "Privacy principles", "privacy", 16}, {"L007", "Data classification", "privacy", 13}, {"L008", "Retention rules", "privacy", 10}, {"L009", "Support workflow", "operations", 9}, {"L010", "Escalation practice", "operations", 18}, {"L011", "Quality checklist", "quality", 12}, {"L012", "Review conversations", "quality", 17},
		{"L013", "Change management", "operations", 13}, {"L014", "Release readiness", "operations", 15}, {"L015", "Accessibility", "quality", 10}, {"L016", "Inclusive language", "quality", 9}, {"L017", "Remote collaboration", "culture", 12}, {"L018", "Feedback skills", "culture", 14}, {"L019", "Coaching basics", "culture", 16}, {"L020", "Mentoring habits", "culture", 15},
		{"L021", "Network boundaries", "security", 12}, {"L022", "Phishing signals", "security", 13}, {"L023", "Safe browsing", "security", 11}, {"L024", "Mobile security", "security", 10}, {"L025", "Backups", "operations", 9}, {"L026", "Recovery drills", "operations", 14}, {"L027", "Service ownership", "operations", 13}, {"L028", "Runbook writing", "operations", 17},
		{"L029", "Metrics literacy", "quality", 12}, {"L030", "Experiment design", "quality", 16}, {"L031", "Root cause analysis", "quality", 18}, {"L032", "Decision records", "quality", 14}, {"L033", "Conflict navigation", "culture", 13}, {"L034", "Meeting design", "culture", 10}, {"L035", "Focus time", "culture", 8}, {"L036", "Healthy boundaries", "culture", 11},
		{"L037", "Threat modeling", "security", 19}, {"L038", "Authorization", "security", 15}, {"L039", "Audit trails", "security", 13}, {"L040", "Secrets handling", "security", 14}, {"L041", "Queue management", "operations", 12}, {"L042", "Capacity planning", "operations", 18}, {"L043", "On-call basics", "operations", 16}, {"L044", "Handoffs", "operations", 10},
		{"L045", "Test strategy", "quality", 19}, {"L046", "Contract tests", "quality", 15}, {"L047", "Load testing", "quality", 18}, {"L048", "Release notes", "quality", 9}, {"L049", "Listening", "culture", 8}, {"L050", "Clear writing", "culture", 10}, {"L051", "Pairing", "culture", 12}, {"L052", "Community", "culture", 14},
	}
}
func CatalogByCategory(category string) []CatalogEntry {
	out := []CatalogEntry{}
	for _, e := range SeedCatalog() {
		if e.Category == category {
			out = append(out, e)
		}
	}
	return out
}
func CatalogMinutes(category string) int {
	n := 0
	for _, e := range CatalogByCategory(category) {
		n += e.Minutes
	}
	return n
}
