Vim�UnDo� T��8�����c����#�r�f:�6��3��	~                                      UA�    _�                              ����                                                                                                                                                                                                                                                                                                                                                             UA�     �                 package main       import "fmt"   import "time"       func main() {       :	cc := time.Now().Local() //.Format("2006-01-02 15:04:05")   	ca := time.Now().Local()   	fmt.Println(ca.After(cc))   	//fmt.Println(cc)   &	//fmt.Println(cc.Sub(30*time.Second))       }�               �                 package main       import "fmt"   import "time"       func main() {   		   :		cc := time.Now().Local()//.Format("2006-01-02 15:04:05")   			ca := time.Now().Local()   				fmt.Println(ca.After(cc))   					//fmt.Println(cc)   +						//fmt.Println(cc.Sub(30*time.Second))   							   					}�               �                   5��