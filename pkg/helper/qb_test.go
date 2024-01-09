package helper

import (
	"reflect"
	"testing"
)

func TestStructQueryInsert(t *testing.T) {
	type TestData struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}

	// Test case 1 with returningID true
	data := TestData{ID: 1, Name: "John"}
	tableName := "test_table"
	returningID := true

	expectedQuery := `
		INSERT INTO
			test_table
		(
			id,name
		)
		VALUES
		(
			?,?
		)
		RETURNING id;
	`
	expectedValues := []interface{}{1, "John"}

	query, values, err := StructQueryInsert(data, tableName, "db", returningID)

	if err != nil {
		t.Errorf("Expected no error, but got error: %v", err)
	}

	if query != expectedQuery {
		t.Errorf("Query does not match. Expected:\n%s\nActual:\n%s", expectedQuery, query)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Values do not match. Expected: %v, Actual: %v", expectedValues, values)
	}

	// Test case 2 with returningID false
	returningID = false

	expectedQuery = `
		INSERT INTO
			test_table
		(
			id,name
		)
		VALUES
		(
			?,?
		);
	`

	query, values, err = StructQueryInsert(data, tableName, "db", returningID)

	if err != nil {
		t.Errorf("Expected no error, but got error: %v", err)
	}

	if query != expectedQuery {
		t.Errorf("Query does not match. Expected:\n%s\nActual:\n%s", expectedQuery, query)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Values do not match. Expected: %v, Actual: %v", expectedValues, values)
	}
}

func TestStructToQueryUpdate(t *testing.T) {
	type UpdateData struct {
		Name  string `db:"name"`
		Email string `db:"email"`
	}

	type WhereData struct {
		ID int `db:"id"`
	}

	updateInput := UpdateData{Name: "John", Email: "john@example.com"}
	whereInput := WhereData{ID: 1}
	tableName := "users"
	tag := "db"

	expectedQuery := "UPDATE users SET name=?, email=? WHERE id=?"
	expectedValues := []interface{}{"John", "john@example.com", 1}

	query, values, err := StructToQueryUpdate(updateInput, whereInput, tableName, tag)

	if err != nil {
		t.Errorf("Expected no error, but got error: %v", err)
	}

	if query != expectedQuery {
		t.Errorf("Query does not match. Expected: %s, Actual: %s", expectedQuery, query)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Values do not match. Expected: %v, Actual: %v", expectedValues, values)
	}
}
