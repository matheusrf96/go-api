package model

import "fmt"

func CreateTodo() error {
	insertQ, err := con.Query("INSERT INTO todo VALUES($1, $2)", "Matheus", "Teste...")

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer insertQ.Close()
	return nil
}
