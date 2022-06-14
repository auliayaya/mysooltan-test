package internal

type Service interface {
	// this function will get file error log
	GetLogFile(path string) ([]string, error)
	// this function will convert error log file by type
	ConvertLogFile(path, typ, saveto string) (string, error)
	// this function will save log file to the destination location
	SaveLog(path string, data string, filename string) (string, error)
}
