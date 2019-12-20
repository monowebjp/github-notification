package historySince

import (
    "bytes"
    "time"
)

func GetTodayDate() string {
    t := time.Now()
    var buffer bytes.Buffer
    buffer.WriteString(t.Format("2006-01-02T00:00:00+00:00"))

    return buffer.String()
}
