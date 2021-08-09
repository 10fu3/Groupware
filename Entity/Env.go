package Entity

type Env struct {
	DB_Host           string `required:"true"`
	DB_Port           int    `required:"true"`
	DB_User           string `required:"true"`
	DB_Password       string `required:"true"`
	DB_Name           string `required:"true"`
	App_Port          int    `required:"true"`
	App_Client_Secret string `required:"true"`
	App_Client_ID     string `required:"true"`
}
