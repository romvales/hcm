package hcmcore

type CoreServiceLogLevel uint

const (
	CoreServiceLog_INFO CoreServiceLogLevel = iota
	CoreServiceLog_WARN
	CoreServiceLog_FATAL
)
