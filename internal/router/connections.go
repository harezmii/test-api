package router

func (r Router) GetConnections() {
	r.App.GET("/api/session/data/postgresql/connections", r.Controller.GetConnections)
}
