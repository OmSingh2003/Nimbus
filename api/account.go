package api 

import (

)
type createAccountRequest struct {

	Owner    string `json:"owner" binding:"required"` 
	Currency string `json:"currency" binding : "required,oneof = INR USD EUR"`
}
func (server *Server) createAccount(ctx gin.Context) {
	var req createAccountRequest
	if err :=ctx.ShouldBindJSON(&req) , err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return 
	}
}
