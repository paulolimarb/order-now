# Order-now (beta)

Esse projeto tem como objetivo a demonstração de como utilizar o rabbitmq como
um barramento de eventos onde o processo de comunicação com um banco de dados é lento.

### Tecnologia
  - Golang
  - RabbitMq
  - MongoDb
  
### Request

```sh
{
    "order": {
        "order_reference" : "123",
        "product": {
            "product_name" : "Light of led",
            "specifications" : "Color white",
            "price_in_cents" : 1200000
        },
        "buyer": {
            "buyer_name":"Nikola Tesla",
            "document_number" : "12312312311",
            "buyer_address" : "W 40th St, New York, NY 10018, EUA"
        }
    }
}
```

