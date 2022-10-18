package global

// UseLog function show how to use global log
func UseLog() error{
	if err := Init(); err != nil {
		return err
	}

	// if this is in other package, this value can get through global.WithField and global.Debug
	WithField("key","value").Debug("hello")
	Debug("test")

	return nil
}