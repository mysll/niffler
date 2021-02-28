package fortune

import (
	"log"
	"niffler/config"
	"niffler/datasource"
	"niffler/datasource/leisu"
)

// Niffler main class
type Niffler struct {
	ds  datasource.BetDataSource
	cfg *config.Config
}

// NewNiffler create new niffler
func NewNiffler(config *config.Config) *Niffler {
	inst := &Niffler{
		cfg: config,
	}

	switch config.DataSource {
	case "bet365":
	case "leisu":
		inst.ds = leisu.New(config)
	}
	return inst
}

func (n *Niffler) prepare() {
	//chat.SendToRecommend("初始化⚽[尴尬][红包][炸弹][忍者]")
}

// Run main process
func (n *Niffler) Run() {
	n.prepare()
	if n.ds == nil {
		log.Fatalln("no datasource define")
	}
	if n.ds != nil {
		if !n.ds.Init() {
			log.Fatalf("init %s datasource failed", n.ds.Name())
		}
	}

}
