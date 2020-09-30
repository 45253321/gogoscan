package plugins

/*
 mysql 爆破引擎
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func MysqlVerify(burstCase BurstCase) bool{
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/",
		burstCase.Username,
		burstCase.Password,
		burstCase.Ip,
		burstCase.Port,
	)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return false
	}
	defer db.Close()

	if err := db.Ping(); err != nil{
		return false
	}
	fmt.Println("connnect success")
	return true
}
