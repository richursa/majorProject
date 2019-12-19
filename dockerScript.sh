sudo docker exec -itd $(sudo docker run -itd -v /Desktop/majorProject/temp/data/db:/data/db blockutil) go run /app/main.go 
