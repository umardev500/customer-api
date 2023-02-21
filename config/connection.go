package config

import (
	"customer-api/pb"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct{}

func NewConn() *Connection {
	return &Connection{}
}

func (c *Connection) getConn(port string) (conn *grpc.ClientConn) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithNoProxy())
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (c *Connection) ProductConn() (client pb.ProductServiceClient) {
	port := os.Getenv("PRODUCT_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewProductServiceClient(conn)

	return
}

func (c *Connection) OrderConn() (client pb.OrderServiceClient) {
	port := os.Getenv("ORDER_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewOrderServiceClient(conn)

	return
}

func (c *Connection) CustomerConn() (client pb.CustomerServiceClient) {
	port := os.Getenv("CUSTOMER_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewCustomerServiceClient(conn)

	return
}

func (c *Connection) UserConn() (client pb.UserServiceClient) {
	port := os.Getenv("USER_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewUserServiceClient(conn)

	return
}
