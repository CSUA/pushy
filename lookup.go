package main

import (
	"fmt"
	"log"
	"os/user"
	"strconv"
	"syscall"
	"unsafe"
)

/*
#include <unistd.h>
#include <sys/types.h>
#include <grp.h>
#include <stdlib.h>
*/
import "C"

type UnknownLookupError string

func (e UnknownLookupError) Error() string {
	return "ERROR: unknown lookup problem " + string(e)
}

func GetUidByName(name string) (uid int, err error) {
	osUser, err := user.Lookup(config.User)
	if err != nil {
		return
	}
	uid, err = strconv.Atoi(osUser.Uid)
	return
}

// TODO: properly attribute http://golang.org/src/pkg/os/user/lookup_unix.go (BSD-style: http://golang.org/LICENSE)
func GetGidByName(name string) (gid int, err error) {
	var group C.struct_group
	var result *C.struct_group

	var bufSize C.long
	bufSize = C.sysconf(C._SC_GETGR_R_SIZE_MAX)
	if bufSize <= 0 || bufSize > 1<<20 {
		log.Fatalf("ERROR: unreasonable _SC_GETGR_R_SIZE_MAX of %d", bufSize)
	}
	buf := C.malloc(C.size_t(bufSize))
	defer C.free(buf)

	var returnValue C.int
	nameC := C.CString(config.Group)
	defer C.free(unsafe.Pointer(nameC))
	returnValue = C.getgrnam_r(nameC,
		&group,
		(*C.char)(buf),
		C.size_t(bufSize),
		&result)
	if returnValue != 0 {
		return -1, fmt.Errorf("ERROR: error looking up group", name, syscall.Errno(returnValue))
	}
	if result == nil {
		return -1, UnknownLookupError(name)
	}
	gid = int(result.gr_gid)
	return
}

func DropPrivileges(username string, groupname string) (err error) {
	uid, err := GetUidByName(username)
	if err != nil {
		log.Fatal("Failed to GetUidByName")
		return
	}
	log.Printf("uid=%d", uid)

	gid, err := GetGidByName(groupname)
	if err != nil {
		log.Fatal("Failed to GetGidByName")
		return
	}
	log.Printf("gid=%d", gid)

	err = syscall.Setgid(gid)
	if err != nil {
		log.Fatal("Failed to SetGid")
		return
	}
	log.Printf("Dropped group privileges")

	err = syscall.Setuid(uid)
	if err != nil {
		return
	}
	log.Printf("Dropped user privileges")
	return
}
