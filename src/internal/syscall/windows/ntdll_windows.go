package windows

const FileAttributeTagInformation = 35

type NTSTATUS int32

type IO_STATUS_BLOCK struct {
	StatusPointer uintptr
	Information   uintptr
}

//sys	NtQueryInformationFile(handle syscall.Handle, ioStatusBlock *IO_STATUS_BLOCK, info *byte, length uint32, class uint32) (status NTSTATUS) = ntdll.NtQueryInformationFile
//sys	RtlNtStatusToDosError(status NTSTATUS) (dosError syscall.Errno) = ntdll.RtlNtStatusToDosErrorNoTeb
