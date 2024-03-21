package main

import (
	"context"
	"fmt"
  "encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
  "github.com/reecerussell/aws-lambda-multipart-parser/parser"
)

type qfullName struct {
  First string `json:"first"`
  Last string `json:"last"`
}
type qPhoneNumber struct {
  Full string `json:"full"`
}

type rawForm struct {
  Slug string `json:"slug"`
  JsExecutionTracker string `json:"jsExecutionTracker"`
  SubmitSource string `json:"submitSource"`
  BuildDate string `json:"buildDate"`
  Q3_fullName3 qfullName `json:"q3_fullName3"`
  Q5_phoneNumber5 qPhoneNumber `json:"q5_phoneNumber5`
  Q12_suggestionsIf string `json:"q12_suggestionsIf"`
  Q31_email string `json:"q31_email"`

}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("This message will show up in the CLI console.")
  
// Parse the request.
	data, err := parser.Parse(request)
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}

  fmt.Printf("formData: %v \n", data)

  phone, ok := data.Get("q5_phoneNumber5")
  if !ok {
    phone = "NOT FOUND"
  }

  var formData rawForm
  rawReq, ok := data.Get("rawRequest")
  json.Unmarshal([]byte(rawReq), &formData)

  // if ok {
  // fmt.Printf("Using GET_Object_RAWREQ phone: %v \n", rawReq.q5_phoneNumber5.full)
  // }

  fmt.Printf("q5_phoneNumber5: %v \n", phone)
  // fmt.Printf("As Object_RAWREQ phone: %v \n", data.rawRequest.q5_phoneNumber5.full)
   fmt.Printf("Using GET_Object_RAWREQ phone: %v \n", formData.Q5_phoneNumber5.Full)


  // fmt.Printf("Request object: %v \n", request)

  
  // fmt.Printf("Request body: %v \n", request.Body)
  // fmt.Printf("Part: %v \n", part)
  // fmt.Printf("Content: %v \n", content)


  // request.ParseMultipartForm(0)
  // fmt.Printf("fullName3: %v", request.FormValue("fullName3"))
  // fmt.Printf("input_5_full: %v", request.FormValue("input_5_full"))
  // fmt.Printf("id = 3 : %v", request.FormValue("3"))
  // use data caught by webhook to send email using SMTP 

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            "Hello, world!",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler, )
}
