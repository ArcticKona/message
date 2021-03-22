// Basic test
// doesnt actually check if message is received
package message
import "testing"

func TestPhone( test * testing.T ) {
	var err error
	var message Message
	message.Receive = "(872) 903-1636"
	message.About = "Test2"
	message.Content = "here is some content"
	message.Link = "https://akona.me"
	err = message.Text( )
	if err != nil {
		test.Fatalf( "%v\r\n" , err ) }
}


