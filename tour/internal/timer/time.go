package timer

import "time"
/**
放回当前本地时间的Time对象。
此处的封装主要是为了便于后续对Time对象做进一步的统一处理
 */
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}
/**
获取当前时间 Timer 加上 duration 后得到的最终时间
 */
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)	// 从字符串中解析出 duration 持续时间
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
