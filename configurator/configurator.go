package configurator

var c *Confor

func init() {
	c = New()
}

// Confor is a configuration, it maintains a set of configuration sources, fetches values
type Confor struct {
	confileName string // the name of config file
	confile     string // the config file, include the path, name and extension
	confType    string // the type of configuration. support: yml, yaml, json, ini

	kvstore map[string]interface{}
}

// New returns an initialized confor instance
func New() *Confor {
	c := new(Confor)
	c.confType = "json"
	c.kvstore = make(map[string]interface{})

	return c
}

// SetConfigFile set the explicitly config file with the path, name and extension
func SetConfigFile(confile string) {
	c.SetConfigFile(confile)
}

func (c *Confor) SetConfigFile(confile string) {
	if confile != "" {
		c.confile = confile
		// TODO: add info log
	}
}

// ReadConfig will load the configuration file form disk, and key/value stores
func ReadConfig() error { return c.ReadConfig() }

func (c *Confor) ReadConfig() error {

	return nil
}
