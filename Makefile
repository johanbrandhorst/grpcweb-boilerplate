generate:
	protoc -I. -Ivendor/ ./proto/web.proto \
		--gopherjs_out=plugins=grpc:$$GOPATH/src \
		--go_out=plugins=grpc:$$GOPATH/src
	go generate ./frontend/

clean:
	rm -f ./proto/client/* ./proto/server/* ./cert.pem ./key.pem \
		./frontend/html/frontend.js ./frontend/html/frontend.js.map

install:
	go install ./vendor/github.com/golang/protobuf/protoc-gen-go \
		./vendor/github.com/johanbrandhorst/protobuf/protoc-gen-gopherjs \
		./vendor/github.com/foobaz/go-zopfli \
		./vendor/github.com/gopherjs/gopherjs

generate_cert:
	go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" \
		--host=localhost,127.0.0.1 \
		--ecdsa-curve=P256 \
		--ca=true

serve:
	go run main.go
