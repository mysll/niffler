package datasource

// BetDataSource 数据源
type BetDataSource interface {
	Name() string
	Init() bool
	Shut() bool
	Fetch() []*Match
	Update(m *Match) bool
}
