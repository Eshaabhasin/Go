package main
import "fmt"
func main(){
	var ages [3] int=[3] int{9,7,8}
	var age=[2] int{8,7}
	names:=[2] string{"eshaa","akanksha"}
	fmt.Println(ages)
	fmt.Println(age)
	fmt.Println(ages,len(ages))
	fmt.Println(age,len(age))
	fmt.Println(names,len(names))

	//slices
	var scores=[]int{8,9,11}
	scores[2]=10
	scores=append(scores, 11)
	fmt.Println(scores,len(scores))

	rangeone:=names[0:2]
	fmt.Println(rangeone)
	rangeone=append(rangeone,"rahul")
	fmt.Println(rangeone)

}