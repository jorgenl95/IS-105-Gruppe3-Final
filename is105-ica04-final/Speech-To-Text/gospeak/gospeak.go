package gospeak

import (speech "github.com/go-speak")
//example


func main(){
  speech.SetWitKey("RLD2DHGWIRNX72FHOZMRNBTGJBGOCEXK") //Wit API Key MUST be set before calling any other Wit.AI functions
  speech.SendWitVoice("audio-file.wav")
  

}


