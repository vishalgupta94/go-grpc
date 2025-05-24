# 🚀 Project Setup Guide

## 🐳 PostgreSQL with Docker

### Start a PostgreSQL Container

```bash
docker run --name drizzle-postgres -e POSTGRES_PASSWORD=mypassword -d -p 5433:5432 postgres
```

### Connect to PostgreSQL

Replace `<containerId>` with your actual container ID:

```bash
docker exec -e PGPASSWORD=mypassword -it <containerId> psql -U postgres -d postgres
```

Example:

```bash
docker exec -e PGPASSWORD=mypassword -it 27f0ca9d76b5 psql -U postgres -d postgres
```

---

## 💼 Project Workspaces Setup

### Create Directories

```bash
mkdir client server proto
```

### Initialize Go Modules

```bash
cd client && go mod init go-grpc/client
cd ../proto && go mod init go-grpc/proto
cd ../server && go mod init go-grpc/server
```

---

## 📁 Workspace Structure

```
go-grpc/
├── client/      # gRPC client code
├── proto/       # Protocol buffer definitions
└── server/      # gRPC server code
```

---

Happy coding! 🚀
