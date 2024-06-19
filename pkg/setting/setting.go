package setting

import (
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Email struct {
	Email    string
	Password string
	Smtp     string
	Port     int64
}

var EmailSetting = &Email{}

var cfg *ini.File

func LoadConfig() error {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		return err
	}
	if err = cfg.Section("app").MapTo(AppSetting); err != nil {
		return err
	}
	if err = cfg.Section("server").MapTo(ServerSetting); err != nil {
		return err
	}
	if err = cfg.Section("database").MapTo(DatabaseSetting); err != nil {
		return err
	}
	if err = cfg.Section("email").MapTo(EmailSetting); err != nil {
		return err
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	return nil
}
