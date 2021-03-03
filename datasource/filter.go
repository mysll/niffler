package datasource

import (
	"fmt"
	"log"
	"math"
	"niffler/chat"
	. "niffler/common"
)

const (
	FILTER_STATE_NONE  = 0
	FILTER_STATE_RED   = 1
	FILTER_STATE_BLACK = 2
)

type Filter struct {
	Id               int64
	Rule             string
	It               string `xorm:"index"`
	Mid              string
	LeagueName       string
	TeamName         string
	Min              int
	Sec              int
	State            int
	HoScore          int     // 主队进球
	HoHalfScore      int     // 主队半场得分
	GuScore          int     // 客队进球
	GuHalfScore      int     // 客队半场得分
	HoRed            int     // 主队红牌
	HoYellow         int     // 主队黄牌
	HoCo             int     // 主队角球
	HoHfCo           int     // 主队半场角球
	GuRed            int     // 客队红牌
	GuYellow         int     // 客队黄牌
	GuCo             int     // 客队角球
	GuHfCo           int     // 客队半场角球
	HalfLet          float64 // 上半场数据
	HalfLetHm        float64
	HalfLetAw        float64
	HalfAvgHm        float64
	HalfAvgAw        float64
	HalfAvgEq        float64
	HalfSize         float64
	HalfSizeBig      float64
	HalfSizeSma      float64
	HalfFirstLet     float64
	HalfFirstLetHm   float64
	HalfFirstLetAw   float64
	HalfFirstAvgHm   float64
	HalfFirstAvgAw   float64
	HalfFirstAvgEq   float64
	HalfFirstSize    float64
	HalfFirstSizeBig float64
	HalfFirstSizeSma float64
	Let              float64 // 下半场数据
	LetHm            float64
	LetAw            float64
	AvgHm            float64
	AvgAw            float64
	AvgEq            float64
	Size             float64
	SizeBig          float64
	SizeSma          float64
	FirstLet         float64
	FirstLetHm       float64
	FirstLetAw       float64
	FirstAvgHm       float64
	FirstAvgAw       float64
	FirstAvgEq       float64
	FirstSize        float64
	FirstSizeBig     float64
	FirstSizeSma     float64
	FilterState      int
	HalfState        int
	WaitOdd          bool
	Inactive         bool  // 未激活
	Created          int64 `xorm:"created"`
	extra            int   `xorm:"-"`
}

func (f *Filter) RuleAlias() string {
	return RuleAlias(f.Rule)
}

func (f *Filter) CheckActive(m *Match) {
	if f.Inactive {
		switch f.Rule {
		case RULE_HALF_05, RULE_334:
			if m.HalfSize-f.AboveSize() < 0.1 &&
				m.HalfSizeBig >= WAIT_ODD &&
				m.Score() == f.Score() {
				f.Inactive = false
				f.Update("inactive")
				msg := f.MakeRuleMessage(m)
				chat.SendToRecommend(msg)
			}
		case RULE_7091, RULE_757:
			if m.Size-f.AboveSize() < 0.1 &&
				m.SizeBig >= WAIT_ODD &&
				m.Score() == f.Score() {
				f.Inactive = false
				f.Update("inactive")
				msg := f.MakeRuleMessage(m)
				log.Println(msg)
				chat.SendToRecommend(msg)
			}
		}

	}
}

func (f *Filter) MakeResultMessage(reset bool, m *Match) string {
	if reset {
		return fmt.Sprintf(TEXT_INVALID, f.RuleAlias(), f.LeagueName, f.TeamName, m.Min, m.Sec, m.HoScore, m.GuScore)
	}

	if f.FilterState == FILTER_STATE_RED {
		return fmt.Sprintf(TEXT_RED, f.RuleAlias(), f.LeagueName, f.TeamName, m.Min, m.Sec, m.HoScore, m.GuScore)
	}
	if f.FilterState == FILTER_STATE_BLACK {
		if f.HalfState == STATUS_FIRSTHALF {
			return fmt.Sprintf(TEXT_BLACK_HALF, f.RuleAlias(), f.LeagueName, f.TeamName, m.HoScore, m.GuScore)
		}
		return fmt.Sprintf(TEXT_BLACK, f.RuleAlias(), f.LeagueName, f.TeamName, m.HoScore, m.GuScore)
	}
	return ""
}

func (f *Filter) Dogfall() int {
	ratio := math.Abs(f.FirstAvgEq - 3.3)
	if ratio > 2.3 {
		ratio = (ratio - 2.3) / 20
		if ratio > 1 {
			ratio = 0.99
		}
		ratio = (1 - ratio) * 50
	} else {
		ratio = 50 + (1-(ratio/2.3))*40
	}

	return int(ratio)
}

func (f *Filter) AboveSize() float64 {
	return float64(f.HoScore+f.GuScore) + 0.5
}

func (f *Filter) getData() [40]float64 {
	return [40]float64{
		float64(f.HoScore), float64(f.HoHalfScore), float64(f.GuScore), float64(f.GuHalfScore),
		f.HalfLet, f.HalfLetHm, f.HalfLetAw,
		f.HalfAvgHm, f.HalfAvgAw, f.HalfAvgEq,
		f.HalfSize, f.HalfSizeBig, f.HalfSizeSma,
		f.HalfFirstLet, f.HalfFirstLetHm, f.HalfFirstLetAw,
		f.HalfFirstAvgHm, f.HalfFirstAvgAw, f.HalfFirstAvgEq,
		f.HalfFirstSize, f.HalfFirstSizeBig, f.HalfFirstSizeSma,
		f.Let, f.LetHm, f.LetAw,
		f.AvgHm, f.AvgAw, f.AvgEq,
		f.Size, f.SizeBig, f.SizeSma,
		f.FirstLet, f.FirstLetHm, f.FirstLetAw,
		f.FirstAvgHm, f.FirstAvgAw, f.FirstAvgEq,
		f.FirstSize, f.FirstSizeBig, f.FirstSizeSma,
	}
}

func (f *Filter) MakeRuleMessage(m *Match) string {
	rescore := f.AboveSize()
	var half string
	var sb float64
	var aiExp float64
	if f.HalfState == STATUS_FIRSTHALF {
		half = "半场"
		sb = m.HalfSizeBig
		aiExp = forecastHalf(f.getData())
	} else {
		half = "全场"
		sb = m.SizeBig
		aiExp = forecast(f.getData())
	}
	s := fmt.Sprintf(TEXT_RULE_MSG,
		f.RuleAlias(), f.LeagueName, f.TeamName,
		m.Min, m.Sec,
		f.HoScore, f.GuScore, f.Dogfall(),
		half, rescore,
		sb,
		aiExp,
		f.It)
	return s
}

func (f *Filter) MakeNoticeOdd() string {
	return fmt.Sprintf(TEXT_NOTICE_ODD_MSG, f.RuleAlias(), f.LeagueName, f.TeamName, f.It)
}

func (f *Filter) Insert() {
	_, err := engine.Insert(f)
	if err != nil {
		panic(err)
	}
}

func (f *Filter) Update(col ...string) {
	engine.Id(f.Id).Cols(col...).Update(f)
}

func (f *Filter) Score() int {
	return f.HoScore + f.GuScore
}

func (f *Filter) LoadFromDB(it string, rule string) bool {
	b, err := engine.Where("it=? and rule=?", it, rule).Get(f)
	if err != nil {
		return false
	}
	return b
}

func (f *Filter) CopyFromMatch(m *Match) {
	f.It = m.It
	f.Mid = m.Mid
	f.LeagueName = m.LeagueName
	f.TeamName = m.TeamName
	f.Min = m.Min
	f.Sec = m.Sec
	f.State = m.State
	f.HoScore = m.HoScore
	f.HoHalfScore = m.HoHalfScore
	f.GuScore = m.GuScore
	f.GuHalfScore = m.GuHalfScore
	f.HoRed = m.HoRed
	f.HoYellow = m.HoYellow
	f.HoCo = m.HoCo
	f.HoHfCo = m.HoHfCo
	f.GuRed = m.GuRed
	f.GuYellow = m.GuYellow
	f.GuCo = m.GuCo
	f.GuHfCo = m.GuHfCo
	f.HalfLet = m.HalfLet
	f.HalfLetHm = m.HalfLetHm
	f.HalfLetAw = m.HalfLetAw
	f.HalfAvgHm = m.HalfAvgHm
	f.HalfAvgAw = m.HalfAvgAw
	f.HalfAvgEq = m.HalfAvgEq
	f.HalfSize = m.HalfSize
	f.HalfSizeBig = m.HalfSizeBig
	f.HalfSizeSma = m.HalfSizeSma
	f.HalfFirstLet = m.HalfFirstLet
	f.HalfFirstLetHm = m.HalfFirstLetHm
	f.HalfFirstLetAw = m.HalfFirstLetAw
	f.HalfFirstAvgHm = m.HalfFirstAvgHm
	f.HalfFirstAvgAw = m.HalfFirstAvgAw
	f.HalfFirstAvgEq = m.HalfFirstAvgEq
	f.HalfFirstSize = m.HalfFirstSize
	f.HalfFirstSizeBig = m.HalfFirstSizeBig
	f.HalfFirstSizeSma = m.HalfFirstSizeSma
	f.Let = m.Let
	f.LetHm = m.LetHm
	f.LetAw = m.LetAw
	f.AvgHm = m.AvgHm
	f.AvgAw = m.AvgAw
	f.AvgEq = m.AvgEq
	f.Size = m.Size
	f.SizeBig = m.SizeBig
	f.SizeSma = m.SizeSma
	f.FirstLet = m.FirstLet
	f.FirstLetHm = m.FirstLetHm
	f.FirstLetAw = m.FirstLetAw
	f.FirstAvgHm = m.FirstAvgHm
	f.FirstAvgAw = m.FirstAvgAw
	f.FirstAvgEq = m.FirstAvgEq
	f.FirstSize = m.FirstSize
	f.FirstSizeBig = m.FirstSizeBig
	f.FirstSizeSma = m.FirstSizeSma
}
