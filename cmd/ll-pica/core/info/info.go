/*
 * SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
 *
 * SPDX-License-Identifier: LGPL-3.0-or-later
 */

package info

import (
	"pkg.deepin.com/linglong/pica/cmd/ll-pica/core/comm"
	"pkg.deepin.com/linglong/pica/cmd/ll-pica/core/linglong"
)

type InfoApp struct {
	Appid          string          `json:"appid"`
	Name           string          `json:"name"`
	Version        string          `json:"version"`
	Kind           string          `json:"kind"`
	Description    string          `json:"description"`
	Runtime        string          `json:"runtime"`
	Module         string          `json:"module"`
	Arch           []string        `json:"arch"`
	Permissions    InfoPermissions `json:"permissions"`
	SupportPlugins []string        `json:"support-plugins"`
	Plugins        []string        `json:"plugins"`
}

type InfoPermissions struct {
	AutoStart     bool `json:"autostart"`
	Notification  bool `json:"notification"`
	Trayicon      bool `json:"trayicon"`
	Clipboard     bool `json:"clipboard"`
	Account       bool `json:"account"`
	Bluetooth     bool `json:"bluetooth"`
	Camera        bool `json:"camera"`
	AudioRecord   bool `json:"audio_record"`
	InstalledApps bool `json:"installed_apps"`
}

func CreateInfo(info comm.Config, debInfo *comm.DebConfig, lb linglong.LinglongBuder) (bool, error) {
	// if ret, err := fs.CheckFileExits(info.ExportDir); !ret && err != nil {
	// 	log.Logger.Errorw("info.json dir not exists! : ", info.ExportDir)
	// 	return false, err
	// }
	// infoFilePath := filepath.Clean(info.ExportDir) + "/info.json"
	// hostArch := runtime.GOARCH
	// if hostArch == "amd64" {
	// 	hostArch = "x86_64"
	// }

	// infoApp := &InfoApp{
	// 	Appid:       debInfo.Info.Appid,
	// 	Name:        debInfo.Info.Name,
	// 	Version:     debInfo.Info.Version,
	// 	Kind:        "app",
	// 	Description: debInfo.Info.Description,
	// 	// "org.deepin.Runtime/20.5.0/x86_64",
	// 	Runtime: fmt.Sprintf("%s/%s/%s", lb.Runtime, lb.Rversion, hostArch),
	// 	Module:  "runtime",
	// }
	// if debInfo.Info.Version == "" || debInfo.Info.Description == "" {
	// 	// Package: deepin-calculator
	// 	// Version: 5.7.16-1
	// 	// Description: Calculator for UOS
	// 	// /var/lib/dpkg/status
	// 	dpkgStatus := info.Basedir + "/var/lib/dpkg/status"
	// 	if ret, err := fs.CheckFileExits(dpkgStatus); err != nil && !ret {
	// 		log.Logger.Warnf("can not found dpkg info %s , %v", dpkgStatus, err)
	// 	}
	// 	if dpkgStatusFile, err := os.Open(dpkgStatus); err != nil {
	// 		log.Logger.Warnf("open status failed:", err)

	// 	} else {
	// 		defer dpkgStatusFile.Close()

	// 		LogFileItor := bufio.NewScanner(dpkgStatusFile)
	// 		LogFileItor.Split(bufio.ScanLines)
	// 		var ReadLine string
	// 		strHeader := fmt.Sprintf("Package: %s", debInfo.Info.Name)
	// 		for LogFileItor.Scan() {
	// 			ReadLine = LogFileItor.Text()

	// 			if ReadLine == strHeader {

	// 				for LogFileItor.Scan() {
	// 					ReadLine = LogFileItor.Text()

	// 					if ReadLine != "" {
	// 						// Version
	// 						if debInfo.Info.Version == "" && strings.HasPrefix(ReadLine, "Version:") {
	// 							ReadVersion := strings.Split(ReadLine, "Version: ")[1]

	// 							if ret := strings.Index(ReadVersion, ":"); ret != -1 {
	// 								ReadVersion = strings.Split(ReadVersion, ":")[1]
	// 							}
	// 							if ret := strings.Index(ReadVersion, "-"); ret != -1 {
	// 								ReadVersion = strings.Split(ReadVersion, "-")[0]
	// 							}
	// 							if ret := strings.Index(ReadVersion, "+"); ret != -1 {
	// 								ReadVersion = strings.Split(ReadVersion, "+")[0]
	// 							}

	// 							verList := []string{}

	// 							regexFP := func() []string {

	// 								regexVerList := strings.Split(ReadVersion, ".")[0:]
	// 								if len(regexVerList) > 3 {
	// 									return regexVerList[:4]
	// 								} else {
	// 									return regexVerList
	// 								}

	// 							}
	// 							regexVer := regexp.MustCompile(`^[-+]?\d+`)
	// 							for _, ver := range regexFP() {
	// 								strVer := regexVer.FindString(ver)
	// 								if strVer == "" {
	// 									verList = append(verList, "0")
	// 								}
	// 								if ret, err := strconv.ParseInt(strVer, 10, 64); err != nil {
	// 									verList = append(verList, "0")
	// 								} else {
	// 									verList = append(verList, fmt.Sprintf("%d", ret))
	// 								}
	// 							}
	// 							infoApp.Version = strings.Join(verList, ".")
	// 							debInfo.Info.Version = strings.Join(verList, ".")
	// 						}
	// 						// Description
	// 						if debInfo.Info.Description == "" && strings.HasPrefix(ReadLine, "Description:") {
	// 							infoApp.Description = strings.Split(ReadLine, "Description: ")[1]
	// 						}

	// 					} else {
	// 						break
	// 					}
	// 				}
	// 			}

	// 		}
	// 	}
	// }
	// infoApp.Permissions.AutoStart = false
	// infoApp.Permissions.Notification = false
	// infoApp.Permissions.Trayicon = false
	// infoApp.Permissions.Clipboard = false
	// infoApp.Permissions.Account = false
	// infoApp.Permissions.Bluetooth = false
	// infoApp.Permissions.Camera = false
	// infoApp.Permissions.AudioRecord = false
	// infoApp.Permissions.InstalledApps = false

	// infoApp.Arch = append(infoApp.Arch, hostArch)
	// infoApp.SupportPlugins = make([]string, 0)
	// infoApp.Plugins = make([]string, 0)

	// data, err := json.MarshalIndent(infoApp, "", "\t")
	// if err != nil {
	// 	log.Logger.Errorw("序列化错误： ", infoFilePath)
	// 	return false, err
	// }

	// // 创建文件
	// file, err := os.Create(infoFilePath)

	// if err != nil {
	// 	log.Logger.Errorw("create file error: ", infoFilePath)
	// 	return false, err
	// }
	// defer file.Close()

	// file.Write(data)

	return true, nil
}
