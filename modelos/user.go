package modelos

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	status    bool
}

func (usuario *User) AddUser(id int, name string, createdat time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdat
	usuario.status = status
}
