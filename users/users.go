package users

import (
	"fmt"
	"time"

	"github.com/faraznicolas/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Nicolas", time.Now(), true)
	fmt.Println(u)
}
