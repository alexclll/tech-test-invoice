package getUsers

type Service struct {
	userRepository UserRepository
}

func (service Service) Execute() (Response, error) {
	users, err := service.userRepository.GetUsers()

	if err != nil {
		return Response{}, err
	}

	return Response{
		Users: users,
	}, nil
}

func NewService(userRepository UserRepository) Service {
	return Service{
		userRepository: userRepository,
	}
}
