package device

import (
	"hu-design-project/application/handler/device/handler_impl"
	"hu-design-project/application/repository"
)

func InitializeDeviceHandler(
	repository repository.DevicesRepository,
) *Handler {
	deviceHandler := Handler{}
	deviceHandler.GetDevices = handler_impl.NewDevicesHandler(repository)
	return &deviceHandler
}
