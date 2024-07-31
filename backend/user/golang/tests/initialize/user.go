package initialize

import (
	"context"
	"database/sql"
	"log"
)

func CreateTableUser(db *sql.DB, ctx context.Context) {
	query := `CREATE TABLE user (
  		id int NOT NULL AUTO_INCREMENT,
  		username varchar(50) NOT NULL,
  		email varchar(100) NOT NULL,
  		password varchar(100) NOT NULL,
  		utc varchar(6) NOT NULL,
  		created_at bigint NOT NULL,
  		PRIMARY KEY (id),
  		UNIQUE KEY username_unique_index (username),
  		UNIQUE KEY email_unique_index (email)
	) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatalln("error when creating table user:", err.Error())
	}
	log.Println("create table user succedded")
}

func CreateDataUser(db *sql.DB, ctx context.Context) {
	query := `INSERT INTO user (id,username,email,password,utc,created_at) VALUES (1,"username","email@email.com","$2a$10$MvEM5qcQFk39jC/3fYzJzOIy7M/xQiGv/PAkkoarCMgsx/rO0UaPG","utc",1695095017);`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatalln("error when creating data user:", err.Error())
	}
	log.Println("create data user succedded")
}

type User struct {
	Id        sql.NullInt32
	Username  sql.NullString
	Email     sql.NullString
	Password  sql.NullString
	Utc       sql.NullString
	CreatedAt sql.NullInt64
}

func GetDataUserByEmail(db *sql.DB, ctx context.Context, email string) (user User) {
	query := `SELECT id, username, email, password, utc, created_at FROM user WHERE email = ?;`
	err := db.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Utc, &user.CreatedAt)
	if err != nil {
		log.Fatalln("error when getting data user:", err.Error())
	}
	log.Println("get data user succedded")
	return
}

func DropTableUser(db *sql.DB, ctx context.Context) {
	query := `DROP TABLE IF EXISTS user;`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatalln("error when dropping table user:", err.Error())
	}
	log.Println("drop table user succedded")
}
