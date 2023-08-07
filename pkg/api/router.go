package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"mvc/pkg/controller"
	"mvc/pkg/models"
	
)

func Start() {

	r := mux.NewRouter()
	r.Use(models.VerifyTokenMiddleware)

	r.HandleFunc("/", controller.Welcome).Methods("GET")
	r.HandleFunc("/signup", controller.AddUserP).Methods("POST")
	r.HandleFunc("/login", controller.LoginUserP).Methods("POST")
	r.HandleFunc("/adminhome", controller.LoginAdminP).Methods("POST")
	r.HandleFunc("/adminhome", controller.WelcomeAdmin).Methods("GET")
	r.HandleFunc("/userlogout", controller.LogoutUser).Methods("POST")

	


	r.HandleFunc("/client/profilepage", controller.ProfilePage).Methods("GET")
	r.HandleFunc("/client/userissue", controller.UserIssueRequests).Methods("GET")
	r.HandleFunc("/client/userreturn", controller.UserReturnRequest).Methods("GET")
	r.HandleFunc("/client/userbrowse", controller.BrowseBooks).Methods("GET")
	r.HandleFunc("/client/userbrowse", controller.RequestIssue).Methods("POST")
	r.HandleFunc("/client/withdrawir", controller.WithdrawIR).Methods("POST")
    r.HandleFunc("/client/withdrawrr", controller.WithdrawRR).Methods("POST")
	r.HandleFunc("/client/requestreturn", controller.RequestReturn).Methods("POST")
    
	r.HandleFunc("/admin/addbook",controller.AddPage).Methods("GET")
	r.HandleFunc("/admin/addbook",controller.Add).Methods("POST")
	r.HandleFunc("/admin/qtyinc",controller.IncQty).Methods("POST")
	r.HandleFunc("/admin/qtydec",controller.DecQty).Methods("POST")
	r.HandleFunc("/admin/booksinv", controller.List).Methods("GET")
	r.HandleFunc("/admin/addadmin",controller.AddAdminPage).Methods("GET")
	r.HandleFunc("/admin/addadmin", controller.AddAdminP).Methods("POST")
	r.HandleFunc("/admin/issuerequests",controller.ListIssueRequest).Methods("GET")
	r.HandleFunc("/admin/returnrequests",controller.ListReturnRequest).Methods("GET")
	r.HandleFunc("/admin/acceptissuereq", controller.AcceptIssue).Methods("POST")
	r.HandleFunc("/admin/rejectissuereq", controller.RejectIssue).Methods("POST")
	r.HandleFunc("/admin/acceptreturnreq", controller.AcceptReturn).Methods("POST")
	r.HandleFunc("/admin/rejectreturnreq", controller.RejectReturn).Methods("POST")
	r.HandleFunc("/admin/viewadmins",controller.ViewAdmins).Methods("GET")
	r.HandleFunc("/adminlogout", controller.LogoutAdmin).Methods("POST")

	http.ListenAndServe(":8000", r)
}