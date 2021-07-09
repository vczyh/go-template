package config

type Source interface {
	//Get(key string) interface{}
	getString(key string) string
	getInt(key string) int
	//GetFloat64(key string) float64
	//GetBool(key string) bool
	//GetTime(key string) time.Time
	//GetDuration(key string) time.Duration
	//GetStringSlice(key string) []string
	//GetIntSlice(key string) []int
}
