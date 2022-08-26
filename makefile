lint:
	staticcheck ./...

docs:
	godoc -http=:8081

tidy:
	go mod tidy && go mod vendor

test:
	go test ./... -count=1

coverage:
	go test ./... -coverprofile=cover.out
	go tool cover -html=cover.out

deploy:
	sam deploy