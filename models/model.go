package models

import(
	"github.com/jinzhu/gorm" //go get -u github.com/jinzhu/gorm
   	_ "github.com/jinzhu/gorm/dialects/mysql" //go get github.com/go-sql-driver/mysql
   	"log"
)

var connection *gorm.DB

var engine_sql string = "mysql"
var username string = "root"
var password string = ""
var database string = "go_web_project" 

func InitializeDataBase(){
	connection = ConnectORM( GetConnectionString() )
}

func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open(engine_sql,stringConnection)
  	if err != nil {
  		log.Fatal(err)
  		return nil
  	}
  	return connection
}

func CloseConnection(){
  connection.Close()
}

func GetConnectionString() string{
  return username + ":" + password + "@/" + database
}

func DeleteRecords(){
	connection.Where("id > ?", "0").Delete(User{})
}




