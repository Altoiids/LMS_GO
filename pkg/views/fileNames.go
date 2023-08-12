package views

import "mvc/pkg/types"

func FileNames() types.FileNames {
	return types.FileNames{
		AcceptIssue:        "accept_issue.html",
		AcceptReturn:       "accept_return.html",
		AddAdmin:           "add_admin.html",
		AddBook:            "add_books.html",
		AdminHome:          "admin_home.html",
		Profile:            "profile.html",
		BooksInventory:     "books_inventory.html",
		BrowseBooks:        "browse_books.html",
		UserHome:           "user_home.html",
		UserIssueRequests:  "user_issue_request.html",
		UserReturnRequests: "user_return_request.html",
		ViewAdmins:         "view_admins.html",
		AdminServerError:   "admin_server_error.html",
		ClientServerError:  "client_server_error.html",
	}
}
