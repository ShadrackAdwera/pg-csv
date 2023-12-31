// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: sales.sql

package internal

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSale = `-- name: CreateSale :one
INSERT INTO sales (region, country,item_type,sales_channel, order_priority, order_date, order_id, ship_date,units_sold, unit_price,unit_cost,total_revenue,total_cost,total_profit)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) 
RETURNING id, region, country, item_type, sales_channel, order_priority, order_date, order_id, ship_date, units_sold, unit_price, unit_cost, total_revenue, total_cost, total_profit, created_at
`

type CreateSaleParams struct {
	Region        string         `json:"region"`
	Country       string         `json:"country"`
	ItemType      string         `json:"item_type"`
	SalesChannel  ESalesChannel  `json:"sales_channel"`
	OrderPriority EOrderPriority `json:"order_priority"`
	OrderDate     time.Time      `json:"order_date"`
	OrderID       int64          `json:"order_id"`
	ShipDate      time.Time      `json:"ship_date"`
	UnitsSold     int32          `json:"units_sold"`
	UnitPrice     float32        `json:"unit_price"`
	UnitCost      float32        `json:"unit_cost"`
	TotalRevenue  pgtype.Numeric `json:"total_revenue"`
	TotalCost     pgtype.Numeric `json:"total_cost"`
	TotalProfit   pgtype.Numeric `json:"total_profit"`
}

func (q *Queries) CreateSale(ctx context.Context, arg CreateSaleParams) (Sale, error) {
	row := q.db.QueryRow(ctx, createSale,
		arg.Region,
		arg.Country,
		arg.ItemType,
		arg.SalesChannel,
		arg.OrderPriority,
		arg.OrderDate,
		arg.OrderID,
		arg.ShipDate,
		arg.UnitsSold,
		arg.UnitPrice,
		arg.UnitCost,
		arg.TotalRevenue,
		arg.TotalCost,
		arg.TotalProfit,
	)
	var i Sale
	err := row.Scan(
		&i.ID,
		&i.Region,
		&i.Country,
		&i.ItemType,
		&i.SalesChannel,
		&i.OrderPriority,
		&i.OrderDate,
		&i.OrderID,
		&i.ShipDate,
		&i.UnitsSold,
		&i.UnitPrice,
		&i.UnitCost,
		&i.TotalRevenue,
		&i.TotalCost,
		&i.TotalProfit,
		&i.CreatedAt,
	)
	return i, err
}

const getSale = `-- name: GetSale :one
SELECT id, region, country, item_type, sales_channel, order_priority, order_date, order_id, ship_date, units_sold, unit_price, unit_cost, total_revenue, total_cost, total_profit, created_at FROM sales
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSale(ctx context.Context, id int64) (Sale, error) {
	row := q.db.QueryRow(ctx, getSale, id)
	var i Sale
	err := row.Scan(
		&i.ID,
		&i.Region,
		&i.Country,
		&i.ItemType,
		&i.SalesChannel,
		&i.OrderPriority,
		&i.OrderDate,
		&i.OrderID,
		&i.ShipDate,
		&i.UnitsSold,
		&i.UnitPrice,
		&i.UnitCost,
		&i.TotalRevenue,
		&i.TotalCost,
		&i.TotalProfit,
		&i.CreatedAt,
	)
	return i, err
}

const getSaleForUpdate = `-- name: GetSaleForUpdate :one
SELECT id, region, country, item_type, sales_channel, order_priority, order_date, order_id, ship_date, units_sold, unit_price, unit_cost, total_revenue, total_cost, total_profit, created_at FROM sales
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetSaleForUpdate(ctx context.Context, id int64) (Sale, error) {
	row := q.db.QueryRow(ctx, getSaleForUpdate, id)
	var i Sale
	err := row.Scan(
		&i.ID,
		&i.Region,
		&i.Country,
		&i.ItemType,
		&i.SalesChannel,
		&i.OrderPriority,
		&i.OrderDate,
		&i.OrderID,
		&i.ShipDate,
		&i.UnitsSold,
		&i.UnitPrice,
		&i.UnitCost,
		&i.TotalRevenue,
		&i.TotalCost,
		&i.TotalProfit,
		&i.CreatedAt,
	)
	return i, err
}

const listSales = `-- name: ListSales :many
SELECT id, region, country, item_type, sales_channel, order_priority, order_date, order_id, ship_date, units_sold, unit_price, unit_cost, total_revenue, total_cost, total_profit, created_at FROM sales
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListSalesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSales(ctx context.Context, arg ListSalesParams) ([]Sale, error) {
	rows, err := q.db.Query(ctx, listSales, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Sale{}
	for rows.Next() {
		var i Sale
		if err := rows.Scan(
			&i.ID,
			&i.Region,
			&i.Country,
			&i.ItemType,
			&i.SalesChannel,
			&i.OrderPriority,
			&i.OrderDate,
			&i.OrderID,
			&i.ShipDate,
			&i.UnitsSold,
			&i.UnitPrice,
			&i.UnitCost,
			&i.TotalRevenue,
			&i.TotalCost,
			&i.TotalProfit,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
