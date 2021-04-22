package message
import "github.com/sendgrid/sendgrid-go"
import "github.com/sendgrid/sendgrid-go/helpers/mail"
import "io/ioutil"
import "strings"
import "os"

// Keep this secret
var EmailSecret string = ternary( os.Getenv( "AKONA_SENDGRID_SECRET" ) != "" , os.Getenv( "AKONA_SENDGRID_SECRET" ) ,
	"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" ).( string )

// Sends emails
func ( self Message )Email( ) error {
	var err error
	var sendtext string
	var sendhtml string

	// Operate templates
	sendtext = self.Content + " " + self.Link
//	template , err := os.Open( "email.htm" )
//	if err != nil {
//		return err }
//	sendhtml = string( mu( ioutil.ReadAll( template ) )[ 0 ].( [ ]byte ) )
	sendhtml = "<!DOCTYPE HTML><html><body><div style="color: white; background: black; font-size: 18px; font-family: sans-serif; padding: 48px; text-align: center;"><img src="https://media.akona.me/logo%20email.png" style="max-width: 480px;" alt="logo"><h1 style="font-size: 48px; font-family: Comfortaa, Quicksand, sans-serif; margin-top: 0px;">akona</h1><p>{{CONTENT}}<br/><a href="{{LINK}}" style="text-decoration: none;">{{LINK}}</a></p><p style="color: grey; font-size: 12;">If you did not make this request, account may be vulnerable.</p></div></body></html>"	// Hard code
	sendhtml = strings.Replace( strings.Replace( sendhtml , "{{LINK}}" , self.Link , -1 ) , "{{CONTENT}}" , self.Content , -1 )

	// Send
	from := mail.NewEmail( "akona" , "noreply@akona.me" )
	to := mail.NewEmail( "" , self.Receive )
	email := mail.NewSingleEmail( from , self.About , to , sendtext , sendhtml )
	client := sendgrid.NewSendClient( EmailSecret )
	_ , err = client.Send( email )

	return err
}


