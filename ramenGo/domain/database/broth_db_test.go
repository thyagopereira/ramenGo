package databases_test

import (
	"database/sql"
	"testing"

	databases "github.com/ramenGo/domain/database"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ramenGo/domain/entity"
	"github.com/stretchr/testify/suite"
)

type BrothDBTestSuite struct {
	suite.Suite
	db      *sql.DB
	brothDB *databases.BrothDB
	broth   *entity.Broth
}

// Test in memory - Should i change to a test container ?
func (bs *BrothDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	bs.Nil(err)
	bs.db = db

	_, err = db.Exec("CREATE TABLE broths( id varchar(255), imageInactive varchar(510), imageActive varchar(510), name varchar(255), description varchar(510), price decimal(16,14) )")
	bs.Nil(err)
	bs.brothDB = databases.NewBrothDB(db)
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
	broth := bs.broth
	err := bs.brothDB.Save(broth)
	bs.Nil(err)
}

func (bs *BrothDBTestSuite) TestFindById() {

	bs.brothDB.Save(bs.broth)
	id := bs.broth.Id

	result, err := bs.brothDB.FindById(id)

	bs.Nil(err)
	bs.NotNil(result)
	broth := result.(*entity.Broth)
	bs.Equal(id, broth.Id)
	bs.EqualValues(bs.broth, broth)
}

func (bs *BrothDBTestSuite) TestGetAll() {
	// Saving some broths
	anotherBroth, _ := entity.NewBroth("someBroth2", "anotherBrothPic2", "Vegetable broth", "Full vegetable broth good for sic ppl", 10.00)
	bs.brothDB.Save(anotherBroth)

	expected := []entity.Entity{bs.broth, anotherBroth}
	result, err := bs.brothDB.GetAll()

	bs.Nil(err)
	bs.NotNil(result)
	bs.EqualValues(expected, result)
}
