go mod download

cd migrations

go run main.go down && go run main.go up

cd ..

go run main.go

