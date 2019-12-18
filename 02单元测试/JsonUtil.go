package mymath

import (
	"encoding/json"
	"fmt"
	"os"
	"errors"
)

type Person struct {
	Name string
	Age int
	Rmb float64
	Sex bool
	Hobby []string
}

func EncodePerson2JsonFile(p *Person,filename string) bool {
	dstFile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer dstFile.Close()

	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(p)
	if err != nil{
		fmt.Println("编码失败，err=",err)
		return false
	}else{
		fmt.Println("编码成功!")
		return true
	}
}

func DecodeJsonFile2Person(filename string) (*Person,error) {
	srcFile, _ := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer srcFile.Close()
	personPtr := new(Person)

	decoder := json.NewDecoder(srcFile)
	err := decoder.Decode(personPtr)
	if err != nil{
		fmt.Println("解码失败，err=",err)
		return nil,errors.New("解码失败")
	}else{
		fmt.Println("解码成功!")
		return personPtr,nil
	}
}
