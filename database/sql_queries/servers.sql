-- name: GetAllServers :many
select * from sa_servers;

-- name: GetFirstServer :one
select * from sa_servers order by id desc limit 1;