package domain

type CardSummary struct{ Total, Active, Used, Expired, Blocked, Exhausted int }
type CatalogFilter struct {
	Category, Query string
	PublishedOnly   bool
}
type Page[T any] struct {
	Items                []T
	Offset, Limit, Total int
}

func (s CardSummary) Healthy() bool { return s.Active > 0 && s.Blocked == 0 }
