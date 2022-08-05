package main

import "fmt"

type Profile struct {
	Name    string
	Age     int
	Married bool
}

func main() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王麻子", Age: 21},
	}
	// Hash 方法
	buildHashIndex(list)
	queryHashData("张三", 30)

	// map 方法
	buildMapIndex(list)
	queryMapData("张三", 30)
}

// 传统 Hash 方法

func simpleHash(str string) (ret int) {
	for i := 0; i < len(str); i++ {
		c := str[i]
		ret += int(c)
	}
	return
}

type classQueryKey struct {
	Name string
	Age  int
}

func (c *classQueryKey) hash() int {
	return simpleHash(c.Name) + c.Age + 100000
}

var hashMapper = make(map[int][]*Profile)

func buildHashIndex(list []*Profile) {
	for _, profile := range list {
		key := classQueryKey{profile.Name, profile.Age}
		exitValue := hashMapper[key.hash()]
		exitValue = append(exitValue, profile)
		hashMapper[key.hash()] = exitValue
	}
}

func queryHashData(name string, age int) {
	keyToQuery := classQueryKey{name, age}
	resultList := hashMapper[keyToQuery.hash()]
	for _, result := range resultList {
		if result.Name == name && result.Age == age {
			fmt.Println(result)
			return
		}
	}
	fmt.Println("no found")
}

// map 方法
type mapQueryKey struct {
	Name string
	Age  int
}

var mapMapper = make(map[mapQueryKey]*Profile)

func buildMapIndex(list []*Profile) {
	for _, profile := range list {
		key := mapQueryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}
		mapMapper[key] = profile
	}
}

func queryMapData(name string, age int) {
	key := mapQueryKey{name, age}
	result, ok := mapMapper[key]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("no found")
	}
}
