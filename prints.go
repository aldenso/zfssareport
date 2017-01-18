package main

import "fmt"

//Header print header for output.
func Header(msg string) {
	fmt.Println("########################################################################################")
	fmt.Printf("## %-82s ##\n", msg)
	fmt.Println("########################################################################################")
}

//PrintPools prints some pools values.
func PrintPools() {
	POOLS = GetPools()
	Header("Pools information")
	fmt.Printf("%6s %8s %-35s %10s %10s %10s %10s %10s %10s\n", "Name", "Status", "Profile", "Total(GB)", "Avail(GB)", "Free(GB)",
		"UsagData(GB)", "UsagSnaps(GB)", "UsagTotal(GB)")
	for _, pool := range POOLS.List {
		fmt.Printf("%6s %8s %-35s %10.2f %10.2f %10.2f %10.2f %10.2f %10.2f\n", pool.Name, pool.Status, pool.Profile,
			pool.Usage.Total/(1024*1024*1024), pool.Usage.Available/(1024*1024*1024),
			pool.Usage.Free/(1024*1024*1024), pool.Usage.UsageData/(1024*1024*1024),
			pool.Usage.UsageSnapshots/(1024*1024*1024), pool.Usage.UsageTotal/(1024*1024*1024))
	}
}

//PrintProjects prints some projects values for all pools.
func PrintProjects() {
	Header("Projects information")
	fmt.Printf("%-22s %10s %10s %8s %10s %-40s\n", "Project", "Reserv(GB)", "Quota(GB)", "Pool", "SpaceTotal(GB)", "mountpoint")
	for _, pool := range POOLS.List {
		projects := GetProjects(pool.Name)
		for _, project := range projects.List {
			POOLSPROJECTS[pool.Name] = append(POOLSPROJECTS[pool.Name], project.Name)
			fmt.Printf("%-22s %10.2f %10.2f %8s %10.2f %-40s\n", project.Name, project.Reservation/(1024*1024*1024),
				project.Quota/(1024*1024*1024), project.Pool, project.SpaceTotal/(1024*1024*1024), project.MountPoint)
		}
	}
}

//PrintFilesystems prints some filesystems values for all projects in all pools.
func PrintFilesystems() {
	Header("Filesystems information")
	fmt.Printf("%-18s %8s %-22s %8s %8s %8s %8s %8s %6s %-35s\n", "Filesystem", "Pool", "Project", "Reserv(GB)", "Quota(GB)", "STotal(GB)",
		"user", "group", "perms", "mountpoint")
	for pool, projects := range POOLSPROJECTS {
		for _, project := range projects {
			filesystems := GetFilesystems(pool, project)
			for _, filesystem := range filesystems.List {
				fmt.Printf("%-18s %8s %-22s %8.2f %8.2f %8.2f %8s %8s %6s %-35s\n", filesystem.Name, filesystem.Pool, filesystem.Project,
					filesystem.Reservation/(1024*1024*1024), filesystem.Quota/(1024*1024*1024), filesystem.SpaceTotal/(1024*1024*1024),
					filesystem.RootUser, filesystem.RootGroup, filesystem.RootPermissions, filesystem.MountPoint)
			}
		}
	}
}

//PrintLUNS prints some luns values for all projects in all pools.
func PrintLUNS() {
	Header("LUNS information")
	fmt.Printf("%-16s %8s %-15s %8s %5s %-15s %8s %32s %8s %8s\n", "LUN", "Pool", "Project", "Status", "ANumber", "IGroup",
		"TGroup", "GUID", "VolSize(GB)", "STotal(GB)")
	for pool, projects := range POOLSPROJECTS {
		for _, project := range projects {
			luns := GetLUNS(pool, project)
			for _, lun := range luns.List {
				fmt.Printf("%-16s %8s %-15s %8s %5d %-15s %8s %32s %8.2f %8.2f\n", lun.Name, lun.Pool, lun.Project, lun.Status,
					lun.AssignedNumber, lun.InitiatorGroup, lun.TargetGroup, lun.LunGUID, lun.VolSize/(1024*1204*1024),
					lun.SpaceTotal/(1024*1024*1024))
			}
		}
	}
}
