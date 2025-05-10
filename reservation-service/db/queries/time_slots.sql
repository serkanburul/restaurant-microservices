-- name: GetTimeSlot :one
SELECT * FROM time_slots
WHERE start_time = $1;

-- name: ListTimeSlots :many
SELECT start_time FROM time_slots;
