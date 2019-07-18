package main

import (
	"fmt"
	"github.com/francoispqt/gojay"
	"math/rand"
	"time"
)

type user struct {
	Id    int
	Name  string
	Email string
}

// implement MarshalerJSONObject
func (u *user) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("id", u.Id)
	enc.StringKey("name", u.Name)
	enc.StringKey("email", u.Email)
}
func (u *user) IsNil() bool {
	return u == nil
}

func (u *user) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.Int(&u.Id)
	case "name":
		return dec.String(&u.Name)
	case "email":
		return dec.String(&u.Email)
	}
	return nil
}
func (u *user) NKeys() int {
	return 3
}

//func main() {
//	u := &user{1, "gojay", "gojay@email.com"}
//	b := strings.Builder{}
//	enc := gojay.NewEncoder(&b)
//	if err := enc.Encode(u); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(b.String())
//}

//func main() {
//	t1 := time.Now().UnixNano()
//	for i := 0; i < 100; i++ {
//		u := &user{1, "gojay", "gojay@email.com"}
//		//b := strings.Builder{}                       // 不推荐这种方式，效率没有下面高
//		//enc := gojay.NewEncoder(&b)
//		//if err := enc.Encode(u); err != nil {
//		//	log.Fatal(err)
//		//}
//		//fmt.Println(b.String())
//		b, err := gojay.MarshalJSONObject(u)
//		if err != nil {
//			log.Fatal(err)
//		}
//		//fmt.Println(string(b))
//		fmt.Println(reflect.TypeOf(string(b)))
//	}
//	t2 := time.Now().UnixNano()
//	fmt.Println(t2 - t1)
//}

//func main() {
//	t1 := time.Now().UnixNano()
//	for i := 0; i < 100000; i++{
//   u := &user{1, "gojay", "gojay@email.com"}
//   buf, err := json.Marshal(u)
//	if err != nil {
//		fmt.Println("err = ", err)
//		fmt.Println("buf = ", string(buf))
//		return
//	}
//	}
//	t2 := time.Now().UnixNano()
//	fmt.Println(t2 - t1)
//}

//func main() {
//	t1 := time.Now().UnixNano()
//	s := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
//	for i := 0; i < 1000; i++ {
//		u := &user{}
//
//		//if err := json.Unmarshal(s, &u); err != nil {
//		//	log.Fatalf("Json unmarshaling failed:%s", err)
//		//}
//		//fmt.Println(u)
//		err := gojay.UnmarshalJSONObject(s, u)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(reflect.TypeOf(u))
//	}
//	t2 := time.Now().UnixNano()
//	fmt.Println(t2 - t1)
//}


func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        b := r.Intn(26) + 65
        bytes[i] = byte(b)
    }
    return string(bytes)
}

func main()  {
	s := RandString(1)
	fmt.Println(s)
}