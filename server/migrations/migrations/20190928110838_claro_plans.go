package migrations

import (
	"bemobi-hire/server/models"

	migrate "github.com/eminetto/mongo-migrate"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

/// Baseado nos exemplos de https://github.com/xakep666/mongo-migrate
func init() {
	carrier := "claro"
	migrate.Register(func(db *mgo.Database) error { // uP
		return db.C(plansCollection).Insert(
			models.DatabasePlan{
				ID:          bson.NewObjectId(),
				PlanCarrier: carrier,
				PlanSKU:     "WEB_CLARO100MB",
				PlanHighlights: []models.PlanHighlights{
					{
						Title: "+10 minutos de ligações para outras operadoras*",
					},
				},
				PlanAdvantages: []models.PlanAdvantage{
					{
						"Minutos ilimitados ",
						"em ligações locais para celulares de outras operadoras",
					},
				},
				PlanTitle:       "100 MB",
				PlanDetails:     "Por dia",
				PlanSmallWords:  "sem renovação automática",
				PlanLimitMbytes: 5000,
				PlanLimitDays:   1,
				PlanPrice:       1.990,
				PlanMoreDetails: models.PlanDetails{
					Franchise: `Prezão R$1,99 / dia
 
					Serviços inclusos:
					Ligações ilimitadas para Claro de todo o Brasil, usando o 21.
					10 minutos de ligações para outras operadoras de todo o Brasil, usando o 21.*
					* Benefício válido apenas para chamadas feitas dentro da área de registro, DDD da linha do cliente.
					100MB para navegar na internet 4.5G da Claro.
					WhatsApp para navegar ilimitado. Envie e receba mensagens, fotos, vídeos e áudios através do aplicativo.
					Chamadas de voz e vídeo (VOIP) e todo acesso a conteúdo externos (links recebidos) que saiam do WhatsApp não estão inclusos no benefício e descontarão da franquia de internet.
					SMS ilimitado para Claro.
					15 SMS para outras operadoras.
					 
					Como contratar: se você já é cliente envie um SMS com a palavra DIA para 2006 e aproveite!`,
					AdditionalInfo: "",
					Regulation:     "",
				},
			},
			models.DatabasePlan{
				ID:          bson.NewObjectId(),
				PlanCarrier: carrier,
				PlanSKU:     "WEB_CLARO1GB",
				Spotted:     true,
				PlanHighlights: []models.PlanHighlights{
					{
						Title: "+1GB de bônus com renovação semanal*",
					},
				},
				PlanAdvantages: []models.PlanAdvantage{
					{
						"Ganhe 1GB de bônus, ",
						"válido por 7 dias, ao acumular R$35 em recarga",
					},
				},
				PlanTitle:       "1GB",
				PlanDetails:     "Por semana",
				PlanSmallWords:  "com renovação automática",
				PlanLimitMbytes: 1000,
				PlanLimitDays:   7,
				PlanPrice:       9.99,
				PlanMoreDetails: models.PlanDetails{
					Franchise: `
<div style="margin-left: 40px;"><strong>2GB (1GB + 1GB* de bônus com renovação semanal)</strong></div>

<div style="margin-left: 40px;"><strong>Apenas: R$9,99 / 7 dias</strong></div>

<div style="margin-left: 40px;">&nbsp;</div>

<div style="margin-left: 40px;"><strong>Serviços inclusos:</strong></div>

<ul>
<li style="margin-left: 40px;">2GB de internet (1GB + 1GB* de bônus com renovação semanal).</li>
</ul>

<div style="margin-left: 80px;">* Bônus de 1GB é concedido automaticamente na renovação da promoção e válido por 7 dias.</div>

<ul>
<li style="margin-left: 40px;">Ligações ilimitadas* para qualquer operadora do Brasil, usando o 21.</li>
</ul>

<div style="margin-left: 80px;">* Benefício válido apenas para chamadas feitas dentro da área de registro, DDD da linha do cliente.</div>

<ul>
<li style="margin-left: 40px;">WhatsApp para navegar ilimitado, válido enquanto o pacote contratado estiver ativo, estão inclusos envio e recebimento de mensagens, fotos, vídeos e áudios através do aplicativo e, promocionalmente, chamadas de voz e vídeo via VOIP sem custo adicional. Todo acesso a conteúdo externos (links recebidos) que saiam do WhatsApp não estão inclusos e descontarão da franquia de internet.</li>
<li style="margin-left: 40px;">SMS à vontade para Claro e 100 SMS para outras operadoras.</li>
</ul>

<div style="margin-left: 40px;"><strong>Vantagens para navegar ilimitado:</strong></div>

<ul>
<li style="margin-left: 40px;">Claro música: escute quando e onde quiser sem descontar da sua internet até 40 playlists que são atualizadas semanalmente.</li>
</ul>

<div style="margin-left: 40px;"><strong>Vantagens de conteúdo ilimitado:</strong></div>

<ul>
<li style="margin-left: 40px;">Claro video: tenha acesso a mais de 15 mil vídeos, séries e desenhos para se divertir a qualquer hora do dia e em qualquer lugar. Acesse todo o conteúdo pelo seu computador, tablet, smartphone.</li>
<li style="margin-left: 40px;">Claro notícias: você bem informado sobre tudo que está acontecendo agora no Brasil e no mundo.</li>
</ul>

<div style="margin-left: 40px;">&nbsp;</div>

<div style="margin-left: 40px;"><strong>Como contratar:</strong> se você já é cliente envie um SMS com a palavra SEMANA para 2006 e aproveite!</div>

<div style="margin-left: 40px;">&nbsp;</div>

<div style="margin-left: 40px;"><strong>Informações gerais:</strong></div>

<ul>
<li style="margin-left: 40px;">Os valores vigentes (sem promoção) poderão ser a cada 12 meses, a contar da data do lançamento do plano. O índice para reajuste deve ser o IGPDI ou na falta deste por qualquer outro índice oficial em vigor. Os valores das promoções poderão ser alterados a qualquer momento.</li>
<li style="margin-left: 40px;">Velocidade média:</li>
<li style="margin-left: 80px;">- Tecnologia 4G = até 5Mbps/512Kbps velocidade máxima de download/upload | 2Mbps velocidade média de download | 128Kbps velocidade mínima</li>
<li style="margin-left: 80px;">- Tecnologia 3G = até 1Mbps/128Kbps velocidade máxima de download/upload | 650Kbps velocidade média de download | 128Kbps velocidade mínima</li>
<li style="margin-left: 80px;">- Tecnologia 2G = até 60Kbps/16Kbps velocidade máxima de download/upload |12Kbps velocidade média de download | 8Kbps velocidade mínima</li>
<li style="margin-left: 40px;">O serviço será ativado dependendo do local de compra do chip e de cadastramento da promoção e dados pessoais, podendo levar até 24 horas após a conclusão do cadastro.</li>
</ul>`,
					AdditionalInfo: "",
					Regulation:     "",
				},
			}, models.DatabasePlan{
				ID:          bson.NewObjectId(),
				PlanCarrier: carrier,
				PlanSKU:     "WEB_CLARO5GB",
				PlanHighlights: []models.PlanHighlights{
					{
						Title: "Ligações ilimitadas*",
					},
				},
				PlanAdvantages: []models.PlanAdvantage{
					{
						"Ligações ilimitadas, ",
						"válido por 7 dias, ao acumular R$35 em recarga",
					},
					{
						"Minutos ilimitados ",
						"em ligações locais para celulares de outras operadoras",
					},
				},
				PlanTitle:       "5GB",
				PlanDetails:     "Por Semana",
				PlanSmallWords:  "com renovação automática",
				PlanLimitMbytes: 1000,
				PlanLimitDays:   7,
				PlanPrice:       14.99,
				PlanMoreDetails: models.PlanDetails{
					Franchise: `
<div style="margin-left: 40px;"><strong>2GB (1GB + 1GB* de bônus com renovação semanal)</strong></div>

<div style="margin-left: 40px;"><strong>Apenas: R$9,99 / 7 dias</strong></div>

<div style="margin-left: 40px;">&nbsp;</div>

<div style="margin-left: 40px;"><strong>Serviços inclusos:</strong></div>

<ul>
<li style="margin-left: 40px;">2GB de internet (1GB + 1GB* de bônus com renovação semanal).</li>
</ul>

<div style="margin-left: 80px;">* Bônus de 1GB é concedido automaticamente na renovação da promoção e válido por 7 dias.</div>

<ul>
<li style="margin-left: 40px;">Ligações ilimitadas* para qualquer operadora do Brasil, usando o 21.</li>
</ul>

<div style="margin-left: 80px;">* Benefício válido apenas para chamadas feitas dentro da área de registro, DDD da linha do cliente.</div>

<ul>
<li style="margin-left: 40px;">WhatsApp para navegar ilimitado, válido enquanto o pacote contratado estiver ativo, estão inclusos envio e recebimento de mensagens, fotos, vídeos e áudios através do aplicativo e, promocionalmente, chamadas de voz e vídeo via VOIP sem custo adicional. Todo acesso a conteúdo externos (links recebidos) que saiam do WhatsApp não estão inclusos e descontarão da franquia de internet.</li>
<li style="margin-left: 40px;">SMS à vontade para Claro e 100 SMS para outras operadoras.</li>
</ul>

<div style="margin-left: 40px;"><strong>Vantagens para navegar ilimitado:</strong></div>

<ul>
<li style="margin-left: 40px;">Claro música: escute quando e onde quiser sem descontar da sua internet até 40 playlists que são atualizadas semanalmente.</li>
</ul>

<div style="margin-left: 40px;"><strong>Vantagens de conteúdo ilimitado:</strong></div>

<ul>
<li style="margin-left: 40px;">Claro video: tenha acesso a mais de 15 mil vídeos, séries e desenhos para se divertir a qualquer hora do dia e em qualquer lugar. Acesse todo o conteúdo pelo seu computador, tablet, smartphone.</li>
<li style="margin-left: 40px;">Claro notícias: você bem informado sobre tudo que está acontecendo agora no Brasil e no mundo.</li>
</ul>

<div style="margin-left: 40px;">&nbsp;</div>

<div style="margin-left: 40px;"><strong>Como contratar:</strong> se você já é cliente envie um SMS com a palavra SEMANA para 2006 e aproveite!</div>

<div style="margin-left: 40px;">&nbsp;</div>

<div style="margin-left: 40px;"><strong>Informações gerais:</strong></div>

<ul>
<li style="margin-left: 40px;">Os valores vigentes (sem promoção) poderão ser a cada 12 meses, a contar da data do lançamento do plano. O índice para reajuste deve ser o IGPDI ou na falta deste por qualquer outro índice oficial em vigor. Os valores das promoções poderão ser alterados a qualquer momento.</li>
<li style="margin-left: 40px;">Velocidade média:</li>
<li style="margin-left: 80px;">- Tecnologia 4G = até 5Mbps/512Kbps velocidade máxima de download/upload | 2Mbps velocidade média de download | 128Kbps velocidade mínima</li>
<li style="margin-left: 80px;">- Tecnologia 3G = até 1Mbps/128Kbps velocidade máxima de download/upload | 650Kbps velocidade média de download | 128Kbps velocidade mínima</li>
<li style="margin-left: 80px;">- Tecnologia 2G = até 60Kbps/16Kbps velocidade máxima de download/upload |12Kbps velocidade média de download | 8Kbps velocidade mínima</li>
<li style="margin-left: 40px;">O serviço será ativado dependendo do local de compra do chip e de cadastramento da promoção e dados pessoais, podendo levar até 24 horas após a conclusão do cadastro.</li>
</ul>`,
					AdditionalInfo: "",
					Regulation:     "",
				},
			},
		)
	}, func(db *mgo.Database) error { // dOWN
		_, err := db.C(plansCollection).RemoveAll(bson.M{
			"plan_carrier": carrier,
		})
		return err
	})
}
