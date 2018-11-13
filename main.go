//go:generate goagen bootstrap -d github.com/samclick/api/design
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kasoshojo/api/app"
)

func main() {
	// Create service
	service := goa.New("kasoshojo")

	dbuser := os.Getenv("DB_USER")
	dbpwd := os.Getenv("DB_PWD")
	dbhost := os.Getenv("DB_HOST")
	dbschema := os.Getenv("DB_SCHEMA")
	log.Println("Connecting to db")
	db := openDBConnection(dbuser, dbpwd, dbhost, dbschema)
	//db.LogMode(true)
	if db.Error != nil {
		panic("failed to connect to DB")
	}
	defer db.Close()

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	keys, _ := loadJWTPublicKeys()
	jwtMiddleware := jwt.New(jwt.NewSimpleResolver(keys), nil, app.NewJWTSecurity())
	app.UseJWTMiddleware(service, jwtMiddleware)

	// Mount "affiliate" controller
	c := NewUsersController(service, db)
	app.MountUsersController(service, c)

	c2 := NewMessagesController(service, db)
	app.MountMessagesController(service, c2)

	c3 := NewNewsController(service, db)
	app.MountNewsController(service, c3)

	c4 := NewHealthController(service)
	app.MountHealthController(service, c4)

	c5 := NewSurveysController(service, db)
	app.MountSurveysController(service, c5)

	c6 := NewAuthController(service, db)
	app.MountAuthController(service, c6)

	c7 := NewSwaggerController(service)
	app.MountSwaggerController(service, c7)

	// Start service

	sslpath := os.Getenv("SSL_PATH")
	if err := service.ListenAndServeTLS(":8080", sslpath+"fullchain.pem", sslpath+"privkey.pem"); err != nil {
		service.LogError("startup", "err", err)
	}

	/*if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}*/
}

func loadJWTPublicKeys() ([]jwt.Key, error) {
	keyFiles, err := filepath.Glob("./key/*.pub")
	if err != nil {
		return nil, err
	}
	keys := make([]jwt.Key, len(keyFiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		key, err := jwtgo.ParseRSAPublicKeyFromPEM(pem)
		if err != nil {
			return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load public keys for JWT security")
	}

	return keys, nil
}

//OpenDBConnection opens a mysql connection over the default port to the database at the given location with the given credentials.
func openDBConnection(dbuser string, dbpwd string, dbhost string, dbschema string) *gorm.DB {
	/*gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "dtb_" + defaultTableName;
	}*/

	connstring := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbuser, dbpwd, dbhost, dbschema)
	log.Println(connstring)
	db, err := gorm.Open("mysql", connstring)

	if err != nil {
		panic(err.Error())
	}
	db.LogMode(false)
	//db.SingularTable(true)
	return db
}

func exitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
	os.Exit(1)
}
