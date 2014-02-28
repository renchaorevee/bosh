package platform

import (
	boshdevicepathresolver "bosh/infrastructure/device_path_resolver"
	boshcmd "bosh/platform/commands"
	boshdisk "bosh/platform/disk"
	boshstats "bosh/platform/stats"
	boshvitals "bosh/platform/vitals"
	boshsettings "bosh/settings"
	boshdir "bosh/settings/directories"
	boshdirs "bosh/settings/directories"
	boshsys "bosh/system"
	"time"
)

type dummyPlatform struct {
	collector          boshstats.StatsCollector
	fs                 boshsys.FileSystem
	cmdRunner          boshsys.CmdRunner
	compressor         boshcmd.Compressor
	copier             boshcmd.Copier
	dirProvider        boshdirs.DirectoriesProvider
	vitalsService      boshvitals.Service
	diskManager        boshdisk.Manager
	devicePathResolver boshdevicepathresolver.DevicePathResolver
}

func NewDummyPlatform(
	collector boshstats.StatsCollector,
	fs boshsys.FileSystem,
	cmdRunner boshsys.CmdRunner,
	dirProvider boshdirs.DirectoriesProvider,
	diskManager boshdisk.Manager,
) (platform dummyPlatform) {
	platform.collector = collector
	platform.fs = fs
	platform.cmdRunner = cmdRunner
	platform.dirProvider = dirProvider
	platform.devicePathResolver = boshdevicepathresolver.NewDevicePathResolver(1*time.Millisecond, platform.fs)

	platform.compressor = boshcmd.NewTarballCompressor(cmdRunner, fs)
	platform.copier = boshcmd.NewCpCopier(cmdRunner, fs)
	platform.vitalsService = boshvitals.NewService(collector, dirProvider)
	platform.diskManager = diskManager
	return
}

func (p dummyPlatform) GetFs() (fs boshsys.FileSystem) {
	return p.fs
}

func (p dummyPlatform) GetRunner() (runner boshsys.CmdRunner) {
	return p.cmdRunner
}

func (p dummyPlatform) GetStatsCollector() (collector boshstats.StatsCollector) {
	return p.collector
}

func (p dummyPlatform) GetCompressor() (compressor boshcmd.Compressor) {
	return p.compressor
}

func (p dummyPlatform) GetCopier() (copier boshcmd.Copier) {
	return p.copier
}

func (p dummyPlatform) GetDirProvider() (dirProvider boshdir.DirectoriesProvider) {
	return p.dirProvider
}

func (p dummyPlatform) GetVitalsService() (service boshvitals.Service) {
	return p.vitalsService
}

func (p dummyPlatform) GetDevicePathResolver() (devicePathResolver boshdevicepathresolver.DevicePathResolver) {
	return p.devicePathResolver
}

func (p dummyPlatform) SetupRuntimeConfiguration() (err error) {
	return
}

func (p dummyPlatform) CreateUser(username, password, basePath string) (err error) {
	return
}

func (p dummyPlatform) AddUserToGroups(username string, groups []string) (err error) {
	return
}

func (p dummyPlatform) DeleteEphemeralUsersMatching(regex string) (err error) {
	return
}

func (p dummyPlatform) SetupSsh(publicKey, username string) (err error) {
	return
}

func (p dummyPlatform) SetUserPassword(user, encryptedPwd string) (err error) {
	return
}

func (p dummyPlatform) SetupHostname(hostname string) (err error) {
	return
}

func (p dummyPlatform) SetupDhcp(networks boshsettings.Networks) (err error) {
	return
}

func (p dummyPlatform) SetupManualNetworking(networks boshsettings.Networks) (err error) {
	return
}

func (p dummyPlatform) SetupLogrotate(groupName, basePath, size string) (err error) {
	return
}

func (p dummyPlatform) SetTimeWithNtpServers(servers []string) (err error) {
	return
}

func (p dummyPlatform) SetupEphemeralDiskWithPath(devicePath string) (err error) {
	return
}

func (p dummyPlatform) SetupTmpDir() (err error) {
	return
}

func (p dummyPlatform) MountPersistentDisk(devicePath, mountPoint string) (err error) {
	return
}

func (p dummyPlatform) UnmountPersistentDisk(devicePath string) (didUnmount bool, err error) {
	return
}

func (p dummyPlatform) NormalizeDiskPath(attachment string) (devicePath string, found bool) {
	return "/dev/sdb", true
}

func (p dummyPlatform) LookupScsiDisk(scsiId string) (devicePath string, found bool) {
	return "/dev/sdb", true
}

func (p dummyPlatform) GetFileContentsFromCDROM(filePath string) (contents []byte, err error) {
	return
}

func (p dummyPlatform) MigratePersistentDisk(fromMountPoint, toMountPoint string) (err error) {
	return
}

func (p dummyPlatform) IsMountPoint(path string) (result bool, err error) {
	return
}

func (p dummyPlatform) IsDevicePathMounted(path string) (result bool, err error) {
	return
}

func (p dummyPlatform) StartMonit() (err error) {
	return
}

func (p dummyPlatform) SetupMonitUser() (err error) {
	return
}

func (p dummyPlatform) GetMonitCredentials() (username, password string, err error) {
	return
}

func (p dummyPlatform) GetDiskManager() (diskManager boshdisk.Manager) {
	return
}
