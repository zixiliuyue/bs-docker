

package scaler

// ScalerProcess through the api interface to the bcs scheduler, to scale up, scale down
// the scaler't target ref deployment application
type ScalerProcess interface {
	// ScaleDeployment TODO
	// scale deployment
	ScaleDeployment(namespace, name string, instance uint) error

	// ScaleApplication TODO
	// scale application
	ScaleApplication(namespace, name string, instance uint) error
}
