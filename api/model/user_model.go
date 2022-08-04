package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetDataUserResponse struct {
		Data []DataUser `json:"users"`
	}
	DataUser struct {
		ID              primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
		IdEmp           string             `json:"id_emp" xml:"id_emp" bson:"id_emp" validate:"required"`
		FirstName       string             `json:"firstName" xml:"firstName" bson:"firstName" validate:"required"`
		LastName        string             `json:"lastName" xml:"lastName" bson:"lastName" validate:"required"`
		Email           string             `json:"email" xml:"email" bson:"email" validate:"required,email"`
		Password        string             `json:"password,omitempty" xml:"password,omitempty" bson:"password" validate:"required"`
		UpdatedPassword bool               `json:"updated_password" xml:"updated_password" bson:"updated_password" validate:"required"`
		Head            bool               `json:"head,omitempty" xml:"head,omitempty" bson:"head" validate:"required"`
		ApproveByHrd    bool               `json:"approve_by_hrd,omitempty" xml:"approve_by_hrd,omitempty" bson:"approve_by_hrd"`
		ApproveByCfo    bool               `json:"approve_by_cfo,omitempty" xml:"approve_by_cfo,omitempty" bson:"approve_by_cfo"`
		Superior        string             `json:"superior" bson:"superior"`
		Department      string             `json:"department" bson:"department"`
		Division        string             `json:"division,omitempty" xml:"division,omitempty" bson:"division"`
		Section         string             `json:"section" bson:"section"`
		Position        string             `json:"position" xml:"position" bson:"position" validate:"required"`
		CreatedBy       string             `json:"created_by" bson:"created_by"`
		CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
		UpdatedBy       string             `json:"update_by" bson:"update_by"`
		UpdatedAt       *time.Time         `json:"update_at" bson:"update_at"`
		DeletedBy       string             `json:"deleted_by" bson:"deleted_by"`
		DeletedAt       *time.Time         `json:"deleted_at" bson:"deleted_at"`
	}

	DataUserRequest struct {
		ID              primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
		IdEmp           string             `json:"id_emp" xml:"id_emp" bson:"id_emp" validate:"required"`
		FirstName       string             `json:"firstName" xml:"firstName" bson:"firstName" validate:"required"`
		LastName        string             `json:"lastName" xml:"lastName" bson:"lastName" validate:"required"`
		Email           string             `json:"email" xml:"email" bson:"email" validate:"required,email"`
		Password        string             `json:"password,omitempty" xml:"password,omitempty" bson:"password" validate:"required"`
		UpdatedPassword bool               `json:"updated_password" xml:"updated_password" bson:"updated_password" validate:"required"`
		Head            bool               `json:"head,omitempty" xml:"head,omitempty" bson:"head" validate:"required"`
		ApproveByHrd    bool               `json:"approve_by_hrd,omitempty" xml:"approve_by_hrd,omitempty" bson:"approve_by_hrd"`
		ApproveByCfo    bool               `json:"approve_by_cfo,omitempty" xml:"approve_by_cfo,omitempty" bson:"approve_by_cfo"`
		Superior        string             `json:"superior" bson:"superior"`
		Department      string             `json:"department" bson:"department"`
		Division        string             `json:"division,omitempty" xml:"division,omitempty" bson:"division"`
		Section         string             `json:"section" bson:"section"`
		Position        string             `json:"position" xml:"position" bson:"position" validate:"required"`
		CreatedBy       string             `json:"created_by" bson:"created_by"`
		CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
		UpdatedBy       string             `json:"update_by" bson:"update_by"`
		UpdatedAt       *time.Time         `json:"update_at" bson:"update_at"`
		DeletedBy       string             `json:"deleted_by" bson:"deleted_by"`
		DeletedAt       *time.Time         `json:"deleted_at" bson:"deleted_at"`
	}

	DataUserPerformance struct {
		ID           primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
		IdEmp        string             `json:"id_emp" xml:"id_emp" bson:"id_emp" validate:"required"`
		FirstName    string             `json:"firstName" xml:"firstName" bson:"firstName" validate:"required"`
		LastName     string             `json:"lastName" xml:"lastName" bson:"lastName" validate:"required"`
		Email        string             `json:"email" xml:"email" bson:"email" validate:"required,email"`
		Head         bool               `json:"head,omitempty" xml:"head,omitempty" bson:"head" validate:"required"`
		ApproveByHrd bool               `json:"approve_by_hrd,omitempty" xml:"approve_by_hrd,omitempty" bson:"approve_by_hrd"`
		ApproveByCfo bool               `json:"approve_by_cfo,omitempty" xml:"approve_by_cfo,omitempty" bson:"approve_by_cfo"`
		Superior     string             `json:"superior" bson:"superior"`
		Department   string             `json:"department" bson:"department"`
		Division     string             `json:"division,omitempty" xml:"division,omitempty" bson:"division"`
		Section      string             `json:"section" bson:"section"`
		Position     string             `json:"position" xml:"position" bson:"position" validate:"required"`
	}
)
