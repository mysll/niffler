package datasource

import (
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	engine *xorm.Engine
)

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	engine.Sync2(new(Match), new(Filter), new(SnapShot), new(Attention))

	engine.DatabaseTZ = time.Local
	engine.TZLocation = time.Local
}
