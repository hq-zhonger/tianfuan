package System

import (
	"golang.org/x/sys/windows"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

/*
#include <windows.h>
#include <tlhelp32.h>
#include <stdbool.h>

bool GetSystem(char path[]) {
	HANDLE hToken;
	LUID Luid;
	TOKEN_PRIVILEGES tp;
	OpenProcessToken(GetCurrentProcess(), TOKEN_ADJUST_PRIVILEGES | TOKEN_QUERY, &hToken);
	LookupPrivilegeValue(NULL, SE_DEBUG_NAME, &Luid);
	tp.PrivilegeCount = 1;
	tp.Privileges[0].Luid = Luid;
	tp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED;
	AdjustTokenPrivileges(hToken, false, &tp, sizeof(tp), NULL, NULL);
	CloseHandle(hToken);

	DWORD idL, idW;
	PROCESSENTRY32 pe;
	pe.dwSize = sizeof(PROCESSENTRY32);
	HANDLE hSnapshot = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
	if (Process32First(hSnapshot, &pe)) {
	do {
	if (0 == _stricmp(pe.szExeFile, "lsass.exe")) {
	idL = pe.th32ProcessID;
	} else if (0 == _stricmp(pe.szExeFile, "winlogon.exe")) {
	idW = pe.th32ProcessID;
	}
	} while (Process32Next(hSnapshot, &pe));
	}
	CloseHandle(hSnapshot);

	HANDLE hProcess = OpenProcess(PROCESS_QUERY_INFORMATION, FALSE, idL);
	if (!hProcess) {
	hProcess = OpenProcess(PROCESS_QUERY_INFORMATION, FALSE, idW);
	}
	HANDLE hTokenx;
	OpenProcessToken(hProcess, TOKEN_DUPLICATE, &hTokenx);
	DuplicateTokenEx(hTokenx, MAXIMUM_ALLOWED, NULL, SecurityIdentification, TokenPrimary, &hToken);
	CloseHandle(hProcess);
	CloseHandle(hTokenx);

	STARTUPINFOW si;
	PROCESS_INFORMATION pi;
	ZeroMemory(&si, sizeof(STARTUPINFOW));
	si.cb = sizeof(STARTUPINFOW);
	si.lpDesktop = L"winsta0\\default";
    wchar_t wstr[strlen(path)+1];
    mbstowcs(wstr, path, strlen(path)+1);
	CreateProcessWithTokenW(hToken, LOGON_NETCREDENTIALS_ONLY, NULL,wstr, NORMAL_PRIORITY_CLASS, NULL, NULL, &si, &pi);
	CloseHandle(hToken);
    return true;
}
*/
import "C"

type System struct{}

func (s *System) GetSystem() {
	abs, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	abs += "\\tianfuan.exe"
	abs = strings.ReplaceAll(abs, "\\", "\\\\")
	if IsTrue := C.GetSystem(C.CString(abs)); IsTrue {
		os.Exit(0)
	}
}

func (s *System) CheckAdministrator() {
	if !s.IsAdmin() {
		err := s.RunMeElevated()
		if err != nil {
			os.Exit(0)
		}
		os.Exit(0)
	}
}

func (s *System) CheckSystem() {
	output, err := exec.Command("cmd", "/c", "whoami").Output()
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(output), "system") != true {
		s.GetSystem()
	}
}

// IsAdmin 判断当前权限是否是管理员
func (s *System) IsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}

// RunMeElevated 以管理员身份运行
func (s *System) RunMeElevated() error {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 0 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		return err
	}
	return nil
}
