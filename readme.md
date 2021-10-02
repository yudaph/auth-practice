# Cara Installasi

## Jika Menggunakan Docker

1. Install Docker

   Untuk cara menginstall docker bisa dilihat pada link [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)

2. Run Docker Compose

   Setelah selesai menginstall docker dan setting variable, dapat menjalankan command berikut untuk menjalankan applikasi :

   ```
   docker compose up
   ```

## Jika tidak menggunakan docker

1. MySQL

   import database termasuk dengan struktul table pada file `init.sql` di dalam folder `MySQL`

2. Config file

   Config file terdapat didalam file `.config` atur variable sesuai dengan perangkatmu, dengan variable `MYSQL_NAME` adalah nama database

   ```
   APP_PORT=8080
   MYSQL_USERNAME=root
   MYSQL_PASSWORD=
   MYSQL_HOST=localhost
   MYSQL_PORT=3306
   MYSQL_NAME=users
   ```

3. Aktifkan godotenv

   Aktifkan line 23 pada file `app.go` di dalam folder `api` untuk membaca environtment variable dalam config file

# API

## Registrasi

Menggunakan metode `POST` pada url `http://localhost:8080/api/v1/users/`
menggunakan body dengan format JSON

```
{
    "email": "test1@email.com",
    "address": "alamatnya",
    "password": "123456",
    "password-confirmation": "123456"
}
```

## Melihat Semua User

Menggunakan metode `GET` pada url `http://localhost:8080/api/v1/users/`

## Melihat Satu User

Menggunakan metode `GET` pada url `http://localhost:8080/api/v1/users/(id_user)` contohnya `http://localhost:8080/api/v1/users/1xgCoeIThojxJydzdRFzfea5ENO`

## Mengubah User

Menggunakan metode `PATCH` pada url `http://localhost:8080/api/v1/users/(id_user)` contohnya `http://localhost:8080/api/v1/users/1xgCoeIThojxJydzdRFzfea5ENO`

dengan contoh body :

```
{
    "email": "emailbaru@email.com",
    "address": "alamat baru",
}
```

## Menghapus User

Menggunakan metode `DELETE` pada url `http://localhost:8080/api/v1/users/(id_user)` contohnya `http://localhost:8080/api/v1/users/1xgCoeIThojxJydzdRFzfea5ENO`

## Login

Menggunakan metode `POST` pada url `http://localhost:8080/api/v1/auth/`
dengan contoh body :

```
{
    "email": "emailbaru@email.com",
    "password": "123456",
}
```

## Change Password

Menggunakan metode `POST` pada url `http://localhost:8080/api/v1/auth/change-password`
dengan contoh body :

```
{
    "old-password": "123456",
    "password": "123456",
    "passwordConfirmation": "123456"
}
```

Jangan lupa menyertakan token dengan tipe `Bearer Token` pada `header` saat melakukan request penggantian password
