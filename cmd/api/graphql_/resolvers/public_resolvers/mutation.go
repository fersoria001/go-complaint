package public_resolvers

import (
	"fmt"
	"go-complaint/application/commands"

	"github.com/graphql-go/graphql"
)

func ContactResolver(p graphql.ResolveParams) (interface{}, error) {
	from := p.Args["email"].(string)
	text := p.Args["text"].(string)
	body := fmt.Sprintf("%s has contact you from Go-Complaint contact page: \n %s", from, text)
	command := commands.SendEmailCommand{
		ToEmail: "bercho001@gmail.com",
		ToName:  "Fernando Agustín Soria",
		Text:    body,
	}
	err := command.Contact(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}
