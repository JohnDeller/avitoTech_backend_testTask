# avitoTech_backend_testTask

Микросервис написан с использованием фреймворка Gin 

SQL файл для создания БД avitoTech_backend_testTask/schema/000001_init.up.sql
Замечу, в ТЗ было сказано, что ненужно реализовывать механизм миграций. Но я использовал утилиту
migrate для быстрого запуска файлов up и down (папка schema), так удобнее.

запуск go run cmd/main.go

Реализовано:
  Зачисление средств на баланс пользователя (если пользователя в БД нет - он создаётся, если есть - обновляется)
  
 POST  localhost:8000/usr/user-balance
  
  входной JSON
    {
    "id":1,
    "balance":500
    }
    
  выходной JSON
  {
    "id": 1
  }
  
  Резервирование средств 
  
 POST  localhost:8000/reserve/reserve-funds
  
  
   входной JSON
  {
    "id":2,
    "user_id":1,
    "order_id":15,
    "balance":15
}
    
  выходной JSON
 {
    "order_id": 15,
    "reserve_id": 2,
    "user_id": 1
}
  Разрезервирование средств
POST  localhost:8000/reserve/unreserve-funds
   входной JSON
  {
    "id":2,
    "user_id":1,
    "order_id":15,
    "balance":15
}
    
  выходной JSON
  {
    "id": 1,
    "order_id": 15,
    "reserve_id": 2
}
  Признание выручки
 POST  localhost:8000/reserve/confirm-profit
   
   входной JSON
  {
    "id":2,
    "user_id":1,
    "order_id":15,
    "balance":15
}
    
  выходной JSON
 {
    "order_id": 15,
    "reserve_id": 2,
    "user_id": 1
    "link to the csv file":"report.csv"
}
  Получение баланса пользователя
  
GET localhost:8000/usr/user-balance/1
  
   выходной JSON
 500
 
 
 При попытках сделать отрицательный баланс
 
  входной JSON
 {
    "id":1,
    "balance":-500
}

выходной JSON

{
    "message": "balance below zero"
}

localhost:8000/reserve/reserve-funds

 входной JSON
{
    "id":1,
    "user_id":1,
    "order_id":15,
    "balance":1500
}

выходной JSON

{
    "message": "balance below zero"
}

При резервировании создаются экземпляры сущностей заказа и операции резервирования, 
при разрезервировании или признании выручки - экземпляр сущности резервирования удаляется, а экземпляр заказа - 
остаётся (можно удалить так:POST  localhost:8000/orders/delete-orders входной JSON { "id":}
