package router

func (r Router) GetTokens() {
	r.App.POST("/api/tokens", r.Controller.AuthTokens)
}

func (r Router) DeleteTokens() {
	r.App.DELETE("/api/tokens/:token", r.Controller.DeleteTokens)
}
