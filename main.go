package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/rs/xid"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

type RSA struct {
	RSA_Ver     int
	RSA_Public  []byte
	RSA_Private []byte
}

type User struct {
	userID   int
	username string
	password string
	nickname string
	userGUID xid.ID
	birthday time.Time
	token    xid.ID
	rsa      RSA
}

func (obj *User) Encode() ([]byte, error) {
	var testStruct = obj
	Len := unsafe.Sizeof(*testStruct)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(testStruct)),
		cap:  int(Len),
		len:  int(Len),
	}
	bin := *(*[]byte)(unsafe.Pointer(testBytes))

	return bin, nil
}

func (u *User) Decode(bin []byte) (*User, error) {
	var obj *User = *(**User)(unsafe.Pointer(&bin))
	return obj, nil
}

var user = User{
	userID:   2323,
	username: "armstrong",
	password: "password",
	nickname: "SuperNova",
	userGUID: xid.New(),
	birthday: time.Now(),

	rsa: RSA{
		RSA_Ver:    999,
		RSA_Public: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

func main() {
	user.token = xid.New()
	user.rsa.RSA_Private = []byte{11, 22, 33, 44, 55}

	fmt.Println(user)
	bin, err := user.Encode()

	fmt.Println("\n[]byte is : ", bin)
	fmt.Println(err)
	fmt.Println()

	user2, err := user.Decode(bin)
	fmt.Println(user2)
	fmt.Println(err)
}
