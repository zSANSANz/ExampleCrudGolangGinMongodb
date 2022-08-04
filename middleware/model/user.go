package model

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	*UserInput `bson:",inline"`
	ID         primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type UserInput struct {
	IdEmp           string `json:"id_emp" xml:"id_emp" bson:"id_emp" validate:"required"`
	FirstName       string `json:"firstName" xml:"firstName" bson:"firstName" validate:"required"`
	LastName        string `json:"lastName" xml:"lastName" bson:"lastName" validate:"required"`
	Email           string `json:"email" xml:"email" bson:"email" validate:"required,email"`
	Password        string `json:"password,omitempty" xml:"password,omitempty" bson:"password" validate:"required"`
	UpdatedPassword bool   `json:"updated_password" xml:"updated_password" bson:"updated_password" validate:"required"`
	Head            bool   `json:"head,omitempty" xml:"head,omitempty" bson:"head" validate:"required"`
	ApproveByHrd    bool   `json:"approve_by_hrd,omitempty" xml:"approve_by_hrd,omitempty" bson:"approve_by_hrd"`
	ApproveByCfo    bool   `json:"approve_by_cfo,omitempty" xml:"approve_by_cfo,omitempty" bson:"approve_by_cfo"`
	Superior        string `json:"superior" bson:"superior"`
	Division        string `json:"division,omitempty" xml:"division,omitempty" bson:"division"`
	Department      string `json:"department" bson:"department"`
	Section         string `json:"section" bson:"section"`
	Position        string `json:"position" xml:"position" bson:"position" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" xml:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" bson:"password" validate:"required"`
	Head     bool   `json:"head,omitempty" xml:"head,omitempty" bson:"head" validate:"required"`
}

type PagedUser struct {
	Data     []User                         `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}

type LoginReturn struct {
	Email string `json:"email" xml:"email" bson:"email" validate:"required,email"`
	Head  bool   `json:"head,omitempty" xml:"head,omitempty" bson:"head" validate:"required"`
}
