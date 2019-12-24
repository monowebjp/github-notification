package main

import (
    . "./keys"
    "bytes"
    "context"
    "fmt"
    "github.com/shurcooL/githubv4"
    "golang.org/x/oauth2"
    "os"
    "strconv"
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
                //CommitContributionsByRepository []struct {
                //    Repository struct {
                //        NameWithOwner githubv4.String
                //    }
                //    Contributions struct {
                //        totalCount githubv4.Int
                //    }
                //}
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

    api := GetTwitterApi()

    countString := strconv.Itoa(int(query.User.ContributionsCollection.TotalCommitContributions))

    var buffer bytes.Buffer
    buffer.WriteString(t.Format("2006年1月2日（"))
    buffer.WriteString(replaceWeekDay(int(t.Weekday())))
    buffer.WriteString(t.Format("）のコミット数: "))
    buffer.WriteString(countString)
    buffer.WriteString("\n#botテスト")
    tweet, err := api.PostTweet(buffer.String(), nil)
    if err != nil {
        panic(err)
    }

    fmt.Println(tweet.Text)
}

func replaceWeekDay(weekDay int) string {
    japaneseWeekDay := [...] string{"日", "月", "火", "水", "木", "金", "土",}

    return japaneseWeekDay[weekDay]
}
