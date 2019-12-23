package main

import (
    "context"
    "fmt"
    "github.com/shurcooL/githubv4"
    "golang.org/x/oauth2"
    "os"
    "time"
)

//import (
//    "context"
//    "fmt"
//    "github.com/shurcooL/githubv4"
//    "golang.org/x/oauth2"
//    "os"
//)
//
//func main() {
//    src := oauth2.StaticTokenSource(
//        &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
//    )
//    httpClient := oauth2.NewClient(context.Background(), src)
//
//    client := githubv4.NewClient(httpClient)
//
//    type Language struct {
//        Name  githubv4.String
//        Color githubv4.String
//    }
//
//    type Repository struct {
//        NameWithOwner githubv4.String
//        Url           githubv4.String
//        Languages     struct {
//            Nodes []struct {
//                Language `graphql:"... on Language"`
//            }
//        } `graphql:"languages(first: 5)"`
//    }
//
//    var query struct {
//        Search struct {
//            Nodes []struct {
//                Repository `graphql:"... on Repository"`
//            }
//        } `graphql:"search(first: 5, query: $q, type: $searchType)"`
//    }
//
//    variables := map[string]interface{}{
//        "q":          githubv4.String("monowebjp"),
//        "searchType": githubv4.SearchTypeRepository,
//    }
//
//    err := client.Query(context.Background(), &query, variables)
//    if err != nil {
//        fmt.Println(err)
//    }
//
//    for _, repo := range query.Search.Nodes {
//        fmt.Println("---------")
//        fmt.Println(repo.NameWithOwner)
//        fmt.Println(repo.Url)
//        for _, lang := range repo.Languages.Nodes {
//            fmt.Println(lang.Name)
//            fmt.Println(lang.Color)
//        }
//    }
//}

func main() {
    src := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
    )
    httpClient := oauth2.NewClient(context.Background(), src)
    client := githubv4.NewClient(httpClient)

    //type node struct {
    //    Name             githubv4.String
    //    DefaultBranchRef struct {
    //        Target struct {
    //            Commit struct {
    //                History struct {
    //                    TotalCount githubv4.Int
    //                } `graphql:"history(since: $t)"`
    //            } `graphql:"... on Commit"`
    //        }
    //    }
    //}
    //
    //var query struct {
    //    Viewer struct {
    //        Repositories struct {
    //            Nodes []node
    //        } `graphql:"repositories(first: 100)"`
    //    }
    //}
    //
    //variables := map[string]interface{}{
    //    "t": githubv4.GitTimestamp(time.Now().Format("2006-01-02T00:00:00+09:00")),

    //{
    //   user(login: "monowebjp") {
    //   name
    //   email
    //   contributionsCollection(from: "2018-12-23T00:00:00", to: "2019-12-23T23:59:59") {
    //       totalRepositoryContributions
    //       totalCommitContributions
    //       commitContributionsByRepository {
    //           repository {
    //               nameWithOwner
    //           }
    //           contributions {
    //               totalCount
    //           }
    //       }
    //   }
    //}
    //}

    var query struct {
        User struct {
            Name                    githubv4.String
            ContributionsCollection struct {
                TotalRepositoryContributions githubv4.Int
                //    TotalCommitContributions        githubv4.Int
                //    CommitContributionsByRepository []struct {
                //        Repository struct {
                //            NameWithOwner githubv4.String
                //        }
                //        Contributions struct {
                //            totalCount githubv4.Int
                //        }
                //    }
            } `graphql:"contributionsCollection(from: $from, to: $to)"`
        } `graphql:"user(login: \"monowebjp\")"`
    }

    variables := map[string]interface{}{
        "from": githubv4.DateTime(time.Now().Format("2006-01-02T00:00:00")),
        "to":   githubv4.DateTime(time.Now().Format("2006-01-02T00:00:00")),
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

    //var buffer bytes.Buffer
    //buffer.WriteString("2019年12月20日（火）のコミット数: ")
    //buffer.WriteString(commitCount)
    //tweet, err := api.PostTweet(buffer.String(), nil)
    //if err != nil {
    //    panic(err)
    //}

    fmt.Println(query.User.ContributionsCollection.TotalRepositoryContributions)
}
