package productcontroller

import (
	"hello-world/helpers"
	"hello-world/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	// Menggunakan helper function helpers.JsonResponse untuk menghasilkan respons
	 helpers.JsonResponse(c, true, "Berhasil Mengambil Data Semua Product", products)
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helpers.NotFoundResponse(c)
			return
		default:
			helpers.InternalServerResponse(c, err.Error())
			return
		}
	}
	helpers.JsonResponse(c, true, "Berhasil Mengambil Data Satu Product", product)
	// Kode untuk menampilkan detail produk
}

func Create(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.BadRequestResponse(c, "Data yang dimasukkan tidak valid")
		return
	}

	if err := models.DB.Create(&input).Error; err != nil {
		helpers.InternalServerResponse(c, err.Error())
		return
	}

	helpers.JsonResponse(c, true, "Berhasil Membuat Produk Baru", input)
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helpers.NotFoundResponse(c)
			return
		default:
			helpers.InternalServerResponse(c, err.Error())
			return
		}
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.BadRequestResponse(c, "Data yang dimasukkan tidak valid")
		return
	}

	if err := models.DB.Model(&product).Updates(&input).Error; err != nil {
		helpers.InternalServerResponse(c, err.Error())
		return
	}

	helpers.JsonResponse(c, true, "Berhasil Memperbarui Produk", input)
}

func Delete(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helpers.NotFoundResponse(c)
			return
		default:
			helpers.InternalServerResponse(c, err.Error())
			return
		}
	}

	if err := models.DB.Delete(&product).Error; err != nil {
		helpers.InternalServerResponse(c, err.Error())
		return
	}

	helpers.JsonResponse(c, true, "Berhasil Menghapus Produk", nil)
}
