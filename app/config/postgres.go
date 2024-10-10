package config

import (
	"chat-api/internal/model/entity"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     int    `envconfig:"POSTGRES_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Dbname   string `envconfig:"POSTGRES_DATABASE" required:"true"`

	MaxConnectionLifetime time.Duration `envconfig:"DB_MAX_CONN_LIFE_TIME" default:"300s"`
	MaxOpenConnection     int           `envconfig:"DB_MAX_OPEN_CONNECTION" default:"100"`
	MaxIdleConnection     int           `envconfig:"DB_MAX_IDLE_CONNECTION" default:"10"`
}

func (p Postgres) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

func OpenPostgresDatabaseConnection(pg Postgres) *gorm.DB {

	db, err := gorm.Open(postgres.Open(pg.ConnectionString()))
	if err != nil {
		log.Errorf("Err > %v", err)
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDb.Ping()
	if err != nil {
		log.Errorf("Err sqlDb.Ping> %v", err)
		panic(err)
	}
	//err = db.Migrator().DropTable(&entity.Chats{})
	//if err != nil {
	//	panic(err)
	//}
	//err = db.Migrator().DropTable(&entity.Users{})
	//if err != nil {
	//	panic(err)
	//}
	//err = db.Migrator().DropTable(&entity.Messages{})
	//if err != nil {
	//	panic(err)
	//}
	//err = db.Migrator().DropTable(&entity.Participants{})
	//if err != nil {
	//	panic(err)
	//}
	//err = db.Migrator().DropTable(&entity.Readers{})
	//if err != nil {
	//	panic(err)
	//}
	err = db.AutoMigrate(&entity.Users{}, &entity.Chats{}, entity.Participants{}, entity.Messages{}, entity.References{})
	if err != nil {
		panic(err)
	}
	if db == nil {
		fmt.Println("db nil")
	}
	return db
}
