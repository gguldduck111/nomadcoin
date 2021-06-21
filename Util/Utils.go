package Util

import "log"

func HandleErr(err error)  {
	if err != nil {
		log.Panicln(err)
	}
}
