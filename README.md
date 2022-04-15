# dummy_service

This is a dummy service that is has two endpoints which simulate integration with two individual systems: ABC and DEF. When request is received this service processes it, prepares its own request to one of the systems and sends it. When response is received service combines data and saves it to DB(Postgre) and responses to a client with a simple counter.


## methods
    1. http://localhost:1111/api/register
    2. http://localhost:1111/api/schedule (upcoming)
## updates
    1. initial upload(15apr22)
    2. working proto (15apr22)