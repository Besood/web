package db

import (
	"golang.org/x/crypto/bcrypt"
)

func UserPassHash(pass string)(string,error){
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), err
}

func UserPassMach(hash,pw string)bool{
		return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}