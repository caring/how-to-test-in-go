package mocks

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	errors "github.com/caring/gopkg-errors"
	logging "github.com/caring/gopkg-logging"

	"github.com/caring/test/1-mocks/db"
)

// Item domain
type Item struct {
	ItemID   string
	ItemName string
}

// FromRow scans a single row into the receiver call.
func (c *Item) FromRow(row *sql.Rows) error {
	if c == nil {
		return errors.WithStack(fmt.Errorf("receiver is nil pointer"))
	}

	err := row.Scan(
		&c.ItemID,
		&c.ItemName,
	)
	return errors.Wrap(err, "error fetching calls")
}

// ListItem is an aggregate of Items
type ListItem []*Item

// SQL defines the get operations for Item
type ItemSQL struct{}

// FromRows scans all rows into the receiver Item list.
// Will fail if called on a nil pointer to list.
// GOOD: var lis ListItem
// BAD:  var lis *ListItem
// OK:   var lis *ListItem = &ListItem{}
func (lis *ListItem) FromRows(rows *sql.Rows) (n int64, err error) {
	if lis == nil {
		return 0, errors.WithStack(fmt.Errorf("receiver is nil pointer"))
	}

	for rows.Next() {
		// Query and store result in model.
		c := Item{}
		if err = c.FromRow(rows); err != nil {
			return 0, err
		}
		*lis = append(*lis, &c)
		n++
	}
	return n, nil
}

// SetOrder aranges Conferences according to the ids provided.
func (lis *ListItem) SetOrder(ids ...string) {
	nlis := make(ListItem, len(ids))

	// if receiver is nil or empty, size to ids and return
	if lis == nil || len(*lis) == 0 {
		*lis = nlis
		return
	}

	// build reverse index
	rev := make(map[string]int, len(ids))
	for i, k := range ids {
		rev[k] = i
	}

	// apply index to new list
	for _, c := range *lis {
		if i, ok := rev[c.ItemID]; ok {
			nlis[i] = c
		}
	}

	// replace list with new order
	*lis = nlis
}

// GetAll returns a list of all Items.
func (s ItemSQL) GetAll(ctx context.Context, log *logging.Logger, qx db.QueryExecutor) (lis ListItem, err error) {
	query := `
	SELECT
	    item_id,
	    item_name
	FROM items
	`

	rows, err := db.QueryContext(ctx, log, qx, query)
	if err != nil {
		return nil, err
	}
	_, err = lis.FromRows(rows)

	return lis, err
}

// GetN returns a slice of Items ordered by provided ids.
// If a row for an id is not found it will set to nil that index. For use with dataloaders.
// If no id's are supplied it will return all results. For use as GetAll
func (s ItemSQL) GetN(ctx context.Context, log *logging.Logger, qx db.QueryExecutor, listIDs ...string) (lis ListItem, err error) {
	if len(listIDs) == 0 {
		return s.GetAll(ctx, log, qx)
	}

	lis = make(ListItem, 0, len(listIDs))
	ids := make([]interface{}, len(listIDs))
	for i, id := range listIDs {
		ids[i] = id
	}

	// Prepare query statement
	query := `
	SELECT
	    item_id,
	    item_name
	FROM items
	` + qryWhereN("item_id", len(listIDs))

	// Query to rows
	log.Debug(fmt.Sprintf("Running query to get: %s", query))
	rows, err := db.QueryContext(ctx, log, qx, query, ids...)
	if err != nil {
		return lis, err
	}
	n, err := lis.FromRows(rows)
	if n == 0 {
		return lis, db.ErrNoRows
	}

	// If more than one id, set order by ids.
	if len(listIDs) > 1 {
		lis.SetOrder(listIDs...)
	}

	return lis, err
}

// qryWhereN is a helper to buld a list of parameter placeholders.
func qryWhereN(col string, n int) string {
	return "WHERE " + col + " IN (?" + strings.Repeat(",?", n-1) + ")\n"
}
