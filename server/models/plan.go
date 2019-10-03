package models

import "gopkg.in/mgo.v2/bson"

// PlanHighlights Pontos Chaves do plano
type PlanHighlights struct {
	Icon        string `bson:"h_icon" json:"h_icon"`
	Title       string `bson:"h_title" json:"h_title"`
	Description string `bson:"h_description" json:"h_description"`
}

// PlanAdvantage Vantagens de um plano
type PlanAdvantage struct {
	Title       string `bson:"a_title" json:"a_title"`
	Description string `bson:"a_description" json:"a_description"`
}

//PlanDetails possui detalhes do plano
type PlanDetails struct {
	Franchise      string `bson:"franchise" json:"franchise"`
	AdditionalInfo string `bson:"additional_info" json:"additional_info"`
	Regulation     string `bson:"regulation" json:"regulation"`
}

// Plan describes an mobile plan showned to users.
type Plan struct {
	ID              bson.ObjectId    `bson:"_id" json:"id"`
	PlanSKU         string           `bson:"plan_sku" json:"plan_sku"`
	Spotted         bool             `bson:"plan_spotted" json:"plan_spotted"`
	PlanCarrier     string           `bson:"plan_carrier" json:"plan_carrier"`
	PlanTitle       string           `bson:"plan_title" json:"plan_title"`
	PlanDetails     string           `bson:"plan_details" json:"plan_details"`
	PlanSmallWords  string           `bson:"plan_small_words" json:"plan_small_words"`
	PlanPrice       float64          `bson:"plan_price" json:"plan_price"`
	PlanHighlights  []PlanHighlights `bson:"plan_highlights" json:"plan_highlights"`
	PlanAdvantages  []PlanAdvantage  `bson:"plan_advantages" json:"plan_advantages"`
	PlanLimitMbytes int              `bson:"plan_limit_mbytes" json:"plan_limit_mbytes"`
	PlanLimitDays   int              `bson:"plan_limit_days" json:"plan_limit_days"`
}

// DatabasePlan is a type used to write into database full plan data .
type DatabasePlan struct {
	ID              bson.ObjectId    `bson:"_id" json:"id"`
	PlanSKU         string           `bson:"plan_sku" json:"plan_sku"`
	Spotted         bool             `bson:"plan_spotted" json:"plan_spotted"`
	PlanCarrier     string           `bson:"plan_carrier" json:"plan_carrier"`
	PlanTitle       string           `bson:"plan_title" json:"plan_title"`
	PlanDetails     string           `bson:"plan_details" json:"plan_details"`
	PlanSmallWords  string           `bson:"plan_small_words" json:"plan_small_words"`
	PlanPrice       float64          `bson:"plan_price" json:"plan_price"`
	PlanHighlights  []PlanHighlights `bson:"plan_highlights" json:"plan_highlights"`
	PlanAdvantages  []PlanAdvantage  `bson:"plan_advantages" json:"plan_advantages"`
	PlanLimitMbytes int              `bson:"plan_limit_mbytes" json:"plan_limit_mbytes"`
	PlanLimitDays   int              `bson:"plan_limit_days" json:"plan_limit_days"`
	PlanMoreDetails PlanDetails      `bson:"plan_more_details" json:"plan_more_details"`
}
