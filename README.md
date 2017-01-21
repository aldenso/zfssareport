# zfssareport
zfs storage appliance report.

TODO: create csv files to write all values for pools, projects, filesystems, luns, etc. 

Issue: When dealing with more than one initiator group for a lun, the zfs api changes AssignedNumber from int to []int, thus we get json: cannot unmarshal. 