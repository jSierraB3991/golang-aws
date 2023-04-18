deploy:
		./bin.sh
		sam deploy --stack-name go-serverless-sam --region us-east-1 --resolve-s3 --capabilities CAPABILITY_IAM  
		rm -rf OrderService/build
		rm -rf PaymentService/build
test:
		cowsay "Testing projects"
