package model

import "fmt"

//Services struct for services.
type Services struct {
	List []Service `json:"services"`
}

//Service struct for service
type Service struct {
	Status string     `json:"<status>"`
	HREF   string     `json:"href"`
	Name   string     `json:"name"`
	Log    ServiceLog `json:"log,omitempty"`
}

//ServiceLog struct for logs in service.
type ServiceLog struct {
	HREF string `json:"href"`
	Size int    `json:"size"`
}

//PrintServiceInfo print general service info.
func (service *Service) PrintServiceInfo() {
	fmt.Printf("%-15s %-10s %-60s %-8d\n",
		service.Name, service.Status, service.Log.HREF, service.Log.Size)
}
