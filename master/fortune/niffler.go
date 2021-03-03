package fortune

import (
	"fmt"
	"log"
	"niffler/chat"
	. "niffler/common"
	"niffler/config"
	"niffler/datasource"
	"niffler/datasource/leisu"
	"niffler/master/rule"
	"time"
)

// Niffler main class
type Niffler struct {
	ds     datasource.BetDataSource
	cfg    *config.Config
	matchs map[string]*datasource.Match
	filter map[string]map[string]*datasource.Filter
}

// NewNiffler create new niffler
func NewNiffler(config *config.Config) *Niffler {
	inst := &Niffler{
		cfg:    config,
		matchs: make(map[string]*datasource.Match),
		filter: make(map[string]map[string]*datasource.Filter),
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

func (n *Niffler) addMatch(match *datasource.Match) {
}

func (n *Niffler) snapShoot(e int, m *datasource.Match) {
	switch e {
	case EVENT_UPDATE_TIME:
		if m.Min%5 == 0 {
			s := new(datasource.SnapShot)
			s.CopyFromMatch(m)
			s.Insert()
		}
	case EVENT_CHANGE_STATE:
		switch m.State {
		case STATUS_FIRSTHALF, STATUS_MIDDLE, STATUS_SECONDHALF, STATUS_COMPLETE:
			s := new(datasource.SnapShot)
			s.CopyFromMatch(m)
			s.Insert()
		}
	}
}

func (n *Niffler) watchIfPresent(m *datasource.Match) bool {
	it := m.It
	// try load from db
	match := new(datasource.Match)
	if match.Load(it) {
		log.Printf("load match %s", it)
		n.addMatch(match)
		match.Update(m)
		return true
	}
	*match = *m
	if !match.Init() {
		log.Printf("match init failed, %s", it)
		return false
	}
	n.addMatch(match)
	match.Insert()
	msg := fmt.Sprintf("[新增] %s %s", match.LeagueName, match.TeamName)
	log.Println(msg)
	chat.SendToBroadcast(msg)
	chat.SendToBroadcast(match.Preview())
	return true
}

func (n *Niffler) filterMatchState(event int, m *datasource.Match) {
}

func (n *Niffler) filterMatch(m *datasource.Match) {
	for _, rs := range config.Setting.Rule.Update {
		if r, ok := rule.Rules[rs]; ok {
			if fi := r.Filter(m); fi != nil {
				fi.Insert()
				n.filter[m.It][rs] = fi
				if !fi.Inactive {
					msg := fi.MakeRuleMessage(m)
					log.Println(msg)
					chat.SendToRecommend(msg)
				}
			}
		}
	}
}

func (n *Niffler) checkActive(m *datasource.Match) {
	if fs, ok := n.filter[m.It]; ok {
		for _, v := range fs {
			if !v.Inactive {
				continue
			}
			v.CheckActive(m)
		}
	}
}

func (n *Niffler) checkFilter(event int, m *datasource.Match) {
}

func (n *Niffler) detect() bool {
	matchs := n.ds.Fetch()
	for _, m := range matchs {
		it := m.It
		var match *datasource.Match
		var ok bool
		if match, ok = n.matchs[it]; !ok {
			n.watchIfPresent(m)
			continue
		}
		events := match.Update(m)
		for _, e := range events {
			n.snapShoot(e, match)
			switch e {
			case EVENT_UPDATE_TIME:
				n.filterMatch(match)
				n.checkActive(match)
			default:
				n.filterMatchState(e, match)
				n.checkFilter(e, match)
			}
		}
	}
	return true
}

func (n *Niffler) work() {
	timer := time.NewTicker(time.Second)
L:
	for {
		select {
		case <-timer.C:
			if !n.detect() {
				log.Println("detect failed, quit")
				break L
			}
		}
	}
	timer.Stop()
}

// Run main process
func (n *Niffler) Run() {
	n.prepare()
	if n.ds == nil {
		log.Fatalln("no datasource define")
	}
	retrys := 0
	delay := time.Second * 5
	for {
		if !n.ds.Init() {
			log.Printf("init %s datasource failed", n.ds.Name())
			retrys++
			if retrys > 3 {
				delay = time.Minute
			}
			time.Sleep(delay)
			continue
		}
		log.Printf("datasource %s init ok", n.ds.Name())
		n.work()
	}
}
