#!/bin/sh

# source <this.sh>

set_path () {
	script=$1
	dir=$(dirname "$script")
	test -n "$dir" && dir=./
	cd "$dir"
	export GOPATH=`pwd`
	# go install: no install location for directory * outside GOPATH
	# For any OS X users and future me, you also need to set GOBIN to
	# avoid this confusing message on install and go get
	# http://stackoverflow.com/questions/18149601/go-install-always-fails-no-install-directory-outside-gopath
	export GOBIN=$GOPATH/bin
	cd - > /dev/null
}

set_path "$BASH_SOURCE"

echo "Install dependencies..."
go get gopkg.in/kataras/iris.v6
go get github.com/julienschmidt/httprouter
go get github.com/labstack/echo
go get github.com/valyala/fasthttp
go get github.com/gin-gonic/gin
go get github.com/emicklei/go-restful
go get github.com/koding/multiconfig
go get github.com/sirupsen/logrus
go get github.com/satori/go.uuid
echo "Done."
