package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/crvini/gambituser/awsgo"
	"github.com/crvini/gambituser/bd"
	"github.com/crvini/gambituser/models"
)

func main() {
	lambda.Start(EjecutoLamba)
}

func EjecutoLamba(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	if !ValidoParamentos() {
		fmt.Println("Error en los parámetros. debe enviar 'SecretName'")
		err := errors.New("error en los parámetros debe enviar SecretName")
		return event, err
	}
	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer secret " + err.Error())
		return event, err
	}
	err = bd.SignUp(datos)
	return event, err
}

func ValidoParamentos() bool {
	var traerParametro bool
	_, traerParametro = os.LookupEnv("SecretName")
	return traerParametro
}
