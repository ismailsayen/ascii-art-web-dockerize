
IMG=my_img

CTR=my_ctr

docker build --no-cache -t $IMG .

docker container run -d -p 8081:8080 --name $CTR $IMG

