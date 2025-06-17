#Usage to run the project

go run .

#To Init new db 
docker-compose up -d

#Command args for the project
1. reset - Resets the database
2. init-db - Inits the database
3. populate-db - Populates the database

for a fresh start
go run . reset init-db populate-db

