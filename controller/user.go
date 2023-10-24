package controller

import (
	"github.com/gin-gonic/gin"
)

// type errMsg struct {
// 	Message string `json:"message,omitempty"`
// }

// //////////////////////////////////////////////////////////////////////////////////////
// var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "Users")

// //////////////////////////////////////////////////////////////////////////////////////
// var validate = validator.New()

// ! Always keep export entities capitalized
func Login(context *gin.Context) {
	// 	var user model.User
	// 	if err := context.BindJSON(&user); err != nil {
	// 		var message = errMsg{Message: "Invalid JSON data"}
	// 		context.JSON(http.StatusBadRequest, message)
	// 		return
	// 	}

	// 	if user.Email == "" || user.Password == "" {
	// 		var message = errMsg{Message: "Please provide with sufficient credentials"}
	// 		context.JSON(http.StatusBadRequest, message)
	// 		return
	// 	}

	// 	context.JSON(http.StatusCreated, user)
}

func Register(r *gin.Context) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	var user model.User
	// 	defer cancel()

	// 	if err := r.BindJSON(&user); err != nil {
	// 		r.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	// 		return
	// 	}

	// 	//* Validating the presence of all the required fields
	// 	if validationErr := validate.Struct(&user); validationErr != nil {
	// 		r.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "Please provide with sufficient credentials", Data: map[string]interface{}{"data": validationErr.Error()}})
	// 		return
	// 	}

	// 	///////////////////////////////////////////////////
	// queries := db.New(conn)

	// iUser, err := queries.CreateUser(ctx, db.CreateUserParams{
	// 	Name:     "Manas",
	// 	Email:    "manas@example.com",
	// 	Password: "password",
	// })

	// fmt.Println(err)
	// if err != nil {
	// 	return err
	// }
	// log.Println(iUser)
	// 	////////////////////////////////////////////////////
	// 	if err != nil {
	// 		r.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	// 		return
	// 	}

	// r.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
