package main

type Sayer interface {
	Say()
}

func Talk(sayer Sayer){
	sayer.Say()
}

func main(){
	Talk(nil)
}

