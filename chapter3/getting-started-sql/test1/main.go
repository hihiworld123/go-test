package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

var db *sqlx.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type Employee struct {
	Id         int             `db:"id"`
	Name       string          `db:"name"`
	Department string          `db:"department"`
	Salary     decimal.Decimal `db:"salary"`
}

// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
// 并将结果映射到一个自定义的 Employee 结构体切片中。
func queryByDepartment(department string) ([]Employee, error) {
	sqlStr := "SELECT id,name,department,salary FROM employees t WHERE t.department =?"
	var employees []Employee
	err := db.Select(&employees, sqlStr, department)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil, err
	}
	fmt.Printf("employees:%#v\n", employees)

	return employees, nil
}

func queryMaxSalary() (Employee, error) {
	sqlStr := "SELECT id,name,department,salary FROM employees t1 " +
		"WHERE t1.salary =(SELECT MAX(t.salary) FROM employees t) LIMIT 1"
	employee := Employee{}
	err := db.Get(&employee, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return employee, err
	}
	fmt.Printf("employee:%#v\n", employee)

	return employee, nil
}

func main() {
	employees, err := queryByDepartment("技术部")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(employees)

	salary, err := queryMaxSalary()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(salary)

}
