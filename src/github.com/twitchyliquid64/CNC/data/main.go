package data

import (
  "github.com/twitchyliquid64/CNC/data/stmdata"
  "github.com/twitchyliquid64/CNC/data/session"
  "github.com/twitchyliquid64/CNC/data/entity"
  "github.com/twitchyliquid64/CNC/data/plugin"
  "github.com/twitchyliquid64/CNC/data/user"
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "database/sql"
)

var DB gorm.DB

func Initialise() {
  logging.Info("data", "Initialise()")
  trackingSetup()

  dbConn, err := sql.Open("postgres", "postgres://" + config.All().Database.Username +
                                     ":" + config.All().Database.Password +
                                     "@" + config.All().Database.Address +
                                     "/" + config.All().Database.Name +
                                     "?sslmode=require")
	if err != nil {
		logging.Error("data", "Error opening DB connection")
    logging.Error("data", "Error: ", err)
    tracking_notifyFault(err)
	}

  DB, err = gorm.Open("postgres", dbConn)
  DB.LogMode(true)

  if err != nil {
    logging.Error("data", "Error launching DB engine")
    logging.Error("data", "Error: ", err)
  }

  checkStructures()

  //make sure that objects in the config BaseObjects are
  //existing, creating them if nessesary.
  for _, usr := range config.All().BaseObjects.AdminUsers{

    tmp := user.User{}

    DB.Where(&user.User{Username:  usr.Username}).First(&tmp)

    if tmp.Username != usr.Username{ //if the user was not found
      logging.Info("data", "Creating admin user: " + usr.Username)
      DB.Create(&user.User{Username: usr.Username,
                        Permissions: []user.Permission{ user.Permission{Name: user.PERM_ADMIN},},
                        AuthMethods: []user.AuthenticationMethod{ user.AuthenticationMethod{
                            MethodType: user.AUTH_PASSWD,
                            Value: usr.Password,
                          }},
                      })
    }
  }

  logging.Info("data", "Initialisation finished.")
}

func autoMigrateTables() {
  logging.Info("data", "Checking structure: Users")
  DB.AutoMigrate(&user.User{})
  logging.Info("data", "Checking structure: Permissions")
  DB.AutoMigrate(&user.Permission{})
  user.Permission{}.Init(DB)
  logging.Info("data", "Checking structure: Emails")
  DB.AutoMigrate(&user.Email{})
  logging.Info("data", "Checking structure: Addresses")
  DB.AutoMigrate(&user.Address{})
  logging.Info("data", "Checking structure: AuthenticationMethods")
  DB.AutoMigrate(&user.AuthenticationMethod{})
  logging.Info("data", "Checking structure: Sessions")
  DB.AutoMigrate(&session.Session{})

  logging.Info("data", "Checking structure: Entity")
  DB.AutoMigrate(&entity.Entity{})
  logging.Info("data", "Checking structure: EntityPivot")
  DB.AutoMigrate(&entity.EntityPivot{})
  logging.Info("data", "Checking structure: EntityLocationRecord")
  DB.AutoMigrate(&entity.EntityLocationRecord{})
  logging.Info("data", "Checking structure: EntityStatusRecord")
  DB.AutoMigrate(&entity.EntityStatusRecord{})
  logging.Info("data", "Checking structure: EntityLogRecord")
  DB.AutoMigrate(&entity.EntityLogRecord{})
  logging.Info("data", "Checking structure: EntityEvent")
  DB.AutoMigrate(&entity.EntityEvent{})

  logging.Info("data", "Checking structure: Plugin")
  DB.AutoMigrate(&plugin.Plugin{})
  logging.Info("data", "Checking structure: Resource")
  DB.AutoMigrate(&plugin.Resource{})
  logging.Info("data", "Checking structure: Stmdata")
  DB.AutoMigrate(&stmdata.Stmdata{})
}

//called during initialisation. Should make sure the schema is intact and up to date.
func checkStructures() {
  logging.Info("Auto migrating tables");
  autoMigrateTables();
  

}
