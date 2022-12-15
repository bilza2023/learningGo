package main
import ("fmt"
	// "github.com/bilza2023/learningGo/util"
		)

	type UserData struct {
		firstName string
		lastName string
		age uint
		isMember bool
		amountDue float32
	}

func main(){
	
fristUser := UserData { "Mike" , "Tyson" , 50 , true , 87.33}
fmt.Println("fristUser.firstName", fristUser.firstName)
fmt.Println("fristUser.lastName", fristUser.lastName)

//--this will not change the original struct (firstUser)
secondUser := createNewUser(fristUser)
 
fmt.Printf("secondUser %v \n",secondUser)

//--this will Change the original struct (firstUser)
createOrignalUser( &fristUser)
 
fmt.Printf("fristUser changed %v \n",fristUser)

}

func createNewUser (userData UserData) UserData {
	userData.firstName = "John"
	userData.amountDue += 100
	return userData
}

func createOrignalUser (userData * UserData) {
	userData.firstName = "Orignal Name changed"
	userData.amountDue = 10000
}