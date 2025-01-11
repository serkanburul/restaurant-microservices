-- name: CreateTableReservation :one
INSERT INTO reservation (token, name, email, status, table_no, time_slot_id, reservation_date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING *;

-- name: GetReservationsByDateAndCapacity :many
SELECT r.table_no, t.start_time
FROM reservation r
         RIGHT JOIN time_slots t
                    ON r.time_slot_id = t.id
         JOIN tables
              ON r.table_no = tables.id
WHERE r.reservation_date = $1 AND
    tables.capacity = $2;

-- name: GetReservationByToken :one
SELECT r.name, r.email, r.status, r.table_no, r.reservation_date, r.created_at, ta.capacity, t.start_time FROM reservation r
JOIN tables ta ON r.table_no = ta.id
JOIN time_slots t ON r.time_slot_id = t.id
WHERE token = $1;

-- name: UpdateReservation :one
UPDATE reservation
SET name = $1, status = $2, table_no = $3, reservation_date = $4, time_slot_id = $5
WHERE token = $6
RETURNING *;

-- name: ListReservations :many
SELECT * FROM reservation
ORDER BY reservation_date DESC;

-- name: DeleteReservationByToken :exec
DELETE FROM reservation
WHERE token = $1;
