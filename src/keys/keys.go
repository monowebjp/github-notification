package keys

import (
    "github.com/ChimeraCoder/anaconda"
    "os"
)

func GetTwitterApi() *anaconda.TwitterApi {
    anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
    anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
    api := anaconda.NewTwitterApi(os.Getenv("TWITTER_API_KEY"), os.Getenv("TWITTER_API_SECRET"))
    return api
}
