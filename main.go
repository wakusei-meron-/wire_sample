package main

import "wire_sample/env"

func main() {
	e := env.Conf{
		DBSetting: env.DBSetting{
			HOST: "http://example.host",
			USER: "user",
			PASS: "pass",
		},
	}
	c := InitController(e)
	c.Start()
}
