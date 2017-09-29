# Recalbox api

[![pipeline status](https://gitlab.com/MikaXII/recalbox-api/badges/master/pipeline.svg)](https://gitlab.com/MikaXII/recalbox-api/commits/master)

* New api for recalbox rewritten in go.
* Old api will be deprecated and no-compatiblity between this 2 differents version.

### With go

Currently I have no make a clean setup.
```
go get -u github.com/kardianos/govendor
go get -u gitlab.com/MikaXII/recalbox-api # or git clone the project inside your $GOPATH
cd $GOPATH/src/gitlab.com/MikaXII/recalbox-api
govendor sync 
go run main.go #or go build
```

### With docker (WIP)
```
docker build -t recalbox-api .
docker run -it --rm -p 8080:8080 -v </recalbox-folder>:/recalbox/share recalbox-api
```

recalbox-folder = same skeleton as /recalbox/share on your fs

| Method | Endpoint                   | Result                                     |
|--------|----------------------------|--------------------------------------------|
| GET    |  /                         |  version && list of available endpoint     |
| GET    | /v1/systems                |  get systems list                          |
| GET    | /v1/systems/:systemId/     | get all roms filename for specified system |
| POST   | /v1/systems/:systemId      | upload a rom file                          |
| GET    | /v1/systems/:systemId/hash | get all roms hash (CRC, MD5, SHA1)         |
| GET    | /v1/bios                   | get bios filename                          |