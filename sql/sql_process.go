package sql_process

import (
	"database/sql"
	"fmt"
	//"strings"
)
func Ebu_create(db *sql.DB, dbname string){

	dbname = "result_"+ dbname
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default character set utf8", dbname)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Printf("create mysql failed,err:%v\n", err)
		return
	}
	sql = fmt.Sprintf("USE %s", dbname)
	db.Exec(sql)
	var tablename = [4]string{"dir_00", "dir_01", "dir_10", "dir_11"}
	for _, table := range tablename {
		sql = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s"+
			"(mac VARCHAR (10) primary key , before_hopoint_n1 TEXT, after_hopoint_n2 TEXT, not_need_hopoint_n3 TEXT, non_hopoint_n4 TEXT,\
			adjust_hopoint TEXT)" +
			"ENGINE=InnoDB DEFAULT CHARSET=utf8", table)
		db.Exec(sql)
	}
}
func Eau_create(db *sql.DB, dbname string){

	dbname = "result_"+ dbname
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default character set utf8", dbname)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Printf("create mysql failed,err:%v\n", err)
		return
	}
	
	sql = fmt.Sprintf("USE %s", dbname)
	db.Exec(sql)
	var tablename = [4]string{"dir_00", "dir_01", "dir_10", "dir_11"}
	for _, table := range tablename {
		sql = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s"+
			"(timestamp BIGINT, cm_sta_dir TEXT, reconnect_flag TEXT, over_hotime TEXT," +
			"non_hopoint TEXT, side_of_hopoint TEXT, " +
			"hotime TEXT, before_ta TEXT, before_rssi TEXT, after_ta TEXT, after_rssil TEXT, eau_mac TEXT, mac TEXT, destmac TEXT, chn TEXT,"+
			"nb_chn TEXT, nb_tath TEXT, nb_rssith TEXT, new_tath TEXT, new_rssith TEXT,before_recnct_rssi TEXT, before_recnct_ta TEXT, join_flag TEXT, recnct_reason TEXT, opt_method TEXT)" +
			"ENGINE=InnoDB DEFAULT CHARSET=utf8", table)
		db.Exec(sql)
	}
	return
}
func Delete_sql(db *sql.DB, dbname string)

	dbname = "result_"+ dbname
	sql := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbname)
	db.Exec(sql)
	if err != nil {
		fmt.Printf("delete mysql failed,err:%v\n", err)
		return
	}
}