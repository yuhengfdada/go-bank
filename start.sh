# !/bin/sh
echo "migration start..."
./migrate -path ./migrations -database "postgresql://root:secret@postgres:5432/simple_bank?sslmode=disable" -verbose up
echo "migration done"
echo "begin executing main.go..."
./main