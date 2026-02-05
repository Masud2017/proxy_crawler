package utils

import "errors"
import "github.com/Masud2017/proxy_crawler/models"

func parseArgs(args []string) ([]models.CommandLineArgument,error) {
	var commandLineArgumentList = make([]models.CommandLineArgument,len(len(args)))

	for idx,arg := range args {
		commandLineArgumentList[idx] = arg	
	}

	if commandLineArgumentList != nil {
		return commandLineArgumentList,errors.New("No error found")
	} else {
		return commandLineArgumentList,errors.New("Something went wrong the command line arguments can not be parsed properly")
	}
}


