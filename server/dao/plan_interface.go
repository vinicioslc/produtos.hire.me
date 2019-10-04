package dao

import (
	"bemobi-hire/server/models"
)

// IPlansDAO Contains all database manipulation relative to Plans
type IPlansDAO interface {
	Connect() IPlansDAO
	GetPlanByCarrierAndSku(Carrier string, Sku string) (models.DatabasePlan, error)
	ListPlansByCarrier(carrier string) ([]models.Plan, error)
	InsertPlanIntoCarrier(carrier string, plan models.DatabasePlan, result *models.DatabasePlan) error
	DeleteByID(id string) error
	DeleteBySku(sku string) error
	UpdateByID(id string, plan models.DatabasePlan) error
	UpdateBySku(sku string, plan models.DatabasePlan) error
}
