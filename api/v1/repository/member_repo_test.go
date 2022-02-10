package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/junminhong/member-center-service/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestGetEmail(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//mock.ExpectPrepare("SELECT TIMEDIFF")
	//mock.ExpectPrepare("SELECT ENGINE")
	// ExpectExec，期望执行一条Exec语句
	// 然后假定会返回(1, 1)，也就是自增主键为1，1条影响结果
	//mock.ExpectExec("INSERT").
	//		WillReturnResult(sqlmock.NewResult(1, 1))
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic(err) // Error here
	}
	rows := sqlmock.NewRows([]string{"uuid", "email", "password"}).
		AddRow("c56df1c0-675e-47ce-add0-123aecba4473", "test@gmail.com", "test")
	//query := "SELECT uuid, email, password FROM members WHERE uuid=?"
	//tmp := mock.ExpectQuery(query).WithArgs("c56df1c0-675e-47ce-add0-123aecba4473").WillReturnRows(rows)
	log.Println(rows)
	testMember := &domain.Member{}
	gormDB.Where("uuid=?", "c56df1c0-675e-47ce-add0-123aecba4473").First(&testMember)
	//postgresql := repository.NewPostgresqlRepository(gormDB, logger.Setup())
	//d := postgresql.GetEmail(&gin.Context{}, "c56df1c0-675e-47ce-add0-123aecba4473")
	//log.Println(d)
	//log.Println(testMember.Email)
}
