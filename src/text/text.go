package text

import (
    "fmt"
    "time"
)

func ChooseTweet() string {

    TweetList := [31]string{
        "投稿の内容を書く",
        //省略
    }

    d := time.Now().Day()
    m := time.Now().Month()

    TweetContent := TweetList[d-1]
    TweetOfToday := fmt.Sprintf("【%d月%d日】\n %s", m, d, TweetContent)
    return TweetOfToday
}
