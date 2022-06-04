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

import "fmt"

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

  var parentCategory Category
  if parentId != nil {
    parentCategory = getCategory(*parentId)
  }
  fmt.Println(parentCategory)

  return category
}

