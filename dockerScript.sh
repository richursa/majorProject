sudo docker exec -itd $(sudo docker run -itd -v /app/temp/data/db:/data/db blockverify) go run /app/main.go 
