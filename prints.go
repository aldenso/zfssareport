package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"sync"

	"github.com/spf13/afero"
)

//Header print header for output.
func Header(msg string) {
	fmt.Println("#####################################################################################################################")
	fmt.Printf("## %-111s ##\n", msg)
	fmt.Println("#####################################################################################################################")
}

//PrintPools prints some pools values and create file to dump all values.
func PrintPools(pools Pools, fs afero.Fs) {
	//POOLS = GetPools()
	file, err := CreateFile(fs, dirname, "pools.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"pool", "status", "profile", "available", "usage_snapshots", "used", "compression", "usage_data",
		"free", "dedupratio", "total", "usage_total", "peer", "owner", "asm"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	Header("Pools information")
	fmt.Printf("%-6s %-8s %-35s %10s %10s %10s %10s %10s %10s\n",
		"Name", "Status", "Profile", "Total(GB)", "Avail(GB)",
		"Free(GB)", "UData(GB)", "USnaps(GB)", "UTotal(GB)")
	for _, pool := range pools.List {
		fmt.Printf("%-6s %-8s %-35s %10.2f %10.2f %10.2f %10.2f %10.2f %10.2f\n",
			pool.Name, pool.Status, pool.Profile, pool.PoolUsage.Total/(1024*1024*1024), pool.PoolUsage.Available/(1024*1024*1024),
			pool.PoolUsage.Free/(1024*1024*1024), pool.PoolUsage.UsageData/(1024*1024*1024), pool.PoolUsage.UsageSnapshots/(1024*1024*1024),
			pool.PoolUsage.UsageTotal/(1024*1024*1024))

		line := fmt.Sprintf("%s;%s;%s;%.2f;%.2f;%.2f;%d;%.2f;%.2f;%d;%.2f;%.2f;%s;%s;%s", pool.Name, pool.Status, pool.Profile,
			pool.PoolUsage.Available, pool.PoolUsage.UsageSnapshots, pool.PoolUsage.Used, pool.PoolUsage.Compression, pool.PoolUsage.UsageData,
			pool.PoolUsage.Free, pool.PoolUsage.Dedupratio, pool.PoolUsage.Total, pool.PoolUsage.UsageTotal, pool.Peer, pool.Owner, pool.ASN)
		record := strings.Split(line, ";")
		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

//CreateMapPoolsProjects create a map for projects in pools.
func CreateMapPoolsProjects(pools Pools) map[string]Projects {
	poolsprojects := make(map[string]Projects)
	var wg sync.WaitGroup
	for _, pool := range pools.List {
		wg.Add(1)
		go func(pool Pool) {
			defer wg.Done()
			projects := GetProjects(pool.Name)
			poolsprojects[pool.Name] = projects
		}(pool)
	}
	wg.Wait()
	return poolsprojects
}

//PrintProjects prints some projects values for all pools.
func PrintProjects(pmap map[string]Projects, fs afero.Fs) {
	file, err := CreateFile(fs, dirname, "projects.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"name", "pool", "reservation", "quota", "space_total", "mountpoint", "aclinherit", "aclmode", "atime",
		"canonical_name", "checksum", "compression", "compressratio", "copies", "creation", "dedup", "default_group",
		"default_permissions", "default_sparce", "default_user", "default_volblocksize", "default_volsize", "defaultgroupquota",
		"defaultuserquota", "encryption", "exported", "href", "id", "keychangedate", "keystatus", "logbias", "maxblocksize",
		"nbmand", "nodestroy", "readonly", "record_size", "rstchown", "secondarycache", "sharedav", "shareftp",
		"sharenfs", "sharesftp", "sharesmb", "sharetftp", "snapdir", "snaplabel", "source_aclinherit", "source_aclmode",
		"source_atime", "source_checksum", "source_compression", "source_copies", "source_dedup", "source_exported",
		"source_keychangedate", "source_logbias", "source_maxblocksize", "source_mountpoint", "source_nbmand", "source_readonly",
		"source_record_size", "source_reservation;", "source_rrsrc_actions", "source_rstchown", "source_secondarycache",
		"source_sharedav", "source_shareftp", "source_sharenfs", "source_sharesftp", "source_sharesmb", "source_sharetftp",
		"source_snapdir", "source_vscan", "space_available", "space_data", "space_snapshots", "space_unused_res",
		"space_unused_res_shares", "vscan"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}

	Header("Projects information")
	fmt.Printf("%-15s %-10s %-10s %-8s %10s %40s\n", "Project", "Reserv(GB)", "Quota(GB)", "Pool", "STotal(GB)", "mountpoint")
	for _, projects := range pmap {
		//projects := GetProjects(pool.Name)
		for _, project := range projects.List {
			fmt.Printf("%-15s %-10.2f %-10.2f %-8s %10.2f %40s\n", project.Name, project.Reservation/(1024*1024*1024),
				project.Quota/(1024*1024*1024), project.Pool, project.SpaceTotal/(1024*1024*1024), project.MountPoint)

			line := fmt.Sprintf("%s;%s;%.2f;%.2f;%.2f;"+
				/*line2*/ "%s;%s;%s;%t;%s;"+
				/*line3*/ "%s;%s;%.2f;%d;%s;"+
				/*line4*/ "%t;%s;%s;%t;%s;"+
				/*line5*/ "%.2f;%.2f;%.2f;%.2f;%s;"+
				/*line6*/ "%t;%s;%s;%s;%s;"+
				/*line7*/ "%s;%d;%t;%t;%t;"+
				/*line8*/ "%f;%t;%s;%s;%s;"+
				/*line9*/ "%s;%s;%s;%s;%s;"+
				/*line10*/ "%s;%s;%s;%s;%s;"+
				/*line11*/ "%s;%s;%s;%s;%s;"+
				/*line12*/ "%s;%s;%s;%s;%s;"+
				/*line13*/ "%s;%s;%s;%s;%s;"+
				/*line14*/ "%s;%s;%s;%s;%s;"+
				/*line15*/ "%s;%s;%s;%.2f;%.2f;"+
				/*line16*/ "%.2f;%.2f;%.2f;%t",
				project.Name, project.Pool, project.Reservation, project.Quota, project.SpaceTotal,
				/*line2*/ project.MountPoint, project.ACLinherit, project.ACLMode, project.ATime, project.CanonicalName,
				/*line3*/ project.CheckSum, project.Compression, project.CompressRatio, project.Copies, project.Creation,
				/*line4*/ project.Dedup, project.DefaultGroup, project.DefaultPermissions, project.DefaultSparse, project.DefaultUser,
				/*line5*/ project.DefaulVolBlockSize, project.DefaultVolSize, project.DefaultGroupQuota, project.DefaultUserQuota, project.Encryption,
				/*line6*/ project.Exported, project.HREF, project.ID, project.KeyChangeDate, project.KeyStatus,
				/*line7*/ project.Logbias, project.MaxBlockSize, project.Nbmand, project.Nodestroy, project.ReadOnly,
				/*line8*/ project.RecordSize, project.Rstchown, project.SecondaryCache, project.ShareDAV, project.ShareFTP,
				/*line9*/ project.ShareNFS, project.ShareSFTP, project.ShareSMB, project.ShareTFTP, project.SnapDir,
				/*line10*/ project.SnapLabel, project.ProjectSource.ACLinherit, project.ProjectSource.ACLMode, project.ProjectSource.ATime, project.ProjectSource.CheckSum,
				/*line11*/ project.ProjectSource.Compression, project.ProjectSource.Copies, project.ProjectSource.Dedup, project.ProjectSource.Exported, project.ProjectSource.KeyChangeDate,
				/*line12*/ project.ProjectSource.Logbias, project.ProjectSource.MaxBlockSize, project.ProjectSource.MountPoint, project.ProjectSource.Nbmand, project.ProjectSource.ReadOnly,
				/*line13*/ project.ProjectSource.RecordSize, project.ProjectSource.Reservation, project.ProjectSource.RRSRCActions, project.ProjectSource.Rstchown, project.ProjectSource.SecondaryCache,
				/*line14*/ project.ProjectSource.ShareDAV, project.ProjectSource.ShareFTP, project.ProjectSource.ShareNFS, project.ProjectSource.ShareSFTP, project.ProjectSource.ShareSMB,
				/*line15*/ project.ProjectSource.ShareTFTP, project.ProjectSource.SnapDir, project.ProjectSource.VScan, project.SpaceAvailable, project.SpaceData,
				/*line16*/ project.SpaceSnapShots, project.SpaceUnusedRes, project.SpaceUnusedResShares, project.VScan)
			record := strings.Split(line, ";")
			if err := writer.Write(record); err != nil {
				log.Fatal(err)
			}
		}
	}
	writer.Flush()
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

//CreateFSSlice create a map for filesystems
func CreateFSSlice(pmap map[string]Projects) []Filesystems {
	var poolsprojectsfs struct {
		List []Filesystems
		mu   sync.Mutex
	}
	var wg sync.WaitGroup
	for pool := range pmap {
		for _, project := range pmap[pool].List {
			wg.Add(1)
			go func(pool string, project Project) {
				defer wg.Done()
				//fmt.Println("get fs proj:", project.Name, "pool", pool)
				filesystems := GetFilesystems(pool, project.Name)
				poolsprojectsfs.mu.Lock()
				poolsprojectsfs.List = append(poolsprojectsfs.List, *filesystems)
				poolsprojectsfs.mu.Unlock()
			}(pool, project)
		}
		wg.Wait()
	}

	return poolsprojectsfs.List
}

//CreateLUNSSlice create a map for filesystems
func CreateLUNSSlice(pmap map[string]Projects) []LUNS {
	var poolsprojectsluns struct {
		List []LUNS
		mu   sync.Mutex
	}
	var wg sync.WaitGroup
	for pool := range pmap {
		for _, project := range pmap[pool].List {
			wg.Add(1)
			go func(pool string, project Project) {
				defer wg.Done()
				//fmt.Println("get lun proj:", project.Name, "pool", pool)
				luns := GetLUNS(pool, project.Name)
				poolsprojectsluns.mu.Lock()
				poolsprojectsluns.List = append(poolsprojectsluns.List, *luns)
				poolsprojectsluns.mu.Unlock()
			}(pool, project)
		}
		wg.Wait()
	}
	return poolsprojectsluns.List
}

//PrintFilesystems prints some filesystems values for all projects in all pools.
func PrintFilesystems(allfs []Filesystems, fs afero.Fs) {
	file, err := CreateFile(fs, dirname, "filesystems.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"name", "pool", "project", "reservation", "quota", "space_total", "root_user", "root_group", "root_permissions",
		"mountpoint", "space_available", "space_data", "space_snapshot", "space_unused_res", "aclinherit", "aclmode", "atime",
		"canonical_name", "case_sensitivity", "checksum", "compression", "compressratio", "copies", "creation", "dedup", "encryption",
		"exported", "href", "id", "keychangedata", "keystatus", "logbias", "maxblocksize", "nbmand", "nodedestroy", "normalization",
		"quota_snap", "readonly", "record_size", "reservation_snap", "rstchown", "secondarycache", "shadow", "sharedav", "shareftp",
		"sharenfs", "sharesftp", "sharesmb", "sharetftp", "snapdir", "snaplabel", "utf8only", "vscan", "source_aclinherit", "source_aclmode",
		"source_atime", "source_checksum", "source_compression", "source_copies", "source_dedup", "source_exported", "source_keychangedata",
		"source_logbias", "source_maxblocksize", "source_mountpoint", "source_nbmand", "source_readonly", "source_record_size", "source_reservation",
		"source_rrsrc_actions", "source_rstchown", "source_secondarycache", "source_sharedav", "source_shareftp", "source_sharenfs", "source_sharesftp",
		"source_sharesmb", "source_sharetftp", "source_snapdir", "source_vscan"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	Header("Filesystems information")
	fmt.Printf("%-12s %-8s %-15s %9s %9s %9s %8s %8s %4s %15s\n",
		"Filesystem", "Pool", "Project", "Reser(GB)", "Quota(GB)", "Total(GB)", "user", "group", "perms", "mountpoint")
	for _, filesystems := range allfs {
		for _, filesystem := range filesystems.List {
			fmt.Printf("%-12s %-8s %-15s %9.2f %9.2f %9.2f %8s %8s %4s %15s\n", filesystem.Name, filesystem.Pool,
				filesystem.Project, filesystem.Reservation/(1024*1024*1024), filesystem.Quota/(1024*1024*1024),
				filesystem.SpaceTotal/(1024*1024*1024), filesystem.RootUser, filesystem.RootGroup, filesystem.RootPermissions,
				filesystem.MountPoint)

			line := fmt.Sprintf("%s;%s;%s;%.2f;%.2f;"+
				/*line2*/ "%.2f;%s;%s;%s;%s;"+
				/*line3*/ "%.2f;%.2f;%.2f;%.2f;%s;"+
				/*line4*/ "%s;%t;%s;%s;%s;"+
				/*line5*/ "%s;%.2f;%d;%s;%t;"+
				/*line6*/ "%s;%t;%s;%s;%s;"+
				/*line7*/ "%s;%s;%d;%t;%t;"+
				/*line8*/ "%s;%t;%t;%.2f;%t;"+
				/*line9*/ "%t;%s;%s;%s;%s;"+
				/*line10*/ "%s;%s;%s;%s;%s;"+
				/*line11*/ "%s;%t;%t;%s;%s;"+
				/*line12*/ "%s;%s;%s;%s;%s;"+
				/*line13*/ "%s;%s;%s;%s;%s;"+
				/*line14*/ "%s;%s;%s;%s;%s;"+
				/*line15*/ "%s;%s;%s;%s;%s;"+
				/*line16*/ "%s;%s;%s;%s;%s",
				filesystem.Name, filesystem.Pool, filesystem.Project, filesystem.Reservation, filesystem.Quota,
				/*line2*/ filesystem.SpaceTotal, filesystem.RootUser, filesystem.RootGroup, filesystem.RootPermissions, filesystem.MountPoint,
				/*line3*/ filesystem.SpaceAvailable, filesystem.SpaceData, filesystem.SpaceSnapShots, filesystem.SpaceUnusedRes, filesystem.ACLinherit,
				/*line4*/ filesystem.ACLMode, filesystem.ATime, filesystem.CanonicalName, filesystem.CaseSensitivity, filesystem.CheckSum,
				/*line5*/ filesystem.Compression, filesystem.CompressRatio, filesystem.Copies, filesystem.Creation, filesystem.Dedup,
				/*line6*/ filesystem.Encryption, filesystem.Exported, filesystem.HREF, filesystem.ID, filesystem.KeyChangeDate,
				/*line7*/ filesystem.KeyStatus, filesystem.Logbias, filesystem.MaxBlockSize, filesystem.Nbmand, filesystem.Nodestroy,
				/*line8*/ filesystem.Normalization, filesystem.QuotaSnap, filesystem.ReadOnly, filesystem.RecordSize, filesystem.ReservationSnap,
				/*line9*/ filesystem.Rstchown, filesystem.SecondaryCache, filesystem.Shadow, filesystem.ShareDAV, filesystem.ShareFTP,
				/*line10*/ filesystem.ShareNFS, filesystem.ShareSFTP, filesystem.ShareSMB, filesystem.ShareTFTP, filesystem.SnapDir,
				/*line11*/ filesystem.SnapLabel, filesystem.UTF8Only, filesystem.VScan, filesystem.FSSource.ACLinherit, filesystem.FSSource.ACLMode,
				/*line12*/ filesystem.FSSource.ATime, filesystem.FSSource.CheckSum, filesystem.FSSource.Compression, filesystem.FSSource.Copies, filesystem.FSSource.Dedup,
				/*line13*/ filesystem.FSSource.Exported, filesystem.FSSource.KeyChangeDate, filesystem.FSSource.Logbias, filesystem.FSSource.MaxBlockSize, filesystem.FSSource.MountPoint,
				/*line14*/ filesystem.FSSource.Nbmand, filesystem.FSSource.ReadOnly, filesystem.FSSource.RecordSize, filesystem.FSSource.Reservation, filesystem.FSSource.RRSRCActions,
				/*line15*/ filesystem.FSSource.Rstchown, filesystem.FSSource.SecondaryCache, filesystem.FSSource.ShareDAV, filesystem.FSSource.ShareFTP, filesystem.FSSource.ShareNFS,
				/*line16*/ filesystem.FSSource.ShareSFTP, filesystem.FSSource.ShareSMB, filesystem.FSSource.ShareTFTP, filesystem.FSSource.SnapDir, filesystem.FSSource.VScan)
			record := strings.Split(line, ";")
			if err := writer.Write(record); err != nil {
				log.Fatal(err)
			}
		}
	}
	writer.Flush()
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

//PrintLUNS prints some luns values for all projects in all pools.
func PrintLUNS(allLuns []LUNS, fs afero.Fs) {
	Header("LUNS information")
	fmt.Printf("%-16s %8s %-15s %8s %5s %-15s %8s %32s %8s %8s\n", "LUN", "Pool", "Project", "Status", "ANumber", "IGroup",
		"TGroup", "GUID", "VolSize(GB)", "STotal(GB)")
	for _, luns := range allLuns {
		for _, lun := range luns.List {
			fmt.Printf("%-16s %8s %-15s %8s %5d %-15s %8s %32s %8.2f %8.2f\n", lun.Name, lun.Pool, lun.Project, lun.Status,
				lun.AssignedNumber, lun.InitiatorGroup, lun.TargetGroup, lun.LunGUID, lun.VolSize/(1024*1204*1024),
				lun.SpaceTotal/(1024*1024*1024))
		}
	}
}
