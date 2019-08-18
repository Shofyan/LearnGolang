//https://dasarpemrogramangolang.novalagung.com/53-sql.html

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/LearnGolang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("SELECT id,name,grade from tb_student where age=?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result []student

	for rows.Next() {
		var item = student{}
		var err = rows.Scan(&item.id, &item.name, &item.grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, item := range result {
		fmt.Println(item.name)
	}

	if err := rows.Close(); err != nil {
		fmt.Println(err.Error())
		return
	}

}

func sqlQueryRow() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var result student
	id := "E001"
	err = db.QueryRow("SELECT name,grade from tb_student where id=?", id).Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name:%s, grade:%d ", result.name, result.grade)

}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name, grade from tb_student where id=?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	_ = stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s, grade: %d\n", result1.name, result1.grade)

	var result2 = student{}
	_ = stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s, grade: %d\n", result2.name, result2.grade)

	var result3 = student{}
	_ = stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("name: %s, grade: %d\n", result3.name, result3.grade)
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values(?,?,?,?) ", "g001", "kina", 23, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success")

	_, err = db.Exec("update  tb_student set age=? where id=?", 25, "g001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success")

	_, err = db.Exec("delete from  tb_student where id=?", "g001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success")
}

func main() {
	sqlQuery()
	sqlQueryRow()
	sqlPrepare()
	sqlExec()
}
