package handlers

import (
    "net/http"
    "fast-line/internal/models"
    "fast-line/internal/services"
    "fast-line/internal/utils"
    "github.com/google/uuid" // Assuming you use this for UUIDs
    "github.com/gin-gonic/gin"
    "strings"
)

type CircuitHandler struct {
    circuitService *services.CircuitService
}

func NewCircuitHandler(circuitService *services.CircuitService) *CircuitHandler {
    return &CircuitHandler{
        circuitService: circuitService,
    }
}

func (h *CircuitHandler) Create(c *gin.Context) {
    var req models.CreateCircuitRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
        return
    }
    
    circuit, err := h.circuitService.Create(&req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create circuit", err.Error())
        return
    }
    
    utils.SuccessResponse(c, http.StatusCreated, "Circuit created successfully", circuit)
}

func (h *CircuitHandler) Update(c *gin.Context) {
    // 1. Get the ID from the URL path parameter
    idStr := c.Param("id")
    circuitID, err := uuid.Parse(idStr)
    if err != nil {
        // If the ID in the URL is not a valid UUID, it's a client error
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid circuit ID format", err.Error())
        return
    }

    // 2. Bind the JSON request body (your code for this is already perfect)
    var req models.UpdateCircuitRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
        return
    }

    // 3. Call the service with BOTH the ID and the request data
    circuit, err := h.circuitService.Update(circuitID, &req)
    if err != nil {
        // --- Refined Error Handling ---
        // Check if the error is a "not found" error from the service
        if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
            utils.ErrorResponse(c, http.StatusNotFound, "Circuit not found", err.Error())
        } else {
            // For all other errors, it's a server-side issue
            utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update circuit", err.Error())
        }
        return
    }

    // 4. Send the success response (your code for this is also perfect)
    utils.SuccessResponse(c, http.StatusOK, "Circuit updated successfully", circuit)
}
func (h *CircuitHandler) Get(c *gin.Context) {
    id, err := utils.GetUUIDFromPath(c, "id")
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
        return
    }
    
    circuit, err := h.circuitService.Get(id)
    if err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "Circuit not found", err.Error())
        return
    }
    
    utils.SuccessResponse(c, http.StatusOK, "Circuit found successfully", circuit)
}

// func (h *CircuitHandler)Delete(c *gin.Context) {
//     id, err := utils.GetUUIDFromPath(c, "id")
//     if err != nil {
//         utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
//         return
//     }

// 	err = h.circuitService.Delete(id)
//     if err != nil {
//         utils.ErrorResponse(c, http.StatusNotFound, "Circuit not found", err.Error())
//         return
//     }
    
//     utils.SuccessResponse(c, http.StatusOK, "Circuit deleted successfully", nil)
// }

func (h *CircuitHandler) Delete(c *gin.Context) {
    // Get the ID from the URL path parameter.
    idStr := c.Param("id")
    circuitID, err := uuid.Parse(idStr)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid circuit ID format", err.Error())
        return
    }

    // Call the service to perform the delete operation.
    err = h.circuitService.Delete(circuitID)
    if err != nil {
        // If the service returns a "not found" error, send a 404 response.
        if strings.Contains(err.Error(), "not found") {
            utils.ErrorResponse(c, http.StatusNotFound, "Circuit to delete not found", err.Error())
            return
        }
        
        // For all other errors, send a 500 internal server error.
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete circuit", err.Error())
        return
    }

    // On successful deletion, send a 204 No Content response.
    // This is the standard RESTful practice for a successful DELETE request.
    c.Status(http.StatusNoContent)
}

func (h *CircuitHandler) List(c *gin.Context) {

    circuits, err := h.circuitService.List(10, 0)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list circuits", err.Error())
        return
    }
    
    utils.SuccessResponse(c, http.StatusOK, "Circuits listed successfully", circuits)
}