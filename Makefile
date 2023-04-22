.DEFAULT_GOAL := build-all

export GO111MODULE=on
export DATE=$(shell date '+%Y%m%d%H%M')
export GOPROXY=https://goproxy.io

firstmakefile:


scan-deps:
	@mkdir -p bin && bash version
	@cp -r conf bin/
	@cp -r static bin/

scan: go-deps
	go build -i -o bin/scan-$(DATE) ./pkg/cmd/scan

all: scan

clean:
	@rm -rf bin

go-deps: scan-deps
	go get -v github.com/fsnotify/fsnotify@v1.4.7
	go get github.com/go-errors/errors
	go get gopkg.in/gormigrate.v1
	go get github.com/Shopify/sarama@v1.29.1
	go get github.com/bsm/sarama-cluster
	go get github.com/bitly/go-simplejson
	go get github.com/influxdata/influxdb/client/v2
	go get github.com/vmihailenco/msgpack
	go get github.com/tidwall/gjson
	go get github.com/goburrow/cache
	go get github.com/satori/go.uuid
	go get github.com/go-errors/errors
	go get github.com/gin-contrib/static
	go get github.com/go-sql-driver/mysql
	go get github.com/go-redis/redis
	go get github.com/golang/glog
	go get github.com/labstack/gommon/log
	go get github.com/spf13/viper
	go get github.com/gin-gonic/gin/binding@v1.6.3
