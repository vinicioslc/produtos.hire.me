package migrations

import (
	"bemobi-hire/server/models"

	migrate "github.com/eminetto/mongo-migrate"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

/// Baseado nos exemplos de https://github.com/xakep666/mongo-migrate
func init() {
	carrier := "vivo"
	migrate.Register(func(db *mgo.Database) error { // uP
		return db.C(plansColl).Insert(
			models.Plan{
				ID:          bson.NewObjectId(),
				PlanCarrier: carrier,
				Spotted:     true,
				PlanSKU:     "WEB_VIVOTURBO15DIAS1499",
				PlanAdvantages: []models.PlanAdvantage{
					{
						"Ganhe 1GB de bônus,",
						" válido por 7 dias, ao acumular R$35 em recarga",
					},
					{
						"Minutos ilimitados",
						" em ligações locais para celulares de outras operadoras",
					},
				},
				PlanHighlights: []models.PlanHighlights{
					{
						"https://celular.vivo.com.br/planos/pre/img/whattsicon.png",
						"WhatsApp Ilimitado",
						" para mensagens, vídeos e fotos",
					},
				},
				PlanTitle:       "3GB",
				PlanDetails:     "de internet 4.5G",
				PlanSmallWords:  "com renovação automática",
				PlanLimitMbytes: 2000,
				PlanLimitDays:   15,
				PlanPrice:       14.990,
				PlanMoreDetails: models.PlanDetails{
					Franchise:      `<div class="tab__content" data-main-content="1"> <div class="r1"> <p>A promoção Vivo Pré Turbo <span class="franquia-container">1.5GB</span> está disponível com as seguintes características:</p> <ul class="features-container"><li><b>WhatsApp Ilimitado</b> para envio de mensagens, vídeos e fotos</li><li><b>Ligações ilimitadas</b> para celulares e fixos Vivo de todo o Brasil, usando 15</li><li><b>SMS ilimitado</b> para celulares Vivo</li><li>R$9,99/7 dias</li><li><b data-target="recebainternet">Receba e Envie internet</b></li><li><b data-target="servicosdigitais">Serviços Digitais:</b> NBA e GoRead</li><li><b data-target="vivobis">Vivo Bis</b></li><li class="bonusgiga"><strong>Ganhe 1GB de bônus</strong>, válido por 7 dias, ao acumular R$35 em recarga.</li></ul> </div> <p>A promoção é renovada automaticamente a cada <span class="renovaDias">7</span> dias, basta ter saldo de recarga válido suficiente.</p> <p>Acesse o regulamento <span class="click__regulamentos" data-choose="3">aqui</span></p> </div>`,
					AdditionalInfo: "",
					Regulation:     "",
				},
			},
			models.Plan{
				ID:              bson.NewObjectId(),
				PlanCarrier:     carrier,
				PlanSKU:         "WEB_VIVOTURBO7DIAS1499",
				PlanTitle:       "1.5GB",
				PlanDetails:     "de internet 4.5G",
				PlanSmallWords:  "com renovação automática",
				PlanLimitMbytes: 2000,
				PlanLimitDays:   7,
				PlanPrice:       8.990,
				PlanAdvantages: []models.PlanAdvantage{
					{
						"Ganhe 1GB de bônus,",
						" válido por 7 dias, ao acumular R$35 em recarga",
					},
					{
						"Minutos ilimitados",
						" em ligações locais para celulares de outras operadoras",
					},
				},
				PlanHighlights: []models.PlanHighlights{
					{
						"https://www.pngarts.com/files/3/Telegram-Logo-PNG-Image-Background.png",
						"Telegram Ilimitado",
						"para mensagens, vídeos e fotos",
					},
				},
				PlanMoreDetails: models.PlanDetails{
					Franchise:      `<div class="tab__content" data-main-content="1"> <div class="r1"> <p>A promoção Vivo Pré Turbo <span class="franquia-container">1.5GB</span> está disponível com as seguintes características:</p> <ul class="features-container"><li><b>WhatsApp Ilimitado</b> para envio de mensagens, vídeos e fotos</li><li><b>Ligações ilimitadas</b> para celulares e fixos Vivo de todo o Brasil, usando 15</li><li><b>SMS ilimitado</b> para celulares Vivo</li><li>R$9,99/7 dias</li><li><b data-target="recebainternet">Receba e Envie internet</b></li><li><b data-target="servicosdigitais">Serviços Digitais:</b> NBA e GoRead</li><li><b data-target="vivobis">Vivo Bis</b></li><li class="bonusgiga"><strong>Ganhe 1GB de bônus</strong>, válido por 7 dias, ao acumular R$35 em recarga.</li></ul> </div> <p>A promoção é renovada automaticamente a cada <span class="renovaDias">7</span> dias, basta ter saldo de recarga válido suficiente.</p> <p>Acesse o regulamento <span class="click__regulamentos" data-choose="3">aqui</span></p> </div>`,
					AdditionalInfo: "",
					Regulation:     "",
				},
			},
			models.Plan{
				ID:              bson.NewObjectId(),
				PlanCarrier:     carrier,
				PlanSKU:         "WEB_VIVOTURBO7DIAS1499",
				PlanTitle:       "200MIN",
				PlanDetails:     "de internet 4.5G",
				PlanSmallWords:  "com renovação automática",
				PlanLimitMbytes: 5000,
				PlanLimitDays:   7,
				PlanPrice:       14.990,
				PlanAdvantages: []models.PlanAdvantage{
					{
						"Ganhe 1GB de bônus,",
						" válido por 7 dias, ao acumular R$35 em recarga",
					},
					{
						"Minutos ilimitados",
						" em ligações locais para celulares de outras operadoras",
					},
				},
				PlanHighlights: []models.PlanHighlights{
					{
						"https://celular.vivo.com.br/planos/pre/img/whattsicon.png",
						"WhatsApp Ilimitado",
						"para mensagens, vídeos e fotos",
					},
				},
				PlanMoreDetails: models.PlanDetails{
					Franchise:      `<div class="tab__content" data-main-content="1"> <div class="r1"> <p>A promoção Vivo Pré Turbo <span class="franquia-container">1.5GB</span> está disponível com as seguintes características:</p> <ul class="features-container"><li><b>WhatsApp Ilimitado</b> para envio de mensagens, vídeos e fotos</li><li><b>Ligações ilimitadas</b> para celulares e fixos Vivo de todo o Brasil, usando 15</li><li><b>SMS ilimitado</b> para celulares Vivo</li><li>R$9,99/7 dias</li><li><b data-target="recebainternet">Receba e Envie internet</b></li><li><b data-target="servicosdigitais">Serviços Digitais:</b> NBA e GoRead</li><li><b data-target="vivobis">Vivo Bis</b></li><li class="bonusgiga"><strong>Ganhe 1GB de bônus</strong>, válido por 7 dias, ao acumular R$35 em recarga.</li></ul> </div> <p>A promoção é renovada automaticamente a cada <span class="renovaDias">7</span> dias, basta ter saldo de recarga válido suficiente.</p> <p>Acesse o regulamento <span class="click__regulamentos" data-choose="3">aqui</span></p> </div>`,
					AdditionalInfo: "",
					Regulation:     "",
				},
			},
		)
	}, func(db *mgo.Database) error { // dOWN
		_, err := db.C(plansColl).RemoveAll(bson.M{
			"plan_carrier": carrier,
		})
		return err
	})
}
