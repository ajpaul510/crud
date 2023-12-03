package initialize

type Iinitialize interface{
	Init(env string) (string, string)
	GetConnectionString() string
	GetPort() string
	SetConnectionString(cs string)
	SetPort(p int) 
}
