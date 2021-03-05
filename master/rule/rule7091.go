package rule

import (
	"niffler/common"
	"niffler/datasource"
)

// Rule7091 7091 rule
type Rule7091 struct {
}

func NewRule7091() Rule {
	r := &Rule7091{}
	return r
}

func (r *Rule7091) Name() string {
	return common.RULE_7091
}

func (r *Rule7091) Filter(m *datasource.Match) *datasource.Filter {
	return nil
}

func (r *Rule7091) LoadFromDb(key string) *datasource.Filter {
	return nil
}
