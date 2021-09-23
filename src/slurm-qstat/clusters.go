package main

import (
	"fmt"
	"log"
	"strings"
)

func getClusterInformation() (map[string]clusterData, error) {
	var result = make(map[string]clusterData)

	raw, err := executeCommand("sacctmgr", "--parsable2", "--noheader", "show", "clusters", "Format=Classification,Cluster,ClusterNodes,ControlHost,ControlPort,DefaultQOS,Fairshare,Flags,GrpTRESMins,GrpTRES,GrpJobs,GrpMemory,GrpNodes,GrpSubmitJob,MaxTRESMins,MaxTRES,MaxJobs,MaxNodes,MaxSubmitJobs,MaxWall,NodeCount,PluginIDSelect,RPC,TRES")
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(raw), "\n") {
		if len(line) == 0 {
			continue
		}

		var cdata = make(map[string]string)
		for i, field := range strings.Split(line, "|") {
			fname := clusterFields[i]
			cdata[fname] = field
		}
		name, found := cdata["Cluster"]
		if !found {
			log.Panic("BUG: No field named Cluster found in sacctmgr output")
		}
		result[name] = cdata
	}
	return result, nil
}

func checkClusterlist(clist []string, all map[string]clusterData) error {
	for _, c := range clist {
		_, found := all[c]
		if !found {
			return fmt.Errorf("Cluster %s is not in list of defined SLURM clusters", c)
		}
	}
	return nil
}

func filterCluster(cdata map[string]clusterData, cl []string) []clusterData {
	var result []clusterData

	if len(cl) > 0 {
		for _, c := range cl {
			cluster, found := cdata[c]
			if found {
				result = append(result, cluster)
			}
		}
	} else {
		for _, v := range cdata {
			result = append(result, v)
		}
	}

	return result
}
