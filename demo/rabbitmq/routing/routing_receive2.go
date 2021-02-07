package routing

import "SpikeSystem/demo/rabbitmq/RabbitMQ"

func RoutingReceive2() {
	routeRec2 := RabbitMQ.NewRabbitMQRouting("exRouteRec2", "routeRec1")
	routeRec2.RecieveRouting()
}
