package controllers

import (
	"bemobi-hire/server/dao"
	serverHelpers "bemobi-hire/server/helpers"
	"bemobi-hire/server/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// PlanController Contains all logic about plans
type PlanController struct {
	planDAO dao.PlansDAO
}

// InjectDAO Injects the dao instance in PlanController
func (ctrl *PlanController) InjectDAO(dao dao.PlansDAO) *PlanController {
	ctrl.planDAO = dao
	return ctrl
}

// GetAllCarrierPlans Return all plans from given carrier
func (ctrl *PlanController) GetAllCarrierPlans(w http.ResponseWriter, r *http.Request) {

	carrier := chi.URLParam(r, "carrier")
	if carrier == "" {
		serverHelpers.RespondWithError(w, 400, "Carrier missing argument")
	}
	plans, err := ctrl.planDAO.ListPlansByCarrier(carrier)
	if err != nil {
		serverHelpers.RespondWithError(w, 500, err.Error())
		return
	}
	if plans == nil {
		serverHelpers.RespondwithJSON(w, 200, ([]models.Plan{}))
		return
	}
	serverHelpers.RespondwithJSON(w, 200, plans)
}

// GetPlanDetails Return details from an plan from a given carrier and sku
func (ctrl *PlanController) GetPlanDetails(w http.ResponseWriter, r *http.Request) {

	carrier := chi.URLParam(r, "carrier")
	sku := chi.URLParam(r, "sku")
	planDetails, err := ctrl.planDAO.GetPlanByCarrierAndSku(carrier, sku)

	if planDetails.PlanSKU == "" && (err.Error() == "not found" || err == nil) {
		serverHelpers.RespondWithError(w, http.StatusNotFound, "Plan not found")
		return
	}
	if err != nil {
		serverHelpers.RespondWithError(w, 500, err.Error())
		return
	}

	serverHelpers.RespondwithJSON(w, http.StatusOK, planDetails.PlanMoreDetails)

}

// InsertPlanFromCarrier Insert an plan into a carrier
func (ctrl *PlanController) InsertPlanFromCarrier(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // good practique following https://stackoverflow.com/questions/23928983/defer-body-close-after-receiving-responses
	carrier := chi.URLParam(r, "carrier")
	receivedPlan := models.DatabasePlan{}
	resultInsert := models.DatabasePlan{}
	if err := json.NewDecoder(r.Body).Decode(&receivedPlan); err != nil {
		serverHelpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := ctrl.planDAO.InsertPlanIntoCarrier(carrier, receivedPlan, &resultInsert); err != nil {
		serverHelpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	serverHelpers.RespondwithJSON(w, http.StatusOK, resultInsert)
}

// UpdatePlanBySKU Insert an plan into a carrier
func (ctrl *PlanController) UpdatePlanBySKU(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // good practique following https://stackoverflow.com/questions/23928983/defer-body-close-after-receiving-responses

	sku := chi.URLParam(r, "sku")
	updatedPlan := models.DatabasePlan{}
	if err := json.NewDecoder(r.Body).Decode(&updatedPlan); err != nil {
		serverHelpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := ctrl.planDAO.UpdateBySku(sku, updatedPlan); err != nil {
		serverHelpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	serverHelpers.RespondwithJSON(w, http.StatusOK, updatedPlan)
}
