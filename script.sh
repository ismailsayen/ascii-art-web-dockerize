
IMG=my_img

CTR=my_ctr

docker rmi -f $(docker images -a -q)

docker rm -f $(docker ps -q)

docker build --no-cache -t $IMG .

docker container run -d -p 8081:8080 --name $CTR $IMG

docker exec -it my_ctr  /bin/bash