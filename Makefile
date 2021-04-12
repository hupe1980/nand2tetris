PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: p6
p6:
	@echo Generate hack files [project 6]
	@go run cmd/assembler/main.go projects/06/add/Add.asm
	@go run cmd/assembler/main.go projects/06/max/Max.asm
	@go run cmd/assembler/main.go projects/06/max/MaxL.asm
	@go run cmd/assembler/main.go projects/06/pong/Pong.asm
	@go run cmd/assembler/main.go projects/06/pong/PongL.asm
	@go run cmd/assembler/main.go projects/06/rect/Rect.asm
	@go run cmd/assembler/main.go projects/06/rect/RectL.asm

.PHONY: test
## test: Runs go test with default values
test:
	go test -v -race -count=1  ./...

.PHONY: help
## help: Prints this help message
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo