// Code generated by go-swagger; DO NOT EDIT.

/*
Package restapi uisvc
API for the user interface service; the service that is directly responsible for presenting data to users.
This service typically runs at the border, and leverages session cookies or authentication tokens that we generate for users. It also is responsible for handling the act of oauth and user creation through its login hooks.
uisvc typically talks to the datasvc and other services to accomplish its goal, it does not save anything locally or carry state.



    Schemes:
      http
    Host: localhost
    BasePath: /
    Version: 1.0.0

    Consumes:
    - application/json

    Produces:
    - application/json

swagger:meta
*/
package restapi
