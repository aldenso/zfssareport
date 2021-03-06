package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/aldenso/zfssareport/model"
	"github.com/aldenso/zfssareport/utils"
	"github.com/spf13/afero"
)

//PrintPools prints some pools values and create file to dump all values.
func PrintPools(pools *model.Pools, fs afero.Fs) {
	//POOLS = GetPools()
	file, err := utils.CreateFile(fs, dirname, "pools.csv")
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
	if !silent {
		utils.Header("Pools information")
		fmt.Printf("%-6s %-8s %-35s %10s %10s %10s %10s %10s %10s\n",
			"Name", "Status", "Profile", "Total(GB)", "Avail(GB)",
			"Free(GB)", "UData(GB)", "USnaps(GB)", "UTotal(GB)")
		fmt.Println("=====================================================================================================================")
	}
	for _, pool := range pools.List {
		if !silent {
			fmt.Printf("%-6s %-8s %-35s %10.2f %10.2f %10.2f %10.2f %10.2f %10.2f\n",
				pool.Name, pool.Status, pool.Profile, pool.PoolUsage.Total/(1024*1024*1024), pool.PoolUsage.Available/(1024*1024*1024),
				pool.PoolUsage.Free/(1024*1024*1024), pool.PoolUsage.UsageData/(1024*1024*1024), pool.PoolUsage.UsageSnapshots/(1024*1024*1024),
				pool.PoolUsage.UsageTotal/(1024*1024*1024))
		}

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

//PrintProjects prints some projects values for all pools.
func PrintProjects(projects *model.Projects, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "projects.csv")
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
	if !silent {
		utils.Header("Projects information")
		fmt.Printf("%-15s %-10s %-10s %-8s %10s %40s\n", "Project", "Reserv(GB)", "Quota(GB)", "Pool", "STotal(GB)", "mountpoint")
		fmt.Println("=====================================================================================================================")
	}
	for _, project := range projects.List {
		//projects := GetProjects(pool.Name)
		if !silent {
			fmt.Printf("%-15s %-10.2f %-10.2f %-8s %10.2f %40s\n", project.Name, project.Reservation/(1024*1024*1024),
				project.Quota/(1024*1024*1024), project.Pool, project.SpaceTotal/(1024*1024*1024), project.MountPoint)
		}
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
	writer.Flush()
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

//PrintFilesystems prints some filesystems values for all projects in all pools.
func PrintFilesystems(filesystems *model.Filesystems, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "filesystems.csv")
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
	if !silent {
		utils.Header("Filesystems information")
		fmt.Printf("%-12s %-8s %-15s %9s %9s %9s %8s %8s %4s %15s\n",
			"Filesystem", "Pool", "Project", "Reser(GB)", "Quota(GB)", "Total(GB)", "user", "group", "perms", "mountpoint")
		fmt.Println("=====================================================================================================================")
	}
	for _, filesystem := range filesystems.List {
		if !silent {
			fmt.Printf("%-12s %-8s %-15s %9.2f %9.2f %9.2f %8s %8s %4s %15s\n", filesystem.Name, filesystem.Pool,
				filesystem.Project, filesystem.Reservation/(1024*1024*1024), filesystem.Quota/(1024*1024*1024),
				filesystem.SpaceTotal/(1024*1024*1024), filesystem.RootUser, filesystem.RootGroup, filesystem.RootPermissions,
				filesystem.MountPoint)
		}

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
	writer.Flush()
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

//PrintLUNS prints some luns values for all projects in all pools.
func PrintLUNS(luns *model.LUNS, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "luns.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"name", "pool", "project", "status", "assignednumber", "initiatorgroup", "targetgroup",
		"lunguid", "volsize", "space_available", "space_data", "space_snapshots", "space_total", "canonical_name",
		"checksum", "compression", "compressratio", "copies", "creation", "dedup", "encryption", "exported",
		"fixednumber", "href", "id", "keychangedate", "keystatus", "logbias", "lunumber", "maxblocksize", "nodestroy",
		"secondarycache", "snaplabel", "sparse", "volblocksize", "writecache", "source_checksum", "source_compression",
		"source_copies", "source_dedup", "source_encryption", "source_exported", "source_keychangedate", "source_logbias",
		"source_maxblocksize", "source_rrsrc_actions", "source_secondary"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("LUNS information")
		fmt.Printf("%-8s %-8s %-15s %9s %4s %-15s %32s %8s %8s\n", "LUN", "Pool", "Project", "Status", "Num", "InitiatorGrp",
			"GUID", "VolS(GB)", "Total(GB)")
		fmt.Println("=====================================================================================================================")
	}
	for _, lun := range luns.List {
		if !silent {
			fmt.Printf("%-8s %-8s %-15s %9s %4d %-15s %32s %8.2f %8.2f\n", lun.Name, lun.Pool, lun.Project, lun.Status,
				lun.AssignedNumber, lun.InitiatorGroup, lun.LunGUID, lun.VolSize/(1024*1204*1024),
				lun.SpaceTotal/(1024*1024*1024))
		}

		line := fmt.Sprintf("%s;%s;%s;%s;%d;"+
			/*line2*/ "%s;%s;%s;%.2f;%.2f;"+
			/*line3*/ "%.2f;%.2f;%.2f;%s;%s;"+
			/*line4*/ "%s;%.2f;%d;%s;%t;"+
			/*line5*/ "%s;%t;%t;%s;%s;"+
			/*line6*/ "%s;%s;%s;%s;%d;"+
			/*line7*/ "%t;%s;%s;%t;%d;"+
			/*line8*/ "%t;%s;%s;%s;%s;"+
			/*line9*/ "%s;%s;%s;%s;%s;"+
			/*line10*/ "%s;%s",
			lun.Name, lun.Pool, lun.Project, lun.Status, lun.AssignedNumber,
			/*line2*/ lun.InitiatorGroup, lun.TargetGroup, lun.LunGUID, lun.VolSize, lun.SpaceAvailable,
			/*line3*/ lun.SpaceData, lun.SpaceSnapShots, lun.SpaceTotal, lun.CanonicalName, lun.Checksum,
			/*line4*/ lun.Compression, lun.CompressRatio, lun.Copies, lun.Creation, lun.Dedup,
			/*line5*/ lun.Encryption, lun.Exported, lun.FixedNumber, lun.HREF, lun.ID,
			/*line6*/ lun.KeyChangeDate, lun.KeyStatus, lun.Logbias, lun.LUNumber, lun.MaxBlockSize,
			/*line7*/ lun.Nodestroy, lun.SecondaryCache, lun.SnapLabel, lun.Sparse, lun.VolBlockSize,
			/*line8*/ lun.WriteCache, lun.LunSource.CheckSum, lun.LunSource.Compression, lun.LunSource.Copies, lun.LunSource.Dedup,
			/*line9*/ lun.LunSource.Encryption, lun.LunSource.Exported, lun.LunSource.KeyChangeDate, lun.LunSource.Logbias, lun.LunSource.MaxBlockSize,
			/*line10*/ lun.LunSource.RRSRCActions, lun.LunSource.SecondaryCache)

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

//PrintNetInterfaces to print some interfaces info and create csv report.
func PrintNetInterfaces(netints *model.NetInterfaces, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "interfaces.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"interface", "class", "links", "label", "state", "v4addrs", "v4dhcp", "v4directnets",
		"v6addrs", "v6dhcp", "v6directnets", "curaddrs", "enable", "href", "admin"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Network Interfaces information")
		fmt.Printf("%-15s %-8s %-25s %-20s %-25s %5s\n", "interface", "class", "links", "label", "v4addrs", "state")
		fmt.Println("=====================================================================================================================")
	}
	for _, interf := range netints.List {
		if !silent {
			interf.PrintNetInterfaceInfo()
		}
		line := fmt.Sprintf("%s;%s;%s;%s;%s;"+
			/*line2*/ "%s;%t;%s;%s;%t;"+
			/*line3*/ "%s;%s;%t;%s;%t",
			interf.Interface, interf.Class, interf.Links, interf.Label, interf.State,
			/*line2*/ interf.V4Addrs, interf.V4DHCP, interf.V4DirectNets, interf.V6Addrs, interf.V6DHCP,
			/*line3*/ interf.V6DirectNets, interf.CurAddrs, interf.Enable, interf.HREF, interf.Admin)

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

//PrintFCInitiators to print some fc initiators info and create csv report.
func PrintFCInitiators(initiators *model.FCInitiators, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "fc-initiators.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"alias", "href", "initiator"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("FC Initiators information")
		fmt.Printf("%-25s %-15s\n", "initiator", "alias")
		fmt.Println("=====================================================================================================================")
	}
	for _, initiator := range initiators.List {
		if !silent {
			initiator.PrintInitiatorInfo()
		}
		line := fmt.Sprintf("%s;%s;%s", initiator.Alias, initiator.HREF, initiator.Initiator)
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

//PrintFCInitiatorGroups to print some fc Initiators info and create csv report.
func PrintFCInitiatorGroups(groups *model.FCInitiatorGroups, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "fc-initiator-groups.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"href", "initiator", "name"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("FC Initiator Groups information")
		fmt.Println("=====================================================================================================================")
	}
	for _, group := range groups.List {
		if !silent {
			group.PrintInitiatorGroupInfo()
		}
		line := fmt.Sprintf("%s;%s;%s", group.HREF, group.Initiators, group.Name)
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

//PrintFCTargets to print some fc targets info and create csv report.
func PrintFCTargets(targets *model.FCTargets, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "fc-targets.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"wwn", "speed", "port", "mode", "discovered_ports", "href", "invalid_crc_count",
		"invalid_tx_word_count", "link_failure_count", "loss_of_signal_count", "loss_of_sync_count",
		"protocol_error_count"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("FC Targets information")
		fmt.Printf("%-25s %-18s %-10s %-8s %-8s %-8s %-8s %8s\n", "wwn", "port", "speed", "mode", "discPorts",
			"LossSyn", "LossSignal", "LinkFail")
		fmt.Println("=====================================================================================================================")
	}
	for _, target := range targets.List {
		if !silent {
			target.PrintTargetInfo()
		}
		line := fmt.Sprintf("%s;%s;%s;%s;%d;%s;%d;%d;%d;%d;%d;%d",
			target.WWN, target.Speed, target.Port, target.Mode, target.DiscoveredPorts,
			target.HREF, target.InvalidCRCCount, target.InvalidTXWordCount, target.LinkFailureCount,
			target.LossOfSignalCount, target.LossOfSyncCount, target.ProtocolErrorCount)
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

//PrintIscsiInitiators to print some initiators info and create csv report.
func PrintIscsiInitiators(initiators *model.IscsiInitiators, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "iscsi-initiators.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"alias", "chapsecret", "chapuser", "href", "initiator"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Iscsi Initiators information")
		fmt.Printf("%-55s %-15s %-11s %20s\n", "initiator", "alias", "chap user", "chap secret")
		fmt.Println("=====================================================================================================================")
	}
	for _, initiator := range initiators.List {
		if !silent {
			initiator.PrintInitiatorInfo()
		}
		line := fmt.Sprintf("%s;%s;%s;%s;%s", initiator.Alias, initiator.ChapSecret, initiator.ChapUser,
			initiator.HREF, initiator.Initiator)
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

//PrintIscsiInitiatorGroups to print some iscsi initiators info and create csv report.
func PrintIscsiInitiatorGroups(groups *model.IscsiInitiatorGroups, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "iscsi-initiator-groups.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"href", "initiator", "name"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Iscsi Initiator Groups information")
		fmt.Println("=====================================================================================================================")
	}
	for _, group := range groups.List {
		if !silent {
			group.PrintInitiatorGroupInfo()
		}
		line := fmt.Sprintf("%s;%s;%s", group.HREF, group.Initiators, group.Name)
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

//PrintChassis to print some chassis components info and create csv report.
func PrintChassis(chassisslice *model.ChassisAll, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "chassis.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"faulted", "href", "locate", "manufacturer", "model", "name", "part",
		"path", "revision", "rpm", "serial", "type"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Chassis information")
		fmt.Printf("%-14s %-25s %-10s %-8s %-12s %s\n", "name", "model", "type", "faulted", "serial", "path")
		fmt.Println("=====================================================================================================================")
	}
	for _, chassis := range chassisslice.List {
		if !silent {
			chassis.PrintChassisInfo()
		}
		line := fmt.Sprintf("%t;%s;%t;%s;%s;"+
			"%s;%s;%d;%s;%d;"+
			"%s;%s",
			chassis.Faulted, chassis.HREF, chassis.Locate, chassis.Manufacturer, chassis.Model,
			chassis.Name, chassis.Part, chassis.Path, chassis.Revision, chassis.RPM,
			chassis.Serial, chassis.Type)
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

//PrintProblems to print some problems info and create csv report.
func PrintProblems(problems *model.Problems, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"action", "code", "description", "diagnosed", "href", "impact", "phoned_home",
		"response", "severity", "type", "url", "uuid"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Problems information")
		fmt.Println("=====================================================================================================================")
	}
	for _, problem := range problems.List {
		if !silent {
			problem.PrintProblemInfo()
		}
		line := fmt.Sprintf("%s;%s;%s;%s;%s;"+
			"%s;%s;%s;%s;%s;"+
			"%s;%s",
			problem.Action, problem.Code, problem.Description, problem.Diagnosed, problem.HREF,
			problem.Impact, problem.PhonedHome, problem.Response, problem.Severity, problem.Type,
			problem.URL, problem.UUID)
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

//PrintNetDevices to print some network devices info and create csv report.
func PrintNetDevices(devices *model.NetDevices, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "devices.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"active", "device", "duplex", "factory_mac", "href", "media", "speed", "up"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Network Devices information")
		fmt.Printf("%-8s %-8s %-8s %-15s %-15s %-18s %12s\n", "device", "active", "up", "speed", "media", "factory_mac", "duplex")
		fmt.Println("=====================================================================================================================")
	}
	for _, device := range devices.List {
		if !silent {
			device.PrintNetDeviceInfo()
		}
		line := fmt.Sprintf("%t;%s;%s;%s;%s;"+
			"%s;%s;%t",
			device.Active, device.Device, device.Duplex, device.FactoryMAC, device.HREF,
			device.Media, device.Speed, device.UP)
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

//PrintNetDatalinks to print some network datalinks info and create csv report.
func PrintNetDatalinks(datalinks *model.NetDatalinks, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "datalinks.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"class", "datalink", "duplex", "href", "label", "links", "mac", "mtu", "speed"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Network Datalinks information")
		fmt.Printf("%-10s %-15s %-15s %-15s %-20s %-10s %s\n", "class", "datalink", "label", "links", "mac", "mtu", "speed")
		fmt.Println("=====================================================================================================================")
	}
	for _, datalink := range datalinks.List {
		if !silent {
			datalink.PrintNetDatalinkInfo()
		}
		line := fmt.Sprintf("%s;%s;%s;%s;%s;"+
			"%s;%s;%d;%s",
			datalink.Class, datalink.Datalink, datalink.Duplex, datalink.HREF, datalink.Label,
			datalink.Links, datalink.MAC, datalink.MTU, datalink.Speed)
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

//PrintZFSSAVersion to print zfs version info and create csv report
func PrintZFSSAVersion(version *model.Version, fs afero.Fs) {
	version.WriteCSV(fs, dirname)
	if !silent {
		version.PrintVersionInfo()
	}
}

//PrintZFSSACluster to print zfs cluster info and create csv report
func PrintZFSSACluster(cluster *model.Cluster, fs afero.Fs) {
	cluster.WriteCSV(fs, dirname)
	if !silent {
		cluster.PrintClusterInfo()
	}
}

//PrintUsers to print some users info and create csv report.
func PrintUsers(users *model.Users, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "users.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"logname", "type", "uid", "fullname", "initial_password", "require_annotation", "roles",
		"kiosk_mode", "kiosk_screen", "href"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Users information")
		fmt.Printf("%-15s %-10s %-11s %-35s %-8s\n", "logname", "type", "uid", "roles", "annotation")
		fmt.Println("=====================================================================================================================")
	}
	for _, user := range users.List {
		if !silent {
			user.PrintUserInfo()
		}
		line := fmt.Sprintf("%s;%s;%d;%s;%s;%v;%s;%v;%s;%s", user.Logname, user.Type, user.UID, user.FullName, user.InitialPassword, user.RequireAnnotation,
			user.Roles, user.KioskMode, user.KioskScreen, user.HREF)
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

//PrintRoutes to print some routes info and create csv report.
func PrintRoutes(routes *model.Routes, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "routes.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"status", "family", "destination", "mask", "href", "interface", "type", "gateway"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("Routes information")
		fmt.Printf("%-18s %-18s %-10s %-8s %-6s %-10s %s\n", "destination", "gateway", "interface", "status", "mask",
			"type", "family")
		fmt.Println("=====================================================================================================================")
	}
	for _, route := range routes.List {
		if !silent {
			route.PrintRouteInfo()
		}
		line := fmt.Sprintf("%s;%s;%s;%d;%s;%s;%s;%s", route.Status, route.Family, route.Destination, route.Mask, route.HREF, route.Interface,
			route.Type, route.Gateway)
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

//PrintServices to print some general services info and create csv report.
func PrintServices(services *model.Services, fs afero.Fs) {
	file, err := utils.CreateFile(fs, dirname, "services.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	//writer.Comma = ';'
	fileheader := []string{"status", "href", "name", "log href", "log size"}
	if err := writer.Write(fileheader); err != nil {
		log.Fatal(err)
	}
	if !silent {
		utils.Header("General Service information")
		fmt.Printf("%-15s %-10s %-60s %-8s\n", "name", "status", "log href", "log size")
		fmt.Println("=====================================================================================================================")
	}
	for _, service := range services.List {
		if !silent {
			service.PrintServiceInfo()
		}
		line := fmt.Sprintf("%s;%s;%s;%s;%d", service.Status, service.HREF, service.Name, service.Log.HREF, service.Log.Size)
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
