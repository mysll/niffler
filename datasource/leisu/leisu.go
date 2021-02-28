package leisu

import (
	"niffler/config"
	"niffler/datasource"
)

// LeiSu lei su data source
type LeiSu struct {
	config *config.Config
}

// New leisu data source
func New(config *config.Config) datasource.BetDataSource {
	l := &LeiSu{
		config: config,
	}
	return l
}

// Name return LeiSu
func (l *LeiSu) Name() string {
	return "LeiSu"
}

// Init init the data source
func (l *LeiSu) Init() bool {
	return true
}

// Shut close the data source
func (l *LeiSu) Shut() bool {
	return false
}

// Fetch all matches
func (l *LeiSu) Fetch() []*datasource.Match {
	return nil
}

// Update single match
func (l *LeiSu) Update(m *datasource.Match) bool {
	return false
}
