package main

func (s *Server) routes() {
	s.router.HandleFunc("/user", s.handleregisteruser()).Methods("POST")
	s.router.HandleFunc("/user", s.handleupdateuser()).Methods("PUT")
	s.router.HandleFunc("/user", s.handledeleteuser()).Methods("DELETE")
	s.router.HandleFunc("/userlogin", s.handleloginuser()).Methods("GET")
	s.router.HandleFunc("/user", s.handlegetuser()).Methods("GET")
	s.router.HandleFunc("/forgetpassword", s.handleforgetpassword()).Methods("POST")
	s.router.HandleFunc("/userlogout", s.handlelogoutuser()).Methods("POST")
}
