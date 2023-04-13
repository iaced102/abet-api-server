package model

type Users struct {
	Id            string   `json:"id" param:"id" db:"id" query:"id" form:"id" pg:"id"   primary:"true"`
	UserType      int      `json:"userType" db:"userType" form:"userType" pg:"user_type,use_zero"`
	UserName      string   `json:"userName" db:"userName" form:"userName" pg:"user_name"`
	CreatedAt     string   `json:"createdAt" db:"createdAt" form:"createdAt" pg:"created_at"`
	UpdatedAt     string   `json:"updatedAt" db:"updatedAt" form:"updatedAt" pg:"updated_at"`
	Password      string   `json:"password" db:"password" form:"password" pg:"password"`
	CryptPassword string   `json:"cryptPassword" db:"cryptPassword" form:"cryptPassword" pg:"crypt_password"`
	tableName     struct{} `pg:"users"`
}