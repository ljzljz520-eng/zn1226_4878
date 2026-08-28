package admin

import (
	"encoding/json"
	"traininggateway/domain"
)

func EncodeSummary(s domain.CardSummary) (string, error) {
	b, e := json.Marshal(s)
	return string(b), e
}
func StatusLabel(s domain.CardSummary) string {
	if s.Blocked > 0 {
		return "attention"
	}
	if s.Active == 0 {
		return "empty"
	}
	return "healthy"
}
