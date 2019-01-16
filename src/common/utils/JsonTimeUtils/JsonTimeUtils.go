package JsonTimeUtils

import (
    "fmt"
    "time"
)

// 格式化
const layout = "2006-01-02 15:04:05"
const layoutDate = "2006-01-02"

// 自定义一个jsonTime类型, 来处理类似于"2018-09-05T11:37:25.5303361+08:00"这样的RFC3339样式
type JsonTime time.Time
type JsonDate time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
    formatTimeByte := fmt.Sprintf("\"%s\"", time.Time(t).Format(layout))
    return []byte(formatTimeByte), nil
}

func (t JsonDate) MarshalJSON() ([]byte, error) {
    formatTimeByte := fmt.Sprintf("\"%s\"", time.Time(t).Format(layoutDate))
    return []byte(formatTimeByte), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
    if string(data) == "null" {
        return nil
    }
    // 这里timeParse返回的结果的第一个参数是time.Time类型, 还需要强制转换一次类型成JsonTime
    baseTime, err := time.Parse(`"`+layout+`"`, string(data))
    *t = JsonTime(baseTime)
    return err
}

func (t *JsonDate) UnmarshalJSON(data []byte) error {
    if string(data) == "null" {
        return nil
    }
    // 这里timeParse返回的结果的第一个参数是time.Time类型, 还需要强制转换一次类型成JsonTime
    baseTime, err := time.Parse(`"`+layoutDate+`"`, string(data))
    *t = JsonDate(baseTime)
    return err
}

func (t JsonTime) String() string {
    return time.Time(t).Format(layout)
}

func (t JsonDate) String() string {
    return time.Time(t).Format(layout)
}
