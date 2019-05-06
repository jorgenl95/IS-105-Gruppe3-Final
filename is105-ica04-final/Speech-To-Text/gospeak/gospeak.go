package gospeak

import (speech "github.com/nicolaifsf/go-speak")
//example


func main(){
  speech.SetWitKey("RLD2DHGWIRNX72FHOZMRNBTGJBGOCEXK") //Wit API key fra sveggen
  speech.SendWitVoice("audio-file.wav")
  

}


