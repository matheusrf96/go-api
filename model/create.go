package model

import "fmt"

func CreateTodo(name string, todo string) error {
	insertQ, err := con.Query("INSERT INTO todo VALUES($1, $2)", name, todo)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer insertQ.Close()
	return nil
}
