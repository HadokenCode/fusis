package api

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/luizbafilho/fusis/api/types"
)

type ServiceResponse struct {
	types.Service
	Destinations []types.Destination
}

func (as ApiService) serviceList(c *gin.Context) {
	services := as.balancer.GetServices()
	if len(services) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	response := []ServiceResponse{}
	for _, s := range services {
		response = append(response, ServiceResponse{
			Service:      s,
			Destinations: as.balancer.GetDestinations(&s),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (as ApiService) serviceGet(c *gin.Context) {
	serviceId := c.Param("service_name")
	service, err := as.balancer.GetService(serviceId)
	if err != nil {
		c.Error(err)
		if err == types.ErrServiceNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("GetService() failed: %v", err)})
		}
		return
	}

	response := ServiceResponse{
		Service:      *service,
		Destinations: as.balancer.GetDestinations(service),
	}

	c.JSON(http.StatusOK, response)
}

func (as ApiService) serviceCreate(c *gin.Context) {
	var newService types.Service
	if err := c.BindJSON(&newService); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Guarantees that no one tries to create a destination together with a service
	// newService.Destinations = []types.Destination{}

	if _, errs := govalidator.ValidateStruct(newService); errs != nil {
		c.Error(errs)
		c.JSON(http.StatusBadRequest, gin.H{"errors": govalidator.ErrorsByField(errs)})
		return
	}

	// If everthing is ok send it to Raft
	err := as.balancer.AddService(&newService)
	if err != nil {
		c.Error(err)
		if err == types.ErrServiceAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("UpsertService() failed: %v", err)})
		}
		return
	}

	c.Header("Location", fmt.Sprintf("/services/%s", newService.Name))
	c.JSON(http.StatusCreated, newService)
}

func (as ApiService) serviceDelete(c *gin.Context) {
	serviceId := c.Param("service_name")
	_, err := as.balancer.GetService(serviceId)
	if err != nil {
		c.Error(err)
		if err == types.ErrServiceNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("GetService() failed: %v", err)})
		}
		return
	}

	err = as.balancer.DeleteService(serviceId)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("DeleteService() failed: %v\n", err)})
		return
	}

	c.Status(http.StatusNoContent)
}

func (as ApiService) destinationCreate(c *gin.Context) {
	serviceName := c.Param("service_name")
	service, err := as.balancer.GetService(serviceName)
	if err != nil {
		c.Error(err)
		if err == types.ErrServiceNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("GetService() failed: %v", err)})
		}
		return
	}

	var destination *types.Destination
	if err := c.BindJSON(&destination); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	destination.ServiceId = serviceName
	if destination.Weight == 0 {
		destination.Weight = 1
	}
	if destination.Mode == "" {
		destination.Mode = "route"
	}

	if _, errs := govalidator.ValidateStruct(destination); errs != nil {
		c.Error(errs)
		c.JSON(http.StatusBadRequest, gin.H{"errors": govalidator.ErrorsByField(errs)})
		return
	}

	err = as.balancer.AddDestination(service, destination)
	if err != nil {
		c.Error(err)
		if err == types.ErrDestinationAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("AddDestination() failed: %v", err)})
		}
		return
	}

	c.Header("Location", fmt.Sprintf("/services/%s/destinations/%s", serviceName, destination.Name))
	c.JSON(http.StatusCreated, destination)
}

func (as ApiService) destinationDelete(c *gin.Context) {
	destinationId := c.Param("destination_name")
	dst, err := as.balancer.GetDestination(destinationId)
	if err != nil {
		c.Error(err)
		if _, ok := err.(types.ErrNotFound); ok {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("GetDestination() failed: %v", err)})
		}
		return
	}

	err = as.balancer.DeleteDestination(dst)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("DeleteDestination() failed: %v", err)})
	}

	err = as.balancer.DelCheck(dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("DelCheck() failed: %v", err)})
		return
	}

	c.Status(http.StatusNoContent)
}

func (as ApiService) flush(c *gin.Context) {
	// err := as.types.Flush()
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	//
	// err = types.Flush()
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
}
