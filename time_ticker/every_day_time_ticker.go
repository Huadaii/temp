package time_ticker

import (
	"log"
	"time"
)

//一天执行一次定时器

func DayilyWork() {
	for {
		//凌晨
		night := time.Now().AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
		parseNight, parseNightErr := time.ParseInLocation("2006-01-02 15:04:05", night, time.Local)
		if parseNightErr != nil {
			log.Println("一天执行一次定时器失败")
			break
		}
		var now = time.Now()
		ticker := time.NewTimer(parseNight.Sub(now))
		<-ticker.C //进行阻塞直到时间到，c管道中输出
		log.Println("----开始执行一天执行一次定时器----")
		//这下面是定时器执行的事情

		log.Println(TimeTickerDayWorkErr, "------------")

	}
}
