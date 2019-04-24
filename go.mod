module github.com/codcodog/boss

go 1.12

replace (
	golang.org/x/net v0.0.0-20180218175443-cbe0f9307d01 => github.com/golang/net v0.0.0-20180218175443-cbe0f9307d01
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
)

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/mattn/go-sqlite3 v1.10.0
)
