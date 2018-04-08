package routes

import (
	"github.com/TiiQu-Network/claim-verifier-prototype/http/controllers"
	"github.com/gobuffalo/buffalo"
)

func Routes(app *buffalo.App) {
	// HomeController routes
	app.GET("/", controllers.Home.Index)
	app.GET("/tutorial/", controllers.Home.Tutorial)
	app.GET("/regenerate/", controllers.Home.Regenerate)

	// Institution routes
	institutionController := new(controllers.Institution)
	app.Resource("/institution/", new(controllers.InstitutionResource))
	app.GET("/institution/{institution}/students/", institutionController.Students)
	app.GET("/institution/{institution}/students/toBlockchain/", institutionController.ToBlockchain)
	app.GET("/institution/{institution}/blockchain/", institutionController.Blockchain)

	// Recipient routes
	studentController := new(controllers.Recipient)
	app.Resource("/student/", new(controllers.RecipientResource))
	app.GET("/student/{student}/to-blockchain/", studentController.ToBlockchain)

	// Platform Members routes
	memberController := new(controllers.Member)
	app.Resource("/member/", new(controllers.MemberResource))
	app.GET("/member/{member}/certifications/", memberController.Certifications)
	app.GET("/member/{member}/certifications/add/", memberController.AddCertification)
	app.GET("/member/{member}/blockchain/", memberController.Blockchain)

	// Certification routes
	memberCertificationController := new(controllers.MemberCertification)
	app.Resource("/member-certification/", new(controllers.MemberCertificationResource))
	app.GET("/member-certification/{memberCertification}/to-blockchain/", memberCertificationController.ToBlockchain)

	// Data-hash routes
	app.Resource("/dataHash/", new(controllers.DataHashResource))
}
