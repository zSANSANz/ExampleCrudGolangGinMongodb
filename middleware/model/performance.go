package model

import (
	"time"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Performance struct {
	*PerformanceInput `bson:",inline"`
	ID                primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type PerformanceInput struct {
	IdEmployee       primitive.ObjectID `json:"id_employee" bson:"id_employee"`
	Employee         []User             `json:"employee" bson:"employee"`
	Devision         []Devision         `json:"devision" bson:"devision"`
	Status           string             `json:"status" bson:"status"`
	RevisionCount    int32              `json:"revision_count" bson:"revision_count"`
	TaskProductivity []TaskProductivity `json:"task_productivity" bson:"task_productivity"`
	TaskPersonal     []TaskPersonal     `json:"task_personal" bson:"task_personal"`
	KpiStatus        string             `json:"kpi_status" bson:"kpi_status"`
	Noted            []Noted            `json:"noted" bson:"noted"`
	TotalScore       float64            `json:"total_score" bson:"total_score"`
	ScoreCategory    string             `json:"score_category" bson:"score_category"`
	CreatedBy        string             `json:"created_by" bson:"created_by"`
	CreatedByID      string             `json:"created_by_id" bson:"created_by_id"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedBy        string             `json:"update_by" bson:"update_by"`
	UpdatedByID      string             `json:"update_by_Id" bson:"update_by_id"`
	UpdatedAt        time.Time          `json:"update_at" bson:"update_at"`
	DeletedBy        string             `json:"deleted_by" bson:"deleted_by"`
	DeletedAt        time.Time          `json:"deleted_at" bson:"deleted_at"`
	Deleted          bool               `json:"deleted" bson:"deleted"`
}

type PagedPerformance struct {
	Data     []Performance                  `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}

type TaskProductivity struct {
	Name   string    `json:"name,omitempty" bson:"name"`
	Nilai  float64   `json:"nilai,omitempty" bson:"nilai"`
	Weight float64   `json:"weight,omitempty" bson:"weight"`
	Date   time.Time `json:"date,omitempty" bson:"date"`
	Reason string    `json:"reason,omitempty" bson:"reason"`
}

type TaskPersonal struct {
	Name   string    `json:"name,omitempty" bson:"name"`
	Nilai  float64   `json:"nilai,omitempty" bson:"nilai"`
	Weight float64   `json:"weight,omitempty" bson:"weight"`
	Date   time.Time `json:"date,omitempty" bson:"date"`
	Reason string    `json:"reason,omitempty" bson:"reason"`
}

type Noted struct {
	Description string    `json:"description,omitempty" bson:"description"`
	CreatedBy   string    `json:"created_by,omitempty" bson:"created_by"`
	Date        time.Time `json:"date,omitempty" bson:"date"`
}
