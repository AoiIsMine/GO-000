package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	name, err := Query(123)
	switch {
	// case err == sql.ErrNoRows:	//warp后,条件为flase
	// fmt.Println("query err = err no rows")
	case errors.Is(err, sql.ErrNoRows):
		err1 := errors.Wrap(err, "111")
		fmt.Printf("no rows,err = %+v", err1)

		fmt.Printf("no rows,err = %+v", err)
		//TODO log
		//TODO logic

	case err != nil:
		// fmt.Println("query err = ", err)
		fmt.Printf("err = %+v", err) //%+v才有详细堆栈输出
		//TODO log
		//TODO logic

	default:
		fmt.Println("query name =", name)
		//TODO logic
	}

}
