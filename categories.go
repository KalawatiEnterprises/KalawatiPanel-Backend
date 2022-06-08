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

type Category struct {
  ID     int
  Parent *Category
  Name   string
}

func getCategory(categoryId int) Category {
  rows, err := db.Prepare("SELECT ID, ParentID, Name FROM Categories WHERE ID = ?")
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  var category Category
  var parentId *int
  if err := rows.QueryRow(categoryId).Scan(&category.ID, &parentId, &category.Name); err != nil {
    panic(err)
  }

  if parentId != nil {
    parentCategory := getCategory(*parentId)
    category.Parent = &parentCategory
  }

  return category
}

func getAllCategories() []Category {
  rows, err := db.Query("SELECT ID, ParentID, Name FROM Categories")
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

func insertCategory(c Category) bool {
  query, err := db.Prepare("INSERT INTO Categories (ParentID, Name) VALUES (?, ?)")
  if err != nil {
    panic(err)
  }
  defer query.Close()

  if c.Parent.ID < 1 {
    _, err = query.Exec(nil, c.Name)
  } else {
    _, err = query.Exec(c.Parent.ID, c.Name)
  }

  if err != nil {
    panic(err)
  }

  return true
}

func getCategoryChildren(parentId int) []int {
  rows, err := db.Query("SELECT ID FROM Categories WHERE ParentID = " + strconv.Itoa(parentId))
  if err != nil {
    panic(err)
  }

  var children []int
  for rows.Next() {
    var c int
    if err := rows.Scan(&c); err != nil {
      panic(err)
    }

    children = append(children, c)
  }

  return children
}

func deleteCategory(categoryId int) bool {
  // else it will delete EVERYTHING
  if categoryId < 1 { return false }

  // delete associated products first
  query, err := db.Prepare("DELETE From Product_Categories WHERE CategoryID = ?")
  if err != nil {
    panic(err)
  }
  _ , err = query.Exec(categoryId)
  if err != nil {
    panic(err)
  }

  // if category is deleted children are also deleted
  query, err = db.Prepare("DELETE FROM Categories WHERE ID = ?")
  if err != nil {
    panic(err)
  }
  _ , err = query.Exec(categoryId)
  if err != nil {
    panic(err)
  }

  for _, i := range getCategoryChildren(categoryId) {
    deleteCategory(i)
  }

  return true
}

func updateCategory(c Category) bool {
  query, err := db.Prepare(`
  UPDATE Categories SET
  ParentID = ?,
  Name     = ?
  WHERE ID = ?`)
  if err != nil {
    panic(err)
  }

  if c.Parent.ID < 1 {
    _, err = query.Exec(nil, c.Name, c.ID)
  } else {
    _, err = query.Exec(c.Parent.ID, c.Name, c.ID)
  }

  if err != nil {
    panic(err)
  }

  return true
}
