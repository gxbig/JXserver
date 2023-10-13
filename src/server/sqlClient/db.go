package sqlClient

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/name5566/leaf/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}

var Db *sql.DB

func init() {
	log.Debug("init db")
	Db = OpenDb()
}
func OpenDb() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Error("数据库连接失败")
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
