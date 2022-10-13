PROTOS=$(shell find . \( -path ./google -o -path ./validate \) -prune -type f -o -name "*.proto")
GOPATH=$(shell go env GOPATH)

.PHONY: init
init:
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		&& go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		&& go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		&& go get -u google.golang.org/protobuf/cmd/protoc-gen-go \
		&& go get -u github.com/envoyproxy/protoc-gen-validate \
		&& go install \
			github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
			google.golang.org/protobuf/cmd/protoc-gen-go \
			google.golang.org/grpc/cmd/protoc-gen-go-grpc \
			github.com/envoyproxy/protoc-gen-validate \
        && mkdir -p google/api \
        && mkdir -p validate \
        && curl 'https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto' > google/api/http.proto \
        && curl 'https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto' > google/api/annotations.proto \
        && curl 'https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto' > google/api/httpbody.proto \
        && curl 'https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto' > google/api/field_behavior.proto \
        && curl 'https://raw.githubusercontent.com/bufbuild/protoc-gen-validate/main/validate/validate.proto' > validate/validate.proto

.PHONY: gen
gen:
	mkdir -p swagger \
	&& rm -rf swagger/* \
	&& protoc \
		--go_out=. --go_opt=module=github.com/transerver/protos \
		--go-grpc_out . --go-grpc_opt=module=github.com/transerver/protos \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt module=github.com/transerver/protos \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out ./swagger \
		--openapiv2_opt logtostderr=true \
  		--validate_out="lang=go:." \
		--validate_opt module=github.com/transerver/protos \
		$(PROTOS)
