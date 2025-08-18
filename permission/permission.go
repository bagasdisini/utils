package permission

var (
	PERMISSION_SECRET_KEY string = "PERMISSION_SECRET_KEY"
)

const (
	ViewAll    = "view_all"
	ViewDetail = "view_detail"
	ViewSelf   = "view_self"
	Create     = "create"
	CreateBulk = "create_bulk"
	Update     = "update"
	UpdateSelf = "update_self"
	Delete     = "delete"
	DeleteBulk = "delete_bulk"
)

type BulkPayload struct {
	App        string
	Menu       string
	Permission string
}

type Permission struct {
	AppName string  `json:"app_name"`
	Menus   []*Menu `json:"menus"`
}

type Menu struct {
	MenuName   string    `json:"menu_name"`
	Permission []*string `json:"permission"`
}
