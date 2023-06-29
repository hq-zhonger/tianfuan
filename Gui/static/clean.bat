@echo off
chcp 65001
echo 正在清理windows常见垃圾
echo.

:: 清理临时文件夹
echo 正在清理临时文件夹
del /q/f/s %temp%\*.*
for /d %%i in ("%temp%\*") do rd /q/s "%%i"
echo.

:: 清理回收站
echo 正在清理回收站
rd /q/s c:\$Recycle.bin
echo.

:: 清理IE缓存
echo 正在清理IE浏览器
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 8
echo.

:: 清理DNS缓存
echo 正在清理DNS缓存
ipconfig /flushdns
echo.

:: 清理系统日志
echo 正在清理系统日志
wevtutil.exe cl System
wevtutil.exe cl Application
echo.


echo 正在清理系统垃圾文件
		
del /f /s /q %systemdrive%\*.tmp
del /f /s /q %systemdrive%\*._mp
del /f /s /q %systemdrive%\*.log
del /f /s /q %systemdrive%\*.gid
del /f /s /q %systemdrive%\*.chk

del /f /s /q %systemdrive%\*.old

del /f /s /q %systemdrive%\recycled\*.*

del /f /s /q %windir%\*.bak

del /f /s /q %windir%\prefetch\*.*

rd /s /q %windir%\temp & md %windir%\temp

del /f /q %userprofile%\小甜饼s\*.*

del /f /q %userprofile%\recent\*.*

del /f /s /q "%userprofile%\Local Settings\Temporary Internet Files\*.*"

del /f /s /q "%userprofile%\Local Settings\Temp\*.*"

del /f /s /q "%userprofile%\recent\*.*"

@echo off
  title @echo off
  color 2
  echo.
  echo.
  echo 请不要关闭此窗口
  echo.
  echo 开始清理垃圾文件 请稍后...
  echo.
  echo 正在清理Thumbs.db数据库文件...
  del c:\\Thumbs.db /f/s/q/a
  del d:\\Thumbs.db /f/s/q/a
  del e:\\Thumbs.db /f/s/q/a
  del f:\\Thumbs.db /f/s/q/a
  del g:\\Thumbs.db /f/s/q/a
  del h:\\Thumbs.db /f/s/q/a
  del i:\\Thumbs.db /f/s/q/a
  echo.
  echo.
  echo.
  echo 正在清理系统分区根目录下tmp文件，请稍等......
  del /f /s /q %systemdrive%\\*.tmp
  echo.
  echo 清理系统分区根目录下tmp文件完成！
  echo.
  echo 正在清理系统分区根目录下_mp文件，请稍等......
  del /f /s /q %systemdrive%\\*._mp
  echo.
  echo 清理系统分区根目录下_mp文件完成！
  echo.
  echo 正在清理系统分区根目录下日志文件，请稍等......
  del /f /s /q %systemdrive%\\*.log
  echo.
  echo 清理系统分区根目录下日志文件完成！
  echo.
  echo 正在清理系统分区根目录下gid文件，请稍等......
  del /f /s /q %systemdrive%\\*.gid
  echo.
  echo 清理系统分区根目录下gid文件完成！
  echo.
  echo 正在清理系统分区根目录下chk文件，请稍等......
  del /f /s /q %systemdrive%\\*.chk
  echo.
  echo 清理系统分区根目录下chk文件完成！
  echo.
  echo 正在清理系统分区根目录下old文件，请稍等......
  del /f /s /q %systemdrive%\\*.old
  echo.
  echo 清理系统分区根目录下old文件完成！
  echo.
  echo 正在清理系统分区根目录下回收站的文件，请稍等......
  del /f /s /q %systemdrive%\\recycled\\*.*
  echo.
  echo 清理系统分区根目录下回收站的文件完成！
  echo.
  echo 正在清理系统windows目录下的所有备份文件，请稍等......
  del /f /s /q %windir%\\*.bak
  echo.
  echo 清理系统windows目录下的所有备份文件完成！
  echo.
  echo 正在清理系统windows\\prefetch\\目录下预读文件，请稍等......
  del /f /s /q %windir%\\prefetch\\*.*
  echo.
  echo 清理系统windows\\prefetch\\目录下预读文件完成！
  echo.
  echo 正在清理系统windows临时目录下的文件，请稍等......
  rd /s /q %windir%\\temp & md %windir%\\temp
  echo.
  echo 清理系统windows临时目录下的文件完成！
  echo.
  echo 正在清理当前用户目录下的cookies文件，请稍等......
  del /f /q %userprofile%\\cookies\\*.*
  echo.
  echo 清理当前用户目录下的cookies文件完成！
  echo.
  echo 正在清理用户目录下用户最近使用的信息文件，请稍等......
  del /f /q %userprofile%\\recent\\*.*
  echo.
  echo 清理用户目录下用户最近使用的信息文件完成！
  echo.
  echo 正在清理用户目录下的Internet临时文件，请稍等......
  del /f /s /q "%userprofile%\\Local Settings\\Temporary Internet Files\\*.*"
  echo.
  echo 清理用户目录下的Internet临时文件完成！
  echo.
  echo 正在清理用户目录下的临时文件，请稍等......
  del /f /s /q "%userprofile%\\Local Settings\\Temp\\*.*"
  echo.
  echo 清理用户目录下的临时文件完成！
  echo.
  echo 正在清理用户目录下最近打开文件记录，请稍等......
  del /f /s /q "%userprofile%\\recent\\*.*"
  echo.
  echo 清理用户目录下最近打开文件记录完成！
  ECHO.


wevtutil.exe cl "Application"
wevtutil.exe cl "HardwareEvents"
wevtutil.exe cl "Internet Explorer"
wevtutil.exe cl "Key Management Service"
wevtutil.exe cl "Microsoft-AppV-Client/Admin"
wevtutil.exe cl "Microsoft-AppV-Client/Operational"
wevtutil.exe cl "Microsoft-AppV-Client/Virtual Applications"
wevtutil.exe cl "Microsoft-User Experience Virtualization-Agent Driver/Operational"
wevtutil.exe cl "Microsoft-User Experience Virtualization-App Agent/Operational"
wevtutil.exe cl "Microsoft-User Experience Virtualization-IPC/Operational"
wevtutil.exe cl "Microsoft-User Experience Virtualization-SQM Uploader/Operational"
wevtutil.exe cl "Microsoft-Windows-AAD/Operational"
wevtutil.exe cl "Microsoft-Windows-All-User-Install-Agent/Admin"
wevtutil.exe cl "Microsoft-Windows-AllJoyn/Operational"
wevtutil.exe cl "Microsoft-Windows-AppHost/Admin"
wevtutil.exe cl "Microsoft-Windows-AppID/Operational"
wevtutil.exe cl "Microsoft-Windows-ApplicabilityEngine/Operational"
wevtutil.exe cl "Microsoft-Windows-Application Server-Applications/Operational"
wevtutil.exe cl "Microsoft-Windows-Application Server-Applications/Admin"
wevtutil.exe cl "Microsoft-Windows-Application-Experience/Program-Compatibility-Assistant"
wevtutil.exe cl "Microsoft-Windows-Application-Experience/Program-Compatibility-Troubleshooter"
wevtutil.exe cl "Microsoft-Windows-Application-Experience/Program-Inventory"
wevtutil.exe cl "Microsoft-Windows-Application-Experience/Program-Telemetry"
wevtutil.exe cl "Microsoft-Windows-Application-Experience/Steps-Recorder"
wevtutil.exe cl "Microsoft-Windows-ApplicationResourceManagementSystem/Operational"
wevtutil.exe cl "Microsoft-Windows-AppLocker/EXE and DLL"
wevtutil.exe cl "Microsoft-Windows-AppLocker/MSI and Script"
wevtutil.exe cl "Microsoft-Windows-AppLocker/Packaged app-Deployment"
wevtutil.exe cl "Microsoft-Windows-AppLocker/Packaged app-Execution"
wevtutil.exe cl "Microsoft-Windows-AppModel-Runtime/Admin"
wevtutil.exe cl "Microsoft-Windows-AppReadiness/Admin"
wevtutil.exe cl "Microsoft-Windows-AppReadiness/Operational"
wevtutil.exe cl "Microsoft-Windows-TWinUI/Operational"
wevtutil.exe cl "Microsoft-Windows-CoreApplication/Operational"
wevtutil.exe cl "Microsoft-Windows-AppXDeployment/Operational"
wevtutil.exe cl "Microsoft-Windows-AppXDeploymentServer/Operational"
wevtutil.exe cl "Microsoft-Windows-Diagnosis-Scripted/Operational"
wevtutil.exe cl "Microsoft-Windows-AppXDeploymentServer/Restricted"
wevtutil.exe cl "Microsoft-Windows-AppxPackaging/Operational"
wevtutil.exe cl "Microsoft-Windows-ASN1/Operational"
wevtutil.exe cl "Microsoft-Windows-AssignedAccess/Admin"
wevtutil.exe cl "Microsoft-Windows-AssignedAccess/Operational"
wevtutil.exe cl "Microsoft-Windows-AssignedAccessBroker/Admin"
wevtutil.exe cl "Microsoft-Windows-Storage-ATAPort/Admin"
wevtutil.exe cl "Microsoft-Windows-Storage-ATAPort/Operational"
wevtutil.exe cl "Microsoft-Windows-Audio/CaptureMonitor"
wevtutil.exe cl "Microsoft-Windows-Audio/Operational"
wevtutil.exe cl "Microsoft-Windows-FMS/Operational"
wevtutil.exe cl "Microsoft-Windows-Audio/PlaybackManager"
wevtutil.exe cl "Microsoft-Windows-Authentication User Interface/Operational"
wevtutil.exe cl "Microsoft-Windows-BackgroundTaskInfrastructure/Operational"
wevtutil.exe cl "Microsoft-Windows-Backup"
wevtutil.exe cl "Microsoft-Windows-Biometrics/Operational"
wevtutil.exe cl "Microsoft-Windows-BitLocker/BitLocker Management"
wevtutil.exe cl "Microsoft-Windows-BitLocker-DrivePreparationTool/Admin"
wevtutil.exe cl "Microsoft-Windows-BitLocker-DrivePreparationTool/Operational"
wevtutil.exe cl "Microsoft-Windows-Bits-Client/Operational"
wevtutil.exe cl "Microsoft-Windows-Bluetooth-BthLEPrepairing/Operational"
wevtutil.exe cl "Microsoft-Windows-Bluetooth-MTPEnum/Operational"
wevtutil.exe cl "Microsoft-Windows-BranchCache/Operational"
wevtutil.exe cl "Microsoft-Windows-BranchCacheSMB/Operational"
wevtutil.exe cl "Microsoft-Windows-Regsvr32/Operational"
wevtutil.exe cl "Microsoft-Windows-CertificateServicesClient-Lifecycle-System/Operational"
wevtutil.exe cl "Microsoft-Windows-CertificateServicesClient-Lifecycle-User/Operational"
wevtutil.exe cl "Microsoft-Client-Licensing-Platform/Admin"
wevtutil.exe cl "Microsoft-Windows-CloudStorageWizard/Operational"
wevtutil.exe cl "Microsoft-Windows-CodeIntegrity/Operational"
wevtutil.exe cl "Microsoft-Windows-Compat-Appraiser/Operational"
wevtutil.exe cl "Microsoft-Windows-Containers-Wcifs/Operational"
wevtutil.exe cl "Microsoft-Windows-Containers-Wcnfs/Operational"
wevtutil.exe cl "Microsoft-Windows-CoreSystem-SmsRouter-Events/Operational"
wevtutil.exe cl "Microsoft-Windows-CorruptedFileRecovery-Client/Operational"
wevtutil.exe cl "Microsoft-Windows-CorruptedFileRecovery-Server/Operational"
wevtutil.exe cl "Microsoft-Windows-Crypto-DPAPI/BackUpKeySvc"
wevtutil.exe cl "Microsoft-Windows-Crypto-DPAPI/Operational"
wevtutil.exe cl "Microsoft-Windows-DAL-Provider/Operational"
wevtutil.exe cl "Microsoft-Windows-DataIntegrityScan/Admin"
wevtutil.exe cl "Microsoft-Windows-DataIntegrityScan/CrashRecovery"
wevtutil.exe cl "Microsoft-Windows-DateTimeControlPanel/Operational"
wevtutil.exe cl "Microsoft-Windows-Deduplication/Diagnostic"
wevtutil.exe cl "Microsoft-Windows-Deduplication/Operational"
wevtutil.exe cl "Microsoft-Windows-Deduplication/Scrubbing"
wevtutil.exe cl "Microsoft-Windows-DSC/Admin"
wevtutil.exe cl "Microsoft-Windows-DSC/Operational"
wevtutil.exe cl "Microsoft-Windows-DeviceGuard/Operational"
wevtutil.exe cl "Microsoft-Windows-DeviceManagement-Enterprise-Diagnostics-Provider/Admin"
wevtutil.exe cl "Microsoft-Windows-Devices-Background/Operational"
wevtutil.exe cl "Microsoft-Windows-DeviceSetupManager/Admin"
wevtutil.exe cl "Microsoft-Windows-DeviceSetupManager/Operational"
wevtutil.exe cl "Microsoft-Windows-DeviceSync/Operational"
wevtutil.exe cl "Microsoft-Windows-Dhcp-Client/Admin"
wevtutil.exe cl "Microsoft-Windows-Dhcpv6-Client/Admin"
wevtutil.exe cl "Microsoft-Windows-Diagnosis-DPS/Operational"
wevtutil.exe cl "Microsoft-Windows-Diagnosis-PCW/Operational"
wevtutil.exe cl "Microsoft-Windows-Diagnosis-PLA/Operational"
wevtutil.exe cl "Microsoft-Windows-Diagnosis-Scheduled/Operational"
wevtutil.exe cl "Microsoft-Windows-Diagnosis-Scripted/Admin"
wevtutil.exe cl "Microsoft-Windows-Diagnosis-Scripted/Operational"
wevtutil.exe cl "Microsoft-Windows-Diagnosis-ScriptedDiagnosticsProvider/Operational"
wevtutil.exe cl "Microsoft-Windows-Diagnostics-Networking/Operational"
wevtutil.exe cl "Microsoft-Windows-Diagnostics-Performance/Operational"
wevtutil.exe cl "Microsoft-Windows-DiskDiagnosticDataCollector/Operational"
wevtutil.exe cl "Microsoft-Windows-DiskDiagnosticResolver/Operational"
wevtutil.exe cl "Microsoft-Windows-EapHost/Operational"
wevtutil.exe cl "Microsoft-Windows-EapMethods-RasChap/Operational"
wevtutil.exe cl "Microsoft-Windows-EapMethods-RasTls/Operational"
wevtutil.exe cl "Microsoft-Windows-EapMethods-Sim/Operational"
wevtutil.exe cl "Microsoft-Windows-EapMethods-Ttls/Operational"
wevtutil.exe cl "Microsoft-Windows-EDP-Audit-Regular/Admin"
wevtutil.exe cl "Microsoft-Windows-EDP-Audit-TCB/Admin"
wevtutil.exe cl "Microsoft-Windows-EmbeddedAppLauncher/Admin"
wevtutil.exe cl "Microsoft-Windows-EventCollector/Operational"
wevtutil.exe cl "Microsoft-Windows-Forwarding/Operational"
wevtutil.exe cl "Microsoft-Windows-Fault-Tolerant-Heap/Operational"
wevtutil.exe cl "Microsoft-Windows-FileHistory-Core/WHC"
wevtutil.exe cl "Microsoft-Windows-FileHistory-Engine/BackupLog"
wevtutil.exe cl "Microsoft-Windows-FMS/Operational"
wevtutil.exe cl "Microsoft-Windows-Folder Redirection/Operational"
wevtutil.exe cl "Microsoft-Windows-GenericRoaming/Admin"
wevtutil.exe cl "Microsoft-Windows-GroupPolicy/Operational"
wevtutil.exe cl "Microsoft-Windows-Help/Operational"
wevtutil.exe cl "Microsoft-Windows-HomeGroup Control Panel/Operational"
wevtutil.exe cl "Microsoft-Windows-HomeGroup Provider Service/Operational"
wevtutil.exe cl "Microsoft-Windows-HomeGroup Listener Service/Operational"
wevtutil.exe cl "Microsoft-Windows-HotspotAuth/Operational"
wevtutil.exe cl "Microsoft-Windows-Hyper-V-Guest-Drivers/Admin"
wevtutil.exe cl "Microsoft-Windows-IdCtrls/Operational"
wevtutil.exe cl "Microsoft-Windows-International/Operational"
wevtutil.exe cl "Microsoft-Windows-International-RegionalOptionsControlPanel/Operational"
wevtutil.exe cl "Microsoft-Windows-Iphlpsvc/Operational"
wevtutil.exe cl "Microsoft-Windows-KdsSvc/Operational"
wevtutil.exe cl "Microsoft-Windows-Kernel-ApphelpCache/Operational"
wevtutil.exe cl "Microsoft-Windows-Kernel-Boot/Operational"
wevtutil.exe cl "Microsoft-Windows-Kernel-EventTracing/Admin"
wevtutil.exe cl "Microsoft-Windows-Kernel-IO/Operational"
wevtutil.exe cl "Microsoft-Windows-Kernel-PnP/Configuration"
wevtutil.exe cl "Microsoft-Windows-Kernel-Power/Thermal-Operational"
wevtutil.exe cl "Microsoft-Windows-Kernel-ShimEngine/Operational"
wevtutil.exe cl "Microsoft-Windows-Kernel-StoreMgr/Operational"
wevtutil.exe cl "Microsoft-Windows-Kernel-WDI/Operational"
wevtutil.exe cl "Microsoft-Windows-Kernel-WHEA/Errors"
wevtutil.exe cl "Microsoft-Windows-Kernel-WHEA/Operational"
wevtutil.exe cl "Microsoft-Windows-Known Folders API Service"
wevtutil.exe cl "Microsoft-Windows-LanguagePackSetup/Operational"
wevtutil.exe cl "Microsoft-Windows-LiveId/Operational"
wevtutil.exe cl "Microsoft-Windows-MemoryDiagnostics-Results/Debug"
wevtutil.exe cl "Microsoft-Windows-Mobile-Broadband-Experience-Parser-Task/Operational"
wevtutil.exe cl "SMSApi"
wevtutil.exe cl "Microsoft-Windows-Mobile-Broadband-Experience-SmsRouter/Admin"
wevtutil.exe cl "Microsoft-Windows-Mprddm/Operational"
wevtutil.exe cl "Microsoft-Windows-MUI/Admin"
wevtutil.exe cl "Microsoft-Windows-MUI/Operational"
wevtutil.exe cl "Microsoft-Windows-NcdAutoSetup/Operational"
wevtutil.exe cl "Microsoft-Windows-NCSI/Operational"
wevtutil.exe cl "Microsoft-Windows-NdisImPlatform/Operational"
wevtutil.exe cl "Microsoft-Windows-NetworkProfile/Operational"
wevtutil.exe cl "Microsoft-Windows-NetworkProvider/Operational"
wevtutil.exe cl "Microsoft-Windows-NetworkProvisioning/Operational"
wevtutil.exe cl "Microsoft-Windows-NlaSvc/Operational"
wevtutil.exe cl "Microsoft-Windows-Ntfs/Operational"
wevtutil.exe cl "Microsoft-Windows-Ntfs/WHC"
wevtutil.exe cl "Microsoft-Windows-NTLM/Operational"
wevtutil.exe cl "Microsoft-Windows-OfflineFiles/Operational"
wevtutil.exe cl "Microsoft-Windows-OneBackup/Debug"
wevtutil.exe cl "Microsoft-Windows-OOBE-Machine-DUI/Operational"
wevtutil.exe cl "Microsoft-Windows-PackageStateRoaming/Operational"
wevtutil.exe cl "Microsoft-Windows-ParentalControls/Operational"
wevtutil.exe cl "Microsoft-Windows-Partition/Diagnostic"
wevtutil.exe cl "Microsoft-Windows-PerceptionRuntime/Operational"
wevtutil.exe cl "Microsoft-Windows-PerceptionSensorDataService/Operational"
wevtutil.exe cl "Microsoft-Windows-Policy/Operational"
wevtutil.exe cl "Microsoft-Windows-PowerShell/Admin"
wevtutil.exe cl "Microsoft-Windows-PowerShell/Operational"
wevtutil.exe cl "Microsoft-Windows-PowerShell-DesiredStateConfiguration-FileDownloadManager/Operational"
wevtutil.exe cl "Microsoft-Windows-NetworkLocationWizard/Operational"
wevtutil.exe cl "Microsoft-Windows-PrintBRM/Admin"
wevtutil.exe cl "Microsoft-Windows-PrintService/Admin"
wevtutil.exe cl "Microsoft-Windows-Program-Compatibility-Assistant/CompatAfterUpgrade"
wevtutil.exe cl "Microsoft-Windows-Provisioning-Diagnostics-Provider/Admin"
wevtutil.exe cl "Microsoft-Windows-PushNotification-Platform/Admin"
wevtutil.exe cl "Microsoft-Windows-PushNotification-Platform/Operational"
wevtutil.exe cl "Microsoft-Windows-ReadyBoost/Operational"
wevtutil.exe cl "Microsoft-Windows-ReadyBoostDriver/Operational"
wevtutil.exe cl "Microsoft-Windows-RemoteApp and Desktop Connections/Admin"
wevtutil.exe cl "Microsoft-Windows-RemoteApp and Desktop Connections/Operational"
wevtutil.exe cl "Microsoft-Windows-RemoteAssistance/Admin"
wevtutil.exe cl "Microsoft-Windows-RemoteAssistance/Operational"
wevtutil.exe cl "Microsoft-Windows-RemoteDesktopServices-RdpCoreTS/Admin"
wevtutil.exe cl "Microsoft-Windows-RemoteDesktopServices-RdpCoreTS/Operational"
wevtutil.exe cl "Microsoft-Windows-RemoteDesktopServices-RemoteFX-Synth3dvsc/Admin"
wevtutil.exe cl "Microsoft-Windows-RemoteDesktopServices-SessionServices/Operational"
wevtutil.exe cl "Microsoft-Windows-Resource-Exhaustion-Detector/Operational"
wevtutil.exe cl "Microsoft-Windows-Resource-Exhaustion-Resolver/Operational"
wevtutil.exe cl "Microsoft-Windows-RestartManager/Operational"
wevtutil.exe cl "Microsoft-Windows-RetailDemo/Admin"
wevtutil.exe cl "Microsoft-Windows-RetailDemo/Operational"
wevtutil.exe cl "Microsoft-Windows-ScmBus/Certification"
wevtutil.exe cl "Microsoft-Windows-ScmDisk0101/Operational"
wevtutil.exe cl "Microsoft-Windows-Security-Audit-Configuration-Client/Operational"
wevtutil.exe cl "Microsoft-Windows-Security-EnterpriseData-FileRevocationManager/Operational"
wevtutil.exe cl "Microsoft-Windows-Security-Netlogon/Operational"
wevtutil.exe cl "Microsoft-Windows-Security-SPP-UX-GenuineCenter-Logging/Operational"
wevtutil.exe cl "Microsoft-Windows-Security-SPP-UX-Notifications/ActionCenter"
wevtutil.exe cl "Microsoft-Windows-Security-UserConsentVerifier/Audit"
wevtutil.exe cl "Microsoft-Windows-SENSE/Operational"
wevtutil.exe cl "Microsoft-Windows-SettingSync/Debug"
wevtutil.exe cl "Microsoft-Windows-SettingSync/Operational"
wevtutil.exe cl "Microsoft-Windows-SettingSync-Azure/Debug"
wevtutil.exe cl "Microsoft-Windows-SettingSync-Azure/Operational"
wevtutil.exe cl "Microsoft-Windows-Shell-ConnectedAccountState/ActionCenter"
wevtutil.exe cl "Microsoft-Windows-Shell-Core/ActionCenter"
wevtutil.exe cl "Microsoft-Windows-Shell-Core/AppDefaults"
wevtutil.exe cl "Microsoft-Windows-Shell-Core/LogonTasksChannel"
wevtutil.exe cl "Microsoft-Windows-Shell-Core/Operational"
wevtutil.exe cl "Microsoft-Windows-SmartCard-Audit/Authentication"
wevtutil.exe cl "Microsoft-Windows-SmartCard-DeviceEnum/Operational"
wevtutil.exe cl "Microsoft-Windows-SmartCard-TPM-VCard-Module/Admin"
wevtutil.exe cl "Microsoft-Windows-SmartCard-TPM-VCard-Module/Operational"
wevtutil.exe cl "Microsoft-Windows-SmbClient/Connectivity"
wevtutil.exe cl "Microsoft-Windows-SMBClient/Operational"
wevtutil.exe cl "Microsoft-Windows-SmbClient/Security"
wevtutil.exe cl "Microsoft-Windows-SMBServer/Audit"
wevtutil.exe cl "Microsoft-Windows-SMBServer/Connectivity"
wevtutil.exe cl "Microsoft-Windows-SMBServer/Operational"
wevtutil.exe cl "Microsoft-Windows-SMBServer/Security"
wevtutil.exe cl "Microsoft-Windows-SMBWitnessClient/Admin"
wevtutil.exe cl "Microsoft-Windows-SMBWitnessClient/Informational"
wevtutil.exe cl "Microsoft-Windows-StateRepository/Operational"
wevtutil.exe cl "Microsoft-Windows-StateRepository/Restricted"
wevtutil.exe cl "Microsoft-Windows-Storage-Tiering/Admin"
wevtutil.exe cl "Microsoft-Windows-StorageManagement/Operational"
wevtutil.exe cl "Microsoft-Windows-StorageSpaces-Driver/Diagnostic"
wevtutil.exe cl "Microsoft-Windows-StorageSpaces-Driver/Operational"
wevtutil.exe cl "Microsoft-Windows-StorageSpaces-ManagementAgent/WHC"
wevtutil.exe cl "Microsoft-Windows-StorageSpaces-SpaceManager/Diagnostic"
wevtutil.exe cl "Microsoft-Windows-StorageSpaces-SpaceManager/Operational"
wevtutil.exe cl "Microsoft-Windows-Storage-ClassPnP/Operational"
wevtutil.exe cl "Microsoft-Windows-Store/Operational"
wevtutil.exe cl "Microsoft-Windows-Storage-Storport/Operational"
wevtutil.exe cl "Microsoft-Windows-SystemSettingsThreshold/Operational"
wevtutil.exe cl "Microsoft-Windows-TaskScheduler/Maintenance"
wevtutil.exe cl "Microsoft-Windows-TCPIP/Operational"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-RDPClient/Operational"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-ClientUSBDevices/Admin"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-ClientUSBDevices/Operational"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-LocalSessionManager/Admin"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-LocalSessionManager/Operational"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-PnPDevices/Admin"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-PnPDevices/Operational"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-Printers/Admin"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-Printers/Operational"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-RemoteConnectionManager/Admin"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-RemoteConnectionManager/Operational"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-ServerUSBDevices/Admin"
wevtutil.exe cl "Microsoft-Windows-TerminalServices-ServerUSBDevices/Operational"
wevtutil.exe cl "Microsoft-Windows-TZSync/Operational"
wevtutil.exe cl "Microsoft-Windows-TZUtil/Operational"
wevtutil.exe cl "Microsoft-Windows-UAC/Operational"
wevtutil.exe cl "Microsoft-Windows-UAC-FileVirtualization/Operational"
wevtutil.exe cl "Microsoft-Windows-SearchUI/Operational"
wevtutil.exe cl "Microsoft-Windows-UniversalTelemetryClient/Operational"
wevtutil.exe cl "Microsoft-Windows-User Control Panel/Operational"
wevtutil.exe cl "Microsoft-Windows-User Device Registration/Admin"
wevtutil.exe cl "Microsoft-Windows-User Profile Service/Operational"
wevtutil.exe cl "Microsoft-Windows-User-Loader/Operational"
wevtutil.exe cl "Microsoft-Windows-UserPnp/ActionCenter"
wevtutil.exe cl "Microsoft-Windows-UserPnp/DeviceInstall"
wevtutil.exe cl "Microsoft-Windows-VDRVROOT/Operational"
wevtutil.exe cl "Microsoft-Windows-VerifyHardwareSecurity/Admin"
wevtutil.exe cl "Microsoft-Windows-VHDMP-Operational"
wevtutil.exe cl "Microsoft-Windows-Volume/Diagnostic"
wevtutil.exe cl "Microsoft-Windows-VolumeSnapshot-Driver/Operational"
wevtutil.exe cl "Microsoft-Windows-VPN-Client/Operational"
wevtutil.exe cl "Microsoft-Windows-Wcmsvc/Operational"
wevtutil.exe cl "Microsoft-Windows-IKE/Operational"
wevtutil.exe cl "Microsoft-Windows-VPN/Operational"
wevtutil.exe cl "Microsoft-Windows-WFP/Operational"
wevtutil.exe cl "Microsoft-WindowsPhone-Connectivity-WiFiConnSvc-Channel"
wevtutil.exe cl "Microsoft-Windows-Win32k/Operational"
wevtutil.exe cl "Microsoft-Windows-Windows Defender/Operational"
wevtutil.exe cl "Microsoft-Windows-Windows Defender/WHC"
wevtutil.exe cl "Microsoft-Windows-Windows Firewall With Advanced Security/ConnectionSecurity"
wevtutil.exe cl "Microsoft-Windows-Windows Firewall With Advanced Security/Firewall"
wevtutil.exe cl "Microsoft-Windows-WinRM/Operational"
wevtutil.exe cl "Microsoft-Windows-WindowsBackup/ActionCenter"
wevtutil.exe cl "Microsoft-Windows-WindowsSystemAssessmentTool/Operational"
wevtutil.exe cl "Microsoft-Windows-WindowsUpdateClient/Operational"
wevtutil.exe cl "Microsoft-Windows-WinINet-Config/ProxyConfigChanged"
wevtutil.exe cl "Microsoft-Windows-Winlogon/Operational"
wevtutil.exe cl "Microsoft-Windows-Winsock-WS2HELP/Operational"
wevtutil.exe cl "Microsoft-Windows-Wired-AutoConfig/Operational"
wevtutil.exe cl "Microsoft-Windows-WLAN-AutoConfig/Operational"
wevtutil.exe cl "Microsoft-Windows-WMI-Activity/Operational"
wevtutil.exe cl "Microsoft-Windows-WorkFolders/WHC"
wevtutil.exe cl "Microsoft-Windows-WorkFolders/Operational"
wevtutil.exe cl "Microsoft-Windows-Workplace Join/Admin"
wevtutil.exe cl "Microsoft-Windows-WPD-CompositeClassDriver/Operational"
wevtutil.exe cl "Microsoft-Windows-WPD-MTPClassDriver/Operational"
wevtutil.exe cl "Microsoft-Windows-WWAN-SVC-Events/Operational
wevtutil.exe cl "OAlerts"
wevtutil.exe cl "Kaspersky Event Log"
wevtutil.exe cl "Security"
wevtutil.exe cl "Setup"
wevtutil.exe cl "System"
wevtutil.exe cl "Windows PowerShell"

echo .
echo Clean Finally !
echo .