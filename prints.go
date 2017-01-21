package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

//Header print header for output.
func Header(msg string) {
	fmt.Println("########################################################################################")
	fmt.Printf("## %-82s ##\n", msg)
	fmt.Println("########################################################################################")
}

//PrintPools prints some pools values and create file to dump all values.
func PrintPools() {
	POOLS = GetPools()
	file, err := CreateFile(Fs, dirname, "pools.csv")
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
	fmt.Printf("%6s %8s %-35s %10s %10s %10s %10s %10s %10s\n",
		"Name", "Status", "Profile", "Total(GB)", "Avail(GB)",
		"Free(GB)", "UsagData(GB)", "UsagSnaps(GB)", "UsagTotal(GB)")
	for _, pool := range POOLS.List {
		fmt.Printf("%6s %8s %-35s %10.2f %10.2f %10.2f %10.2f %10.2f %10.2f\n",
			pool.Name, pool.Status, pool.Profile, pool.Usage.Total/(1024*1024*1024), pool.Usage.Available/(1024*1024*1024),
			pool.Usage.Free/(1024*1024*1024), pool.Usage.UsageData/(1024*1024*1024), pool.Usage.UsageSnapshots/(1024*1024*1024),
			pool.Usage.UsageTotal/(1024*1024*1024))

		line := fmt.Sprintf("%s;%s;%s;%.2f;%.2f;%.2f;%d;%.2f;%.2f;%d;%.2f;%.2f;%s;%s;%s", pool.Name, pool.Status, pool.Profile,
			pool.Usage.Available, pool.Usage.UsageSnapshots, pool.Usage.Used, pool.Usage.Compression, pool.Usage.UsageData,
			pool.Usage.Free, pool.Usage.Dedupratio, pool.Usage.Total, pool.Usage.UsageTotal, pool.Peer, pool.Owner, pool.ASN)
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
func PrintProjects() {
	file, err := CreateFile(Fs, dirname, "projects.csv")
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
	fmt.Printf("%-22s %10s %10s %8s %10s %-40s\n", "Project", "Reserv(GB)", "Quota(GB)", "Pool", "SpaceTotal(GB)", "mountpoint")
	for _, pool := range POOLS.List {
		projects := GetProjects(pool.Name)
		for _, project := range projects.List {
			POOLSPROJECTS[pool.Name] = append(POOLSPROJECTS[pool.Name], project.Name)
			fmt.Printf("%-22s %10.2f %10.2f %8s %10.2f %-40s\n", project.Name, project.Reservation/(1024*1024*1024),
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
				/*line10*/ project.SnapLabel, project.Source.ACLinherit, project.Source.ACLMode, project.Source.ATime, project.Source.CheckSum,
				/*line11*/ project.Source.Compression, project.Source.Copies, project.Source.Dedup, project.Source.Exported, project.Source.KeyChangeDate,
				/*line12*/ project.Source.Logbias, project.Source.MaxBlockSize, project.Source.MountPoint, project.Source.Nbmand, project.Source.ReadOnly,
				/*line13*/ project.Source.RecordSize, project.Source.Reservation, project.Source.RRSRCActions, project.Source.Rstchown, project.Source.SecondaryCache,
				/*line14*/ project.Source.ShareDAV, project.Source.ShareFTP, project.Source.ShareNFS, project.Source.ShareSFTP, project.Source.ShareSMB,
				/*line15*/ project.Source.ShareTFTP, project.Source.SnapDir, project.Source.VScan, project.SpaceAvailable, project.SpaceData,
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

//PrintFilesystems prints some filesystems values for all projects in all pools.
func PrintFilesystems() {
	Header("Filesystems information")
	fmt.Printf("%-18s %8s %-22s %8s %8s %8s %8s %8s %6s %-35s\n",
		"Filesystem", "Pool", "Project", "Reserv(GB)", "Quota(GB)", "STotal(GB)", "user", "group", "perms", "mountpoint")
	for pool, projects := range POOLSPROJECTS {
		for _, project := range projects {
			filesystems := GetFilesystems(pool, project)
			for _, filesystem := range filesystems.List {
				fmt.Printf("%-18s %8s %-22s %8.2f %8.2f %8.2f %8s %8s %6s %-35s\n", filesystem.Name, filesystem.Pool,
					filesystem.Project, filesystem.Reservation/(1024*1024*1024), filesystem.Quota/(1024*1024*1024),
					filesystem.SpaceTotal/(1024*1024*1024), filesystem.RootUser, filesystem.RootGroup, filesystem.RootPermissions,
					filesystem.MountPoint)
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
				//initiator := strings.Join(lun.InitiatorGroup, "|")
				fmt.Printf("%-16s %8s %-15s %8s %5d %-15s %8s %32s %8.2f %8.2f\n", lun.Name, lun.Pool, lun.Project, lun.Status,
					lun.AssignedNumber, lun.InitiatorGroup, lun.TargetGroup, lun.LunGUID, lun.VolSize/(1024*1204*1024),
					lun.SpaceTotal/(1024*1024*1024))
			}
		}
	}
}
