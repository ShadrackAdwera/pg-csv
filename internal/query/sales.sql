-- name: CreateSale :one
INSERT INTO sales (region, country,item_type,sales_channel, order_priority, order_date, order_id, ship_date,units_sold, unit_price,unit_cost,total_revenue,total_cost,total_profit)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) 
RETURNING *;

-- name: ListSales :many
SELECT * FROM sales
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetSale :one
SELECT * FROM sales
WHERE id = $1 LIMIT 1;

-- name: GetSaleForUpdate :one
SELECT * FROM sales
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

