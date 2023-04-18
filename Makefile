deploy:
		./bin.sh
		sam deploy --stack-name go-serverless-sam --region us-east-1 --resolve-s3 --capabilities CAPABILITY_IAM  
test:
		cowsay "Testing projects"
