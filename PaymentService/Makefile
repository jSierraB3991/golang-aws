run:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main cmd/app/main.go
		zip -jrm build/payment.zip build/main
