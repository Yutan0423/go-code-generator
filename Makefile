build_chapter2:
	@cd cmd/chapter2/testdata/chapter2 && go build -o array_helper ../../../

generate_chapter2:
	@cd cmd/chapter2/testdata/chapter2 && go generate ./...