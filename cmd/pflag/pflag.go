package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper1()
}

func pflag1() {
	name := pflag.StringP("name", "n", "liuyibao", "Input your name")
	age := pflag.IntP("age", "a", 22, "Input your age")

	pflag.Parse()

	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
}

func pflag2()  {
	var version bool
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.BoolVar(&version, "version", true, "Print version information and quit")

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(version)

}

func pflag3() {
	var name = pflag.String("name", "colin", "Input your name")
	pflag.Parse()

	fmt.Println(*name)
}

func pflag4() {
	var name = pflag.StringP("name", "n", "colin", "Input your name")
	pflag.Parse()

	fmt.Println(*name)
}

type OssConfig struct {
	Port int
	Name string
	Password string
}

func viper1()  {
	cfg := pflag.StringP("config", "c", "", "Configuration file.")
	help := pflag.BoolP("help", "h", false, "Show this help message.")

	pflag.Parse()

	if *help {
		pflag.Usage()
		return
	}

	if *cfg != "" {
		viper.SetConfigFile(*cfg)
		viper.SetConfigType("yaml")
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.iam")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//viper.SetDefault()
	//viper.AutomaticEnv()
	//viper.BindEnv()
	fmt.Println(viper.GetString("db_address"))

	var ossConfig OssConfig
	err := viper.UnmarshalKey("oss_config", &ossConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ossConfig)

	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
}
