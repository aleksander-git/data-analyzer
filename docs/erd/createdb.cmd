createdb.exe -h localhost -p 5432 -U postgres data-analyzer
psql -f data-analyzer-schema.sql -U postgres -h localhost -p 5432 -d data-analyzer

