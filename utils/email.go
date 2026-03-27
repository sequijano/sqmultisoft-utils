package utils

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
)

type SMTPConfig struct {
	Host string
	Port string
	User string
	Pass string
	From string
}

func SendEmail(cfg SMTPConfig, to, subject, body string) error {
	if cfg.User == "" || cfg.Pass == "" {
		return fmt.Errorf("SMTP no configurado")
	}
	msg := []byte(
		"From: " + cfg.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n" +
			"\r\n" +
			body + "\r\n",
	)
	addr := cfg.Host + ":" + cfg.Port
	auth := smtp.PlainAuth("", cfg.User, cfg.Pass, cfg.Host)

	if cfg.Port == "465" {
		tlsCfg := &tls.Config{ServerName: cfg.Host}
		conn, err := tls.Dial("tcp", addr, tlsCfg)
		if err != nil {
			return fmt.Errorf("tls dial: %w", err)
		}
		defer conn.Close()
		client, err := smtp.NewClient(conn, cfg.Host)
		if err != nil {
			return fmt.Errorf("smtp client: %w", err)
		}
		defer client.Close()
		if err = client.Auth(auth); err != nil {
			return fmt.Errorf("smtp auth: %w", err)
		}
		if err = client.Mail(cfg.From); err != nil {
			return err
		}
		if err = client.Rcpt(to); err != nil {
			return err
		}
		w, err := client.Data()
		if err != nil {
			return err
		}
		if _, err = w.Write(msg); err != nil {
			return err
		}
		return w.Close()
	}

	host, _, _ := net.SplitHostPort(addr)
	_ = host
	return smtp.SendMail(addr, auth, cfg.From, []string{to}, msg)
}

func TenantWelcomeEmail(tenantName, tenantSlug, username, password string) (subject, body string) {
	subject = "Bienvenido a Multi-Currency POS - Credenciales de acceso"
	body = fmt.Sprintf(`Hola,

Tu cuenta en Multi-Currency POS ha sido creada exitosamente.

Empresa: %s
Código de empresa: %s
Usuario: %s
Contraseña temporal: %s

Por seguridad, el sistema te pedirá cambiar tu contraseña en el primer inicio de sesión.

Saludos,
Equipo Multi-Currency POS
`, tenantName, tenantSlug, username, password)
	return
}

func SuperAdminOTPEmail(code string) (subject, body string) {
	subject = "Multi-Currency POS - Código de verificación"
	body = fmt.Sprintf(`Hola,

Se ha solicitado acceso al panel de administración de Multi-Currency POS.

Tu código de verificación es: %s

Este código expira en 5 minutos. Si no solicitaste este acceso, ignora este mensaje.

Saludos,
Equipo Multi-Currency POS
`, code)
	return
}

func PasswordResetEmail(tenantName, tenantSlug, username, password string) (subject, body string) {
	subject = "Multi-Currency POS - Nueva contraseña asignada"
	body = fmt.Sprintf(`Hola,

Se ha generado una nueva contraseña para tu cuenta en Multi-Currency POS.

Empresa: %s
Código de empresa: %s
Usuario: %s
Nueva contraseña temporal: %s

Por seguridad, el sistema te pedirá cambiar tu contraseña en el próximo inicio de sesión.

Saludos,
Equipo Multi-Currency POS
`, tenantName, tenantSlug, username, password)
	return
}
