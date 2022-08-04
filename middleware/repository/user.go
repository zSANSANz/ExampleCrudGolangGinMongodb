package repository

import (
	"chatnews-api/middleware/exception"
	"chatnews-api/middleware/model"
	"context"

	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cntx context.Context = context.TODO()

type UserRepository interface {
	GetAllUser(page int64, limit int64) (*model.PagedUser, error)
	SaveUser(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	GetUser(id string) (*model.User, error)
	GetUserByToken(id string) (*model.User, error)
	GetUserHod(id string) ([]model.User, error)
	GetUserHodAll(id string) ([]model.User, error)
	GetUserSupervision(id string) ([]model.User, error)
	GetUserName(id string) ([]model.User, error)
	GetUserSupervisionAll(id string) ([]model.User, error)
	UpdateUser(id string, user *model.User) (*model.User, error)
	UpdateUserPassword(id string, user *model.User) (*model.User, error)
	UpdateUserPosition(id string, user *model.User) (*model.User, error)
	UpdateUserSuperior(id string, user *model.User) (*model.User, error)
	UpdateUserSection(id string, user *model.User) (*model.User, error)
	UpdateUserEmpId(id string, user *model.User) (*model.User, error)
	DeleteUser(id string) error
}

type userRepositoryImpl struct {
	Connection *mongo.Database
}

func (userRepository *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var existingUser model.User
	filter := bson.M{"email": email}
	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	if err != nil {
		return nil, err
	}
	return &existingUser, nil
}

func NewUserRepository(Connection *mongo.Database) UserRepository {
	return &userRepositoryImpl{Connection: Connection}
}

func (userRepository *userRepositoryImpl) GetAllUser(page int64, limit int64) (*model.PagedUser, error) {
	var users []model.User

	filter := bson.M{}

	collection := userRepository.Connection.Collection("users")

	projection := bson.D{
		{"id", 1},
		{"id_emp", 1},
		{"firstName", 1},
		{"lastName", 1},
		{"email", 1},
		{"head", 1},
		{"superior", 1},
		{"division", 1},
		{"department", 1},
		{"section", 1},
		{"position", 1},
	}

	paginatedData, err := paginate.New(collection).Context(cntx).Limit(limit).Page(page).Select(projection).Filter(filter).Decode(&users).Find()
	if err != nil {
		return nil, err
	}

	return &model.PagedUser{
		Data:     users,
		PageInfo: paginatedData.Pagination,
	}, nil
}

func (userRepository *userRepositoryImpl) GetUserHod(id string) ([]model.User, error) {
	var existingUsers []model.User
	filter := bson.M{
		"division": bson.M{"$regex": id},
	}

	cursor, err := userRepository.Connection.Collection("users").Find(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "division", id)
	}

	if err = cursor.All(cntx, &existingUsers); err != nil {
		return nil, exception.ResourceNotFoundException("User", "division", id)
	}

	return existingUsers, nil
}

func (userRepository *userRepositoryImpl) GetUserHodAll(id string) ([]model.User, error) {
	var existingUsers []model.User
	filter := bson.M{}

	cursor, err := userRepository.Connection.Collection("users").Find(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "supervision", id)
	}

	if err = cursor.All(cntx, &existingUsers); err != nil {
		return nil, exception.ResourceNotFoundException("User", "supervision", id)
	}

	return existingUsers, nil
}

func (userRepository *userRepositoryImpl) GetUserSupervision(id string) ([]model.User, error) {
	var existingUsers []model.User
	filter := bson.M{
		"superior": bson.M{"$regex": id},
	}

	cursor, err := userRepository.Connection.Collection("users").Find(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "supervision", id)
	}

	if err = cursor.All(cntx, &existingUsers); err != nil {
		return nil, exception.ResourceNotFoundException("User", "supervision", id)
	}

	return existingUsers, nil
}

func (userRepository *userRepositoryImpl) GetUserName(id string) ([]model.User, error) {
	var existingUsers []model.User
	filter := bson.M{
		"firstName": bson.M{"$regex": id},
	}

	cursor, err := userRepository.Connection.Collection("users").Find(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "name", id)
	}

	if err = cursor.All(cntx, &existingUsers); err != nil {
		return nil, exception.ResourceNotFoundException("User", "name", id)
	}

	return existingUsers, nil
}

func (userRepository *userRepositoryImpl) GetUserSupervisionAll(id string) ([]model.User, error) {
	var existingUsers []model.User
	filter := bson.M{
		"superior": bson.M{"$regex": " "},
	}

	cursor, err := userRepository.Connection.Collection("users").Find(cntx, filter)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "supervision", id)
	}

	if err = cursor.All(cntx, &existingUsers); err != nil {
		return nil, exception.ResourceNotFoundException("User", "supervision", id)
	}

	return existingUsers, nil
}

func (userRepository *userRepositoryImpl) GetUser(id string) (*model.User, error) {
	var existingUser model.User
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	existingUser.Password = ""
	return &existingUser, nil
}

func (userRepository *userRepositoryImpl) GetUserByToken(id string) (*model.User, error) {
	var existingUser model.User
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	err := userRepository.Connection.Collection("users").FindOne(cntx, filter).Decode(&existingUser)
	if err != nil {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	existingUser.Password = ""
	return &existingUser, nil
}

func (userRepository *userRepositoryImpl) SaveUser(user *model.User) (*model.User, error) {
	user.ID = primitive.NewObjectID()

	_, err := userRepository.Connection.Collection("users").InsertOne(cntx, user)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (userRepository *userRepositoryImpl) UpdateUser(id string, user *model.User) (*model.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	result, err := userRepository.Connection.Collection("users").ReplaceOne(cntx, filter, user)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.ID = objectId
	user.Password = ""
	return user, nil
}

func (userRepository *userRepositoryImpl) UpdateUserPassword(id string, user *model.User) (*model.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	dataObjectID := bson.M{"$set": bson.M{
		"password":         user.Password,
		"updated_password": true,
	}}

	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, filter, dataObjectID)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.ID = objectId
	user.FirstName = "doesnt changed"
	user.LastName = "doesnt changed"
	user.UpdatedPassword = true
	user.Email = "doesnt changed"
	return user, nil
}

func (userRepository *userRepositoryImpl) UpdateUserPosition(id string, user *model.User) (*model.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	dataObjectID := bson.M{"$set": bson.M{
		"position": user.Position,
	}}

	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, filter, dataObjectID)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.ID = objectId
	return user, nil
}

func (userRepository *userRepositoryImpl) UpdateUserSection(id string, user *model.User) (*model.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	dataObjectID := bson.M{"$set": bson.M{
		"section": user.Section,
	}}

	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, filter, dataObjectID)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.ID = objectId
	return user, nil
}

func (userRepository *userRepositoryImpl) UpdateUserSuperior(id string, user *model.User) (*model.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	dataObjectID := bson.M{"$set": bson.M{
		"superior": user.Superior,
	}}

	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, filter, dataObjectID)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.ID = objectId
	return user, nil
}

func (userRepository *userRepositoryImpl) UpdateUserEmpId(id string, user *model.User) (*model.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	dataObjectID := bson.M{"$set": bson.M{
		"id_emp": user.IdEmp,
	}}

	result, err := userRepository.Connection.Collection("users").UpdateOne(cntx, filter, dataObjectID)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}

	user.ID = objectId
	return user, nil
}

func (userRepository *userRepositoryImpl) DeleteUser(id string) error {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}

	result, err := userRepository.Connection.Collection("users").DeleteOne(cntx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return exception.ResourceNotFoundException("User", "id", id)
	}

	return nil
}
