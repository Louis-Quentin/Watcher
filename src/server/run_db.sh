set -e

Container="Area_postgres"

[ "$(docker ps -a | grep $Container)" ] && docker stop $Container
[ "$(docker ps -a | grep $Container)" ] && docker rm $Container
docker build -t area_db .
docker run -d -p 5432:5432 --name $Container area_db
sleep 2
docker exec -it $Container bash -c "sh /tmp/entrypoint.sh"
