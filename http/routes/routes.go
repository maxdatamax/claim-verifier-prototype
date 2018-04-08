package routes

import (
	"github.com/TiiQu-Network/claim-verifier-prototype/http/controllers"
	"github.com/gobuffalo/buffalo"
)

func Routes(app *buffalo.App) {
	// HomeController routes
	home := controllers.Home
	app.GET(home.Index())
	app.GET(home.Tutorial())
	app.GET(home.Regenerate())

	// InstitutionController routes
	inst := controllers.Institution
	app.Resource(inst.Url(""), inst.Resource)
	app.GET(inst.Students())
	app.GET(inst.ToBlockchain())
	app.GET(inst.Blockchain())

	// Recipient routes
	recp := controllers.Recipient
	app.Resource(recp.Url(""), recp.Resource)
	app.GET(recp.ToBlockchain())

	// Platform Members routes
	mbr := controllers.Member
	app.Resource(mbr.Url(""), mbr.Resource)
	app.GET(mbr.Certifications())
	app.GET(mbr.AddCertification())
	app.GET(mbr.Blockchain())

	// Certification routes
	memberCertificationController := new(controllers.MemberCertification)
	app.Resource("/member-certification/", new(controllers.MemberCertificationResource))
	app.GET("/member-certification/{memberCertification}/to-blockchain/", memberCertificationController.ToBlockchain)

	// Data-hash routes
	app.Resource("/dataHash/", new(controllers.DataHashResource))
}
