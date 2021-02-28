package datasource

// Attention data for at team
type Attention struct {
	Id   int64
	Team string `xorm:"index"`
}

// Insert to db
func (a *Attention) Insert() {
	_, err := engine.Insert(a)
	if err != nil {
		panic(err)
	}
}

// Remove from db
func (a *Attention) Remove() {
	engine.Where("team=?", a.Team).Delete(a)
}

// Find from db
func (a *Attention) Find(team string) bool {
	b, err := engine.Where("team=?", team).Get(a)
	if err != nil {
		return false
	}
	return b
}
