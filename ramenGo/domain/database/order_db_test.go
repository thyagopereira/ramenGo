package databases_test

import (
	"database/sql"
	"testing"

	databases "github.com/ramenGo/domain/database"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ramenGo/domain/entity"
	"github.com/stretchr/testify/suite"
)

type OrderDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	OrderDB   *databases.OrderDB
	ProteinDB *databases.ProteinDB
	BrothDB   *databases.BrothDB
	Order     *entity.Order
	Protein   *entity.Protein
	Broth     *entity.Broth
}

// Test in memory - Should i change to a test container ?
func (os *OrderDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	os.Nil(err)
	os.db = db

	_, err = db.Exec("CREATE TABLE broths ( id VARCHAR(255) NOT NULL PRIMARY KEY, imageInactive VARCHAR(510), imageActive VARCHAR(510), name VARCHAR(255), description VARCHAR(510), price DECIMAL(16,14))")
	os.Nil(err)
	os.BrothDB = databases.NewBrothDB(db)

	_, err = db.Exec("CREATE TABLE proteins ( id VARCHAR(255) NOT NULL PRIMARY KEY, imageInactive VARCHAR(510), imageActive VARCHAR(510), name VARCHAR(255), description VARCHAR(510), price DECIMAL(16,14))")
	os.Nil(err)
	os.ProteinDB = databases.NewProteinDB(db)

	_, err = db.Exec("CREATE TABLE orders ( id VARCHAR(255) NOT NULL PRIMARY KEY, description VARCHAR(510), image VARCHAR(510), brothId VARCHAR(255), proteinId VARCHAR(255), CONSTRAINT FK_BrothId FOREIGN KEY (brothId) REFERENCES broths(id), CONSTRAINT FK_ProteinId FOREIGN KEY (proteinId) REFERENCES proteins(id))")
	os.Nil(err)
	os.OrderDB = databases.NewOrderDB(db)

	os.Broth, err = entity.NewBroth("someBrothPic", "anotherBrothPic", "Chicken broth", "Caipira Chicken broth good for sic ppl", 12.33)
	os.Nil(err)

	os.Protein, err = entity.NewProtein("someProteinPic", "anotherProteinPic", "Chicken", "Caipira Chicken", 15.204)
	os.Nil(err)

	os.Order, err = entity.NewOrder(os.Broth.Id, os.Protein.Id, "someImage", "someDescription")
	os.Nil(err)
}

func (os *OrderDBTestSuite) TearDownSuite() {
	defer os.db.Close()
	os.db.Exec("DROP TABLE orders")
	os.db.Exec("DROP TABLE proteins")
	os.db.Exec("DROP TABLE broths")
}

func TestOrderDBTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDBTestSuite))
}

func (os *OrderDBTestSuite) TestSave() {
	os.BrothDB.Save(os.Broth)
	os.ProteinDB.Save(os.Protein)

	order, err := entity.NewOrder(os.Broth.Id, os.Protein.Id, "someImage", "someDescription")
	os.Nil(err)

	err = os.OrderDB.Save(order)
	os.Nil(err)
}

func (os *OrderDBTestSuite) TestFindById() {
	os.BrothDB.Save(os.Broth)
	os.ProteinDB.Save(os.Protein)

	err := os.OrderDB.Save(os.Order)
	os.Nil(err)

	id := os.Order.Id
	result, err := os.OrderDB.FindById(id)
	os.Nil(err)
	os.NotNil(result)

	order := result.(*entity.Order)
	os.Equal(id, order.Id)
	os.EqualValues(os.Order, order)
}

func (os *OrderDBTestSuite) TestGetAll() {
	// Saving some orders

	os.BrothDB.Save(os.Broth)
	os.ProteinDB.Save(os.Protein)
	os.OrderDB.Save(os.Order)

	anotherOrder, _ := entity.NewOrder(os.Broth.Id, os.Protein.Id, "someImage", "someDescription")
	os.OrderDB.Save(anotherOrder)

	expected := []entity.Entity{os.Order, anotherOrder}
	result, err := os.OrderDB.GetAll()

	os.Nil(err)
	os.NotNil(result)
	os.EqualValues(expected, result)
}
