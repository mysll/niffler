package rule

import (
	"niffler/common"
	"niffler/datasource"
)

var Rules = make(map[string]Rule)

// Rule define
type Rule interface {
	Name() string
	Filter(m *datasource.Match) *datasource.Filter
	LoadFromDb(key string) *datasource.Filter
}

func init() {
	Rules[common.RULE_7091] = NewRule7091()
}
