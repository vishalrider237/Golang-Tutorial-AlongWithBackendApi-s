package CrudWithMySql

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Create a new employee with the input payload
// @Tags employees
// @Accept  json
// @Produce  json
// @Param employee body models.Employee true "Employee"
// @Success 200 {object} models.Employee
// @Failure 400 {object} gin.H
// @Router /employees [post]
func CreateEmployee(c *gin.Context) {
	var employee Employee
	err := c.ShouldBindJSON(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&employee)
	c.JSON(http.StatusOK, employee)
}

// GetEmployees godoc
// @Summary Get all employees
// @Description Get a list of all employees
// @Tags employees
// @Produce  json
// @Success 200 {array} models.Employee
// @Router /employees [get]
func GetAllEmployee(c *gin.Context) {
	var employees []Employee
	DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}

// GetEmployee godoc
// @Summary Get an employee by ID
// @Description Get a single employee by ID
// @Tags employees
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee
// @Failure 404 {object} gin.H
// @Router /employees/{id} [get]
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee Employee
	err := DB.First(&employee, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// UpdateEmployee godoc
// @Summary Update an employee
// @Description Update an existing employee by ID
// @Tags employees
// @Accept  json
// @Produce  json
// @Param id path int true "Employee ID"
// @Param employee body models.Employee true "Employee"
// @Success 200 {object} models.Employee
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /employees/{id} [put]
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee Employee
	err := DB.First(&employee, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	err = c.ShouldBindJSON(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Save(employee)
	c.JSON(http.StatusOK, employee)
}

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Delete an employee by ID
// @Tags employees
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /employees/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	err := DB.Delete(Employee{}, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
