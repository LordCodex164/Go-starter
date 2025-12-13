package main

type User struct {
	id int
	Name string
}

type MailRepo interface {
	send() error
}

type UserRepo interface {
	CreateUser(*User) error
}

type UserService struct {
	repo UserRepo
}

type MailService struct {
	repo MailRepo
}

type Repos struct {
	userRepo UserRepo
}

type Services struct {
	userService UserService
	mailService MailService
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func NewMailService(mailRepo MailRepo) *MailService {
	return &MailService{
		repo: mailRepo,
	}
}

func NewServices (repo *Repos) *Services {
	return &Services{
		userService: *NewUserService(repo.userRepo),
	}
}

func main() {
	//here we can pass the newUserService a db repo
	//NewUserService()
}
