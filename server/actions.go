package server

import (
	"aino-spring.com/aino_site/misc"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetHost(context *gin.Context) string {
	scheme := "http"
	if context.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + context.Request.Host
}

func (server *Server) SendVerificationEmail(host string, email string) error {
	id, err := server.Database.FetchUserByEmail(email)
	if err != nil {
		return err
	}

	verifyKey := misc.GenerateVerificationKey(email, server.Config.VerifySalt)
	verifyLink := host + "/api/users/" + id + "/verify/" + verifyKey + "?redirect=/login"
	verifyMessage, err := server.EmailTemplate.Render("verify", gin.H{"link": verifyLink})
	if err != nil {
		return err
	}

	return server.Emailer.SendMail(email, "Verify your email", verifyMessage)
}

func (server *Server) SendPasswordResetEmail(host string, email string) error {
	id, err := server.Database.FetchUserByEmail(email)
	if err != nil {
		return err
	}

	resetKey := misc.GenerateVerificationKey(email, server.Config.VerifySalt)
	resetLink := "http://" + host + "/api/users/" + id + "/password/reset/" + resetKey + "?redirect=/login"
	resetMessage, err := server.EmailTemplate.Render("reset_password", gin.H{"link": resetLink})
	if err != nil {
		return err
	}

	return server.Emailer.SendMail(email, "Reset your password", resetMessage)
}
