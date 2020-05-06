package configuration

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/absinsekt/pnk/utils/strings"

	"github.com/gorilla/securecookie"
)

var (
	Debug                = getEnv("PNK_DEBUG", false).(bool)
	HostAddress          = getEnv("PNK_HOST", "127.0.0.1").(string)
	Port                 = getEnv("PNK_PORT", 5000).(int)
	TemplatePath         = getEnv("PNK_TEMPLATE_PATH", "templates").(string)
	DbHost               = getEnv("PNK_DB_HOST", "127.0.0.1").(string)
	DbName               = getEnv("PNK_DB_NAME", "pnk_db").(string)
	DbUser               = getEnv("PNK_DB_USER", "punk").(string)
	DbPassword           = getEnv("PNK_DB_PASSWORD", "punksnotdead").(string)
	SMTPHost             = getEnv("PNK_SMTP_HOST", "smtp.gmail.com").(string)
	SMTPPort             = getEnv("PNK_SMTP_PORT", 465).(int)
	SMTPUser             = getEnv("PNK_SMTP_USER", "google@gmail.com").(string)
	SMTPPassword         = getEnv("PNK_SMTP_PASSWORD", "punksnotdead").(string)
	SessionNS            = getEnv("PNK_SESSION_NS", "punksnotdead").(string)
	SessionVersion       = getEnv("PNK_SESSION_VERSION", strings.GenerateRandomString(8)).(string)
	SessionAuthKey       = getEnv("PNK_SESSION_AUTH_KEY", string(securecookie.GenerateRandomKey(64))).(string)
	SessionEncryptionKey = getEnv("PNK_SESSION_ENCRYPTION_KEY", string(securecookie.GenerateRandomKey(32))).(string)

	SecondsRarely, _ = time.ParseDuration(
		fmt.Sprintf("%ds", getEnv("PNK_SECONDS_RARELY", 180).(int)))

	SecondsFrequently, _ = time.ParseDuration(
		fmt.Sprintf("%ds", getEnv("PNK_SECONDS_FREQUENTLY", 30).(int)))
)

func getEnv(key string, fallback interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		switch fallback.(type) {
		case int:
			if v, err := strconv.Atoi(value); err == nil {
				return v
			}
		case bool:
			if v, err := strconv.ParseBool(value); err == nil {
				return v
			}
		default:
			return value
		}
	}

	return fallback
}
