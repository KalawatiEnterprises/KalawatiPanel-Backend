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
  "github.com/gin-gonic/gin"
  "os"
  "io/ioutil"
)

func getImages(productId string) []string {
  imgs, err := ioutil.ReadDir(imagesDir + productId)
  if err != nil {
    panic(err)
  }

  var imageLinks []string
  for _, i := range imgs {
    imageLinks = append(imageLinks, productId + "/" + i.Name())
  }

  return imageLinks
}

func insertImage(ctx *gin.Context) bool {
  img, err := ctx.FormFile("image")
  if err != nil {
    panic(err)
	return false
  }

  productId := ctx.PostForm("product-id")
  imageId := ctx.PostForm("image-id")

  err = ctx.SaveUploadedFile(img, imagesDir + productId + "/" + imageId + ".webp")
  if err != nil {
    panic(err)
    return false
  }

  return true
}

func deleteImage(path string) bool {
  err := os.Remove(imagesDir + path)
  if err != nil {
    panic(err)
    return false
  }

  return true
}
