package main
import ("fmt")

func main(){
	var user int
	fmt.Print("Enter any five digit to check if it's a pallindrome or not: ")
	fmt.Scanln(&user)

	original := user
	var num int

	for user !=0{
		num = num * 10 + user % 10
		user = user / 10
		}
				if(original == num){
			fmt.Println("It's a pallindrome")
		}else{
			fmt.Println("It's not a pallindrome")
		}
}