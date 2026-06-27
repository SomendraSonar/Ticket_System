package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Ticket
func CreateTicket(c *gin.Context) {
	var ticket Ticket

	if err := c.BindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	userID, _ := c.Get("user_id")
	id := userID.(int)

	ticket.ID = len(tickets) + 1
	ticket.UserID = id
	ticket.Status = "open"

	tickets = append(tickets, ticket)

	c.JSON(http.StatusCreated, ticket)
}

// Get all tickets of logged-in user
func GetTickets(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(int)

	var myTickets []Ticket

	for _, ticket := range tickets {
		if ticket.UserID == id {
			myTickets = append(myTickets, ticket)
		}
	}

	c.JSON(http.StatusOK, myTickets)
}

// Get single ticket by ID
func GetTicket(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(int)

	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ticket id",
		})
		return
	}

	for _, ticket := range tickets {
		if ticket.ID == ticketID &&
			ticket.UserID == id {

			c.JSON(http.StatusOK, ticket)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Ticket not found",
	})
}

// Update ticket status
func UpdateStatus(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := userID.(int)

	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ticket id",
		})
		return
	}

	var input struct {
		Status string `json:"status"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	for i, ticket := range tickets {

		if ticket.ID == ticketID &&
			ticket.UserID == id {

			// Closed ticket cannot be reopened
			if ticket.Status == "closed" {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Closed ticket cannot be reopened",
				})
				return
			}

			// Allowed statuses
			if input.Status != "open" &&
				input.Status != "in_progress" &&
				input.Status != "closed" {

				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid status",
				})
				return
			}

			tickets[i].Status = input.Status

			c.JSON(http.StatusOK, tickets[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Ticket not found",
	})
}
