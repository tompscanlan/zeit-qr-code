module zeitqrcode

go 1.13

require (
	common v0.0.0

	github.com/DataDog/zstd v1.4.4 // indirect
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/go-oidc v2.1.0+incompatible
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/sessions v1.2.0
	github.com/kidstuff/mongostore v0.0.0-20181113001930-e650cd85ee4b
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/skip2/go-qrcode v0.0.0-20191027152451-9434209cb086
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.2.0
	golang.org/x/crypto v0.0.0-20191219195013-becbf705a915 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
)

replace common => ./common
