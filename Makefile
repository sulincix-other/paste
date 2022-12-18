STATIC_FLAGS=-buildmode=pie -ldflags '-linkmode external -extldflags "-static"'
SHELL=/bin/bash
build: clean
	mkdir -p build
	for file in $$(find src -type f) ; do \
	    echo "Generate: $$file build/`basename $$file`" ; \
	    cat $$file | gcc -E - -o build/`basename $$file` ; \
	    sed -i "s/#.*//g" build/`basename $$file` ; \
	done
	for file in $$(find data -type f) ; do \
	    fname=`basename $$file` ; \
	    bash tool/gendata.sh $${fname/./_} $$file > build/$$fname.go ; \
	done
	cd build ; set -x ; go build $${STATIC:+$(STATIC_FLAGS)}  -a -o paste *.go

clean:
	rm -rf build
