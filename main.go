package main

import (
	_ "time"
	_ "cloud-api-go/amazon/monitor"
	"cloud-api-go/amazon/resources"
)

func main()  {

	/*now:=time.Now()
	later :=now.Add(-time.Hour)
    monitor.RdsCPUUsagFilterDBName(now,later,"accountdb01")
	monitor.RdsConnFilterDBName(now,later,"accountdb01")

	 */

	resources.Test()

}