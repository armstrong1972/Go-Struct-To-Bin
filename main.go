package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/rs/xid"
)

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

func main() {
	// Init the value of user object
	var user1 = User{
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
	user1.token = xid.New()
	user1.rsa.RSA_Private = []byte{11, 22, 33, 44, 55}
	fmt.Println("\n---- Original : User1 ------\n", user1)

	// Encoding to Bin
	buf, _ := EncodeStruct(&user1)
	fmt.Println("\n---- Encoded : []byte ------\n ", buf)

	// Decoding to Struct
	var user2 *User = *(**User)(unsafe.Pointer(&buf))
	fmt.Println("\n---- Decoded : User2 -------\n", user2)

	// Check them are same address
	user2.username = "Rainbow"
	fmt.Println("\n\n  User1.name=", user1.username, " @ ", unsafe.Pointer(&user1))
	fmt.Println("  User2.name=", user2.username, " @ ", unsafe.Pointer(user2))
}
