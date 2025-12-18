package models

type RoomInfo struct {
	Name  string
	Users []string
}

func (r *RoomInfo) Join(username string) {
	r.Users = append(r.Users, username)
}

type RoomEvent struct {
	Name    string
	Content string
}

type PlayerCommand struct {
	Name    string
	Content string
}
