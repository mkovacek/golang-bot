package main

import(
	"log"
	"net/http"
	"gopkg.in/maciekmm/messenger-platform-go-sdk.v4"
	"fmt"
)
func main(){
	mess.MessageReceived = MessageReceived
	http.HandleFunc("/",index)
	http.HandleFunc("/webhook", mess.Handler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Worldd"))
}

var mess = &messenger.Messenger{
	AccessToken: "ACCESS_TOKEN",
}


func MessageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
	profile, err := mess.GetProfile(opts.Sender.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("Hello, %s %s, %s", profile.FirstName, profile.LastName, msg.Text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}






