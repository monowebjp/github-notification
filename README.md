# github-notification
Twitterに毎日その日のコミット数をTweetするbot。

## .env
ルートに配置し、下記を記入する
```
GITHUB_TOKEN=
TWITTER_CONSUMER_KEY=
TWITTER_CONSUMER_SECRET=
TWITTER_API_KEY=
TWITTER_API_SECRET=
```

## セットアップコマンド
```
docker-compose build
```

```
docker-compose up -d
```

```
docker-compose exec app go get
```

```
docker-compose exec app go build main.go
```

```
docker-compose run app env
```

## 起動コマンド
```
docker-compose exec app go run main.go
```
