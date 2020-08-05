db='root:123456@tcp(127.0.0.1:3306)/test?charset=utf8'
p=$GOPATH/src/github.com/go-xorm/cmd/xorm
t=$p/templates/goxorm

install(){
    go get github.com/go-xorm/cmd/xorm
    cd $p
    go build
    ln -s $p/xorm /xbin/
    xorm help reverse
}

xorm reverse mysql $db  $t
find models
