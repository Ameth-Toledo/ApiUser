package dependencies

import (
	"UsersFree/src/core"
	"UsersFree/src/users/application"
	"UsersFree/src/users/infrastructure/adapters"
	"UsersFree/src/users/infrastructure/controllers"
)

func Init(pool *core.Conn_MySQL) (
	*controllers.CreateUserController,
	*controllers.ViewUserController,
	*controllers.EditUserController,
	*controllers.DeleteUserController,
	*controllers.ViewUserByIdController,
	*controllers.AuthController,
	error,
) {

	ps := adapters.NewMySQL(pool.DB)

	createClient := application.NewCreateUser(ps)
	viewClient := application.NewListUser(ps)
	editClient := application.NewEditUser(ps)
	deleteClient := application.NewDeleteUser(ps)
	viewClientById := application.NewUserById(ps)
	authService := application.NewAuthService(ps)

	authController := controllers.NewAuthController(authService)
	createClientController := controllers.NewCreateUserController(*createClient)
	viewClientController := controllers.NewViewUserController(*viewClient)
	editClientController := controllers.NewEditUserController(*editClient)
	deleteClientController := controllers.NewDeleteUserController(*deleteClient)
	viewClientByIdController := controllers.NewViewUserByIdController(*viewClientById)

	return createClientController, viewClientController, editClientController, deleteClientController, viewClientByIdController, authController, nil
}
