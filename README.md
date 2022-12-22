# Sulincix Paste
Simple paste service written go

## Features
* uses sqlite3 and simple html (nojs)
* paste pages are simple text
* less size ( < 10mb )

## Installing
### From source
1. run `make` or `make static`
2. copy build/paste file
### From docker image
```shell
docker run -d \
    -p 2023:<port> \
    registry.gitlab.com/sulincix/paste \
    /paste
```
