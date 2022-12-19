STATIC_FLAGS=-buildmode=pie -ldflags '-linkmode external -extldflags "-static"'
SHELL=/bin/bash
build: clean
	mkdir -p build
	GOPATH=`pwd`/modules go mod download github.com/mattn/go-sqlite3
	@for file in $$(find src -type f) ; do \
	    echo "Generate: $$file build/`basename $$file`" ; \
	    cat $$file | gcc -E - -o build/`basename $$file` ; \
	    sed -i "s/#.*//g" build/`basename $$file` ; \
	done
	@for file in $$(find data -type f) ; do \
	    fname=`basename $$file` ; \
	    echo "Generate: $$file build/$$fname.go" ;\
	    bash tool/gendata.sh $${fname/./_} $$file text/$${fname/*./} > build/$$fname.go ; \
	done
	cd build ; set -x ; GOPATH=`pwd`/../modules go build $${STATIC:+$(STATIC_FLAGS)}  -a -o paste *.go

clean:
	mkdir -p build
	chmod 777 -R build
	rm -rf build go.sum
