# Technical Test Sharing Vision Indonesia

- **Backend**: Golang, Echo Framework, MySQL, GORM
- **Frontend**: React Vite, shadcn/ui, TanStack Query

Server backend berjalan di `http://localhost:4000`  
Frontend berjalan di `http://localhost:5173` (default Vite)

---

## ⚙️ Setup Backend (Golang)

### 1. Masuk ke folder backend

```bash
cd backend
```

### 2. Copy environment

```bash
cp .env.example .env
```

Isi konfigurasi MySQL:

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=yourdb
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Jalankan migration

```bash
go run cmd/migrate.go
```

### 5. Jalankan server

```bash
go run main.go
```

Server backend tersedia di:

```
http://localhost:4000
```

---

## ⚙️ Setup Frontend (React Vite)

### 1. Masuk ke folder frontend

```bash
cd frontend
```

### 2. Install dependencies

```bash
npm install
```

### 3. Jalankan aplikasi

```bash
npm run dev
```

Aplikasi frontend tersedia di:

```
http://localhost:5173
```

## 🛠️ Teknologi yang digunakan

- **Backend**
  - Go Lang
  - Echo Framework
  - MySQL
  - GORM
- **Frontend**
  - React Typescript
  - Vite
  - shadcn/ui
  - TanStack Query
