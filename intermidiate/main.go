package main
//writting files
import ("fmt"
        "os")

func main(){
     file,err :=os.Create("output.txt")
	 if err !=nil{
		fmt.Println(err)
		return
	 }
	 fmt.Println("File created")

	 defer file.Close()

	 data :=[]byte("HIi it's me")
	 n,err :=file.Write(data)
	 if err !=nil{
		fmt.Println(err)
	 }
	 fmt.Println("data has been written",n)

	 //appending directly string insead of creating byte type
	 file ,err =os.Create("String.txt")
	 if err!=nil{
		fmt.Println("error creating file",err)
		return
	 }
	 defer file.Close()

	 fmt.Println("File created successfully")

	 m,err :=file.WriteString("shut the fuck off")
	 if err!=nil{
		fmt.Println("Error appending file",err)
	 }
	 fmt.Println(m,"data has been appended")

	 
	}
