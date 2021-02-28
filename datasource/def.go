package datasource

const WAIT_ODD = 1.9

const (
	STATUS_UNKNOWN    = -1
	STATUS_NONE       = 0 // 未开始
	STATUS_FIRSTHALF  = 1 // 上半场
	STATUS_MIDDLE     = 2 // 中场
	STATUS_SECONDHALF = 3 // 下半场
	STATUS_COMPLETE   = 4 // 完赛
)

func State(s int) string {
	switch s {
	case STATUS_UNKNOWN:
		return "未知"
	case STATUS_NONE:
		return "未开始"
	case STATUS_FIRSTHALF:
		return "上半场"
	case STATUS_MIDDLE:
		return "中场"
	case STATUS_SECONDHALF:
		return "下半场"
	case STATUS_COMPLETE:
		return "完赛"
	default:
		return "未知"
	}
}

const (
	RULE_334     = "334"
	RULE_7091    = "7091"
	RULE_757     = "757"
	RULE_HALF_05 = "half0.5"
	RULE_HALF_EQ = "halfeq"
	RULE_LZ_001  = "lz001"
)

const (
	RULE_334_ALIAS     = "F4"
	RULE_7091_ALIAS    = "S+1"
	RULE_757_ALIAS     = "S7"
	RULE_HALF_05_ALIAS = "F0.5"
	RULE_HALF_EQ_ALIAS = "F0.5="
	RULE_LZ_001_ALIAS  = "S="
)

func RuleAlias(rule string) string {
	switch rule {
	case RULE_334:
		return RULE_334_ALIAS
	case RULE_7091:
		return RULE_7091_ALIAS
	case RULE_757:
		return RULE_757_ALIAS
	case RULE_HALF_05:
		return RULE_HALF_05_ALIAS
	case RULE_HALF_EQ:
		return RULE_HALF_EQ_ALIAS
	case RULE_LZ_001:
		return RULE_LZ_001_ALIAS
	default:
		return "unknown"
	}
}
