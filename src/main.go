package main

import (
    . "./keys"
    "bytes"
    "context"
    "fmt"
    "github.com/shurcooL/githubv4"
    "golang.org/x/oauth2"
    "os"
)

var query struct {
    Viewer struct {
        Login     githubv4.String
        CreatedAt githubv4.DateTime
    }
}

func main() {
    src := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
    )
    httpClient := oauth2.NewClient(context.Background(), src)
    client := githubv4.NewClient(httpClient)
    err := client.Query(context.Background(), &query, nil)
    if err != nil {
        // Handle error.
    }
    fmt.Println("    Login:", query.Viewer.Login)
    fmt.Println("CreatedAt:", query.Viewer.CreatedAt)

    api := GetTwitterApi()

    var buffer bytes.Buffer
    buffer.WriteString("なんかできたっぽいから環境変数のもたせ方変えてみる")
    //buffer.WriteString("なんかできたっぽいから\n    Login:")
    //buffer.WriteString(string(query.Viewer.Login))
    //buffer.WriteString("\nCreatedAt:")
    //buffer.WriteString(query.Viewer.CreatedAt.Format("2006-01-02 03:04:05"))
    //text := "test"
    tweet, err := api.PostTweet(buffer.String(), nil)
    if err != nil {
        panic(err)
    }

    fmt.Print(tweet.Text)
}
