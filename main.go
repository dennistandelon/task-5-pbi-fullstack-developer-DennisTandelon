package main;

import (
	"github.com/dennistandelon/task-5-pbi-fullstack-developer-DennisTandelon/routes"
	"fmt"
);

func main() {

	route := routes.SetupRouter();
	
	port := 8080;
	address := fmt.Sprintf("localhost:%d", port);
	fmt.Println("Server running on ", address);
	
	route.Run(address);
}