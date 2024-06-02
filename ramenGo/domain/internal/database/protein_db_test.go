package database_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ramenGo/domain/internal/database"
	"github.com/ramenGo/domain/internal/entity"
	"github.com/stretchr/testify/suite"
)

type ProteinDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	proteinDB *database.ProteinDB
	protein   *entity.Protein
}

// Test in memory - Should i change to a test container ?
func (ps *ProteinDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	ps.Nil(err)
	ps.db = db

	_, err = db.Exec("CREATE TABLE proteins( id varchar(255), imageInactive varchar(510), imageActive varchar(510), name varchar(255), description varchar(510), price decimal(16,14) )")
	ps.Nil(err)
	ps.proteinDB = database.NewProteinDB(db)
	ps.protein, _ = entity.NewProtein("someProteinPic", "anotherProteinPic", "Chicken", "Caipira Chicken", 15.204)
}

func (ps *ProteinDBTestSuite) TearDownSuite() {
	defer ps.db.Close()
	ps.db.Exec("DOP TABLE broths")
}

func TestProteinDBTestSuite(t *testing.T) {
	suite.Run(t, new(ProteinDBTestSuite))
}

func (ps *ProteinDBTestSuite) TestSave() {
	protein := ps.protein
	err := ps.proteinDB.Save(protein)
	ps.Nil(err)
}

func (ps *ProteinDBTestSuite) TestFindById() {

	ps.proteinDB.Save(ps.protein)
	id := ps.protein.Id

	result, err := ps.proteinDB.FindById(id)

	ps.Nil(err)
	ps.NotNil(result)
	protein := result.(*entity.Protein)
	ps.Equal(id, protein.Id)
	ps.EqualValues(ps.protein, protein)
}

func (ps *ProteinDBTestSuite) TestGetAll() {
	// Saving some broths
	anotherProtein, _ := entity.NewProtein("veganProteinpic", "anotherProteinPic2", "Vegetable Protein", "Full vegetable Protein", 10.02)
	ps.proteinDB.Save(anotherProtein)

	expected := []entity.Entity{ps.protein, anotherProtein}
	result, err := ps.proteinDB.GetAll()

	ps.Nil(err)
	ps.NotNil(result)
	ps.EqualValues(expected, result)
}
