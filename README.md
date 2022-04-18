# dummy_service

This is a dummy service with two endpoints that integrate with two individual systems: ABC and DEF. When a user sends a request, this service processes it, prepares and sends its own request to one of the systems. As soon as a response is received this service combines data, saves it to DB(Postgre) and responses to a client with a simple counter and a statusOk.

## methods
    1. http://localhost:1111/api/register
    2. http://localhost:1111/api/schedule (upcoming)

## updates
    1. initial upload(15apr22)
    2. working proto (15apr22)

## swagger
    1. http://localhost:1111/swagger/index.html