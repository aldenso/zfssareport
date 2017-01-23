package main

import "github.com/aldenso/zfssareport/model"

var (
	testPools = model.Pools{[]model.Pool{{Status: "online",
		Profile: "mirror",
		Name:    "pool_0",
		PoolUsage: model.PoolUsage{Available: 3.7345023488e+10,
			UsageSnapshots: 38912,
			Used:           1.059328e+06,
			Compression:    1,
			UsageData:      1.0785024e+09,
			Free:           3.4614091264e+10,
			Dedupratio:     100,
			Total:          3.7346082816e+10,
			UsageTotal:     2.1477632e+09},
		Peer:  "00000000-0000-0000-0000-000000000000",
		Owner: "zfs",
		ASN:   "c79d8f07-4146-4ca8-d0a1-ef0d2bcea217"}}}
	testProjects = model.Projects{[]model.Project{{ACLinherit: "restricted",
		ACLMode: "discard", ATime: true, CanonicalName: "pool_0/local/default",
		CheckSum: "fletcher4", Compression: "off", CompressRatio: 100, Copies: 1,
		Creation: "20170108T18:31:33", Dedup: false, DefaultGroup: "other",
		DefaultPermissions: "700", DefaultSparse: false, DefaultUser: "nobody",
		DefaulVolBlockSize: 8192, DefaultVolSize: 0, DefaultGroupQuota: 0,
		DefaultUserQuota: 0, Encryption: "off", Exported: true,
		HREF: "/api/storage/v1/pools/pool_0/projects/default",
		ID:   "dc91107b-462f-3503-0000-000000000000", KeyChangeDate: "",
		KeyStatus: "none", Logbias: "latency", MaxBlockSize: 1048576,
		MountPoint: "/export", Name: "default", Nbmand: false, Nodestroy: false,
		Pool: "pool_0", Quota: 0, ReadOnly: false, RecordSize: 0, Reservation: 0,
		Rstchown: true, SecondaryCache: "all", ShareDAV: "", ShareFTP: "", ShareNFS: "on",
		ShareSFTP: "", ShareSMB: "off", ShareTFTP: "", SnapDir: "hidden",
		ProjectSource: model.ProjectSource{ACLinherit: "inherited",
			ACLMode: "inherited", ATime: "inherited", CheckSum: "inherited",
			Compression: "inherited", Copies: "inherited", Dedup: "inherited",
			Exported: "local", KeyChangeDate: "inherited", Logbias: "inherited",
			MaxBlockSize: "inherited", MountPoint: "local", Nbmand: "inherited",
			ReadOnly: "inherited", RecordSize: "", Reservation: "local",
			RRSRCActions: "local", Rstchown: "inherited", SecondaryCache: "inherited",
			ShareDAV: "local", ShareFTP: "local", ShareNFS: "local", ShareSFTP: "local",
			ShareSMB: "local", ShareTFTP: "local", SnapDir: "inherited", VScan: "inherited"},
		SpaceAvailable: 3.4614091264e+10, SpaceData: 31744, SpaceSnapShots: 0,
		SpaceTotal: 31744, SpaceUnusedRes: 0, SpaceUnusedResShares: 0, VScan: false}}}

	testFilesystems = []model.Filesystems{model.Filesystems{[]model.Filesystem{model.Filesystem{
		ACLinherit: "restricted", ACLMode: "discard", ATime: true, CanonicalName: "pool_0/local/test1/test1",
		CaseSensitivity: "mixed", CheckSum: "fletcher4", Compression: "off", CompressRatio: 100, Copies: 1,
		Creation: "20170116T19:11:01", Dedup: false, Encryption: "off", Exported: true,
		HREF: "api/storage/v1/pools/pool_0/projects/test1/filesystems/test1", ID: "81574547-b532-a8c9-0000-000000000000",
		KeyChangeDate: "", KeyStatus: "none", Logbias: "latency", MaxBlockSize: 1048576, MountPoint: "export/test1/test1",
		Name: "test1", Nbmand: false, Nodestroy: false, Normalization: "none", Pool: "pool_0", Project: "test1",
		Quota: 1.44703488e+08, QuotaSnap: true, ReadOnly: false, RecordSize: 0, Reservation: 0, ReservationSnap: true,
		RootGroup: "other", RootPermissions: "700", RootUser: "nobody", Rstchown: true, SecondaryCache: "all",
		Shadow: "none", ShareDAV: "", ShareFTP: "", ShareNFS: "on", ShareSFTP: "", ShareSMB: "off", ShareTFTP: "",
		SnapDir: "hidden", SnapLabel: "",
		FSSource: model.FSSource{
			ACLinherit: "inherited", ACLMode: "inherited", ATime: "inherited", CheckSum: "inherited", Compression: "inherited",
			Copies: "inherited", Dedup: "inherited", Exported: "inherited", KeyChangeDate: "inherited", Logbias: "inherited",
			MaxBlockSize: "inherited", MountPoint: "inherited", Nbmand: "inherited", ReadOnly: "inherited", RecordSize: "inherited",
			Reservation: "local", RRSRCActions: "inherited", Rstchown: "inherited", SecondaryCache: "inherited", ShareDAV: "inherited",
			ShareFTP: "inherited", ShareNFS: "inherited", ShareSFTP: "inherited", ShareSMB: "inherited", ShareTFTP: "inherited",
			SnapDir: "inherited", VScan: "inherited"},
		SpaceAvailable: 1.44631808e+08, SpaceData: 32768, SpaceSnapShots: 38912, SpaceTotal: 71680, SpaceUnusedRes: 0,
		UTF8Only: true, VScan: false},
	}}}
	/*testFilesystem = Filesystems{
	[]Filesystem{Filesystem{
		ACLinherit: "restricted", ACLMode: "discard", ATime: true, CanonicalName: "pool_0/local/test1/test1",
		CaseSensitivity: "mixed", CheckSum: "fletcher4", Compression: "off", CompressRatio: 100, Copies: 1,
		Creation: "20170116T19:11:01", Dedup: false, Encryption: "off", Exported: true,
		HREF: "api/storage/v1/pools/pool_0/projects/test1/filesystems/test1", ID: "81574547-b532-a8c9-0000-000000000000",
		KeyChangeDate: "", KeyStatus: "none", Logbias: "latency", MaxBlockSize: 1048576, MountPoint: "export/test1/test1",
		Name: "test1", Nbmand: false, Nodestroy: false, Normalization: "none", Pool: "pool_0", Project: "test1",
		Quota: 1.44703488e+08, QuotaSnap: true, ReadOnly: false, RecordSize: 0, Reservation: 0, ReservationSnap: true,
		RootGroup: "other", RootPermissions: "700", RootUser: "nobody", Rstchown: true, SecondaryCache: "all",
		Shadow: "none", ShareDAV: "", ShareFTP: "", ShareNFS: "on", ShareSFTP: "", ShareSMB: "off", ShareTFTP: "",
		SnapDir: "hidden", SnapLabel: "",
		FSSource: FSSource{
			ACLinherit: "inherited", ACLMode: "inherited", ATime: "inherited", CheckSum: "inherited", Compression: "inherited",
			Copies: "inherited", Dedup: "inherited", Exported: "inherited", KeyChangeDate: "inherited", Logbias: "inherited",
			MaxBlockSize: "inherited", MountPoint: "inherited", Nbmand: "inherited", ReadOnly: "inherited", RecordSize: "inherited",
			Reservation: "local", RRSRCActions: "inherited", Rstchown: "inherited", SecondaryCache: "inherited", ShareDAV: "inherited",
			ShareFTP: "inherited", ShareNFS: "inherited", ShareSFTP: "inherited", ShareSMB: "inherited", ShareTFTP: "inherited",
			SnapDir: "inherited", VScan: "inherited"},
		SpaceAvailable: 1.44631808e+08, SpaceData: 32768, SpaceSnapShots: 38912, SpaceTotal: 71680, SpaceUnusedRes: 0,
		UTF8Only: true, VScan: false},
	}}*/
)

func Example_Header() {
	msg := "this is a test"
	Header(msg)
	//Output:
	//#####################################################################################################################
	//## this is a test                                                                                                  ##
	//#####################################################################################################################
}

func Example_PrintPools() {
	PrintPools(testPools, FsMem)
	//Output:
	//#####################################################################################################################
	//## Pools information                                                                                               ##
	//#####################################################################################################################
	//Name   Status   Profile                              Total(GB)  Avail(GB)   Free(GB)  UData(GB) USnaps(GB) UTotal(GB)
	//=====================================================================================================================
	//pool_0 online   mirror                                   34.78      34.78      32.24       1.00       0.00       2.00
}

func Example_PrintProjects() {
	pmap := make(map[string]model.Projects)
	for _, pool := range testPools.List {
		pmap[pool.Name] = testProjects
	}
	PrintProjects(pmap, FsMem)
	//Output:
	//#####################################################################################################################
	//## Projects information                                                                                            ##
	//#####################################################################################################################
	//Project         Reserv(GB) Quota(GB)  Pool     STotal(GB)                               mountpoint
	//=====================================================================================================================
	//default         0.00       0.00       pool_0         0.00                                  /export
}

func Example_PrintFilesystems() {
	pmap := make(map[string]model.Projects)
	for _, pool := range testPools.List {
		pmap[pool.Name] = testProjects
	}
	PrintFilesystems(testFilesystems, FsMem)
	//Output:
	//#####################################################################################################################
	//## Filesystems information                                                                                         ##
	//#####################################################################################################################
	//Filesystem   Pool     Project         Reser(GB) Quota(GB) Total(GB)     user    group perms      mountpoint
	//=====================================================================================================================
	//test1        pool_0   test1                0.00      0.13      0.00   nobody    other  700 export/test1/test1
}
