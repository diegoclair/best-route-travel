package domain

import "time"

//DefaultConnectionTimeout is the timeout for http requests
const DefaultConnectionTimeout time.Duration = 1 * time.Minute

//constants to interage with safra api
const (
	TypePrescriptionFile string = "prescription"
)
