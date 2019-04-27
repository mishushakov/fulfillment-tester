GOOS=darwin go build -o ./bin/fulfillment-tester-darwin -ldflags="-s -w" ./main.go 
GOARCH=amd64 GOOS=linux go build -o ./bin/fulfillment-tester-linux -ldflags="-s -w" ./main.go
GOOS=windows go build -o ./bin/fulfillment-tester-win -ldflags="-s -w" ./main.go