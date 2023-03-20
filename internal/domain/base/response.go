package base

import "github.com/google/uuid"

type Blame string

const (
	BlameUser     Blame = "User"
	BlameKafka    Blame = "Kafka"
	BlamePostgres Blame = "Postgres"
	BlameS3       Blame = "S3"
	BlameServer   Blame = "Server"
	BlameUnknown  Blame = "Unknown"
)

// ResponseOK is a base OK response from server.
type ResponseOK struct {
	Status     string `json:"status" example:"OK"`
	TrackingID string `json:"trackingID" example:"12345678-1234-1234-1234-000000000000"`
}

// ResponseOKWithID is a base OK response from server with additional ID in answer.
type ResponseOKWithID struct {
	Status     string    `json:"status" example:"OK"`
	TrackingID string    `json:"trackingID" example:"12345678-1234-1234-1234-000000000000"`
	ID         uuid.UUID `json:"ID" example:"12345678-1234-1234-1234-000000000000"`
}

// ResponseFailure is a general error response from server.
type ResponseFailure struct {
	Status     string `json:"status" example:"Error"`
	Blame      Blame  `json:"blame" example:"Guilty System"`
	TrackingID string `json:"trackingID" example:"12345678-1234-1234-1234-000000000000"`
	Message    string `json:"message" example:"error occurred"`
}

type ResponseOKWithJWT struct {
	Status     string `json:"status" example:"OK"`
	TrackingID string `json:"trackingID" example:"12345678-1234-1234-1234-000000000000"`
	JWT        string `json:"JWT"`
}
