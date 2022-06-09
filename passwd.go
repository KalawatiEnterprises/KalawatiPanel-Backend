/*
 * KalawatiPanel-Backend - Backend for KalawatiPanel
 * Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@protonmail.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
  "fmt"
  "time"
  "math/rand"
  "github.com/gin-gonic/gin"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ/*&^%$#@!()[]{}")

var passwd string

func init() {
  rand.Seed(time.Now().UnixNano())
  passwd = genPasswd(128)
  // passwd = "testPasswd"

  fmt.Println(passwd)
}

func genPasswd(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func passwdMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
	password := c.PostForm("passwd")

    if password != passwd {
      c.AbortWithStatusJSON(401, gin.H{"error": "Incorrect Password."})
      fmt.Println("Wrong password entered:", password)
      return
    }

    c.Next()
  }
}
