package main

import "fmt"

func main() {
	// var in *bufio.Reader
	// var out *bufio.Writer
	// in = bufio.NewReader(os.Stdin)
	// out = bufio.NewWriter(os.Stdout)
	// defer out.Flush()
	// var t string
	// t, _ = in.ReadString('\n')
	// tINT, _ := strconv.Atoi(t[:len(t)-1])
	// data := ""
	//
	//	for i := 0; i < tINT; i++ {
	//		n, _ := in.ReadString('\n')
	//		nINT, _ := strconv.Atoi(n[:len(n)-1])
	//		data = ""
	//		for j := 0; j < nINT; j++ {
	//			str, _ := in.ReadString('\n')
	//			data += str
	//		}
	//		result := ReadData([]byte(data))
	//		fmt.Fprintf(out, "%d\n", virusFile(result, 0, false))
	//	}
	result := ReadJson("example.json")
	fmt.Println(virusFile(result, 0, false))
}
