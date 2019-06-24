// Author: James Mallon <jamesmallondev@gmail.com>
// Main package
package main

import (
	//"axeman/libs/databases"
	//"axeman/models"
	//"axeman/models/dao"
	//"fmt"
	"axeman/routes"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// starting point...
func main() {
	//conn, _ := databases.SQLConn()
	//conn.Migrate(&models.User{})
	//user := models.User{"test@test.com", "12345", "- peace bru! - peace, shit... - fuck that shit off!!!! WAR!!!!!!!!!!!!!"}
	//userDAO := dao.UserDAO()
	//for i := 0; i < 10; i++ {
	//	//res := userDAO.InsertUser(conn, &user)
	//	res := userDAO.InsertUser(&user)
	//	if res {
	//		fmt.Println("User created!")
	//	}
	//}

	//conn.DB.Create(&user)
	//conn.DB.Create(&user)
	//conn.DB.Create(&user)
	//conn.DB.Create(&user)

	router := httprouter.New()

	routes.GetRoutes(router)
	routes.GetAuthRoutes(router)

	router.ServeFiles("/static/*filepath", http.Dir("public/"))
	router.ServeFiles("/libs/*filepath", http.Dir("node_modules/"))
	panic(http.ListenAndServe("localhost:8080", router))

}
