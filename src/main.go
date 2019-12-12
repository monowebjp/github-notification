package main

import (
    . "./keys"
    . "fmt"
)

func main(){

    api := GetTwitterApi()

    text := "テスト・テスト"
    tweet, err := api.PostTweet(text, nil)
    if err != nil {
        panic(err)
    }

    Print(tweet.Text)
}
