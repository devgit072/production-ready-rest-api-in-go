Docker command to run postgres locally:
sudo docker run --name prod-rest-api -e POSTGRES_PASSWORD=postgres -p5432:5432 -d postgres
Check if docker postgres is running:
docker ps


