package main

//Pools struct for ZFS pools.
type Pools struct {
	List []Pool `json:"pools"`
}

//Pool struct for ZFS pool.
type Pool struct {
	Status    string `json:"status"`
	Profile   string `json:"profile"`
	Name      string `json:"name"`
	PoolUsage `json:"usage"`
	Peer      string `json:"peer"`
	Owner     string `json:"owner"`
	ASN       string `json:"asn"`
}

//PoolUsage struct for Usage in Pool struct.
type PoolUsage struct {
	Available      float64 `json:"available"`
	UsageSnapshots float64 `json:"usage_snapshots"`
	Used           float64 `json:"used"`
	Compression    int     `json:"compression"`
	UsageData      float64 `json:"usage_data"`
	Free           float64 `json:"free"`
	Dedupratio     int     `json:"dedupratio"`
	Total          float64 `json:"total"`
	UsageTotal     float64 `json:"usage_total"`
}

//Projects struct for Projects in Pool.
type Projects struct {
	List []Project `json:"projects"`
}

//Project struct for project in Pool.
type Project struct {
	ACLinherit    string  `json:"aclinherit"`
	ACLMode       string  `json:"aclmode"`
	ATime         bool    `json:"atime"`
	CanonicalName string  `json:"canonical_name"`
	CheckSum      string  `json:"checksum"`
	Compression   string  `json:"compression"`
	CompressRatio float64 `json:"compressratio"`
	Copies        int     `json:"copies"`
	Creation      string  `json:"creation"`
	//Creation           time.Time `json:"creation"`
	Dedup                bool    `json:"dedup"`
	DefaultGroup         string  `json:"default_group"`
	DefaultPermissions   string  `json:"default_permissions"`
	DefaultSparse        bool    `json:"default_sparce"`
	DefaultUser          string  `json:"default_user"`
	DefaulVolBlockSize   float64 `json:"default_volblocksize"`
	DefaultVolSize       float64 `json:"default_volsize"`
	DefaultGroupQuota    float64 `json:"defaultgroupquota"`
	DefaultUserQuota     float64 `json:"defaultuserquota"`
	Encryption           string  `json:"encryption"`
	Exported             bool    `json:"exported"`
	HREF                 string  `json:"href"`
	ID                   string  `json:"id"`
	KeyChangeDate        string  `json:"keychangedate,omitempty"`
	KeyStatus            string  `json:"keystatus"`
	Logbias              string  `json:"logbias"`
	MaxBlockSize         int     `json:"maxblocksize"`
	MountPoint           string  `json:"mountpoint"`
	Name                 string  `json:"name"`
	Nbmand               bool    `json:"nbmand"`
	Nodestroy            bool    `json:"nodestroy"`
	Pool                 string  `json:"pool"`
	Quota                float64 `json:"quota"`
	ReadOnly             bool    `json:"readonly"`
	RecordSize           float64 `json:"record_size"`
	Reservation          float64 `json:"reservation"`
	Rstchown             bool    `json:"rstchown"`
	SecondaryCache       string  `json:"secondarycache"`
	ShareDAV             string  `json:"sharedav,omitempty"`
	ShareFTP             string  `json:"shareftp,omitempty"`
	ShareNFS             string  `json:"sharenfs,omitempty"`
	ShareSFTP            string  `json:"sharesftp,omitempty"`
	ShareSMB             string  `json:"sharesmb,omitempty"`
	ShareTFTP            string  `json:"sharetftp,omitempty"`
	SnapDir              string  `json:"snapdir"`
	SnapLabel            string  `json:"snaplabel,omitempty"`
	ProjectSource        `json:"source"`
	SpaceAvailable       float64 `json:"space_available"`
	SpaceData            float64 `json:"space_data"`
	SpaceSnapShots       float64 `json:"space_snapshots"`
	SpaceTotal           float64 `json:"space_total"`
	SpaceUnusedRes       float64 `json:"space_unused_res"`
	SpaceUnusedResShares float64 `json:"space_unused_res_shares"`
	VScan                bool    `json:"vscan"`
}

//ProjectSource struct for source in Project struct.
type ProjectSource struct {
	ACLinherit     string `json:"aclinherit"`
	ACLMode        string `json:"aclmode"`
	ATime          string `json:"atime"`
	CheckSum       string `json:"checksum"`
	Compression    string `json:"compression"`
	Copies         string `json:"copies"`
	Dedup          string `json:"dedup"`
	Exported       string `json:"exported"`
	KeyChangeDate  string `json:"keychangedate"`
	Logbias        string `json:"logbias"`
	MaxBlockSize   string `json:"maxblocksize"`
	MountPoint     string `json:"mountpoint"`
	Nbmand         string `json:"nbmand"`
	ReadOnly       string `json:"readonly"`
	RecordSize     string `json:"record_size"`
	Reservation    string `json:"reservation"`
	RRSRCActions   string `json:"rrsrc_actions"`
	Rstchown       string `json:"rstchown"`
	SecondaryCache string `json:"secondarycache"`
	ShareDAV       string `json:"sharedav"`
	ShareFTP       string `json:"shareftp"`
	ShareNFS       string `json:"sharenfs"`
	ShareSFTP      string `json:"sharesftp"`
	ShareSMB       string `json:"sharesmb"`
	ShareTFTP      string `json:"sharetftp"`
	SnapDir        string `json:"snapdir"`
	VScan          string `json:"vscan"`
}

//Filesystems struct for Filesystems in a Project.
type Filesystems struct {
	List []Filesystem `json:"Filesystems"`
}

//Filesystem struct for Filesystem in a Project.
type Filesystem struct {
	ACLinherit      string  `json:"aclinherit"`
	ACLMode         string  `json:"aclmode"`
	ATime           bool    `json:"atime"`
	CanonicalName   string  `json:"canonical_name"`
	CaseSensitivity string  `json:"casesensitivity"`
	CheckSum        string  `json:"checksum"`
	Compression     string  `json:"compression"`
	CompressRatio   float64 `json:"compressratio"`
	Copies          int     `json:"copies"`
	Creation        string  `json:"creation"`
	//Creation           time.Time `json:"creation"`
	Dedup           bool    `json:"dedup"`
	Encryption      string  `json:"encryption"`
	Exported        bool    `json:"exported"`
	HREF            string  `json:"href"`
	ID              string  `json:"id"`
	KeyChangeDate   string  `json:"keychangedate,omitempty"`
	KeyStatus       string  `json:"keystatus"`
	Logbias         string  `json:"logbias"`
	MaxBlockSize    int     `json:"maxblocksize"`
	MountPoint      string  `json:"mountpoint"`
	Name            string  `json:"name"`
	Nbmand          bool    `json:"nbmand"`
	Nodestroy       bool    `json:"nodestroy"`
	Normalization   string  `json:"normalization"`
	Pool            string  `json:"pool"`
	Project         string  `json:"project"`
	Quota           float64 `json:"quota"`
	QuotaSnap       bool    `json:"quota_snap"`
	ReadOnly        bool    `json:"readonly"`
	RecordSize      float64 `json:"record_size"`
	Reservation     float64 `json:"reservation"`
	ReservationSnap bool    `json:"reservation_snap"`
	RootGroup       string  `json:"root_group"`
	RootPermissions string  `json:"root_permissions"`
	RootUser        string  `json:"root_user"`
	Rstchown        bool    `json:"rstchown"`
	SecondaryCache  string  `json:"secondarycache"`
	Shadow          string  `json:"shadow"`
	ShareDAV        string  `json:"sharedav,omitempty"`
	ShareFTP        string  `json:"shareftp,omitempty"`
	ShareNFS        string  `json:"sharenfs,omitempty"`
	ShareSFTP       string  `json:"sharesftp,omitempty"`
	ShareSMB        string  `json:"sharesmb,omitempty"`
	ShareTFTP       string  `json:"sharetftp,omitempty"`
	SnapDir         string  `json:"snapdir"`
	SnapLabel       string  `json:"snaplabel,omitempty"`
	FSSource        `json:"source"`
	SpaceAvailable  float64 `json:"space_available"`
	SpaceData       float64 `json:"space_data"`
	SpaceSnapShots  float64 `json:"space_snapshots"`
	SpaceTotal      float64 `json:"space_total"`
	SpaceUnusedRes  float64 `json:"space_unused_res"`
	UTF8Only        bool    `json:"utf8only"`
	VScan           bool    `json:"vscan"`
}

//FSSource struct for source in Filesystem struct.
type FSSource struct {
	ACLinherit     string `json:"aclinherit"`
	ACLMode        string `json:"aclmode"`
	ATime          string `json:"atime"`
	CheckSum       string `json:"checksum"`
	Compression    string `json:"compression"`
	Copies         string `json:"copies"`
	Dedup          string `json:"dedup"`
	Exported       string `json:"exported"`
	KeyChangeDate  string `json:"keychangedate"`
	Logbias        string `json:"logbias"`
	MaxBlockSize   string `json:"maxblocksize"`
	MountPoint     string `json:"mountpoint"`
	Nbmand         string `json:"nbmand"`
	ReadOnly       string `json:"readonly"`
	RecordSize     string `json:"record_size"`
	Reservation    string `json:"reservation"`
	RRSRCActions   string `json:"rrsrc_actions"`
	Rstchown       string `json:"rstchown"`
	SecondaryCache string `json:"secondarycache"`
	ShareDAV       string `json:"sharedav"`
	ShareFTP       string `json:"shareftp"`
	ShareNFS       string `json:"sharenfs"`
	ShareSFTP      string `json:"sharesftp"`
	ShareSMB       string `json:"sharesmb"`
	ShareTFTP      string `json:"sharetftp"`
	SnapDir        string `json:"snapdir"`
	VScan          string `json:"vscan"`
}

//LUNS struct for luns in a project.
type LUNS struct {
	List []LUN `json:"luns"`
}

//LUN struct for lun in project.
type LUN struct {
	AssignedNumber int     `json:"assignednumber"`
	CanonicalName  string  `json:"canonical_name"`
	Checksum       string  `json:"checksum"`
	Compression    string  `json:"compression"`
	CompressRatio  float64 `json:"compressratio"`
	Copies         int     `json:"copies"`
	Creation       string  `json:"creation"`
	//Creation           time.Time `json:"creation"`
	Dedup          bool     `json:"dedup"`
	Encryption     string   `json:"encryption"`
	Exported       bool     `json:"exported"`
	FixedNumber    bool     `json:"fixednumber"`
	HREF           string   `json:"href"`
	ID             string   `json:"id"`
	InitiatorGroup []string `json:"initiatorgroup"`
	KeyChangeDate  string   `json:"keychangedate,omitempty"`
	KeyStatus      string   `json:"keystatus"`
	Logbias        string   `json:"logbias"`
	LunGUID        string   `json:"lunguid"`
	LUNumber       string   `json:"lunumber"`
	MaxBlockSize   int      `json:"maxblocksize"`
	Name           string   `json:"name"`
	Nodestroy      bool     `json:"nodestroy"`
	Pool           string   `json:"pool"`
	Project        string   `json:"project"`
	SecondaryCache string   `json:"secondarycache"`
	SnapLabel      string   `json:"snaplabel,omitempty"`
	Source         struct {
		CheckSum       string `json:"checksum"`
		Compression    string `json:"compression"`
		Copies         string `json:"copies"`
		Dedup          string `json:"dedup"`
		Encryption     string `json:"encryption"`
		Exported       string `json:"exported"`
		KeyChangeDate  string `json:"keychangedate"`
		Logbias        string `json:"logbias"`
		MaxBlockSize   string `json:"maxblocksize"`
		RRSRCActions   string `json:"rrsrc_actions"`
		SecondaryCache string `json:"secondarycache"`
	}
	SpaceAvailable float64 `json:"space_available"`
	SpaceData      float64 `json:"space_data"`
	SpaceSnapShots float64 `json:"space_snapshots"`
	SpaceTotal     float64 `json:"space_total"`
	Sparse         bool    `json:"sparse"`
	Status         string  `json:"status"`
	TargetGroup    string  `json:"targetgroup"`
	VolBlockSize   int     `json:"volblocksize"`
	VolSize        float64 `json:"volsize"`
	WriteCache     bool    `json:"writecache"`
}
