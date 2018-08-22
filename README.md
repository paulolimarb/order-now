# Order-now (beta)

Esse projeto tem como objetivo demonstrar a utilização do rabbitmq como
um barramento de eventos.

### Tecnologia
  - Golang
  - RabbitMq
  - MongoDb
  - Docker (em breve)
  
### Aplicações
Order-Api - A api é responsavel por receber o request do cliente e
enviar para o rabbitmq.

```sh
cd order-api
go run main.go -dev
```
    
Order-Consumer - O consumer retira as mensagens da fila e persiste no mongodb.

```sh
cd order-consumer
go run main.go -dev
```

### Request

```sh
{
	"order": {
		"order_reference" : "123",
		"product": {
			"product_name" : "Light of led",
			"specifications" : "color white",
			"price_in_cents" : 12000
		},
		"buyer": {
			"buyer_name":"Nikola Tesla",
			"document_number" : "12312312311",
			"buyer_address" : "W 40th St, New York, NY 10018, EUA",
			"buyer_email":"nik@ge.com"
		}
	}
}
```

 ### Observação
 No momento é necessário estar com o rabbitmq e mongodb executando local. 
 


