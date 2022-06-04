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

type Brand struct {
  ID          int
  Name        *string
  DisplayName string
  LogoURL     *string
  Website     *string
}

func getAllBrands() []Brand {
  rows, err := db.Query("SELECT ID, Name, DisplayName, LogoURL, Website FROM Brands")
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  var brands []Brand
  for rows.Next() {
    var b Brand
    if err := rows.Scan(&b.ID, &b.Name, &b.DisplayName, &b.LogoURL, &b.Website); err != nil {
      panic(err)
    }
    brands = append(brands, b)
  }

  return brands
}

func insertBrand(b Brand) bool {
  query, err := db.Prepare("INSERT INTO Brands (Name, DisplayName, LogoURL, Website) VALUES (?, ?, ?, ?)")
  if err != nil {
    panic(err)
  }
  defer query.Close()

  _, err = query.Exec(b.Name, b.DisplayName, b.LogoURL, b.Website)
  if err != nil {
    panic(err)
  }

  return true
}

func deleteBrand(brandId int) bool {
  query, err := db.Prepare("DELETE From Brands WHERE ID = ?")
  if err != nil {
    panic(err)
  }
  _ , err = query.Exec(brandId)
  if err != nil {
    panic(err)
  }

  return true
}

func updateBrand(b Brand) bool {
  // update product details
  query, err := db.Prepare(`
  UPDATE Brands SET
  Name        = ?,
  DisplayName = ?,
  LogoURL     = ?,
  Website     = ?
  WHERE ID    = ?`)
  if err != nil {
    panic(err)
  }
  _ , err = query.Exec(b.Name, b.DisplayName, b.LogoURL, b.Website, b.ID)
  if err != nil {
    panic(err)
  }

  return true
}
