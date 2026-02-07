package configs

import "fmt"
import "os"

type Configuration struct {
	IsProd bool
}

IsProd := false
if os.Getenv("dev_type") == "prod" {
	IsProd = true
}


func (conf Configuration) InitiateConfiguration () (Configuration, error){
	// this will be the entry point for all the initial configuration that is needed for runtime execution
	
	conf := Configuration{ IsProd : IsProd}

	if conf != nil {
		return conf, Errors.new("The configuration is created successfully")
	} else{
		return conf, Errors.new("Something went wrong the configuration is unsuccessful.")
	}

}
