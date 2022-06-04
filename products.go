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
	"strconv"
)

type Product struct {
  ID          int
  Name        string
  Description *string
  Brand       Brand
  Categories  []Category
}

func getAllProducts() []Product {
  qry := `
  SELECT Products.ID, Products.Name, Products.Description,
  Brands.ID, Brands.Name, Brands.DisplayName, Brands.LogoURL, Brands.Website
  FROM Products
  INNER JOIN Brands ON Products.BrandID = Brands.ID`
  rows, err := db.Query(qry)
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  var products []Product
  for rows.Next() {
    var p Product

    if err := rows.Scan(
      &p.ID, &p.Name, &p.Description, &p.Brand.ID, 
      &p.Brand.Name, &p.Brand.DisplayName, &p.Brand.LogoURL, &p.Brand.Website,
    ); err != nil {
      panic(err)
    }

    p.Categories = getProductCategories(p.ID)
    products = append(products, p)
  }

  return products
}

func getProductCategories(productId int) []Category {
  qry := `
  SELECT Product_Categories.CategoryID, Categories.ParentID, Categories.Name
  FROM Product_Categories 
  INNER JOIN Categories ON Product_Categories.CategoryID = Categories.ID 
  WHERE Product_Categories.ProductID = ` + strconv.Itoa(productId)
  rows, err := db.Query(qry)
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  var categories []Category
  for rows.Next() {
    var c Category
    var parentId *int

    if err := rows.Scan(&c.ID, &parentId, &c.Name); err != nil {
      panic(err)
    }

    if parentId != nil {
      parentCategory := getCategory(*parentId)
      c.Parent = &parentCategory
    }

    categories = append(categories, c)
  }

  return categories
}

func insertProduct(p Product) bool {
  query, err := db.Prepare("INSERT INTO Products (Name, Description, BrandID) VALUES (?, ?, ?)")
  if err != nil {
    panic(err)
  }
  defer query.Close()

  res, err := query.Exec(p.Name, p.Description, p.Brand.ID)
  if err != nil {
    panic(err)
  }

  newID, err := res.LastInsertId()
  if err != nil {
    panic(err)
  }

  for _, i := range p.Categories {
    addProductCategory(int(newID), i.ID)
  }

  return true
}
