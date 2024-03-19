/*
 * SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
 *
 * SPDX-License-Identifier: LGPL-3.0-or-later
 */

package linglong

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"

	"pkg.deepin.com/linglong/pica/cmd/ll-pica/core/comm"
	"pkg.deepin.com/linglong/pica/cmd/ll-pica/utils/fs"
	"pkg.deepin.com/linglong/pica/cmd/ll-pica/utils/log"
)

type LinglongBuder struct {
	Appid       string
	Name        string
	Version     string
	Runtime     string
	Rversion    string
	Description string
	// Kind        string
	// Hash        string
	// Url         string
	Sources   []comm.Source
	Configure []string
	Install   []string
}

type RuntimeJson struct {
	Appid       string   `json:"appid"`
	Arch        []string `json:"arch"`
	Base        string   `json:"base"`
	Description string   `json:"description"`
	Kind        string   `json:"kind"`
	Name        string   `json:"name"`
	Runtime     string   `json:"runtime"`
	Version     string   `json:"version"`
}

// LoadRuntimeInfo
func (ts *LinglongBuder) LoadRuntimeInfo(path string) bool {
	// load runtime info from file
	if ret, err := fs.CheckFileExits(path); err != nil && !ret {
		log.Logger.Warnf("load runtime info failed: %v", err)
		return false
	}
	var runtimedir RuntimeJson
	runtimedirFd, err := ioutil.ReadFile(path)
	if err != nil {
		log.Logger.Errorf("get %s error: %v", path, err)
		return false
	}
	err = json.Unmarshal(runtimedirFd, &runtimedir)
	if err != nil {
		log.Logger.Errorf("error: %v", err)
		return false
	}
	// copy to LinglongBuder
	if runtimedir.Appid != "" && runtimedir.Version != "" {
		ts.Runtime = runtimedir.Appid
		ts.Rversion = runtimedir.Version
		return true
	}

	return false
}

const LinglongBuilderTMPL = `
package:
  id: {{.Appid}}
  name: {{.Name}}
  version: {{.Version}}
  kind: app
  description: |
    {{.Description}}

runtime:
  id: {{.Runtime}}
  version: {{.Rversion}}

sources:
{{- range .Sources}}
{{- if eq .Kind "local"}}
  - kind: local
{{- else}}
  - kind: {{.Kind}}
    url: {{.Url}}
    digest: {{.Digest}}
{{end}}
{{- end}}
build:
  kind: manual
  manual:
    configure: |
      # extract deb
      ls *.deb | xargs -I {} dpkg -x {} workdir

      #>>> auto generate by ll-pica begin
      {{- range $line := .Configure}}
        {{- printf "\n      %s" $line}}
      {{- end}}
      #>>> auto generate by ll-pica end

    install: |
      install -d $PREFIX/share
      install -d $PREFIX/bin

      #>>> auto generate by ll-pica begin
      {{- range $line := .Install}}
        {{- printf "\n      %s" $line}}
      {{- end}}
      #>>> auto generate by ll-pica end
`

// CreateLinglongYamlBuilder
func (ts *LinglongBuder) CreateLinglongYamlBuilder(path string) bool {

	tpl, err := template.New("linglong").Parse(LinglongBuilderTMPL)

	if err != nil {
		log.Logger.Fatalf("parse deb shell template failed! ", err)
		return false
	}

	// create save file
	log.Logger.Debug("create save file: ", path)
	saveFd, ret := os.Create(path)
	if ret != nil {
		log.Logger.Fatalf("save to %s failed!", path)
		return false
	}
	defer saveFd.Close()

	// render template
	log.Logger.Debug("render template: ", ts)
	tpl.Execute(saveFd, ts)

	return true

}

// CreateLinglongBuilder
func (ts *LinglongBuder) CreateLinglongBuilder(path string) bool {

	log.Logger.Debugf("create save file: ", path)

	// check workstation
	if ret, err := fs.CheckFileExits(path); err != nil && !ret {
		log.Logger.Errorf("workstation witch convert not found: %s", path)
		return false
	} else {
		err := os.Chdir(path)
		if err != nil {
			log.Logger.Errorf("workstation can not enter directory: %s", path)
			return false
		}
	}

	// caller ll-builder build
	if ret, msg, err := comm.ExecAndWait(10, "ll-builder", "build"); err != nil {
		log.Logger.Fatalf("ll-builder failed: ", err, msg, ret)
		return false
	} else {
		log.Logger.Infof("ll-builder succeeded: ", path, ret)
		return true
	}
}

func (ts *LinglongBuder) LinglongExport(path string) bool {
	log.Logger.Debugf("ll-builder import : ", ts.Appid)
	appExportPath := fs.GetFilePPath(path)
	appExportPath = fs.GetFilePPath(appExportPath)
	// check workstation
	if ret, err := fs.CheckFileExits(path); err != nil && !ret {
		log.Logger.Errorf("workstation witch convert not found: %s", path)
		return false
	} else {
		err := os.Chdir(appExportPath)
		if err != nil {
			log.Logger.Errorf("workstation can not enter directory: %s", appExportPath)
			return false
		}
	}
	// caller ll-builder export --local
	if ret, msg, err := comm.ExecAndWait(120, "ll-builder", "export", path); err != nil {
		log.Logger.Fatalf("ll-builder export failed: ", err, msg, ret)
		return false
	} else {
		log.Logger.Infof("ll-builder export succeeded: ", path, ret)
	}

	// chmod 755 uab
	if bundleList, err := fs.FindBundlePath(appExportPath); err != nil {
		log.Logger.Errorf("not found bundle")
		return false
	} else {
		for _, bundle := range bundleList {
			log.Logger.Infof("chmod 0755 for %s", bundle)
			if err := os.Chmod(bundle, 0755); err != nil {
				log.Logger.Errorf("chmod 0755 for %s failed！", bundle)
				return false
			}
		}
	}
	return true
}
