package core

import (
	"fmt"
	"time"

	"github.com/absinsekt/pnk/lib/strings"
	"github.com/gorilla/securecookie"
)

// Config singletone
var Config *configuration

// InitConfiguration sets Config singletone by ENV variables overriden by map
func InitConfiguration(overrides map[string]interface{}) {
	overrider := func(varName string, fallback interface{}) interface{} {
		result := overrides[varName]

		if result == nil {
			result = GetEnv(varName, fallback)
		}

		return result
	}

	secureAuthKey := overrider(envSecureAuthKey, strings.GenerateRandomString(64)).(string)
	secureEncryptionKey := overrider(envSecureEncryptionKey, strings.GenerateRandomString(32)).(string)
	secondsRarely, _ := time.ParseDuration(fmt.Sprintf("%ds", overrider(envSecondsRarely, 180).(int)))
	secondsFrequently, _ := time.ParseDuration(fmt.Sprintf("%ds", overrider(envSecondsFrequently, 30).(int)))
	tzLocation, _ := time.LoadLocation(overrider(envTimezoneName, "Etc/UTC").(string))

	Config = &configuration{
		ProjectName:      overrider(envProjectName, "PNK").(string),
		BaseURL:          overrider(envBaseURL, "http://localhost/").(string),
		Debug:            overrider(envDebug, false).(bool),
		HostAddress:      overrider(envHost, "127.0.0.1").(string),
		Port:             overrider(envPort, 5000).(int),
		TemplatePath:     overrider(envTemplatePath, "templates").(string),
		MediaPath:        overrider(envMediaPath, "uploads").(string),
		MediaURL:         overrider(envMediaURL, "/media").(string),
		ThumbnailsPrefix: overrider(envThumbnailsPrefix, "thumbs").(string),
		CacheEnabled:     overrider(envCacheEnabled, true).(bool),
		DbHost:           overrider(envDBHost, "127.0.0.1").(string),
		DbName:           overrider(envDBName, "pnk_db").(string),
		DbUser:           overrider(envDBUser, "punk").(string),
		DbPassword:       overrider(envDBPassword, "punksnotdead").(string),
		SMTPHost:         overrider(envSMTPHost, "smtp.gmail.com").(string),
		SMTPPort:         overrider(envSMTPPort, 465).(int),
		SMTPUser:         overrider(envSMTPUser, "google@gmail.com").(string),
		SMTPPassword:     overrider(envSMTPPassword, "punksnotdead").(string),
		SessionVersion:   overrider(envSessionVersion, strings.GenerateRandomString(8)).(string),
		SecureVault: securecookie.New(
			[]byte(secureAuthKey),
			[]byte(secureEncryptionKey),
		),
		SecondsRarely:     secondsRarely,
		SecondsFrequently: secondsFrequently,
		Timezone:          tzLocation,
	}
}

type configuration struct {
	ProjectName       string
	BaseURL           string
	Debug             bool
	HostAddress       string
	Port              int
	TemplatePath      string
	MediaPath         string
	MediaURL          string
	ThumbnailsPrefix  string
	CacheEnabled      bool
	DbHost            string
	DbName            string
	DbUser            string
	DbPassword        string
	SMTPHost          string
	SMTPPort          int
	SMTPUser          string
	SMTPPassword      string
	SessionVersion    string
	SecureVault       *securecookie.SecureCookie
	SecondsRarely     time.Duration
	SecondsFrequently time.Duration
	Timezone          *time.Location
}

const (
	envProjectName         = "PNK_PROJECT_NAME"
	envBaseURL             = "PNK_BASE_URL"
	envDebug               = "PNK_DEBUG"
	envHost                = "PNK_HOST"
	envPort                = "PNK_PORT"
	envTemplatePath        = "PNK_TEMPLATE_PATH"
	envMediaPath           = "PNK_MEDIA_PATH"
	envMediaURL            = "PNK_MEDIA_URL"
	envThumbnailsPrefix    = "PNK_THUMBNAILS_PREFIX"
	envCacheEnabled        = "PNK_CACHE_ENABLED"
	envDBHost              = "PNK_DB_HOST"
	envDBName              = "PNK_DB_NAME"
	envDBUser              = "PNK_DB_USER"
	envDBPassword          = "PNK_DB_PASSWORD"
	envSecureAuthKey       = "PNK_SECURE_AUTH_KEY"
	envSecureEncryptionKey = "PNK_SECURE_ENCRYPTION_KEY"
	envSecondsRarely       = "PNK_SECONDS_RARELY"
	envSecondsFrequently   = "PNK_SECONDS_FREQUENTLY"
	envSMTPHost            = "PNK_SMTP_HOST"
	envSMTPPort            = "PNK_SMTP_PORT"
	envSMTPUser            = "PNK_SMTP_USER"
	envSMTPPassword        = "PNK_SMTP_PASSWORD"
	envSessionVersion      = "PNK_SESSION_VERSION"
	envTimezoneName        = "PNK_TIMEZONE_NAME"
)
