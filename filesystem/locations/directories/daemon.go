package directories

func DaemonDirectory() string {
	return AppDirectory() + "/daemon"
}

func DefinitionsDirectory() string {
	return AppDirectory() + "/definitions"
}

func ReportsDirectory() string {
	return AppDirectory() + "/reports"
}
