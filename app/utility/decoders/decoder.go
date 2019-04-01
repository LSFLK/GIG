package decoders

import (
	"GIG/app/models"
	"io"
)

type Decoder interface {
	DecodeSource(reader io.Reader) models.Entity
}