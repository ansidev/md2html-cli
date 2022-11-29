.PHONY: prepare test mdtest

prepare:
	@echo "Installing golangci-lint"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Install Husky"
	@go install github.com/go-courier/husky/cmd/husky@latest && husky init

test:
	@go test -v ./...

demo:
	@go run main.go -output dist/index.html --minify --style 'light' test/fixtures/index.md
	@go run main.go -output dist/sample.min.html --excerpt-separator '<!-- more -->' --minify --style 'light' test/fixtures/sample.md
	@go run main.go -output dist/sample_with_custom_template.min.html --excerpt-separator '<!-- more -->' --minify --style 'light' --template test/fixtures/template.html test/fixtures/sample.md
	@go run main.go -output dist/sample_with_custom_markdown_style.html --excerpt-separator '<!-- more -->' --style 'dark' --template test/fixtures/template.html test/fixtures/sample.md
	@go run main.go -output dist/sample.html --excerpt-separator '<!-- more -->' --style 'light' test/fixtures/sample.md
