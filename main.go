package main;

import (
	"rakamin/routes"
	"fmt"
);

func main() {

	route := routes.SetupRouter();
	
	port := 8080;
	address := fmt.Sprintf("localhost:%d", port);
	fmt.Println("Server running on ", address);
	
	route.Run(address);
}