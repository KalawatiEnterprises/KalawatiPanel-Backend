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
  "encoding/json"
  "fmt"
)

func main() {
  r := gin.New()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  // serve static (image) files
  r.Static("/images", os.Getenv("PRODUCT_IMAGES_DIR"))

  // product routes
  r.GET("/api/products", func (ctx *gin.Context) {
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(200, getAllProducts())
  })

  r.POST("/api/products", passwdMiddleware(), func (ctx *gin.Context) {
    var product Product
    json.Unmarshal([]byte(ctx.PostForm("data")), &product)
    ctx.JSON(200, insertProduct(product))
  })

  r.DELETE("/api/products", passwdMiddleware(), func (ctx *gin.Context) {
    var product Product
    json.Unmarshal([]byte(ctx.PostForm("data")), &product)
    ctx.JSON(200, deleteProduct(product.ID))
  })

  r.PUT("/api/products", passwdMiddleware(), func (ctx *gin.Context) {
    var product Product
    json.Unmarshal([]byte(ctx.PostForm("data")), &product)
    ctx.JSON(200, updateProduct(product))
  })

  // brand routes
  r.GET("/api/brands", func (ctx *gin.Context) {
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(200, getAllBrands())
  })

  r.POST("/api/brands", passwdMiddleware(), func (ctx *gin.Context) {
    var brand Brand
    json.Unmarshal([]byte(ctx.PostForm("data")), &brand)
    ctx.JSON(200, insertBrand(brand))
  })

  r.DELETE("/api/brands", passwdMiddleware(), func (ctx *gin.Context) {
    var brand Brand
    json.Unmarshal([]byte(ctx.PostForm("data")), &brand)
    ctx.JSON(200, deleteBrand(brand.ID))
  })

  r.PUT("/api/brands", passwdMiddleware(), func (ctx *gin.Context) {
    var brand Brand
    json.Unmarshal([]byte(ctx.PostForm("data")), &brand)
    ctx.JSON(200, updateBrand(brand))
  })

  // category routes
  r.GET("/api/categories", func (ctx *gin.Context) {
    ctx.Header("Content-Type", "application/json")
    ctx.JSON(200, getAllCategories())
  })

  r.POST("/api/categories", passwdMiddleware(), func (ctx *gin.Context) {
    var category Category
    json.Unmarshal([]byte(ctx.PostForm("data")), &category)
    ctx.JSON(200, insertCategory(category))
  })

  r.DELETE("/api/categories", passwdMiddleware(), func (ctx *gin.Context) {
    var category Category
    json.Unmarshal([]byte(ctx.PostForm("data")), &category)
    ctx.JSON(200, deleteCategory(category.ID))
  })

  r.PUT("/api/categories", passwdMiddleware(), func (ctx *gin.Context) {
    var category Category
    json.Unmarshal([]byte(ctx.PostForm("data")), &category)
    ctx.JSON(200, updateCategory(category))
  })

  // image routes
  r.GET("/api/images/:productId", func (ctx *gin.Context) {
    ctx.JSON(200, getImages(ctx.Param("productId")))
  })

  r.POST("/api/images/", passwdMiddleware(), func (ctx *gin.Context) {
    ctx.JSON(200, insertImage(ctx))
  })

  r.DELETE("/api/images/", passwdMiddleware(), func (ctx *gin.Context) {
    ctx.JSON(200, deleteImage(ctx.PostForm("path")))
  })

  r.Run(":4001")
}
