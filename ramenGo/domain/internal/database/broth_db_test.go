package database_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ramenGo/domain/internal/database"
	"github.com/ramenGo/domain/internal/entity"
	"github.com/stretchr/testify/suite"
)

type BrothDBTestSuite struct {
	suite.Suite
	db      *sql.DB
	brothDB *database.BrothDB
	broth   *entity.Broth
}

// Test in memory - Should i change to a test container ?
func (bs *BrothDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	bs.Nil(err)
	bs.db = db

	_, err = db.Exec("CREATE TABLE broths( id varchar(255), imageInactive varchar(510), imageActive varchar(510), name varchar(255), description varchar(510), price decimal(16,14) )")
	fmt.Println(err)
	bs.brothDB = database.NewBrothDB(db)
	bs.broth, _ = entity.NewBroth("someBrothPic", "anotherBrothPic", "Chicken broth", "Caipira Chicken broth good for sic ppl", 12.33)
}

func (bs *BrothDBTestSuite) TearDownSuite() {
	defer bs.db.Close()
	bs.db.Exec("DOP TABLE broths")
}

func TestBrothDBTestSuite(t *testing.T) {
	suite.Run(t, new(BrothDBTestSuite))
}

func (bs *BrothDBTestSuite) TestSave() {
	bs.SetupSuite()
	broth := bs.broth
	err := bs.brothDB.Save(broth)
	bs.Nil(err)
}
