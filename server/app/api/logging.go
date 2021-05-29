package api

const callingAPIFormat = "Calling API: %s\n"

func callingAPI(named string, moreMessages ...interface{}) {
	logger.Printf(callingAPIFormat, named)
	for _, msg := range moreMessages {
		logger.Println(msg)
	}
}
