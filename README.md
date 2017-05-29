# zfssareport

zfs storage appliance report.

Create a config file (config.yml):

```bash
./zfssareport -template
config.yml created.
```

```yml
# ZFSSA REPORT CONFIG"
ip: 192.168.56.150
user: root
password: password
```

or

```yml
# ZFSSA REPORT CONFIG"
ip: 192.168.56.150
user: root
password:
```

If the password is empty, then the program will ask for it.

Output example:

```bash
./zfssareport -t config.yml
Please enter your password:
#####################################################################################################################
## ZFS Storage Appliance Version                                                                                   ##
#####################################################################################################################
nodename        version                  product                 csn                                        boot_time
=====================================================================================================================
ZFSNODE-001     2013.06.05.6.8,1-1.1     Sun ZFS Storage 7420    YYYXXXZZZZ   Sun Nov 17 2016 08:42:31 GMT+0000 (UTC)
#####################################################################################################################
## ZFS Storage Appliance Cluster                                                                                   ##
#####################################################################################################################
description                    peer description               peer Hostname   peer state
=====================================================================================================================
Active (takeover completed)    Ready (waiting for failback)   ZFSNODE-002     AKCS_STRIPPED
#####################################################################################################################
resource        type            owner           details                   label
=====================================================================================================================
igb1            private         ZFSNODE-002     [192.168.56.150 ] Untitled Interface
ipmp1           singleton       ZFSNODE-001     [192.168.32.150 ] ipmp1
ipmp2           singleton       ZFSNODE-002     [192.168.32.150 ] Ipmp2
pool_0          singleton       ZFSNODE-002     [20.0T          ]
pool_1          singleton       ZFSNODE-002     [5.07T          ]
pool_2          singleton       ZFSNODE-002     [37.4T          ]
pool_3          singleton       ZFSNODE-001     [               ]
pool_4          singleton       ZFSNODE-001     [               ]
#####################################################################################################################
## Network Interfaces information                                                                                  ##
#####################################################################################################################
interface       class    links                     label                v4addrs                   state
=====================================================================================================================
igb1            ip       igb1                      Untitled Interface   192.168.56.150/24           up
ipmp1           ipmp     ixgbe505000,ixgbe505002   ipmp0                192.168.32.150/24         offline
ipmp2           ipmp     ixgbe505001,ixgbe505003   Ipmp1                192.168.32.151/24           up
ixgbe505000     ip       ixgbe505000               ixgbe0_505           0.0.0.0/0                 offline
ixgbe505001     ip       ixgbe505001               ixgbe1_505           0.0.0.0/0                   up
ixgbe505002     ip       ixgbe505002               ixgbe2_505           0.0.0.0/0                 offline
ixgbe505003     ip       ixgbe505003               ixgbe3_505           0.0.0.0/0                   up
#####################################################################################################################
## Pools information                                                                                               ##
#####################################################################################################################
Name   Status   Profile                              Total(GB)  Avail(GB)   Free(GB)  UData(GB) USnaps(GB) UTotal(GB)
=====================================================================================================================
pool_3 exported                                           0.00       0.00       0.00       0.00       0.00       0.00
pool_4 exported                                           0.00       0.00       0.00       0.00       0.00       0.00
pool_0 degraded mirror_nspf:log_stripe:cache_stripe   20800.00    9112.32      49.44   20421.04       0.00   20421.04
pool_2 online   raidz2:log_stripe:cache_stripe        50176.00   28873.33    4171.45   23982.22       0.00   34374.65
pool_1 online   mirror:log_mirror                      5272.00    1935.20    1501.69    3685.18       0.00    3685.18
#####################################################################################################################
## Projects information                                                                                            ##
#####################################################################################################################
Project         Reserv(GB) Quota(GB)  Pool     STotal(GB)                               mountpoint
=====================================================================================================================
apps_prod        0.00       730.00     pool_1       229.93                         /export/apps-prod
apps-dev         0.00       0.00       pool_1       234.08                         /export/apps-dev
.
.
.
#####################################################################################################################
## LUNS information                                                                                                ##
#####################################################################################################################
LUN      Pool     Project            Status  Num InitiatorGrp                                GUID VolS(GB) Total(GB)
=====================================================================================================================
VS00     pool_3   VMWPROD00          online  10  [EVVEP000-009   ] 600144YYYYYYYXXXXXXXXXZZZZZZZZZZ     1.70     2.00
VS01     pool_3   VMWPROD00          online  11  [EVVEP000-009   ] 600144YYYYYYYXXXXXXXXXZZZZZZZZZZ     1.70     2.00
VS02     pool_3   VMWPROD00          online  12  [EVVEP000-009   ] 600144YYYYYYYXXXXXXXXXZZZZZZZZZZ     1.70     2.00
.
.
.
+++ results file '192.168.56.150_2017-01-24T214531-0400' created +++

############# DONE in 1m9.5939805s #############
```

The csv reports are placed in a zip file, named after de ip/hostname and the time.

```text
192.168.56.150_2017-01-24T214531-0400.zip
```

Issue: When dealing with more than one initiator group for a lun, the zfs api changes AssignedNumber from int to []int, thus we get json: cannot unmarshal.

You can use silent mode (-silent) to suppress the items output and just know the steps running.

```text
./zfssareport -t config.yml -silent
getting version info.
getting cluster info.
getting chassis info.
getting problems info.
getting network interfaces info.
getting network datalinks info.
getting network devices info.
getting pools info.
getting projects info.
getting filesystems info.
getting luns info.
getting FC initiators info.
getting FC initiators groups info.
getting FC targets info.
getting ISCSI initiators info.
getting ISCSI initiators groups info.

+++ results file '10.246.2.83_2017-05-29T080858-0400' created +++

############# DONE in 1m47.0031203s #############
```