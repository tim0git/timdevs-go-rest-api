start_db:
	docker run --detach --name dynamodb -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -inMemory

test:
	go test -count=1 ./...

debug:
	go test -count=1 ./... -v

coverage:
	go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out

run:
	PORT="8443" TABLE_NAME="Vehicles" AWS_ACCESS_KEY_ID="mock-key" AWS_SECRET_ACCESS_KEY="mock-secret" DYNAMODB_ENDPOINT="http://localhost:8000" go run main.go

build:
	go build -ldflags="-w -s" -o build

run_build:
	PORT="8443" TABLE_NAME="Vehicles" AWS_ACCESS_KEY_ID="mock-key" AWS_SECRET_ACCESS_KEY="mock-secret" DYNAMODB_ENDPOINT="http://localhost:8000" ./build

package:
	docker build -t vehicles-api -f Dockerfile .

run_package:
	docker run --add-host localhost:host-gateway -p8443:8443 --env-file .env.local vehicles-api

dive:
	CI=true dive vehicles-api --ci-config docker/.dive.yaml

clean:
	rm -rf build
