package types

type Book struct {
	UserName       string `json:"username"`
	UserID         int    `json:"userId"`
	BookID         int    `json:"bookId"`
	BookName       string `json:"bookName"`
	RequestID      int    `json:"requestId"`
	Publisher      string `json:"publisher"`
	ISBN           string `json:"isbn"`
	Edition        int    `json:"edition"`
	Quantity       int    `json:"quantity"`
	IssuedQuantity int    `json:"issuedQuantity"`
}
type UserRegistration struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ErrorMessage struct {
	Message string `json:"message"`
}
type DBInfo struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
	JWT_KEY     string `yaml:"JWT_SECRETKEY"`
}

type FileNames struct {
	AcceptIssue        string
	AcceptReturn       string
	AddAdmin           string
	AddBook            string
	AdminHome          string
	Profile            string
	BooksInventory     string
	BrowseBooks        string
	UserHome           string
	UserIssueRequests  string
	UserReturnRequests string
	ViewAdmins         string
	AdminServerError   string
	ClientServerError  string
}
