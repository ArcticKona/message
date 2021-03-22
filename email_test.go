// Basic test
package message
import "testing"

func TestEmail( test * testing.T ) {
	var err error
	var message Message
	message.Receive = "arcticjieer@gmail.com"
	message.About = "Test2"
	message.Content = "here is some content"
	message.Link = "https://akona.me"
	err = message.Email( )
	if err != nil {
		test.Fatalf( "%v\r\n" , err ) }
}


