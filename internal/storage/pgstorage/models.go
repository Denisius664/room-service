package pgstorage

type RoomInfo struct {
	Name  string   `db:"name"`
	Users []string `db:"users"`
}

const (
	tableName = "roomsInfo"

	NameColumnName  = "name"
	UsersColumnName = "users"
)
