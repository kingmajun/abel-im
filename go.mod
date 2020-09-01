module abel-im

go 1.14

require (
	github.com/astaxie/beego v1.10.1
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.5 // indirect
	github.com/jonboulle/clockwork v0.2.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kingmajun/king-orm v0.0.0-20200901062423-1a1edff7dbe7
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.1.3 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.4 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	golang.org/x/tools v0.0.0-20200117065230-39095c1d176c // indirect
	google.golang.org/grpc v1.30.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

replace go.etcd.io/bbolt v1.3.4 => github.com/coreos/bbolt v1.3.4
