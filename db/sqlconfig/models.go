// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import ()

type User struct {
	ID         int64
	Name       string
	Email      string
	Password   string
	Isverified bool
	Otp        string
}
