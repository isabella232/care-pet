// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/care-pet/go/cmd/server/models"
)

// FindSensorsByPetIDReader is a Reader for the FindSensorsByPetID structure.
type FindSensorsByPetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindSensorsByPetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindSensorsByPetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindSensorsByPetIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindSensorsByPetIDOK creates a FindSensorsByPetIDOK with default headers values
func NewFindSensorsByPetIDOK() *FindSensorsByPetIDOK {
	return &FindSensorsByPetIDOK{}
}

/*FindSensorsByPetIDOK handles this case with default header values.

sensors response
*/
type FindSensorsByPetIDOK struct {
	Payload []*models.Sensor
}

func (o *FindSensorsByPetIDOK) Error() string {
	return fmt.Sprintf("[GET /pet/{id}/sensors][%d] findSensorsByPetIdOK  %+v", 200, o.Payload)
}

func (o *FindSensorsByPetIDOK) GetPayload() []*models.Sensor {
	return o.Payload
}

func (o *FindSensorsByPetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindSensorsByPetIDDefault creates a FindSensorsByPetIDDefault with default headers values
func NewFindSensorsByPetIDDefault(code int) *FindSensorsByPetIDDefault {
	return &FindSensorsByPetIDDefault{
		_statusCode: code,
	}
}

/*FindSensorsByPetIDDefault handles this case with default header values.

error
*/
type FindSensorsByPetIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find sensors by pet id default response
func (o *FindSensorsByPetIDDefault) Code() int {
	return o._statusCode
}

func (o *FindSensorsByPetIDDefault) Error() string {
	return fmt.Sprintf("[GET /pet/{id}/sensors][%d] find sensors by pet id default  %+v", o._statusCode, o.Payload)
}

func (o *FindSensorsByPetIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindSensorsByPetIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}