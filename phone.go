package message
import "errors"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"
import "os"

// TextNow Configs
var TextApiurl string = ternary( os.Getenv( "AKONA_TEXTNOW_APIURL" ) != "" , os.Getenv( "AKONA_TEXTNOW_APIURL" ) ,
	"https://www.textnow.com/api/users/mdd55svnlr3pqkeq5iq9/messages" ).( string )
var TextCookie string = ternary( os.Getenv( "AKONA_TEXTNOW_COOKIE" ) != "" , os.Getenv( "AKONA_TEXTNOW_COOKIE" ) ,
	"connect.sid=s:8x1SgmtNw4eNBbA8OQ0mtQqq07DivwAk.LmPINML7UXkusEDqGr/vIvt7qn9/tfcEcvjv/QgWRb8;" ).( string )

// Sends out messages with TextNow. Probably illegal.
func ( self Message )Text( ) error {
	var err error
	var smstext string

	// TextNow doesnt like email addresses
	if strings.Contains( self.Receive , "@" ) {
		return errors.New( "that doesnt look like a phone number" ) }

	// Build message
	smstext = strings.Join( [ ]string{ self.About , self.Content , self.Link } , " " )
	smstext = strings.Replace( smstext , "\\" , "\\\\" , -1 )
	smstext = strings.Replace( smstext , "\"" , "\\\"" , -1 )
	smstext = url.QueryEscape( smstext )
	self.Receive = strings.Replace( self.Receive , "\\" , "\\\\" , -1 )
	self.Receive = strings.Replace( self.Receive , "\"" , "\\\"" , -1 )
	self.Receive = url.QueryEscape( self.Receive )
	smstext = strings.Join( [ ]string{ "json=%7B%22contact_value%22%3A%22" , self.Receive , "%22%2C%22message%22%3A%22" , smstext , "%22%7D" } , "" )

	// Send it
	var client http.Client
	var request http.Request
	request.Method = http.MethodPost
	request.URL = mu( url.Parse( TextApiurl ) )[ 0 ].( * url.URL )
	request.Header = map[ string ][ ]string{
		"Cookie" : { TextCookie } ,
		"Content-Type" : { "application/x-www-form-urlencoded" } , }
	request.Body = ioutil.NopCloser( strings.NewReader( smstext ) )
	request.ContentLength = int64( len( smstext ) )	// Official docs can be misleading
	response , err := client.Do( & request )
	if err != nil {
		return err }
	if response.StatusCode != http.StatusOK {
		return errors.New( "returned unexpected code: " + response.Status ) }

	return nil
}


