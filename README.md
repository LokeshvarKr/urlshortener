# This is URL shortener service

1. You will be able to get short url for a long url
2. You will be able to get your original url from short url


## To build docker image
    docker build -t urlshortner


## To run docker container 
    
    docker run -d -p 8081:8081 urlshortener

or
    
    docker run -d -p 8081:8081 <urlshortener ImageId>


## To check if service is runnig inside container
go to any web browser and type http://localhost:8081/home and send request 

you will be able to see output like "URL shortner home page"


## To make request for getting short url for a long url

go to terminal and make http request using curl
    
    curl -X POST -H "Content-Type: application/json" -d "{\"actualURL\": \"http://veryLoNguRl/v1/api/LLongGUrrrLl/8988\"}" http://localhost:8081/shorturl

You will be able to see output like this 
    
    {"actualURL":"http://veryLoNguRl/v1/api/LLongGUrrrLl/8988","shortURL":"http://infra.cd/b2hSBS"}


## To make request for getting actual url for a received short url

go to terminal and make http request using curl

    curl -X POST -H "Content-Type: application/json" -d "{\"url\": \"http://infra.cd/gaYyrp\"}" http://localhost:8081/actualurl

You will be able to see output like this 

    {"actualURL":"http://veryLoNguRl/v1/api/LLongGUrrrLl/8988","shortURL":"http://infra.cd/gaYyrp"}

