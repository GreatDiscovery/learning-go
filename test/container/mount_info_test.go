package container

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sys/unix"
	"testing"
)

// Info reveals information about a particular mounted filesystem. This
// struct is populated from the content in the /proc/<pid>/mountinfo file.
type Info struct {
	// ID is a unique identifier of the mount (may be reused after umount).
	ID int

	// Parent is the ID of the parent mount (or of self for the root
	// of this mount namespace's mount tree).
	Parent int

	// Major and Minor are the major and the minor components of the Dev
	// field of unix.Stat_t structure returned by unix.*Stat calls for
	// files on this filesystem.
	Major, Minor int

	// Root is the pathname of the directory in the filesystem which forms
	// the root of this mount.
	Root string

	// Mountpoint is the pathname of the mount point relative to the
	// process's root directory.
	Mountpoint string

	// Options is a comma-separated list of mount options.
	Options string

	// Optional are zero or more fields of the form "tag[:value]",
	// separated by a space.  Currently, the possible optional fields are
	// "shared", "master", "propagate_from", and "unbindable". For more
	// information, see mount_namespaces(7) Linux man page.
	Optional string

	// FSType is the filesystem type in the form "type[.subtype]".
	FSType string

	// Source is filesystem-specific information, or "none".
	Source string

	// VFSOptions is a comma-separated list of superblock options.
	VFSOptions string
}

func (i *Info) String() string {
	marshal, _ := json.Marshal(i)
	return string(marshal)
}

func TestGetMountInfo(t *testing.T) {
	count, err := unix.Getfsstat(nil, unix.MNT_WAIT)
	if err != nil {
		panic(err)
	}

	entries := make([]unix.Statfs_t, count)
	_, err = unix.Getfsstat(entries, unix.MNT_WAIT)
	if err != nil {
		panic(err)
	}

	var out []*Info
	for _, entry := range entries {
		mountinfo := getMountinfo(&entry)
		out = append(out, mountinfo)
	}
	for _, info := range out {
		fmt.Printf("out=%+v\n", info)
	}
}

func getMountinfo(entry *unix.Statfs_t) *Info {
	return &Info{
		Mountpoint: unix.ByteSliceToString(entry.Mntonname[:]),
		FSType:     unix.ByteSliceToString(entry.Fstypename[:]),
		Source:     unix.ByteSliceToString(entry.Mntfromname[:]),
	}
}
