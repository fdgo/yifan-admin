package randId

var (
	s *Snowflake
	e error
)

func init() {
	s, e = NewSnowflake(int64(0), int64(0))
	if e != nil {
		return
	}
}
func RandID() uint {
	return uint(s.NextVal())
}
