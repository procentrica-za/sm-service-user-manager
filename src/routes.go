package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/user", s.handleregisteruser()).Methods("POST")
	s.router.HandleFunc("/user", s.handleupdateuser()).Methods("PUT")
	s.router.HandleFunc("/user", s.handledeleteuser()).Methods("DELETE")
	s.router.HandleFunc("/userlogin", s.handleloginuser()).Methods("GET")
	s.router.HandleFunc("/user", s.handlegetuser()).Methods("GET")
	s.router.HandleFunc("/forgotpassword", s.handleforgotpassword()).Methods("GET")
	s.router.HandleFunc("/userpassword", s.handleupdatepassword()).Methods("PUT")
	s.router.HandleFunc("/institution", s.handlegetinstitutions()).Methods("GET")

}
