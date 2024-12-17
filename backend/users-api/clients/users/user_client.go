package clientUsers

import (
	e "backend/errors"
	Model "backend/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	//"gorm.io/gorm"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type MySQLConfig struct {
	Name string
	User string
	Pass string
	Host string
}

type SQL struct {
	db       *gorm.DB
	Database string
}

// DeleteReserva implements reservas.Repository.

func NewSql(config MySQLConfig) SQL {
	db, err := gorm.Open("mysql", config.User+":"+config.Pass+"@tcp("+config.Host+":3306)/"+config.Name+"?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Println("Connection Established gg")
	}
	db.AutoMigrate(&Model.User{})
	return SQL{
		db:       db,
		Database: config.Name,
	}

}

/*
var (

	migrate = []interface{}{
		Model.User{},
	}

)

	func NewMySQL(config MySQLConfig) SQL {
		// Build DSN (Data Source Name)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Username, config.Password, config.Host, config.Port, config.Database)

		// Open connection to MySQL using GORM
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to MySQL: %s", err.Error())
		}

		// Automigrate structs to Gorm
		for _, target := range migrate {
			if err := db.AutoMigrate(target); err != nil {
				log.Fatalf("error automigrating structs: %s", err.Error())
			}
		}

		return SQL{
			db: db,
		}
	}
*/
func (repository SQL) GetUserByName(Usuario Model.User) (Model.User, e.ApiError) {
	var user Model.User

	result := repository.db.Where("User = ?", Usuario.User).First(&user)
	log.Debug("User: ", user)
	if result.Error != nil {
		log.Error("Error al buscar el usuario")
		log.Error(result.Error)
		return user, e.NewBadRequestApiError("Error al buscar el usuario")
	}
	println("se busci user en main")
	return user, nil
}

func (repository SQL) InsertUsuario(user Model.User) (Model.User, e.ApiError) {
	result := repository.db.Create(&user)

	if result.Error != nil {
		log.Error("Error al crear el usuario")
		log.Error(result.Error)
		return user, e.NewBadRequestApiError("Error al crear el usuario")
	}
	log.Debug("User Created: ", user.Id)
	return user, nil
}

func (repository SQL) GetUserById(Id int) (Model.User, e.ApiError) {
	var userId Model.User

	result := repository.db.Where("id = ?", Id).First(&userId)
	log.Debug("id: ", userId)
	if result.Error != nil {
		log.Error("Error al buscar el usuario")
		log.Error(result.Error)
		return userId, e.NewBadRequestApiError("Error al buscar el usuario")
	}

	return userId, nil
}

func (repository SQL) GetuserName(buscado int) (string, e.ApiError) {
	var userId Model.User

	result := repository.db.Where("id = ?", buscado).First(&userId)
	if result.Error != nil {
		log.Error("Error al buscar el usuario")
		log.Error(result.Error)
		return userId.User, e.NewBadRequestApiError("Error al buscar el usuario")
	}

	return userId.User, nil
}
