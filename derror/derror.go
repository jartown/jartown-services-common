package derror

import "github.com/singkorn/jartown-services-common/log"

type Derror struct {
	Err        error  `json:"-"`
	ErrStr     string `json:"error"`
	SourceFunc string `json:"source_func"`
	SourceLine int    `json:"source_line"`
	Debug      string `json:"debug_message"`
}

func (e Derror) Error() string {
	return e.Err.Error()
}

func ErrorDebug(err error, debug string) Derror {
	derror, ok := err.(Derror)
	if !ok {
		file, line := log.GetCaller()
		derror = Derror{Err: err, ErrStr: err.Error(), SourceFunc: file, SourceLine: line, Debug: debug}
		log.Error(err, debug)
	}
	return derror
}

func Error(err error) Derror {
	return ErrorDebug(err, "")
}
