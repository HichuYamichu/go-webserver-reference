package user

func (UserService) Delete(id int) {
	userRepo.DeleteOneByID(id)
}
