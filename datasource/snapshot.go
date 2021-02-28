package datasource

type SnapShot struct {
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
}

func (s *SnapShot) Insert() {
	_, err := engine.Insert(s)
	if err != nil {
		panic(err)
	}
}

func (s *SnapShot) CopyFromMatch(m *Match) {
	s.It = m.It
	s.Mid = m.Mid
	s.LeagueName = m.LeagueName
	s.TeamName = m.TeamName
	s.Min = m.Min
	s.Sec = m.Sec
	s.State = m.State
	s.HoScore = m.HoScore
	s.HoHalfScore = m.HoHalfScore
	s.GuScore = m.GuScore
	s.GuHalfScore = m.GuHalfScore
	s.HoRed = m.HoRed
	s.HoYellow = m.HoYellow
	s.HoCo = m.HoCo
	s.HoHfCo = m.HoHfCo
	s.GuRed = m.GuRed
	s.GuYellow = m.GuYellow
	s.GuCo = m.GuCo
	s.GuHfCo = m.GuHfCo
	s.HalfLet = m.HalfLet
	s.HalfLetHm = m.HalfLetHm
	s.HalfLetAw = m.HalfLetAw
	s.HalfAvgHm = m.HalfAvgHm
	s.HalfAvgAw = m.HalfAvgAw
	s.HalfAvgEq = m.HalfAvgEq
	s.HalfSize = m.HalfSize
	s.HalfSizeBig = m.HalfSizeBig
	s.HalfSizeSma = m.HalfSizeSma
	s.HalfFirstLet = m.HalfFirstLet
	s.HalfFirstLetHm = m.HalfFirstLetHm
	s.HalfFirstLetAw = m.HalfFirstLetAw
	s.HalfFirstAvgHm = m.HalfFirstAvgHm
	s.HalfFirstAvgAw = m.HalfFirstAvgAw
	s.HalfFirstAvgEq = m.HalfFirstAvgEq
	s.HalfFirstSize = m.HalfFirstSize
	s.HalfFirstSizeBig = m.HalfFirstSizeBig
	s.HalfFirstSizeSma = m.HalfFirstSizeSma

	s.Let = m.Let
	s.LetHm = m.LetHm
	s.LetAw = m.LetAw
	s.AvgHm = m.AvgHm
	s.AvgAw = m.AvgAw
	s.AvgEq = m.AvgEq
	s.Size = m.Size
	s.SizeBig = m.SizeBig
	s.SizeSma = m.SizeSma
	s.FirstLet = m.FirstLet
	s.FirstLetHm = m.FirstLetHm
	s.FirstLetAw = m.FirstLetAw
	s.FirstAvgHm = m.FirstAvgHm
	s.FirstAvgAw = m.FirstAvgAw
	s.FirstAvgEq = m.FirstAvgEq
	s.FirstSize = m.FirstSize
	s.FirstSizeBig = m.FirstSizeBig
	s.FirstSizeSma = m.FirstSizeSma
}
