hello:
	@echo "hello"

env:
	CGO_ENABLED=1 GOOS=ios GOARCH=arm64

update:
	go get -u ./...

entg:
	go run ./cmd/entg .

entmigrate:
	go run ./cmd/api core migrate

# make name=Website entinit
entinit:
	go run -mod=mod entgo.io/ent/cmd/ent new $(name)

wire:
	cd wireapp && wire