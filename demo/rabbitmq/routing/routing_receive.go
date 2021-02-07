package routing

import "GoLearningUtils/demo/rabbitmq/RabbitMQ"

func RoutingReceive() {
	routeRec1 := RabbitMQ.NewRabbitMQRouting("exRouteRec1", "routeRec1")
	routeRec1.RecieveRouting()
}
