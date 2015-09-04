package demo

import (
	"fmt"
	"github.com/pivotal-golang/bytefmt"
	"syscall"
)

/*
Filesystem types:

              ADFS_SUPER_MAGIC      0xadf5
              AFFS_SUPER_MAGIC      0xADFF
              BDEVFS_MAGIC          0x62646576
              BEFS_SUPER_MAGIC      0x42465331
              BFS_MAGIC             0x1BADFACE
              BINFMTFS_MAGIC        0x42494e4d
              BTRFS_SUPER_MAGIC     0x9123683E
              CGROUP_SUPER_MAGIC    0x27e0eb
              CIFS_MAGIC_NUMBER     0xFF534D42
              CODA_SUPER_MAGIC      0x73757245
              COH_SUPER_MAGIC       0x012FF7B7
              CRAMFS_MAGIC          0x28cd3d45
              DEBUGFS_MAGIC         0x64626720
              DEVFS_SUPER_MAGIC     0x1373
              DEVPTS_SUPER_MAGIC    0x1cd1
              EFIVARFS_MAGIC        0xde5e81e4
              EFS_SUPER_MAGIC       0x00414A53
              EXT_SUPER_MAGIC       0x137D
              EXT2_OLD_SUPER_MAGIC  0xEF51
              EXT2_SUPER_MAGIC      0xEF53
              EXT3_SUPER_MAGIC      0xEF53
              EXT4_SUPER_MAGIC      0xEF53
              FUSE_SUPER_MAGIC      0x65735546
              FUTEXFS_SUPER_MAGIC   0xBAD1DEA
              HFS_SUPER_MAGIC       0x4244
              HOSTFS_SUPER_MAGIC    0x00c0ffee
              HPFS_SUPER_MAGIC      0xF995E849
              HUGETLBFS_MAGIC       0x958458f6
              ISOFS_SUPER_MAGIC     0x9660
              JFFS2_SUPER_MAGIC     0x72b6
              JFS_SUPER_MAGIC       0x3153464a
              MINIX_SUPER_MAGIC     0x137F
              MINIX_SUPER_MAGIC2    0x138F
              MINIX2_SUPER_MAGIC    0x2468
              MINIX2_SUPER_MAGIC2   0x2478
              MINIX3_SUPER_MAGIC    0x4d5a
              MQUEUE_MAGIC          0x19800202
              MSDOS_SUPER_MAGIC     0x4d44
              NCP_SUPER_MAGIC       0x564c
              NFS_SUPER_MAGIC       0x6969
              NILFS_SUPER_MAGIC     0x3434
              NTFS_SB_MAGIC         0x5346544e
              OCFS2_SUPER_MAGIC     0x7461636f
              OPENPROM_SUPER_MAGIC  0x9fa1
              PIPEFS_MAGIC          0x50495045
              PROC_SUPER_MAGIC      0x9fa0
              PSTOREFS_MAGIC        0x6165676C
              QNX4_SUPER_MAGIC      0x002f
              QNX6_SUPER_MAGIC      0x68191122
              RAMFS_MAGIC           0x858458f6
              REISERFS_SUPER_MAGIC  0x52654973
              ROMFS_MAGIC           0x7275
              SELINUX_MAGIC         0xf97cff8c
              SMACK_MAGIC           0x43415d53
              SMB_SUPER_MAGIC       0x517B
              SOCKFS_MAGIC          0x534F434B
              SQUASHFS_MAGIC        0x73717368
              SYSFS_MAGIC           0x62656572
              SYSV2_SUPER_MAGIC     0x012FF7B6
              SYSV4_SUPER_MAGIC     0x012FF7B5
              TMPFS_MAGIC           0x01021994
              UDF_SUPER_MAGIC       0x15013346
              UFS_MAGIC             0x00011954
              USBDEVICE_SUPER_MAGIC 0x9fa2
              V9FS_MAGIC            0x01021997
              VXFS_SUPER_MAGIC      0xa501FCF5
              XENFS_SUPER_MAGIC     0xabba1974
              XENIX_SUPER_MAGIC     0x012FF7B4
              XFS_SUPER_MAGIC       0x58465342
              _XIAFS_SUPER_MAGIC    0x012FD16D
*/
type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	fmt.Print("",
		"Type:", fs.Type, "\n", /* Type of filesystem (see below) */
		"Bsize:", fs.Bsize, "\n", /* Optimal transfer block size */
		"Blocks:", fs.Blocks, "\n", /* Total data blocks in filesystem */
		"Bfree:", fs.Bfree, "\n", /* Free blocks in filesystem */
		"Bavail:", fs.Bavail, "\n", /* Free blocks available to unprivileged user*/
		"Files:", fs.Files, "\n", /* Total file nodes in filesystem */
		"Ffree:", fs.Ffree, "\n", /* Free file nodes in filesystem */
		//"Fsid:", fs.Bsize, "\n",/* Filesystem ID */
		"Namelen:", fs.Namelen, "\n", /* Maximum length of filenames */
		"Frsize:", fs.Frsize, "\n", /* Fragment size (since Linux 2.6) */
		"Flags:", fs.Flags, "\n", /* Mount flags of filesystem (since Linux 2.6.36) */
		//"Spare:", fs.Bsize, "\n", /* Padding bytes reserved for future use */
	)

	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	fmt.Println("",
		"Total:", disk.All, bytefmt.ByteSize(disk.All), "\n",
		"Used:", disk.Used, bytefmt.ByteSize(disk.Used), "\n",
		"Free:", disk.Free, bytefmt.ByteSize(disk.Free))
	return
}
