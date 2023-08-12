

/*
Package pluginManager provides dymamic getting mesos slave attributes implements.

Scheduler support for the use of plugins to get mesos slave attributes.
It is mainly applicable to the acquisition of dynamic attributes, example for container ip resources,
net flow.

The types of plugin are mainly including dynamic, executable.
User can implement specific plugin based on their own scenarios.

	//mesos slave attribute plugin's names
	pluginsNames := []string{"ip-resources","net-flow"}

	pluginer,err := NewPluginManager(pluginsNames)
	if err != nil {
		//...
		return err
	}

	//get mesos slave's dynamic attributes
	attrs,err := pluginer.GetHostAttributes(para)
	if err != nil {
		//...
		return err
	}

	//set dynamic attributes to mesos slave
	//...
*/
package pluginManager
