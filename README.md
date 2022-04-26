
# Potentivio apps

![Logo](https://github.com/ALTA-Potentivio/BE-Potentivio/blob/main/logo.png)

Readme in english [click here](https://github.com/ALTA-Potentivio/BE-Potentivio/blob/readme/README-en.md).

Potentivio apps adalah sebuah aplikasi yang menjembatani antara artist dengan cafe untuk menjalin sebuah hubungan bisnis. Dengan aplikasi ini, cafe akan lebih mudah dalam mencari artist dan dari sisi artist berkemungkinan mendapatkan job lebih banyak dan mudah. 


## Fitur Cafe 

- Register
- Login User
- Edit User
- User cafe dapat mencari artist seusai kategori, genre, lokasi dan mengurutkan sesuai dengan harga tertinggi atau terendah
- User cafe dapat menghire artist sesuai dengan waktu yang di tentukan
- User cafe dapat membatalkan artist yang sudah di boking dengan alasan yang valid
- User cafe dapat melakukan pemabayaran melalui xendit yang sudah di integrasikan
- User cafe dapat memberikan rating dan comment ketika acara sudah selesai

## Fitur artis

- Register
- Login User
- Edit User
- User artist akan mendapatkan informasi jika ada cafe yang mengundang mereka
- User artist dapat menolak atau menerima tawaran
- User artist dapat melakukan cancel dari job yang sudah diterima

## Open APIs

Untuk Open API bisa lihat selengkapnya [disini](https://app.swaggerhub.com/apis-docs/satriacening/project_group3/1.0.0#/)


## Menjalankan Lokal

Kloning project

```bash
  $ git clone git@github.com:ALTA-BE7-SATRIA-PUTRA/group_project3.git
```

Masuk ke direktori project

```bash
  $ cd ~/project
```
Buat `database` baru

Buat sebuah file dengan nama di dalam folder root project `.env` dengan format dibawah ini. Sesuaikan configurasi di komputer lokal

```bash
  export APP_PORT=""
  export JWT_SECRET="S3CR3T"
  export DB_PORT="3306"
  export DB_DRIVER="mysql"
  export DB_NAME=""
  export DB_ADDRESS="127.0.0.1"
  export DB_USERNAME=""
  export DB_PASSWORD=""
```

Jalankan aplikasi 

```bash
  $ go run main.go
```


## Authors

- [@satriacening](https://github.com/satriacening)
- [@usamaha07](https://github.com/usamaha07)
- [@abbiyudha](https://github.com/abbiyudha)

 
