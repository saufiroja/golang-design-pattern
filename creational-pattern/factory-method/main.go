package main

import "fmt"

func main() {
	mysql := NewDatabaseConnection("MySQL")
	println(mysql.Connect())

	postgres := NewDatabaseConnection("PostgreSQL")
	println(postgres.Connect())
}

// product
type DatabaseConnection interface {
	Connect() string
}

// concrete product
type MySQLConnection struct {
	Dsn string
}

func (m *MySQLConnection) Connect() string {
	return fmt.Sprintf("Connecting to %s", m.Dsn)
}

// concrete product
type PostgreSQLConnection struct {
	Dsn string
}

func (p *PostgreSQLConnection) Connect() string {
	return fmt.Sprintf("Connecting to %s", p.Dsn)
}

// creator
type DatabaseConnectionFactory interface {
	CreateConnection() DatabaseConnection
}

// concrete creator
type MySQLConnectionFactory struct{}

func NewMySQLConnectionFactory() *MySQLConnectionFactory {
	return &MySQLConnectionFactory{}
}

func (m *MySQLConnectionFactory) CreateConnection() DatabaseConnection {
	return &MySQLConnection{
		Dsn: "mysql://root:root@localhost:3306/db",
	}
}

// concrete creator
type PostgreSQLConnectionFactory struct{}

func NewPostgreSQLConnectionFactory() *PostgreSQLConnectionFactory {
	return &PostgreSQLConnectionFactory{}
}

func (p *PostgreSQLConnectionFactory) CreateConnection() DatabaseConnection {
	return &PostgreSQLConnection{
		Dsn: "postgres://root:root@localhost:5432/db",
	}
}

// factory method
func NewDatabaseConnection(dbType string) DatabaseConnection {
	switch dbType {
	case "MySQL":
		return NewMySQLConnectionFactory().CreateConnection()
	case "PostgreSQL":
		return NewPostgreSQLConnectionFactory().CreateConnection()
	default:
		return nil
	}
}
