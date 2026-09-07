package header

type Definition struct {
	Name   string
	Values []string
}

func CommonHeaders() []Definition {
	return []Definition{
		{
			Name: "Accept",
			Values: []string{
				"application/json",
				"application/xml",
				"text/plain",
				"text/html",
				"*/*",
			},
		},
		{
			Name: "Authorization",
			Values: []string{
				"Bearer",
				"Basic",
			},
		},
		{
			Name: "Content-Type",
			Values: []string{
				"application/json",
				"application/xml",
				"text/plain",
				"application/x-www-form-urlencoded",
				"multipart/form-data",
			},
		},
		{
			Name:   "User-Agent",
			Values: []string{},
		},
	}
}
