-- name: CreateSubscription :exec
INSERT INTO subscription( email )
VALUES ( $1 );

-- name: ListSubscriptionEmail :many
SELECT email FROM subscription
WHERE status = true;