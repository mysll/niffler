package datasource

import (
	"fmt"
	"math"
)

type Match struct {
	Id               int64
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
	Let              float64 // 全场数据
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
}

func (m *Match) Init() bool {
	return false
}

func (m *Match) Score() int {
	return m.HoScore + m.GuScore
}

func (m *Match) Dogfall() int {
	ratio := math.Abs(m.FirstAvgEq - 3.3)
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

func (m *Match) Load(it string) bool {
	b, err := engine.Where("it=?", it).Get(m)
	if err != nil {
		return false
	}
	return b
}

func (m *Match) Insert() {
	_, err := engine.Insert(m)
	if err != nil {
		panic(err)
	}
}

func (m *Match) Preview() string {
	return fmt.Sprintf(TEXT_PREVIEW,
		m.LeagueName, m.TeamName,
		m.FirstAvgHm, m.FirstAvgEq, m.FirstAvgAw,
		m.FirstLet, m.FirstLetHm, m.FirstLetAw,
		m.FirstSize, m.FirstSizeBig, m.FirstSizeSma,
		m.HalfFirstAvgHm, m.HalfFirstAvgEq, m.HalfFirstAvgAw,
		m.HalfFirstLet, m.HalfFirstLetHm, m.HalfFirstLetAw,
		m.HalfFirstSize, m.HalfFirstSizeBig, m.HalfFirstSizeSma,
		m.It,
	)
}

func (m *Match) String() string {
	return fmt.Sprintf(TEXT_FULL,
		m.LeagueName, m.TeamName,
		m.Min, m.Sec,
		m.HoScore, m.GuScore,
		m.AvgHm, m.AvgEq, m.AvgAw,
		m.Let, m.LetHm, m.LetAw,
		m.Size, m.SizeBig, m.SizeSma,
		m.FirstAvgHm, m.FirstAvgEq, m.FirstAvgAw,
		m.FirstLet, m.FirstLetHm, m.FirstLetAw,
		m.FirstSize, m.FirstSizeBig, m.FirstSizeSma,
		m.HalfFirstAvgHm, m.HalfFirstAvgEq, m.HalfFirstAvgAw,
		m.HalfFirstLet, m.HalfFirstLetHm, m.HalfFirstLetAw,
		m.HalfFirstSize, m.HalfFirstSizeBig, m.HalfFirstSizeSma,
	)
}

func (m *Match) Update(latest *Match) []int {
	var event []int
	return event
}
