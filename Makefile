start_db:
	docker run --detach --name dynamodb -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -inMemory

delete_db:
	docker rm -f dynamodb

create_table:
	aws dynamodb create-table --table-name Vehicles --attribute-definitions AttributeName=vin,AttributeType=S --key-schema AttributeName=vin,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --endpoint-url http://localhost:8000

test: start_db
	go test -count=1 ./... ; make delete_db

debug: start_db
	go test -count=1 ./... -v ; make delete_db

coverage: start_db
	go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out ; make delete_db

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

compose:
	docker-compose up -d

dive:
	CI=true dive vehicles-api --ci-config docker/.dive.yaml

clean:
	rm -rf build ; docker rm -f dynamodb ; docker rmi timdevs-go-rest-api-vehicle-api:latest ; docker rmi vehicles-api:latest ; docker-compose down ; docker image prune -f
