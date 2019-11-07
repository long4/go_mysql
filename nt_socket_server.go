package main

import (
	"database/sql"
	"net"
	"fmt"
	"log"
	//"github.com/axgle/mahonia"
	"strings"
	_"github.com/go-sql-driver/mysql"
	"sql"
)
func get_ip(host *string) (flag bool){

	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return false
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						*host = ipnet.IP.String()
						fmt.Println(ipnet.IP.String())
						return true
					}
				}
			}
		}
	}
	return false
}
func main () {
	var host string
	var sqlport string = "3306"
	var buf = make([]byte, 1024)
	port := "9000"
	get_ip(&host)
	addr:= host + ":" + port
	
	listener, err := net.Listen("tcp", addr)
 	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err == nil {
			_, err := conn.Read(buf)		
			if err == nil {
				//desDecoder := mahonia.NewDecoder("utf-8")
				//_, resBytes, _ := desDecoder.Translate([]byte(buf), true)
				//desDecoder.Translate([]byte(buf), true)
				var str string = string(buf)
				strSplit := strings.Split(str, ",")
				fmt.Println(strSplit)
				for _, str1 := range strSplit {
					str1Split := strings.Split(str1, "=")
					var a [2]string
					copy(a[:], str1Split)
					fmt.Println(a[0])
					fmt.Println(a[1])
					if "sample_dev1_cm_value" == a[0]{
						fmt.Println(a[1])
						//dsn := fmt.Sprintf("%s:%s@%s(%s:%d/)", "nt", "nufront", "tcp", host, sqlport)
						DB, err := sql.Open("mysql", "nt:nufront@tcp(" + host + ":" + sqlport +")/")
						if err != nil {
							fmt.Printf("Open mysql failed,err:%v\n", err)
							return
						}
						sql_process.Eau_create(DB, a[1])
						DB.Close()
					}
					
				}
			}
		}
	}
}
