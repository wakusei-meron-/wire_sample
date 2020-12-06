package env

type (
	DBSetting struct {
		HOST string
		USER string
		PASS string
	}

	Conf struct {
		DBSetting DBSetting
	}
)
