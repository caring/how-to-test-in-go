package mocks_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matryer/is"

	logging "github.com/caring/gopkg-logging"
	mocks "github.com/caring/test/1-mocks"
)

func TestItem_GetN(t *testing.T) {
	log := logging.NewNopLogger()

	var expectCols = []string{
		"item_id", "item_name",
	}
	var expectValues = func(item_id string) []driver.Value {
		return []driver.Value{item_id, ""}
	}

	tests := []struct {
		args      []string
		expectSQL string
	}{
		{
			args: nil,
			expectSQL: `
			SELECT item_id,
			       item_name
			FROM items
			`,
		},
		{
			args: []string{"one"},
			expectSQL: `
			SELECT item_id,
			       item_name
			FROM items
			WHERE item_id IN \(\?\)
			`,
		},
		{
			args: []string{"one", "two", "three"},
			expectSQL: `
			SELECT item_id,
			       item_name
			FROM items
			WHERE item_id IN \(\?,\?,\?\)
			`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint("Test_", i), func(t *testing.T) {
			is := is.New(t)
			db, mock, err := sqlmock.New()
			is.NoErr(err)

			qry := mock.ExpectQuery(tt.expectSQL)

			rows := sqlmock.NewRows(expectCols)
			for _, v := range tt.args {
				rows.AddRow(expectValues(v)...)
			}
			qry.WillReturnRows(rows)

			lis, err := mocks.ItemSQL{}.GetN(context.TODO(), log, db, tt.args...)
			is.NoErr(err)

			is.Equal(len(tt.args), len(lis))

			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
