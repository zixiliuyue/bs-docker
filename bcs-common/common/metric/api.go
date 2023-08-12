

package metric

// NewMetricController metric controller
func NewMetricController(conf Config, healthFunc HealthFunc, metrics ...*MetricContructor) error {
	return newMetricHandler(conf, healthFunc, metrics...)
}

// RunModeType xxx
type RunModeType string

// RoleType xxx
// used when your module running with Master_Slave_Mode mode
type RoleType string

const (
	// Master_Slave_Mode xxx
	Master_Slave_Mode RunModeType = "master-slave"
	// Master_Master_Mode xxx
	Master_Master_Mode RunModeType = "master-master"
	// MasterRole xxx
	MasterRole RoleType = "master"
	// SlaveRole xxx
	SlaveRole RoleType = "slave"
	// UnknownRole xxx
	UnknownRole RoleType = "unknown"
)

// Config xxx
type Config struct {
	// name of your module
	ModuleName string
	// running mode of your module
	// could be one of Master_Slave_Mode or Master_Master_Mode
	RunMode RunModeType
	// ip address of this module running on
	IP string
	// port number of the metric's http handler depends on.
	MetricPort uint
	// cluster id of your module belongs to.
	ClusterID string
	// self defined info labeled on your metrics.
	// deprecated, unused now.
	Labels map[string]string
	// whether disable golang's metric, default is false.
	DisableGolangMetric bool
	// metric http server's ssl configuration
	SvrCaFile   string
	SvrCertFile string
	SvrKeyFile  string
	SvrKeyPwd   string
}

// HealthFunc xxx
type HealthFunc func() HealthMeta

// HealthMeta xxx
type HealthMeta struct {
	// the running role of your module when you are running with Master_Slave_Mode.
	// must be not empty. if you set with an empty value, an error will be occurred.
	// when your module is running in Master_Master_Mode,  this filed should be set
	// with value of "Slave".
	CurrentRole RoleType `json:"current_role"`
	// if this module is healthy
	IsHealthy bool `json:"healthy"`
	// messages which describes the health status
	Message string `json:"message"`
}

// MetricMeta xxx
type MetricMeta struct {
	// metric's name
	Name string
	// metric's help info, which should be short and briefly.
	Help string
	// metric labels, which can describe the special info about this metric.
	ConstLables map[string]string
}

// GetMetaFunc xxx
type GetMetaFunc func() *MetricMeta

// GetResultFunc xxx
type GetResultFunc func() (*MetricResult, error)

// MetricContructor xxx
type MetricContructor struct {
	GetMeta   GetMetaFunc
	GetResult GetResultFunc
}

// MetricResult xxx
type MetricResult struct {
	Value *FloatOrString
	// variable labels means that this labels value can be changed with each call.
	VariableLabels map[string]string
}
