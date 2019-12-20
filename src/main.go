package main

import (
    //. "./key"
    "context"
    "fmt"
    "github.com/shurcooL/githubv4"
    "golang.org/x/oauth2"
    "os"
)

type node struct {
    Name             githubv4.String
    DefaultBranchRef struct {
        Target struct {
            Commit struct {
                History struct {
                    TotalCount githubv4.Int
                } `graphql:"history(since: \"2019-12-20T00:00:00+00:00\")"`
            } `graphql:"... on Commit"`
        }
    }
}

var query struct {
    Viewer struct {
        Login        githubv4.String
        CreatedAt    githubv4.DateTime
        Repositories struct {
            Nodes []node
        } `graphql:"repositories(first: 100)"`
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

    commitCount := 0

    for _, i := range query.Viewer.Repositories.Nodes {
        commitCount = commitCount + int(i.DefaultBranchRef.Target.Commit.History.TotalCount)
    }

    fmt.Println("CommitCount:", commitCount)

    //api := GetTwitterApi()

    //var buffer bytes.Buffer
    //buffer.WriteString("なんかできたっぽいから\n    Login:")
    //buffer.WriteString(string(query.Viewer.Login))
    //buffer.WriteString("\nCreatedAt:")
    //buffer.WriteString(query.Viewer.CreatedAt.Format("2006-01-02 03:04:05"))
    //tweet, err := api.PostTweet(buffer.String(), nil)
    //if err != nil {
    //    panic(err)
    //}

    //fmt.Print(tweet.Text)
}
