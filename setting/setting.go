package setting

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/config"
	"github.com/ngaut/log"
)

var (
	//Global Config
	AppName string
	Usage   string
	Version string
	Author  string
	Email   string

	//Runtime Config
	RunMode        string
	ListenMode     string
	HttpsCertFile  string
	HttpsKeyFile   string
	LogPath        string
	LogLevel       string
	DatabaseDriver string
	DatabaseURI    string
)

//
func init() {
	if err := setConfig("conf/containerops.conf"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//set log level
	log.SetLevelByString(LogLevel)
}

func setConfig(path string) error {
	conf, err := config.NewConfig("ini", path)
	if err != nil {
		return fmt.Errorf("Read %s error: %v", path, err.Error())
	}

	//config globals
	if appname := conf.String("appname"); appname != "" {
		AppName = appname
	} else if appname == "" {
		return fmt.Errorf("AppName config value is null")
	}

	if usage := conf.String("usage"); usage != "" {
		Usage = usage
	} else if usage == "" {
		return fmt.Errorf("Usage config value is null")
	}

	if version := conf.String("version"); version != "" {
		Version = version
	} else if version == "" {
		return fmt.Errorf("Version config value is null")
	}

	if author := conf.String("author"); author != "" {
		Author = author
	} else if author == "" {
		return fmt.Errorf("Author config value is null")
	}

	if email := conf.String("email"); email != "" {
		Email = email
	} else if email == "" {
		return fmt.Errorf("Email config value is null")
	}

	//config runtime
	if runmode := conf.String("runmode"); runmode != "" {
		RunMode = runmode
	} else if runmode == "" {
		return fmt.Errorf("RunMode config value is null")
	}

	if listenmode := conf.String("listenmode"); listenmode != "" {
		ListenMode = listenmode
	} else if listenmode == "" {
		return fmt.Errorf("ListenMode config value is null")
	}

	if httpscertfile := conf.String("httpscertfile"); httpscertfile != "" {
		HttpsCertFile = httpscertfile
	} else if httpscertfile == "" {
		return fmt.Errorf("HttpsCertFile config value is null")
	}

	if httpskeyfile := conf.String("httpskeyfile"); httpskeyfile != "" {
		HttpsKeyFile = httpskeyfile
	} else if httpskeyfile == "" {
		return fmt.Errorf("HttpsKeyFile config value is null")
	}

	if logpath := conf.String("log::filepath"); logpath != "" {
		LogPath = logpath
	} else if logpath == "" {
		return fmt.Errorf("LogPath config value is null")
	}

	if loglevel := conf.String("log::level"); loglevel != "" {
		LogLevel = loglevel
	} else if loglevel == "" {
		return fmt.Errorf("LogLevel config value is null")
	}

	if databasedriver := conf.String("database::driver"); databasedriver != "" {
		DatabaseDriver = databasedriver
	} else if databasedriver == "" {
		return fmt.Errorf("Database Driver config value is null")
	}

	if databaseuri := conf.String("database::uri"); databaseuri != "" {
		DatabaseURI = databaseuri
	} else if databaseuri == "" {
		return fmt.Errorf("Database URI config vaule is null")
	}

	return nil
}
