package taskrunner 

const (
 	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE = "c"

	VIDEO_PATH = "./videos/"
)

//Control channel
type controlChan chan string

//Data channel
//User interface{} for generic type
type dataChan chan interface{}


type fn func(dc dataChan) error