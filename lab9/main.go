package main
import ("log")


type User struct {
	id int
	name string
	init bool
}

func (u *User) define(id int, n string){
	u.id = id
	u.name = n
	u.init = true
} 

func define(user *User, id int, n string){
	user.id = id
	user.name = n
	user.init = true
} 

//działa
func NewUser(id int, n string) User{
	return User{
		id:id,
		name:n,
		init:true,
	}
}

func main() {

	user := User{}
	define(&user,2,"Mateusz")
	//_user2 := NewUser(1, "John")
	if user.init {
		log.Println("User initialized")
	} else {
		log.Fatal("User not initialized")
	}
}