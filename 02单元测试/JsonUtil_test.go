package mymath

import "testing"

func TestEncodePerson2JsonFile(t *testing.T) {
	filename := "D:/BJBlockChain1803/demos/W2/day4/file/人员.json"
	p := Person{"于谦",50,123.45,true,[]string{"抽烟","喝酒","烫头"}}

	ok := EncodePerson2JsonFile(&p, filename)
	if ok {
		pBack, _ := DecodeJsonFile2Person(filename)
		if pBack.Name == p.Name && pBack.Age == p.Age && pBack.Rmb == p.Rmb && pBack.Sex==p.Sex {
			t.Log("EncodePerson2JsonFile 测试成功！")
		}else{
			t.Fatal("编码前后数据不一致！")
		}
	}else{
		t.Fatal("编码失败！")
	}
}

func TestDecodeJsonFile2Person(t *testing.T) {
	filename := "D:/BJBlockChain1803/demos/W2/day4/file/人员2.json"
	p := Person{"于谦",50,123.45,true,[]string{"抽烟","喝酒","烫头"}}

	ok := EncodePerson2JsonFile(&p, filename)
	if ok {
		pBack, _ := DecodeJsonFile2Person(filename)
		if pBack.Name == p.Name && pBack.Age == p.Age && pBack.Rmb == p.Rmb && pBack.Sex==p.Sex {
			t.Log("EncodePerson2JsonFile 测试成功！")
		}else{
			t.Fatal("编码前后数据不一致！")
		}
	}else{
		t.Fatal("编码失败！")
	}
}