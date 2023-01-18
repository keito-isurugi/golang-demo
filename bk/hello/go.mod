module github.com/keito-isurugi/golang-demo/module/hello

go 1.18

// go mod edit -replace github.com/keito-isurugi/golang-demo/module/greeting=../greeting

replace github.com/keito-isurugi/golang-demo/module/greeting => ../greeting

require github.com/keito-isurugi/golang-demo/module/greeting v0.0.0-00010101000000-000000000000
