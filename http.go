package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{&http.Server{
		Addr:    addr,
		Handler: handler,
		// https://blog.cloudflare.com/exposing-go-on-the-internet/
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig: &tls.Config{
			NextProtos:       []string{"h2", "http/1.1"},
			MinVersion:       tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{tls.CurveP256, tls.X25519},
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
			PreferServerCipherSuites: true,
		},
	}}
}

func NewServerTLS(addr string, cert tls.Certificate, handler http.Handler) *Server {
	srv := NewServer(addr, handler)
	srv.TLSConfig.Certificates = []tls.Certificate{cert}

	return srv
}

type Server struct {
	*http.Server
}

func (srv *Server) IsTLS() bool {
	return len(srv.TLSConfig.Certificates) > 0 || srv.TLSConfig.GetCertificate != nil
}

func (srv *Server) Listen() (net.Listener, error) {
	return net.Listen("tcp", srv.Addr)
}

func (srv *Server) ListenAndServe() error {
	ln, err := srv.Listen()
	if err != nil {
		return err
	}

	return srv.Serve(ln)
}

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
	ln, err := srv.Listen()
	if err != nil {
		return err
	}

	return srv.ServeTLS(ln, certFile, keyFile)
}

func (srv *Server) Start() error {
	ln, err := srv.Listen()
	if err != nil {
		return err
	}

	if srv.IsTLS() {
		ln = tls.NewListener(ln, srv.TLSConfig)
	}

	return srv.Serve(ln)
}

func HttpExample()  {
	addr := os.Getenv("HTTP_LISTEN")
	if addr == "" {
		addr = ":8080"
	}

	// use something like chi here
	r := http.NewServeMux()

	// certFile := "" // todo
	// keyFile := "" // todo
	// cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	// if err != nil {
	// 	log.Fatalf("error loading certificate: %v", err)
	// }
	// server := NewServerTLS(addr, cert, r)

	server := NewServer(addr, r)

	serveErr := server.Start()
	if serveErr != nil {
		log.Printf("error starting http server: %v", serveErr)
	}
}
