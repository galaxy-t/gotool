package main

//import (
//	"database/sql"
//	"fmt"
//	"github.com/go-sql-driver/mysql"
//	"log"
//	"time"
//)
//import _ "github.com/go-sql-driver/mysql"
//
//var DB *sql.DB
//
//// MySQL 可连接的数据库种类, 暂时仅支持 MySQL
//const (
//	MySQL = "mysql" // MySQL
//)
//
//// Open 打开一个数据库连接, 需要提供一个合法的数据库源连接信息, 如: username:password@tcp(ip:port)/dbname
//// 以下列出默认设置
//// 最大打开连接数: 10
//// 空闲连接池的最大连接数: 5
//// 一个连接的最长可用时间: 6小时
//// 一个连接可能处于空闲状态的最长时间: 1小时
//func Open(dataSourceName string) {
//
//	var err error
//	DB, err = sql.Open(MySQL, dataSourceName)
//
//	// 数据库连接失败, 实际上此时还未进行连接, 也没有验证驱动程序连接参数, 按照官方说法, 此时只是创建了一个数据库对象,
//	// 只有在第一次需要时, 才会延迟建立与数据库的第一个实际连接
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 设置与数据库的最大打开连接数。
//	// 如果 MaxIdleConns 大于 0 并且新的 MaxOpenConns 小于 MaxIdleConns，那么 MaxIdleConns 将被减少以匹配新的 MaxOpenConns 限制。
//	// 如果 n <= 0，则对打开的连接数没有限制。
//	// 默认为 0（无限制）。
//	DB.SetMaxOpenConns(10)
//
//	// 设置空闲连接池的最大连接数。
//	// 如果 MaxOpenConns 大于 0 但小于新的 MaxIdleConns，则新的 MaxIdleConns 将减少以匹配 MaxOpenConns 限制。
//	// 如果 n <= 0，则不保留任何空闲连接。
//	// 当前的默认最大空闲连接数为 2。这可能会在未来版本中更改。
//	// 官方建议该值与 SetMaxOpenConns 保持一致, 如果该值小于 SetMaxOpenConns , 那么连接打开和关闭的频率可能比预期的要高很多
//	DB.SetMaxIdleConns(10)
//
//	// 设置连接可以重用的最长时间。
//	// 过期的连接可能会在重用前被延迟关闭。
//	// 如果 d <= 0，连接不会因为连接的年龄而关闭。
//	// 官方建议小于五分钟
//	DB.SetConnMaxLifetime(time.Minute * 4)
//
//	// 设置连接可能处于空闲状态的最长时间。
//	// 过期的连接可能会在重用前被延迟关闭。
//	// 如果 d <= 0，连接不会因为连接的空闲时间而关闭。
//	// 如果向更快的关闭空闲连接, 可以使用该值
//	// DB.SetConnMaxIdleTime(time.Hour)
//
//	// 验证数据库是否连接正常
//	err = DB.Ping()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
//
//// Query 执行查询
//func Query() {
//
//	var username string
//	var department_id int64
//
//	// 允许为空的字段, 该包中提供, 但是 go 语言本身不是一种允许可空的语言, 所以之后的数据库设计中要求不允许出现空字段
//	// var s sql.NullString
//
//	rows, err := DB.Query("SELECT * FROM t_aaa WHERE department_id > ?", 100)
//
//	// 将 err 处理成一个包装好的异常
//	if driverErr, ok := err.(*mysql.MySQLError); ok {
//		if driverErr.Number == 1045 {
//			// 此处用于检查是否用户名密码错误
//		}
//	}
//	/*if err != nil {
//		log.Fatal(err)
//	}*/
//
//	// rows 会在 rows.Next() 结束或者 异常终止时自动调用 rows.Close(), 但是我们依然需要主动使用 defer 来标记一下
//	// 因为我们可能会主动退出 rows.Next(), rows.Close() 会返回一个 err, 我们可以在有必要的时候去检查他
//	defer rows.Close()
//
//	for rows.Next() {
//
//		err = rows.Scan(&username, &department_id)
//		if err != nil {
//			log.Fatal(err)
//		}
//		log.Println(username, department_id)
//	}
//
//	// rows.Next() 可能会因为某种原因退出, 而不是正常完成循环, 因此需要检查循环是否正常中止, 异常终止会自动调用 rows.Close()
//	// 检查 rows.Err() 是否存在异常
//	if rows.Err() != nil {
//		log.Fatal(rows.Err())
//	}
//
//}
//
//// QuerySingleRow 只查询一行, 即使返回值有多行, 也只取第一行
//func QuerySingleRow() {
//
//	var username string
//	var department_id int64
//	err := DB.QueryRow("SELECT * FROM t_aaa WHERE department_id > ?", 100).Scan(&username, &department_id)
//	if err != nil {
//
//		if err == sql.ErrNoRows {
//			// 这证明没有获取到任何一条记录
//		} else {
//			// 其它异常
//			log.Fatal(err)
//		}
//	}
//	fmt.Println(username, department_id)
//
//}
//
//func QueryCount() {
//
//	var count int
//
//	rows, _ := DB.Query("SELECT COUNT(*) FROM t_aaa")
//
//	defer rows.Close()
//	rows.Next()
//	ct, _ := rows.ColumnTypes()
//
//	for a, b := range ct {
//		fmt.Println(a, b)
//
//	}
//
//	fmt.Println(ct)
//	rows.Scan(&count)
//	fmt.Println(count)
//
//}
//
//// Exec 用于执行 INSERT,UPDATE,DELETE和其它不返回行的语句(不能查询类似于 SELECT COUNT(*) 之类的语句)
//func Exec() {
//
//	res, err := DB.Exec("INSERT INTO t_aaa(username,department_id) VALUES(?,?)", "kkkkkkk", 123456)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	lastId, err := res.LastInsertId() // 返回自增主键id, 若无自增主键则返回 0, 我们实际上可以使用雪花算法来自主生成主键 id
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	rowCnt, err := res.RowsAffected() // 受影响的行数, 这个值是有意义的
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println(lastId, rowCnt)
//
//}
//
//// TX 事务
//func TX() {
//
//	tx, err := DB.Begin()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	res, err := tx.Exec("INSERT INTO t_aaa(username,department_id) VALUES(?,?)", "kkkkkkk", 123456)
//	if err != nil {
//		tx.Rollback()
//	}
//
//	rowCnt, err := res.RowsAffected()
//	if err != nil {
//		tx.Rollback()
//	}
//
//	fmt.Println(rowCnt)
//
//	tx.Commit()
//
//}
