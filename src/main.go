package main

import (
    "bytes"
    "context"
    "fmt"
    "github.com/shurcooL/githubv4"
    "golang.org/x/oauth2"
    "os"
    "time"
)

func main() {
    src := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
    )
    httpClient := oauth2.NewClient(context.Background(), src)
    client := githubv4.NewClient(httpClient)

    var query struct {
        User struct {
            Name                    githubv4.String
            ContributionsCollection struct {
                TotalRepositoryContributions    githubv4.Int
                TotalCommitContributions        githubv4.Int
                CommitContributionsByRepository []struct {
                    Repository struct {
                        NameWithOwner githubv4.String
                    }
                    Contributions struct {
                        totalCount githubv4.Int
                    }
                }
            } `graphql:"contributionsCollection(from: $from, to: $to)"`
        } `graphql:"user(login: $name)"`
    }

    t := time.Now()
    t1 := t.Truncate(time.Hour).Add(time.Duration(- t.Hour()) * time.Hour)
    t2 := t1.AddDate(0, 0, 1).Add(time.Duration(- 1) * time.Second)

    variables := map[string]interface{}{
        "name": githubv4.String("monowebjp"),
        "from": githubv4.DateTime{t1},
        "to":   githubv4.DateTime{t2},
    }

    err := client.Query(context.Background(), &query, variables)
    if err != nil {
        fmt.Println(err)
    }

    //commitCount := 0
    //
    //for _, i := range query.ContributionsCollection.Repositories.Nodes {
    //    commitCount = commitCount + int(i.DefaultBranchRef.Target.Commit.History.TotalCount)
    //}

    //api := GetTwitterApi()

    fmt.Println(query.User.ContributionsCollection.TotalRepositoryContributions)
    fmt.Println(string(query.User.ContributionsCollection.TotalRepositoryContributions))

    var buffer bytes.Buffer
    buffer.WriteString("2019年12月20日（火）のコミット数: ")
    buffer.WriteString(string(query.User.ContributionsCollection.TotalRepositoryContributions))
    //tweet, err := api.PostTweet(buffer.String(), nil)
    //if err != nil {
    //   panic(err)
    //}

    fmt.Println(buffer.String())
}
