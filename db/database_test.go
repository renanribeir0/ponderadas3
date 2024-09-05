package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type DatabaseTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *DatabaseTestSuite) SetupSuite() {
	dsn := "user=testuser password=testpassword dbname=testdb sslmode=disable"
	err := InitDB(dsn)
	suite.Require().NoError(err, "Erro ao conectar ao banco de dados de teste")
	suite.db = DB
}

func (suite *DatabaseTestSuite) TestUserInsertion() {
	user := User{Name: "John Doe"}
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err, "Erro ao criar registro de usuário")

	var retrievedUser User
	err = suite.db.First(&retrievedUser, "name = ?", "John Doe").Error
	suite.Require().NoError(err, "Erro ao recuperar registro de usuário")

	suite.Equal(user.Name, retrievedUser.Name, "Os nomes devem corresponder")
}

func (suite *DatabaseTestSuite) TearDownSuite() {
	suite.db.Exec("DROP TABLE users;")
}

func TestDatabaseTestSuite(t *testing.T) {
	if os.Getenv("POSTGRES_DSN") == "" {
		t.Skip("Testes do PostgreSQL ignorados; variável de ambiente POSTGRES_DSN não fornecida.")
	}

	suite.Run(t, new(DatabaseTestSuite))
}
