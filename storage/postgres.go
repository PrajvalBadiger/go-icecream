package storage

import (
	"database/sql"
	"fmt"

	"github.com/PrajvalBadiger/go-icecream/types"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

// Create a new Postgres DB
func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Check the connection to the database
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

// Init: Initalize a DB
func (s *PostgresStore) Init() error {
	return s.Create_flavour_table()
}

// CreateAccountTable: create a table for the db
func (s *PostgresStore) Create_flavour_table() error {
	// only create a table if it doesn't exists
	query := `create table if not exists flavours (
		id serial primary key,
		name varchar(50),
		price serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) Create_flavour(f *types.Flavour) error {
	query := `insert into flavours
	(name, price, created_at)
	values ($1, $2, $3)`

	_, err := s.db.Query(query, f.Name, f.Price, f.Created_at)
	if err != nil {
		return err
	}

	return nil
}

// Get_flavours: get all the flavours in the Storage
func (s *PostgresStore) Get_flavours() ([]*types.Flavour, error) {
	rows, err := s.db.Query("select * from flavours")
	if err != nil {
		return nil, err
	}

	flavours := []*types.Flavour{}

	for rows.Next() {
		flavour, err := scan_into_row(rows)
		if err != nil {
			return nil, err
		}
		flavours = append(flavours, flavour)
	}
	return flavours, nil
}

// Get_flavour_by_id: get the flavour by its id
func (s *PostgresStore) Get_flavour_by_id(id int) (*types.Flavour, error) {
	rows, err := s.db.Query("select * from flavours where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scan_into_row(rows)
	}

	return nil, fmt.Errorf("Flavour %d not found", id)
}

// Delete_flavour: remove a flavour form the Storage
func (s *PostgresStore) Delete_flavour(id int) error {

	_, err := s.db.Query("delete from flavours where id = $1", id)
	if err != nil {
		return err
	}

	return err
}

// Update_flavour: update a flavour which already exists in the Storage
func (s *PostgresStore) Update_flavour(id int, f *types.Flavour) error {
	query := `update flavours
	set name = $1, price = $2, created_at = $3
	where id = $4`

	_, err := s.db.Query(query, f.Name, f.Price, f.Created_at, id)
	if err != nil {
		return err
	}

	return nil
}

// utils

// scan the row to get all the columns of the row
func scan_into_row(rows *sql.Rows) (*types.Flavour, error) {
	flavour := new(types.Flavour)
	err := rows.Scan(
		&flavour.ID,
		&flavour.Name,
		&flavour.Price,
		&flavour.Created_at)

	return flavour, err
}
