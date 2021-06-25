// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const rightJoin = `-- name: RightJoin :many
SELECT f.id, f.bar_id, b.id
FROM foo f
RIGHT JOIN bar b ON b.id = f.bar_id
WHERE f.id = $1
`

type RightJoinRow struct {
	ID    sql.NullInt32
	BarID sql.NullInt32
	ID_2  int32
}

func (q *Queries) RightJoin(ctx context.Context, id int32) ([]RightJoinRow, error) {
	rows, err := q.db.QueryContext(ctx, rightJoin, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RightJoinRow
	for rows.Next() {
		var i RightJoinRow
		if err := rows.Scan(&i.ID, &i.BarID, &i.ID_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}