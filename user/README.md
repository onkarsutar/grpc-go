
docker build -t onkarsutar/grpc-go_server .

docker run -p 50001:50001 -it onkarsutar/grpc-go_server sh