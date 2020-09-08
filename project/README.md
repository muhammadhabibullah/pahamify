## Pokemon Project

### Tech stack:
1. Go (1.14)
2. MySQL (5.7)
3. Docker Compose (3.3) & Docker

### Framework:
1. [Echo](https://github.com/labstack/echo)
2. [UUID](https://github.com/google/uuid)
3. [SQL-Migrate](https://github.com/rubenv/sql-migrate)
4. [Viper](https://github.com/spf13/viper)
5. [Gorm v2](https://gorm.io)

### Instalasi
+ Jalankan MySQL dengan docker compose. Pastikan port 3306 tidak digunakan di local machine.
```
docker-compose up --build
```

+ Configurasi file `config.json` untuk arahkan host ke IP docker (`172.19.0.1`)

+ Build program dengan docker.
```
docker build -t pokemon .
```

+ Run migration untuk mengisi data di database ketika pertama kali menjalankan program docker. Sesuaikan konfigurasi port yang tersedia di local machine.
```
docker run -p 3333:3000 pokemon /bin/bash -c "./pokemon migrate"
```

+ Jika program distop, untuk menjalankan program lagi cukup run docker yang sudah dibuild tanpa melakukan migrasi lagi. 
```
docker run -p 3333:3000 pokemon
```
