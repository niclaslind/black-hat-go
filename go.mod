module github.com/niclaslind/black-hat-go

go 1.14

replace github.com/niclaslind/black-hat-go/ch3/bing-metadata/metadata => ./ch3/bing-metadata/metadata

require (
	github.com/PuerkitoBio/goquery v1.6.1 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.1
)
