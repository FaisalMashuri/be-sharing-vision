# Backend Sharing Vising

### Soal No 1.
Query SQL : 

`CREATE DATABASE article;`

`USE article`

`CREATE TABLE posts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(200),
  content TEXT,
  status ENUM('Publish', 'Draft', 'Thrash') DEFAULT 'Draft'
  category VARCHAR(100),
  created_date TIMESTAMP,
  updated_date TIMESTAMP 
);`

### Soal No 2.
Sudah ada auto migrate, jadi ketika program dijalankan jika table belum ada, maka akan langsung terbuat table dengan kolom kolom yang sudah di define

### Soal No 3.
Sudah Sudah dibuat


### Soal No 4
Sudah dibuat

### Soal No 5
Sudah Ada collection postman


# Untuk Menjalankan
- buka terminal dan arahkan ke project folder, ketikan docker-compose up
- Lalu untuk membuka dbms, kunjungi `localhost:8080`
  - pada field server isi dengan `mysql`
  - pada field pengguna(user) isi dengan `docker`
  - pada field password isi dengan `admin`
  - pada field basis data(database) isi dengan `submission`

- Buka terminal baru dan arahkan ke folder project, lalu ketik `go run main.go`
  - lalu kunjungi localhost:3000

- Endpoint
  - POST : /article
  - GET : /article/:offset/:limit
  - GET : /article/:id
  - PUT : /article/:id
  - DELETE : /article/:id
