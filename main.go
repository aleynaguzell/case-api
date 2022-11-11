package main

import "case-api/api"

//  @title           Swagger Case API
// @version         1.0
// @description     An application with Swagger
// @termsOfService  http://swagger.io/terms/

// @contact.name   #aleyna
// @contact.email  aleynaguzel2109@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /
func main() {
	api.Init()
	api.App.Start()
}
