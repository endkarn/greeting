# GREETING PROJECT
## คำอธิบาย

เป็นโปรเจคเกี่ยวกับการสร้าง REST API ด้วยภาษา Go โดยจะเอา API ไปเรียกข้อมูลใน greeting ใน database ออกมาแสดง ซึ่ง database ที่ใช้เป็น PostgreSQL

## Installment
- [Golang version 1.10 ขึ่้นไป](https://golang.org/dl/)
- [Docker](https://docs.docker.com/install/)

## How to build and how to use

ในโปรเจคนี้มี 

- dockerfile ไว้สำหรับ build application 
- docker-compose สำหรับ run service
- makefile เพื่อให้ง่ายต่อการทำงาน โดยมีคำสั่ง build-app, up, down

#### Build Application
1. build application ให้เป็น docker images

```
make build-app
```

#### Run Application
2. สั่ง docker-compose up เพิ่ม start PostgreSQL และ greeting

```
make up
```

#### Stop Application
3. หยุดการทำงานของ application
```
control + c 
make down
```
