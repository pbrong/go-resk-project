package test

import (
	"fmt"
	"github.com/json-iterator/go"
)

func main() {
	jsonTest()
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := intersection(nums1, nums2)
	for _, v := range res {
		print(v)
	}
}

type Person struct {
	Name     string `json:"username"`
	Age      int    `json:"age"`
	Describe string `json:"describe"`
	//Describe string `json:"-"`   // - 表示不进行json化
	sex string `json:"sex"` // 不可见字段不进行json化
}

func jsonTest() {
	// 序列化
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	p := &Person{
		Name:     "ARong",
		Age:      1,
		Describe: "Crud Boy",
		sex:      "男",
	}
	data, err := json.Marshal(p)
	if err != nil {
		println(err)
	}
	fmt.Println(string(data))
	// 反序列化
	p2 := &Person{}
	err = json.Unmarshal(data, p2)
	if err != nil {
		println(err)
	}
	fmt.Println(p2)
}

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	set := map[int]bool{}
	n1 := len(nums1)
	if n1 == 0 {
		return res
	}
	for _, v := range nums1 {
		set[v] = true
	}
	for _, v := range nums2 {
		if flag, _ := set[v]; flag {
			set[v] = false
			res = append(res, v)
		}
	}
	return res
}
