package services

import "argentina-tresury/model"

func GetUsers() ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := model.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (*model.User, error) {
	user := &model.User{}
	if err := model.DB.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *model.User) error {
	user.Password = HashAndSalt([]byte(user.Password))
	return model.DB.Create(user).Error
}

func UpdateUser(id uint, user *model.User) (*model.User, error) {
	u := &model.User{}
	if err := model.DB.First(u, id).Error; err != nil {
		return nil, err
	}
	if u == nil {
		user = nil
		return nil, nil
	}
	u.ChapterID = user.ChapterID
	u.Profile = user.Profile
	if user.Password != u.Password {
		u.Password = HashAndSalt([]byte(user.Password))
	}
	if err := model.DB.Save(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func DeleteUser(id uint) error {
	return model.DB.Delete(&model.User{}, id).Error
}
