include .envrc

.PHONY: gen-docs
gen-docs:
	@swag init -g main.go -d cmd/api && swag fmt