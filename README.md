# DDD

## 🚀 Start the Container

Run the following command in the directory containing `docker-compose.yml`:

```bash
docker compose up -d
```

This starts PostgreSQL in the background.

---

## 🔑 Access the Container

Connect to PostgreSQL with:

```bash
docker exec -it go-ddd bash
```

---

## 🔑 Access the Database

Connect to PostgreSQL with:

```bash
docker exec -it go-ddd psql -U postgres -d postgres
```
