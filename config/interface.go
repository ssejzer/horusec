// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/spf13/cobra"

	customImages "github.com/ZupIT/horusec/internal/entities/custom_images"
	"github.com/ZupIT/horusec/internal/entities/toolsconfig"
	"github.com/ZupIT/horusec/internal/entities/workdir"
)

type IConfig interface {
	NewConfigsFromCobraAndLoadsCmdGlobalFlags(cmd *cobra.Command) IConfig
	NewConfigsFromCobraAndLoadsCmdStartFlags(cmd *cobra.Command) IConfig
	NewConfigsFromViper() IConfig
	NewConfigsFromEnvironments() IConfig

	GetVersion() string

	GetDefaultConfigFilePath() string
	GetConfigFilePath() string
	SetConfigFilePath(configFilePath string)

	GetLogLevel() string
	SetLogLevel(logLevel string)

	GetHorusecAPIUri() string
	SetHorusecAPIURI(horusecAPIURI string)

	GetTimeoutInSecondsRequest() int64
	SetTimeoutInSecondsRequest(timeoutInSecondsRequest int64)

	GetTimeoutInSecondsAnalysis() int64
	SetTimeoutInSecondsAnalysis(timeoutInSecondsAnalysis int64)

	GetMonitorRetryInSeconds() int64
	SetMonitorRetryInSeconds(retryInterval int64)

	GetRepositoryAuthorization() string
	SetRepositoryAuthorization(repositoryAuthorization string)

	GetPrintOutputType() string
	SetPrintOutputType(printOutputType string)

	GetJSONOutputFilePath() string
	SetJSONOutputFilePath(jsonOutputFilePath string)

	GetSeveritiesToIgnore() []string
	SetSeveritiesToIgnore(severitiesToIgnore []string)

	GetFilesOrPathsToIgnore() []string
	SetFilesOrPathsToIgnore(filesOrPaths []string)

	GetReturnErrorIfFoundVulnerability() bool
	SetReturnErrorIfFoundVulnerability(returnError bool)

	GetProjectPath() string
	SetProjectPath(projectPath string)

	GetWorkDir() *workdir.WorkDir
	SetWorkDir(toParse interface{})

	GetEnableGitHistoryAnalysis() bool
	SetEnableGitHistoryAnalysis(enableGitHistoryAnalysis bool)

	GetCertInsecureSkipVerify() bool
	SetCertInsecureSkipVerify(certInsecureSkipVerify bool)

	GetCertPath() string
	SetCertPath(certPath string)

	GetEnableCommitAuthor() bool
	SetEnableCommitAuthor(isEnable bool)

	GetRepositoryName() string
	SetRepositoryName(repositoryName string)

	GetRiskAcceptHashes() (output []string)
	SetRiskAcceptHashes(riskAccept []string)

	GetFalsePositiveHashes() (output []string)
	SetFalsePositiveHashes(falsePositive []string)

	GetHeaders() (headers map[string]string)
	SetHeaders(headers interface{})

	GetContainerBindProjectPath() string
	SetContainerBindProjectPath(containerBindProjectPath string)

	GetIsTimeout() bool
	SetIsTimeout(isTimeout bool)

	GetToolsConfig() toolsconfig.MapToolConfig
	SetToolsConfig(toolsConfig interface{})

	GetDisableDocker() bool
	SetDisableDocker(disableDocker bool)

	GetEnableInformationSeverity() bool
	SetEnableInformationSeverity(enableInformationSeverity bool)

	GetCustomRulesPath() string
	SetCustomRulesPath(customRulesPath string)

	IsEmptyRepositoryAuthorization() bool
	ToBytes(isMarshalIndent bool) (bytes []byte)
	ToMapLowerCase() map[string]interface{}
	NormalizeConfigs() IConfig

	GetCustomImages() customImages.CustomImages
	SetCustomImages(configData interface{})

	SetShowVulnerabilitiesTypes(vulnerabilitiesTypes []string)
	GetShowVulnerabilitiesTypes() []string

	GetEnableOwaspDependencyCheck() bool
	SetEnableOwaspDependencyCheck(enableOwaspDependencyCheck bool)

	GetEnableShellCheck() bool
	SetEnableShellCheck(enableShellCheck bool)
}
