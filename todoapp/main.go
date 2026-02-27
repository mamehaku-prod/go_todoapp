package main

import (
	"Sec_17-18/todoapp/app/controllers"
	"Sec_17-18/todoapp/app/models"
	"fmt"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	fmt.Println(models.Db)

	controllers.StartMainServer()

	/*
		user, _ := models.GetUserByEmail("test@example.com")
		fmt.Println(user)

		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(session)

		valid, _ := session.CheckSession()
		fmt.Println(valid)
	*/

	/*
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@example.com"
		u.Password = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	/*
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdaterUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo")
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	/*
		user, _ := models.GetUser(3)
		user.CreateTodo("Third Todo")
	*/

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
		fmt.Println(v)
		}
	*/

	/*
		t, _ := models.GetTodo(3)
		t.DeleteTodo()
	*/

}
