package LoginCCNU

import (
	"fmt"
	"github.com/ShiinaOrez/LoginCCNU/spoc"
)

const (
	YOURSNO      = "YOUR_STUDENT_NO"
	YOURPASSWORD = "YOUR_PASSWORD"
)

func trySpoc() {
	response, err := spoc.Login(YOURSNO, YOURPASSWORD)
	if response == nil {
		fmt.Println("Error... !!!")
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(response.Data.RoleDepartment)
	}
}
