// This package hashes password
package utills

import "golang.org/x/crypto/bcrypt"



// Will hash the password using bcrypt and return the hashed password or an error if hashing fails.
func HashPassword(password string) (string,error){

	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",err
	}

	return string(bytes),nil
}

// Will compare the provided password with the hashed password and return true if they match, or false if they do not match or if an error occurs during comparison.
func CheckPasswordHash(password,hash string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	if err!=nil{
		return false
	}

	return true
}