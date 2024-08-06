.DEFAULT_GOAL:=build

BINARY:=password-api

clean:
	rm -f ${BINARY}


build: clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BINARY}
	gosec .