package video

import "traininggateway/domain"

func PublishedOnly() domain.CatalogFilter { return domain.CatalogFilter{PublishedOnly: true} }
func ByCategory(c string) domain.CatalogFilter {
	return domain.CatalogFilter{Category: c, PublishedOnly: true}
}
