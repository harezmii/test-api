package router

func (r Router) GetUsers() {
	r.App.GET("/api/session/data/:source/users", r.Controller.Users)
}
