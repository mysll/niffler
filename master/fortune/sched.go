package fortune

import (
	"github.com/robfig/cron"
)

var c = cron.New()

func sched() {
	/*C.AddFunc("0 0 22 * * 0,2,4", func() {
		h := ssq.Histroy(true)
		if len(h) > 0 {
			chat.SendToRecommend(fmt.Sprintf("update ok, %v", h[0]))
		}
	})*/
	c.Start()
}
