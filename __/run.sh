#cd ..
#make swagger


## Chạy file .sql trên docker
## "cat ./db/main/001_init.sql | docker exec -i <container-name> psql -U <user> -d <database>"
cat ./../db/main/001_init.sql | docker exec -i hapo_postgres psql -U postgres -d postgres