package windows

import "syscall"

const (
	FileModeInformation         = 16
	FileAttributeTagInformation = 35

	FILE_SYNCHRONOUS_IO_NONALERT uint32 = 0x00000020

	OBJ_CASE_INSENSITIVE uint32 = 0x00000040

	FILE_CREATE uint32 = 0x00000002
)

type NTSTATUS int32

type IO_STATUS_BLOCK struct {
	StatusPointer uintptr
	Information   uintptr
}

type OBJECT_ATTRIBUTES struct {
	Length                   uint32
	RootDirectory            syscall.Handle
	ObjectName               *UNICODE_STRING
	Attributes               uint32
	SecurityDescriptor       *byte
	SecurityQualityOfService *byte
}

type UNICODE_STRING struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *uint16
}

//sys	NtCreateNamedPipeFile(handle *syscall.Handle, desiredAccess uint32, objectAttributes *OBJECT_ATTRIBUTES, ioStatusBlock *IO_STATUS_BLOCK, shareAccess uint32, createDisposition uint32, createOptions uint32, writeModeMessage uint32, readModeMessage uint32, nonBlocking uint32, maxInstances uint32, inBufferSize uint32, outBufferSize uint32, defaultTimeout *int64) (status NTSTATUS) = ntdll.NtCreateNamedPipeFile
//sys	NtOpenFile(handle *syscall.Handle, desiredAccess uint32, objectAttributes *OBJECT_ATTRIBUTES, ioStatusBlock *IO_STATUS_BLOCK, shareAccess uint32, openOptions uint32) (status NTSTATUS) = ntdll.NtOpenFile
//sys	NtQueryInformationFile(handle syscall.Handle, ioStatusBlock *IO_STATUS_BLOCK, info *byte, length uint32, class uint32) (status NTSTATUS) = ntdll.NtQueryInformationFile
//sys	RtlNtStatusToDosError(status NTSTATUS) (dosError syscall.Errno) = ntdll.RtlNtStatusToDosErrorNoTeb
