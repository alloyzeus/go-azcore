package azcore

/**/ /**/

// ServiceServerConfig holds the configuration for a service client.
type ServiceServerConfig interface {
	ServiceConfig
}

// ServiceServer provides an abstraction for all service clients.
type ServiceServer interface {
	Service
	AZServiceServer()
}

// ServiceServerModule provides all the required to instantiate a service
// client.
type ServiceServerModule struct {
	ServiceServerConfigSkeleton func() ServiceServerConfig
	NewServiceServer            func(ServiceServerConfig) (ServiceServer, Error)
}

var _ ServiceModule = ServiceServerModule{}

// AZServiceModule is required for conformance with ServiceModule.
func (ServiceServerModule) AZServiceModule() {}