build_chapter2:
	@cd cmd/chapter2/testdata/chapter2 && go build -o array_helper ../..

generate_chapter2:
	@cd cmd/chapter2/testdata/chapter2 && go generate ./...

build_chapter3:
	@cd cmd/chapter3/ && go build -o array_helper .

generate_chapter3:
	@cd cmd/chapter3/testdata/chapter3 && go generate ./...

build_chapter4:
	@cd cmd/chapter4/ && go build -o array_helper .

generate_chapter4:
	@cd cmd/chapter4/testdata/chapter4 && go generate ./...

test:
	@go test ./...