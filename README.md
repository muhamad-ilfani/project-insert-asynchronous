# Saving Many Datas to DB

## Technologies
1. Go Language
2. Echo Framework
3. Postgresql DB
4. Kafka Producer and Consumer

## DB Setting
Please see .env file. Default value as follow :
>DB_HOST=localhost  
DB_DRIVER=postgres  
DB_USER=xxxx    
DB_PASSWORD=xxxx    
DB_NAME=postgres    
DB_PORT=5432

Change above configuration based on your local connection for Postgresql.

Make sure user has access to create schema and table. This service will automathically create schema and table.

 Or you can directly create schema with name "servicea" and table with name "users"

## Kafka Setting
running in localhost:9092 (default)

## How to Run
1. Clone this repository with
    > git clone https://github.com/muhamad-ilfani/coding-test-be.git
2. Setting your DB connection and match it with .env file
3. Run kafka on local. If you are linux user, you can use this command:  
    -> Go to kafka directory    
    -> Run zookeeper    
    > bin/zookeeper-server-start.sh config/zookeeper.properties    

    -> Run kafka server 
    > bin/kafka-server-start.sh config/server.properties
4. Create topic kafka   
This topic is used for retry process if there is an error when insert data to database. Create topic with this command
    > bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic registration.notification.retry

    You can check topic list with
    > bin/kafka-topics.sh --bootstrap-server=localhost:9092 --list

5. Run service with command
    > go run main.go

6. Make sure service is running in port 8000 with command
    > curl -X GET localhost:8000/

    Response :
    > {"message":"welcome"}
7. You can insert data to DB with API : localhost:8000/register

    payload: 
    
    {
    "request_id":123456,
    "data" : [
        {
            "id":12345,
            "customer":"test1",
            "quantity":1,
            "price":10.00,
            "timestamp":"2022-01-01 22:10:44"
        },
        {
            "id":12346,
            "customer":"test2",
            "quantity":2,
            "price":20.00,
            "timestamp":"2022-01-01 22:10:44"
        }
    ]
    }