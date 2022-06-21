package time_ticker

const  (
    TimeTickerDayWorkErr="----------每天执行一次的定时器失败-----------"

)



//注册定时器
func RegisterTimeTicker(){
    go DayilyWork()
}
