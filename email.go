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
	template , err := os.Open( "email.htm" )
	if err != nil {
		return err }
	sendhtml = string( mu( ioutil.ReadAll( template ) )[ 0 ].( [ ]byte ) )
	sendhtml = strings.Replace( strings.Replace( sendhtml , "{{LINK}}" , self.Link , -1 ) , "{{CONTENT}}" , self.Content , -1 )

	// Send
	from := mail.NewEmail( "akona" , "noreply@akona.me" )
	to := mail.NewEmail( "" , self.Receive )
	email := mail.NewSingleEmail( from , self.About , to , sendtext , sendhtml )
	client := sendgrid.NewSendClient( secret )
	_ , err = client.Send( email )

	return err
}


