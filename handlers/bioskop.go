package handlers

import (
	"go-bioskop/db"
	"go-bioskop/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// POST /bioskop
func CreateBioskop(c *gin.Context) {
	var b models.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil || b.Nama == "" || b.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi wajib diisi"})
		return
	}

	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	err := db.Conn.QueryRow(query, b.Nama, b.Lokasi, b.Rating).Scan(&b.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, b)
}

// GET /bioskop
func GetAllBioskop(c *gin.Context) {
	rows, err := db.Conn.Query("SELECT id, nama, lokasi, rating FROM bioskop")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Bioskop
	for rows.Next() {
		var b models.Bioskop
		if err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		list = append(list, b)
	}

	c.JSON(http.StatusOK, list)
}

// GET /bioskop/:id
func GetBioskopByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var b models.Bioskop
	query := "SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1"
	err = db.Conn.QueryRow(query, id).Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// PUT /bioskop/:id
func UpdateBioskop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var b models.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil || b.Nama == "" || b.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi wajib diisi"})
		return
	}

	query := "UPDATE bioskop SET nama=$1, lokasi=$2, rating=$3 WHERE id=$4"
	res, err := db.Conn.Exec(query, b.Nama, b.Lokasi, b.Rating, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	b.ID = id
	c.JSON(http.StatusOK, b)
}

// DELETE /bioskop/:id
func DeleteBioskop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	query := "DELETE FROM bioskop WHERE id = $1"
	res, err := db.Conn.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
