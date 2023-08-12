

/*
Package scheduler provides scheduler main logic implements.

Transactions
There are followed transactions to do applications or taskgroups operation:
LAUCH: create and launch an application from version definition
DELETE: delete application
UPDATE: update application
SCALE: scale up or scale down application's instances
RESCHEDULE: reschedule taskgroup when it is fail or required by API

Service
When applications are running, sometimes they are binded to some services, and need to export to services,
Service Manager is implemented to do application bind and export, it watches followed events:
Taskgroup Add
Taskgroup Delete
Taskgroup Update
Service Add
Service Update
Service delete

Status Report
When tasks run on slave machine, the status will reported by mesos slave, the report message is processed by function StatusReport

Health Check Report
If a running taskgroup is configured to do health check, the health-check result will reported by healthy module, the messeages are processed by HealthyReport

Deployment related functions
The deployments' rollingupdate is implemented by using application transactions, refer to function DeploymentCheck

DataChecker
DataChecker is responsable for dirty or error data in ZK
refer to DataCheckManage

Message
Message is used to send message to executor, just as localfile, signal ...
*/
package scheduler
