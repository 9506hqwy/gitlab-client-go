package gitlab

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestEditorFn func(ctx context.Context, req *http.Request) error
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}
type ClientOption func(*Client) error
type ClientInterface interface {
	// GetApiV4AdminBatchedBackgroundMigrations request
	GetApiV4AdminBatchedBackgroundMigrations(ctx context.Context, params *GetApiV4AdminBatchedBackgroundMigrationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4AdminBatchedBackgroundMigrationsId request
	GetApiV4AdminBatchedBackgroundMigrationsId(ctx context.Context, id int32, params *GetApiV4AdminBatchedBackgroundMigrationsIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithBody request with any body
	PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4AdminBatchedBackgroundMigrationsIdPause(ctx context.Context, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithBody request with any body
	PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4AdminBatchedBackgroundMigrationsIdResume(ctx context.Context, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4AdminCiVariables request
	GetApiV4AdminCiVariables(ctx context.Context, params *GetApiV4AdminCiVariablesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4AdminCiVariablesWithBody request with any body
	PostApiV4AdminCiVariablesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4AdminCiVariables(ctx context.Context, body PostApiV4AdminCiVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4AdminCiVariablesKey request
	DeleteApiV4AdminCiVariablesKey(ctx context.Context, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4AdminCiVariablesKey request
	GetApiV4AdminCiVariablesKey(ctx context.Context, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4AdminCiVariablesKeyWithBody request with any body
	PutApiV4AdminCiVariablesKeyWithBody(ctx context.Context, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4AdminCiVariablesKey(ctx context.Context, key string, body PutApiV4AdminCiVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4AdminClusters request
	GetApiV4AdminClusters(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4AdminClustersAddWithBody request with any body
	PostApiV4AdminClustersAddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4AdminClustersAdd(ctx context.Context, body PostApiV4AdminClustersAddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4AdminClustersClusterId request
	DeleteApiV4AdminClustersClusterId(ctx context.Context, clusterId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4AdminClustersClusterId request
	GetApiV4AdminClustersClusterId(ctx context.Context, clusterId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4AdminClustersClusterIdWithBody request with any body
	PutApiV4AdminClustersClusterIdWithBody(ctx context.Context, clusterId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4AdminClustersClusterId(ctx context.Context, clusterId int32, body PutApiV4AdminClustersClusterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName request
	GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName(ctx context.Context, databaseName string, tableName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4AdminMigrationsTimestampMarkWithBody request with any body
	PostApiV4AdminMigrationsTimestampMarkWithBody(ctx context.Context, timestamp int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4AdminMigrationsTimestampMark(ctx context.Context, timestamp int32, body PostApiV4AdminMigrationsTimestampMarkJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ApplicationAppearance request
	GetApiV4ApplicationAppearance(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ApplicationAppearanceWithBody request with any body
	PutApiV4ApplicationAppearanceWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ApplicationPlanLimits request
	GetApiV4ApplicationPlanLimits(ctx context.Context, params *GetApiV4ApplicationPlanLimitsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ApplicationPlanLimitsWithBody request with any body
	PutApiV4ApplicationPlanLimitsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ApplicationPlanLimits(ctx context.Context, body PutApiV4ApplicationPlanLimitsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ApplicationStatistics request
	GetApiV4ApplicationStatistics(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Applications request
	GetApiV4Applications(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ApplicationsWithBody request with any body
	PostApiV4ApplicationsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Applications(ctx context.Context, body PostApiV4ApplicationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ApplicationsId request
	DeleteApiV4ApplicationsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ApplicationsIdRenewSecret request
	PostApiV4ApplicationsIdRenewSecret(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Avatar request
	GetApiV4Avatar(ctx context.Context, params *GetApiV4AvatarParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4BroadcastMessages request
	GetApiV4BroadcastMessages(ctx context.Context, params *GetApiV4BroadcastMessagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4BroadcastMessagesWithBody request with any body
	PostApiV4BroadcastMessagesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4BroadcastMessages(ctx context.Context, body PostApiV4BroadcastMessagesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4BroadcastMessagesId request
	DeleteApiV4BroadcastMessagesId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4BroadcastMessagesId request
	GetApiV4BroadcastMessagesId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4BroadcastMessagesIdWithBody request with any body
	PutApiV4BroadcastMessagesIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4BroadcastMessagesId(ctx context.Context, id int32, body PutApiV4BroadcastMessagesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4BulkImports request
	GetApiV4BulkImports(ctx context.Context, params *GetApiV4BulkImportsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4BulkImportsWithBody request with any body
	PostApiV4BulkImportsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4BulkImportsWithFormdataBody(ctx context.Context, body PostApiV4BulkImportsFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4BulkImportsEntities request
	GetApiV4BulkImportsEntities(ctx context.Context, params *GetApiV4BulkImportsEntitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4BulkImportsImportId request
	GetApiV4BulkImportsImportId(ctx context.Context, importId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4BulkImportsImportIdCancel request
	PostApiV4BulkImportsImportIdCancel(ctx context.Context, importId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4BulkImportsImportIdEntities request
	GetApiV4BulkImportsImportIdEntities(ctx context.Context, importId int32, params *GetApiV4BulkImportsImportIdEntitiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4BulkImportsImportIdEntitiesEntityId request
	GetApiV4BulkImportsImportIdEntitiesEntityId(ctx context.Context, importId int32, entityId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4BulkImportsImportIdEntitiesEntityIdFailures request
	GetApiV4BulkImportsImportIdEntitiesEntityIdFailures(ctx context.Context, importId int32, entityId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ContainerRegistryEventEvents request
	PostApiV4ContainerRegistryEventEvents(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4DeployKeys request
	GetApiV4DeployKeys(ctx context.Context, params *GetApiV4DeployKeysParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4DeployKeysWithBody request with any body
	PostApiV4DeployKeysWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4DeployKeys(ctx context.Context, body PostApiV4DeployKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4DeployTokens request
	GetApiV4DeployTokens(ctx context.Context, params *GetApiV4DeployTokensParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4DiscoverCertBasedClusters request
	GetApiV4DiscoverCertBasedClusters(ctx context.Context, params *GetApiV4DiscoverCertBasedClustersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Events request
	GetApiV4Events(ctx context.Context, params *GetApiV4EventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4FeatureFlagsUnleashProjectId request
	GetApiV4FeatureFlagsUnleashProjectId(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4FeatureFlagsUnleashProjectIdClientFeatures request
	GetApiV4FeatureFlagsUnleashProjectIdClientFeatures(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithBody request with any body
	PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithBody(ctx context.Context, projectId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4FeatureFlagsUnleashProjectIdClientMetrics(ctx context.Context, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithBody request with any body
	PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithBody(ctx context.Context, projectId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4FeatureFlagsUnleashProjectIdClientRegister(ctx context.Context, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4FeatureFlagsUnleashProjectIdFeatures request
	GetApiV4FeatureFlagsUnleashProjectIdFeatures(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Features request
	GetApiV4Features(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4FeaturesDefinitions request
	GetApiV4FeaturesDefinitions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4FeaturesName request
	DeleteApiV4FeaturesName(ctx context.Context, name int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4FeaturesNameWithBody request with any body
	PostApiV4FeaturesNameWithBody(ctx context.Context, name int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4FeaturesName(ctx context.Context, name int32, body PostApiV4FeaturesNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GeoNodeProxyIdGraphql request
	PostApiV4GeoNodeProxyIdGraphql(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GeoProxy request
	GetApiV4GeoProxy(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GeoProxyGitSshInfoRefsReceivePackWithBody request with any body
	PostApiV4GeoProxyGitSshInfoRefsReceivePackWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GeoProxyGitSshInfoRefsReceivePack(ctx context.Context, body PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GeoProxyGitSshInfoRefsUploadPackWithBody request with any body
	PostApiV4GeoProxyGitSshInfoRefsUploadPackWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GeoProxyGitSshInfoRefsUploadPack(ctx context.Context, body PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GeoProxyGitSshReceivePackWithBody request with any body
	PostApiV4GeoProxyGitSshReceivePackWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GeoProxyGitSshReceivePack(ctx context.Context, body PostApiV4GeoProxyGitSshReceivePackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GeoProxyGitSshUploadPackWithBody request with any body
	PostApiV4GeoProxyGitSshUploadPackWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GeoProxyGitSshUploadPack(ctx context.Context, body PostApiV4GeoProxyGitSshUploadPackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GeoRepositoriesGlRepositoryPipelineRefs request
	GetApiV4GeoRepositoriesGlRepositoryPipelineRefs(ctx context.Context, glRepository string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GeoRetrieveReplicableNameReplicableId request
	GetApiV4GeoRetrieveReplicableNameReplicableId(ctx context.Context, replicableName string, replicableId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GeoStatusWithBody request with any body
	PostApiV4GeoStatusWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GeoStatus(ctx context.Context, body PostApiV4GeoStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupIdPackagesComposerpackageName request
	GetApiV4GroupIdPackagesComposerpackageName(ctx context.Context, id string, params *GetApiV4GroupIdPackagesComposerpackageNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupIdPackagesComposerPSha request
	GetApiV4GroupIdPackagesComposerPSha(ctx context.Context, id string, sha string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupIdPackagesComposerP2packageName request
	GetApiV4GroupIdPackagesComposerP2packageName(ctx context.Context, id string, params *GetApiV4GroupIdPackagesComposerP2packageNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupIdPackagesComposerPackages request
	GetApiV4GroupIdPackagesComposerPackages(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Groups request
	GetApiV4Groups(ctx context.Context, params *GetApiV4GroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsWithBody request with any body
	PostApiV4GroupsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Groups(ctx context.Context, body PostApiV4GroupsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsImportWithBody request with any body
	PostApiV4GroupsImportWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsImportAuthorize request
	PostApiV4GroupsImportAuthorize(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsId request
	DeleteApiV4GroupsId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsId request
	GetApiV4GroupsId(ctx context.Context, id string, params *GetApiV4GroupsIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdWithBody request with any body
	PutApiV4GroupsIdWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsId(ctx context.Context, id string, body PutApiV4GroupsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdDebianDistributions request
	GetApiV4GroupsIdDebianDistributions(ctx context.Context, id string, params *GetApiV4GroupsIdDebianDistributionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdDebianDistributionsWithBody request with any body
	PostApiV4GroupsIdDebianDistributionsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdDebianDistributions(ctx context.Context, id string, body PostApiV4GroupsIdDebianDistributionsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdDebianDistributionsCodename request
	DeleteApiV4GroupsIdDebianDistributionsCodename(ctx context.Context, id string, codename string, params *DeleteApiV4GroupsIdDebianDistributionsCodenameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdDebianDistributionsCodename request
	GetApiV4GroupsIdDebianDistributionsCodename(ctx context.Context, id string, codename string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdDebianDistributionsCodenameWithBody request with any body
	PutApiV4GroupsIdDebianDistributionsCodenameWithBody(ctx context.Context, id string, codename string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdDebianDistributionsCodename(ctx context.Context, id string, codename string, body PutApiV4GroupsIdDebianDistributionsCodenameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdDebianDistributionsCodenameKeyAsc request
	GetApiV4GroupsIdDebianDistributionsCodenameKeyAsc(ctx context.Context, id string, codename string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionInrelease request
	GetApiV4GroupsIdPackagesDebianDistsdistributionInrelease(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionInreleaseParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionRelease request
	GetApiV4GroupsIdPackagesDebianDistsdistributionRelease(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpg request
	GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpg(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpgParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages(ctx context.Context, id string, component string, architecture string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256 request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256(ctx context.Context, id string, component string, architecture string, fileSha256 int32, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages(ctx context.Context, id string, component string, architecture string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256 request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256(ctx context.Context, id string, component string, architecture string, fileSha256 int32, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSources request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSources(ctx context.Context, id string, component string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSourcesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256 request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256(ctx context.Context, id string, component string, fileSha256 int32, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName request
	GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileName(ctx context.Context, id string, distribution string, projectId int32, letter string, packageName string, packageVersion string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesMavenpathFileName request
	GetApiV4GroupsIdPackagesMavenpathFileName(ctx context.Context, id string, fileName string, params *GetApiV4GroupsIdPackagesMavenpathFileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesNpmpackageName request
	GetApiV4GroupsIdPackagesNpmpackageName(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesNpmpackageNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk request
	PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuick request
	PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuick(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTags request
	GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTags(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag request
	DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag(ctx context.Context, id string, tag string, params *DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagWithBody request with any body
	PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagWithBody(ctx context.Context, id string, tag string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTag(ctx context.Context, id string, tag string, body PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesNugetIndex request
	GetApiV4GroupsIdPackagesNugetIndex(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesNugetQuery request
	GetApiV4GroupsIdPackagesNugetQuery(ctx context.Context, id int32, params *GetApiV4GroupsIdPackagesNugetQueryParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName request
	GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName(ctx context.Context, id int32, params *GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesNugetV2 request
	GetApiV4GroupsIdPackagesNugetV2(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesNugetV2Metadata request
	GetApiV4GroupsIdPackagesNugetV2Metadata(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifier request
	GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifier(ctx context.Context, id int32, sha256 string, params *GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifierParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesPypiSimple request
	GetApiV4GroupsIdPackagesPypiSimple(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackagesPypiSimplepackageName request
	GetApiV4GroupsIdPackagesPypiSimplepackageName(ctx context.Context, id int32, params *GetApiV4GroupsIdPackagesPypiSimplepackageNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdAccessRequests request
	GetApiV4GroupsIdAccessRequests(ctx context.Context, id string, params *GetApiV4GroupsIdAccessRequestsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdAccessRequests request
	PostApiV4GroupsIdAccessRequests(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdAccessRequestsUserId request
	DeleteApiV4GroupsIdAccessRequestsUserId(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdAccessRequestsUserIdApproveWithBody request with any body
	PutApiV4GroupsIdAccessRequestsUserIdApproveWithBody(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdAccessRequestsUserIdApprove(ctx context.Context, id string, userId int32, body PutApiV4GroupsIdAccessRequestsUserIdApproveJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdAccessTokensSelfRotateWithBody request with any body
	PostApiV4GroupsIdAccessTokensSelfRotateWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdAccessTokensSelfRotate(ctx context.Context, id string, body PostApiV4GroupsIdAccessTokensSelfRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdArchive request
	PostApiV4GroupsIdArchive(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdAuditEvents request
	GetApiV4GroupsIdAuditEvents(ctx context.Context, id int32, params *GetApiV4GroupsIdAuditEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdAuditEventsAuditEventId request
	GetApiV4GroupsIdAuditEventsAuditEventId(ctx context.Context, id int32, auditEventId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdAvatar request
	GetApiV4GroupsIdAvatar(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdBadges request
	GetApiV4GroupsIdBadges(ctx context.Context, id string, params *GetApiV4GroupsIdBadgesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdBadgesWithBody request with any body
	PostApiV4GroupsIdBadgesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdBadges(ctx context.Context, id string, body PostApiV4GroupsIdBadgesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdBadgesRender request
	GetApiV4GroupsIdBadgesRender(ctx context.Context, id string, params *GetApiV4GroupsIdBadgesRenderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdBadgesBadgeId request
	DeleteApiV4GroupsIdBadgesBadgeId(ctx context.Context, id string, badgeId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdBadgesBadgeId request
	GetApiV4GroupsIdBadgesBadgeId(ctx context.Context, id string, badgeId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdBadgesBadgeIdWithBody request with any body
	PutApiV4GroupsIdBadgesBadgeIdWithBody(ctx context.Context, id string, badgeId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdBadgesBadgeId(ctx context.Context, id string, badgeId int32, body PutApiV4GroupsIdBadgesBadgeIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdBillableMembers request
	GetApiV4GroupsIdBillableMembers(ctx context.Context, id string, params *GetApiV4GroupsIdBillableMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdBillableMembersUserId request
	DeleteApiV4GroupsIdBillableMembersUserId(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdBillableMembersUserIdIndirect request
	GetApiV4GroupsIdBillableMembersUserIdIndirect(ctx context.Context, id string, userId int32, params *GetApiV4GroupsIdBillableMembersUserIdIndirectParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdBillableMembersUserIdMemberships request
	GetApiV4GroupsIdBillableMembersUserIdMemberships(ctx context.Context, id string, userId int32, params *GetApiV4GroupsIdBillableMembersUserIdMembershipsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdClusters request
	GetApiV4GroupsIdClusters(ctx context.Context, id string, params *GetApiV4GroupsIdClustersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdClustersUserWithBody request with any body
	PostApiV4GroupsIdClustersUserWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdClustersUser(ctx context.Context, id string, body PostApiV4GroupsIdClustersUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdClustersClusterId request
	DeleteApiV4GroupsIdClustersClusterId(ctx context.Context, id string, clusterId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdClustersClusterId request
	GetApiV4GroupsIdClustersClusterId(ctx context.Context, id string, clusterId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdClustersClusterIdWithBody request with any body
	PutApiV4GroupsIdClustersClusterIdWithBody(ctx context.Context, id string, clusterId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdClustersClusterId(ctx context.Context, id string, clusterId int32, body PutApiV4GroupsIdClustersClusterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdCustomAttributes request
	GetApiV4GroupsIdCustomAttributes(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdCustomAttributesKey request
	DeleteApiV4GroupsIdCustomAttributesKey(ctx context.Context, id int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdCustomAttributesKey request
	GetApiV4GroupsIdCustomAttributesKey(ctx context.Context, id int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdCustomAttributesKeyWithBody request with any body
	PutApiV4GroupsIdCustomAttributesKeyWithBody(ctx context.Context, id int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdCustomAttributesKey(ctx context.Context, id int32, key string, body PutApiV4GroupsIdCustomAttributesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdDependencyProxyCache request
	DeleteApiV4GroupsIdDependencyProxyCache(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdDeployTokens request
	GetApiV4GroupsIdDeployTokens(ctx context.Context, id int32, params *GetApiV4GroupsIdDeployTokensParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdDeployTokensWithBody request with any body
	PostApiV4GroupsIdDeployTokensWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdDeployTokens(ctx context.Context, id int32, body PostApiV4GroupsIdDeployTokensJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdDeployTokensTokenId request
	DeleteApiV4GroupsIdDeployTokensTokenId(ctx context.Context, id int32, tokenId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdDeployTokensTokenId request
	GetApiV4GroupsIdDeployTokensTokenId(ctx context.Context, id int32, tokenId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdDescendantGroups request
	GetApiV4GroupsIdDescendantGroups(ctx context.Context, id string, params *GetApiV4GroupsIdDescendantGroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdEpicsEpicIidAwardEmoji request
	GetApiV4GroupsIdEpicsEpicIidAwardEmoji(ctx context.Context, id string, epicIid int32, params *GetApiV4GroupsIdEpicsEpicIidAwardEmojiParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdEpicsEpicIidAwardEmojiWithBody request with any body
	PostApiV4GroupsIdEpicsEpicIidAwardEmojiWithBody(ctx context.Context, id int32, epicIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdEpicsEpicIidAwardEmoji(ctx context.Context, id int32, epicIid int32, body PostApiV4GroupsIdEpicsEpicIidAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdEpicsEpicIidAwardEmojiAwardId request
	DeleteApiV4GroupsIdEpicsEpicIidAwardEmojiAwardId(ctx context.Context, id int32, epicIid int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdEpicsEpicIidAwardEmojiAwardId request
	GetApiV4GroupsIdEpicsEpicIidAwardEmojiAwardId(ctx context.Context, id int32, epicIid int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmoji request
	GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(ctx context.Context, id int32, epicIid int32, noteId int32, params *GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiWithBody request with any body
	PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiWithBody(ctx context.Context, id int32, epicIid int32, noteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmoji(ctx context.Context, id int32, epicIid int32, noteId int32, body PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId request
	DeleteApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(ctx context.Context, id int32, epicIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId request
	GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardId(ctx context.Context, id int32, epicIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdExport request
	PostApiV4GroupsIdExport(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdExportDownload request
	GetApiV4GroupsIdExportDownload(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdExportRelationsWithBody request with any body
	PostApiV4GroupsIdExportRelationsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdExportRelations(ctx context.Context, id string, body PostApiV4GroupsIdExportRelationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdExportRelationsDownload request
	GetApiV4GroupsIdExportRelationsDownload(ctx context.Context, id string, params *GetApiV4GroupsIdExportRelationsDownloadParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdExportRelationsStatus request
	GetApiV4GroupsIdExportRelationsStatus(ctx context.Context, id string, params *GetApiV4GroupsIdExportRelationsStatusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdGroupsShared request
	GetApiV4GroupsIdGroupsShared(ctx context.Context, id string, params *GetApiV4GroupsIdGroupsSharedParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdIntegrations request
	GetApiV4GroupsIdIntegrations(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsAppleAppStoreWithBody request with any body
	PutApiV4GroupsIdIntegrationsAppleAppStoreWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsAppleAppStore(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsAppleAppStoreJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsAsanaWithBody request with any body
	PutApiV4GroupsIdIntegrationsAsanaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsAsana(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsAsanaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsAssemblaWithBody request with any body
	PutApiV4GroupsIdIntegrationsAssemblaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsAssembla(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsAssemblaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsBambooWithBody request with any body
	PutApiV4GroupsIdIntegrationsBambooWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsBamboo(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsBambooJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsBugzillaWithBody request with any body
	PutApiV4GroupsIdIntegrationsBugzillaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsBugzilla(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsBugzillaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsBuildkiteWithBody request with any body
	PutApiV4GroupsIdIntegrationsBuildkiteWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsBuildkite(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsBuildkiteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsCampfireWithBody request with any body
	PutApiV4GroupsIdIntegrationsCampfireWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsCampfire(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsCampfireJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsClickupWithBody request with any body
	PutApiV4GroupsIdIntegrationsClickupWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsClickup(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsClickupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsConfluenceWithBody request with any body
	PutApiV4GroupsIdIntegrationsConfluenceWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsConfluence(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsConfluenceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsCustomIssueTrackerWithBody request with any body
	PutApiV4GroupsIdIntegrationsCustomIssueTrackerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsCustomIssueTracker(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsCustomIssueTrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsDatadogWithBody request with any body
	PutApiV4GroupsIdIntegrationsDatadogWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsDatadog(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsDatadogJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsDiffblueCoverWithBody request with any body
	PutApiV4GroupsIdIntegrationsDiffblueCoverWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsDiffblueCover(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsDiffblueCoverJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsDiscordWithBody request with any body
	PutApiV4GroupsIdIntegrationsDiscordWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsDiscord(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsDiscordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsDroneCiWithBody request with any body
	PutApiV4GroupsIdIntegrationsDroneCiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsDroneCi(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsDroneCiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsEmailsOnPushWithBody request with any body
	PutApiV4GroupsIdIntegrationsEmailsOnPushWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsEmailsOnPush(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsEmailsOnPushJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsEwmWithBody request with any body
	PutApiV4GroupsIdIntegrationsEwmWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsEwm(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsEwmJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsExternalWikiWithBody request with any body
	PutApiV4GroupsIdIntegrationsExternalWikiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsExternalWiki(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsExternalWikiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsGitGuardianWithBody request with any body
	PutApiV4GroupsIdIntegrationsGitGuardianWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsGitGuardian(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGitGuardianJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsGithubWithBody request with any body
	PutApiV4GroupsIdIntegrationsGithubWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsGithub(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsGitlabSlackApplicationWithBody request with any body
	PutApiV4GroupsIdIntegrationsGitlabSlackApplicationWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsGitlabSlackApplication(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGitlabSlackApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithBody request with any body
	PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistry(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithBody request with any body
	PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederation(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsGooglePlayWithBody request with any body
	PutApiV4GroupsIdIntegrationsGooglePlayWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsGooglePlay(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGooglePlayJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsHangoutsChatWithBody request with any body
	PutApiV4GroupsIdIntegrationsHangoutsChatWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsHangoutsChat(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsHangoutsChatJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsHarborWithBody request with any body
	PutApiV4GroupsIdIntegrationsHarborWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsHarbor(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsHarborJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsIrkerWithBody request with any body
	PutApiV4GroupsIdIntegrationsIrkerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsIrker(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsIrkerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsJenkinsWithBody request with any body
	PutApiV4GroupsIdIntegrationsJenkinsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsJenkins(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsJenkinsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsJiraWithBody request with any body
	PutApiV4GroupsIdIntegrationsJiraWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsJira(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsJiraJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsJiraCloudAppWithBody request with any body
	PutApiV4GroupsIdIntegrationsJiraCloudAppWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsJiraCloudApp(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsJiraCloudAppJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsMatrixWithBody request with any body
	PutApiV4GroupsIdIntegrationsMatrixWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsMatrix(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMatrixJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsMattermostWithBody request with any body
	PutApiV4GroupsIdIntegrationsMattermostWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsMattermost(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMattermostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsMattermostSlashCommandsWithBody request with any body
	PutApiV4GroupsIdIntegrationsMattermostSlashCommandsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsMattermostSlashCommands(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMattermostSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsMicrosoftTeamsWithBody request with any body
	PutApiV4GroupsIdIntegrationsMicrosoftTeamsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsMicrosoftTeams(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMicrosoftTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsMockCiWithBody request with any body
	PutApiV4GroupsIdIntegrationsMockCiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsMockCi(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMockCiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsMockMonitoringWithBody request with any body
	PutApiV4GroupsIdIntegrationsMockMonitoringWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsMockMonitoring(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMockMonitoringJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsPackagistWithBody request with any body
	PutApiV4GroupsIdIntegrationsPackagistWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsPackagist(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPackagistJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsPhorgeWithBody request with any body
	PutApiV4GroupsIdIntegrationsPhorgeWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsPhorge(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPhorgeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsPipelinesEmailWithBody request with any body
	PutApiV4GroupsIdIntegrationsPipelinesEmailWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsPipelinesEmail(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPipelinesEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsPivotaltrackerWithBody request with any body
	PutApiV4GroupsIdIntegrationsPivotaltrackerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsPivotaltracker(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPivotaltrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsPumbleWithBody request with any body
	PutApiV4GroupsIdIntegrationsPumbleWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsPumble(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPumbleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsPushoverWithBody request with any body
	PutApiV4GroupsIdIntegrationsPushoverWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsPushover(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPushoverJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsRedmineWithBody request with any body
	PutApiV4GroupsIdIntegrationsRedmineWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsRedmine(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsRedmineJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsSlackWithBody request with any body
	PutApiV4GroupsIdIntegrationsSlackWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsSlack(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsSlackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsSlackSlashCommandsWithBody request with any body
	PutApiV4GroupsIdIntegrationsSlackSlashCommandsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsSlackSlashCommands(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsSlackSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsSquashTmWithBody request with any body
	PutApiV4GroupsIdIntegrationsSquashTmWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsSquashTm(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsSquashTmJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsTeamcityWithBody request with any body
	PutApiV4GroupsIdIntegrationsTeamcityWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsTeamcity(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsTeamcityJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsTelegramWithBody request with any body
	PutApiV4GroupsIdIntegrationsTelegramWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsTelegram(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsTelegramJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsUnifyCircuitWithBody request with any body
	PutApiV4GroupsIdIntegrationsUnifyCircuitWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsUnifyCircuit(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsUnifyCircuitJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsWebexTeamsWithBody request with any body
	PutApiV4GroupsIdIntegrationsWebexTeamsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsWebexTeams(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsWebexTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsYoutrackWithBody request with any body
	PutApiV4GroupsIdIntegrationsYoutrackWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsYoutrack(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsYoutrackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdIntegrationsZentaoWithBody request with any body
	PutApiV4GroupsIdIntegrationsZentaoWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdIntegrationsZentao(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsZentaoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdIntegrationsSlug request
	DeleteApiV4GroupsIdIntegrationsSlug(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdIntegrationsSlug request
	GetApiV4GroupsIdIntegrationsSlug(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdInvitations request
	GetApiV4GroupsIdInvitations(ctx context.Context, id string, params *GetApiV4GroupsIdInvitationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdInvitationsWithBody request with any body
	PostApiV4GroupsIdInvitationsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdInvitations(ctx context.Context, id string, body PostApiV4GroupsIdInvitationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdInvitationsEmail request
	DeleteApiV4GroupsIdInvitationsEmail(ctx context.Context, id string, email string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdInvitationsEmailWithBody request with any body
	PutApiV4GroupsIdInvitationsEmailWithBody(ctx context.Context, id string, email string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdInvitationsEmail(ctx context.Context, id string, email string, body PutApiV4GroupsIdInvitationsEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdInvitedGroups request
	GetApiV4GroupsIdInvitedGroups(ctx context.Context, id string, params *GetApiV4GroupsIdInvitedGroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdIssues request
	GetApiV4GroupsIdIssues(ctx context.Context, id string, params *GetApiV4GroupsIdIssuesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdLdapSync request
	PostApiV4GroupsIdLdapSync(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdMembers request
	GetApiV4GroupsIdMembers(ctx context.Context, id string, params *GetApiV4GroupsIdMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdMembersWithBody request with any body
	PostApiV4GroupsIdMembersWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdMembers(ctx context.Context, id string, body PostApiV4GroupsIdMembersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdMembersAll request
	GetApiV4GroupsIdMembersAll(ctx context.Context, id string, params *GetApiV4GroupsIdMembersAllParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdMembersAllUserId request
	GetApiV4GroupsIdMembersAllUserId(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdMembersApproveAll request
	PostApiV4GroupsIdMembersApproveAll(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdMembersMemberIdApprove request
	PutApiV4GroupsIdMembersMemberIdApprove(ctx context.Context, id string, memberId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdMembersUserId request
	DeleteApiV4GroupsIdMembersUserId(ctx context.Context, id string, userId int32, params *DeleteApiV4GroupsIdMembersUserIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdMembersUserId request
	GetApiV4GroupsIdMembersUserId(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdMembersUserIdWithBody request with any body
	PutApiV4GroupsIdMembersUserIdWithBody(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdMembersUserId(ctx context.Context, id string, userId int32, body PutApiV4GroupsIdMembersUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdMembersUserIdOverride request
	DeleteApiV4GroupsIdMembersUserIdOverride(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdMembersUserIdOverride request
	PostApiV4GroupsIdMembersUserIdOverride(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdMembersUserIdStateWithBody request with any body
	PutApiV4GroupsIdMembersUserIdStateWithBody(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdMembersUserIdState(ctx context.Context, id string, userId int32, body PutApiV4GroupsIdMembersUserIdStateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdMergeRequests request
	GetApiV4GroupsIdMergeRequests(ctx context.Context, id string, params *GetApiV4GroupsIdMergeRequestsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPackages request
	GetApiV4GroupsIdPackages(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPendingMembers request
	GetApiV4GroupsIdPendingMembers(ctx context.Context, id string, params *GetApiV4GroupsIdPendingMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdPlaceholderReassignments request
	GetApiV4GroupsIdPlaceholderReassignments(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdPlaceholderReassignmentsWithBody request with any body
	PostApiV4GroupsIdPlaceholderReassignmentsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdPlaceholderReassignments(ctx context.Context, id string, body PostApiV4GroupsIdPlaceholderReassignmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdPlaceholderReassignmentsAuthorize request
	PostApiV4GroupsIdPlaceholderReassignmentsAuthorize(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdProjects request
	GetApiV4GroupsIdProjects(ctx context.Context, id string, params *GetApiV4GroupsIdProjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdProjectsShared request
	GetApiV4GroupsIdProjectsShared(ctx context.Context, id string, params *GetApiV4GroupsIdProjectsSharedParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdProjectsProjectId request
	PostApiV4GroupsIdProjectsProjectId(ctx context.Context, id string, projectId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdProvisionedUsers request
	GetApiV4GroupsIdProvisionedUsers(ctx context.Context, id int32, params *GetApiV4GroupsIdProvisionedUsersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdRegistryRepositories request
	GetApiV4GroupsIdRegistryRepositories(ctx context.Context, id string, params *GetApiV4GroupsIdRegistryRepositoriesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdReleases request
	GetApiV4GroupsIdReleases(ctx context.Context, id string, params *GetApiV4GroupsIdReleasesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdRestore request
	PostApiV4GroupsIdRestore(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdRunners request
	GetApiV4GroupsIdRunners(ctx context.Context, id string, params *GetApiV4GroupsIdRunnersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdRunnersResetRegistrationToken request
	PostApiV4GroupsIdRunnersResetRegistrationToken(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdSamlUsers request
	GetApiV4GroupsIdSamlUsers(ctx context.Context, id int32, params *GetApiV4GroupsIdSamlUsersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdShareWithBody request with any body
	PostApiV4GroupsIdShareWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdShare(ctx context.Context, id string, body PostApiV4GroupsIdShareJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdShareGroupId request
	DeleteApiV4GroupsIdShareGroupId(ctx context.Context, id string, groupId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdSshCertificates request
	GetApiV4GroupsIdSshCertificates(ctx context.Context, id int32, params *GetApiV4GroupsIdSshCertificatesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdSshCertificatesWithBody request with any body
	PostApiV4GroupsIdSshCertificatesWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdSshCertificates(ctx context.Context, id int32, body PostApiV4GroupsIdSshCertificatesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdSshCertificatesSshCertificatesId request
	DeleteApiV4GroupsIdSshCertificatesSshCertificatesId(ctx context.Context, id int32, sshCertificatesId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdSubgroups request
	GetApiV4GroupsIdSubgroups(ctx context.Context, id string, params *GetApiV4GroupsIdSubgroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdTokensRevokeWithBody request with any body
	PostApiV4GroupsIdTokensRevokeWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdTokensRevoke(ctx context.Context, id string, body PostApiV4GroupsIdTokensRevokeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdTransferWithBody request with any body
	PostApiV4GroupsIdTransferWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdTransfer(ctx context.Context, id string, body PostApiV4GroupsIdTransferJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdTransferLocations request
	GetApiV4GroupsIdTransferLocations(ctx context.Context, id string, params *GetApiV4GroupsIdTransferLocationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdUnarchive request
	PostApiV4GroupsIdUnarchive(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdUploads request
	GetApiV4GroupsIdUploads(ctx context.Context, id int32, params *GetApiV4GroupsIdUploadsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdUploadsSecretFilename request
	DeleteApiV4GroupsIdUploadsSecretFilename(ctx context.Context, id int32, secret string, filename string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdUploadsSecretFilename request
	GetApiV4GroupsIdUploadsSecretFilename(ctx context.Context, id int32, secret string, filename string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdUploadsUploadId request
	DeleteApiV4GroupsIdUploadsUploadId(ctx context.Context, id int32, uploadId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdUploadsUploadId request
	GetApiV4GroupsIdUploadsUploadId(ctx context.Context, id int32, uploadId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdUsers request
	GetApiV4GroupsIdUsers(ctx context.Context, id int32, params *GetApiV4GroupsIdUsersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdVariables request
	GetApiV4GroupsIdVariables(ctx context.Context, id string, params *GetApiV4GroupsIdVariablesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdVariablesWithBody request with any body
	PostApiV4GroupsIdVariablesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdVariables(ctx context.Context, id string, body PostApiV4GroupsIdVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdVariablesKey request
	DeleteApiV4GroupsIdVariablesKey(ctx context.Context, id string, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdVariablesKey request
	GetApiV4GroupsIdVariablesKey(ctx context.Context, id string, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdVariablesKeyWithBody request with any body
	PutApiV4GroupsIdVariablesKeyWithBody(ctx context.Context, id string, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdVariablesKey(ctx context.Context, id string, key string, body PutApiV4GroupsIdVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdWikis request
	GetApiV4GroupsIdWikis(ctx context.Context, id int32, params *GetApiV4GroupsIdWikisParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdWikisWithBody request with any body
	PostApiV4GroupsIdWikisWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdWikis(ctx context.Context, id int32, body PostApiV4GroupsIdWikisJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4GroupsIdWikisAttachmentsWithBody request with any body
	PostApiV4GroupsIdWikisAttachmentsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4GroupsIdWikisAttachments(ctx context.Context, id int32, body PostApiV4GroupsIdWikisAttachmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4GroupsIdWikisSlug request
	DeleteApiV4GroupsIdWikisSlug(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4GroupsIdWikisSlug request
	GetApiV4GroupsIdWikisSlug(ctx context.Context, id int32, slug string, params *GetApiV4GroupsIdWikisSlugParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4GroupsIdWikisSlugWithBody request with any body
	PutApiV4GroupsIdWikisSlugWithBody(ctx context.Context, id int32, slug int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4GroupsIdWikisSlug(ctx context.Context, id int32, slug int32, body PutApiV4GroupsIdWikisSlugJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Hooks request
	GetApiV4Hooks(ctx context.Context, params *GetApiV4HooksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4HooksWithBody request with any body
	PostApiV4HooksWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Hooks(ctx context.Context, body PostApiV4HooksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4HooksHookId request
	DeleteApiV4HooksHookId(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4HooksHookId request
	GetApiV4HooksHookId(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4HooksHookId request
	PostApiV4HooksHookId(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4HooksHookIdWithBody request with any body
	PutApiV4HooksHookIdWithBody(ctx context.Context, hookId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4HooksHookId(ctx context.Context, hookId int32, body PutApiV4HooksHookIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4HooksHookIdCustomHeadersKey request
	DeleteApiV4HooksHookIdCustomHeadersKey(ctx context.Context, hookId int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4HooksHookIdCustomHeadersKeyWithBody request with any body
	PutApiV4HooksHookIdCustomHeadersKeyWithBody(ctx context.Context, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4HooksHookIdCustomHeadersKey(ctx context.Context, hookId int32, key string, body PutApiV4HooksHookIdCustomHeadersKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4HooksHookIdUrlVariablesKey request
	DeleteApiV4HooksHookIdUrlVariablesKey(ctx context.Context, hookId int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4HooksHookIdUrlVariablesKeyWithBody request with any body
	PutApiV4HooksHookIdUrlVariablesKeyWithBody(ctx context.Context, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4HooksHookIdUrlVariablesKey(ctx context.Context, hookId int32, key string, body PutApiV4HooksHookIdUrlVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ImportBitbucketWithBody request with any body
	PostApiV4ImportBitbucketWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ImportBitbucket(ctx context.Context, body PostApiV4ImportBitbucketJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ImportBitbucketServerWithBody request with any body
	PostApiV4ImportBitbucketServerWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ImportBitbucketServer(ctx context.Context, body PostApiV4ImportBitbucketServerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ImportGithubWithBody request with any body
	PostApiV4ImportGithubWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ImportGithub(ctx context.Context, body PostApiV4ImportGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ImportGithubCancelWithBody request with any body
	PostApiV4ImportGithubCancelWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ImportGithubCancel(ctx context.Context, body PostApiV4ImportGithubCancelJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ImportGithubGistsWithBody request with any body
	PostApiV4ImportGithubGistsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ImportGithubGists(ctx context.Context, body PostApiV4ImportGithubGistsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4IntegrationsJiraConnectSubscriptionsWithBody request with any body
	PostApiV4IntegrationsJiraConnectSubscriptionsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4IntegrationsJiraConnectSubscriptions(ctx context.Context, body PostApiV4IntegrationsJiraConnectSubscriptionsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4IntegrationsSlackEventsWithBody request with any body
	PostApiV4IntegrationsSlackEventsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4IntegrationsSlackEvents(ctx context.Context, body PostApiV4IntegrationsSlackEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4IntegrationsSlackInteractions request
	PostApiV4IntegrationsSlackInteractions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4IntegrationsSlackOptions request
	PostApiV4IntegrationsSlackOptions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Issues request
	GetApiV4Issues(ctx context.Context, params *GetApiV4IssuesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4IssuesId request
	GetApiV4IssuesId(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Job request
	GetApiV4Job(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4JobAllowedAgents request
	GetApiV4JobAllowedAgents(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4JobsRequestWithBody request with any body
	PostApiV4JobsRequestWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4JobsRequest(ctx context.Context, body PostApiV4JobsRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4JobsIdWithBody request with any body
	PutApiV4JobsIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4JobsId(ctx context.Context, id int32, body PutApiV4JobsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4JobsIdArtifacts request
	GetApiV4JobsIdArtifacts(ctx context.Context, id int32, params *GetApiV4JobsIdArtifactsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4JobsIdArtifactsWithBody request with any body
	PostApiV4JobsIdArtifactsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4JobsIdArtifacts(ctx context.Context, id int32, body PostApiV4JobsIdArtifactsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4JobsIdArtifactsAuthorizeWithBody request with any body
	PostApiV4JobsIdArtifactsAuthorizeWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4JobsIdArtifactsAuthorize(ctx context.Context, id int32, body PostApiV4JobsIdArtifactsAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchApiV4JobsIdTraceWithBody request with any body
	PatchApiV4JobsIdTraceWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchApiV4JobsIdTrace(ctx context.Context, id int32, body PatchApiV4JobsIdTraceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Keys request
	GetApiV4Keys(ctx context.Context, params *GetApiV4KeysParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4KeysId request
	GetApiV4KeysId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4MarkdownWithBody request with any body
	PostApiV4MarkdownWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Markdown(ctx context.Context, body PostApiV4MarkdownJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4MergeRequests request
	GetApiV4MergeRequests(ctx context.Context, params *GetApiV4MergeRequestsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Metadata request
	GetApiV4Metadata(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Namespaces request
	GetApiV4Namespaces(ctx context.Context, params *GetApiV4NamespacesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4NamespacesStorageLimitExclusions request
	GetApiV4NamespacesStorageLimitExclusions(ctx context.Context, params *GetApiV4NamespacesStorageLimitExclusionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4NamespacesId request
	GetApiV4NamespacesId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4NamespacesIdWithBody request with any body
	PutApiV4NamespacesIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4NamespacesId(ctx context.Context, id int32, body PutApiV4NamespacesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4NamespacesIdExists request
	GetApiV4NamespacesIdExists(ctx context.Context, id string, params *GetApiV4NamespacesIdExistsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4NamespacesIdGitlabSubscription request
	GetApiV4NamespacesIdGitlabSubscription(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4NamespacesIdStorageLimitExclusion request
	DeleteApiV4NamespacesIdStorageLimitExclusion(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4NamespacesIdStorageLimitExclusionWithBody request with any body
	PostApiV4NamespacesIdStorageLimitExclusionWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4NamespacesIdStorageLimitExclusion(ctx context.Context, id int32, body PostApiV4NamespacesIdStorageLimitExclusionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4OrganizationsWithBody request with any body
	PostApiV4OrganizationsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Organizations(ctx context.Context, body PostApiV4OrganizationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1ConansSearch request
	GetApiV4PackagesConanV1ConansSearch(ctx context.Context, params *GetApiV4PackagesConanV1ConansSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel request
	DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls request
	PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls request
	PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName request
	GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithBody request with any body
	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithBody(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, body PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize request
	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName request
	GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithBody request with any body
	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithBody(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, body PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize request
	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1Ping request
	GetApiV4PackagesConanV1Ping(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1UsersAuthenticate request
	GetApiV4PackagesConanV1UsersAuthenticate(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesConanV1UsersCheckCredentials request
	GetApiV4PackagesConanV1UsersCheckCredentials(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesMavenpathFileName request
	GetApiV4PackagesMavenpathFileName(ctx context.Context, fileName string, params *GetApiV4PackagesMavenpathFileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesNpmpackageName request
	GetApiV4PackagesNpmpackageName(ctx context.Context, params *GetApiV4PackagesNpmpackageNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulk request
	PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulk(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4PackagesNpmNpmV1SecurityAuditsQuick request
	PostApiV4PackagesNpmNpmV1SecurityAuditsQuick(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesNpmPackagepackageNameDistTags request
	GetApiV4PackagesNpmPackagepackageNameDistTags(ctx context.Context, params *GetApiV4PackagesNpmPackagepackageNameDistTagsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4PackagesNpmPackagepackageNameDistTagsTag request
	DeleteApiV4PackagesNpmPackagepackageNameDistTagsTag(ctx context.Context, tag string, params *DeleteApiV4PackagesNpmPackagepackageNameDistTagsTagParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4PackagesNpmPackagepackageNameDistTagsTagWithBody request with any body
	PutApiV4PackagesNpmPackagepackageNameDistTagsTagWithBody(ctx context.Context, tag string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4PackagesNpmPackagepackageNameDistTagsTag(ctx context.Context, tag string, body PutApiV4PackagesNpmPackagepackageNameDistTagsTagJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystem(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersion request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersion(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, params *GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownload request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownload(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, params *GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownloadParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFile request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFile(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, params *GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFileParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownload(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersions(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PagesDomains request
	GetApiV4PagesDomains(ctx context.Context, params *GetApiV4PagesDomainsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PersonalAccessTokens request
	GetApiV4PersonalAccessTokens(ctx context.Context, params *GetApiV4PersonalAccessTokensParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4PersonalAccessTokensSelf request
	DeleteApiV4PersonalAccessTokensSelf(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PersonalAccessTokensSelf request
	GetApiV4PersonalAccessTokensSelf(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PersonalAccessTokensSelfAssociations request
	GetApiV4PersonalAccessTokensSelfAssociations(ctx context.Context, params *GetApiV4PersonalAccessTokensSelfAssociationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4PersonalAccessTokensSelfRotateWithBody request with any body
	PostApiV4PersonalAccessTokensSelfRotateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4PersonalAccessTokensSelfRotate(ctx context.Context, body PostApiV4PersonalAccessTokensSelfRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4PersonalAccessTokensId request
	DeleteApiV4PersonalAccessTokensId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4PersonalAccessTokensId request
	GetApiV4PersonalAccessTokensId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4PersonalAccessTokensIdRotateWithBody request with any body
	PostApiV4PersonalAccessTokensIdRotateWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4PersonalAccessTokensIdRotate(ctx context.Context, id int32, body PostApiV4PersonalAccessTokensIdRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Projects request
	GetApiV4Projects(ctx context.Context, params *GetApiV4ProjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsWithBody request with any body
	PostApiV4ProjectsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Projects(ctx context.Context, body PostApiV4ProjectsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsImportWithBody request with any body
	PostApiV4ProjectsImportWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsImportRelationWithBody request with any body
	PostApiV4ProjectsImportRelationWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsImportRelationAuthorize request
	PostApiV4ProjectsImportRelationAuthorize(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsImportAuthorize request
	PostApiV4ProjectsImportAuthorize(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsRemoteImportWithBody request with any body
	PostApiV4ProjectsRemoteImportWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsRemoteImportS3WithBody request with any body
	PostApiV4ProjectsRemoteImportS3WithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsUserUserIdWithBody request with any body
	PostApiV4ProjectsUserUserIdWithBody(ctx context.Context, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsUserUserId(ctx context.Context, userId int32, body PostApiV4ProjectsUserUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsId request
	DeleteApiV4ProjectsId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsId request
	GetApiV4ProjectsId(ctx context.Context, id string, params *GetApiV4ProjectsIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdWithBody request with any body
	PutApiV4ProjectsIdWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsId(ctx context.Context, id string, body PutApiV4ProjectsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRefRefTriggerPipelineWithBody request with any body
	PostApiV4ProjectsIdRefRefTriggerPipelineWithBody(ctx context.Context, id string, ref string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRefRefTriggerPipeline(ctx context.Context, id string, ref string, body PostApiV4ProjectsIdRefRefTriggerPipelineJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdAccessRequests request
	GetApiV4ProjectsIdAccessRequests(ctx context.Context, id string, params *GetApiV4ProjectsIdAccessRequestsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdAccessRequests request
	PostApiV4ProjectsIdAccessRequests(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdAccessRequestsUserId request
	DeleteApiV4ProjectsIdAccessRequestsUserId(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdAccessRequestsUserIdApproveWithBody request with any body
	PutApiV4ProjectsIdAccessRequestsUserIdApproveWithBody(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdAccessRequestsUserIdApprove(ctx context.Context, id string, userId int32, body PutApiV4ProjectsIdAccessRequestsUserIdApproveJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdAccessTokensSelfRotateWithBody request with any body
	PostApiV4ProjectsIdAccessTokensSelfRotateWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdAccessTokensSelfRotate(ctx context.Context, id string, body PostApiV4ProjectsIdAccessTokensSelfRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages request
	GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImages(ctx context.Context, id string, alertIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesWithBody request with any body
	PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesWithBody(ctx context.Context, id string, alertIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize request
	PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorize(ctx context.Context, id string, alertIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId request
	DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageId(ctx context.Context, id string, alertIid int32, metricImageId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdWithBody request with any body
	PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdWithBody(ctx context.Context, id string, alertIid int32, metricImageId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdArchive request
	PostApiV4ProjectsIdArchive(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdArtifacts request
	DeleteApiV4ProjectsIdArtifacts(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdAuditEvents request
	GetApiV4ProjectsIdAuditEvents(ctx context.Context, id int32, params *GetApiV4ProjectsIdAuditEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdAuditEventsAuditEventId request
	GetApiV4ProjectsIdAuditEventsAuditEventId(ctx context.Context, id int32, auditEventId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdAvatar request
	GetApiV4ProjectsIdAvatar(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdBadges request
	GetApiV4ProjectsIdBadges(ctx context.Context, id string, params *GetApiV4ProjectsIdBadgesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdBadgesWithBody request with any body
	PostApiV4ProjectsIdBadgesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdBadges(ctx context.Context, id string, body PostApiV4ProjectsIdBadgesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdBadgesRender request
	GetApiV4ProjectsIdBadgesRender(ctx context.Context, id string, params *GetApiV4ProjectsIdBadgesRenderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdBadgesBadgeId request
	DeleteApiV4ProjectsIdBadgesBadgeId(ctx context.Context, id string, badgeId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdBadgesBadgeId request
	GetApiV4ProjectsIdBadgesBadgeId(ctx context.Context, id string, badgeId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdBadgesBadgeIdWithBody request with any body
	PutApiV4ProjectsIdBadgesBadgeIdWithBody(ctx context.Context, id string, badgeId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdBadgesBadgeId(ctx context.Context, id string, badgeId int32, body PutApiV4ProjectsIdBadgesBadgeIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdCatalogPublishWithBody request with any body
	PostApiV4ProjectsIdCatalogPublishWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdCatalogPublish(ctx context.Context, id string, body PostApiV4ProjectsIdCatalogPublishJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdCiLint request
	GetApiV4ProjectsIdCiLint(ctx context.Context, id int32, params *GetApiV4ProjectsIdCiLintParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdCiLintWithBody request with any body
	PostApiV4ProjectsIdCiLintWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdCiLint(ctx context.Context, id int32, body PostApiV4ProjectsIdCiLintJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdClusterAgents request
	GetApiV4ProjectsIdClusterAgents(ctx context.Context, id string, params *GetApiV4ProjectsIdClusterAgentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdClusterAgentsWithBody request with any body
	PostApiV4ProjectsIdClusterAgentsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdClusterAgents(ctx context.Context, id string, body PostApiV4ProjectsIdClusterAgentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdClusterAgentsAgentId request
	DeleteApiV4ProjectsIdClusterAgentsAgentId(ctx context.Context, id string, agentId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdClusterAgentsAgentId request
	GetApiV4ProjectsIdClusterAgentsAgentId(ctx context.Context, id string, agentId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdClusterAgentsAgentIdTokens request
	GetApiV4ProjectsIdClusterAgentsAgentIdTokens(ctx context.Context, id string, agentId int32, params *GetApiV4ProjectsIdClusterAgentsAgentIdTokensParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdClusterAgentsAgentIdTokensWithBody request with any body
	PostApiV4ProjectsIdClusterAgentsAgentIdTokensWithBody(ctx context.Context, id string, agentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdClusterAgentsAgentIdTokens(ctx context.Context, id string, agentId int32, body PostApiV4ProjectsIdClusterAgentsAgentIdTokensJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId request
	DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId(ctx context.Context, id string, agentId int32, tokenId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId request
	GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenId(ctx context.Context, id string, agentId int32, tokenId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdClusters request
	GetApiV4ProjectsIdClusters(ctx context.Context, id string, params *GetApiV4ProjectsIdClustersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdClustersUserWithBody request with any body
	PostApiV4ProjectsIdClustersUserWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdClustersUser(ctx context.Context, id string, body PostApiV4ProjectsIdClustersUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdClustersClusterId request
	DeleteApiV4ProjectsIdClustersClusterId(ctx context.Context, id string, clusterId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdClustersClusterId request
	GetApiV4ProjectsIdClustersClusterId(ctx context.Context, id string, clusterId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdClustersClusterIdWithBody request with any body
	PutApiV4ProjectsIdClustersClusterIdWithBody(ctx context.Context, id string, clusterId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdClustersClusterId(ctx context.Context, id string, clusterId int32, body PutApiV4ProjectsIdClustersClusterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdCreateCiConfig request
	PostApiV4ProjectsIdCreateCiConfig(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdCustomAttributes request
	GetApiV4ProjectsIdCustomAttributes(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdCustomAttributesKey request
	DeleteApiV4ProjectsIdCustomAttributesKey(ctx context.Context, id int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdCustomAttributesKey request
	GetApiV4ProjectsIdCustomAttributesKey(ctx context.Context, id int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdCustomAttributesKeyWithBody request with any body
	PutApiV4ProjectsIdCustomAttributesKeyWithBody(ctx context.Context, id int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdCustomAttributesKey(ctx context.Context, id int32, key string, body PutApiV4ProjectsIdCustomAttributesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDebianDistributions request
	GetApiV4ProjectsIdDebianDistributions(ctx context.Context, id string, params *GetApiV4ProjectsIdDebianDistributionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdDebianDistributionsWithBody request with any body
	PostApiV4ProjectsIdDebianDistributionsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdDebianDistributions(ctx context.Context, id string, body PostApiV4ProjectsIdDebianDistributionsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdDebianDistributionsCodename request
	DeleteApiV4ProjectsIdDebianDistributionsCodename(ctx context.Context, id string, codename string, params *DeleteApiV4ProjectsIdDebianDistributionsCodenameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDebianDistributionsCodename request
	GetApiV4ProjectsIdDebianDistributionsCodename(ctx context.Context, id string, codename string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdDebianDistributionsCodenameWithBody request with any body
	PutApiV4ProjectsIdDebianDistributionsCodenameWithBody(ctx context.Context, id string, codename string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdDebianDistributionsCodename(ctx context.Context, id string, codename string, body PutApiV4ProjectsIdDebianDistributionsCodenameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDebianDistributionsCodenameKeyAsc request
	GetApiV4ProjectsIdDebianDistributionsCodenameKeyAsc(ctx context.Context, id string, codename string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDeployKeys request
	GetApiV4ProjectsIdDeployKeys(ctx context.Context, id string, params *GetApiV4ProjectsIdDeployKeysParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdDeployKeysWithBody request with any body
	PostApiV4ProjectsIdDeployKeysWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdDeployKeys(ctx context.Context, id string, body PostApiV4ProjectsIdDeployKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdDeployKeysKeyId request
	DeleteApiV4ProjectsIdDeployKeysKeyId(ctx context.Context, id string, keyId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDeployKeysKeyId request
	GetApiV4ProjectsIdDeployKeysKeyId(ctx context.Context, id string, keyId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdDeployKeysKeyIdWithBody request with any body
	PutApiV4ProjectsIdDeployKeysKeyIdWithBody(ctx context.Context, id string, keyId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdDeployKeysKeyId(ctx context.Context, id string, keyId int32, body PutApiV4ProjectsIdDeployKeysKeyIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdDeployKeysKeyIdEnable request
	PostApiV4ProjectsIdDeployKeysKeyIdEnable(ctx context.Context, id string, keyId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDeployTokens request
	GetApiV4ProjectsIdDeployTokens(ctx context.Context, id string, params *GetApiV4ProjectsIdDeployTokensParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdDeployTokensWithBody request with any body
	PostApiV4ProjectsIdDeployTokensWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdDeployTokens(ctx context.Context, id string, body PostApiV4ProjectsIdDeployTokensJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdDeployTokensTokenId request
	DeleteApiV4ProjectsIdDeployTokensTokenId(ctx context.Context, id string, tokenId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDeployTokensTokenId request
	GetApiV4ProjectsIdDeployTokensTokenId(ctx context.Context, id string, tokenId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDeployments request
	GetApiV4ProjectsIdDeployments(ctx context.Context, id string, params *GetApiV4ProjectsIdDeploymentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdDeploymentsWithBody request with any body
	PostApiV4ProjectsIdDeploymentsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdDeployments(ctx context.Context, id string, body PostApiV4ProjectsIdDeploymentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdDeploymentsDeploymentId request
	DeleteApiV4ProjectsIdDeploymentsDeploymentId(ctx context.Context, id string, deploymentId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDeploymentsDeploymentId request
	GetApiV4ProjectsIdDeploymentsDeploymentId(ctx context.Context, id string, deploymentId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdDeploymentsDeploymentIdWithBody request with any body
	PutApiV4ProjectsIdDeploymentsDeploymentIdWithBody(ctx context.Context, id string, deploymentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdDeploymentsDeploymentId(ctx context.Context, id string, deploymentId int32, body PutApiV4ProjectsIdDeploymentsDeploymentIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalWithBody request with any body
	PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalWithBody(ctx context.Context, id string, deploymentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdDeploymentsDeploymentIdApproval(ctx context.Context, id string, deploymentId int32, body PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequests request
	GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequests(ctx context.Context, id string, deploymentId int32, params *GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdEnvironments request
	GetApiV4ProjectsIdEnvironments(ctx context.Context, id string, params *GetApiV4ProjectsIdEnvironmentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdEnvironmentsWithBody request with any body
	PostApiV4ProjectsIdEnvironmentsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdEnvironments(ctx context.Context, id string, body PostApiV4ProjectsIdEnvironmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdEnvironmentsReviewApps request
	DeleteApiV4ProjectsIdEnvironmentsReviewApps(ctx context.Context, id string, params *DeleteApiV4ProjectsIdEnvironmentsReviewAppsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdEnvironmentsStopStaleWithBody request with any body
	PostApiV4ProjectsIdEnvironmentsStopStaleWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdEnvironmentsStopStale(ctx context.Context, id string, body PostApiV4ProjectsIdEnvironmentsStopStaleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdEnvironmentsEnvironmentId request
	DeleteApiV4ProjectsIdEnvironmentsEnvironmentId(ctx context.Context, id string, environmentId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdEnvironmentsEnvironmentId request
	GetApiV4ProjectsIdEnvironmentsEnvironmentId(ctx context.Context, id string, environmentId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdEnvironmentsEnvironmentIdWithBody request with any body
	PutApiV4ProjectsIdEnvironmentsEnvironmentIdWithBody(ctx context.Context, id string, environmentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdEnvironmentsEnvironmentId(ctx context.Context, id string, environmentId int32, body PutApiV4ProjectsIdEnvironmentsEnvironmentIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopWithBody request with any body
	PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopWithBody(ctx context.Context, id string, environmentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdEnvironmentsEnvironmentIdStop(ctx context.Context, id string, environmentId int32, body PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdErrorTrackingClientKeys request
	GetApiV4ProjectsIdErrorTrackingClientKeys(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdErrorTrackingClientKeys request
	PostApiV4ProjectsIdErrorTrackingClientKeys(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyId request
	DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyId(ctx context.Context, id string, keyId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdErrorTrackingSettings request
	GetApiV4ProjectsIdErrorTrackingSettings(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchApiV4ProjectsIdErrorTrackingSettingsWithBody request with any body
	PatchApiV4ProjectsIdErrorTrackingSettingsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchApiV4ProjectsIdErrorTrackingSettings(ctx context.Context, id string, body PatchApiV4ProjectsIdErrorTrackingSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdErrorTrackingSettingsWithBody request with any body
	PutApiV4ProjectsIdErrorTrackingSettingsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdErrorTrackingSettings(ctx context.Context, id string, body PutApiV4ProjectsIdErrorTrackingSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdEvents request
	GetApiV4ProjectsIdEvents(ctx context.Context, id string, params *GetApiV4ProjectsIdEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdExport request
	GetApiV4ProjectsIdExport(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdExportWithBody request with any body
	PostApiV4ProjectsIdExportWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdExport(ctx context.Context, id string, body PostApiV4ProjectsIdExportJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdExportDownload request
	GetApiV4ProjectsIdExportDownload(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdExportRelationsWithBody request with any body
	PostApiV4ProjectsIdExportRelationsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdExportRelations(ctx context.Context, id string, body PostApiV4ProjectsIdExportRelationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdExportRelationsDownload request
	GetApiV4ProjectsIdExportRelationsDownload(ctx context.Context, id string, params *GetApiV4ProjectsIdExportRelationsDownloadParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdExportRelationsStatus request
	GetApiV4ProjectsIdExportRelationsStatus(ctx context.Context, id string, params *GetApiV4ProjectsIdExportRelationsStatusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdFeatureFlags request
	GetApiV4ProjectsIdFeatureFlags(ctx context.Context, id string, params *GetApiV4ProjectsIdFeatureFlagsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdFeatureFlagsWithBody request with any body
	PostApiV4ProjectsIdFeatureFlagsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdFeatureFlags(ctx context.Context, id string, body PostApiV4ProjectsIdFeatureFlagsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagName request
	DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx context.Context, id string, featureFlagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdFeatureFlagsFeatureFlagName request
	GetApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx context.Context, id string, featureFlagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithBody request with any body
	PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithBody(ctx context.Context, id string, featureFlagName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdFeatureFlagsFeatureFlagName(ctx context.Context, id string, featureFlagName string, body PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdFeatureFlagsUserLists request
	GetApiV4ProjectsIdFeatureFlagsUserLists(ctx context.Context, id string, params *GetApiV4ProjectsIdFeatureFlagsUserListsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdFeatureFlagsUserListsWithBody request with any body
	PostApiV4ProjectsIdFeatureFlagsUserListsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdFeatureFlagsUserLists(ctx context.Context, id string, body PostApiV4ProjectsIdFeatureFlagsUserListsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdFeatureFlagsUserListsIid request
	DeleteApiV4ProjectsIdFeatureFlagsUserListsIid(ctx context.Context, id string, iid string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdFeatureFlagsUserListsIid request
	GetApiV4ProjectsIdFeatureFlagsUserListsIid(ctx context.Context, id string, iid string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdFeatureFlagsUserListsIidWithBody request with any body
	PutApiV4ProjectsIdFeatureFlagsUserListsIidWithBody(ctx context.Context, id string, iid string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdFeatureFlagsUserListsIid(ctx context.Context, id string, iid string, body PutApiV4ProjectsIdFeatureFlagsUserListsIidJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdFork request
	DeleteApiV4ProjectsIdFork(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdForkWithBody request with any body
	PostApiV4ProjectsIdForkWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdFork(ctx context.Context, id string, body PostApiV4ProjectsIdForkJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdForkForkedFromId request
	PostApiV4ProjectsIdForkForkedFromId(ctx context.Context, id string, forkedFromId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdForks request
	GetApiV4ProjectsIdForks(ctx context.Context, id string, params *GetApiV4ProjectsIdForksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdFreezePeriods request
	GetApiV4ProjectsIdFreezePeriods(ctx context.Context, id string, params *GetApiV4ProjectsIdFreezePeriodsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdFreezePeriodsWithBody request with any body
	PostApiV4ProjectsIdFreezePeriodsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdFreezePeriods(ctx context.Context, id string, body PostApiV4ProjectsIdFreezePeriodsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodId request
	DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx context.Context, id string, freezePeriodId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdFreezePeriodsFreezePeriodId request
	GetApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx context.Context, id string, freezePeriodId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithBody request with any body
	PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithBody(ctx context.Context, id string, freezePeriodId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdFreezePeriodsFreezePeriodId(ctx context.Context, id string, freezePeriodId int32, body PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdGroups request
	GetApiV4ProjectsIdGroups(ctx context.Context, id string, params *GetApiV4ProjectsIdGroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdHooks request
	GetApiV4ProjectsIdHooks(ctx context.Context, id string, params *GetApiV4ProjectsIdHooksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdHooksWithBody request with any body
	PostApiV4ProjectsIdHooksWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdHooks(ctx context.Context, id string, body PostApiV4ProjectsIdHooksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdHooksHookId request
	DeleteApiV4ProjectsIdHooksHookId(ctx context.Context, id string, hookId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdHooksHookId request
	GetApiV4ProjectsIdHooksHookId(ctx context.Context, id string, hookId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdHooksHookIdWithBody request with any body
	PutApiV4ProjectsIdHooksHookIdWithBody(ctx context.Context, id string, hookId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdHooksHookId(ctx context.Context, id string, hookId int32, body PutApiV4ProjectsIdHooksHookIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdHooksHookIdCustomHeadersKey request
	DeleteApiV4ProjectsIdHooksHookIdCustomHeadersKey(ctx context.Context, id int32, hookId int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyWithBody request with any body
	PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyWithBody(ctx context.Context, id int32, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdHooksHookIdCustomHeadersKey(ctx context.Context, id int32, hookId int32, key string, body PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdHooksHookIdEvents request
	GetApiV4ProjectsIdHooksHookIdEvents(ctx context.Context, id int32, hookId int32, params *GetApiV4ProjectsIdHooksHookIdEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdHooksHookIdEventsHookLogIdResend request
	PostApiV4ProjectsIdHooksHookIdEventsHookLogIdResend(ctx context.Context, id int32, hookId int32, hookLogId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdHooksHookIdTestTrigger request
	PostApiV4ProjectsIdHooksHookIdTestTrigger(ctx context.Context, id int32, hookId int32, trigger string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKey request
	DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKey(ctx context.Context, id int32, hookId int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyWithBody request with any body
	PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyWithBody(ctx context.Context, id int32, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdHooksHookIdUrlVariablesKey(ctx context.Context, id int32, hookId int32, key string, body PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdHousekeepingWithBody request with any body
	PostApiV4ProjectsIdHousekeepingWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdHousekeeping(ctx context.Context, id string, body PostApiV4ProjectsIdHousekeepingJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdImport request
	GetApiV4ProjectsIdImport(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdImportProjectMembersProjectId request
	PostApiV4ProjectsIdImportProjectMembersProjectId(ctx context.Context, id string, projectId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIntegrations request
	GetApiV4ProjectsIdIntegrations(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsAppleAppStoreWithBody request with any body
	PutApiV4ProjectsIdIntegrationsAppleAppStoreWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsAppleAppStore(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsAppleAppStoreJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsAsanaWithBody request with any body
	PutApiV4ProjectsIdIntegrationsAsanaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsAsana(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsAsanaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsAssemblaWithBody request with any body
	PutApiV4ProjectsIdIntegrationsAssemblaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsAssembla(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsAssemblaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsBambooWithBody request with any body
	PutApiV4ProjectsIdIntegrationsBambooWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsBamboo(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsBambooJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsBugzillaWithBody request with any body
	PutApiV4ProjectsIdIntegrationsBugzillaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsBugzilla(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsBugzillaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsBuildkiteWithBody request with any body
	PutApiV4ProjectsIdIntegrationsBuildkiteWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsBuildkite(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsBuildkiteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsCampfireWithBody request with any body
	PutApiV4ProjectsIdIntegrationsCampfireWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsCampfire(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsCampfireJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsClickupWithBody request with any body
	PutApiV4ProjectsIdIntegrationsClickupWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsClickup(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsClickupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsConfluenceWithBody request with any body
	PutApiV4ProjectsIdIntegrationsConfluenceWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsConfluence(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsConfluenceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsCustomIssueTrackerWithBody request with any body
	PutApiV4ProjectsIdIntegrationsCustomIssueTrackerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsCustomIssueTracker(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsCustomIssueTrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsDatadogWithBody request with any body
	PutApiV4ProjectsIdIntegrationsDatadogWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsDatadog(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsDatadogJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsDiffblueCoverWithBody request with any body
	PutApiV4ProjectsIdIntegrationsDiffblueCoverWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsDiffblueCover(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsDiffblueCoverJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsDiscordWithBody request with any body
	PutApiV4ProjectsIdIntegrationsDiscordWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsDiscord(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsDiscordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsDroneCiWithBody request with any body
	PutApiV4ProjectsIdIntegrationsDroneCiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsDroneCi(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsDroneCiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsEmailsOnPushWithBody request with any body
	PutApiV4ProjectsIdIntegrationsEmailsOnPushWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsEmailsOnPush(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsEmailsOnPushJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsEwmWithBody request with any body
	PutApiV4ProjectsIdIntegrationsEwmWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsEwm(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsEwmJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsExternalWikiWithBody request with any body
	PutApiV4ProjectsIdIntegrationsExternalWikiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsExternalWiki(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsExternalWikiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsGitGuardianWithBody request with any body
	PutApiV4ProjectsIdIntegrationsGitGuardianWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsGitGuardian(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGitGuardianJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsGithubWithBody request with any body
	PutApiV4ProjectsIdIntegrationsGithubWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsGithub(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationWithBody request with any body
	PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsGitlabSlackApplication(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithBody request with any body
	PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistry(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithBody request with any body
	PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederation(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsGooglePlayWithBody request with any body
	PutApiV4ProjectsIdIntegrationsGooglePlayWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsGooglePlay(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGooglePlayJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsHangoutsChatWithBody request with any body
	PutApiV4ProjectsIdIntegrationsHangoutsChatWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsHangoutsChat(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsHangoutsChatJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsHarborWithBody request with any body
	PutApiV4ProjectsIdIntegrationsHarborWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsHarbor(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsHarborJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsIrkerWithBody request with any body
	PutApiV4ProjectsIdIntegrationsIrkerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsIrker(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsIrkerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsJenkinsWithBody request with any body
	PutApiV4ProjectsIdIntegrationsJenkinsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsJenkins(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsJenkinsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsJiraWithBody request with any body
	PutApiV4ProjectsIdIntegrationsJiraWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsJira(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsJiraJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsJiraCloudAppWithBody request with any body
	PutApiV4ProjectsIdIntegrationsJiraCloudAppWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsJiraCloudApp(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsJiraCloudAppJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsMatrixWithBody request with any body
	PutApiV4ProjectsIdIntegrationsMatrixWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsMatrix(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMatrixJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsMattermostWithBody request with any body
	PutApiV4ProjectsIdIntegrationsMattermostWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsMattermost(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMattermostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsWithBody request with any body
	PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsMattermostSlashCommands(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerWithBody request with any body
	PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTrigger(ctx context.Context, id string, body PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsMicrosoftTeamsWithBody request with any body
	PutApiV4ProjectsIdIntegrationsMicrosoftTeamsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsMicrosoftTeams(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMicrosoftTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsMockCiWithBody request with any body
	PutApiV4ProjectsIdIntegrationsMockCiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsMockCi(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMockCiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsMockMonitoringWithBody request with any body
	PutApiV4ProjectsIdIntegrationsMockMonitoringWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsMockMonitoring(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMockMonitoringJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsPackagistWithBody request with any body
	PutApiV4ProjectsIdIntegrationsPackagistWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsPackagist(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPackagistJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsPhorgeWithBody request with any body
	PutApiV4ProjectsIdIntegrationsPhorgeWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsPhorge(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPhorgeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsPipelinesEmailWithBody request with any body
	PutApiV4ProjectsIdIntegrationsPipelinesEmailWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsPipelinesEmail(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPipelinesEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsPivotaltrackerWithBody request with any body
	PutApiV4ProjectsIdIntegrationsPivotaltrackerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsPivotaltracker(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPivotaltrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsPumbleWithBody request with any body
	PutApiV4ProjectsIdIntegrationsPumbleWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsPumble(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPumbleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsPushoverWithBody request with any body
	PutApiV4ProjectsIdIntegrationsPushoverWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsPushover(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPushoverJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsRedmineWithBody request with any body
	PutApiV4ProjectsIdIntegrationsRedmineWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsRedmine(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsRedmineJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsSlackWithBody request with any body
	PutApiV4ProjectsIdIntegrationsSlackWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsSlack(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsSlackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsSlackSlashCommandsWithBody request with any body
	PutApiV4ProjectsIdIntegrationsSlackSlashCommandsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsSlackSlashCommands(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsSlackSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerWithBody request with any body
	PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTrigger(ctx context.Context, id string, body PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsSquashTmWithBody request with any body
	PutApiV4ProjectsIdIntegrationsSquashTmWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsSquashTm(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsSquashTmJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsTeamcityWithBody request with any body
	PutApiV4ProjectsIdIntegrationsTeamcityWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsTeamcity(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsTeamcityJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsTelegramWithBody request with any body
	PutApiV4ProjectsIdIntegrationsTelegramWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsTelegram(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsTelegramJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsUnifyCircuitWithBody request with any body
	PutApiV4ProjectsIdIntegrationsUnifyCircuitWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsUnifyCircuit(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsUnifyCircuitJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsWebexTeamsWithBody request with any body
	PutApiV4ProjectsIdIntegrationsWebexTeamsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsWebexTeams(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsWebexTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsYoutrackWithBody request with any body
	PutApiV4ProjectsIdIntegrationsYoutrackWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsYoutrack(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsYoutrackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIntegrationsZentaoWithBody request with any body
	PutApiV4ProjectsIdIntegrationsZentaoWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdIntegrationsZentao(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsZentaoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdIntegrationsSlug request
	DeleteApiV4ProjectsIdIntegrationsSlug(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIntegrationsSlug request
	GetApiV4ProjectsIdIntegrationsSlug(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdInvitations request
	GetApiV4ProjectsIdInvitations(ctx context.Context, id string, params *GetApiV4ProjectsIdInvitationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdInvitationsWithBody request with any body
	PostApiV4ProjectsIdInvitationsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdInvitations(ctx context.Context, id string, body PostApiV4ProjectsIdInvitationsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdInvitationsEmail request
	DeleteApiV4ProjectsIdInvitationsEmail(ctx context.Context, id string, email string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdInvitationsEmailWithBody request with any body
	PutApiV4ProjectsIdInvitationsEmailWithBody(ctx context.Context, id string, email string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdInvitationsEmail(ctx context.Context, id string, email string, body PutApiV4ProjectsIdInvitationsEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdInvitedGroups request
	GetApiV4ProjectsIdInvitedGroups(ctx context.Context, id string, params *GetApiV4ProjectsIdInvitedGroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssues request
	PostApiV4ProjectsIdIssues(ctx context.Context, id string, params *PostApiV4ProjectsIdIssuesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEvents request
	GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEvents(ctx context.Context, id string, eventableId int32, params *GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventId request
	GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventId(ctx context.Context, id string, eventableId int32, eventId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdIssuesIssueIid request
	DeleteApiV4ProjectsIdIssuesIssueIid(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidAddSpentTime request
	PostApiV4ProjectsIdIssuesIssueIidAddSpentTime(ctx context.Context, id string, issueIid int, params *PostApiV4ProjectsIdIssuesIssueIidAddSpentTimeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidAwardEmoji request
	GetApiV4ProjectsIdIssuesIssueIidAwardEmoji(ctx context.Context, id string, issueIid int32, params *GetApiV4ProjectsIdIssuesIssueIidAwardEmojiParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidAwardEmojiWithBody request with any body
	PostApiV4ProjectsIdIssuesIssueIidAwardEmojiWithBody(ctx context.Context, id int32, issueIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdIssuesIssueIidAwardEmoji(ctx context.Context, id int32, issueIid int32, body PostApiV4ProjectsIdIssuesIssueIidAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardId request
	DeleteApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardId(ctx context.Context, id int32, issueIid int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardId request
	GetApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardId(ctx context.Context, id int32, issueIid int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidClone request
	PostApiV4ProjectsIdIssuesIssueIidClone(ctx context.Context, id string, issueIid int, params *PostApiV4ProjectsIdIssuesIssueIidCloneParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidClosedBy request
	GetApiV4ProjectsIdIssuesIssueIidClosedBy(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidLinks request
	GetApiV4ProjectsIdIssuesIssueIidLinks(ctx context.Context, id string, issueIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidLinksWithBody request with any body
	PostApiV4ProjectsIdIssuesIssueIidLinksWithBody(ctx context.Context, id string, issueIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdIssuesIssueIidLinks(ctx context.Context, id string, issueIid int32, body PostApiV4ProjectsIdIssuesIssueIidLinksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId request
	DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId(ctx context.Context, id string, issueIid int32, issueLinkId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId request
	GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkId(ctx context.Context, id string, issueIid int32, issueLinkId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidMetricImages request
	GetApiV4ProjectsIdIssuesIssueIidMetricImages(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidMetricImagesImageId request
	DeleteApiV4ProjectsIdIssuesIssueIidMetricImagesImageId(ctx context.Context, id string, issueIid int, imageId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidMoveWithBody request with any body
	PostApiV4ProjectsIdIssuesIssueIidMoveWithBody(ctx context.Context, id string, issueIid int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdIssuesIssueIidMoveWithFormdataBody(ctx context.Context, id string, issueIid int, body PostApiV4ProjectsIdIssuesIssueIidMoveFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidNotes request
	PostApiV4ProjectsIdIssuesIssueIidNotes(ctx context.Context, id string, issueIid int, params *PostApiV4ProjectsIdIssuesIssueIidNotesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji request
	GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(ctx context.Context, id int32, issueIid int32, noteId int32, params *GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiWithBody request with any body
	PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiWithBody(ctx context.Context, id int32, issueIid int32, noteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmoji(ctx context.Context, id int32, issueIid int32, noteId int32, body PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId request
	DeleteApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(ctx context.Context, id int32, issueIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId request
	GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardId(ctx context.Context, id int32, issueIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidParticipants request
	GetApiV4ProjectsIdIssuesIssueIidParticipants(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidRelatedMergeRequests request
	GetApiV4ProjectsIdIssuesIssueIidRelatedMergeRequests(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdIssuesIssueIidReorder request
	PutApiV4ProjectsIdIssuesIssueIidReorder(ctx context.Context, id string, issueIid int, params *PutApiV4ProjectsIdIssuesIssueIidReorderParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidResetSpentTime request
	PostApiV4ProjectsIdIssuesIssueIidResetSpentTime(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidResetTimeEstimate request
	PostApiV4ProjectsIdIssuesIssueIidResetTimeEstimate(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidSubscribe request
	PostApiV4ProjectsIdIssuesIssueIidSubscribe(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidTimeEstimate request
	PostApiV4ProjectsIdIssuesIssueIidTimeEstimate(ctx context.Context, id string, issueIid int, params *PostApiV4ProjectsIdIssuesIssueIidTimeEstimateParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidTimeStats request
	GetApiV4ProjectsIdIssuesIssueIidTimeStats(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidTodo request
	PostApiV4ProjectsIdIssuesIssueIidTodo(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdIssuesIssueIidUnsubscribe request
	PostApiV4ProjectsIdIssuesIssueIidUnsubscribe(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdIssuesIssueIidUserAgentDetail request
	GetApiV4ProjectsIdIssuesIssueIidUserAgentDetail(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobTokenScope request
	GetApiV4ProjectsIdJobTokenScope(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchApiV4ProjectsIdJobTokenScopeWithBody request with any body
	PatchApiV4ProjectsIdJobTokenScopeWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchApiV4ProjectsIdJobTokenScope(ctx context.Context, id int32, body PatchApiV4ProjectsIdJobTokenScopeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobTokenScopeAllowlist request
	GetApiV4ProjectsIdJobTokenScopeAllowlist(ctx context.Context, id int32, params *GetApiV4ProjectsIdJobTokenScopeAllowlistParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdJobTokenScopeAllowlistWithBody request with any body
	PostApiV4ProjectsIdJobTokenScopeAllowlistWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdJobTokenScopeAllowlist(ctx context.Context, id int32, body PostApiV4ProjectsIdJobTokenScopeAllowlistJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectId request
	DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectId(ctx context.Context, id int32, targetProjectId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobTokenScopeGroupsAllowlist request
	GetApiV4ProjectsIdJobTokenScopeGroupsAllowlist(ctx context.Context, id int32, params *GetApiV4ProjectsIdJobTokenScopeGroupsAllowlistParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistWithBody request with any body
	PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdJobTokenScopeGroupsAllowlist(ctx context.Context, id int32, body PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId request
	DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupId(ctx context.Context, id int32, targetGroupId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobs request
	GetApiV4ProjectsIdJobs(ctx context.Context, id string, params *GetApiV4ProjectsIdJobsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobsArtifactsRefNameDownload request
	GetApiV4ProjectsIdJobsArtifactsRefNameDownload(ctx context.Context, id string, refName string, params *GetApiV4ProjectsIdJobsArtifactsRefNameDownloadParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPath request
	GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPath(ctx context.Context, id string, refName string, params *GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPathParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobsJobId request
	GetApiV4ProjectsIdJobsJobId(ctx context.Context, id int32, jobId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdJobsJobIdArtifacts request
	DeleteApiV4ProjectsIdJobsJobIdArtifacts(ctx context.Context, id string, jobId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobsJobIdArtifacts request
	GetApiV4ProjectsIdJobsJobIdArtifacts(ctx context.Context, id string, jobId int32, params *GetApiV4ProjectsIdJobsJobIdArtifactsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobsJobIdArtifactsartifactPath request
	GetApiV4ProjectsIdJobsJobIdArtifactsartifactPath(ctx context.Context, id string, jobId int32, params *GetApiV4ProjectsIdJobsJobIdArtifactsartifactPathParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdJobsJobIdArtifactsKeep request
	PostApiV4ProjectsIdJobsJobIdArtifactsKeep(ctx context.Context, id string, jobId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdJobsJobIdCancelWithBody request with any body
	PostApiV4ProjectsIdJobsJobIdCancelWithBody(ctx context.Context, id int32, jobId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdJobsJobIdCancel(ctx context.Context, id int32, jobId int32, body PostApiV4ProjectsIdJobsJobIdCancelJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdJobsJobIdErase request
	PostApiV4ProjectsIdJobsJobIdErase(ctx context.Context, id int32, jobId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdJobsJobIdPlayWithBody request with any body
	PostApiV4ProjectsIdJobsJobIdPlayWithBody(ctx context.Context, id int32, jobId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdJobsJobIdPlay(ctx context.Context, id int32, jobId int32, body PostApiV4ProjectsIdJobsJobIdPlayJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdJobsJobIdRetry request
	PostApiV4ProjectsIdJobsJobIdRetry(ctx context.Context, id int32, jobId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdJobsJobIdTrace request
	GetApiV4ProjectsIdJobsJobIdTrace(ctx context.Context, id int32, jobId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdLanguages request
	GetApiV4ProjectsIdLanguages(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMembers request
	GetApiV4ProjectsIdMembers(ctx context.Context, id string, params *GetApiV4ProjectsIdMembersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMembersWithBody request with any body
	PostApiV4ProjectsIdMembersWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMembers(ctx context.Context, id string, body PostApiV4ProjectsIdMembersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMembersAll request
	GetApiV4ProjectsIdMembersAll(ctx context.Context, id string, params *GetApiV4ProjectsIdMembersAllParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMembersAllUserId request
	GetApiV4ProjectsIdMembersAllUserId(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdMembersUserId request
	DeleteApiV4ProjectsIdMembersUserId(ctx context.Context, id string, userId int32, params *DeleteApiV4ProjectsIdMembersUserIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMembersUserId request
	GetApiV4ProjectsIdMembersUserId(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdMembersUserIdWithBody request with any body
	PutApiV4ProjectsIdMembersUserIdWithBody(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdMembersUserId(ctx context.Context, id string, userId int32, body PutApiV4ProjectsIdMembersUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequests request
	GetApiV4ProjectsIdMergeRequests(ctx context.Context, id string, params *GetApiV4ProjectsIdMergeRequestsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequests(ctx context.Context, id string, body PostApiV4ProjectsIdMergeRequestsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEvents request
	GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEvents(ctx context.Context, id string, eventableId int32, params *GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId request
	GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventId(ctx context.Context, id string, eventableId int32, eventId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIid request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIid request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidWithBody request with any body
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdMergeRequestsMergeRequestIid(ctx context.Context, id string, mergeRequestIid int32, body PutApiV4ProjectsIdMergeRequestsMergeRequestIidJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTime(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalState request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalState(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals(ctx context.Context, id int32, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovals(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveWithBody(ctx context.Context, id int32, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprove(ctx context.Context, id int32, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmoji request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmoji(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiWithBody(ctx context.Context, id int32, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmoji(ctx context.Context, id int32, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(ctx context.Context, id int32, mergeRequestIid int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardId(ctx context.Context, id int32, mergeRequestIid int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceeds(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidChanges request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidChanges(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssues request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssues(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommits request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommits(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommitsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx context.Context, id string, mergeRequestIid int32, params *DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommits(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffs request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffs(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotes(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublish(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithBody request with any body
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithBody(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteId(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, body PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish request
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublish(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeWithBody request with any body
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdMergeRequestsMergeRequestIidMerge(ctx context.Context, id string, mergeRequestIid int32, body PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRef request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRef(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiWithBody(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmoji(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardId(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipants request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipants(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelines(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidRawDiffs request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidRawDiffs(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseWithBody request with any body
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebase(ctx context.Context, id string, mergeRequestIid int32, body PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssues request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssues(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssuesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovals request
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovals(ctx context.Context, id int32, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTime request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTime(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimate(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewers request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewers(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateWithBody request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateWithBody(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimate(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStats request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStats(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapprove request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapprove(ctx context.Context, id int32, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersions request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersions(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionId request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionId(ctx context.Context, id string, mergeRequestIid int32, versionId int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackages request
	GetApiV4ProjectsIdPackages(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesCargoConfigJson request
	GetApiV4ProjectsIdPackagesCargoConfigJson(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesComposerWithBody request with any body
	PostApiV4ProjectsIdPackagesComposerWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPackagesComposer(ctx context.Context, id string, body PostApiV4ProjectsIdPackagesComposerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesComposerArchivespackageName request
	GetApiV4ProjectsIdPackagesComposerArchivespackageName(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesComposerArchivespackageNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansSearch request
	GetApiV4ProjectsIdPackagesConanV1ConansSearch(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesConanV1ConansSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel request
	DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannel(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigest(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrls(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReference(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigest(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrls(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls request
	PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrls(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls request
	PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrls(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName request
	GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithBody request with any body
	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithBody(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileName(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, body PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize request
	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorize(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName request
	GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithBody request with any body
	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithBody(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileName(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, body PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize request
	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorize(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1Ping request
	GetApiV4ProjectsIdPackagesConanV1Ping(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1UsersAuthenticate request
	GetApiV4ProjectsIdPackagesConanV1UsersAuthenticate(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentials request
	GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentials(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansSearch request
	GetApiV4ProjectsIdPackagesConanV2ConansSearch(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesConanV2ConansSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatest request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatest(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisions request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisions(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevision request
	DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevision(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFiles request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFiles(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameWithBody request with any body
	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameWithBody(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileName(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, body PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorize request
	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorize(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatest(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisions(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevision request
	DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevision(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFiles(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameWithBody request with any body
	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameWithBody(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileName(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, body PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorize request
	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorize(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearch(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearch(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2UsersAuthenticate request
	GetApiV4ProjectsIdPackagesConanV2UsersAuthenticate(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesConanV2UsersCheckCredentials request
	GetApiV4ProjectsIdPackagesConanV2UsersCheckCredentials(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionInrelease request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionInrelease(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionInreleaseParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionRelease request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionRelease(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpg request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpg(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpgParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackages(ctx context.Context, id string, component string, architecture string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256 request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256(ctx context.Context, id string, component string, architecture string, fileSha256 int32, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackages(ctx context.Context, id string, component string, architecture string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256 request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256(ctx context.Context, id string, component string, architecture string, fileSha256 int32, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSources request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSources(ctx context.Context, id string, component string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSourcesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256 request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256(ctx context.Context, id string, component string, fileSha256 int32, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName request
	GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileName(ctx context.Context, id string, distribution string, letter string, packageName string, packageVersion string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesDebianFileNameWithBody request with any body
	PutApiV4ProjectsIdPackagesDebianFileNameWithBody(ctx context.Context, id string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesDebianFileName(ctx context.Context, id string, fileName string, body PutApiV4ProjectsIdPackagesDebianFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeWithBody request with any body
	PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeWithBody(ctx context.Context, id string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesDebianFileNameAuthorize(ctx context.Context, id string, fileName string, body PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileName request
	GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileName(ctx context.Context, id string, packageName string, fileName string, params *GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameWithBody request with any body
	PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameWithBody(ctx context.Context, id string, packageName string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileName(ctx context.Context, id string, packageName string, fileName string, body PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeWithBody request with any body
	PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeWithBody(ctx context.Context, id string, packageName string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorize(ctx context.Context, id string, packageName string, fileName string, body PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesGomoduleNameVList request
	GetApiV4ProjectsIdPackagesGomoduleNameVList(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesGomoduleNameVListParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionInfo request
	GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionInfo(ctx context.Context, id string, moduleVersion string, params *GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionInfoParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionMod request
	GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionMod(ctx context.Context, id string, moduleVersion string, params *GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionModParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionZip request
	GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionZip(ctx context.Context, id string, moduleVersion string, params *GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionZipParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesHelmApiChannelChartsWithBody request with any body
	PostApiV4ProjectsIdPackagesHelmApiChannelChartsWithBody(ctx context.Context, id int32, channel string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPackagesHelmApiChannelCharts(ctx context.Context, id int32, channel string, body PostApiV4ProjectsIdPackagesHelmApiChannelChartsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorize request
	PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorize(ctx context.Context, id int32, channel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgz request
	GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgz(ctx context.Context, id int32, channel string, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesHelmChannelIndexYaml request
	GetApiV4ProjectsIdPackagesHelmChannelIndexYaml(ctx context.Context, id int32, channel string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesMavenpathFileName request
	GetApiV4ProjectsIdPackagesMavenpathFileName(ctx context.Context, id string, fileName string, params *GetApiV4ProjectsIdPackagesMavenpathFileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesMavenpathFileNameWithBody request with any body
	PutApiV4ProjectsIdPackagesMavenpathFileNameWithBody(ctx context.Context, id string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesMavenpathFileName(ctx context.Context, id string, fileName string, body PutApiV4ProjectsIdPackagesMavenpathFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeWithBody request with any body
	PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeWithBody(ctx context.Context, id string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorize(ctx context.Context, id string, fileName string, body PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileName request
	GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileName(ctx context.Context, id string, modelVersionId string, fileName string, params *GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameWithBody request with any body
	PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameWithBody(ctx context.Context, id string, modelVersionId string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileName(ctx context.Context, id string, modelVersionId string, fileName string, body PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeWithBody request with any body
	PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeWithBody(ctx context.Context, id string, modelVersionId string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorize(ctx context.Context, id string, modelVersionId string, fileName string, body PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNpmpackageName request
	GetApiV4ProjectsIdPackagesNpmpackageName(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNpmpackageNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNpmpackageNamefileName request
	GetApiV4ProjectsIdPackagesNpmpackageNamefileName(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNpmpackageNamefileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk request
	PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulk(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuick request
	PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuick(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTags request
	GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTags(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag request
	DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag(ctx context.Context, id string, tag string, params *DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagWithBody request with any body
	PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagWithBody(ctx context.Context, id string, tag string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTag(ctx context.Context, id string, tag string, body PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesNpmPackageNameWithBody request with any body
	PutApiV4ProjectsIdPackagesNpmPackageNameWithBody(ctx context.Context, id string, packageName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesNpmPackageName(ctx context.Context, id string, packageName string, body PutApiV4ProjectsIdPackagesNpmPackageNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesNugetWithBody request with any body
	PutApiV4ProjectsIdPackagesNugetWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesNuget(ctx context.Context, id string, body PutApiV4ProjectsIdPackagesNugetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersion request
	DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersion(ctx context.Context, id string, params *DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesNugetAuthorize request
	PutApiV4ProjectsIdPackagesNugetAuthorize(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilename request
	GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilename(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilenameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndex request
	GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndex(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndexParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNugetIndex request
	GetApiV4ProjectsIdPackagesNugetIndex(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNugetQuery request
	GetApiV4ProjectsIdPackagesNugetQuery(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNugetQueryParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName request
	GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileName(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesNugetSymbolpackageWithBody request with any body
	PutApiV4ProjectsIdPackagesNugetSymbolpackageWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesNugetSymbolpackage(ctx context.Context, id string, body PutApiV4ProjectsIdPackagesNugetSymbolpackageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorize request
	PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorize(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNugetV2 request
	GetApiV4ProjectsIdPackagesNugetV2(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesNugetV2WithBody request with any body
	PutApiV4ProjectsIdPackagesNugetV2WithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesNugetV2(ctx context.Context, id string, body PutApiV4ProjectsIdPackagesNugetV2JSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesNugetV2Metadata request
	GetApiV4ProjectsIdPackagesNugetV2Metadata(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesNugetV2Authorize request
	PutApiV4ProjectsIdPackagesNugetV2Authorize(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesProtectionRules request
	GetApiV4ProjectsIdPackagesProtectionRules(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesProtectionRulesWithBody request with any body
	PostApiV4ProjectsIdPackagesProtectionRulesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPackagesProtectionRules(ctx context.Context, id string, body PostApiV4ProjectsIdPackagesProtectionRulesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleId request
	DeleteApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleId(ctx context.Context, id string, packageProtectionRuleId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdWithBody request with any body
	PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdWithBody(ctx context.Context, id string, packageProtectionRuleId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleId(ctx context.Context, id string, packageProtectionRuleId int32, body PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesPypiWithBody request with any body
	PostApiV4ProjectsIdPackagesPypiWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPackagesPypi(ctx context.Context, id string, body PostApiV4ProjectsIdPackagesPypiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesPypiAuthorize request
	PostApiV4ProjectsIdPackagesPypiAuthorize(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifier request
	GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifier(ctx context.Context, id string, sha256 string, params *GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifierParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesPypiSimple request
	GetApiV4ProjectsIdPackagesPypiSimple(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesPypiSimplepackageName request
	GetApiV4ProjectsIdPackagesPypiSimplepackageName(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesPypiSimplepackageNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesRpm request
	PostApiV4ProjectsIdPackagesRpm(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesRpmpackageFileIdfileName request
	GetApiV4ProjectsIdPackagesRpmpackageFileIdfileName(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesRpmpackageFileIdfileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesRpmAuthorize request
	PostApiV4ProjectsIdPackagesRpmAuthorize(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesRpmRepodatafileName request
	GetApiV4ProjectsIdPackagesRpmRepodatafileName(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesRpmRepodatafileNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesRubygemsApiV1Dependencies request
	GetApiV4ProjectsIdPackagesRubygemsApiV1Dependencies(ctx context.Context, id int32, params *GetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesRubygemsApiV1GemsWithBody request with any body
	PostApiV4ProjectsIdPackagesRubygemsApiV1GemsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPackagesRubygemsApiV1Gems(ctx context.Context, id int32, body PostApiV4ProjectsIdPackagesRubygemsApiV1GemsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorize request
	PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorize(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesRubygemsGemsFileName request
	GetApiV4ProjectsIdPackagesRubygemsGemsFileName(ctx context.Context, id int32, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileName request
	GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileName(ctx context.Context, id int32, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesRubygemsFileName request
	GetApiV4ProjectsIdPackagesRubygemsFileName(ctx context.Context, id int32, fileName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystem request
	GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystem(ctx context.Context, id string, moduleName string, moduleSystem string, params *GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersion request
	GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersion(ctx context.Context, id string, moduleName string, moduleSystem string, params *GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileWithBody request with any body
	PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileWithBody(ctx context.Context, id string, moduleName string, moduleSystem string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeWithBody request with any body
	PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeWithBody(ctx context.Context, id string, moduleName string, moduleSystem string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorize(ctx context.Context, id string, moduleName string, moduleSystem string, body PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPackagesPackageId request
	DeleteApiV4ProjectsIdPackagesPackageId(ctx context.Context, id string, packageId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesPackageId request
	GetApiV4ProjectsIdPackagesPackageId(ctx context.Context, id string, packageId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesPackageIdPackageFiles request
	GetApiV4ProjectsIdPackagesPackageIdPackageFiles(ctx context.Context, id string, packageId int32, params *GetApiV4ProjectsIdPackagesPackageIdPackageFilesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileId request
	DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileId(ctx context.Context, id string, packageId int32, packageFileId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPackagesPackageIdPipelines request
	GetApiV4ProjectsIdPackagesPackageIdPipelines(ctx context.Context, id string, packageId int32, params *GetApiV4ProjectsIdPackagesPackageIdPipelinesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPages request
	DeleteApiV4ProjectsIdPages(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPages request
	GetApiV4ProjectsIdPages(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchApiV4ProjectsIdPagesWithBody request with any body
	PatchApiV4ProjectsIdPagesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchApiV4ProjectsIdPages(ctx context.Context, id string, body PatchApiV4ProjectsIdPagesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPagesDomains request
	GetApiV4ProjectsIdPagesDomains(ctx context.Context, id string, params *GetApiV4ProjectsIdPagesDomainsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPagesDomainsWithBody request with any body
	PostApiV4ProjectsIdPagesDomainsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPagesDomains(ctx context.Context, id string, body PostApiV4ProjectsIdPagesDomainsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPagesDomainsDomain request
	DeleteApiV4ProjectsIdPagesDomainsDomain(ctx context.Context, id string, domain string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPagesDomainsDomain request
	GetApiV4ProjectsIdPagesDomainsDomain(ctx context.Context, id string, domain string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPagesDomainsDomainWithBody request with any body
	PutApiV4ProjectsIdPagesDomainsDomainWithBody(ctx context.Context, id string, domain string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPagesDomainsDomain(ctx context.Context, id string, domain string, body PutApiV4ProjectsIdPagesDomainsDomainJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPagesDomainsDomainVerify request
	PutApiV4ProjectsIdPagesDomainsDomainVerify(ctx context.Context, id string, domain string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPagesAccess request
	GetApiV4ProjectsIdPagesAccess(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPipelineWithBody request with any body
	PostApiV4ProjectsIdPipelineWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPipeline(ctx context.Context, id string, body PostApiV4ProjectsIdPipelineJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelineSchedules request
	GetApiV4ProjectsIdPipelineSchedules(ctx context.Context, id string, params *GetApiV4ProjectsIdPipelineSchedulesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPipelineSchedulesWithBody request with any body
	PostApiV4ProjectsIdPipelineSchedulesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPipelineSchedules(ctx context.Context, id string, body PostApiV4ProjectsIdPipelineSchedulesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleId request
	DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleId request
	GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithBody request with any body
	PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithBody(ctx context.Context, id string, pipelineScheduleId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleId(ctx context.Context, id string, pipelineScheduleId int32, body PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelines request
	GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelines(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlay request
	PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlay(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership request
	PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnership(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesWithBody request with any body
	PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesWithBody(ctx context.Context, id string, pipelineScheduleId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariables(ctx context.Context, id string, pipelineScheduleId int32, body PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey request
	DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(ctx context.Context, id string, pipelineScheduleId int32, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyWithBody request with any body
	PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyWithBody(ctx context.Context, id string, pipelineScheduleId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKey(ctx context.Context, id string, pipelineScheduleId int32, key string, body PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelines request
	GetApiV4ProjectsIdPipelines(ctx context.Context, id string, params *GetApiV4ProjectsIdPipelinesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelinesLatest request
	GetApiV4ProjectsIdPipelinesLatest(ctx context.Context, id string, params *GetApiV4ProjectsIdPipelinesLatestParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdPipelinesPipelineId request
	DeleteApiV4ProjectsIdPipelinesPipelineId(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelinesPipelineId request
	GetApiV4ProjectsIdPipelinesPipelineId(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdBridges request
	GetApiV4ProjectsIdPipelinesPipelineIdBridges(ctx context.Context, id string, pipelineId int32, params *GetApiV4ProjectsIdPipelinesPipelineIdBridgesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPipelinesPipelineIdCancel request
	PostApiV4ProjectsIdPipelinesPipelineIdCancel(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdJobs request
	GetApiV4ProjectsIdPipelinesPipelineIdJobs(ctx context.Context, id string, pipelineId int32, params *GetApiV4ProjectsIdPipelinesPipelineIdJobsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdPipelinesPipelineIdMetadataWithBody request with any body
	PutApiV4ProjectsIdPipelinesPipelineIdMetadataWithBody(ctx context.Context, id string, pipelineId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdPipelinesPipelineIdMetadata(ctx context.Context, id string, pipelineId int32, body PutApiV4ProjectsIdPipelinesPipelineIdMetadataJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdPipelinesPipelineIdRetry request
	PostApiV4ProjectsIdPipelinesPipelineIdRetry(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdTestReport request
	GetApiV4ProjectsIdPipelinesPipelineIdTestReport(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummary request
	GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummary(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdVariables request
	GetApiV4ProjectsIdPipelinesPipelineIdVariables(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdProtectedBranches request
	GetApiV4ProjectsIdProtectedBranches(ctx context.Context, id string, params *GetApiV4ProjectsIdProtectedBranchesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdProtectedBranchesWithBody request with any body
	PostApiV4ProjectsIdProtectedBranchesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdProtectedBranches(ctx context.Context, id string, body PostApiV4ProjectsIdProtectedBranchesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdProtectedBranchesName request
	DeleteApiV4ProjectsIdProtectedBranchesName(ctx context.Context, id string, name string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdProtectedBranchesName request
	GetApiV4ProjectsIdProtectedBranchesName(ctx context.Context, id string, name string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchApiV4ProjectsIdProtectedBranchesNameWithBody request with any body
	PatchApiV4ProjectsIdProtectedBranchesNameWithBody(ctx context.Context, id string, name string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchApiV4ProjectsIdProtectedBranchesName(ctx context.Context, id string, name string, body PatchApiV4ProjectsIdProtectedBranchesNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdProtectedTags request
	GetApiV4ProjectsIdProtectedTags(ctx context.Context, id string, params *GetApiV4ProjectsIdProtectedTagsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdProtectedTagsWithBody request with any body
	PostApiV4ProjectsIdProtectedTagsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdProtectedTags(ctx context.Context, id string, body PostApiV4ProjectsIdProtectedTagsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdProtectedTagsName request
	DeleteApiV4ProjectsIdProtectedTagsName(ctx context.Context, id string, name string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdProtectedTagsName request
	GetApiV4ProjectsIdProtectedTagsName(ctx context.Context, id string, name string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRegistryProtectionRepositoryRules request
	GetApiV4ProjectsIdRegistryProtectionRepositoryRules(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRegistryProtectionRepositoryRulesWithBody request with any body
	PostApiV4ProjectsIdRegistryProtectionRepositoryRulesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRegistryProtectionRepositoryRules(ctx context.Context, id string, body PostApiV4ProjectsIdRegistryProtectionRepositoryRulesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleId request
	DeleteApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleId(ctx context.Context, id string, protectionRuleId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdWithBody request with any body
	PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdWithBody(ctx context.Context, id string, protectionRuleId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleId(ctx context.Context, id string, protectionRuleId int32, body PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRegistryRepositories request
	GetApiV4ProjectsIdRegistryRepositories(ctx context.Context, id string, params *GetApiV4ProjectsIdRegistryRepositoriesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryId request
	DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryId(ctx context.Context, id string, repositoryId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags request
	DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags(ctx context.Context, id string, repositoryId int32, params *DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags request
	GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTags(ctx context.Context, id string, repositoryId int32, params *GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName request
	DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(ctx context.Context, id string, repositoryId int32, tagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName request
	GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagName(ctx context.Context, id string, repositoryId int32, tagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRelationImports request
	GetApiV4ProjectsIdRelationImports(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdReleases request
	GetApiV4ProjectsIdReleases(ctx context.Context, id string, params *GetApiV4ProjectsIdReleasesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdReleasesWithBody request with any body
	PostApiV4ProjectsIdReleasesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdReleases(ctx context.Context, id string, body PostApiV4ProjectsIdReleasesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdReleasesPermalinkLatestsuffixPath request
	GetApiV4ProjectsIdReleasesPermalinkLatestsuffixPath(ctx context.Context, id string, params *GetApiV4ProjectsIdReleasesPermalinkLatestsuffixPathParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdReleasesTagName request
	DeleteApiV4ProjectsIdReleasesTagName(ctx context.Context, id string, tagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdReleasesTagName request
	GetApiV4ProjectsIdReleasesTagName(ctx context.Context, id string, tagName string, params *GetApiV4ProjectsIdReleasesTagNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdReleasesTagNameWithBody request with any body
	PutApiV4ProjectsIdReleasesTagNameWithBody(ctx context.Context, id string, tagName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdReleasesTagName(ctx context.Context, id string, tagName string, body PutApiV4ProjectsIdReleasesTagNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdReleasesTagNameAssetsLinks request
	GetApiV4ProjectsIdReleasesTagNameAssetsLinks(ctx context.Context, id string, tagName string, params *GetApiV4ProjectsIdReleasesTagNameAssetsLinksParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdReleasesTagNameAssetsLinksWithBody request with any body
	PostApiV4ProjectsIdReleasesTagNameAssetsLinksWithBody(ctx context.Context, id string, tagName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdReleasesTagNameAssetsLinks(ctx context.Context, id string, tagName string, body PostApiV4ProjectsIdReleasesTagNameAssetsLinksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId request
	DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx context.Context, id string, tagName string, linkId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId request
	GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx context.Context, id string, tagName string, linkId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithBody request with any body
	PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithBody(ctx context.Context, id string, tagName string, linkId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkId(ctx context.Context, id string, tagName string, linkId int32, body PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPath request
	GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPath(ctx context.Context, id string, tagName string, params *GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPathParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdReleasesTagNameEvidence request
	PostApiV4ProjectsIdReleasesTagNameEvidence(ctx context.Context, id int32, tagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRemoteMirrors request
	GetApiV4ProjectsIdRemoteMirrors(ctx context.Context, id string, params *GetApiV4ProjectsIdRemoteMirrorsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRemoteMirrorsWithBody request with any body
	PostApiV4ProjectsIdRemoteMirrorsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRemoteMirrors(ctx context.Context, id string, body PostApiV4ProjectsIdRemoteMirrorsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRemoteMirrorsMirrorId request
	DeleteApiV4ProjectsIdRemoteMirrorsMirrorId(ctx context.Context, id string, mirrorId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRemoteMirrorsMirrorId request
	GetApiV4ProjectsIdRemoteMirrorsMirrorId(ctx context.Context, id string, mirrorId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdRemoteMirrorsMirrorIdWithBody request with any body
	PutApiV4ProjectsIdRemoteMirrorsMirrorIdWithBody(ctx context.Context, id string, mirrorId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdRemoteMirrorsMirrorId(ctx context.Context, id string, mirrorId string, body PutApiV4ProjectsIdRemoteMirrorsMirrorIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRemoteMirrorsMirrorIdPublicKey request
	GetApiV4ProjectsIdRemoteMirrorsMirrorIdPublicKey(ctx context.Context, id string, mirrorId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRemoteMirrorsMirrorIdSync request
	PostApiV4ProjectsIdRemoteMirrorsMirrorIdSync(ctx context.Context, id string, mirrorId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryArchive request
	GetApiV4ProjectsIdRepositoryArchive(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryArchiveParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryBlobsSha request
	GetApiV4ProjectsIdRepositoryBlobsSha(ctx context.Context, id string, sha string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryBlobsShaRaw request
	GetApiV4ProjectsIdRepositoryBlobsShaRaw(ctx context.Context, id string, sha string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryBranches request
	GetApiV4ProjectsIdRepositoryBranches(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryBranchesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositoryBranchesWithBody request with any body
	PostApiV4ProjectsIdRepositoryBranchesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRepositoryBranches(ctx context.Context, id string, body PostApiV4ProjectsIdRepositoryBranchesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRepositoryBranchesBranch request
	DeleteApiV4ProjectsIdRepositoryBranchesBranch(ctx context.Context, id string, branch string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryBranchesBranch request
	GetApiV4ProjectsIdRepositoryBranchesBranch(ctx context.Context, id string, branch int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// HeadApiV4ProjectsIdRepositoryBranchesBranch request
	HeadApiV4ProjectsIdRepositoryBranchesBranch(ctx context.Context, id string, branch string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdRepositoryBranchesBranchProtectWithBody request with any body
	PutApiV4ProjectsIdRepositoryBranchesBranchProtectWithBody(ctx context.Context, id string, branch string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdRepositoryBranchesBranchProtect(ctx context.Context, id string, branch string, body PutApiV4ProjectsIdRepositoryBranchesBranchProtectJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdRepositoryBranchesBranchUnprotect request
	PutApiV4ProjectsIdRepositoryBranchesBranchUnprotect(ctx context.Context, id string, branch string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryChangelog request
	GetApiV4ProjectsIdRepositoryChangelog(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryChangelogParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositoryChangelogWithBody request with any body
	PostApiV4ProjectsIdRepositoryChangelogWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRepositoryChangelog(ctx context.Context, id string, body PostApiV4ProjectsIdRepositoryChangelogJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommits request
	GetApiV4ProjectsIdRepositoryCommits(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryCommitsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositoryCommitsWithBody request with any body
	PostApiV4ProjectsIdRepositoryCommitsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRepositoryCommits(ctx context.Context, id string, body PostApiV4ProjectsIdRepositoryCommitsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommitsSha request
	GetApiV4ProjectsIdRepositoryCommitsSha(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositoryCommitsShaCherryPickWithBody request with any body
	PostApiV4ProjectsIdRepositoryCommitsShaCherryPickWithBody(ctx context.Context, id string, sha string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRepositoryCommitsShaCherryPick(ctx context.Context, id string, sha string, body PostApiV4ProjectsIdRepositoryCommitsShaCherryPickJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaComments request
	GetApiV4ProjectsIdRepositoryCommitsShaComments(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaCommentsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositoryCommitsShaCommentsWithBody request with any body
	PostApiV4ProjectsIdRepositoryCommitsShaCommentsWithBody(ctx context.Context, id string, sha string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRepositoryCommitsShaComments(ctx context.Context, id string, sha string, body PostApiV4ProjectsIdRepositoryCommitsShaCommentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaDiff request
	GetApiV4ProjectsIdRepositoryCommitsShaDiff(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaDiffParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaMergeRequests request
	GetApiV4ProjectsIdRepositoryCommitsShaMergeRequests(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaRefs request
	GetApiV4ProjectsIdRepositoryCommitsShaRefs(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaRefsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositoryCommitsShaRevertWithBody request with any body
	PostApiV4ProjectsIdRepositoryCommitsShaRevertWithBody(ctx context.Context, id string, sha string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRepositoryCommitsShaRevert(ctx context.Context, id string, sha string, body PostApiV4ProjectsIdRepositoryCommitsShaRevertJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaSequence request
	GetApiV4ProjectsIdRepositoryCommitsShaSequence(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaSequenceParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaSignature request
	GetApiV4ProjectsIdRepositoryCommitsShaSignature(ctx context.Context, id string, sha string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaStatuses request
	GetApiV4ProjectsIdRepositoryCommitsShaStatuses(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaStatusesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryCompare request
	GetApiV4ProjectsIdRepositoryCompare(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryCompareParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryContributors request
	GetApiV4ProjectsIdRepositoryContributors(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryContributorsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRepositoryFilesFilePath request
	DeleteApiV4ProjectsIdRepositoryFilesFilePath(ctx context.Context, id string, filePath string, params *DeleteApiV4ProjectsIdRepositoryFilesFilePathParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryFilesFilePath request
	GetApiV4ProjectsIdRepositoryFilesFilePath(ctx context.Context, id string, filePath string, params *GetApiV4ProjectsIdRepositoryFilesFilePathParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// HeadApiV4ProjectsIdRepositoryFilesFilePath request
	HeadApiV4ProjectsIdRepositoryFilesFilePath(ctx context.Context, id string, filePath string, params *HeadApiV4ProjectsIdRepositoryFilesFilePathParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositoryFilesFilePathWithBody request with any body
	PostApiV4ProjectsIdRepositoryFilesFilePathWithBody(ctx context.Context, id string, filePath string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRepositoryFilesFilePath(ctx context.Context, id string, filePath string, body PostApiV4ProjectsIdRepositoryFilesFilePathJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdRepositoryFilesFilePathWithBody request with any body
	PutApiV4ProjectsIdRepositoryFilesFilePathWithBody(ctx context.Context, id string, filePath string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdRepositoryFilesFilePath(ctx context.Context, id string, filePath string, body PutApiV4ProjectsIdRepositoryFilesFilePathJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryFilesFilePathBlame request
	GetApiV4ProjectsIdRepositoryFilesFilePathBlame(ctx context.Context, id string, filePath string, params *GetApiV4ProjectsIdRepositoryFilesFilePathBlameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// HeadApiV4ProjectsIdRepositoryFilesFilePathBlame request
	HeadApiV4ProjectsIdRepositoryFilesFilePathBlame(ctx context.Context, id string, filePath string, params *HeadApiV4ProjectsIdRepositoryFilesFilePathBlameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryFilesFilePathRaw request
	GetApiV4ProjectsIdRepositoryFilesFilePathRaw(ctx context.Context, id string, filePath string, params *GetApiV4ProjectsIdRepositoryFilesFilePathRawParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryHealth request
	GetApiV4ProjectsIdRepositoryHealth(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryHealthParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryMergeBase request
	GetApiV4ProjectsIdRepositoryMergeBase(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryMergeBaseParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRepositoryMergedBranches request
	DeleteApiV4ProjectsIdRepositoryMergedBranches(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdRepositorySubmodulesSubmoduleWithBody request with any body
	PutApiV4ProjectsIdRepositorySubmodulesSubmoduleWithBody(ctx context.Context, id string, submodule string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdRepositorySubmodulesSubmodule(ctx context.Context, id string, submodule string, body PutApiV4ProjectsIdRepositorySubmodulesSubmoduleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryTags request
	GetApiV4ProjectsIdRepositoryTags(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryTagsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositoryTagsWithBody request with any body
	PostApiV4ProjectsIdRepositoryTagsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRepositoryTags(ctx context.Context, id string, body PostApiV4ProjectsIdRepositoryTagsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRepositoryTagsTagName request
	DeleteApiV4ProjectsIdRepositoryTagsTagName(ctx context.Context, id string, tagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryTagsTagName request
	GetApiV4ProjectsIdRepositoryTagsTagName(ctx context.Context, id string, tagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryTagsTagNameSignature request
	GetApiV4ProjectsIdRepositoryTagsTagNameSignature(ctx context.Context, id string, tagName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRepositoryTree request
	GetApiV4ProjectsIdRepositoryTree(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryTreeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRepositorySize request
	PostApiV4ProjectsIdRepositorySize(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdResourceGroups request
	GetApiV4ProjectsIdResourceGroups(ctx context.Context, id string, params *GetApiV4ProjectsIdResourceGroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdResourceGroupsKey request
	GetApiV4ProjectsIdResourceGroupsKey(ctx context.Context, id string, key string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdResourceGroupsKeyWithBody request with any body
	PutApiV4ProjectsIdResourceGroupsKeyWithBody(ctx context.Context, id string, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdResourceGroupsKey(ctx context.Context, id string, key string, body PutApiV4ProjectsIdResourceGroupsKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobs request
	GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobs(ctx context.Context, id string, key string, params *GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRestore request
	PostApiV4ProjectsIdRestore(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdRunners request
	GetApiV4ProjectsIdRunners(ctx context.Context, id string, params *GetApiV4ProjectsIdRunnersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRunnersWithBody request with any body
	PostApiV4ProjectsIdRunnersWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdRunners(ctx context.Context, id string, body PostApiV4ProjectsIdRunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdRunnersResetRegistrationToken request
	PostApiV4ProjectsIdRunnersResetRegistrationToken(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdRunnersRunnerId request
	DeleteApiV4ProjectsIdRunnersRunnerId(ctx context.Context, id string, runnerId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSecureFiles request
	GetApiV4ProjectsIdSecureFiles(ctx context.Context, id string, params *GetApiV4ProjectsIdSecureFilesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdSecureFilesWithBody request with any body
	PostApiV4ProjectsIdSecureFilesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdSecureFiles(ctx context.Context, id string, body PostApiV4ProjectsIdSecureFilesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdSecureFilesSecureFileId request
	DeleteApiV4ProjectsIdSecureFilesSecureFileId(ctx context.Context, id string, secureFileId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSecureFilesSecureFileId request
	GetApiV4ProjectsIdSecureFilesSecureFileId(ctx context.Context, id string, secureFileId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSecureFilesSecureFileIdDownload request
	GetApiV4ProjectsIdSecureFilesSecureFileIdDownload(ctx context.Context, id string, secureFileId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdServices request
	GetApiV4ProjectsIdServices(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesAppleAppStoreWithBody request with any body
	PutApiV4ProjectsIdServicesAppleAppStoreWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesAppleAppStore(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesAppleAppStoreJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesAsanaWithBody request with any body
	PutApiV4ProjectsIdServicesAsanaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesAsana(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesAsanaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesAssemblaWithBody request with any body
	PutApiV4ProjectsIdServicesAssemblaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesAssembla(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesAssemblaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesBambooWithBody request with any body
	PutApiV4ProjectsIdServicesBambooWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesBamboo(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesBambooJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesBugzillaWithBody request with any body
	PutApiV4ProjectsIdServicesBugzillaWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesBugzilla(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesBugzillaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesBuildkiteWithBody request with any body
	PutApiV4ProjectsIdServicesBuildkiteWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesBuildkite(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesBuildkiteJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesCampfireWithBody request with any body
	PutApiV4ProjectsIdServicesCampfireWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesCampfire(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesCampfireJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesClickupWithBody request with any body
	PutApiV4ProjectsIdServicesClickupWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesClickup(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesClickupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesConfluenceWithBody request with any body
	PutApiV4ProjectsIdServicesConfluenceWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesConfluence(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesConfluenceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesCustomIssueTrackerWithBody request with any body
	PutApiV4ProjectsIdServicesCustomIssueTrackerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesCustomIssueTracker(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesCustomIssueTrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesDatadogWithBody request with any body
	PutApiV4ProjectsIdServicesDatadogWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesDatadog(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesDatadogJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesDiffblueCoverWithBody request with any body
	PutApiV4ProjectsIdServicesDiffblueCoverWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesDiffblueCover(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesDiffblueCoverJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesDiscordWithBody request with any body
	PutApiV4ProjectsIdServicesDiscordWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesDiscord(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesDiscordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesDroneCiWithBody request with any body
	PutApiV4ProjectsIdServicesDroneCiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesDroneCi(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesDroneCiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesEmailsOnPushWithBody request with any body
	PutApiV4ProjectsIdServicesEmailsOnPushWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesEmailsOnPush(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesEmailsOnPushJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesEwmWithBody request with any body
	PutApiV4ProjectsIdServicesEwmWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesEwm(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesEwmJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesExternalWikiWithBody request with any body
	PutApiV4ProjectsIdServicesExternalWikiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesExternalWiki(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesExternalWikiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesGitGuardianWithBody request with any body
	PutApiV4ProjectsIdServicesGitGuardianWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesGitGuardian(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGitGuardianJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesGithubWithBody request with any body
	PutApiV4ProjectsIdServicesGithubWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesGithub(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesGitlabSlackApplicationWithBody request with any body
	PutApiV4ProjectsIdServicesGitlabSlackApplicationWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesGitlabSlackApplication(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGitlabSlackApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryWithBody request with any body
	PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistry(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationWithBody request with any body
	PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederation(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesGooglePlayWithBody request with any body
	PutApiV4ProjectsIdServicesGooglePlayWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesGooglePlay(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGooglePlayJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesHangoutsChatWithBody request with any body
	PutApiV4ProjectsIdServicesHangoutsChatWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesHangoutsChat(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesHangoutsChatJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesHarborWithBody request with any body
	PutApiV4ProjectsIdServicesHarborWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesHarbor(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesHarborJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesIrkerWithBody request with any body
	PutApiV4ProjectsIdServicesIrkerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesIrker(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesIrkerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesJenkinsWithBody request with any body
	PutApiV4ProjectsIdServicesJenkinsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesJenkins(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesJenkinsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesJiraWithBody request with any body
	PutApiV4ProjectsIdServicesJiraWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesJira(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesJiraJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesJiraCloudAppWithBody request with any body
	PutApiV4ProjectsIdServicesJiraCloudAppWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesJiraCloudApp(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesJiraCloudAppJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesMatrixWithBody request with any body
	PutApiV4ProjectsIdServicesMatrixWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesMatrix(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMatrixJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesMattermostWithBody request with any body
	PutApiV4ProjectsIdServicesMattermostWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesMattermost(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMattermostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesMattermostSlashCommandsWithBody request with any body
	PutApiV4ProjectsIdServicesMattermostSlashCommandsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesMattermostSlashCommands(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMattermostSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerWithBody request with any body
	PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdServicesMattermostSlashCommandsTrigger(ctx context.Context, id string, body PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesMicrosoftTeamsWithBody request with any body
	PutApiV4ProjectsIdServicesMicrosoftTeamsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesMicrosoftTeams(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMicrosoftTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesMockCiWithBody request with any body
	PutApiV4ProjectsIdServicesMockCiWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesMockCi(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMockCiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesMockMonitoringWithBody request with any body
	PutApiV4ProjectsIdServicesMockMonitoringWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesMockMonitoring(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMockMonitoringJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesPackagistWithBody request with any body
	PutApiV4ProjectsIdServicesPackagistWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesPackagist(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPackagistJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesPhorgeWithBody request with any body
	PutApiV4ProjectsIdServicesPhorgeWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesPhorge(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPhorgeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesPipelinesEmailWithBody request with any body
	PutApiV4ProjectsIdServicesPipelinesEmailWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesPipelinesEmail(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPipelinesEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesPivotaltrackerWithBody request with any body
	PutApiV4ProjectsIdServicesPivotaltrackerWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesPivotaltracker(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPivotaltrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesPumbleWithBody request with any body
	PutApiV4ProjectsIdServicesPumbleWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesPumble(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPumbleJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesPushoverWithBody request with any body
	PutApiV4ProjectsIdServicesPushoverWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesPushover(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPushoverJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesRedmineWithBody request with any body
	PutApiV4ProjectsIdServicesRedmineWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesRedmine(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesRedmineJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesSlackWithBody request with any body
	PutApiV4ProjectsIdServicesSlackWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesSlack(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesSlackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesSlackSlashCommandsWithBody request with any body
	PutApiV4ProjectsIdServicesSlackSlashCommandsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesSlackSlashCommands(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesSlackSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerWithBody request with any body
	PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdServicesSlackSlashCommandsTrigger(ctx context.Context, id string, body PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesSquashTmWithBody request with any body
	PutApiV4ProjectsIdServicesSquashTmWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesSquashTm(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesSquashTmJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesTeamcityWithBody request with any body
	PutApiV4ProjectsIdServicesTeamcityWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesTeamcity(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesTeamcityJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesTelegramWithBody request with any body
	PutApiV4ProjectsIdServicesTelegramWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesTelegram(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesTelegramJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesUnifyCircuitWithBody request with any body
	PutApiV4ProjectsIdServicesUnifyCircuitWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesUnifyCircuit(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesUnifyCircuitJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesWebexTeamsWithBody request with any body
	PutApiV4ProjectsIdServicesWebexTeamsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesWebexTeams(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesWebexTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesYoutrackWithBody request with any body
	PutApiV4ProjectsIdServicesYoutrackWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesYoutrack(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesYoutrackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdServicesZentaoWithBody request with any body
	PutApiV4ProjectsIdServicesZentaoWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdServicesZentao(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesZentaoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdServicesSlug request
	DeleteApiV4ProjectsIdServicesSlug(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdServicesSlug request
	GetApiV4ProjectsIdServicesSlug(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdShareWithBody request with any body
	PostApiV4ProjectsIdShareWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdShare(ctx context.Context, id string, body PostApiV4ProjectsIdShareJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdShareGroupId request
	DeleteApiV4ProjectsIdShareGroupId(ctx context.Context, id string, groupId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdShareLocations request
	GetApiV4ProjectsIdShareLocations(ctx context.Context, id int32, params *GetApiV4ProjectsIdShareLocationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnapshot request
	GetApiV4ProjectsIdSnapshot(ctx context.Context, id int32, params *GetApiV4ProjectsIdSnapshotParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippets request
	GetApiV4ProjectsIdSnippets(ctx context.Context, id string, params *GetApiV4ProjectsIdSnippetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdSnippetsWithBody request with any body
	PostApiV4ProjectsIdSnippetsWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdSnippets(ctx context.Context, id string, body PostApiV4ProjectsIdSnippetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdSnippetsSnippetId request
	DeleteApiV4ProjectsIdSnippetsSnippetId(ctx context.Context, id string, snippetId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippetsSnippetId request
	GetApiV4ProjectsIdSnippetsSnippetId(ctx context.Context, id string, snippetId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdSnippetsSnippetIdWithBody request with any body
	PutApiV4ProjectsIdSnippetsSnippetIdWithBody(ctx context.Context, id string, snippetId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdSnippetsSnippetId(ctx context.Context, id string, snippetId int32, body PutApiV4ProjectsIdSnippetsSnippetIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdAwardEmoji request
	GetApiV4ProjectsIdSnippetsSnippetIdAwardEmoji(ctx context.Context, id string, snippetId int32, params *GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiWithBody request with any body
	PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiWithBody(ctx context.Context, id int32, snippetId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdSnippetsSnippetIdAwardEmoji(ctx context.Context, id int32, snippetId int32, body PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardId request
	DeleteApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardId(ctx context.Context, id int32, snippetId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardId request
	GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardId(ctx context.Context, id int32, snippetId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRaw request
	GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRaw(ctx context.Context, id string, snippetId int32, ref string, filePath string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji request
	GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(ctx context.Context, id int32, snippetId int32, noteId int32, params *GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiWithBody request with any body
	PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiWithBody(ctx context.Context, id int32, snippetId int32, noteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmoji(ctx context.Context, id int32, snippetId int32, noteId int32, body PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId request
	DeleteApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(ctx context.Context, id int32, snippetId int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId request
	GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardId(ctx context.Context, id int32, snippetId int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdRaw request
	GetApiV4ProjectsIdSnippetsSnippetIdRaw(ctx context.Context, id string, snippetId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetail request
	GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetail(ctx context.Context, id string, snippetId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdStar request
	PostApiV4ProjectsIdStar(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdStarrers request
	GetApiV4ProjectsIdStarrers(ctx context.Context, id string, params *GetApiV4ProjectsIdStarrersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdStatistics request
	GetApiV4ProjectsIdStatistics(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdStatusesShaWithBody request with any body
	PostApiV4ProjectsIdStatusesShaWithBody(ctx context.Context, id string, sha string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdStatusesSha(ctx context.Context, id string, sha string, body PostApiV4ProjectsIdStatusesShaJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdStorage request
	GetApiV4ProjectsIdStorage(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdTemplatesType request
	GetApiV4ProjectsIdTemplatesType(ctx context.Context, id string, pType string, params *GetApiV4ProjectsIdTemplatesTypeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdTemplatesTypeName request
	GetApiV4ProjectsIdTemplatesTypeName(ctx context.Context, id string, pType string, name string, params *GetApiV4ProjectsIdTemplatesTypeNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdTerraformStateName request
	DeleteApiV4ProjectsIdTerraformStateName(ctx context.Context, id string, name int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdTerraformStateName request
	GetApiV4ProjectsIdTerraformStateName(ctx context.Context, id string, name string, params *GetApiV4ProjectsIdTerraformStateNameParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdTerraformStateName request
	PostApiV4ProjectsIdTerraformStateName(ctx context.Context, id string, name int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdTerraformStateNameLock request
	DeleteApiV4ProjectsIdTerraformStateNameLock(ctx context.Context, id string, name int32, params *DeleteApiV4ProjectsIdTerraformStateNameLockParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdTerraformStateNameLockWithBody request with any body
	PostApiV4ProjectsIdTerraformStateNameLockWithBody(ctx context.Context, id string, name int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdTerraformStateNameLock(ctx context.Context, id string, name int32, body PostApiV4ProjectsIdTerraformStateNameLockJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdTerraformStateNameVersionsSerial request
	DeleteApiV4ProjectsIdTerraformStateNameVersionsSerial(ctx context.Context, id string, name int32, serial int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdTerraformStateNameVersionsSerial request
	GetApiV4ProjectsIdTerraformStateNameVersionsSerial(ctx context.Context, id string, name string, serial int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdTransferWithBody request with any body
	PutApiV4ProjectsIdTransferWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdTransfer(ctx context.Context, id string, body PutApiV4ProjectsIdTransferJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdTransferLocations request
	GetApiV4ProjectsIdTransferLocations(ctx context.Context, id string, params *GetApiV4ProjectsIdTransferLocationsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdTriggers request
	GetApiV4ProjectsIdTriggers(ctx context.Context, id string, params *GetApiV4ProjectsIdTriggersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdTriggersWithBody request with any body
	PostApiV4ProjectsIdTriggersWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdTriggers(ctx context.Context, id string, body PostApiV4ProjectsIdTriggersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdTriggersTriggerId request
	DeleteApiV4ProjectsIdTriggersTriggerId(ctx context.Context, id string, triggerId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdTriggersTriggerId request
	GetApiV4ProjectsIdTriggersTriggerId(ctx context.Context, id string, triggerId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdTriggersTriggerIdWithBody request with any body
	PutApiV4ProjectsIdTriggersTriggerIdWithBody(ctx context.Context, id string, triggerId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdTriggersTriggerId(ctx context.Context, id string, triggerId int32, body PutApiV4ProjectsIdTriggersTriggerIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdUnarchive request
	PostApiV4ProjectsIdUnarchive(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdUnstar request
	PostApiV4ProjectsIdUnstar(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdUploads request
	GetApiV4ProjectsIdUploads(ctx context.Context, id int32, params *GetApiV4ProjectsIdUploadsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdUploadsWithBody request with any body
	PostApiV4ProjectsIdUploadsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdUploads(ctx context.Context, id int32, body PostApiV4ProjectsIdUploadsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdUploadsAuthorize request
	PostApiV4ProjectsIdUploadsAuthorize(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdUploadsSecretFilename request
	DeleteApiV4ProjectsIdUploadsSecretFilename(ctx context.Context, id int32, secret string, filename string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdUploadsSecretFilename request
	GetApiV4ProjectsIdUploadsSecretFilename(ctx context.Context, id int32, secret string, filename string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdUploadsUploadId request
	DeleteApiV4ProjectsIdUploadsUploadId(ctx context.Context, id int32, uploadId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdUploadsUploadId request
	GetApiV4ProjectsIdUploadsUploadId(ctx context.Context, id int32, uploadId int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdUsers request
	GetApiV4ProjectsIdUsers(ctx context.Context, id string, params *GetApiV4ProjectsIdUsersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdVariables request
	GetApiV4ProjectsIdVariables(ctx context.Context, id string, params *GetApiV4ProjectsIdVariablesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdVariablesWithBody request with any body
	PostApiV4ProjectsIdVariablesWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdVariables(ctx context.Context, id string, body PostApiV4ProjectsIdVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdVariablesKey request
	DeleteApiV4ProjectsIdVariablesKey(ctx context.Context, id string, key string, params *DeleteApiV4ProjectsIdVariablesKeyParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdVariablesKey request
	GetApiV4ProjectsIdVariablesKey(ctx context.Context, id string, key string, params *GetApiV4ProjectsIdVariablesKeyParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdVariablesKeyWithBody request with any body
	PutApiV4ProjectsIdVariablesKeyWithBody(ctx context.Context, id string, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdVariablesKey(ctx context.Context, id string, key string, body PutApiV4ProjectsIdVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdWikis request
	GetApiV4ProjectsIdWikis(ctx context.Context, id int32, params *GetApiV4ProjectsIdWikisParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdWikisWithBody request with any body
	PostApiV4ProjectsIdWikisWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdWikis(ctx context.Context, id int32, body PostApiV4ProjectsIdWikisJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4ProjectsIdWikisAttachmentsWithBody request with any body
	PostApiV4ProjectsIdWikisAttachmentsWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4ProjectsIdWikisAttachments(ctx context.Context, id int32, body PostApiV4ProjectsIdWikisAttachmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4ProjectsIdWikisSlug request
	DeleteApiV4ProjectsIdWikisSlug(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4ProjectsIdWikisSlug request
	GetApiV4ProjectsIdWikisSlug(ctx context.Context, id int32, slug string, params *GetApiV4ProjectsIdWikisSlugParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4ProjectsIdWikisSlugWithBody request with any body
	PutApiV4ProjectsIdWikisSlugWithBody(ctx context.Context, id int32, slug int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4ProjectsIdWikisSlug(ctx context.Context, id int32, slug int32, body PutApiV4ProjectsIdWikisSlugJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4RegistryRepositoriesId request
	GetApiV4RegistryRepositoriesId(ctx context.Context, id string, params *GetApiV4RegistryRepositoriesIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4Runners request
	DeleteApiV4Runners(ctx context.Context, params *DeleteApiV4RunnersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Runners request
	GetApiV4Runners(ctx context.Context, params *GetApiV4RunnersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4RunnersWithBody request with any body
	PostApiV4RunnersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Runners(ctx context.Context, body PostApiV4RunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4RunnersAll request
	GetApiV4RunnersAll(ctx context.Context, params *GetApiV4RunnersAllParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4RunnersManagers request
	DeleteApiV4RunnersManagers(ctx context.Context, params *DeleteApiV4RunnersManagersParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4RunnersResetAuthenticationTokenWithBody request with any body
	PostApiV4RunnersResetAuthenticationTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4RunnersResetAuthenticationToken(ctx context.Context, body PostApiV4RunnersResetAuthenticationTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4RunnersResetRegistrationToken request
	PostApiV4RunnersResetRegistrationToken(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4RunnersVerifyWithBody request with any body
	PostApiV4RunnersVerifyWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4RunnersVerify(ctx context.Context, body PostApiV4RunnersVerifyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4RunnersId request
	DeleteApiV4RunnersId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4RunnersId request
	GetApiV4RunnersId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4RunnersIdWithBody request with any body
	PutApiV4RunnersIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4RunnersId(ctx context.Context, id int32, body PutApiV4RunnersIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4RunnersIdJobs request
	GetApiV4RunnersIdJobs(ctx context.Context, id int32, params *GetApiV4RunnersIdJobsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4RunnersIdManagers request
	GetApiV4RunnersIdManagers(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4RunnersIdResetAuthenticationToken request
	PostApiV4RunnersIdResetAuthenticationToken(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4SlackTriggerWithBody request with any body
	PostApiV4SlackTriggerWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4SlackTrigger(ctx context.Context, body PostApiV4SlackTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Snippets request
	GetApiV4Snippets(ctx context.Context, params *GetApiV4SnippetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4SnippetsWithBody request with any body
	PostApiV4SnippetsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Snippets(ctx context.Context, body PostApiV4SnippetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4SnippetsAll request
	GetApiV4SnippetsAll(ctx context.Context, params *GetApiV4SnippetsAllParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4SnippetsPublic request
	GetApiV4SnippetsPublic(ctx context.Context, params *GetApiV4SnippetsPublicParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4SnippetsId request
	DeleteApiV4SnippetsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4SnippetsId request
	GetApiV4SnippetsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4SnippetsIdWithBody request with any body
	PutApiV4SnippetsIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4SnippetsId(ctx context.Context, id int32, body PutApiV4SnippetsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4SnippetsIdFilesRefFilePathRaw request
	GetApiV4SnippetsIdFilesRefFilePathRaw(ctx context.Context, id int32, ref string, filePath string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4SnippetsIdRaw request
	GetApiV4SnippetsIdRaw(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4SnippetsIdUserAgentDetail request
	GetApiV4SnippetsIdUserAgentDetail(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4SuggestionsBatchApplyWithBody request with any body
	PutApiV4SuggestionsBatchApplyWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4SuggestionsBatchApply(ctx context.Context, body PutApiV4SuggestionsBatchApplyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4SuggestionsIdApplyWithBody request with any body
	PutApiV4SuggestionsIdApplyWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4SuggestionsIdApply(ctx context.Context, id int32, body PutApiV4SuggestionsIdApplyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Topics request
	GetApiV4Topics(ctx context.Context, params *GetApiV4TopicsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4TopicsWithBody request with any body
	PostApiV4TopicsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4Topics(ctx context.Context, body PostApiV4TopicsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4TopicsMergeWithBody request with any body
	PostApiV4TopicsMergeWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4TopicsMerge(ctx context.Context, body PostApiV4TopicsMergeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV4TopicsId request
	DeleteApiV4TopicsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4TopicsId request
	GetApiV4TopicsId(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV4TopicsIdWithBody request with any body
	PutApiV4TopicsIdWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV4TopicsId(ctx context.Context, id int32, body PutApiV4TopicsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4UsageDataIncrementCounterWithBody request with any body
	PostApiV4UsageDataIncrementCounterWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4UsageDataIncrementCounter(ctx context.Context, body PostApiV4UsageDataIncrementCounterJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4UsageDataIncrementUniqueUsersWithBody request with any body
	PostApiV4UsageDataIncrementUniqueUsersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4UsageDataIncrementUniqueUsers(ctx context.Context, body PostApiV4UsageDataIncrementUniqueUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UsageDataMetricDefinitions request
	GetApiV4UsageDataMetricDefinitions(ctx context.Context, params *GetApiV4UsageDataMetricDefinitionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UsageDataNonSqlMetrics request
	GetApiV4UsageDataNonSqlMetrics(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UsageDataQueries request
	GetApiV4UsageDataQueries(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UsageDataServicePing request
	GetApiV4UsageDataServicePing(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4UsageDataTrackEventWithBody request with any body
	PostApiV4UsageDataTrackEventWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4UsageDataTrackEvent(ctx context.Context, body PostApiV4UsageDataTrackEventJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4UsageDataTrackEventsWithBody request with any body
	PostApiV4UsageDataTrackEventsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4UsageDataTrackEvents(ctx context.Context, body PostApiV4UsageDataTrackEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV4UserRunnersWithBody request with any body
	PostApiV4UserRunnersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV4UserRunners(ctx context.Context, body PostApiV4UserRunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UserCounts request
	GetApiV4UserCounts(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UsersIdEvents request
	GetApiV4UsersIdEvents(ctx context.Context, id string, params *GetApiV4UsersIdEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UsersUserIdContributedProjects request
	GetApiV4UsersUserIdContributedProjects(ctx context.Context, userId string, params *GetApiV4UsersUserIdContributedProjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UsersUserIdProjects request
	GetApiV4UsersUserIdProjects(ctx context.Context, userId string, params *GetApiV4UsersUserIdProjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4UsersUserIdStarredProjects request
	GetApiV4UsersUserIdStarredProjects(ctx context.Context, userId string, params *GetApiV4UsersUserIdStarredProjectsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4Version request
	GetApiV4Version(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV4WebCommitsPublicKey request
	GetApiV4WebCommitsPublicKey(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}
type ClientWithResponses struct {
	ClientInterface
}
type ClientWithResponsesInterface interface {
	// GetApiV4AdminBatchedBackgroundMigrationsWithResponse request
	GetApiV4AdminBatchedBackgroundMigrationsWithResponse(ctx context.Context, params *GetApiV4AdminBatchedBackgroundMigrationsParams, reqEditors ...RequestEditorFn) (*GetApiV4AdminBatchedBackgroundMigrationsResponse, error)

	// GetApiV4AdminBatchedBackgroundMigrationsIdWithResponse request
	GetApiV4AdminBatchedBackgroundMigrationsIdWithResponse(ctx context.Context, id int32, params *GetApiV4AdminBatchedBackgroundMigrationsIdParams, reqEditors ...RequestEditorFn) (*GetApiV4AdminBatchedBackgroundMigrationsIdResponse, error)

	// PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithBodyWithResponse request with any body
	PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse, error)

	PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithResponse(ctx context.Context, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse, error)

	// PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithBodyWithResponse request with any body
	PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse, error)

	PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithResponse(ctx context.Context, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse, error)

	// GetApiV4AdminCiVariablesWithResponse request
	GetApiV4AdminCiVariablesWithResponse(ctx context.Context, params *GetApiV4AdminCiVariablesParams, reqEditors ...RequestEditorFn) (*GetApiV4AdminCiVariablesResponse, error)

	// PostApiV4AdminCiVariablesWithBodyWithResponse request with any body
	PostApiV4AdminCiVariablesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4AdminCiVariablesResponse, error)

	PostApiV4AdminCiVariablesWithResponse(ctx context.Context, body PostApiV4AdminCiVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4AdminCiVariablesResponse, error)

	// DeleteApiV4AdminCiVariablesKeyWithResponse request
	DeleteApiV4AdminCiVariablesKeyWithResponse(ctx context.Context, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4AdminCiVariablesKeyResponse, error)

	// GetApiV4AdminCiVariablesKeyWithResponse request
	GetApiV4AdminCiVariablesKeyWithResponse(ctx context.Context, key string, reqEditors ...RequestEditorFn) (*GetApiV4AdminCiVariablesKeyResponse, error)

	// PutApiV4AdminCiVariablesKeyWithBodyWithResponse request with any body
	PutApiV4AdminCiVariablesKeyWithBodyWithResponse(ctx context.Context, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4AdminCiVariablesKeyResponse, error)

	PutApiV4AdminCiVariablesKeyWithResponse(ctx context.Context, key string, body PutApiV4AdminCiVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4AdminCiVariablesKeyResponse, error)

	// GetApiV4AdminClustersWithResponse request
	GetApiV4AdminClustersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4AdminClustersResponse, error)

	// PostApiV4AdminClustersAddWithBodyWithResponse request with any body
	PostApiV4AdminClustersAddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4AdminClustersAddResponse, error)

	PostApiV4AdminClustersAddWithResponse(ctx context.Context, body PostApiV4AdminClustersAddJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4AdminClustersAddResponse, error)

	// DeleteApiV4AdminClustersClusterIdWithResponse request
	DeleteApiV4AdminClustersClusterIdWithResponse(ctx context.Context, clusterId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4AdminClustersClusterIdResponse, error)

	// GetApiV4AdminClustersClusterIdWithResponse request
	GetApiV4AdminClustersClusterIdWithResponse(ctx context.Context, clusterId int32, reqEditors ...RequestEditorFn) (*GetApiV4AdminClustersClusterIdResponse, error)

	// PutApiV4AdminClustersClusterIdWithBodyWithResponse request with any body
	PutApiV4AdminClustersClusterIdWithBodyWithResponse(ctx context.Context, clusterId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4AdminClustersClusterIdResponse, error)

	PutApiV4AdminClustersClusterIdWithResponse(ctx context.Context, clusterId int32, body PutApiV4AdminClustersClusterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4AdminClustersClusterIdResponse, error)

	// GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameWithResponse request
	GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameWithResponse(ctx context.Context, databaseName string, tableName string, reqEditors ...RequestEditorFn) (*GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse, error)

	// PostApiV4AdminMigrationsTimestampMarkWithBodyWithResponse request with any body
	PostApiV4AdminMigrationsTimestampMarkWithBodyWithResponse(ctx context.Context, timestamp int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4AdminMigrationsTimestampMarkResponse, error)

	PostApiV4AdminMigrationsTimestampMarkWithResponse(ctx context.Context, timestamp int32, body PostApiV4AdminMigrationsTimestampMarkJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4AdminMigrationsTimestampMarkResponse, error)

	// GetApiV4ApplicationAppearanceWithResponse request
	GetApiV4ApplicationAppearanceWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4ApplicationAppearanceResponse, error)

	// PutApiV4ApplicationAppearanceWithBodyWithResponse request with any body
	PutApiV4ApplicationAppearanceWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ApplicationAppearanceResponse, error)

	// GetApiV4ApplicationPlanLimitsWithResponse request
	GetApiV4ApplicationPlanLimitsWithResponse(ctx context.Context, params *GetApiV4ApplicationPlanLimitsParams, reqEditors ...RequestEditorFn) (*GetApiV4ApplicationPlanLimitsResponse, error)

	// PutApiV4ApplicationPlanLimitsWithBodyWithResponse request with any body
	PutApiV4ApplicationPlanLimitsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ApplicationPlanLimitsResponse, error)

	PutApiV4ApplicationPlanLimitsWithResponse(ctx context.Context, body PutApiV4ApplicationPlanLimitsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ApplicationPlanLimitsResponse, error)

	// GetApiV4ApplicationStatisticsWithResponse request
	GetApiV4ApplicationStatisticsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4ApplicationStatisticsResponse, error)

	// GetApiV4ApplicationsWithResponse request
	GetApiV4ApplicationsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4ApplicationsResponse, error)

	// PostApiV4ApplicationsWithBodyWithResponse request with any body
	PostApiV4ApplicationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ApplicationsResponse, error)

	PostApiV4ApplicationsWithResponse(ctx context.Context, body PostApiV4ApplicationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ApplicationsResponse, error)

	// DeleteApiV4ApplicationsIdWithResponse request
	DeleteApiV4ApplicationsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ApplicationsIdResponse, error)

	// PostApiV4ApplicationsIdRenewSecretWithResponse request
	PostApiV4ApplicationsIdRenewSecretWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4ApplicationsIdRenewSecretResponse, error)

	// GetApiV4AvatarWithResponse request
	GetApiV4AvatarWithResponse(ctx context.Context, params *GetApiV4AvatarParams, reqEditors ...RequestEditorFn) (*GetApiV4AvatarResponse, error)

	// GetApiV4BroadcastMessagesWithResponse request
	GetApiV4BroadcastMessagesWithResponse(ctx context.Context, params *GetApiV4BroadcastMessagesParams, reqEditors ...RequestEditorFn) (*GetApiV4BroadcastMessagesResponse, error)

	// PostApiV4BroadcastMessagesWithBodyWithResponse request with any body
	PostApiV4BroadcastMessagesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4BroadcastMessagesResponse, error)

	PostApiV4BroadcastMessagesWithResponse(ctx context.Context, body PostApiV4BroadcastMessagesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4BroadcastMessagesResponse, error)

	// DeleteApiV4BroadcastMessagesIdWithResponse request
	DeleteApiV4BroadcastMessagesIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4BroadcastMessagesIdResponse, error)

	// GetApiV4BroadcastMessagesIdWithResponse request
	GetApiV4BroadcastMessagesIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4BroadcastMessagesIdResponse, error)

	// PutApiV4BroadcastMessagesIdWithBodyWithResponse request with any body
	PutApiV4BroadcastMessagesIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4BroadcastMessagesIdResponse, error)

	PutApiV4BroadcastMessagesIdWithResponse(ctx context.Context, id int32, body PutApiV4BroadcastMessagesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4BroadcastMessagesIdResponse, error)

	// GetApiV4BulkImportsWithResponse request
	GetApiV4BulkImportsWithResponse(ctx context.Context, params *GetApiV4BulkImportsParams, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsResponse, error)

	// PostApiV4BulkImportsWithBodyWithResponse request with any body
	PostApiV4BulkImportsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4BulkImportsResponse, error)

	PostApiV4BulkImportsWithFormdataBodyWithResponse(ctx context.Context, body PostApiV4BulkImportsFormdataRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4BulkImportsResponse, error)

	// GetApiV4BulkImportsEntitiesWithResponse request
	GetApiV4BulkImportsEntitiesWithResponse(ctx context.Context, params *GetApiV4BulkImportsEntitiesParams, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsEntitiesResponse, error)

	// GetApiV4BulkImportsImportIdWithResponse request
	GetApiV4BulkImportsImportIdWithResponse(ctx context.Context, importId int32, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsImportIdResponse, error)

	// PostApiV4BulkImportsImportIdCancelWithResponse request
	PostApiV4BulkImportsImportIdCancelWithResponse(ctx context.Context, importId int32, reqEditors ...RequestEditorFn) (*PostApiV4BulkImportsImportIdCancelResponse, error)

	// GetApiV4BulkImportsImportIdEntitiesWithResponse request
	GetApiV4BulkImportsImportIdEntitiesWithResponse(ctx context.Context, importId int32, params *GetApiV4BulkImportsImportIdEntitiesParams, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsImportIdEntitiesResponse, error)

	// GetApiV4BulkImportsImportIdEntitiesEntityIdWithResponse request
	GetApiV4BulkImportsImportIdEntitiesEntityIdWithResponse(ctx context.Context, importId int32, entityId int32, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsImportIdEntitiesEntityIdResponse, error)

	// GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresWithResponse request
	GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresWithResponse(ctx context.Context, importId int32, entityId int32, reqEditors ...RequestEditorFn) (*GetApiV4BulkImportsImportIdEntitiesEntityIdFailuresResponse, error)

	// PostApiV4ContainerRegistryEventEventsWithResponse request
	PostApiV4ContainerRegistryEventEventsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4ContainerRegistryEventEventsResponse, error)

	// GetApiV4DeployKeysWithResponse request
	GetApiV4DeployKeysWithResponse(ctx context.Context, params *GetApiV4DeployKeysParams, reqEditors ...RequestEditorFn) (*GetApiV4DeployKeysResponse, error)

	// PostApiV4DeployKeysWithBodyWithResponse request with any body
	PostApiV4DeployKeysWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4DeployKeysResponse, error)

	PostApiV4DeployKeysWithResponse(ctx context.Context, body PostApiV4DeployKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4DeployKeysResponse, error)

	// GetApiV4DeployTokensWithResponse request
	GetApiV4DeployTokensWithResponse(ctx context.Context, params *GetApiV4DeployTokensParams, reqEditors ...RequestEditorFn) (*GetApiV4DeployTokensResponse, error)

	// GetApiV4DiscoverCertBasedClustersWithResponse request
	GetApiV4DiscoverCertBasedClustersWithResponse(ctx context.Context, params *GetApiV4DiscoverCertBasedClustersParams, reqEditors ...RequestEditorFn) (*GetApiV4DiscoverCertBasedClustersResponse, error)

	// GetApiV4EventsWithResponse request
	GetApiV4EventsWithResponse(ctx context.Context, params *GetApiV4EventsParams, reqEditors ...RequestEditorFn) (*GetApiV4EventsResponse, error)

	// GetApiV4FeatureFlagsUnleashProjectIdWithResponse request
	GetApiV4FeatureFlagsUnleashProjectIdWithResponse(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdParams, reqEditors ...RequestEditorFn) (*GetApiV4FeatureFlagsUnleashProjectIdResponse, error)

	// GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesWithResponse request
	GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesWithResponse(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesParams, reqEditors ...RequestEditorFn) (*GetApiV4FeatureFlagsUnleashProjectIdClientFeaturesResponse, error)

	// PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithBodyWithResponse request with any body
	PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithBodyWithResponse(ctx context.Context, projectId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse, error)

	PostApiV4FeatureFlagsUnleashProjectIdClientMetricsWithResponse(ctx context.Context, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientMetricsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4FeatureFlagsUnleashProjectIdClientMetricsResponse, error)

	// PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithBodyWithResponse request with any body
	PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithBodyWithResponse(ctx context.Context, projectId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse, error)

	PostApiV4FeatureFlagsUnleashProjectIdClientRegisterWithResponse(ctx context.Context, projectId string, body PostApiV4FeatureFlagsUnleashProjectIdClientRegisterJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4FeatureFlagsUnleashProjectIdClientRegisterResponse, error)

	// GetApiV4FeatureFlagsUnleashProjectIdFeaturesWithResponse request
	GetApiV4FeatureFlagsUnleashProjectIdFeaturesWithResponse(ctx context.Context, projectId string, params *GetApiV4FeatureFlagsUnleashProjectIdFeaturesParams, reqEditors ...RequestEditorFn) (*GetApiV4FeatureFlagsUnleashProjectIdFeaturesResponse, error)

	// GetApiV4FeaturesWithResponse request
	GetApiV4FeaturesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4FeaturesResponse, error)

	// GetApiV4FeaturesDefinitionsWithResponse request
	GetApiV4FeaturesDefinitionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4FeaturesDefinitionsResponse, error)

	// DeleteApiV4FeaturesNameWithResponse request
	DeleteApiV4FeaturesNameWithResponse(ctx context.Context, name int32, reqEditors ...RequestEditorFn) (*DeleteApiV4FeaturesNameResponse, error)

	// PostApiV4FeaturesNameWithBodyWithResponse request with any body
	PostApiV4FeaturesNameWithBodyWithResponse(ctx context.Context, name int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4FeaturesNameResponse, error)

	PostApiV4FeaturesNameWithResponse(ctx context.Context, name int32, body PostApiV4FeaturesNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4FeaturesNameResponse, error)

	// PostApiV4GeoNodeProxyIdGraphqlWithResponse request
	PostApiV4GeoNodeProxyIdGraphqlWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4GeoNodeProxyIdGraphqlResponse, error)

	// GetApiV4GeoProxyWithResponse request
	GetApiV4GeoProxyWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4GeoProxyResponse, error)

	// PostApiV4GeoProxyGitSshInfoRefsReceivePackWithBodyWithResponse request with any body
	PostApiV4GeoProxyGitSshInfoRefsReceivePackWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse, error)

	PostApiV4GeoProxyGitSshInfoRefsReceivePackWithResponse(ctx context.Context, body PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse, error)

	// PostApiV4GeoProxyGitSshInfoRefsUploadPackWithBodyWithResponse request with any body
	PostApiV4GeoProxyGitSshInfoRefsUploadPackWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse, error)

	PostApiV4GeoProxyGitSshInfoRefsUploadPackWithResponse(ctx context.Context, body PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse, error)

	// PostApiV4GeoProxyGitSshReceivePackWithBodyWithResponse request with any body
	PostApiV4GeoProxyGitSshReceivePackWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshReceivePackResponse, error)

	PostApiV4GeoProxyGitSshReceivePackWithResponse(ctx context.Context, body PostApiV4GeoProxyGitSshReceivePackJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshReceivePackResponse, error)

	// PostApiV4GeoProxyGitSshUploadPackWithBodyWithResponse request with any body
	PostApiV4GeoProxyGitSshUploadPackWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshUploadPackResponse, error)

	PostApiV4GeoProxyGitSshUploadPackWithResponse(ctx context.Context, body PostApiV4GeoProxyGitSshUploadPackJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshUploadPackResponse, error)

	// GetApiV4GeoRepositoriesGlRepositoryPipelineRefsWithResponse request
	GetApiV4GeoRepositoriesGlRepositoryPipelineRefsWithResponse(ctx context.Context, glRepository string, reqEditors ...RequestEditorFn) (*GetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse, error)

	// GetApiV4GeoRetrieveReplicableNameReplicableIdWithResponse request
	GetApiV4GeoRetrieveReplicableNameReplicableIdWithResponse(ctx context.Context, replicableName string, replicableId int32, reqEditors ...RequestEditorFn) (*GetApiV4GeoRetrieveReplicableNameReplicableIdResponse, error)

	// PostApiV4GeoStatusWithBodyWithResponse request with any body
	PostApiV4GeoStatusWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoStatusResponse, error)

	PostApiV4GeoStatusWithResponse(ctx context.Context, body PostApiV4GeoStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoStatusResponse, error)

	// GetApiV4GroupIdPackagesComposerpackageNameWithResponse request
	GetApiV4GroupIdPackagesComposerpackageNameWithResponse(ctx context.Context, id string, params *GetApiV4GroupIdPackagesComposerpackageNameParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupIdPackagesComposerpackageNameResponse, error)

	// GetApiV4GroupIdPackagesComposerPShaWithResponse request
	GetApiV4GroupIdPackagesComposerPShaWithResponse(ctx context.Context, id string, sha string, reqEditors ...RequestEditorFn) (*GetApiV4GroupIdPackagesComposerPShaResponse, error)

	// GetApiV4GroupIdPackagesComposerP2packageNameWithResponse request
	GetApiV4GroupIdPackagesComposerP2packageNameWithResponse(ctx context.Context, id string, params *GetApiV4GroupIdPackagesComposerP2packageNameParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupIdPackagesComposerP2packageNameResponse, error)

	// GetApiV4GroupIdPackagesComposerPackagesWithResponse request
	GetApiV4GroupIdPackagesComposerPackagesWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4GroupIdPackagesComposerPackagesResponse, error)

	// GetApiV4GroupsWithResponse request
	GetApiV4GroupsWithResponse(ctx context.Context, params *GetApiV4GroupsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsResponse, error)

	// PostApiV4GroupsWithBodyWithResponse request with any body
	PostApiV4GroupsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsResponse, error)

	PostApiV4GroupsWithResponse(ctx context.Context, body PostApiV4GroupsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsResponse, error)

	// PostApiV4GroupsImportWithBodyWithResponse request with any body
	PostApiV4GroupsImportWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsImportResponse, error)

	// PostApiV4GroupsImportAuthorizeWithResponse request
	PostApiV4GroupsImportAuthorizeWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4GroupsImportAuthorizeResponse, error)

	// DeleteApiV4GroupsIdWithResponse request
	DeleteApiV4GroupsIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdResponse, error)

	// GetApiV4GroupsIdWithResponse request
	GetApiV4GroupsIdWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdResponse, error)

	// PutApiV4GroupsIdWithBodyWithResponse request with any body
	PutApiV4GroupsIdWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdResponse, error)

	PutApiV4GroupsIdWithResponse(ctx context.Context, id string, body PutApiV4GroupsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdResponse, error)

	// GetApiV4GroupsIdDebianDistributionsWithResponse request
	GetApiV4GroupsIdDebianDistributionsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdDebianDistributionsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdDebianDistributionsResponse, error)

	// PostApiV4GroupsIdDebianDistributionsWithBodyWithResponse request with any body
	PostApiV4GroupsIdDebianDistributionsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdDebianDistributionsResponse, error)

	PostApiV4GroupsIdDebianDistributionsWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdDebianDistributionsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdDebianDistributionsResponse, error)

	// DeleteApiV4GroupsIdDebianDistributionsCodenameWithResponse request
	DeleteApiV4GroupsIdDebianDistributionsCodenameWithResponse(ctx context.Context, id string, codename string, params *DeleteApiV4GroupsIdDebianDistributionsCodenameParams, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdDebianDistributionsCodenameResponse, error)

	// GetApiV4GroupsIdDebianDistributionsCodenameWithResponse request
	GetApiV4GroupsIdDebianDistributionsCodenameWithResponse(ctx context.Context, id string, codename string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdDebianDistributionsCodenameResponse, error)

	// PutApiV4GroupsIdDebianDistributionsCodenameWithBodyWithResponse request with any body
	PutApiV4GroupsIdDebianDistributionsCodenameWithBodyWithResponse(ctx context.Context, id string, codename string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdDebianDistributionsCodenameResponse, error)

	PutApiV4GroupsIdDebianDistributionsCodenameWithResponse(ctx context.Context, id string, codename string, body PutApiV4GroupsIdDebianDistributionsCodenameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdDebianDistributionsCodenameResponse, error)

	// GetApiV4GroupsIdDebianDistributionsCodenameKeyAscWithResponse request
	GetApiV4GroupsIdDebianDistributionsCodenameKeyAscWithResponse(ctx context.Context, id string, codename string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdDebianDistributionsCodenameKeyAscResponse, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionInreleaseWithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionInreleaseWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionInreleaseParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionInreleaseResponse, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseWithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseResponse, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpgWithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpgWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpgParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionReleaseGpgResponse, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesWithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesWithResponse(ctx context.Context, id string, component string, architecture string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesResponse, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256WithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256WithResponse(ctx context.Context, id string, component string, architecture string, fileSha256 int32, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesWithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesWithResponse(ctx context.Context, id string, component string, architecture string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesResponse, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256WithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256WithResponse(ctx context.Context, id string, component string, architecture string, fileSha256 int32, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256Response, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSourcesWithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSourcesWithResponse(ctx context.Context, id string, component string, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSourcesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceSourcesResponse, error)

	// GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256WithResponse request
	GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256WithResponse(ctx context.Context, id string, component string, fileSha256 int32, params *GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256Response, error)

	// GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameWithResponse request
	GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameWithResponse(ctx context.Context, id string, distribution string, projectId int32, letter string, packageName string, packageVersion string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesDebianPoolDistributionProjectIdLetterPackageNamePackageVersionFileNameResponse, error)

	// GetApiV4GroupsIdPackagesMavenpathFileNameWithResponse request
	GetApiV4GroupsIdPackagesMavenpathFileNameWithResponse(ctx context.Context, id string, fileName string, params *GetApiV4GroupsIdPackagesMavenpathFileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesMavenpathFileNameResponse, error)

	// GetApiV4GroupsIdPackagesNpmpackageNameWithResponse request
	GetApiV4GroupsIdPackagesNpmpackageNameWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesNpmpackageNameParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesNpmpackageNameResponse, error)

	// PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkWithResponse request
	PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdPackagesNpmNpmV1SecurityAdvisoriesBulkResponse, error)

	// PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuickWithResponse request
	PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuickWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdPackagesNpmNpmV1SecurityAuditsQuickResponse, error)

	// GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsWithResponse request
	GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsResponse, error)

	// DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagWithResponse request
	DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagWithResponse(ctx context.Context, id string, tag string, params *DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagParams, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagResponse, error)

	// PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagWithBodyWithResponse request with any body
	PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagWithBodyWithResponse(ctx context.Context, id string, tag string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagResponse, error)

	PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagWithResponse(ctx context.Context, id string, tag string, body PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdPackagesNpmPackagepackageNameDistTagsTagResponse, error)

	// GetApiV4GroupsIdPackagesNugetIndexWithResponse request
	GetApiV4GroupsIdPackagesNugetIndexWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesNugetIndexResponse, error)

	// GetApiV4GroupsIdPackagesNugetQueryWithResponse request
	GetApiV4GroupsIdPackagesNugetQueryWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdPackagesNugetQueryParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesNugetQueryResponse, error)

	// GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameWithResponse request
	GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameResponse, error)

	// GetApiV4GroupsIdPackagesNugetV2WithResponse request
	GetApiV4GroupsIdPackagesNugetV2WithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesNugetV2Response, error)

	// GetApiV4GroupsIdPackagesNugetV2MetadataWithResponse request
	GetApiV4GroupsIdPackagesNugetV2MetadataWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesNugetV2MetadataResponse, error)

	// GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifierWithResponse request
	GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifierWithResponse(ctx context.Context, id int32, sha256 string, params *GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifierParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesPypiFilesSha256fileIdentifierResponse, error)

	// GetApiV4GroupsIdPackagesPypiSimpleWithResponse request
	GetApiV4GroupsIdPackagesPypiSimpleWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesPypiSimpleResponse, error)

	// GetApiV4GroupsIdPackagesPypiSimplepackageNameWithResponse request
	GetApiV4GroupsIdPackagesPypiSimplepackageNameWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdPackagesPypiSimplepackageNameParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesPypiSimplepackageNameResponse, error)

	// GetApiV4GroupsIdAccessRequestsWithResponse request
	GetApiV4GroupsIdAccessRequestsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdAccessRequestsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdAccessRequestsResponse, error)

	// PostApiV4GroupsIdAccessRequestsWithResponse request
	PostApiV4GroupsIdAccessRequestsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdAccessRequestsResponse, error)

	// DeleteApiV4GroupsIdAccessRequestsUserIdWithResponse request
	DeleteApiV4GroupsIdAccessRequestsUserIdWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdAccessRequestsUserIdResponse, error)

	// PutApiV4GroupsIdAccessRequestsUserIdApproveWithBodyWithResponse request with any body
	PutApiV4GroupsIdAccessRequestsUserIdApproveWithBodyWithResponse(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdAccessRequestsUserIdApproveResponse, error)

	PutApiV4GroupsIdAccessRequestsUserIdApproveWithResponse(ctx context.Context, id string, userId int32, body PutApiV4GroupsIdAccessRequestsUserIdApproveJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdAccessRequestsUserIdApproveResponse, error)

	// PostApiV4GroupsIdAccessTokensSelfRotateWithBodyWithResponse request with any body
	PostApiV4GroupsIdAccessTokensSelfRotateWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdAccessTokensSelfRotateResponse, error)

	PostApiV4GroupsIdAccessTokensSelfRotateWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdAccessTokensSelfRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdAccessTokensSelfRotateResponse, error)

	// PostApiV4GroupsIdArchiveWithResponse request
	PostApiV4GroupsIdArchiveWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdArchiveResponse, error)

	// GetApiV4GroupsIdAuditEventsWithResponse request
	GetApiV4GroupsIdAuditEventsWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdAuditEventsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdAuditEventsResponse, error)

	// GetApiV4GroupsIdAuditEventsAuditEventIdWithResponse request
	GetApiV4GroupsIdAuditEventsAuditEventIdWithResponse(ctx context.Context, id int32, auditEventId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdAuditEventsAuditEventIdResponse, error)

	// GetApiV4GroupsIdAvatarWithResponse request
	GetApiV4GroupsIdAvatarWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdAvatarResponse, error)

	// GetApiV4GroupsIdBadgesWithResponse request
	GetApiV4GroupsIdBadgesWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdBadgesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdBadgesResponse, error)

	// PostApiV4GroupsIdBadgesWithBodyWithResponse request with any body
	PostApiV4GroupsIdBadgesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdBadgesResponse, error)

	PostApiV4GroupsIdBadgesWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdBadgesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdBadgesResponse, error)

	// GetApiV4GroupsIdBadgesRenderWithResponse request
	GetApiV4GroupsIdBadgesRenderWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdBadgesRenderParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdBadgesRenderResponse, error)

	// DeleteApiV4GroupsIdBadgesBadgeIdWithResponse request
	DeleteApiV4GroupsIdBadgesBadgeIdWithResponse(ctx context.Context, id string, badgeId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdBadgesBadgeIdResponse, error)

	// GetApiV4GroupsIdBadgesBadgeIdWithResponse request
	GetApiV4GroupsIdBadgesBadgeIdWithResponse(ctx context.Context, id string, badgeId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdBadgesBadgeIdResponse, error)

	// PutApiV4GroupsIdBadgesBadgeIdWithBodyWithResponse request with any body
	PutApiV4GroupsIdBadgesBadgeIdWithBodyWithResponse(ctx context.Context, id string, badgeId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdBadgesBadgeIdResponse, error)

	PutApiV4GroupsIdBadgesBadgeIdWithResponse(ctx context.Context, id string, badgeId int32, body PutApiV4GroupsIdBadgesBadgeIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdBadgesBadgeIdResponse, error)

	// GetApiV4GroupsIdBillableMembersWithResponse request
	GetApiV4GroupsIdBillableMembersWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdBillableMembersParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdBillableMembersResponse, error)

	// DeleteApiV4GroupsIdBillableMembersUserIdWithResponse request
	DeleteApiV4GroupsIdBillableMembersUserIdWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdBillableMembersUserIdResponse, error)

	// GetApiV4GroupsIdBillableMembersUserIdIndirectWithResponse request
	GetApiV4GroupsIdBillableMembersUserIdIndirectWithResponse(ctx context.Context, id string, userId int32, params *GetApiV4GroupsIdBillableMembersUserIdIndirectParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdBillableMembersUserIdIndirectResponse, error)

	// GetApiV4GroupsIdBillableMembersUserIdMembershipsWithResponse request
	GetApiV4GroupsIdBillableMembersUserIdMembershipsWithResponse(ctx context.Context, id string, userId int32, params *GetApiV4GroupsIdBillableMembersUserIdMembershipsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdBillableMembersUserIdMembershipsResponse, error)

	// GetApiV4GroupsIdClustersWithResponse request
	GetApiV4GroupsIdClustersWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdClustersParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdClustersResponse, error)

	// PostApiV4GroupsIdClustersUserWithBodyWithResponse request with any body
	PostApiV4GroupsIdClustersUserWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdClustersUserResponse, error)

	PostApiV4GroupsIdClustersUserWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdClustersUserJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdClustersUserResponse, error)

	// DeleteApiV4GroupsIdClustersClusterIdWithResponse request
	DeleteApiV4GroupsIdClustersClusterIdWithResponse(ctx context.Context, id string, clusterId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdClustersClusterIdResponse, error)

	// GetApiV4GroupsIdClustersClusterIdWithResponse request
	GetApiV4GroupsIdClustersClusterIdWithResponse(ctx context.Context, id string, clusterId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdClustersClusterIdResponse, error)

	// PutApiV4GroupsIdClustersClusterIdWithBodyWithResponse request with any body
	PutApiV4GroupsIdClustersClusterIdWithBodyWithResponse(ctx context.Context, id string, clusterId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdClustersClusterIdResponse, error)

	PutApiV4GroupsIdClustersClusterIdWithResponse(ctx context.Context, id string, clusterId int32, body PutApiV4GroupsIdClustersClusterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdClustersClusterIdResponse, error)

	// GetApiV4GroupsIdCustomAttributesWithResponse request
	GetApiV4GroupsIdCustomAttributesWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdCustomAttributesResponse, error)

	// DeleteApiV4GroupsIdCustomAttributesKeyWithResponse request
	DeleteApiV4GroupsIdCustomAttributesKeyWithResponse(ctx context.Context, id int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdCustomAttributesKeyResponse, error)

	// GetApiV4GroupsIdCustomAttributesKeyWithResponse request
	GetApiV4GroupsIdCustomAttributesKeyWithResponse(ctx context.Context, id int32, key string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdCustomAttributesKeyResponse, error)

	// PutApiV4GroupsIdCustomAttributesKeyWithBodyWithResponse request with any body
	PutApiV4GroupsIdCustomAttributesKeyWithBodyWithResponse(ctx context.Context, id int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdCustomAttributesKeyResponse, error)

	PutApiV4GroupsIdCustomAttributesKeyWithResponse(ctx context.Context, id int32, key string, body PutApiV4GroupsIdCustomAttributesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdCustomAttributesKeyResponse, error)

	// DeleteApiV4GroupsIdDependencyProxyCacheWithResponse request
	DeleteApiV4GroupsIdDependencyProxyCacheWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdDependencyProxyCacheResponse, error)

	// GetApiV4GroupsIdDeployTokensWithResponse request
	GetApiV4GroupsIdDeployTokensWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdDeployTokensParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdDeployTokensResponse, error)

	// PostApiV4GroupsIdDeployTokensWithBodyWithResponse request with any body
	PostApiV4GroupsIdDeployTokensWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdDeployTokensResponse, error)

	PostApiV4GroupsIdDeployTokensWithResponse(ctx context.Context, id int32, body PostApiV4GroupsIdDeployTokensJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdDeployTokensResponse, error)

	// DeleteApiV4GroupsIdDeployTokensTokenIdWithResponse request
	DeleteApiV4GroupsIdDeployTokensTokenIdWithResponse(ctx context.Context, id int32, tokenId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdDeployTokensTokenIdResponse, error)

	// GetApiV4GroupsIdDeployTokensTokenIdWithResponse request
	GetApiV4GroupsIdDeployTokensTokenIdWithResponse(ctx context.Context, id int32, tokenId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdDeployTokensTokenIdResponse, error)

	// GetApiV4GroupsIdDescendantGroupsWithResponse request
	GetApiV4GroupsIdDescendantGroupsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdDescendantGroupsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdDescendantGroupsResponse, error)

	// GetApiV4GroupsIdEpicsEpicIidAwardEmojiWithResponse request
	GetApiV4GroupsIdEpicsEpicIidAwardEmojiWithResponse(ctx context.Context, id string, epicIid int32, params *GetApiV4GroupsIdEpicsEpicIidAwardEmojiParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdEpicsEpicIidAwardEmojiResponse, error)

	// PostApiV4GroupsIdEpicsEpicIidAwardEmojiWithBodyWithResponse request with any body
	PostApiV4GroupsIdEpicsEpicIidAwardEmojiWithBodyWithResponse(ctx context.Context, id int32, epicIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdEpicsEpicIidAwardEmojiResponse, error)

	PostApiV4GroupsIdEpicsEpicIidAwardEmojiWithResponse(ctx context.Context, id int32, epicIid int32, body PostApiV4GroupsIdEpicsEpicIidAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdEpicsEpicIidAwardEmojiResponse, error)

	// DeleteApiV4GroupsIdEpicsEpicIidAwardEmojiAwardIdWithResponse request
	DeleteApiV4GroupsIdEpicsEpicIidAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, epicIid int32, awardId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdEpicsEpicIidAwardEmojiAwardIdResponse, error)

	// GetApiV4GroupsIdEpicsEpicIidAwardEmojiAwardIdWithResponse request
	GetApiV4GroupsIdEpicsEpicIidAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, epicIid int32, awardId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdEpicsEpicIidAwardEmojiAwardIdResponse, error)

	// GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiWithResponse request
	GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiWithResponse(ctx context.Context, id int32, epicIid int32, noteId int32, params *GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiResponse, error)

	// PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiWithBodyWithResponse request with any body
	PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiWithBodyWithResponse(ctx context.Context, id int32, epicIid int32, noteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiResponse, error)

	PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiWithResponse(ctx context.Context, id int32, epicIid int32, noteId int32, body PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiResponse, error)

	// DeleteApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdWithResponse request
	DeleteApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, epicIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdResponse, error)

	// GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdWithResponse request
	GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, epicIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdEpicsEpicIidNotesNoteIdAwardEmojiAwardIdResponse, error)

	// PostApiV4GroupsIdExportWithResponse request
	PostApiV4GroupsIdExportWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdExportResponse, error)

	// GetApiV4GroupsIdExportDownloadWithResponse request
	GetApiV4GroupsIdExportDownloadWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdExportDownloadResponse, error)

	// PostApiV4GroupsIdExportRelationsWithBodyWithResponse request with any body
	PostApiV4GroupsIdExportRelationsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdExportRelationsResponse, error)

	PostApiV4GroupsIdExportRelationsWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdExportRelationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdExportRelationsResponse, error)

	// GetApiV4GroupsIdExportRelationsDownloadWithResponse request
	GetApiV4GroupsIdExportRelationsDownloadWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdExportRelationsDownloadParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdExportRelationsDownloadResponse, error)

	// GetApiV4GroupsIdExportRelationsStatusWithResponse request
	GetApiV4GroupsIdExportRelationsStatusWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdExportRelationsStatusParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdExportRelationsStatusResponse, error)

	// GetApiV4GroupsIdGroupsSharedWithResponse request
	GetApiV4GroupsIdGroupsSharedWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdGroupsSharedParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdGroupsSharedResponse, error)

	// GetApiV4GroupsIdIntegrationsWithResponse request
	GetApiV4GroupsIdIntegrationsWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdIntegrationsResponse, error)

	// PutApiV4GroupsIdIntegrationsAppleAppStoreWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsAppleAppStoreWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsAppleAppStoreResponse, error)

	PutApiV4GroupsIdIntegrationsAppleAppStoreWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsAppleAppStoreJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsAppleAppStoreResponse, error)

	// PutApiV4GroupsIdIntegrationsAsanaWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsAsanaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsAsanaResponse, error)

	PutApiV4GroupsIdIntegrationsAsanaWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsAsanaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsAsanaResponse, error)

	// PutApiV4GroupsIdIntegrationsAssemblaWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsAssemblaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsAssemblaResponse, error)

	PutApiV4GroupsIdIntegrationsAssemblaWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsAssemblaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsAssemblaResponse, error)

	// PutApiV4GroupsIdIntegrationsBambooWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsBambooWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsBambooResponse, error)

	PutApiV4GroupsIdIntegrationsBambooWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsBambooJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsBambooResponse, error)

	// PutApiV4GroupsIdIntegrationsBugzillaWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsBugzillaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsBugzillaResponse, error)

	PutApiV4GroupsIdIntegrationsBugzillaWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsBugzillaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsBugzillaResponse, error)

	// PutApiV4GroupsIdIntegrationsBuildkiteWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsBuildkiteWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsBuildkiteResponse, error)

	PutApiV4GroupsIdIntegrationsBuildkiteWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsBuildkiteJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsBuildkiteResponse, error)

	// PutApiV4GroupsIdIntegrationsCampfireWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsCampfireWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsCampfireResponse, error)

	PutApiV4GroupsIdIntegrationsCampfireWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsCampfireJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsCampfireResponse, error)

	// PutApiV4GroupsIdIntegrationsClickupWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsClickupWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsClickupResponse, error)

	PutApiV4GroupsIdIntegrationsClickupWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsClickupJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsClickupResponse, error)

	// PutApiV4GroupsIdIntegrationsConfluenceWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsConfluenceWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsConfluenceResponse, error)

	PutApiV4GroupsIdIntegrationsConfluenceWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsConfluenceJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsConfluenceResponse, error)

	// PutApiV4GroupsIdIntegrationsCustomIssueTrackerWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsCustomIssueTrackerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsCustomIssueTrackerResponse, error)

	PutApiV4GroupsIdIntegrationsCustomIssueTrackerWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsCustomIssueTrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsCustomIssueTrackerResponse, error)

	// PutApiV4GroupsIdIntegrationsDatadogWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsDatadogWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsDatadogResponse, error)

	PutApiV4GroupsIdIntegrationsDatadogWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsDatadogJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsDatadogResponse, error)

	// PutApiV4GroupsIdIntegrationsDiffblueCoverWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsDiffblueCoverWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsDiffblueCoverResponse, error)

	PutApiV4GroupsIdIntegrationsDiffblueCoverWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsDiffblueCoverJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsDiffblueCoverResponse, error)

	// PutApiV4GroupsIdIntegrationsDiscordWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsDiscordWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsDiscordResponse, error)

	PutApiV4GroupsIdIntegrationsDiscordWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsDiscordJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsDiscordResponse, error)

	// PutApiV4GroupsIdIntegrationsDroneCiWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsDroneCiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsDroneCiResponse, error)

	PutApiV4GroupsIdIntegrationsDroneCiWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsDroneCiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsDroneCiResponse, error)

	// PutApiV4GroupsIdIntegrationsEmailsOnPushWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsEmailsOnPushWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsEmailsOnPushResponse, error)

	PutApiV4GroupsIdIntegrationsEmailsOnPushWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsEmailsOnPushJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsEmailsOnPushResponse, error)

	// PutApiV4GroupsIdIntegrationsEwmWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsEwmWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsEwmResponse, error)

	PutApiV4GroupsIdIntegrationsEwmWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsEwmJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsEwmResponse, error)

	// PutApiV4GroupsIdIntegrationsExternalWikiWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsExternalWikiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsExternalWikiResponse, error)

	PutApiV4GroupsIdIntegrationsExternalWikiWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsExternalWikiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsExternalWikiResponse, error)

	// PutApiV4GroupsIdIntegrationsGitGuardianWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsGitGuardianWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGitGuardianResponse, error)

	PutApiV4GroupsIdIntegrationsGitGuardianWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGitGuardianJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGitGuardianResponse, error)

	// PutApiV4GroupsIdIntegrationsGithubWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsGithubWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGithubResponse, error)

	PutApiV4GroupsIdIntegrationsGithubWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGithubResponse, error)

	// PutApiV4GroupsIdIntegrationsGitlabSlackApplicationWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsGitlabSlackApplicationWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGitlabSlackApplicationResponse, error)

	PutApiV4GroupsIdIntegrationsGitlabSlackApplicationWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGitlabSlackApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGitlabSlackApplicationResponse, error)

	// PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryResponse, error)

	PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGoogleCloudPlatformArtifactRegistryResponse, error)

	// PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationResponse, error)

	PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationResponse, error)

	// PutApiV4GroupsIdIntegrationsGooglePlayWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsGooglePlayWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGooglePlayResponse, error)

	PutApiV4GroupsIdIntegrationsGooglePlayWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsGooglePlayJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsGooglePlayResponse, error)

	// PutApiV4GroupsIdIntegrationsHangoutsChatWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsHangoutsChatWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsHangoutsChatResponse, error)

	PutApiV4GroupsIdIntegrationsHangoutsChatWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsHangoutsChatJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsHangoutsChatResponse, error)

	// PutApiV4GroupsIdIntegrationsHarborWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsHarborWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsHarborResponse, error)

	PutApiV4GroupsIdIntegrationsHarborWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsHarborJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsHarborResponse, error)

	// PutApiV4GroupsIdIntegrationsIrkerWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsIrkerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsIrkerResponse, error)

	PutApiV4GroupsIdIntegrationsIrkerWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsIrkerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsIrkerResponse, error)

	// PutApiV4GroupsIdIntegrationsJenkinsWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsJenkinsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsJenkinsResponse, error)

	PutApiV4GroupsIdIntegrationsJenkinsWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsJenkinsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsJenkinsResponse, error)

	// PutApiV4GroupsIdIntegrationsJiraWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsJiraWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsJiraResponse, error)

	PutApiV4GroupsIdIntegrationsJiraWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsJiraJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsJiraResponse, error)

	// PutApiV4GroupsIdIntegrationsJiraCloudAppWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsJiraCloudAppWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsJiraCloudAppResponse, error)

	PutApiV4GroupsIdIntegrationsJiraCloudAppWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsJiraCloudAppJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsJiraCloudAppResponse, error)

	// PutApiV4GroupsIdIntegrationsMatrixWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsMatrixWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMatrixResponse, error)

	PutApiV4GroupsIdIntegrationsMatrixWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMatrixJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMatrixResponse, error)

	// PutApiV4GroupsIdIntegrationsMattermostWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsMattermostWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMattermostResponse, error)

	PutApiV4GroupsIdIntegrationsMattermostWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMattermostJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMattermostResponse, error)

	// PutApiV4GroupsIdIntegrationsMattermostSlashCommandsWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsMattermostSlashCommandsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMattermostSlashCommandsResponse, error)

	PutApiV4GroupsIdIntegrationsMattermostSlashCommandsWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMattermostSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMattermostSlashCommandsResponse, error)

	// PutApiV4GroupsIdIntegrationsMicrosoftTeamsWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsMicrosoftTeamsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMicrosoftTeamsResponse, error)

	PutApiV4GroupsIdIntegrationsMicrosoftTeamsWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMicrosoftTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMicrosoftTeamsResponse, error)

	// PutApiV4GroupsIdIntegrationsMockCiWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsMockCiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMockCiResponse, error)

	PutApiV4GroupsIdIntegrationsMockCiWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMockCiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMockCiResponse, error)

	// PutApiV4GroupsIdIntegrationsMockMonitoringWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsMockMonitoringWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMockMonitoringResponse, error)

	PutApiV4GroupsIdIntegrationsMockMonitoringWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsMockMonitoringJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsMockMonitoringResponse, error)

	// PutApiV4GroupsIdIntegrationsPackagistWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsPackagistWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPackagistResponse, error)

	PutApiV4GroupsIdIntegrationsPackagistWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPackagistJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPackagistResponse, error)

	// PutApiV4GroupsIdIntegrationsPhorgeWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsPhorgeWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPhorgeResponse, error)

	PutApiV4GroupsIdIntegrationsPhorgeWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPhorgeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPhorgeResponse, error)

	// PutApiV4GroupsIdIntegrationsPipelinesEmailWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsPipelinesEmailWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPipelinesEmailResponse, error)

	PutApiV4GroupsIdIntegrationsPipelinesEmailWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPipelinesEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPipelinesEmailResponse, error)

	// PutApiV4GroupsIdIntegrationsPivotaltrackerWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsPivotaltrackerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPivotaltrackerResponse, error)

	PutApiV4GroupsIdIntegrationsPivotaltrackerWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPivotaltrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPivotaltrackerResponse, error)

	// PutApiV4GroupsIdIntegrationsPumbleWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsPumbleWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPumbleResponse, error)

	PutApiV4GroupsIdIntegrationsPumbleWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPumbleJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPumbleResponse, error)

	// PutApiV4GroupsIdIntegrationsPushoverWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsPushoverWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPushoverResponse, error)

	PutApiV4GroupsIdIntegrationsPushoverWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsPushoverJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsPushoverResponse, error)

	// PutApiV4GroupsIdIntegrationsRedmineWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsRedmineWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsRedmineResponse, error)

	PutApiV4GroupsIdIntegrationsRedmineWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsRedmineJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsRedmineResponse, error)

	// PutApiV4GroupsIdIntegrationsSlackWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsSlackWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsSlackResponse, error)

	PutApiV4GroupsIdIntegrationsSlackWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsSlackJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsSlackResponse, error)

	// PutApiV4GroupsIdIntegrationsSlackSlashCommandsWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsSlackSlashCommandsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsSlackSlashCommandsResponse, error)

	PutApiV4GroupsIdIntegrationsSlackSlashCommandsWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsSlackSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsSlackSlashCommandsResponse, error)

	// PutApiV4GroupsIdIntegrationsSquashTmWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsSquashTmWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsSquashTmResponse, error)

	PutApiV4GroupsIdIntegrationsSquashTmWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsSquashTmJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsSquashTmResponse, error)

	// PutApiV4GroupsIdIntegrationsTeamcityWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsTeamcityWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsTeamcityResponse, error)

	PutApiV4GroupsIdIntegrationsTeamcityWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsTeamcityJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsTeamcityResponse, error)

	// PutApiV4GroupsIdIntegrationsTelegramWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsTelegramWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsTelegramResponse, error)

	PutApiV4GroupsIdIntegrationsTelegramWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsTelegramJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsTelegramResponse, error)

	// PutApiV4GroupsIdIntegrationsUnifyCircuitWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsUnifyCircuitWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsUnifyCircuitResponse, error)

	PutApiV4GroupsIdIntegrationsUnifyCircuitWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsUnifyCircuitJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsUnifyCircuitResponse, error)

	// PutApiV4GroupsIdIntegrationsWebexTeamsWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsWebexTeamsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsWebexTeamsResponse, error)

	PutApiV4GroupsIdIntegrationsWebexTeamsWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsWebexTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsWebexTeamsResponse, error)

	// PutApiV4GroupsIdIntegrationsYoutrackWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsYoutrackWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsYoutrackResponse, error)

	PutApiV4GroupsIdIntegrationsYoutrackWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsYoutrackJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsYoutrackResponse, error)

	// PutApiV4GroupsIdIntegrationsZentaoWithBodyWithResponse request with any body
	PutApiV4GroupsIdIntegrationsZentaoWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsZentaoResponse, error)

	PutApiV4GroupsIdIntegrationsZentaoWithResponse(ctx context.Context, id int32, body PutApiV4GroupsIdIntegrationsZentaoJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdIntegrationsZentaoResponse, error)

	// DeleteApiV4GroupsIdIntegrationsSlugWithResponse request
	DeleteApiV4GroupsIdIntegrationsSlugWithResponse(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdIntegrationsSlugResponse, error)

	// GetApiV4GroupsIdIntegrationsSlugWithResponse request
	GetApiV4GroupsIdIntegrationsSlugWithResponse(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdIntegrationsSlugResponse, error)

	// GetApiV4GroupsIdInvitationsWithResponse request
	GetApiV4GroupsIdInvitationsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdInvitationsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdInvitationsResponse, error)

	// PostApiV4GroupsIdInvitationsWithBodyWithResponse request with any body
	PostApiV4GroupsIdInvitationsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdInvitationsResponse, error)

	PostApiV4GroupsIdInvitationsWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdInvitationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdInvitationsResponse, error)

	// DeleteApiV4GroupsIdInvitationsEmailWithResponse request
	DeleteApiV4GroupsIdInvitationsEmailWithResponse(ctx context.Context, id string, email string, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdInvitationsEmailResponse, error)

	// PutApiV4GroupsIdInvitationsEmailWithBodyWithResponse request with any body
	PutApiV4GroupsIdInvitationsEmailWithBodyWithResponse(ctx context.Context, id string, email string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdInvitationsEmailResponse, error)

	PutApiV4GroupsIdInvitationsEmailWithResponse(ctx context.Context, id string, email string, body PutApiV4GroupsIdInvitationsEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdInvitationsEmailResponse, error)

	// GetApiV4GroupsIdInvitedGroupsWithResponse request
	GetApiV4GroupsIdInvitedGroupsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdInvitedGroupsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdInvitedGroupsResponse, error)

	// GetApiV4GroupsIdIssuesWithResponse request
	GetApiV4GroupsIdIssuesWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdIssuesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdIssuesResponse, error)

	// PostApiV4GroupsIdLdapSyncWithResponse request
	PostApiV4GroupsIdLdapSyncWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdLdapSyncResponse, error)

	// GetApiV4GroupsIdMembersWithResponse request
	GetApiV4GroupsIdMembersWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdMembersParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdMembersResponse, error)

	// PostApiV4GroupsIdMembersWithBodyWithResponse request with any body
	PostApiV4GroupsIdMembersWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdMembersResponse, error)

	PostApiV4GroupsIdMembersWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdMembersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdMembersResponse, error)

	// GetApiV4GroupsIdMembersAllWithResponse request
	GetApiV4GroupsIdMembersAllWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdMembersAllParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdMembersAllResponse, error)

	// GetApiV4GroupsIdMembersAllUserIdWithResponse request
	GetApiV4GroupsIdMembersAllUserIdWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdMembersAllUserIdResponse, error)

	// PostApiV4GroupsIdMembersApproveAllWithResponse request
	PostApiV4GroupsIdMembersApproveAllWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdMembersApproveAllResponse, error)

	// PutApiV4GroupsIdMembersMemberIdApproveWithResponse request
	PutApiV4GroupsIdMembersMemberIdApproveWithResponse(ctx context.Context, id string, memberId int32, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdMembersMemberIdApproveResponse, error)

	// DeleteApiV4GroupsIdMembersUserIdWithResponse request
	DeleteApiV4GroupsIdMembersUserIdWithResponse(ctx context.Context, id string, userId int32, params *DeleteApiV4GroupsIdMembersUserIdParams, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdMembersUserIdResponse, error)

	// GetApiV4GroupsIdMembersUserIdWithResponse request
	GetApiV4GroupsIdMembersUserIdWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdMembersUserIdResponse, error)

	// PutApiV4GroupsIdMembersUserIdWithBodyWithResponse request with any body
	PutApiV4GroupsIdMembersUserIdWithBodyWithResponse(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdMembersUserIdResponse, error)

	PutApiV4GroupsIdMembersUserIdWithResponse(ctx context.Context, id string, userId int32, body PutApiV4GroupsIdMembersUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdMembersUserIdResponse, error)

	// DeleteApiV4GroupsIdMembersUserIdOverrideWithResponse request
	DeleteApiV4GroupsIdMembersUserIdOverrideWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdMembersUserIdOverrideResponse, error)

	// PostApiV4GroupsIdMembersUserIdOverrideWithResponse request
	PostApiV4GroupsIdMembersUserIdOverrideWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdMembersUserIdOverrideResponse, error)

	// PutApiV4GroupsIdMembersUserIdStateWithBodyWithResponse request with any body
	PutApiV4GroupsIdMembersUserIdStateWithBodyWithResponse(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdMembersUserIdStateResponse, error)

	PutApiV4GroupsIdMembersUserIdStateWithResponse(ctx context.Context, id string, userId int32, body PutApiV4GroupsIdMembersUserIdStateJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdMembersUserIdStateResponse, error)

	// GetApiV4GroupsIdMergeRequestsWithResponse request
	GetApiV4GroupsIdMergeRequestsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdMergeRequestsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdMergeRequestsResponse, error)

	// GetApiV4GroupsIdPackagesWithResponse request
	GetApiV4GroupsIdPackagesWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdPackagesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPackagesResponse, error)

	// GetApiV4GroupsIdPendingMembersWithResponse request
	GetApiV4GroupsIdPendingMembersWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdPendingMembersParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPendingMembersResponse, error)

	// GetApiV4GroupsIdPlaceholderReassignmentsWithResponse request
	GetApiV4GroupsIdPlaceholderReassignmentsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdPlaceholderReassignmentsResponse, error)

	// PostApiV4GroupsIdPlaceholderReassignmentsWithBodyWithResponse request with any body
	PostApiV4GroupsIdPlaceholderReassignmentsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdPlaceholderReassignmentsResponse, error)

	PostApiV4GroupsIdPlaceholderReassignmentsWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdPlaceholderReassignmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdPlaceholderReassignmentsResponse, error)

	// PostApiV4GroupsIdPlaceholderReassignmentsAuthorizeWithResponse request
	PostApiV4GroupsIdPlaceholderReassignmentsAuthorizeWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdPlaceholderReassignmentsAuthorizeResponse, error)

	// GetApiV4GroupsIdProjectsWithResponse request
	GetApiV4GroupsIdProjectsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdProjectsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdProjectsResponse, error)

	// GetApiV4GroupsIdProjectsSharedWithResponse request
	GetApiV4GroupsIdProjectsSharedWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdProjectsSharedParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdProjectsSharedResponse, error)

	// PostApiV4GroupsIdProjectsProjectIdWithResponse request
	PostApiV4GroupsIdProjectsProjectIdWithResponse(ctx context.Context, id string, projectId string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdProjectsProjectIdResponse, error)

	// GetApiV4GroupsIdProvisionedUsersWithResponse request
	GetApiV4GroupsIdProvisionedUsersWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdProvisionedUsersParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdProvisionedUsersResponse, error)

	// GetApiV4GroupsIdRegistryRepositoriesWithResponse request
	GetApiV4GroupsIdRegistryRepositoriesWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdRegistryRepositoriesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdRegistryRepositoriesResponse, error)

	// GetApiV4GroupsIdReleasesWithResponse request
	GetApiV4GroupsIdReleasesWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdReleasesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdReleasesResponse, error)

	// PostApiV4GroupsIdRestoreWithResponse request
	PostApiV4GroupsIdRestoreWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdRestoreResponse, error)

	// GetApiV4GroupsIdRunnersWithResponse request
	GetApiV4GroupsIdRunnersWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdRunnersParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdRunnersResponse, error)

	// PostApiV4GroupsIdRunnersResetRegistrationTokenWithResponse request
	PostApiV4GroupsIdRunnersResetRegistrationTokenWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdRunnersResetRegistrationTokenResponse, error)

	// GetApiV4GroupsIdSamlUsersWithResponse request
	GetApiV4GroupsIdSamlUsersWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdSamlUsersParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdSamlUsersResponse, error)

	// PostApiV4GroupsIdShareWithBodyWithResponse request with any body
	PostApiV4GroupsIdShareWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdShareResponse, error)

	PostApiV4GroupsIdShareWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdShareJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdShareResponse, error)

	// DeleteApiV4GroupsIdShareGroupIdWithResponse request
	DeleteApiV4GroupsIdShareGroupIdWithResponse(ctx context.Context, id string, groupId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdShareGroupIdResponse, error)

	// GetApiV4GroupsIdSshCertificatesWithResponse request
	GetApiV4GroupsIdSshCertificatesWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdSshCertificatesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdSshCertificatesResponse, error)

	// PostApiV4GroupsIdSshCertificatesWithBodyWithResponse request with any body
	PostApiV4GroupsIdSshCertificatesWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdSshCertificatesResponse, error)

	PostApiV4GroupsIdSshCertificatesWithResponse(ctx context.Context, id int32, body PostApiV4GroupsIdSshCertificatesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdSshCertificatesResponse, error)

	// DeleteApiV4GroupsIdSshCertificatesSshCertificatesIdWithResponse request
	DeleteApiV4GroupsIdSshCertificatesSshCertificatesIdWithResponse(ctx context.Context, id int32, sshCertificatesId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdSshCertificatesSshCertificatesIdResponse, error)

	// GetApiV4GroupsIdSubgroupsWithResponse request
	GetApiV4GroupsIdSubgroupsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdSubgroupsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdSubgroupsResponse, error)

	// PostApiV4GroupsIdTokensRevokeWithBodyWithResponse request with any body
	PostApiV4GroupsIdTokensRevokeWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdTokensRevokeResponse, error)

	PostApiV4GroupsIdTokensRevokeWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdTokensRevokeJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdTokensRevokeResponse, error)

	// PostApiV4GroupsIdTransferWithBodyWithResponse request with any body
	PostApiV4GroupsIdTransferWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdTransferResponse, error)

	PostApiV4GroupsIdTransferWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdTransferJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdTransferResponse, error)

	// GetApiV4GroupsIdTransferLocationsWithResponse request
	GetApiV4GroupsIdTransferLocationsWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdTransferLocationsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdTransferLocationsResponse, error)

	// PostApiV4GroupsIdUnarchiveWithResponse request
	PostApiV4GroupsIdUnarchiveWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdUnarchiveResponse, error)

	// GetApiV4GroupsIdUploadsWithResponse request
	GetApiV4GroupsIdUploadsWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdUploadsParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdUploadsResponse, error)

	// DeleteApiV4GroupsIdUploadsSecretFilenameWithResponse request
	DeleteApiV4GroupsIdUploadsSecretFilenameWithResponse(ctx context.Context, id int32, secret string, filename string, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdUploadsSecretFilenameResponse, error)

	// GetApiV4GroupsIdUploadsSecretFilenameWithResponse request
	GetApiV4GroupsIdUploadsSecretFilenameWithResponse(ctx context.Context, id int32, secret string, filename string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdUploadsSecretFilenameResponse, error)

	// DeleteApiV4GroupsIdUploadsUploadIdWithResponse request
	DeleteApiV4GroupsIdUploadsUploadIdWithResponse(ctx context.Context, id int32, uploadId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdUploadsUploadIdResponse, error)

	// GetApiV4GroupsIdUploadsUploadIdWithResponse request
	GetApiV4GroupsIdUploadsUploadIdWithResponse(ctx context.Context, id int32, uploadId int32, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdUploadsUploadIdResponse, error)

	// GetApiV4GroupsIdUsersWithResponse request
	GetApiV4GroupsIdUsersWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdUsersParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdUsersResponse, error)

	// GetApiV4GroupsIdVariablesWithResponse request
	GetApiV4GroupsIdVariablesWithResponse(ctx context.Context, id string, params *GetApiV4GroupsIdVariablesParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdVariablesResponse, error)

	// PostApiV4GroupsIdVariablesWithBodyWithResponse request with any body
	PostApiV4GroupsIdVariablesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdVariablesResponse, error)

	PostApiV4GroupsIdVariablesWithResponse(ctx context.Context, id string, body PostApiV4GroupsIdVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdVariablesResponse, error)

	// DeleteApiV4GroupsIdVariablesKeyWithResponse request
	DeleteApiV4GroupsIdVariablesKeyWithResponse(ctx context.Context, id string, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdVariablesKeyResponse, error)

	// GetApiV4GroupsIdVariablesKeyWithResponse request
	GetApiV4GroupsIdVariablesKeyWithResponse(ctx context.Context, id string, key string, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdVariablesKeyResponse, error)

	// PutApiV4GroupsIdVariablesKeyWithBodyWithResponse request with any body
	PutApiV4GroupsIdVariablesKeyWithBodyWithResponse(ctx context.Context, id string, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdVariablesKeyResponse, error)

	PutApiV4GroupsIdVariablesKeyWithResponse(ctx context.Context, id string, key string, body PutApiV4GroupsIdVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdVariablesKeyResponse, error)

	// GetApiV4GroupsIdWikisWithResponse request
	GetApiV4GroupsIdWikisWithResponse(ctx context.Context, id int32, params *GetApiV4GroupsIdWikisParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdWikisResponse, error)

	// PostApiV4GroupsIdWikisWithBodyWithResponse request with any body
	PostApiV4GroupsIdWikisWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdWikisResponse, error)

	PostApiV4GroupsIdWikisWithResponse(ctx context.Context, id int32, body PostApiV4GroupsIdWikisJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdWikisResponse, error)

	// PostApiV4GroupsIdWikisAttachmentsWithBodyWithResponse request with any body
	PostApiV4GroupsIdWikisAttachmentsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdWikisAttachmentsResponse, error)

	PostApiV4GroupsIdWikisAttachmentsWithResponse(ctx context.Context, id int32, body PostApiV4GroupsIdWikisAttachmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GroupsIdWikisAttachmentsResponse, error)

	// DeleteApiV4GroupsIdWikisSlugWithResponse request
	DeleteApiV4GroupsIdWikisSlugWithResponse(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*DeleteApiV4GroupsIdWikisSlugResponse, error)

	// GetApiV4GroupsIdWikisSlugWithResponse request
	GetApiV4GroupsIdWikisSlugWithResponse(ctx context.Context, id int32, slug string, params *GetApiV4GroupsIdWikisSlugParams, reqEditors ...RequestEditorFn) (*GetApiV4GroupsIdWikisSlugResponse, error)

	// PutApiV4GroupsIdWikisSlugWithBodyWithResponse request with any body
	PutApiV4GroupsIdWikisSlugWithBodyWithResponse(ctx context.Context, id int32, slug int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdWikisSlugResponse, error)

	PutApiV4GroupsIdWikisSlugWithResponse(ctx context.Context, id int32, slug int32, body PutApiV4GroupsIdWikisSlugJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4GroupsIdWikisSlugResponse, error)

	// GetApiV4HooksWithResponse request
	GetApiV4HooksWithResponse(ctx context.Context, params *GetApiV4HooksParams, reqEditors ...RequestEditorFn) (*GetApiV4HooksResponse, error)

	// PostApiV4HooksWithBodyWithResponse request with any body
	PostApiV4HooksWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4HooksResponse, error)

	PostApiV4HooksWithResponse(ctx context.Context, body PostApiV4HooksJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4HooksResponse, error)

	// DeleteApiV4HooksHookIdWithResponse request
	DeleteApiV4HooksHookIdWithResponse(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4HooksHookIdResponse, error)

	// GetApiV4HooksHookIdWithResponse request
	GetApiV4HooksHookIdWithResponse(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*GetApiV4HooksHookIdResponse, error)

	// PostApiV4HooksHookIdWithResponse request
	PostApiV4HooksHookIdWithResponse(ctx context.Context, hookId int32, reqEditors ...RequestEditorFn) (*PostApiV4HooksHookIdResponse, error)

	// PutApiV4HooksHookIdWithBodyWithResponse request with any body
	PutApiV4HooksHookIdWithBodyWithResponse(ctx context.Context, hookId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdResponse, error)

	PutApiV4HooksHookIdWithResponse(ctx context.Context, hookId int32, body PutApiV4HooksHookIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdResponse, error)

	// DeleteApiV4HooksHookIdCustomHeadersKeyWithResponse request
	DeleteApiV4HooksHookIdCustomHeadersKeyWithResponse(ctx context.Context, hookId int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4HooksHookIdCustomHeadersKeyResponse, error)

	// PutApiV4HooksHookIdCustomHeadersKeyWithBodyWithResponse request with any body
	PutApiV4HooksHookIdCustomHeadersKeyWithBodyWithResponse(ctx context.Context, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdCustomHeadersKeyResponse, error)

	PutApiV4HooksHookIdCustomHeadersKeyWithResponse(ctx context.Context, hookId int32, key string, body PutApiV4HooksHookIdCustomHeadersKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdCustomHeadersKeyResponse, error)

	// DeleteApiV4HooksHookIdUrlVariablesKeyWithResponse request
	DeleteApiV4HooksHookIdUrlVariablesKeyWithResponse(ctx context.Context, hookId int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4HooksHookIdUrlVariablesKeyResponse, error)

	// PutApiV4HooksHookIdUrlVariablesKeyWithBodyWithResponse request with any body
	PutApiV4HooksHookIdUrlVariablesKeyWithBodyWithResponse(ctx context.Context, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdUrlVariablesKeyResponse, error)

	PutApiV4HooksHookIdUrlVariablesKeyWithResponse(ctx context.Context, hookId int32, key string, body PutApiV4HooksHookIdUrlVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4HooksHookIdUrlVariablesKeyResponse, error)

	// PostApiV4ImportBitbucketWithBodyWithResponse request with any body
	PostApiV4ImportBitbucketWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportBitbucketResponse, error)

	PostApiV4ImportBitbucketWithResponse(ctx context.Context, body PostApiV4ImportBitbucketJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportBitbucketResponse, error)

	// PostApiV4ImportBitbucketServerWithBodyWithResponse request with any body
	PostApiV4ImportBitbucketServerWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportBitbucketServerResponse, error)

	PostApiV4ImportBitbucketServerWithResponse(ctx context.Context, body PostApiV4ImportBitbucketServerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportBitbucketServerResponse, error)

	// PostApiV4ImportGithubWithBodyWithResponse request with any body
	PostApiV4ImportGithubWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubResponse, error)

	PostApiV4ImportGithubWithResponse(ctx context.Context, body PostApiV4ImportGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubResponse, error)

	// PostApiV4ImportGithubCancelWithBodyWithResponse request with any body
	PostApiV4ImportGithubCancelWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubCancelResponse, error)

	PostApiV4ImportGithubCancelWithResponse(ctx context.Context, body PostApiV4ImportGithubCancelJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubCancelResponse, error)

	// PostApiV4ImportGithubGistsWithBodyWithResponse request with any body
	PostApiV4ImportGithubGistsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubGistsResponse, error)

	PostApiV4ImportGithubGistsWithResponse(ctx context.Context, body PostApiV4ImportGithubGistsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ImportGithubGistsResponse, error)

	// PostApiV4IntegrationsJiraConnectSubscriptionsWithBodyWithResponse request with any body
	PostApiV4IntegrationsJiraConnectSubscriptionsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsJiraConnectSubscriptionsResponse, error)

	PostApiV4IntegrationsJiraConnectSubscriptionsWithResponse(ctx context.Context, body PostApiV4IntegrationsJiraConnectSubscriptionsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsJiraConnectSubscriptionsResponse, error)

	// PostApiV4IntegrationsSlackEventsWithBodyWithResponse request with any body
	PostApiV4IntegrationsSlackEventsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsSlackEventsResponse, error)

	PostApiV4IntegrationsSlackEventsWithResponse(ctx context.Context, body PostApiV4IntegrationsSlackEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsSlackEventsResponse, error)

	// PostApiV4IntegrationsSlackInteractionsWithResponse request
	PostApiV4IntegrationsSlackInteractionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsSlackInteractionsResponse, error)

	// PostApiV4IntegrationsSlackOptionsWithResponse request
	PostApiV4IntegrationsSlackOptionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4IntegrationsSlackOptionsResponse, error)

	// GetApiV4IssuesWithResponse request
	GetApiV4IssuesWithResponse(ctx context.Context, params *GetApiV4IssuesParams, reqEditors ...RequestEditorFn) (*GetApiV4IssuesResponse, error)

	// GetApiV4IssuesIdWithResponse request
	GetApiV4IssuesIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetApiV4IssuesIdResponse, error)

	// GetApiV4JobWithResponse request
	GetApiV4JobWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4JobResponse, error)

	// GetApiV4JobAllowedAgentsWithResponse request
	GetApiV4JobAllowedAgentsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4JobAllowedAgentsResponse, error)

	// PostApiV4JobsRequestWithBodyWithResponse request with any body
	PostApiV4JobsRequestWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4JobsRequestResponse, error)

	PostApiV4JobsRequestWithResponse(ctx context.Context, body PostApiV4JobsRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4JobsRequestResponse, error)

	// PutApiV4JobsIdWithBodyWithResponse request with any body
	PutApiV4JobsIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4JobsIdResponse, error)

	PutApiV4JobsIdWithResponse(ctx context.Context, id int32, body PutApiV4JobsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4JobsIdResponse, error)

	// GetApiV4JobsIdArtifactsWithResponse request
	GetApiV4JobsIdArtifactsWithResponse(ctx context.Context, id int32, params *GetApiV4JobsIdArtifactsParams, reqEditors ...RequestEditorFn) (*GetApiV4JobsIdArtifactsResponse, error)

	// PostApiV4JobsIdArtifactsWithBodyWithResponse request with any body
	PostApiV4JobsIdArtifactsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4JobsIdArtifactsResponse, error)

	PostApiV4JobsIdArtifactsWithResponse(ctx context.Context, id int32, body PostApiV4JobsIdArtifactsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4JobsIdArtifactsResponse, error)

	// PostApiV4JobsIdArtifactsAuthorizeWithBodyWithResponse request with any body
	PostApiV4JobsIdArtifactsAuthorizeWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4JobsIdArtifactsAuthorizeResponse, error)

	PostApiV4JobsIdArtifactsAuthorizeWithResponse(ctx context.Context, id int32, body PostApiV4JobsIdArtifactsAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4JobsIdArtifactsAuthorizeResponse, error)

	// PatchApiV4JobsIdTraceWithBodyWithResponse request with any body
	PatchApiV4JobsIdTraceWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchApiV4JobsIdTraceResponse, error)

	PatchApiV4JobsIdTraceWithResponse(ctx context.Context, id int32, body PatchApiV4JobsIdTraceJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchApiV4JobsIdTraceResponse, error)

	// GetApiV4KeysWithResponse request
	GetApiV4KeysWithResponse(ctx context.Context, params *GetApiV4KeysParams, reqEditors ...RequestEditorFn) (*GetApiV4KeysResponse, error)

	// GetApiV4KeysIdWithResponse request
	GetApiV4KeysIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4KeysIdResponse, error)

	// PostApiV4MarkdownWithBodyWithResponse request with any body
	PostApiV4MarkdownWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4MarkdownResponse, error)

	PostApiV4MarkdownWithResponse(ctx context.Context, body PostApiV4MarkdownJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4MarkdownResponse, error)

	// GetApiV4MergeRequestsWithResponse request
	GetApiV4MergeRequestsWithResponse(ctx context.Context, params *GetApiV4MergeRequestsParams, reqEditors ...RequestEditorFn) (*GetApiV4MergeRequestsResponse, error)

	// GetApiV4MetadataWithResponse request
	GetApiV4MetadataWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4MetadataResponse, error)

	// GetApiV4NamespacesWithResponse request
	GetApiV4NamespacesWithResponse(ctx context.Context, params *GetApiV4NamespacesParams, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesResponse, error)

	// GetApiV4NamespacesStorageLimitExclusionsWithResponse request
	GetApiV4NamespacesStorageLimitExclusionsWithResponse(ctx context.Context, params *GetApiV4NamespacesStorageLimitExclusionsParams, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesStorageLimitExclusionsResponse, error)

	// GetApiV4NamespacesIdWithResponse request
	GetApiV4NamespacesIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesIdResponse, error)

	// PutApiV4NamespacesIdWithBodyWithResponse request with any body
	PutApiV4NamespacesIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4NamespacesIdResponse, error)

	PutApiV4NamespacesIdWithResponse(ctx context.Context, id int32, body PutApiV4NamespacesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4NamespacesIdResponse, error)

	// GetApiV4NamespacesIdExistsWithResponse request
	GetApiV4NamespacesIdExistsWithResponse(ctx context.Context, id string, params *GetApiV4NamespacesIdExistsParams, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesIdExistsResponse, error)

	// GetApiV4NamespacesIdGitlabSubscriptionWithResponse request
	GetApiV4NamespacesIdGitlabSubscriptionWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4NamespacesIdGitlabSubscriptionResponse, error)

	// DeleteApiV4NamespacesIdStorageLimitExclusionWithResponse request
	DeleteApiV4NamespacesIdStorageLimitExclusionWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4NamespacesIdStorageLimitExclusionResponse, error)

	// PostApiV4NamespacesIdStorageLimitExclusionWithBodyWithResponse request with any body
	PostApiV4NamespacesIdStorageLimitExclusionWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4NamespacesIdStorageLimitExclusionResponse, error)

	PostApiV4NamespacesIdStorageLimitExclusionWithResponse(ctx context.Context, id int32, body PostApiV4NamespacesIdStorageLimitExclusionJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4NamespacesIdStorageLimitExclusionResponse, error)

	// PostApiV4OrganizationsWithBodyWithResponse request with any body
	PostApiV4OrganizationsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4OrganizationsResponse, error)

	PostApiV4OrganizationsWithResponse(ctx context.Context, body PostApiV4OrganizationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4OrganizationsResponse, error)

	// GetApiV4PackagesConanV1ConansSearchWithResponse request
	GetApiV4PackagesConanV1ConansSearchWithResponse(ctx context.Context, params *GetApiV4PackagesConanV1ConansSearchParams, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1ConansSearchResponse, error)

	// DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelWithResponse request
	DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*DeleteApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelResponse, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelWithResponse request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelResponse, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestWithResponse request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestResponse, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsWithResponse request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsResponse, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceWithResponse request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceResponse, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestWithResponse request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestResponse, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsWithResponse request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsResponse, error)

	// PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsWithResponse request
	PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsResponse, error)

	// GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchWithResponse request
	GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchResponse, error)

	// PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsWithResponse request
	PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*PostApiV4PackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsResponse, error)

	// GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithResponse request
	GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameResponse, error)

	// PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithBodyWithResponse request with any body
	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithBodyWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameResponse, error)

	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, body PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameResponse, error)

	// PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeWithResponse request
	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeResponse, error)

	// GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithResponse request
	GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameResponse, error)

	// PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithBodyWithResponse request with any body
	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithBodyWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameResponse, error)

	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, body PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameResponse, error)

	// PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeWithResponse request
	PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeWithResponse(ctx context.Context, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*PutApiV4PackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeResponse, error)

	// GetApiV4PackagesConanV1PingWithResponse request
	GetApiV4PackagesConanV1PingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1PingResponse, error)

	// GetApiV4PackagesConanV1UsersAuthenticateWithResponse request
	GetApiV4PackagesConanV1UsersAuthenticateWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1UsersAuthenticateResponse, error)

	// GetApiV4PackagesConanV1UsersCheckCredentialsWithResponse request
	GetApiV4PackagesConanV1UsersCheckCredentialsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4PackagesConanV1UsersCheckCredentialsResponse, error)

	// GetApiV4PackagesMavenpathFileNameWithResponse request
	GetApiV4PackagesMavenpathFileNameWithResponse(ctx context.Context, fileName string, params *GetApiV4PackagesMavenpathFileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4PackagesMavenpathFileNameResponse, error)

	// GetApiV4PackagesNpmpackageNameWithResponse request
	GetApiV4PackagesNpmpackageNameWithResponse(ctx context.Context, params *GetApiV4PackagesNpmpackageNameParams, reqEditors ...RequestEditorFn) (*GetApiV4PackagesNpmpackageNameResponse, error)

	// PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulkWithResponse request
	PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulkWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4PackagesNpmNpmV1SecurityAdvisoriesBulkResponse, error)

	// PostApiV4PackagesNpmNpmV1SecurityAuditsQuickWithResponse request
	PostApiV4PackagesNpmNpmV1SecurityAuditsQuickWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4PackagesNpmNpmV1SecurityAuditsQuickResponse, error)

	// GetApiV4PackagesNpmPackagepackageNameDistTagsWithResponse request
	GetApiV4PackagesNpmPackagepackageNameDistTagsWithResponse(ctx context.Context, params *GetApiV4PackagesNpmPackagepackageNameDistTagsParams, reqEditors ...RequestEditorFn) (*GetApiV4PackagesNpmPackagepackageNameDistTagsResponse, error)

	// DeleteApiV4PackagesNpmPackagepackageNameDistTagsTagWithResponse request
	DeleteApiV4PackagesNpmPackagepackageNameDistTagsTagWithResponse(ctx context.Context, tag string, params *DeleteApiV4PackagesNpmPackagepackageNameDistTagsTagParams, reqEditors ...RequestEditorFn) (*DeleteApiV4PackagesNpmPackagepackageNameDistTagsTagResponse, error)

	// PutApiV4PackagesNpmPackagepackageNameDistTagsTagWithBodyWithResponse request with any body
	PutApiV4PackagesNpmPackagepackageNameDistTagsTagWithBodyWithResponse(ctx context.Context, tag string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4PackagesNpmPackagepackageNameDistTagsTagResponse, error)

	PutApiV4PackagesNpmPackagepackageNameDistTagsTagWithResponse(ctx context.Context, tag string, body PutApiV4PackagesNpmPackagepackageNameDistTagsTagJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4PackagesNpmPackagepackageNameDistTagsTagResponse, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemWithResponse request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemWithResponse(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemResponse, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionWithResponse request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionWithResponse(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, params *GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionParams, reqEditors ...RequestEditorFn) (*GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionResponse, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownloadWithResponse request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownloadWithResponse(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, params *GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownloadParams, reqEditors ...RequestEditorFn) (*GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionDownloadResponse, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFileWithResponse request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFileWithResponse(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, params *GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFileParams, reqEditors ...RequestEditorFn) (*GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemmoduleVersionFileResponse, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadWithResponse request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadWithResponse(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemDownloadResponse, error)

	// GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsWithResponse request
	GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsWithResponse(ctx context.Context, moduleNamespace string, moduleName string, moduleSystem string, reqEditors ...RequestEditorFn) (*GetApiV4PackagesTerraformModulesV1ModuleNamespaceModuleNameModuleSystemVersionsResponse, error)

	// GetApiV4PagesDomainsWithResponse request
	GetApiV4PagesDomainsWithResponse(ctx context.Context, params *GetApiV4PagesDomainsParams, reqEditors ...RequestEditorFn) (*GetApiV4PagesDomainsResponse, error)

	// GetApiV4PersonalAccessTokensWithResponse request
	GetApiV4PersonalAccessTokensWithResponse(ctx context.Context, params *GetApiV4PersonalAccessTokensParams, reqEditors ...RequestEditorFn) (*GetApiV4PersonalAccessTokensResponse, error)

	// DeleteApiV4PersonalAccessTokensSelfWithResponse request
	DeleteApiV4PersonalAccessTokensSelfWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DeleteApiV4PersonalAccessTokensSelfResponse, error)

	// GetApiV4PersonalAccessTokensSelfWithResponse request
	GetApiV4PersonalAccessTokensSelfWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4PersonalAccessTokensSelfResponse, error)

	// GetApiV4PersonalAccessTokensSelfAssociationsWithResponse request
	GetApiV4PersonalAccessTokensSelfAssociationsWithResponse(ctx context.Context, params *GetApiV4PersonalAccessTokensSelfAssociationsParams, reqEditors ...RequestEditorFn) (*GetApiV4PersonalAccessTokensSelfAssociationsResponse, error)

	// PostApiV4PersonalAccessTokensSelfRotateWithBodyWithResponse request with any body
	PostApiV4PersonalAccessTokensSelfRotateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4PersonalAccessTokensSelfRotateResponse, error)

	PostApiV4PersonalAccessTokensSelfRotateWithResponse(ctx context.Context, body PostApiV4PersonalAccessTokensSelfRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4PersonalAccessTokensSelfRotateResponse, error)

	// DeleteApiV4PersonalAccessTokensIdWithResponse request
	DeleteApiV4PersonalAccessTokensIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4PersonalAccessTokensIdResponse, error)

	// GetApiV4PersonalAccessTokensIdWithResponse request
	GetApiV4PersonalAccessTokensIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4PersonalAccessTokensIdResponse, error)

	// PostApiV4PersonalAccessTokensIdRotateWithBodyWithResponse request with any body
	PostApiV4PersonalAccessTokensIdRotateWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4PersonalAccessTokensIdRotateResponse, error)

	PostApiV4PersonalAccessTokensIdRotateWithResponse(ctx context.Context, id int32, body PostApiV4PersonalAccessTokensIdRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4PersonalAccessTokensIdRotateResponse, error)

	// GetApiV4ProjectsWithResponse request
	GetApiV4ProjectsWithResponse(ctx context.Context, params *GetApiV4ProjectsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsResponse, error)

	// PostApiV4ProjectsWithBodyWithResponse request with any body
	PostApiV4ProjectsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsResponse, error)

	PostApiV4ProjectsWithResponse(ctx context.Context, body PostApiV4ProjectsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsResponse, error)

	// PostApiV4ProjectsImportWithBodyWithResponse request with any body
	PostApiV4ProjectsImportWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsImportResponse, error)

	// PostApiV4ProjectsImportRelationWithBodyWithResponse request with any body
	PostApiV4ProjectsImportRelationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsImportRelationResponse, error)

	// PostApiV4ProjectsImportRelationAuthorizeWithResponse request
	PostApiV4ProjectsImportRelationAuthorizeWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsImportRelationAuthorizeResponse, error)

	// PostApiV4ProjectsImportAuthorizeWithResponse request
	PostApiV4ProjectsImportAuthorizeWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsImportAuthorizeResponse, error)

	// PostApiV4ProjectsRemoteImportWithBodyWithResponse request with any body
	PostApiV4ProjectsRemoteImportWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsRemoteImportResponse, error)

	// PostApiV4ProjectsRemoteImportS3WithBodyWithResponse request with any body
	PostApiV4ProjectsRemoteImportS3WithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsRemoteImportS3Response, error)

	// PostApiV4ProjectsUserUserIdWithBodyWithResponse request with any body
	PostApiV4ProjectsUserUserIdWithBodyWithResponse(ctx context.Context, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsUserUserIdResponse, error)

	PostApiV4ProjectsUserUserIdWithResponse(ctx context.Context, userId int32, body PostApiV4ProjectsUserUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsUserUserIdResponse, error)

	// DeleteApiV4ProjectsIdWithResponse request
	DeleteApiV4ProjectsIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdResponse, error)

	// GetApiV4ProjectsIdWithResponse request
	GetApiV4ProjectsIdWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdResponse, error)

	// PutApiV4ProjectsIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdResponse, error)

	PutApiV4ProjectsIdWithResponse(ctx context.Context, id string, body PutApiV4ProjectsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdResponse, error)

	// PostApiV4ProjectsIdRefRefTriggerPipelineWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRefRefTriggerPipelineWithBodyWithResponse(ctx context.Context, id string, ref string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRefRefTriggerPipelineResponse, error)

	PostApiV4ProjectsIdRefRefTriggerPipelineWithResponse(ctx context.Context, id string, ref string, body PostApiV4ProjectsIdRefRefTriggerPipelineJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRefRefTriggerPipelineResponse, error)

	// GetApiV4ProjectsIdAccessRequestsWithResponse request
	GetApiV4ProjectsIdAccessRequestsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdAccessRequestsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdAccessRequestsResponse, error)

	// PostApiV4ProjectsIdAccessRequestsWithResponse request
	PostApiV4ProjectsIdAccessRequestsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdAccessRequestsResponse, error)

	// DeleteApiV4ProjectsIdAccessRequestsUserIdWithResponse request
	DeleteApiV4ProjectsIdAccessRequestsUserIdWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdAccessRequestsUserIdResponse, error)

	// PutApiV4ProjectsIdAccessRequestsUserIdApproveWithBodyWithResponse request with any body
	PutApiV4ProjectsIdAccessRequestsUserIdApproveWithBodyWithResponse(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdAccessRequestsUserIdApproveResponse, error)

	PutApiV4ProjectsIdAccessRequestsUserIdApproveWithResponse(ctx context.Context, id string, userId int32, body PutApiV4ProjectsIdAccessRequestsUserIdApproveJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdAccessRequestsUserIdApproveResponse, error)

	// PostApiV4ProjectsIdAccessTokensSelfRotateWithBodyWithResponse request with any body
	PostApiV4ProjectsIdAccessTokensSelfRotateWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdAccessTokensSelfRotateResponse, error)

	PostApiV4ProjectsIdAccessTokensSelfRotateWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdAccessTokensSelfRotateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdAccessTokensSelfRotateResponse, error)

	// GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesWithResponse request
	GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesWithResponse(ctx context.Context, id string, alertIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesResponse, error)

	// PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesWithBodyWithResponse(ctx context.Context, id string, alertIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesResponse, error)

	// PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorizeWithResponse request
	PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorizeWithResponse(ctx context.Context, id string, alertIid int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesAuthorizeResponse, error)

	// DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdWithResponse request
	DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdWithResponse(ctx context.Context, id string, alertIid int32, metricImageId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdResponse, error)

	// PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdWithBodyWithResponse(ctx context.Context, id string, alertIid int32, metricImageId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdAlertManagementAlertsAlertIidMetricImagesMetricImageIdResponse, error)

	// PostApiV4ProjectsIdArchiveWithResponse request
	PostApiV4ProjectsIdArchiveWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdArchiveResponse, error)

	// DeleteApiV4ProjectsIdArtifactsWithResponse request
	DeleteApiV4ProjectsIdArtifactsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdArtifactsResponse, error)

	// GetApiV4ProjectsIdAuditEventsWithResponse request
	GetApiV4ProjectsIdAuditEventsWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdAuditEventsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdAuditEventsResponse, error)

	// GetApiV4ProjectsIdAuditEventsAuditEventIdWithResponse request
	GetApiV4ProjectsIdAuditEventsAuditEventIdWithResponse(ctx context.Context, id int32, auditEventId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdAuditEventsAuditEventIdResponse, error)

	// GetApiV4ProjectsIdAvatarWithResponse request
	GetApiV4ProjectsIdAvatarWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdAvatarResponse, error)

	// GetApiV4ProjectsIdBadgesWithResponse request
	GetApiV4ProjectsIdBadgesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdBadgesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdBadgesResponse, error)

	// PostApiV4ProjectsIdBadgesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdBadgesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdBadgesResponse, error)

	PostApiV4ProjectsIdBadgesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdBadgesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdBadgesResponse, error)

	// GetApiV4ProjectsIdBadgesRenderWithResponse request
	GetApiV4ProjectsIdBadgesRenderWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdBadgesRenderParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdBadgesRenderResponse, error)

	// DeleteApiV4ProjectsIdBadgesBadgeIdWithResponse request
	DeleteApiV4ProjectsIdBadgesBadgeIdWithResponse(ctx context.Context, id string, badgeId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdBadgesBadgeIdResponse, error)

	// GetApiV4ProjectsIdBadgesBadgeIdWithResponse request
	GetApiV4ProjectsIdBadgesBadgeIdWithResponse(ctx context.Context, id string, badgeId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdBadgesBadgeIdResponse, error)

	// PutApiV4ProjectsIdBadgesBadgeIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdBadgesBadgeIdWithBodyWithResponse(ctx context.Context, id string, badgeId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdBadgesBadgeIdResponse, error)

	PutApiV4ProjectsIdBadgesBadgeIdWithResponse(ctx context.Context, id string, badgeId int32, body PutApiV4ProjectsIdBadgesBadgeIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdBadgesBadgeIdResponse, error)

	// PostApiV4ProjectsIdCatalogPublishWithBodyWithResponse request with any body
	PostApiV4ProjectsIdCatalogPublishWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdCatalogPublishResponse, error)

	PostApiV4ProjectsIdCatalogPublishWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdCatalogPublishJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdCatalogPublishResponse, error)

	// GetApiV4ProjectsIdCiLintWithResponse request
	GetApiV4ProjectsIdCiLintWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdCiLintParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdCiLintResponse, error)

	// PostApiV4ProjectsIdCiLintWithBodyWithResponse request with any body
	PostApiV4ProjectsIdCiLintWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdCiLintResponse, error)

	PostApiV4ProjectsIdCiLintWithResponse(ctx context.Context, id int32, body PostApiV4ProjectsIdCiLintJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdCiLintResponse, error)

	// GetApiV4ProjectsIdClusterAgentsWithResponse request
	GetApiV4ProjectsIdClusterAgentsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdClusterAgentsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdClusterAgentsResponse, error)

	// PostApiV4ProjectsIdClusterAgentsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdClusterAgentsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdClusterAgentsResponse, error)

	PostApiV4ProjectsIdClusterAgentsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdClusterAgentsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdClusterAgentsResponse, error)

	// DeleteApiV4ProjectsIdClusterAgentsAgentIdWithResponse request
	DeleteApiV4ProjectsIdClusterAgentsAgentIdWithResponse(ctx context.Context, id string, agentId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdClusterAgentsAgentIdResponse, error)

	// GetApiV4ProjectsIdClusterAgentsAgentIdWithResponse request
	GetApiV4ProjectsIdClusterAgentsAgentIdWithResponse(ctx context.Context, id string, agentId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdClusterAgentsAgentIdResponse, error)

	// GetApiV4ProjectsIdClusterAgentsAgentIdTokensWithResponse request
	GetApiV4ProjectsIdClusterAgentsAgentIdTokensWithResponse(ctx context.Context, id string, agentId int32, params *GetApiV4ProjectsIdClusterAgentsAgentIdTokensParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdClusterAgentsAgentIdTokensResponse, error)

	// PostApiV4ProjectsIdClusterAgentsAgentIdTokensWithBodyWithResponse request with any body
	PostApiV4ProjectsIdClusterAgentsAgentIdTokensWithBodyWithResponse(ctx context.Context, id string, agentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdClusterAgentsAgentIdTokensResponse, error)

	PostApiV4ProjectsIdClusterAgentsAgentIdTokensWithResponse(ctx context.Context, id string, agentId int32, body PostApiV4ProjectsIdClusterAgentsAgentIdTokensJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdClusterAgentsAgentIdTokensResponse, error)

	// DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenIdWithResponse request
	DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenIdWithResponse(ctx context.Context, id string, agentId int32, tokenId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdClusterAgentsAgentIdTokensTokenIdResponse, error)

	// GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenIdWithResponse request
	GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenIdWithResponse(ctx context.Context, id string, agentId int32, tokenId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdClusterAgentsAgentIdTokensTokenIdResponse, error)

	// GetApiV4ProjectsIdClustersWithResponse request
	GetApiV4ProjectsIdClustersWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdClustersParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdClustersResponse, error)

	// PostApiV4ProjectsIdClustersUserWithBodyWithResponse request with any body
	PostApiV4ProjectsIdClustersUserWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdClustersUserResponse, error)

	PostApiV4ProjectsIdClustersUserWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdClustersUserJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdClustersUserResponse, error)

	// DeleteApiV4ProjectsIdClustersClusterIdWithResponse request
	DeleteApiV4ProjectsIdClustersClusterIdWithResponse(ctx context.Context, id string, clusterId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdClustersClusterIdResponse, error)

	// GetApiV4ProjectsIdClustersClusterIdWithResponse request
	GetApiV4ProjectsIdClustersClusterIdWithResponse(ctx context.Context, id string, clusterId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdClustersClusterIdResponse, error)

	// PutApiV4ProjectsIdClustersClusterIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdClustersClusterIdWithBodyWithResponse(ctx context.Context, id string, clusterId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdClustersClusterIdResponse, error)

	PutApiV4ProjectsIdClustersClusterIdWithResponse(ctx context.Context, id string, clusterId int32, body PutApiV4ProjectsIdClustersClusterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdClustersClusterIdResponse, error)

	// PostApiV4ProjectsIdCreateCiConfigWithResponse request
	PostApiV4ProjectsIdCreateCiConfigWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdCreateCiConfigResponse, error)

	// GetApiV4ProjectsIdCustomAttributesWithResponse request
	GetApiV4ProjectsIdCustomAttributesWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdCustomAttributesResponse, error)

	// DeleteApiV4ProjectsIdCustomAttributesKeyWithResponse request
	DeleteApiV4ProjectsIdCustomAttributesKeyWithResponse(ctx context.Context, id int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdCustomAttributesKeyResponse, error)

	// GetApiV4ProjectsIdCustomAttributesKeyWithResponse request
	GetApiV4ProjectsIdCustomAttributesKeyWithResponse(ctx context.Context, id int32, key string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdCustomAttributesKeyResponse, error)

	// PutApiV4ProjectsIdCustomAttributesKeyWithBodyWithResponse request with any body
	PutApiV4ProjectsIdCustomAttributesKeyWithBodyWithResponse(ctx context.Context, id int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdCustomAttributesKeyResponse, error)

	PutApiV4ProjectsIdCustomAttributesKeyWithResponse(ctx context.Context, id int32, key string, body PutApiV4ProjectsIdCustomAttributesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdCustomAttributesKeyResponse, error)

	// GetApiV4ProjectsIdDebianDistributionsWithResponse request
	GetApiV4ProjectsIdDebianDistributionsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdDebianDistributionsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDebianDistributionsResponse, error)

	// PostApiV4ProjectsIdDebianDistributionsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdDebianDistributionsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDebianDistributionsResponse, error)

	PostApiV4ProjectsIdDebianDistributionsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdDebianDistributionsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDebianDistributionsResponse, error)

	// DeleteApiV4ProjectsIdDebianDistributionsCodenameWithResponse request
	DeleteApiV4ProjectsIdDebianDistributionsCodenameWithResponse(ctx context.Context, id string, codename string, params *DeleteApiV4ProjectsIdDebianDistributionsCodenameParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdDebianDistributionsCodenameResponse, error)

	// GetApiV4ProjectsIdDebianDistributionsCodenameWithResponse request
	GetApiV4ProjectsIdDebianDistributionsCodenameWithResponse(ctx context.Context, id string, codename string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDebianDistributionsCodenameResponse, error)

	// PutApiV4ProjectsIdDebianDistributionsCodenameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdDebianDistributionsCodenameWithBodyWithResponse(ctx context.Context, id string, codename string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdDebianDistributionsCodenameResponse, error)

	PutApiV4ProjectsIdDebianDistributionsCodenameWithResponse(ctx context.Context, id string, codename string, body PutApiV4ProjectsIdDebianDistributionsCodenameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdDebianDistributionsCodenameResponse, error)

	// GetApiV4ProjectsIdDebianDistributionsCodenameKeyAscWithResponse request
	GetApiV4ProjectsIdDebianDistributionsCodenameKeyAscWithResponse(ctx context.Context, id string, codename string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDebianDistributionsCodenameKeyAscResponse, error)

	// GetApiV4ProjectsIdDeployKeysWithResponse request
	GetApiV4ProjectsIdDeployKeysWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdDeployKeysParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDeployKeysResponse, error)

	// PostApiV4ProjectsIdDeployKeysWithBodyWithResponse request with any body
	PostApiV4ProjectsIdDeployKeysWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeployKeysResponse, error)

	PostApiV4ProjectsIdDeployKeysWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdDeployKeysJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeployKeysResponse, error)

	// DeleteApiV4ProjectsIdDeployKeysKeyIdWithResponse request
	DeleteApiV4ProjectsIdDeployKeysKeyIdWithResponse(ctx context.Context, id string, keyId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdDeployKeysKeyIdResponse, error)

	// GetApiV4ProjectsIdDeployKeysKeyIdWithResponse request
	GetApiV4ProjectsIdDeployKeysKeyIdWithResponse(ctx context.Context, id string, keyId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDeployKeysKeyIdResponse, error)

	// PutApiV4ProjectsIdDeployKeysKeyIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdDeployKeysKeyIdWithBodyWithResponse(ctx context.Context, id string, keyId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdDeployKeysKeyIdResponse, error)

	PutApiV4ProjectsIdDeployKeysKeyIdWithResponse(ctx context.Context, id string, keyId int32, body PutApiV4ProjectsIdDeployKeysKeyIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdDeployKeysKeyIdResponse, error)

	// PostApiV4ProjectsIdDeployKeysKeyIdEnableWithResponse request
	PostApiV4ProjectsIdDeployKeysKeyIdEnableWithResponse(ctx context.Context, id string, keyId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeployKeysKeyIdEnableResponse, error)

	// GetApiV4ProjectsIdDeployTokensWithResponse request
	GetApiV4ProjectsIdDeployTokensWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdDeployTokensParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDeployTokensResponse, error)

	// PostApiV4ProjectsIdDeployTokensWithBodyWithResponse request with any body
	PostApiV4ProjectsIdDeployTokensWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeployTokensResponse, error)

	PostApiV4ProjectsIdDeployTokensWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdDeployTokensJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeployTokensResponse, error)

	// DeleteApiV4ProjectsIdDeployTokensTokenIdWithResponse request
	DeleteApiV4ProjectsIdDeployTokensTokenIdWithResponse(ctx context.Context, id string, tokenId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdDeployTokensTokenIdResponse, error)

	// GetApiV4ProjectsIdDeployTokensTokenIdWithResponse request
	GetApiV4ProjectsIdDeployTokensTokenIdWithResponse(ctx context.Context, id string, tokenId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDeployTokensTokenIdResponse, error)

	// GetApiV4ProjectsIdDeploymentsWithResponse request
	GetApiV4ProjectsIdDeploymentsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdDeploymentsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDeploymentsResponse, error)

	// PostApiV4ProjectsIdDeploymentsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdDeploymentsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeploymentsResponse, error)

	PostApiV4ProjectsIdDeploymentsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdDeploymentsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeploymentsResponse, error)

	// DeleteApiV4ProjectsIdDeploymentsDeploymentIdWithResponse request
	DeleteApiV4ProjectsIdDeploymentsDeploymentIdWithResponse(ctx context.Context, id string, deploymentId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdDeploymentsDeploymentIdResponse, error)

	// GetApiV4ProjectsIdDeploymentsDeploymentIdWithResponse request
	GetApiV4ProjectsIdDeploymentsDeploymentIdWithResponse(ctx context.Context, id string, deploymentId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDeploymentsDeploymentIdResponse, error)

	// PutApiV4ProjectsIdDeploymentsDeploymentIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdDeploymentsDeploymentIdWithBodyWithResponse(ctx context.Context, id string, deploymentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdDeploymentsDeploymentIdResponse, error)

	PutApiV4ProjectsIdDeploymentsDeploymentIdWithResponse(ctx context.Context, id string, deploymentId int32, body PutApiV4ProjectsIdDeploymentsDeploymentIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdDeploymentsDeploymentIdResponse, error)

	// PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalWithBodyWithResponse request with any body
	PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalWithBodyWithResponse(ctx context.Context, id string, deploymentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalResponse, error)

	PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalWithResponse(ctx context.Context, id string, deploymentId int32, body PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdDeploymentsDeploymentIdApprovalResponse, error)

	// GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsWithResponse request
	GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsWithResponse(ctx context.Context, id string, deploymentId int32, params *GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdDeploymentsDeploymentIdMergeRequestsResponse, error)

	// GetApiV4ProjectsIdEnvironmentsWithResponse request
	GetApiV4ProjectsIdEnvironmentsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdEnvironmentsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdEnvironmentsResponse, error)

	// PostApiV4ProjectsIdEnvironmentsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdEnvironmentsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdEnvironmentsResponse, error)

	PostApiV4ProjectsIdEnvironmentsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdEnvironmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdEnvironmentsResponse, error)

	// DeleteApiV4ProjectsIdEnvironmentsReviewAppsWithResponse request
	DeleteApiV4ProjectsIdEnvironmentsReviewAppsWithResponse(ctx context.Context, id string, params *DeleteApiV4ProjectsIdEnvironmentsReviewAppsParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdEnvironmentsReviewAppsResponse, error)

	// PostApiV4ProjectsIdEnvironmentsStopStaleWithBodyWithResponse request with any body
	PostApiV4ProjectsIdEnvironmentsStopStaleWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdEnvironmentsStopStaleResponse, error)

	PostApiV4ProjectsIdEnvironmentsStopStaleWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdEnvironmentsStopStaleJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdEnvironmentsStopStaleResponse, error)

	// DeleteApiV4ProjectsIdEnvironmentsEnvironmentIdWithResponse request
	DeleteApiV4ProjectsIdEnvironmentsEnvironmentIdWithResponse(ctx context.Context, id string, environmentId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdEnvironmentsEnvironmentIdResponse, error)

	// GetApiV4ProjectsIdEnvironmentsEnvironmentIdWithResponse request
	GetApiV4ProjectsIdEnvironmentsEnvironmentIdWithResponse(ctx context.Context, id string, environmentId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdEnvironmentsEnvironmentIdResponse, error)

	// PutApiV4ProjectsIdEnvironmentsEnvironmentIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdEnvironmentsEnvironmentIdWithBodyWithResponse(ctx context.Context, id string, environmentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdEnvironmentsEnvironmentIdResponse, error)

	PutApiV4ProjectsIdEnvironmentsEnvironmentIdWithResponse(ctx context.Context, id string, environmentId int32, body PutApiV4ProjectsIdEnvironmentsEnvironmentIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdEnvironmentsEnvironmentIdResponse, error)

	// PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopWithBodyWithResponse request with any body
	PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopWithBodyWithResponse(ctx context.Context, id string, environmentId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopResponse, error)

	PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopWithResponse(ctx context.Context, id string, environmentId int32, body PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdEnvironmentsEnvironmentIdStopResponse, error)

	// GetApiV4ProjectsIdErrorTrackingClientKeysWithResponse request
	GetApiV4ProjectsIdErrorTrackingClientKeysWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdErrorTrackingClientKeysResponse, error)

	// PostApiV4ProjectsIdErrorTrackingClientKeysWithResponse request
	PostApiV4ProjectsIdErrorTrackingClientKeysWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdErrorTrackingClientKeysResponse, error)

	// DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyIdWithResponse request
	DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyIdWithResponse(ctx context.Context, id string, keyId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdErrorTrackingClientKeysKeyIdResponse, error)

	// GetApiV4ProjectsIdErrorTrackingSettingsWithResponse request
	GetApiV4ProjectsIdErrorTrackingSettingsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdErrorTrackingSettingsResponse, error)

	// PatchApiV4ProjectsIdErrorTrackingSettingsWithBodyWithResponse request with any body
	PatchApiV4ProjectsIdErrorTrackingSettingsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdErrorTrackingSettingsResponse, error)

	PatchApiV4ProjectsIdErrorTrackingSettingsWithResponse(ctx context.Context, id string, body PatchApiV4ProjectsIdErrorTrackingSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdErrorTrackingSettingsResponse, error)

	// PutApiV4ProjectsIdErrorTrackingSettingsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdErrorTrackingSettingsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdErrorTrackingSettingsResponse, error)

	PutApiV4ProjectsIdErrorTrackingSettingsWithResponse(ctx context.Context, id string, body PutApiV4ProjectsIdErrorTrackingSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdErrorTrackingSettingsResponse, error)

	// GetApiV4ProjectsIdEventsWithResponse request
	GetApiV4ProjectsIdEventsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdEventsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdEventsResponse, error)

	// GetApiV4ProjectsIdExportWithResponse request
	GetApiV4ProjectsIdExportWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdExportResponse, error)

	// PostApiV4ProjectsIdExportWithBodyWithResponse request with any body
	PostApiV4ProjectsIdExportWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdExportResponse, error)

	PostApiV4ProjectsIdExportWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdExportJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdExportResponse, error)

	// GetApiV4ProjectsIdExportDownloadWithResponse request
	GetApiV4ProjectsIdExportDownloadWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdExportDownloadResponse, error)

	// PostApiV4ProjectsIdExportRelationsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdExportRelationsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdExportRelationsResponse, error)

	PostApiV4ProjectsIdExportRelationsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdExportRelationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdExportRelationsResponse, error)

	// GetApiV4ProjectsIdExportRelationsDownloadWithResponse request
	GetApiV4ProjectsIdExportRelationsDownloadWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdExportRelationsDownloadParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdExportRelationsDownloadResponse, error)

	// GetApiV4ProjectsIdExportRelationsStatusWithResponse request
	GetApiV4ProjectsIdExportRelationsStatusWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdExportRelationsStatusParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdExportRelationsStatusResponse, error)

	// GetApiV4ProjectsIdFeatureFlagsWithResponse request
	GetApiV4ProjectsIdFeatureFlagsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdFeatureFlagsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdFeatureFlagsResponse, error)

	// PostApiV4ProjectsIdFeatureFlagsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdFeatureFlagsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdFeatureFlagsResponse, error)

	PostApiV4ProjectsIdFeatureFlagsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdFeatureFlagsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdFeatureFlagsResponse, error)

	// DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithResponse request
	DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithResponse(ctx context.Context, id string, featureFlagName string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdFeatureFlagsFeatureFlagNameResponse, error)

	// GetApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithResponse request
	GetApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithResponse(ctx context.Context, id string, featureFlagName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdFeatureFlagsFeatureFlagNameResponse, error)

	// PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithBodyWithResponse(ctx context.Context, id string, featureFlagName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameResponse, error)

	PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameWithResponse(ctx context.Context, id string, featureFlagName string, body PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdFeatureFlagsFeatureFlagNameResponse, error)

	// GetApiV4ProjectsIdFeatureFlagsUserListsWithResponse request
	GetApiV4ProjectsIdFeatureFlagsUserListsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdFeatureFlagsUserListsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdFeatureFlagsUserListsResponse, error)

	// PostApiV4ProjectsIdFeatureFlagsUserListsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdFeatureFlagsUserListsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdFeatureFlagsUserListsResponse, error)

	PostApiV4ProjectsIdFeatureFlagsUserListsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdFeatureFlagsUserListsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdFeatureFlagsUserListsResponse, error)

	// DeleteApiV4ProjectsIdFeatureFlagsUserListsIidWithResponse request
	DeleteApiV4ProjectsIdFeatureFlagsUserListsIidWithResponse(ctx context.Context, id string, iid string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdFeatureFlagsUserListsIidResponse, error)

	// GetApiV4ProjectsIdFeatureFlagsUserListsIidWithResponse request
	GetApiV4ProjectsIdFeatureFlagsUserListsIidWithResponse(ctx context.Context, id string, iid string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdFeatureFlagsUserListsIidResponse, error)

	// PutApiV4ProjectsIdFeatureFlagsUserListsIidWithBodyWithResponse request with any body
	PutApiV4ProjectsIdFeatureFlagsUserListsIidWithBodyWithResponse(ctx context.Context, id string, iid string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdFeatureFlagsUserListsIidResponse, error)

	PutApiV4ProjectsIdFeatureFlagsUserListsIidWithResponse(ctx context.Context, id string, iid string, body PutApiV4ProjectsIdFeatureFlagsUserListsIidJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdFeatureFlagsUserListsIidResponse, error)

	// DeleteApiV4ProjectsIdForkWithResponse request
	DeleteApiV4ProjectsIdForkWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdForkResponse, error)

	// PostApiV4ProjectsIdForkWithBodyWithResponse request with any body
	PostApiV4ProjectsIdForkWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdForkResponse, error)

	PostApiV4ProjectsIdForkWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdForkJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdForkResponse, error)

	// PostApiV4ProjectsIdForkForkedFromIdWithResponse request
	PostApiV4ProjectsIdForkForkedFromIdWithResponse(ctx context.Context, id string, forkedFromId string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdForkForkedFromIdResponse, error)

	// GetApiV4ProjectsIdForksWithResponse request
	GetApiV4ProjectsIdForksWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdForksParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdForksResponse, error)

	// GetApiV4ProjectsIdFreezePeriodsWithResponse request
	GetApiV4ProjectsIdFreezePeriodsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdFreezePeriodsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdFreezePeriodsResponse, error)

	// PostApiV4ProjectsIdFreezePeriodsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdFreezePeriodsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdFreezePeriodsResponse, error)

	PostApiV4ProjectsIdFreezePeriodsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdFreezePeriodsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdFreezePeriodsResponse, error)

	// DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithResponse request
	DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithResponse(ctx context.Context, id string, freezePeriodId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdFreezePeriodsFreezePeriodIdResponse, error)

	// GetApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithResponse request
	GetApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithResponse(ctx context.Context, id string, freezePeriodId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdFreezePeriodsFreezePeriodIdResponse, error)

	// PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithBodyWithResponse(ctx context.Context, id string, freezePeriodId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdResponse, error)

	PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdWithResponse(ctx context.Context, id string, freezePeriodId int32, body PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdFreezePeriodsFreezePeriodIdResponse, error)

	// GetApiV4ProjectsIdGroupsWithResponse request
	GetApiV4ProjectsIdGroupsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdGroupsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdGroupsResponse, error)

	// GetApiV4ProjectsIdHooksWithResponse request
	GetApiV4ProjectsIdHooksWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdHooksParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdHooksResponse, error)

	// PostApiV4ProjectsIdHooksWithBodyWithResponse request with any body
	PostApiV4ProjectsIdHooksWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdHooksResponse, error)

	PostApiV4ProjectsIdHooksWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdHooksJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdHooksResponse, error)

	// DeleteApiV4ProjectsIdHooksHookIdWithResponse request
	DeleteApiV4ProjectsIdHooksHookIdWithResponse(ctx context.Context, id string, hookId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdHooksHookIdResponse, error)

	// GetApiV4ProjectsIdHooksHookIdWithResponse request
	GetApiV4ProjectsIdHooksHookIdWithResponse(ctx context.Context, id string, hookId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdHooksHookIdResponse, error)

	// PutApiV4ProjectsIdHooksHookIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdHooksHookIdWithBodyWithResponse(ctx context.Context, id string, hookId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdHooksHookIdResponse, error)

	PutApiV4ProjectsIdHooksHookIdWithResponse(ctx context.Context, id string, hookId int32, body PutApiV4ProjectsIdHooksHookIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdHooksHookIdResponse, error)

	// DeleteApiV4ProjectsIdHooksHookIdCustomHeadersKeyWithResponse request
	DeleteApiV4ProjectsIdHooksHookIdCustomHeadersKeyWithResponse(ctx context.Context, id int32, hookId int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdHooksHookIdCustomHeadersKeyResponse, error)

	// PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyWithBodyWithResponse request with any body
	PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyWithBodyWithResponse(ctx context.Context, id int32, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyResponse, error)

	PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyWithResponse(ctx context.Context, id int32, hookId int32, key string, body PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdHooksHookIdCustomHeadersKeyResponse, error)

	// GetApiV4ProjectsIdHooksHookIdEventsWithResponse request
	GetApiV4ProjectsIdHooksHookIdEventsWithResponse(ctx context.Context, id int32, hookId int32, params *GetApiV4ProjectsIdHooksHookIdEventsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdHooksHookIdEventsResponse, error)

	// PostApiV4ProjectsIdHooksHookIdEventsHookLogIdResendWithResponse request
	PostApiV4ProjectsIdHooksHookIdEventsHookLogIdResendWithResponse(ctx context.Context, id int32, hookId int32, hookLogId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdHooksHookIdEventsHookLogIdResendResponse, error)

	// PostApiV4ProjectsIdHooksHookIdTestTriggerWithResponse request
	PostApiV4ProjectsIdHooksHookIdTestTriggerWithResponse(ctx context.Context, id int32, hookId int32, trigger string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdHooksHookIdTestTriggerResponse, error)

	// DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKeyWithResponse request
	DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKeyWithResponse(ctx context.Context, id int32, hookId int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdHooksHookIdUrlVariablesKeyResponse, error)

	// PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyWithBodyWithResponse request with any body
	PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyWithBodyWithResponse(ctx context.Context, id int32, hookId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyResponse, error)

	PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyWithResponse(ctx context.Context, id int32, hookId int32, key string, body PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdHooksHookIdUrlVariablesKeyResponse, error)

	// PostApiV4ProjectsIdHousekeepingWithBodyWithResponse request with any body
	PostApiV4ProjectsIdHousekeepingWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdHousekeepingResponse, error)

	PostApiV4ProjectsIdHousekeepingWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdHousekeepingJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdHousekeepingResponse, error)

	// GetApiV4ProjectsIdImportWithResponse request
	GetApiV4ProjectsIdImportWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdImportResponse, error)

	// PostApiV4ProjectsIdImportProjectMembersProjectIdWithResponse request
	PostApiV4ProjectsIdImportProjectMembersProjectIdWithResponse(ctx context.Context, id string, projectId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdImportProjectMembersProjectIdResponse, error)

	// GetApiV4ProjectsIdIntegrationsWithResponse request
	GetApiV4ProjectsIdIntegrationsWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIntegrationsResponse, error)

	// PutApiV4ProjectsIdIntegrationsAppleAppStoreWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsAppleAppStoreWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsAppleAppStoreResponse, error)

	PutApiV4ProjectsIdIntegrationsAppleAppStoreWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsAppleAppStoreJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsAppleAppStoreResponse, error)

	// PutApiV4ProjectsIdIntegrationsAsanaWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsAsanaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsAsanaResponse, error)

	PutApiV4ProjectsIdIntegrationsAsanaWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsAsanaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsAsanaResponse, error)

	// PutApiV4ProjectsIdIntegrationsAssemblaWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsAssemblaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsAssemblaResponse, error)

	PutApiV4ProjectsIdIntegrationsAssemblaWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsAssemblaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsAssemblaResponse, error)

	// PutApiV4ProjectsIdIntegrationsBambooWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsBambooWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsBambooResponse, error)

	PutApiV4ProjectsIdIntegrationsBambooWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsBambooJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsBambooResponse, error)

	// PutApiV4ProjectsIdIntegrationsBugzillaWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsBugzillaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsBugzillaResponse, error)

	PutApiV4ProjectsIdIntegrationsBugzillaWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsBugzillaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsBugzillaResponse, error)

	// PutApiV4ProjectsIdIntegrationsBuildkiteWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsBuildkiteWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsBuildkiteResponse, error)

	PutApiV4ProjectsIdIntegrationsBuildkiteWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsBuildkiteJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsBuildkiteResponse, error)

	// PutApiV4ProjectsIdIntegrationsCampfireWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsCampfireWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsCampfireResponse, error)

	PutApiV4ProjectsIdIntegrationsCampfireWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsCampfireJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsCampfireResponse, error)

	// PutApiV4ProjectsIdIntegrationsClickupWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsClickupWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsClickupResponse, error)

	PutApiV4ProjectsIdIntegrationsClickupWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsClickupJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsClickupResponse, error)

	// PutApiV4ProjectsIdIntegrationsConfluenceWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsConfluenceWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsConfluenceResponse, error)

	PutApiV4ProjectsIdIntegrationsConfluenceWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsConfluenceJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsConfluenceResponse, error)

	// PutApiV4ProjectsIdIntegrationsCustomIssueTrackerWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsCustomIssueTrackerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsCustomIssueTrackerResponse, error)

	PutApiV4ProjectsIdIntegrationsCustomIssueTrackerWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsCustomIssueTrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsCustomIssueTrackerResponse, error)

	// PutApiV4ProjectsIdIntegrationsDatadogWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsDatadogWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsDatadogResponse, error)

	PutApiV4ProjectsIdIntegrationsDatadogWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsDatadogJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsDatadogResponse, error)

	// PutApiV4ProjectsIdIntegrationsDiffblueCoverWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsDiffblueCoverWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsDiffblueCoverResponse, error)

	PutApiV4ProjectsIdIntegrationsDiffblueCoverWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsDiffblueCoverJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsDiffblueCoverResponse, error)

	// PutApiV4ProjectsIdIntegrationsDiscordWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsDiscordWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsDiscordResponse, error)

	PutApiV4ProjectsIdIntegrationsDiscordWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsDiscordJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsDiscordResponse, error)

	// PutApiV4ProjectsIdIntegrationsDroneCiWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsDroneCiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsDroneCiResponse, error)

	PutApiV4ProjectsIdIntegrationsDroneCiWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsDroneCiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsDroneCiResponse, error)

	// PutApiV4ProjectsIdIntegrationsEmailsOnPushWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsEmailsOnPushWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsEmailsOnPushResponse, error)

	PutApiV4ProjectsIdIntegrationsEmailsOnPushWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsEmailsOnPushJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsEmailsOnPushResponse, error)

	// PutApiV4ProjectsIdIntegrationsEwmWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsEwmWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsEwmResponse, error)

	PutApiV4ProjectsIdIntegrationsEwmWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsEwmJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsEwmResponse, error)

	// PutApiV4ProjectsIdIntegrationsExternalWikiWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsExternalWikiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsExternalWikiResponse, error)

	PutApiV4ProjectsIdIntegrationsExternalWikiWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsExternalWikiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsExternalWikiResponse, error)

	// PutApiV4ProjectsIdIntegrationsGitGuardianWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsGitGuardianWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGitGuardianResponse, error)

	PutApiV4ProjectsIdIntegrationsGitGuardianWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGitGuardianJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGitGuardianResponse, error)

	// PutApiV4ProjectsIdIntegrationsGithubWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsGithubWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGithubResponse, error)

	PutApiV4ProjectsIdIntegrationsGithubWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGithubResponse, error)

	// PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationResponse, error)

	PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGitlabSlackApplicationResponse, error)

	// PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryResponse, error)

	PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformArtifactRegistryResponse, error)

	// PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationResponse, error)

	PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGoogleCloudPlatformWorkloadIdentityFederationResponse, error)

	// PutApiV4ProjectsIdIntegrationsGooglePlayWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsGooglePlayWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGooglePlayResponse, error)

	PutApiV4ProjectsIdIntegrationsGooglePlayWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsGooglePlayJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsGooglePlayResponse, error)

	// PutApiV4ProjectsIdIntegrationsHangoutsChatWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsHangoutsChatWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsHangoutsChatResponse, error)

	PutApiV4ProjectsIdIntegrationsHangoutsChatWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsHangoutsChatJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsHangoutsChatResponse, error)

	// PutApiV4ProjectsIdIntegrationsHarborWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsHarborWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsHarborResponse, error)

	PutApiV4ProjectsIdIntegrationsHarborWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsHarborJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsHarborResponse, error)

	// PutApiV4ProjectsIdIntegrationsIrkerWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsIrkerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsIrkerResponse, error)

	PutApiV4ProjectsIdIntegrationsIrkerWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsIrkerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsIrkerResponse, error)

	// PutApiV4ProjectsIdIntegrationsJenkinsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsJenkinsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsJenkinsResponse, error)

	PutApiV4ProjectsIdIntegrationsJenkinsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsJenkinsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsJenkinsResponse, error)

	// PutApiV4ProjectsIdIntegrationsJiraWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsJiraWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsJiraResponse, error)

	PutApiV4ProjectsIdIntegrationsJiraWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsJiraJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsJiraResponse, error)

	// PutApiV4ProjectsIdIntegrationsJiraCloudAppWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsJiraCloudAppWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsJiraCloudAppResponse, error)

	PutApiV4ProjectsIdIntegrationsJiraCloudAppWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsJiraCloudAppJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsJiraCloudAppResponse, error)

	// PutApiV4ProjectsIdIntegrationsMatrixWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsMatrixWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMatrixResponse, error)

	PutApiV4ProjectsIdIntegrationsMatrixWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMatrixJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMatrixResponse, error)

	// PutApiV4ProjectsIdIntegrationsMattermostWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsMattermostWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMattermostResponse, error)

	PutApiV4ProjectsIdIntegrationsMattermostWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMattermostJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMattermostResponse, error)

	// PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsResponse, error)

	PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMattermostSlashCommandsResponse, error)

	// PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerWithBodyWithResponse request with any body
	PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerResponse, error)

	PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIntegrationsMattermostSlashCommandsTriggerResponse, error)

	// PutApiV4ProjectsIdIntegrationsMicrosoftTeamsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsMicrosoftTeamsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMicrosoftTeamsResponse, error)

	PutApiV4ProjectsIdIntegrationsMicrosoftTeamsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMicrosoftTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMicrosoftTeamsResponse, error)

	// PutApiV4ProjectsIdIntegrationsMockCiWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsMockCiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMockCiResponse, error)

	PutApiV4ProjectsIdIntegrationsMockCiWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMockCiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMockCiResponse, error)

	// PutApiV4ProjectsIdIntegrationsMockMonitoringWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsMockMonitoringWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMockMonitoringResponse, error)

	PutApiV4ProjectsIdIntegrationsMockMonitoringWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsMockMonitoringJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsMockMonitoringResponse, error)

	// PutApiV4ProjectsIdIntegrationsPackagistWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsPackagistWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPackagistResponse, error)

	PutApiV4ProjectsIdIntegrationsPackagistWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPackagistJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPackagistResponse, error)

	// PutApiV4ProjectsIdIntegrationsPhorgeWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsPhorgeWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPhorgeResponse, error)

	PutApiV4ProjectsIdIntegrationsPhorgeWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPhorgeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPhorgeResponse, error)

	// PutApiV4ProjectsIdIntegrationsPipelinesEmailWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsPipelinesEmailWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPipelinesEmailResponse, error)

	PutApiV4ProjectsIdIntegrationsPipelinesEmailWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPipelinesEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPipelinesEmailResponse, error)

	// PutApiV4ProjectsIdIntegrationsPivotaltrackerWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsPivotaltrackerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPivotaltrackerResponse, error)

	PutApiV4ProjectsIdIntegrationsPivotaltrackerWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPivotaltrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPivotaltrackerResponse, error)

	// PutApiV4ProjectsIdIntegrationsPumbleWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsPumbleWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPumbleResponse, error)

	PutApiV4ProjectsIdIntegrationsPumbleWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPumbleJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPumbleResponse, error)

	// PutApiV4ProjectsIdIntegrationsPushoverWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsPushoverWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPushoverResponse, error)

	PutApiV4ProjectsIdIntegrationsPushoverWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsPushoverJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsPushoverResponse, error)

	// PutApiV4ProjectsIdIntegrationsRedmineWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsRedmineWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsRedmineResponse, error)

	PutApiV4ProjectsIdIntegrationsRedmineWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsRedmineJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsRedmineResponse, error)

	// PutApiV4ProjectsIdIntegrationsSlackWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsSlackWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsSlackResponse, error)

	PutApiV4ProjectsIdIntegrationsSlackWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsSlackJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsSlackResponse, error)

	// PutApiV4ProjectsIdIntegrationsSlackSlashCommandsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsSlackSlashCommandsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsSlackSlashCommandsResponse, error)

	PutApiV4ProjectsIdIntegrationsSlackSlashCommandsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsSlackSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsSlackSlashCommandsResponse, error)

	// PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerWithBodyWithResponse request with any body
	PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerResponse, error)

	PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIntegrationsSlackSlashCommandsTriggerResponse, error)

	// PutApiV4ProjectsIdIntegrationsSquashTmWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsSquashTmWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsSquashTmResponse, error)

	PutApiV4ProjectsIdIntegrationsSquashTmWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsSquashTmJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsSquashTmResponse, error)

	// PutApiV4ProjectsIdIntegrationsTeamcityWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsTeamcityWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsTeamcityResponse, error)

	PutApiV4ProjectsIdIntegrationsTeamcityWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsTeamcityJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsTeamcityResponse, error)

	// PutApiV4ProjectsIdIntegrationsTelegramWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsTelegramWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsTelegramResponse, error)

	PutApiV4ProjectsIdIntegrationsTelegramWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsTelegramJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsTelegramResponse, error)

	// PutApiV4ProjectsIdIntegrationsUnifyCircuitWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsUnifyCircuitWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsUnifyCircuitResponse, error)

	PutApiV4ProjectsIdIntegrationsUnifyCircuitWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsUnifyCircuitJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsUnifyCircuitResponse, error)

	// PutApiV4ProjectsIdIntegrationsWebexTeamsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsWebexTeamsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsWebexTeamsResponse, error)

	PutApiV4ProjectsIdIntegrationsWebexTeamsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsWebexTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsWebexTeamsResponse, error)

	// PutApiV4ProjectsIdIntegrationsYoutrackWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsYoutrackWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsYoutrackResponse, error)

	PutApiV4ProjectsIdIntegrationsYoutrackWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsYoutrackJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsYoutrackResponse, error)

	// PutApiV4ProjectsIdIntegrationsZentaoWithBodyWithResponse request with any body
	PutApiV4ProjectsIdIntegrationsZentaoWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsZentaoResponse, error)

	PutApiV4ProjectsIdIntegrationsZentaoWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdIntegrationsZentaoJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIntegrationsZentaoResponse, error)

	// DeleteApiV4ProjectsIdIntegrationsSlugWithResponse request
	DeleteApiV4ProjectsIdIntegrationsSlugWithResponse(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdIntegrationsSlugResponse, error)

	// GetApiV4ProjectsIdIntegrationsSlugWithResponse request
	GetApiV4ProjectsIdIntegrationsSlugWithResponse(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIntegrationsSlugResponse, error)

	// GetApiV4ProjectsIdInvitationsWithResponse request
	GetApiV4ProjectsIdInvitationsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdInvitationsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdInvitationsResponse, error)

	// PostApiV4ProjectsIdInvitationsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdInvitationsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdInvitationsResponse, error)

	PostApiV4ProjectsIdInvitationsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdInvitationsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdInvitationsResponse, error)

	// DeleteApiV4ProjectsIdInvitationsEmailWithResponse request
	DeleteApiV4ProjectsIdInvitationsEmailWithResponse(ctx context.Context, id string, email string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdInvitationsEmailResponse, error)

	// PutApiV4ProjectsIdInvitationsEmailWithBodyWithResponse request with any body
	PutApiV4ProjectsIdInvitationsEmailWithBodyWithResponse(ctx context.Context, id string, email string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdInvitationsEmailResponse, error)

	PutApiV4ProjectsIdInvitationsEmailWithResponse(ctx context.Context, id string, email string, body PutApiV4ProjectsIdInvitationsEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdInvitationsEmailResponse, error)

	// GetApiV4ProjectsIdInvitedGroupsWithResponse request
	GetApiV4ProjectsIdInvitedGroupsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdInvitedGroupsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdInvitedGroupsResponse, error)

	// PostApiV4ProjectsIdIssuesWithResponse request
	PostApiV4ProjectsIdIssuesWithResponse(ctx context.Context, id string, params *PostApiV4ProjectsIdIssuesParams, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesResponse, error)

	// GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsWithResponse request
	GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsWithResponse(ctx context.Context, id string, eventableId int32, params *GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsResponse, error)

	// GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdWithResponse request
	GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdWithResponse(ctx context.Context, id string, eventableId int32, eventId string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesEventableIdResourceMilestoneEventsEventIdResponse, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidWithResponse request
	DeleteApiV4ProjectsIdIssuesIssueIidWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdIssuesIssueIidResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidAddSpentTimeWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidAddSpentTimeWithResponse(ctx context.Context, id string, issueIid int, params *PostApiV4ProjectsIdIssuesIssueIidAddSpentTimeParams, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidAddSpentTimeResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidAwardEmojiWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidAwardEmojiWithResponse(ctx context.Context, id string, issueIid int32, params *GetApiV4ProjectsIdIssuesIssueIidAwardEmojiParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidAwardEmojiResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidAwardEmojiWithBodyWithResponse request with any body
	PostApiV4ProjectsIdIssuesIssueIidAwardEmojiWithBodyWithResponse(ctx context.Context, id int32, issueIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidAwardEmojiResponse, error)

	PostApiV4ProjectsIdIssuesIssueIidAwardEmojiWithResponse(ctx context.Context, id int32, issueIid int32, body PostApiV4ProjectsIdIssuesIssueIidAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidAwardEmojiResponse, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardIdWithResponse request
	DeleteApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, issueIid int32, awardId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardIdWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, issueIid int32, awardId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidAwardEmojiAwardIdResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidCloneWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidCloneWithResponse(ctx context.Context, id string, issueIid int, params *PostApiV4ProjectsIdIssuesIssueIidCloneParams, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidCloneResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidClosedByWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidClosedByWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidClosedByResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidLinksWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidLinksWithResponse(ctx context.Context, id string, issueIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidLinksResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidLinksWithBodyWithResponse request with any body
	PostApiV4ProjectsIdIssuesIssueIidLinksWithBodyWithResponse(ctx context.Context, id string, issueIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidLinksResponse, error)

	PostApiV4ProjectsIdIssuesIssueIidLinksWithResponse(ctx context.Context, id string, issueIid int32, body PostApiV4ProjectsIdIssuesIssueIidLinksJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidLinksResponse, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkIdWithResponse request
	DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkIdWithResponse(ctx context.Context, id string, issueIid int32, issueLinkId string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdIssuesIssueIidLinksIssueLinkIdResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkIdWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkIdWithResponse(ctx context.Context, id string, issueIid int32, issueLinkId string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidLinksIssueLinkIdResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidMetricImagesWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidMetricImagesWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidMetricImagesResponse, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidMetricImagesImageIdWithResponse request
	DeleteApiV4ProjectsIdIssuesIssueIidMetricImagesImageIdWithResponse(ctx context.Context, id string, issueIid int, imageId int, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdIssuesIssueIidMetricImagesImageIdResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidMoveWithBodyWithResponse request with any body
	PostApiV4ProjectsIdIssuesIssueIidMoveWithBodyWithResponse(ctx context.Context, id string, issueIid int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidMoveResponse, error)

	PostApiV4ProjectsIdIssuesIssueIidMoveWithFormdataBodyWithResponse(ctx context.Context, id string, issueIid int, body PostApiV4ProjectsIdIssuesIssueIidMoveFormdataRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidMoveResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidNotesWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidNotesWithResponse(ctx context.Context, id string, issueIid int, params *PostApiV4ProjectsIdIssuesIssueIidNotesParams, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidNotesResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiWithResponse(ctx context.Context, id int32, issueIid int32, noteId int32, params *GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiWithBodyWithResponse request with any body
	PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiWithBodyWithResponse(ctx context.Context, id int32, issueIid int32, noteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiResponse, error)

	PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiWithResponse(ctx context.Context, id int32, issueIid int32, noteId int32, body PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiResponse, error)

	// DeleteApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdWithResponse request
	DeleteApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, issueIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, issueIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidNotesNoteIdAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidParticipantsWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidParticipantsWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidParticipantsResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidRelatedMergeRequestsWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidRelatedMergeRequestsWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidRelatedMergeRequestsResponse, error)

	// PutApiV4ProjectsIdIssuesIssueIidReorderWithResponse request
	PutApiV4ProjectsIdIssuesIssueIidReorderWithResponse(ctx context.Context, id string, issueIid int, params *PutApiV4ProjectsIdIssuesIssueIidReorderParams, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdIssuesIssueIidReorderResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidResetSpentTimeWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidResetSpentTimeWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidResetSpentTimeResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidResetTimeEstimateWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidResetTimeEstimateWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidResetTimeEstimateResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidSubscribeWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidSubscribeWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidSubscribeResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidTimeEstimateWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidTimeEstimateWithResponse(ctx context.Context, id string, issueIid int, params *PostApiV4ProjectsIdIssuesIssueIidTimeEstimateParams, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidTimeEstimateResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidTimeStatsWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidTimeStatsWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidTimeStatsResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidTodoWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidTodoWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidTodoResponse, error)

	// PostApiV4ProjectsIdIssuesIssueIidUnsubscribeWithResponse request
	PostApiV4ProjectsIdIssuesIssueIidUnsubscribeWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdIssuesIssueIidUnsubscribeResponse, error)

	// GetApiV4ProjectsIdIssuesIssueIidUserAgentDetailWithResponse request
	GetApiV4ProjectsIdIssuesIssueIidUserAgentDetailWithResponse(ctx context.Context, id string, issueIid int, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdIssuesIssueIidUserAgentDetailResponse, error)

	// GetApiV4ProjectsIdJobTokenScopeWithResponse request
	GetApiV4ProjectsIdJobTokenScopeWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobTokenScopeResponse, error)

	// PatchApiV4ProjectsIdJobTokenScopeWithBodyWithResponse request with any body
	PatchApiV4ProjectsIdJobTokenScopeWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdJobTokenScopeResponse, error)

	PatchApiV4ProjectsIdJobTokenScopeWithResponse(ctx context.Context, id int32, body PatchApiV4ProjectsIdJobTokenScopeJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdJobTokenScopeResponse, error)

	// GetApiV4ProjectsIdJobTokenScopeAllowlistWithResponse request
	GetApiV4ProjectsIdJobTokenScopeAllowlistWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdJobTokenScopeAllowlistParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobTokenScopeAllowlistResponse, error)

	// PostApiV4ProjectsIdJobTokenScopeAllowlistWithBodyWithResponse request with any body
	PostApiV4ProjectsIdJobTokenScopeAllowlistWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobTokenScopeAllowlistResponse, error)

	PostApiV4ProjectsIdJobTokenScopeAllowlistWithResponse(ctx context.Context, id int32, body PostApiV4ProjectsIdJobTokenScopeAllowlistJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobTokenScopeAllowlistResponse, error)

	// DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectIdWithResponse request
	DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectIdWithResponse(ctx context.Context, id int32, targetProjectId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdJobTokenScopeAllowlistTargetProjectIdResponse, error)

	// GetApiV4ProjectsIdJobTokenScopeGroupsAllowlistWithResponse request
	GetApiV4ProjectsIdJobTokenScopeGroupsAllowlistWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdJobTokenScopeGroupsAllowlistParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobTokenScopeGroupsAllowlistResponse, error)

	// PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistWithBodyWithResponse request with any body
	PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistResponse, error)

	PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistWithResponse(ctx context.Context, id int32, body PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobTokenScopeGroupsAllowlistResponse, error)

	// DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupIdWithResponse request
	DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupIdWithResponse(ctx context.Context, id int32, targetGroupId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdJobTokenScopeGroupsAllowlistTargetGroupIdResponse, error)

	// GetApiV4ProjectsIdJobsWithResponse request
	GetApiV4ProjectsIdJobsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdJobsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobsResponse, error)

	// GetApiV4ProjectsIdJobsArtifactsRefNameDownloadWithResponse request
	GetApiV4ProjectsIdJobsArtifactsRefNameDownloadWithResponse(ctx context.Context, id string, refName string, params *GetApiV4ProjectsIdJobsArtifactsRefNameDownloadParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobsArtifactsRefNameDownloadResponse, error)

	// GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPathWithResponse request
	GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPathWithResponse(ctx context.Context, id string, refName string, params *GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPathParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobsArtifactsRefNameRawartifactPathResponse, error)

	// GetApiV4ProjectsIdJobsJobIdWithResponse request
	GetApiV4ProjectsIdJobsJobIdWithResponse(ctx context.Context, id int32, jobId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobsJobIdResponse, error)

	// DeleteApiV4ProjectsIdJobsJobIdArtifactsWithResponse request
	DeleteApiV4ProjectsIdJobsJobIdArtifactsWithResponse(ctx context.Context, id string, jobId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdJobsJobIdArtifactsResponse, error)

	// GetApiV4ProjectsIdJobsJobIdArtifactsWithResponse request
	GetApiV4ProjectsIdJobsJobIdArtifactsWithResponse(ctx context.Context, id string, jobId int32, params *GetApiV4ProjectsIdJobsJobIdArtifactsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobsJobIdArtifactsResponse, error)

	// GetApiV4ProjectsIdJobsJobIdArtifactsartifactPathWithResponse request
	GetApiV4ProjectsIdJobsJobIdArtifactsartifactPathWithResponse(ctx context.Context, id string, jobId int32, params *GetApiV4ProjectsIdJobsJobIdArtifactsartifactPathParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobsJobIdArtifactsartifactPathResponse, error)

	// PostApiV4ProjectsIdJobsJobIdArtifactsKeepWithResponse request
	PostApiV4ProjectsIdJobsJobIdArtifactsKeepWithResponse(ctx context.Context, id string, jobId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobsJobIdArtifactsKeepResponse, error)

	// PostApiV4ProjectsIdJobsJobIdCancelWithBodyWithResponse request with any body
	PostApiV4ProjectsIdJobsJobIdCancelWithBodyWithResponse(ctx context.Context, id int32, jobId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobsJobIdCancelResponse, error)

	PostApiV4ProjectsIdJobsJobIdCancelWithResponse(ctx context.Context, id int32, jobId int32, body PostApiV4ProjectsIdJobsJobIdCancelJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobsJobIdCancelResponse, error)

	// PostApiV4ProjectsIdJobsJobIdEraseWithResponse request
	PostApiV4ProjectsIdJobsJobIdEraseWithResponse(ctx context.Context, id int32, jobId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobsJobIdEraseResponse, error)

	// PostApiV4ProjectsIdJobsJobIdPlayWithBodyWithResponse request with any body
	PostApiV4ProjectsIdJobsJobIdPlayWithBodyWithResponse(ctx context.Context, id int32, jobId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobsJobIdPlayResponse, error)

	PostApiV4ProjectsIdJobsJobIdPlayWithResponse(ctx context.Context, id int32, jobId int32, body PostApiV4ProjectsIdJobsJobIdPlayJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobsJobIdPlayResponse, error)

	// PostApiV4ProjectsIdJobsJobIdRetryWithResponse request
	PostApiV4ProjectsIdJobsJobIdRetryWithResponse(ctx context.Context, id int32, jobId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdJobsJobIdRetryResponse, error)

	// GetApiV4ProjectsIdJobsJobIdTraceWithResponse request
	GetApiV4ProjectsIdJobsJobIdTraceWithResponse(ctx context.Context, id int32, jobId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdJobsJobIdTraceResponse, error)

	// GetApiV4ProjectsIdLanguagesWithResponse request
	GetApiV4ProjectsIdLanguagesWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdLanguagesResponse, error)

	// GetApiV4ProjectsIdMembersWithResponse request
	GetApiV4ProjectsIdMembersWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdMembersParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMembersResponse, error)

	// PostApiV4ProjectsIdMembersWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMembersWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMembersResponse, error)

	PostApiV4ProjectsIdMembersWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdMembersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMembersResponse, error)

	// GetApiV4ProjectsIdMembersAllWithResponse request
	GetApiV4ProjectsIdMembersAllWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdMembersAllParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMembersAllResponse, error)

	// GetApiV4ProjectsIdMembersAllUserIdWithResponse request
	GetApiV4ProjectsIdMembersAllUserIdWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMembersAllUserIdResponse, error)

	// DeleteApiV4ProjectsIdMembersUserIdWithResponse request
	DeleteApiV4ProjectsIdMembersUserIdWithResponse(ctx context.Context, id string, userId int32, params *DeleteApiV4ProjectsIdMembersUserIdParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdMembersUserIdResponse, error)

	// GetApiV4ProjectsIdMembersUserIdWithResponse request
	GetApiV4ProjectsIdMembersUserIdWithResponse(ctx context.Context, id string, userId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMembersUserIdResponse, error)

	// PutApiV4ProjectsIdMembersUserIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdMembersUserIdWithBodyWithResponse(ctx context.Context, id string, userId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMembersUserIdResponse, error)

	PutApiV4ProjectsIdMembersUserIdWithResponse(ctx context.Context, id string, userId int32, body PutApiV4ProjectsIdMembersUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMembersUserIdResponse, error)

	// GetApiV4ProjectsIdMergeRequestsWithResponse request
	GetApiV4ProjectsIdMergeRequestsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdMergeRequestsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsResponse, error)

	// PostApiV4ProjectsIdMergeRequestsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsResponse, error)

	PostApiV4ProjectsIdMergeRequestsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdMergeRequestsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsResponse, error)

	// GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsWithResponse request
	GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsWithResponse(ctx context.Context, id string, eventableId int32, params *GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsResponse, error)

	// GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdWithResponse request
	GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdWithResponse(ctx context.Context, id string, eventableId int32, eventId string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsEventableIdResourceMilestoneEventsEventIdResponse, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidWithResponse request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidResponse, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidWithBodyWithResponse request with any body
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidResponse, error)

	PutApiV4ProjectsIdMergeRequestsMergeRequestIidWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PutApiV4ProjectsIdMergeRequestsMergeRequestIidJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidAddSpentTimeResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalStateWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalStateWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalStateResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsWithResponse(ctx context.Context, id int32, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidApprovalsResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveWithBodyWithResponse(ctx context.Context, id int32, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveWithResponse(ctx context.Context, id int32, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidApproveResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiWithBodyWithResponse(ctx context.Context, id int32, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiWithResponse(ctx context.Context, id int32, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiResponse, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdWithResponse request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, mergeRequestIid int32, awardId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, mergeRequestIid int32, awardId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidAwardEmojiAwardIdResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceedsWithResponse request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceedsWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidCancelMergeWhenPipelineSucceedsResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidChangesResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidClosesIssuesResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommitsWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommitsWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommitsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidCommitsResponse, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithResponse request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidContextCommitsResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidDiffsResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublishWithResponse request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublishWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesBulkPublishResponse, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithResponse request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithResponse(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithResponse(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdResponse, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdResponse, error)

	PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdWithResponse(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, body PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdResponse, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublishWithResponse request
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublishWithResponse(ctx context.Context, id string, mergeRequestIid int32, draftNoteId int32, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidDraftNotesDraftNoteIdPublishResponse, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeWithBodyWithResponse request with any body
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeResponse, error)

	PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidMergeResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRefWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRefWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidMergeRefResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiWithResponse(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiWithBodyWithResponse(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiWithResponse(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiResponse, error)

	// DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdWithResponse request
	DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, mergeRequestIid int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidNotesNoteIdAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipantsWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipantsWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidParticipantsResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidPipelinesResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidRawDiffsWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidRawDiffsWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidRawDiffsResponse, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseWithBodyWithResponse request with any body
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseResponse, error)

	PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidRebaseResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssuesWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssuesWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssuesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidRelatedIssuesResponse, error)

	// PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovalsWithResponse request
	PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovalsWithResponse(ctx context.Context, id int32, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdMergeRequestsMergeRequestIidResetApprovalsResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTimeWithResponse request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTimeWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetSpentTimeResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimateWithResponse request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimateWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidResetTimeEstimateResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewersWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewersWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidReviewersResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateWithBodyWithResponse request with any body
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateWithBodyWithResponse(ctx context.Context, id string, mergeRequestIid int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateResponse, error)

	PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateWithResponse(ctx context.Context, id string, mergeRequestIid int32, body PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidTimeEstimateResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStatsWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStatsWithResponse(ctx context.Context, id string, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidTimeStatsResponse, error)

	// PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapproveWithResponse request
	PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapproveWithResponse(ctx context.Context, id int32, mergeRequestIid int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdMergeRequestsMergeRequestIidUnapproveResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsWithResponse(ctx context.Context, id string, mergeRequestIid int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsResponse, error)

	// GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdWithResponse request
	GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdWithResponse(ctx context.Context, id string, mergeRequestIid int32, versionId int32, params *GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdMergeRequestsMergeRequestIidVersionsVersionIdResponse, error)

	// GetApiV4ProjectsIdPackagesWithResponse request
	GetApiV4ProjectsIdPackagesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesResponse, error)

	// GetApiV4ProjectsIdPackagesCargoConfigJsonWithResponse request
	GetApiV4ProjectsIdPackagesCargoConfigJsonWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesCargoConfigJsonResponse, error)

	// PostApiV4ProjectsIdPackagesComposerWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPackagesComposerWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesComposerResponse, error)

	PostApiV4ProjectsIdPackagesComposerWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdPackagesComposerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesComposerResponse, error)

	// GetApiV4ProjectsIdPackagesComposerArchivespackageNameWithResponse request
	GetApiV4ProjectsIdPackagesComposerArchivespackageNameWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesComposerArchivespackageNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesComposerArchivespackageNameResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansSearchWithResponse request
	GetApiV4ProjectsIdPackagesConanV1ConansSearchWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesConanV1ConansSearchParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1ConansSearchResponse, error)

	// DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelWithResponse request
	DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelWithResponse request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestWithResponse request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDigestResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsWithResponse request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelDownloadUrlsResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceWithResponse request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestWithResponse request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDigestResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsWithResponse request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceDownloadUrlsResponse, error)

	// PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsWithResponse request
	PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, conanPackageReference string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelPackagesConanPackageReferenceUploadUrlsResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchWithResponse request
	GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchResponse, error)

	// PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsWithResponse request
	PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesConanV1ConansPackageNamePackageVersionPackageUsernamePackageChannelUploadUrlsResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithResponse request
	GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithBodyWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameResponse, error)

	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, body PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeWithResponse request
	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionExportFileNameAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithResponse request
	GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithBodyWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameResponse, error)

	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, body PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeWithResponse request
	PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV1FilesPackageNamePackageVersionPackageUsernamePackageChannelRecipeRevisionPackageConanPackageReferencePackageRevisionFileNameAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1PingWithResponse request
	GetApiV4ProjectsIdPackagesConanV1PingWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1PingResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1UsersAuthenticateWithResponse request
	GetApiV4ProjectsIdPackagesConanV1UsersAuthenticateWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1UsersAuthenticateResponse, error)

	// GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentialsWithResponse request
	GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentialsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV1UsersCheckCredentialsResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansSearchWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansSearchWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesConanV2ConansSearchParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansSearchResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelLatestResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsResponse, error)

	// DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionWithResponse request
	DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameWithBodyWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameResponse, error)

	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, body PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorizeWithResponse request
	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorizeWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, fileName string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionFilesFileNameAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceLatestResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsResponse, error)

	// DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionWithResponse request
	DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameWithBodyWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameResponse, error)

	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, body PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorizeWithResponse request
	PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorizeWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, conanPackageReference string, packageRevision string, fileName string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionPackagesConanPackageReferenceRevisionsPackageRevisionFilesFileNameAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, recipeRevision string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelRevisionsRecipeRevisionSearchResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchWithResponse request
	GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchWithResponse(ctx context.Context, id string, packageName string, packageVersion string, packageUsername string, packageChannel string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2ConansPackageNamePackageVersionPackageUsernamePackageChannelSearchResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2UsersAuthenticateWithResponse request
	GetApiV4ProjectsIdPackagesConanV2UsersAuthenticateWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2UsersAuthenticateResponse, error)

	// GetApiV4ProjectsIdPackagesConanV2UsersCheckCredentialsWithResponse request
	GetApiV4ProjectsIdPackagesConanV2UsersCheckCredentialsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesConanV2UsersCheckCredentialsResponse, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionInreleaseWithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionInreleaseWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionInreleaseParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionInreleaseResponse, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseWithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseResponse, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpgWithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpgWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpgParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionReleaseGpgResponse, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesWithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesWithResponse(ctx context.Context, id string, component string, architecture string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitecturePackagesResponse, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256WithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256WithResponse(ctx context.Context, id string, component string, architecture string, fileSha256 int32, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentBinaryArchitectureByHashSha256FileSha256Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesWithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesWithResponse(ctx context.Context, id string, component string, architecture string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitecturePackagesResponse, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256WithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256WithResponse(ctx context.Context, id string, component string, architecture string, fileSha256 int32, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentDebianInstallerBinaryArchitectureByHashSha256FileSha256Response, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSourcesWithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSourcesWithResponse(ctx context.Context, id string, component string, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSourcesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceSourcesResponse, error)

	// GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256WithResponse request
	GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256WithResponse(ctx context.Context, id string, component string, fileSha256 int32, params *GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256Params, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianDistsdistributionComponentSourceByHashSha256FileSha256Response, error)

	// GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameWithResponse request
	GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameWithResponse(ctx context.Context, id string, distribution string, letter string, packageName string, packageVersion string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesDebianPoolDistributionLetterPackageNamePackageVersionFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesDebianFileNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesDebianFileNameWithBodyWithResponse(ctx context.Context, id string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesDebianFileNameResponse, error)

	PutApiV4ProjectsIdPackagesDebianFileNameWithResponse(ctx context.Context, id string, fileName string, body PutApiV4ProjectsIdPackagesDebianFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesDebianFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeWithBodyWithResponse(ctx context.Context, id string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeResponse, error)

	PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeWithResponse(ctx context.Context, id string, fileName string, body PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesDebianFileNameAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameWithResponse request
	GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameWithResponse(ctx context.Context, id string, packageName string, fileName string, params *GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameWithBodyWithResponse(ctx context.Context, id string, packageName string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameResponse, error)

	PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameWithResponse(ctx context.Context, id string, packageName string, fileName string, body PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeWithBodyWithResponse(ctx context.Context, id string, packageName string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeResponse, error)

	PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeWithResponse(ctx context.Context, id string, packageName string, fileName string, body PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesGenericPackageNamepackageVersionpathFileNameAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesGomoduleNameVListWithResponse request
	GetApiV4ProjectsIdPackagesGomoduleNameVListWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesGomoduleNameVListParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesGomoduleNameVListResponse, error)

	// GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionInfoWithResponse request
	GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionInfoWithResponse(ctx context.Context, id string, moduleVersion string, params *GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionInfoParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionInfoResponse, error)

	// GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionModWithResponse request
	GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionModWithResponse(ctx context.Context, id string, moduleVersion string, params *GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionModParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionModResponse, error)

	// GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionZipWithResponse request
	GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionZipWithResponse(ctx context.Context, id string, moduleVersion string, params *GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionZipParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesGomoduleNameVModuleVersionZipResponse, error)

	// PostApiV4ProjectsIdPackagesHelmApiChannelChartsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPackagesHelmApiChannelChartsWithBodyWithResponse(ctx context.Context, id int32, channel string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesHelmApiChannelChartsResponse, error)

	PostApiV4ProjectsIdPackagesHelmApiChannelChartsWithResponse(ctx context.Context, id int32, channel string, body PostApiV4ProjectsIdPackagesHelmApiChannelChartsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesHelmApiChannelChartsResponse, error)

	// PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorizeWithResponse request
	PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorizeWithResponse(ctx context.Context, id int32, channel string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesHelmApiChannelChartsAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgzWithResponse request
	GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgzWithResponse(ctx context.Context, id int32, channel string, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesHelmChannelChartsFileNameTgzResponse, error)

	// GetApiV4ProjectsIdPackagesHelmChannelIndexYamlWithResponse request
	GetApiV4ProjectsIdPackagesHelmChannelIndexYamlWithResponse(ctx context.Context, id int32, channel string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesHelmChannelIndexYamlResponse, error)

	// GetApiV4ProjectsIdPackagesMavenpathFileNameWithResponse request
	GetApiV4ProjectsIdPackagesMavenpathFileNameWithResponse(ctx context.Context, id string, fileName string, params *GetApiV4ProjectsIdPackagesMavenpathFileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesMavenpathFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesMavenpathFileNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesMavenpathFileNameWithBodyWithResponse(ctx context.Context, id string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesMavenpathFileNameResponse, error)

	PutApiV4ProjectsIdPackagesMavenpathFileNameWithResponse(ctx context.Context, id string, fileName string, body PutApiV4ProjectsIdPackagesMavenpathFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesMavenpathFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeWithBodyWithResponse(ctx context.Context, id string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeResponse, error)

	PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeWithResponse(ctx context.Context, id string, fileName string, body PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesMavenpathFileNameAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameWithResponse request
	GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameWithResponse(ctx context.Context, id string, modelVersionId string, fileName string, params *GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameWithBodyWithResponse(ctx context.Context, id string, modelVersionId string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameResponse, error)

	PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameWithResponse(ctx context.Context, id string, modelVersionId string, fileName string, body PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeWithBodyWithResponse(ctx context.Context, id string, modelVersionId string, fileName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeResponse, error)

	PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeWithResponse(ctx context.Context, id string, modelVersionId string, fileName string, body PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesMlModelsModelVersionIdFilespathFileNameAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesNpmpackageNameWithResponse request
	GetApiV4ProjectsIdPackagesNpmpackageNameWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNpmpackageNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNpmpackageNameResponse, error)

	// GetApiV4ProjectsIdPackagesNpmpackageNamefileNameWithResponse request
	GetApiV4ProjectsIdPackagesNpmpackageNamefileNameWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNpmpackageNamefileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNpmpackageNamefileNameResponse, error)

	// PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulkWithResponse request
	PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulkWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAdvisoriesBulkResponse, error)

	// PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuickWithResponse request
	PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuickWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesNpmNpmV1SecurityAuditsQuickResponse, error)

	// GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsWithResponse request
	GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsResponse, error)

	// DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagWithResponse request
	DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagWithResponse(ctx context.Context, id string, tag string, params *DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagResponse, error)

	// PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagWithBodyWithResponse(ctx context.Context, id string, tag string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagResponse, error)

	PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagWithResponse(ctx context.Context, id string, tag string, body PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNpmPackagepackageNameDistTagsTagResponse, error)

	// PutApiV4ProjectsIdPackagesNpmPackageNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesNpmPackageNameWithBodyWithResponse(ctx context.Context, id string, packageName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNpmPackageNameResponse, error)

	PutApiV4ProjectsIdPackagesNpmPackageNameWithResponse(ctx context.Context, id string, packageName string, body PutApiV4ProjectsIdPackagesNpmPackageNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNpmPackageNameResponse, error)

	// PutApiV4ProjectsIdPackagesNugetWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesNugetWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetResponse, error)

	PutApiV4ProjectsIdPackagesNugetWithResponse(ctx context.Context, id string, body PutApiV4ProjectsIdPackagesNugetJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetResponse, error)

	// DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersionWithResponse request
	DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersionWithResponse(ctx context.Context, id string, params *DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersionParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPackagesNugetpackageNamepackageVersionResponse, error)

	// PutApiV4ProjectsIdPackagesNugetAuthorizeWithResponse request
	PutApiV4ProjectsIdPackagesNugetAuthorizeWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilenameWithResponse request
	GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilenameWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilenameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNugetDownloadpackageNamepackageVersionpackageFilenameResponse, error)

	// GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndexWithResponse request
	GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndexWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndexParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNugetDownloadpackageNameIndexResponse, error)

	// GetApiV4ProjectsIdPackagesNugetIndexWithResponse request
	GetApiV4ProjectsIdPackagesNugetIndexWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNugetIndexResponse, error)

	// GetApiV4ProjectsIdPackagesNugetQueryWithResponse request
	GetApiV4ProjectsIdPackagesNugetQueryWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNugetQueryParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNugetQueryResponse, error)

	// GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameWithResponse request
	GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNugetSymbolfilesfileNamesignaturesameFileNameResponse, error)

	// PutApiV4ProjectsIdPackagesNugetSymbolpackageWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesNugetSymbolpackageWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetSymbolpackageResponse, error)

	PutApiV4ProjectsIdPackagesNugetSymbolpackageWithResponse(ctx context.Context, id string, body PutApiV4ProjectsIdPackagesNugetSymbolpackageJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetSymbolpackageResponse, error)

	// PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorizeWithResponse request
	PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorizeWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetSymbolpackageAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesNugetV2WithResponse request
	GetApiV4ProjectsIdPackagesNugetV2WithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNugetV2Response, error)

	// PutApiV4ProjectsIdPackagesNugetV2WithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesNugetV2WithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetV2Response, error)

	PutApiV4ProjectsIdPackagesNugetV2WithResponse(ctx context.Context, id string, body PutApiV4ProjectsIdPackagesNugetV2JSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetV2Response, error)

	// GetApiV4ProjectsIdPackagesNugetV2MetadataWithResponse request
	GetApiV4ProjectsIdPackagesNugetV2MetadataWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesNugetV2MetadataResponse, error)

	// PutApiV4ProjectsIdPackagesNugetV2AuthorizeWithResponse request
	PutApiV4ProjectsIdPackagesNugetV2AuthorizeWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesNugetV2AuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesProtectionRulesWithResponse request
	GetApiV4ProjectsIdPackagesProtectionRulesWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesProtectionRulesResponse, error)

	// PostApiV4ProjectsIdPackagesProtectionRulesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPackagesProtectionRulesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesProtectionRulesResponse, error)

	PostApiV4ProjectsIdPackagesProtectionRulesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdPackagesProtectionRulesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesProtectionRulesResponse, error)

	// DeleteApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdWithResponse request
	DeleteApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdWithResponse(ctx context.Context, id string, packageProtectionRuleId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdResponse, error)

	// PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdWithBodyWithResponse request with any body
	PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdWithBodyWithResponse(ctx context.Context, id string, packageProtectionRuleId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdResponse, error)

	PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdWithResponse(ctx context.Context, id string, packageProtectionRuleId int32, body PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdPackagesProtectionRulesPackageProtectionRuleIdResponse, error)

	// PostApiV4ProjectsIdPackagesPypiWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPackagesPypiWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesPypiResponse, error)

	PostApiV4ProjectsIdPackagesPypiWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdPackagesPypiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesPypiResponse, error)

	// PostApiV4ProjectsIdPackagesPypiAuthorizeWithResponse request
	PostApiV4ProjectsIdPackagesPypiAuthorizeWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesPypiAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifierWithResponse request
	GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifierWithResponse(ctx context.Context, id string, sha256 string, params *GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifierParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesPypiFilesSha256fileIdentifierResponse, error)

	// GetApiV4ProjectsIdPackagesPypiSimpleWithResponse request
	GetApiV4ProjectsIdPackagesPypiSimpleWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesPypiSimpleResponse, error)

	// GetApiV4ProjectsIdPackagesPypiSimplepackageNameWithResponse request
	GetApiV4ProjectsIdPackagesPypiSimplepackageNameWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesPypiSimplepackageNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesPypiSimplepackageNameResponse, error)

	// PostApiV4ProjectsIdPackagesRpmWithResponse request
	PostApiV4ProjectsIdPackagesRpmWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesRpmResponse, error)

	// GetApiV4ProjectsIdPackagesRpmpackageFileIdfileNameWithResponse request
	GetApiV4ProjectsIdPackagesRpmpackageFileIdfileNameWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesRpmpackageFileIdfileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesRpmpackageFileIdfileNameResponse, error)

	// PostApiV4ProjectsIdPackagesRpmAuthorizeWithResponse request
	PostApiV4ProjectsIdPackagesRpmAuthorizeWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesRpmAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesRpmRepodatafileNameWithResponse request
	GetApiV4ProjectsIdPackagesRpmRepodatafileNameWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPackagesRpmRepodatafileNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesRpmRepodatafileNameResponse, error)

	// GetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesWithResponse request
	GetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesRubygemsApiV1DependenciesResponse, error)

	// PostApiV4ProjectsIdPackagesRubygemsApiV1GemsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPackagesRubygemsApiV1GemsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesRubygemsApiV1GemsResponse, error)

	PostApiV4ProjectsIdPackagesRubygemsApiV1GemsWithResponse(ctx context.Context, id int32, body PostApiV4ProjectsIdPackagesRubygemsApiV1GemsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesRubygemsApiV1GemsResponse, error)

	// PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorizeWithResponse request
	PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorizeWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPackagesRubygemsApiV1GemsAuthorizeResponse, error)

	// GetApiV4ProjectsIdPackagesRubygemsGemsFileNameWithResponse request
	GetApiV4ProjectsIdPackagesRubygemsGemsFileNameWithResponse(ctx context.Context, id int32, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesRubygemsGemsFileNameResponse, error)

	// GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileNameWithResponse request
	GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileNameWithResponse(ctx context.Context, id int32, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesRubygemsQuickMarshal48FileNameResponse, error)

	// GetApiV4ProjectsIdPackagesRubygemsFileNameWithResponse request
	GetApiV4ProjectsIdPackagesRubygemsFileNameWithResponse(ctx context.Context, id int32, fileName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesRubygemsFileNameResponse, error)

	// GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemWithResponse request
	GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemWithResponse(ctx context.Context, id string, moduleName string, moduleSystem string, params *GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemResponse, error)

	// GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionWithResponse request
	GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionWithResponse(ctx context.Context, id string, moduleName string, moduleSystem string, params *GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionResponse, error)

	// PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileWithBodyWithResponse(ctx context.Context, id string, moduleName string, moduleSystem string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileResponse, error)

	// PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeWithBodyWithResponse(ctx context.Context, id string, moduleName string, moduleSystem string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeResponse, error)

	PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeWithResponse(ctx context.Context, id string, moduleName string, moduleSystem string, body PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPackagesTerraformModulesModuleNameModuleSystemmoduleVersionFileAuthorizeResponse, error)

	// DeleteApiV4ProjectsIdPackagesPackageIdWithResponse request
	DeleteApiV4ProjectsIdPackagesPackageIdWithResponse(ctx context.Context, id string, packageId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPackagesPackageIdResponse, error)

	// GetApiV4ProjectsIdPackagesPackageIdWithResponse request
	GetApiV4ProjectsIdPackagesPackageIdWithResponse(ctx context.Context, id string, packageId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesPackageIdResponse, error)

	// GetApiV4ProjectsIdPackagesPackageIdPackageFilesWithResponse request
	GetApiV4ProjectsIdPackagesPackageIdPackageFilesWithResponse(ctx context.Context, id string, packageId int32, params *GetApiV4ProjectsIdPackagesPackageIdPackageFilesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesPackageIdPackageFilesResponse, error)

	// DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileIdWithResponse request
	DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileIdWithResponse(ctx context.Context, id string, packageId int32, packageFileId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPackagesPackageIdPackageFilesPackageFileIdResponse, error)

	// GetApiV4ProjectsIdPackagesPackageIdPipelinesWithResponse request
	GetApiV4ProjectsIdPackagesPackageIdPipelinesWithResponse(ctx context.Context, id string, packageId int32, params *GetApiV4ProjectsIdPackagesPackageIdPipelinesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPackagesPackageIdPipelinesResponse, error)

	// DeleteApiV4ProjectsIdPagesWithResponse request
	DeleteApiV4ProjectsIdPagesWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPagesResponse, error)

	// GetApiV4ProjectsIdPagesWithResponse request
	GetApiV4ProjectsIdPagesWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPagesResponse, error)

	// PatchApiV4ProjectsIdPagesWithBodyWithResponse request with any body
	PatchApiV4ProjectsIdPagesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdPagesResponse, error)

	PatchApiV4ProjectsIdPagesWithResponse(ctx context.Context, id string, body PatchApiV4ProjectsIdPagesJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdPagesResponse, error)

	// GetApiV4ProjectsIdPagesDomainsWithResponse request
	GetApiV4ProjectsIdPagesDomainsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPagesDomainsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPagesDomainsResponse, error)

	// PostApiV4ProjectsIdPagesDomainsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPagesDomainsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPagesDomainsResponse, error)

	PostApiV4ProjectsIdPagesDomainsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdPagesDomainsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPagesDomainsResponse, error)

	// DeleteApiV4ProjectsIdPagesDomainsDomainWithResponse request
	DeleteApiV4ProjectsIdPagesDomainsDomainWithResponse(ctx context.Context, id string, domain string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPagesDomainsDomainResponse, error)

	// GetApiV4ProjectsIdPagesDomainsDomainWithResponse request
	GetApiV4ProjectsIdPagesDomainsDomainWithResponse(ctx context.Context, id string, domain string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPagesDomainsDomainResponse, error)

	// PutApiV4ProjectsIdPagesDomainsDomainWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPagesDomainsDomainWithBodyWithResponse(ctx context.Context, id string, domain string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPagesDomainsDomainResponse, error)

	PutApiV4ProjectsIdPagesDomainsDomainWithResponse(ctx context.Context, id string, domain string, body PutApiV4ProjectsIdPagesDomainsDomainJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPagesDomainsDomainResponse, error)

	// PutApiV4ProjectsIdPagesDomainsDomainVerifyWithResponse request
	PutApiV4ProjectsIdPagesDomainsDomainVerifyWithResponse(ctx context.Context, id string, domain string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPagesDomainsDomainVerifyResponse, error)

	// GetApiV4ProjectsIdPagesAccessWithResponse request
	GetApiV4ProjectsIdPagesAccessWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPagesAccessResponse, error)

	// PostApiV4ProjectsIdPipelineWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPipelineWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelineResponse, error)

	PostApiV4ProjectsIdPipelineWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdPipelineJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelineResponse, error)

	// GetApiV4ProjectsIdPipelineSchedulesWithResponse request
	GetApiV4ProjectsIdPipelineSchedulesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPipelineSchedulesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelineSchedulesResponse, error)

	// PostApiV4ProjectsIdPipelineSchedulesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPipelineSchedulesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelineSchedulesResponse, error)

	PostApiV4ProjectsIdPipelineSchedulesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdPipelineSchedulesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelineSchedulesResponse, error)

	// DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithResponse request
	DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithResponse(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdResponse, error)

	// GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithResponse request
	GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithResponse(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdResponse, error)

	// PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithBodyWithResponse(ctx context.Context, id string, pipelineScheduleId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdResponse, error)

	PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdWithResponse(ctx context.Context, id string, pipelineScheduleId int32, body PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdResponse, error)

	// GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesWithResponse request
	GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesWithResponse(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPipelinesResponse, error)

	// PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlayWithResponse request
	PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlayWithResponse(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdPlayResponse, error)

	// PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnershipWithResponse request
	PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnershipWithResponse(ctx context.Context, id string, pipelineScheduleId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdTakeOwnershipResponse, error)

	// PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesWithBodyWithResponse(ctx context.Context, id string, pipelineScheduleId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesResponse, error)

	PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesWithResponse(ctx context.Context, id string, pipelineScheduleId int32, body PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesResponse, error)

	// DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyWithResponse request
	DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyWithResponse(ctx context.Context, id string, pipelineScheduleId int32, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyResponse, error)

	// PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyWithBodyWithResponse(ctx context.Context, id string, pipelineScheduleId int32, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyResponse, error)

	PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyWithResponse(ctx context.Context, id string, pipelineScheduleId int32, key string, body PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPipelineSchedulesPipelineScheduleIdVariablesKeyResponse, error)

	// GetApiV4ProjectsIdPipelinesWithResponse request
	GetApiV4ProjectsIdPipelinesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPipelinesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelinesResponse, error)

	// GetApiV4ProjectsIdPipelinesLatestWithResponse request
	GetApiV4ProjectsIdPipelinesLatestWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdPipelinesLatestParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelinesLatestResponse, error)

	// DeleteApiV4ProjectsIdPipelinesPipelineIdWithResponse request
	DeleteApiV4ProjectsIdPipelinesPipelineIdWithResponse(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdPipelinesPipelineIdResponse, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdWithResponse request
	GetApiV4ProjectsIdPipelinesPipelineIdWithResponse(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelinesPipelineIdResponse, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdBridgesWithResponse request
	GetApiV4ProjectsIdPipelinesPipelineIdBridgesWithResponse(ctx context.Context, id string, pipelineId int32, params *GetApiV4ProjectsIdPipelinesPipelineIdBridgesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelinesPipelineIdBridgesResponse, error)

	// PostApiV4ProjectsIdPipelinesPipelineIdCancelWithResponse request
	PostApiV4ProjectsIdPipelinesPipelineIdCancelWithResponse(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelinesPipelineIdCancelResponse, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdJobsWithResponse request
	GetApiV4ProjectsIdPipelinesPipelineIdJobsWithResponse(ctx context.Context, id string, pipelineId int32, params *GetApiV4ProjectsIdPipelinesPipelineIdJobsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelinesPipelineIdJobsResponse, error)

	// PutApiV4ProjectsIdPipelinesPipelineIdMetadataWithBodyWithResponse request with any body
	PutApiV4ProjectsIdPipelinesPipelineIdMetadataWithBodyWithResponse(ctx context.Context, id string, pipelineId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPipelinesPipelineIdMetadataResponse, error)

	PutApiV4ProjectsIdPipelinesPipelineIdMetadataWithResponse(ctx context.Context, id string, pipelineId int32, body PutApiV4ProjectsIdPipelinesPipelineIdMetadataJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdPipelinesPipelineIdMetadataResponse, error)

	// PostApiV4ProjectsIdPipelinesPipelineIdRetryWithResponse request
	PostApiV4ProjectsIdPipelinesPipelineIdRetryWithResponse(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdPipelinesPipelineIdRetryResponse, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdTestReportWithResponse request
	GetApiV4ProjectsIdPipelinesPipelineIdTestReportWithResponse(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelinesPipelineIdTestReportResponse, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummaryWithResponse request
	GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummaryWithResponse(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelinesPipelineIdTestReportSummaryResponse, error)

	// GetApiV4ProjectsIdPipelinesPipelineIdVariablesWithResponse request
	GetApiV4ProjectsIdPipelinesPipelineIdVariablesWithResponse(ctx context.Context, id string, pipelineId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdPipelinesPipelineIdVariablesResponse, error)

	// GetApiV4ProjectsIdProtectedBranchesWithResponse request
	GetApiV4ProjectsIdProtectedBranchesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdProtectedBranchesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdProtectedBranchesResponse, error)

	// PostApiV4ProjectsIdProtectedBranchesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdProtectedBranchesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdProtectedBranchesResponse, error)

	PostApiV4ProjectsIdProtectedBranchesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdProtectedBranchesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdProtectedBranchesResponse, error)

	// DeleteApiV4ProjectsIdProtectedBranchesNameWithResponse request
	DeleteApiV4ProjectsIdProtectedBranchesNameWithResponse(ctx context.Context, id string, name string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdProtectedBranchesNameResponse, error)

	// GetApiV4ProjectsIdProtectedBranchesNameWithResponse request
	GetApiV4ProjectsIdProtectedBranchesNameWithResponse(ctx context.Context, id string, name string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdProtectedBranchesNameResponse, error)

	// PatchApiV4ProjectsIdProtectedBranchesNameWithBodyWithResponse request with any body
	PatchApiV4ProjectsIdProtectedBranchesNameWithBodyWithResponse(ctx context.Context, id string, name string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdProtectedBranchesNameResponse, error)

	PatchApiV4ProjectsIdProtectedBranchesNameWithResponse(ctx context.Context, id string, name string, body PatchApiV4ProjectsIdProtectedBranchesNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdProtectedBranchesNameResponse, error)

	// GetApiV4ProjectsIdProtectedTagsWithResponse request
	GetApiV4ProjectsIdProtectedTagsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdProtectedTagsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdProtectedTagsResponse, error)

	// PostApiV4ProjectsIdProtectedTagsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdProtectedTagsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdProtectedTagsResponse, error)

	PostApiV4ProjectsIdProtectedTagsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdProtectedTagsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdProtectedTagsResponse, error)

	// DeleteApiV4ProjectsIdProtectedTagsNameWithResponse request
	DeleteApiV4ProjectsIdProtectedTagsNameWithResponse(ctx context.Context, id string, name string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdProtectedTagsNameResponse, error)

	// GetApiV4ProjectsIdProtectedTagsNameWithResponse request
	GetApiV4ProjectsIdProtectedTagsNameWithResponse(ctx context.Context, id string, name string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdProtectedTagsNameResponse, error)

	// GetApiV4ProjectsIdRegistryProtectionRepositoryRulesWithResponse request
	GetApiV4ProjectsIdRegistryProtectionRepositoryRulesWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRegistryProtectionRepositoryRulesResponse, error)

	// PostApiV4ProjectsIdRegistryProtectionRepositoryRulesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRegistryProtectionRepositoryRulesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRegistryProtectionRepositoryRulesResponse, error)

	PostApiV4ProjectsIdRegistryProtectionRepositoryRulesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdRegistryProtectionRepositoryRulesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRegistryProtectionRepositoryRulesResponse, error)

	// DeleteApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdWithResponse request
	DeleteApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdWithResponse(ctx context.Context, id string, protectionRuleId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdResponse, error)

	// PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdWithBodyWithResponse request with any body
	PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdWithBodyWithResponse(ctx context.Context, id string, protectionRuleId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdResponse, error)

	PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdWithResponse(ctx context.Context, id string, protectionRuleId int32, body PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchApiV4ProjectsIdRegistryProtectionRepositoryRulesProtectionRuleIdResponse, error)

	// GetApiV4ProjectsIdRegistryRepositoriesWithResponse request
	GetApiV4ProjectsIdRegistryRepositoriesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRegistryRepositoriesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRegistryRepositoriesResponse, error)

	// DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdWithResponse request
	DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdWithResponse(ctx context.Context, id string, repositoryId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdResponse, error)

	// DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsWithResponse request
	DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsWithResponse(ctx context.Context, id string, repositoryId int32, params *DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsResponse, error)

	// GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsWithResponse request
	GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsWithResponse(ctx context.Context, id string, repositoryId int32, params *GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsResponse, error)

	// DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameWithResponse request
	DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameWithResponse(ctx context.Context, id string, repositoryId int32, tagName string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameResponse, error)

	// GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameWithResponse request
	GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameWithResponse(ctx context.Context, id string, repositoryId int32, tagName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRegistryRepositoriesRepositoryIdTagsTagNameResponse, error)

	// GetApiV4ProjectsIdRelationImportsWithResponse request
	GetApiV4ProjectsIdRelationImportsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRelationImportsResponse, error)

	// GetApiV4ProjectsIdReleasesWithResponse request
	GetApiV4ProjectsIdReleasesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdReleasesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdReleasesResponse, error)

	// PostApiV4ProjectsIdReleasesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdReleasesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdReleasesResponse, error)

	PostApiV4ProjectsIdReleasesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdReleasesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdReleasesResponse, error)

	// GetApiV4ProjectsIdReleasesPermalinkLatestsuffixPathWithResponse request
	GetApiV4ProjectsIdReleasesPermalinkLatestsuffixPathWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdReleasesPermalinkLatestsuffixPathParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdReleasesPermalinkLatestsuffixPathResponse, error)

	// DeleteApiV4ProjectsIdReleasesTagNameWithResponse request
	DeleteApiV4ProjectsIdReleasesTagNameWithResponse(ctx context.Context, id string, tagName string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdReleasesTagNameResponse, error)

	// GetApiV4ProjectsIdReleasesTagNameWithResponse request
	GetApiV4ProjectsIdReleasesTagNameWithResponse(ctx context.Context, id string, tagName string, params *GetApiV4ProjectsIdReleasesTagNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdReleasesTagNameResponse, error)

	// PutApiV4ProjectsIdReleasesTagNameWithBodyWithResponse request with any body
	PutApiV4ProjectsIdReleasesTagNameWithBodyWithResponse(ctx context.Context, id string, tagName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdReleasesTagNameResponse, error)

	PutApiV4ProjectsIdReleasesTagNameWithResponse(ctx context.Context, id string, tagName string, body PutApiV4ProjectsIdReleasesTagNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdReleasesTagNameResponse, error)

	// GetApiV4ProjectsIdReleasesTagNameAssetsLinksWithResponse request
	GetApiV4ProjectsIdReleasesTagNameAssetsLinksWithResponse(ctx context.Context, id string, tagName string, params *GetApiV4ProjectsIdReleasesTagNameAssetsLinksParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdReleasesTagNameAssetsLinksResponse, error)

	// PostApiV4ProjectsIdReleasesTagNameAssetsLinksWithBodyWithResponse request with any body
	PostApiV4ProjectsIdReleasesTagNameAssetsLinksWithBodyWithResponse(ctx context.Context, id string, tagName string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdReleasesTagNameAssetsLinksResponse, error)

	PostApiV4ProjectsIdReleasesTagNameAssetsLinksWithResponse(ctx context.Context, id string, tagName string, body PostApiV4ProjectsIdReleasesTagNameAssetsLinksJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdReleasesTagNameAssetsLinksResponse, error)

	// DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithResponse request
	DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithResponse(ctx context.Context, id string, tagName string, linkId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdResponse, error)

	// GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithResponse request
	GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithResponse(ctx context.Context, id string, tagName string, linkId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdResponse, error)

	// PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithBodyWithResponse(ctx context.Context, id string, tagName string, linkId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdResponse, error)

	PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdWithResponse(ctx context.Context, id string, tagName string, linkId int32, body PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdReleasesTagNameAssetsLinksLinkIdResponse, error)

	// GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPathWithResponse request
	GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPathWithResponse(ctx context.Context, id string, tagName string, params *GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPathParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdReleasesTagNameDownloadsdirectAssetPathResponse, error)

	// PostApiV4ProjectsIdReleasesTagNameEvidenceWithResponse request
	PostApiV4ProjectsIdReleasesTagNameEvidenceWithResponse(ctx context.Context, id int32, tagName string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdReleasesTagNameEvidenceResponse, error)

	// GetApiV4ProjectsIdRemoteMirrorsWithResponse request
	GetApiV4ProjectsIdRemoteMirrorsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRemoteMirrorsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRemoteMirrorsResponse, error)

	// PostApiV4ProjectsIdRemoteMirrorsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRemoteMirrorsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRemoteMirrorsResponse, error)

	PostApiV4ProjectsIdRemoteMirrorsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdRemoteMirrorsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRemoteMirrorsResponse, error)

	// DeleteApiV4ProjectsIdRemoteMirrorsMirrorIdWithResponse request
	DeleteApiV4ProjectsIdRemoteMirrorsMirrorIdWithResponse(ctx context.Context, id string, mirrorId string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRemoteMirrorsMirrorIdResponse, error)

	// GetApiV4ProjectsIdRemoteMirrorsMirrorIdWithResponse request
	GetApiV4ProjectsIdRemoteMirrorsMirrorIdWithResponse(ctx context.Context, id string, mirrorId string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRemoteMirrorsMirrorIdResponse, error)

	// PutApiV4ProjectsIdRemoteMirrorsMirrorIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdRemoteMirrorsMirrorIdWithBodyWithResponse(ctx context.Context, id string, mirrorId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRemoteMirrorsMirrorIdResponse, error)

	PutApiV4ProjectsIdRemoteMirrorsMirrorIdWithResponse(ctx context.Context, id string, mirrorId string, body PutApiV4ProjectsIdRemoteMirrorsMirrorIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRemoteMirrorsMirrorIdResponse, error)

	// GetApiV4ProjectsIdRemoteMirrorsMirrorIdPublicKeyWithResponse request
	GetApiV4ProjectsIdRemoteMirrorsMirrorIdPublicKeyWithResponse(ctx context.Context, id string, mirrorId string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRemoteMirrorsMirrorIdPublicKeyResponse, error)

	// PostApiV4ProjectsIdRemoteMirrorsMirrorIdSyncWithResponse request
	PostApiV4ProjectsIdRemoteMirrorsMirrorIdSyncWithResponse(ctx context.Context, id string, mirrorId string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRemoteMirrorsMirrorIdSyncResponse, error)

	// GetApiV4ProjectsIdRepositoryArchiveWithResponse request
	GetApiV4ProjectsIdRepositoryArchiveWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryArchiveParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryArchiveResponse, error)

	// GetApiV4ProjectsIdRepositoryBlobsShaWithResponse request
	GetApiV4ProjectsIdRepositoryBlobsShaWithResponse(ctx context.Context, id string, sha string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryBlobsShaResponse, error)

	// GetApiV4ProjectsIdRepositoryBlobsShaRawWithResponse request
	GetApiV4ProjectsIdRepositoryBlobsShaRawWithResponse(ctx context.Context, id string, sha string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryBlobsShaRawResponse, error)

	// GetApiV4ProjectsIdRepositoryBranchesWithResponse request
	GetApiV4ProjectsIdRepositoryBranchesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryBranchesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryBranchesResponse, error)

	// PostApiV4ProjectsIdRepositoryBranchesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRepositoryBranchesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryBranchesResponse, error)

	PostApiV4ProjectsIdRepositoryBranchesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdRepositoryBranchesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryBranchesResponse, error)

	// DeleteApiV4ProjectsIdRepositoryBranchesBranchWithResponse request
	DeleteApiV4ProjectsIdRepositoryBranchesBranchWithResponse(ctx context.Context, id string, branch string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRepositoryBranchesBranchResponse, error)

	// GetApiV4ProjectsIdRepositoryBranchesBranchWithResponse request
	GetApiV4ProjectsIdRepositoryBranchesBranchWithResponse(ctx context.Context, id string, branch int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryBranchesBranchResponse, error)

	// HeadApiV4ProjectsIdRepositoryBranchesBranchWithResponse request
	HeadApiV4ProjectsIdRepositoryBranchesBranchWithResponse(ctx context.Context, id string, branch string, reqEditors ...RequestEditorFn) (*HeadApiV4ProjectsIdRepositoryBranchesBranchResponse, error)

	// PutApiV4ProjectsIdRepositoryBranchesBranchProtectWithBodyWithResponse request with any body
	PutApiV4ProjectsIdRepositoryBranchesBranchProtectWithBodyWithResponse(ctx context.Context, id string, branch string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRepositoryBranchesBranchProtectResponse, error)

	PutApiV4ProjectsIdRepositoryBranchesBranchProtectWithResponse(ctx context.Context, id string, branch string, body PutApiV4ProjectsIdRepositoryBranchesBranchProtectJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRepositoryBranchesBranchProtectResponse, error)

	// PutApiV4ProjectsIdRepositoryBranchesBranchUnprotectWithResponse request
	PutApiV4ProjectsIdRepositoryBranchesBranchUnprotectWithResponse(ctx context.Context, id string, branch string, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRepositoryBranchesBranchUnprotectResponse, error)

	// GetApiV4ProjectsIdRepositoryChangelogWithResponse request
	GetApiV4ProjectsIdRepositoryChangelogWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryChangelogParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryChangelogResponse, error)

	// PostApiV4ProjectsIdRepositoryChangelogWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRepositoryChangelogWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryChangelogResponse, error)

	PostApiV4ProjectsIdRepositoryChangelogWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdRepositoryChangelogJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryChangelogResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryCommitsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsResponse, error)

	// PostApiV4ProjectsIdRepositoryCommitsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRepositoryCommitsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryCommitsResponse, error)

	PostApiV4ProjectsIdRepositoryCommitsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdRepositoryCommitsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryCommitsResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsShaWithResponse(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsShaResponse, error)

	// PostApiV4ProjectsIdRepositoryCommitsShaCherryPickWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRepositoryCommitsShaCherryPickWithBodyWithResponse(ctx context.Context, id string, sha string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryCommitsShaCherryPickResponse, error)

	PostApiV4ProjectsIdRepositoryCommitsShaCherryPickWithResponse(ctx context.Context, id string, sha string, body PostApiV4ProjectsIdRepositoryCommitsShaCherryPickJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryCommitsShaCherryPickResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaCommentsWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsShaCommentsWithResponse(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaCommentsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsShaCommentsResponse, error)

	// PostApiV4ProjectsIdRepositoryCommitsShaCommentsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRepositoryCommitsShaCommentsWithBodyWithResponse(ctx context.Context, id string, sha string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryCommitsShaCommentsResponse, error)

	PostApiV4ProjectsIdRepositoryCommitsShaCommentsWithResponse(ctx context.Context, id string, sha string, body PostApiV4ProjectsIdRepositoryCommitsShaCommentsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryCommitsShaCommentsResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaDiffWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsShaDiffWithResponse(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaDiffParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsShaDiffResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsWithResponse(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsShaMergeRequestsResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaRefsWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsShaRefsWithResponse(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaRefsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsShaRefsResponse, error)

	// PostApiV4ProjectsIdRepositoryCommitsShaRevertWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRepositoryCommitsShaRevertWithBodyWithResponse(ctx context.Context, id string, sha string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryCommitsShaRevertResponse, error)

	PostApiV4ProjectsIdRepositoryCommitsShaRevertWithResponse(ctx context.Context, id string, sha string, body PostApiV4ProjectsIdRepositoryCommitsShaRevertJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryCommitsShaRevertResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaSequenceWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsShaSequenceWithResponse(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaSequenceParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsShaSequenceResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaSignatureWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsShaSignatureWithResponse(ctx context.Context, id string, sha string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsShaSignatureResponse, error)

	// GetApiV4ProjectsIdRepositoryCommitsShaStatusesWithResponse request
	GetApiV4ProjectsIdRepositoryCommitsShaStatusesWithResponse(ctx context.Context, id string, sha string, params *GetApiV4ProjectsIdRepositoryCommitsShaStatusesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCommitsShaStatusesResponse, error)

	// GetApiV4ProjectsIdRepositoryCompareWithResponse request
	GetApiV4ProjectsIdRepositoryCompareWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryCompareParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryCompareResponse, error)

	// GetApiV4ProjectsIdRepositoryContributorsWithResponse request
	GetApiV4ProjectsIdRepositoryContributorsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryContributorsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryContributorsResponse, error)

	// DeleteApiV4ProjectsIdRepositoryFilesFilePathWithResponse request
	DeleteApiV4ProjectsIdRepositoryFilesFilePathWithResponse(ctx context.Context, id string, filePath string, params *DeleteApiV4ProjectsIdRepositoryFilesFilePathParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRepositoryFilesFilePathResponse, error)

	// GetApiV4ProjectsIdRepositoryFilesFilePathWithResponse request
	GetApiV4ProjectsIdRepositoryFilesFilePathWithResponse(ctx context.Context, id string, filePath string, params *GetApiV4ProjectsIdRepositoryFilesFilePathParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryFilesFilePathResponse, error)

	// HeadApiV4ProjectsIdRepositoryFilesFilePathWithResponse request
	HeadApiV4ProjectsIdRepositoryFilesFilePathWithResponse(ctx context.Context, id string, filePath string, params *HeadApiV4ProjectsIdRepositoryFilesFilePathParams, reqEditors ...RequestEditorFn) (*HeadApiV4ProjectsIdRepositoryFilesFilePathResponse, error)

	// PostApiV4ProjectsIdRepositoryFilesFilePathWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRepositoryFilesFilePathWithBodyWithResponse(ctx context.Context, id string, filePath string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryFilesFilePathResponse, error)

	PostApiV4ProjectsIdRepositoryFilesFilePathWithResponse(ctx context.Context, id string, filePath string, body PostApiV4ProjectsIdRepositoryFilesFilePathJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryFilesFilePathResponse, error)

	// PutApiV4ProjectsIdRepositoryFilesFilePathWithBodyWithResponse request with any body
	PutApiV4ProjectsIdRepositoryFilesFilePathWithBodyWithResponse(ctx context.Context, id string, filePath string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRepositoryFilesFilePathResponse, error)

	PutApiV4ProjectsIdRepositoryFilesFilePathWithResponse(ctx context.Context, id string, filePath string, body PutApiV4ProjectsIdRepositoryFilesFilePathJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRepositoryFilesFilePathResponse, error)

	// GetApiV4ProjectsIdRepositoryFilesFilePathBlameWithResponse request
	GetApiV4ProjectsIdRepositoryFilesFilePathBlameWithResponse(ctx context.Context, id string, filePath string, params *GetApiV4ProjectsIdRepositoryFilesFilePathBlameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryFilesFilePathBlameResponse, error)

	// HeadApiV4ProjectsIdRepositoryFilesFilePathBlameWithResponse request
	HeadApiV4ProjectsIdRepositoryFilesFilePathBlameWithResponse(ctx context.Context, id string, filePath string, params *HeadApiV4ProjectsIdRepositoryFilesFilePathBlameParams, reqEditors ...RequestEditorFn) (*HeadApiV4ProjectsIdRepositoryFilesFilePathBlameResponse, error)

	// GetApiV4ProjectsIdRepositoryFilesFilePathRawWithResponse request
	GetApiV4ProjectsIdRepositoryFilesFilePathRawWithResponse(ctx context.Context, id string, filePath string, params *GetApiV4ProjectsIdRepositoryFilesFilePathRawParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryFilesFilePathRawResponse, error)

	// GetApiV4ProjectsIdRepositoryHealthWithResponse request
	GetApiV4ProjectsIdRepositoryHealthWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryHealthParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryHealthResponse, error)

	// GetApiV4ProjectsIdRepositoryMergeBaseWithResponse request
	GetApiV4ProjectsIdRepositoryMergeBaseWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryMergeBaseParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryMergeBaseResponse, error)

	// DeleteApiV4ProjectsIdRepositoryMergedBranchesWithResponse request
	DeleteApiV4ProjectsIdRepositoryMergedBranchesWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRepositoryMergedBranchesResponse, error)

	// PutApiV4ProjectsIdRepositorySubmodulesSubmoduleWithBodyWithResponse request with any body
	PutApiV4ProjectsIdRepositorySubmodulesSubmoduleWithBodyWithResponse(ctx context.Context, id string, submodule string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRepositorySubmodulesSubmoduleResponse, error)

	PutApiV4ProjectsIdRepositorySubmodulesSubmoduleWithResponse(ctx context.Context, id string, submodule string, body PutApiV4ProjectsIdRepositorySubmodulesSubmoduleJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdRepositorySubmodulesSubmoduleResponse, error)

	// GetApiV4ProjectsIdRepositoryTagsWithResponse request
	GetApiV4ProjectsIdRepositoryTagsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryTagsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryTagsResponse, error)

	// PostApiV4ProjectsIdRepositoryTagsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRepositoryTagsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryTagsResponse, error)

	PostApiV4ProjectsIdRepositoryTagsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdRepositoryTagsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositoryTagsResponse, error)

	// DeleteApiV4ProjectsIdRepositoryTagsTagNameWithResponse request
	DeleteApiV4ProjectsIdRepositoryTagsTagNameWithResponse(ctx context.Context, id string, tagName string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRepositoryTagsTagNameResponse, error)

	// GetApiV4ProjectsIdRepositoryTagsTagNameWithResponse request
	GetApiV4ProjectsIdRepositoryTagsTagNameWithResponse(ctx context.Context, id string, tagName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryTagsTagNameResponse, error)

	// GetApiV4ProjectsIdRepositoryTagsTagNameSignatureWithResponse request
	GetApiV4ProjectsIdRepositoryTagsTagNameSignatureWithResponse(ctx context.Context, id string, tagName string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryTagsTagNameSignatureResponse, error)

	// GetApiV4ProjectsIdRepositoryTreeWithResponse request
	GetApiV4ProjectsIdRepositoryTreeWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRepositoryTreeParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRepositoryTreeResponse, error)

	// PostApiV4ProjectsIdRepositorySizeWithResponse request
	PostApiV4ProjectsIdRepositorySizeWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRepositorySizeResponse, error)

	// GetApiV4ProjectsIdResourceGroupsWithResponse request
	GetApiV4ProjectsIdResourceGroupsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdResourceGroupsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdResourceGroupsResponse, error)

	// GetApiV4ProjectsIdResourceGroupsKeyWithResponse request
	GetApiV4ProjectsIdResourceGroupsKeyWithResponse(ctx context.Context, id string, key string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdResourceGroupsKeyResponse, error)

	// PutApiV4ProjectsIdResourceGroupsKeyWithBodyWithResponse request with any body
	PutApiV4ProjectsIdResourceGroupsKeyWithBodyWithResponse(ctx context.Context, id string, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdResourceGroupsKeyResponse, error)

	PutApiV4ProjectsIdResourceGroupsKeyWithResponse(ctx context.Context, id string, key string, body PutApiV4ProjectsIdResourceGroupsKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdResourceGroupsKeyResponse, error)

	// GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsWithResponse request
	GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsWithResponse(ctx context.Context, id string, key string, params *GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdResourceGroupsKeyUpcomingJobsResponse, error)

	// PostApiV4ProjectsIdRestoreWithResponse request
	PostApiV4ProjectsIdRestoreWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRestoreResponse, error)

	// GetApiV4ProjectsIdRunnersWithResponse request
	GetApiV4ProjectsIdRunnersWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdRunnersParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdRunnersResponse, error)

	// PostApiV4ProjectsIdRunnersWithBodyWithResponse request with any body
	PostApiV4ProjectsIdRunnersWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRunnersResponse, error)

	PostApiV4ProjectsIdRunnersWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdRunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRunnersResponse, error)

	// PostApiV4ProjectsIdRunnersResetRegistrationTokenWithResponse request
	PostApiV4ProjectsIdRunnersResetRegistrationTokenWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdRunnersResetRegistrationTokenResponse, error)

	// DeleteApiV4ProjectsIdRunnersRunnerIdWithResponse request
	DeleteApiV4ProjectsIdRunnersRunnerIdWithResponse(ctx context.Context, id string, runnerId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdRunnersRunnerIdResponse, error)

	// GetApiV4ProjectsIdSecureFilesWithResponse request
	GetApiV4ProjectsIdSecureFilesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdSecureFilesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSecureFilesResponse, error)

	// PostApiV4ProjectsIdSecureFilesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdSecureFilesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdSecureFilesResponse, error)

	PostApiV4ProjectsIdSecureFilesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdSecureFilesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdSecureFilesResponse, error)

	// DeleteApiV4ProjectsIdSecureFilesSecureFileIdWithResponse request
	DeleteApiV4ProjectsIdSecureFilesSecureFileIdWithResponse(ctx context.Context, id string, secureFileId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdSecureFilesSecureFileIdResponse, error)

	// GetApiV4ProjectsIdSecureFilesSecureFileIdWithResponse request
	GetApiV4ProjectsIdSecureFilesSecureFileIdWithResponse(ctx context.Context, id string, secureFileId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSecureFilesSecureFileIdResponse, error)

	// GetApiV4ProjectsIdSecureFilesSecureFileIdDownloadWithResponse request
	GetApiV4ProjectsIdSecureFilesSecureFileIdDownloadWithResponse(ctx context.Context, id string, secureFileId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSecureFilesSecureFileIdDownloadResponse, error)

	// GetApiV4ProjectsIdServicesWithResponse request
	GetApiV4ProjectsIdServicesWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdServicesResponse, error)

	// PutApiV4ProjectsIdServicesAppleAppStoreWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesAppleAppStoreWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesAppleAppStoreResponse, error)

	PutApiV4ProjectsIdServicesAppleAppStoreWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesAppleAppStoreJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesAppleAppStoreResponse, error)

	// PutApiV4ProjectsIdServicesAsanaWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesAsanaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesAsanaResponse, error)

	PutApiV4ProjectsIdServicesAsanaWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesAsanaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesAsanaResponse, error)

	// PutApiV4ProjectsIdServicesAssemblaWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesAssemblaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesAssemblaResponse, error)

	PutApiV4ProjectsIdServicesAssemblaWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesAssemblaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesAssemblaResponse, error)

	// PutApiV4ProjectsIdServicesBambooWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesBambooWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesBambooResponse, error)

	PutApiV4ProjectsIdServicesBambooWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesBambooJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesBambooResponse, error)

	// PutApiV4ProjectsIdServicesBugzillaWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesBugzillaWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesBugzillaResponse, error)

	PutApiV4ProjectsIdServicesBugzillaWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesBugzillaJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesBugzillaResponse, error)

	// PutApiV4ProjectsIdServicesBuildkiteWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesBuildkiteWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesBuildkiteResponse, error)

	PutApiV4ProjectsIdServicesBuildkiteWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesBuildkiteJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesBuildkiteResponse, error)

	// PutApiV4ProjectsIdServicesCampfireWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesCampfireWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesCampfireResponse, error)

	PutApiV4ProjectsIdServicesCampfireWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesCampfireJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesCampfireResponse, error)

	// PutApiV4ProjectsIdServicesClickupWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesClickupWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesClickupResponse, error)

	PutApiV4ProjectsIdServicesClickupWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesClickupJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesClickupResponse, error)

	// PutApiV4ProjectsIdServicesConfluenceWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesConfluenceWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesConfluenceResponse, error)

	PutApiV4ProjectsIdServicesConfluenceWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesConfluenceJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesConfluenceResponse, error)

	// PutApiV4ProjectsIdServicesCustomIssueTrackerWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesCustomIssueTrackerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesCustomIssueTrackerResponse, error)

	PutApiV4ProjectsIdServicesCustomIssueTrackerWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesCustomIssueTrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesCustomIssueTrackerResponse, error)

	// PutApiV4ProjectsIdServicesDatadogWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesDatadogWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesDatadogResponse, error)

	PutApiV4ProjectsIdServicesDatadogWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesDatadogJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesDatadogResponse, error)

	// PutApiV4ProjectsIdServicesDiffblueCoverWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesDiffblueCoverWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesDiffblueCoverResponse, error)

	PutApiV4ProjectsIdServicesDiffblueCoverWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesDiffblueCoverJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesDiffblueCoverResponse, error)

	// PutApiV4ProjectsIdServicesDiscordWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesDiscordWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesDiscordResponse, error)

	PutApiV4ProjectsIdServicesDiscordWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesDiscordJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesDiscordResponse, error)

	// PutApiV4ProjectsIdServicesDroneCiWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesDroneCiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesDroneCiResponse, error)

	PutApiV4ProjectsIdServicesDroneCiWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesDroneCiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesDroneCiResponse, error)

	// PutApiV4ProjectsIdServicesEmailsOnPushWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesEmailsOnPushWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesEmailsOnPushResponse, error)

	PutApiV4ProjectsIdServicesEmailsOnPushWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesEmailsOnPushJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesEmailsOnPushResponse, error)

	// PutApiV4ProjectsIdServicesEwmWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesEwmWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesEwmResponse, error)

	PutApiV4ProjectsIdServicesEwmWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesEwmJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesEwmResponse, error)

	// PutApiV4ProjectsIdServicesExternalWikiWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesExternalWikiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesExternalWikiResponse, error)

	PutApiV4ProjectsIdServicesExternalWikiWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesExternalWikiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesExternalWikiResponse, error)

	// PutApiV4ProjectsIdServicesGitGuardianWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesGitGuardianWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGitGuardianResponse, error)

	PutApiV4ProjectsIdServicesGitGuardianWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGitGuardianJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGitGuardianResponse, error)

	// PutApiV4ProjectsIdServicesGithubWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesGithubWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGithubResponse, error)

	PutApiV4ProjectsIdServicesGithubWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGithubJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGithubResponse, error)

	// PutApiV4ProjectsIdServicesGitlabSlackApplicationWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesGitlabSlackApplicationWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGitlabSlackApplicationResponse, error)

	PutApiV4ProjectsIdServicesGitlabSlackApplicationWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGitlabSlackApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGitlabSlackApplicationResponse, error)

	// PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryResponse, error)

	PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGoogleCloudPlatformArtifactRegistryResponse, error)

	// PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationResponse, error)

	PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGoogleCloudPlatformWorkloadIdentityFederationResponse, error)

	// PutApiV4ProjectsIdServicesGooglePlayWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesGooglePlayWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGooglePlayResponse, error)

	PutApiV4ProjectsIdServicesGooglePlayWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesGooglePlayJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesGooglePlayResponse, error)

	// PutApiV4ProjectsIdServicesHangoutsChatWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesHangoutsChatWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesHangoutsChatResponse, error)

	PutApiV4ProjectsIdServicesHangoutsChatWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesHangoutsChatJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesHangoutsChatResponse, error)

	// PutApiV4ProjectsIdServicesHarborWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesHarborWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesHarborResponse, error)

	PutApiV4ProjectsIdServicesHarborWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesHarborJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesHarborResponse, error)

	// PutApiV4ProjectsIdServicesIrkerWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesIrkerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesIrkerResponse, error)

	PutApiV4ProjectsIdServicesIrkerWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesIrkerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesIrkerResponse, error)

	// PutApiV4ProjectsIdServicesJenkinsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesJenkinsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesJenkinsResponse, error)

	PutApiV4ProjectsIdServicesJenkinsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesJenkinsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesJenkinsResponse, error)

	// PutApiV4ProjectsIdServicesJiraWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesJiraWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesJiraResponse, error)

	PutApiV4ProjectsIdServicesJiraWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesJiraJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesJiraResponse, error)

	// PutApiV4ProjectsIdServicesJiraCloudAppWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesJiraCloudAppWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesJiraCloudAppResponse, error)

	PutApiV4ProjectsIdServicesJiraCloudAppWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesJiraCloudAppJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesJiraCloudAppResponse, error)

	// PutApiV4ProjectsIdServicesMatrixWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesMatrixWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMatrixResponse, error)

	PutApiV4ProjectsIdServicesMatrixWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMatrixJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMatrixResponse, error)

	// PutApiV4ProjectsIdServicesMattermostWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesMattermostWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMattermostResponse, error)

	PutApiV4ProjectsIdServicesMattermostWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMattermostJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMattermostResponse, error)

	// PutApiV4ProjectsIdServicesMattermostSlashCommandsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesMattermostSlashCommandsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMattermostSlashCommandsResponse, error)

	PutApiV4ProjectsIdServicesMattermostSlashCommandsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMattermostSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMattermostSlashCommandsResponse, error)

	// PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerWithBodyWithResponse request with any body
	PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerResponse, error)

	PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdServicesMattermostSlashCommandsTriggerResponse, error)

	// PutApiV4ProjectsIdServicesMicrosoftTeamsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesMicrosoftTeamsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMicrosoftTeamsResponse, error)

	PutApiV4ProjectsIdServicesMicrosoftTeamsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMicrosoftTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMicrosoftTeamsResponse, error)

	// PutApiV4ProjectsIdServicesMockCiWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesMockCiWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMockCiResponse, error)

	PutApiV4ProjectsIdServicesMockCiWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMockCiJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMockCiResponse, error)

	// PutApiV4ProjectsIdServicesMockMonitoringWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesMockMonitoringWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMockMonitoringResponse, error)

	PutApiV4ProjectsIdServicesMockMonitoringWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesMockMonitoringJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesMockMonitoringResponse, error)

	// PutApiV4ProjectsIdServicesPackagistWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesPackagistWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPackagistResponse, error)

	PutApiV4ProjectsIdServicesPackagistWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPackagistJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPackagistResponse, error)

	// PutApiV4ProjectsIdServicesPhorgeWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesPhorgeWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPhorgeResponse, error)

	PutApiV4ProjectsIdServicesPhorgeWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPhorgeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPhorgeResponse, error)

	// PutApiV4ProjectsIdServicesPipelinesEmailWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesPipelinesEmailWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPipelinesEmailResponse, error)

	PutApiV4ProjectsIdServicesPipelinesEmailWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPipelinesEmailJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPipelinesEmailResponse, error)

	// PutApiV4ProjectsIdServicesPivotaltrackerWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesPivotaltrackerWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPivotaltrackerResponse, error)

	PutApiV4ProjectsIdServicesPivotaltrackerWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPivotaltrackerJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPivotaltrackerResponse, error)

	// PutApiV4ProjectsIdServicesPumbleWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesPumbleWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPumbleResponse, error)

	PutApiV4ProjectsIdServicesPumbleWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPumbleJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPumbleResponse, error)

	// PutApiV4ProjectsIdServicesPushoverWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesPushoverWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPushoverResponse, error)

	PutApiV4ProjectsIdServicesPushoverWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesPushoverJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesPushoverResponse, error)

	// PutApiV4ProjectsIdServicesRedmineWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesRedmineWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesRedmineResponse, error)

	PutApiV4ProjectsIdServicesRedmineWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesRedmineJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesRedmineResponse, error)

	// PutApiV4ProjectsIdServicesSlackWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesSlackWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesSlackResponse, error)

	PutApiV4ProjectsIdServicesSlackWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesSlackJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesSlackResponse, error)

	// PutApiV4ProjectsIdServicesSlackSlashCommandsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesSlackSlashCommandsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesSlackSlashCommandsResponse, error)

	PutApiV4ProjectsIdServicesSlackSlashCommandsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesSlackSlashCommandsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesSlackSlashCommandsResponse, error)

	// PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerWithBodyWithResponse request with any body
	PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerResponse, error)

	PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdServicesSlackSlashCommandsTriggerResponse, error)

	// PutApiV4ProjectsIdServicesSquashTmWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesSquashTmWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesSquashTmResponse, error)

	PutApiV4ProjectsIdServicesSquashTmWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesSquashTmJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesSquashTmResponse, error)

	// PutApiV4ProjectsIdServicesTeamcityWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesTeamcityWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesTeamcityResponse, error)

	PutApiV4ProjectsIdServicesTeamcityWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesTeamcityJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesTeamcityResponse, error)

	// PutApiV4ProjectsIdServicesTelegramWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesTelegramWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesTelegramResponse, error)

	PutApiV4ProjectsIdServicesTelegramWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesTelegramJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesTelegramResponse, error)

	// PutApiV4ProjectsIdServicesUnifyCircuitWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesUnifyCircuitWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesUnifyCircuitResponse, error)

	PutApiV4ProjectsIdServicesUnifyCircuitWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesUnifyCircuitJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesUnifyCircuitResponse, error)

	// PutApiV4ProjectsIdServicesWebexTeamsWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesWebexTeamsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesWebexTeamsResponse, error)

	PutApiV4ProjectsIdServicesWebexTeamsWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesWebexTeamsJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesWebexTeamsResponse, error)

	// PutApiV4ProjectsIdServicesYoutrackWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesYoutrackWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesYoutrackResponse, error)

	PutApiV4ProjectsIdServicesYoutrackWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesYoutrackJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesYoutrackResponse, error)

	// PutApiV4ProjectsIdServicesZentaoWithBodyWithResponse request with any body
	PutApiV4ProjectsIdServicesZentaoWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesZentaoResponse, error)

	PutApiV4ProjectsIdServicesZentaoWithResponse(ctx context.Context, id int32, body PutApiV4ProjectsIdServicesZentaoJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdServicesZentaoResponse, error)

	// DeleteApiV4ProjectsIdServicesSlugWithResponse request
	DeleteApiV4ProjectsIdServicesSlugWithResponse(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdServicesSlugResponse, error)

	// GetApiV4ProjectsIdServicesSlugWithResponse request
	GetApiV4ProjectsIdServicesSlugWithResponse(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdServicesSlugResponse, error)

	// PostApiV4ProjectsIdShareWithBodyWithResponse request with any body
	PostApiV4ProjectsIdShareWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdShareResponse, error)

	PostApiV4ProjectsIdShareWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdShareJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdShareResponse, error)

	// DeleteApiV4ProjectsIdShareGroupIdWithResponse request
	DeleteApiV4ProjectsIdShareGroupIdWithResponse(ctx context.Context, id string, groupId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdShareGroupIdResponse, error)

	// GetApiV4ProjectsIdShareLocationsWithResponse request
	GetApiV4ProjectsIdShareLocationsWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdShareLocationsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdShareLocationsResponse, error)

	// GetApiV4ProjectsIdSnapshotWithResponse request
	GetApiV4ProjectsIdSnapshotWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdSnapshotParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnapshotResponse, error)

	// GetApiV4ProjectsIdSnippetsWithResponse request
	GetApiV4ProjectsIdSnippetsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdSnippetsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsResponse, error)

	// PostApiV4ProjectsIdSnippetsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdSnippetsWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdSnippetsResponse, error)

	PostApiV4ProjectsIdSnippetsWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdSnippetsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdSnippetsResponse, error)

	// DeleteApiV4ProjectsIdSnippetsSnippetIdWithResponse request
	DeleteApiV4ProjectsIdSnippetsSnippetIdWithResponse(ctx context.Context, id string, snippetId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdSnippetsSnippetIdResponse, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdWithResponse request
	GetApiV4ProjectsIdSnippetsSnippetIdWithResponse(ctx context.Context, id string, snippetId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsSnippetIdResponse, error)

	// PutApiV4ProjectsIdSnippetsSnippetIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdSnippetsSnippetIdWithBodyWithResponse(ctx context.Context, id string, snippetId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdSnippetsSnippetIdResponse, error)

	PutApiV4ProjectsIdSnippetsSnippetIdWithResponse(ctx context.Context, id string, snippetId int32, body PutApiV4ProjectsIdSnippetsSnippetIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdSnippetsSnippetIdResponse, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiWithResponse request
	GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiWithResponse(ctx context.Context, id string, snippetId int32, params *GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiResponse, error)

	// PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiWithBodyWithResponse request with any body
	PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiWithBodyWithResponse(ctx context.Context, id int32, snippetId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiResponse, error)

	PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiWithResponse(ctx context.Context, id int32, snippetId int32, body PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdSnippetsSnippetIdAwardEmojiResponse, error)

	// DeleteApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardIdWithResponse request
	DeleteApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, snippetId int32, awardId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardIdWithResponse request
	GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, snippetId int32, awardId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsSnippetIdAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRawWithResponse request
	GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRawWithResponse(ctx context.Context, id string, snippetId int32, ref string, filePath string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsSnippetIdFilesRefFilePathRawResponse, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiWithResponse request
	GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiWithResponse(ctx context.Context, id int32, snippetId int32, noteId int32, params *GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiResponse, error)

	// PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiWithBodyWithResponse request with any body
	PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiWithBodyWithResponse(ctx context.Context, id int32, snippetId int32, noteId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiResponse, error)

	PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiWithResponse(ctx context.Context, id int32, snippetId int32, noteId int32, body PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiResponse, error)

	// DeleteApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdWithResponse request
	DeleteApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, snippetId int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdWithResponse request
	GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdWithResponse(ctx context.Context, id int32, snippetId int32, noteId int32, awardId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsSnippetIdNotesNoteIdAwardEmojiAwardIdResponse, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdRawWithResponse request
	GetApiV4ProjectsIdSnippetsSnippetIdRawWithResponse(ctx context.Context, id string, snippetId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsSnippetIdRawResponse, error)

	// GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetailWithResponse request
	GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetailWithResponse(ctx context.Context, id string, snippetId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdSnippetsSnippetIdUserAgentDetailResponse, error)

	// PostApiV4ProjectsIdStarWithResponse request
	PostApiV4ProjectsIdStarWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdStarResponse, error)

	// GetApiV4ProjectsIdStarrersWithResponse request
	GetApiV4ProjectsIdStarrersWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdStarrersParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdStarrersResponse, error)

	// GetApiV4ProjectsIdStatisticsWithResponse request
	GetApiV4ProjectsIdStatisticsWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdStatisticsResponse, error)

	// PostApiV4ProjectsIdStatusesShaWithBodyWithResponse request with any body
	PostApiV4ProjectsIdStatusesShaWithBodyWithResponse(ctx context.Context, id string, sha string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdStatusesShaResponse, error)

	PostApiV4ProjectsIdStatusesShaWithResponse(ctx context.Context, id string, sha string, body PostApiV4ProjectsIdStatusesShaJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdStatusesShaResponse, error)

	// GetApiV4ProjectsIdStorageWithResponse request
	GetApiV4ProjectsIdStorageWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdStorageResponse, error)

	// GetApiV4ProjectsIdTemplatesTypeWithResponse request
	GetApiV4ProjectsIdTemplatesTypeWithResponse(ctx context.Context, id string, pType string, params *GetApiV4ProjectsIdTemplatesTypeParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdTemplatesTypeResponse, error)

	// GetApiV4ProjectsIdTemplatesTypeNameWithResponse request
	GetApiV4ProjectsIdTemplatesTypeNameWithResponse(ctx context.Context, id string, pType string, name string, params *GetApiV4ProjectsIdTemplatesTypeNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdTemplatesTypeNameResponse, error)

	// DeleteApiV4ProjectsIdTerraformStateNameWithResponse request
	DeleteApiV4ProjectsIdTerraformStateNameWithResponse(ctx context.Context, id string, name int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdTerraformStateNameResponse, error)

	// GetApiV4ProjectsIdTerraformStateNameWithResponse request
	GetApiV4ProjectsIdTerraformStateNameWithResponse(ctx context.Context, id string, name string, params *GetApiV4ProjectsIdTerraformStateNameParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdTerraformStateNameResponse, error)

	// PostApiV4ProjectsIdTerraformStateNameWithResponse request
	PostApiV4ProjectsIdTerraformStateNameWithResponse(ctx context.Context, id string, name int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdTerraformStateNameResponse, error)

	// DeleteApiV4ProjectsIdTerraformStateNameLockWithResponse request
	DeleteApiV4ProjectsIdTerraformStateNameLockWithResponse(ctx context.Context, id string, name int32, params *DeleteApiV4ProjectsIdTerraformStateNameLockParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdTerraformStateNameLockResponse, error)

	// PostApiV4ProjectsIdTerraformStateNameLockWithBodyWithResponse request with any body
	PostApiV4ProjectsIdTerraformStateNameLockWithBodyWithResponse(ctx context.Context, id string, name int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdTerraformStateNameLockResponse, error)

	PostApiV4ProjectsIdTerraformStateNameLockWithResponse(ctx context.Context, id string, name int32, body PostApiV4ProjectsIdTerraformStateNameLockJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdTerraformStateNameLockResponse, error)

	// DeleteApiV4ProjectsIdTerraformStateNameVersionsSerialWithResponse request
	DeleteApiV4ProjectsIdTerraformStateNameVersionsSerialWithResponse(ctx context.Context, id string, name int32, serial int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdTerraformStateNameVersionsSerialResponse, error)

	// GetApiV4ProjectsIdTerraformStateNameVersionsSerialWithResponse request
	GetApiV4ProjectsIdTerraformStateNameVersionsSerialWithResponse(ctx context.Context, id string, name string, serial int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdTerraformStateNameVersionsSerialResponse, error)

	// PutApiV4ProjectsIdTransferWithBodyWithResponse request with any body
	PutApiV4ProjectsIdTransferWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdTransferResponse, error)

	PutApiV4ProjectsIdTransferWithResponse(ctx context.Context, id string, body PutApiV4ProjectsIdTransferJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdTransferResponse, error)

	// GetApiV4ProjectsIdTransferLocationsWithResponse request
	GetApiV4ProjectsIdTransferLocationsWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdTransferLocationsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdTransferLocationsResponse, error)

	// GetApiV4ProjectsIdTriggersWithResponse request
	GetApiV4ProjectsIdTriggersWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdTriggersParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdTriggersResponse, error)

	// PostApiV4ProjectsIdTriggersWithBodyWithResponse request with any body
	PostApiV4ProjectsIdTriggersWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdTriggersResponse, error)

	PostApiV4ProjectsIdTriggersWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdTriggersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdTriggersResponse, error)

	// DeleteApiV4ProjectsIdTriggersTriggerIdWithResponse request
	DeleteApiV4ProjectsIdTriggersTriggerIdWithResponse(ctx context.Context, id string, triggerId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdTriggersTriggerIdResponse, error)

	// GetApiV4ProjectsIdTriggersTriggerIdWithResponse request
	GetApiV4ProjectsIdTriggersTriggerIdWithResponse(ctx context.Context, id string, triggerId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdTriggersTriggerIdResponse, error)

	// PutApiV4ProjectsIdTriggersTriggerIdWithBodyWithResponse request with any body
	PutApiV4ProjectsIdTriggersTriggerIdWithBodyWithResponse(ctx context.Context, id string, triggerId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdTriggersTriggerIdResponse, error)

	PutApiV4ProjectsIdTriggersTriggerIdWithResponse(ctx context.Context, id string, triggerId int32, body PutApiV4ProjectsIdTriggersTriggerIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdTriggersTriggerIdResponse, error)

	// PostApiV4ProjectsIdUnarchiveWithResponse request
	PostApiV4ProjectsIdUnarchiveWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdUnarchiveResponse, error)

	// PostApiV4ProjectsIdUnstarWithResponse request
	PostApiV4ProjectsIdUnstarWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdUnstarResponse, error)

	// GetApiV4ProjectsIdUploadsWithResponse request
	GetApiV4ProjectsIdUploadsWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdUploadsParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdUploadsResponse, error)

	// PostApiV4ProjectsIdUploadsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdUploadsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdUploadsResponse, error)

	PostApiV4ProjectsIdUploadsWithResponse(ctx context.Context, id int32, body PostApiV4ProjectsIdUploadsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdUploadsResponse, error)

	// PostApiV4ProjectsIdUploadsAuthorizeWithResponse request
	PostApiV4ProjectsIdUploadsAuthorizeWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdUploadsAuthorizeResponse, error)

	// DeleteApiV4ProjectsIdUploadsSecretFilenameWithResponse request
	DeleteApiV4ProjectsIdUploadsSecretFilenameWithResponse(ctx context.Context, id int32, secret string, filename string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdUploadsSecretFilenameResponse, error)

	// GetApiV4ProjectsIdUploadsSecretFilenameWithResponse request
	GetApiV4ProjectsIdUploadsSecretFilenameWithResponse(ctx context.Context, id int32, secret string, filename string, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdUploadsSecretFilenameResponse, error)

	// DeleteApiV4ProjectsIdUploadsUploadIdWithResponse request
	DeleteApiV4ProjectsIdUploadsUploadIdWithResponse(ctx context.Context, id int32, uploadId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdUploadsUploadIdResponse, error)

	// GetApiV4ProjectsIdUploadsUploadIdWithResponse request
	GetApiV4ProjectsIdUploadsUploadIdWithResponse(ctx context.Context, id int32, uploadId int32, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdUploadsUploadIdResponse, error)

	// GetApiV4ProjectsIdUsersWithResponse request
	GetApiV4ProjectsIdUsersWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdUsersParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdUsersResponse, error)

	// GetApiV4ProjectsIdVariablesWithResponse request
	GetApiV4ProjectsIdVariablesWithResponse(ctx context.Context, id string, params *GetApiV4ProjectsIdVariablesParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdVariablesResponse, error)

	// PostApiV4ProjectsIdVariablesWithBodyWithResponse request with any body
	PostApiV4ProjectsIdVariablesWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdVariablesResponse, error)

	PostApiV4ProjectsIdVariablesWithResponse(ctx context.Context, id string, body PostApiV4ProjectsIdVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdVariablesResponse, error)

	// DeleteApiV4ProjectsIdVariablesKeyWithResponse request
	DeleteApiV4ProjectsIdVariablesKeyWithResponse(ctx context.Context, id string, key string, params *DeleteApiV4ProjectsIdVariablesKeyParams, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdVariablesKeyResponse, error)

	// GetApiV4ProjectsIdVariablesKeyWithResponse request
	GetApiV4ProjectsIdVariablesKeyWithResponse(ctx context.Context, id string, key string, params *GetApiV4ProjectsIdVariablesKeyParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdVariablesKeyResponse, error)

	// PutApiV4ProjectsIdVariablesKeyWithBodyWithResponse request with any body
	PutApiV4ProjectsIdVariablesKeyWithBodyWithResponse(ctx context.Context, id string, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdVariablesKeyResponse, error)

	PutApiV4ProjectsIdVariablesKeyWithResponse(ctx context.Context, id string, key string, body PutApiV4ProjectsIdVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdVariablesKeyResponse, error)

	// GetApiV4ProjectsIdWikisWithResponse request
	GetApiV4ProjectsIdWikisWithResponse(ctx context.Context, id int32, params *GetApiV4ProjectsIdWikisParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdWikisResponse, error)

	// PostApiV4ProjectsIdWikisWithBodyWithResponse request with any body
	PostApiV4ProjectsIdWikisWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdWikisResponse, error)

	PostApiV4ProjectsIdWikisWithResponse(ctx context.Context, id int32, body PostApiV4ProjectsIdWikisJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdWikisResponse, error)

	// PostApiV4ProjectsIdWikisAttachmentsWithBodyWithResponse request with any body
	PostApiV4ProjectsIdWikisAttachmentsWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdWikisAttachmentsResponse, error)

	PostApiV4ProjectsIdWikisAttachmentsWithResponse(ctx context.Context, id int32, body PostApiV4ProjectsIdWikisAttachmentsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4ProjectsIdWikisAttachmentsResponse, error)

	// DeleteApiV4ProjectsIdWikisSlugWithResponse request
	DeleteApiV4ProjectsIdWikisSlugWithResponse(ctx context.Context, id int32, slug string, reqEditors ...RequestEditorFn) (*DeleteApiV4ProjectsIdWikisSlugResponse, error)

	// GetApiV4ProjectsIdWikisSlugWithResponse request
	GetApiV4ProjectsIdWikisSlugWithResponse(ctx context.Context, id int32, slug string, params *GetApiV4ProjectsIdWikisSlugParams, reqEditors ...RequestEditorFn) (*GetApiV4ProjectsIdWikisSlugResponse, error)

	// PutApiV4ProjectsIdWikisSlugWithBodyWithResponse request with any body
	PutApiV4ProjectsIdWikisSlugWithBodyWithResponse(ctx context.Context, id int32, slug int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdWikisSlugResponse, error)

	PutApiV4ProjectsIdWikisSlugWithResponse(ctx context.Context, id int32, slug int32, body PutApiV4ProjectsIdWikisSlugJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4ProjectsIdWikisSlugResponse, error)

	// GetApiV4RegistryRepositoriesIdWithResponse request
	GetApiV4RegistryRepositoriesIdWithResponse(ctx context.Context, id string, params *GetApiV4RegistryRepositoriesIdParams, reqEditors ...RequestEditorFn) (*GetApiV4RegistryRepositoriesIdResponse, error)

	// DeleteApiV4RunnersWithResponse request
	DeleteApiV4RunnersWithResponse(ctx context.Context, params *DeleteApiV4RunnersParams, reqEditors ...RequestEditorFn) (*DeleteApiV4RunnersResponse, error)

	// GetApiV4RunnersWithResponse request
	GetApiV4RunnersWithResponse(ctx context.Context, params *GetApiV4RunnersParams, reqEditors ...RequestEditorFn) (*GetApiV4RunnersResponse, error)

	// PostApiV4RunnersWithBodyWithResponse request with any body
	PostApiV4RunnersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResponse, error)

	PostApiV4RunnersWithResponse(ctx context.Context, body PostApiV4RunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResponse, error)

	// GetApiV4RunnersAllWithResponse request
	GetApiV4RunnersAllWithResponse(ctx context.Context, params *GetApiV4RunnersAllParams, reqEditors ...RequestEditorFn) (*GetApiV4RunnersAllResponse, error)

	// DeleteApiV4RunnersManagersWithResponse request
	DeleteApiV4RunnersManagersWithResponse(ctx context.Context, params *DeleteApiV4RunnersManagersParams, reqEditors ...RequestEditorFn) (*DeleteApiV4RunnersManagersResponse, error)

	// PostApiV4RunnersResetAuthenticationTokenWithBodyWithResponse request with any body
	PostApiV4RunnersResetAuthenticationTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResetAuthenticationTokenResponse, error)

	PostApiV4RunnersResetAuthenticationTokenWithResponse(ctx context.Context, body PostApiV4RunnersResetAuthenticationTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResetAuthenticationTokenResponse, error)

	// PostApiV4RunnersResetRegistrationTokenWithResponse request
	PostApiV4RunnersResetRegistrationTokenWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV4RunnersResetRegistrationTokenResponse, error)

	// PostApiV4RunnersVerifyWithBodyWithResponse request with any body
	PostApiV4RunnersVerifyWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4RunnersVerifyResponse, error)

	PostApiV4RunnersVerifyWithResponse(ctx context.Context, body PostApiV4RunnersVerifyJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4RunnersVerifyResponse, error)

	// DeleteApiV4RunnersIdWithResponse request
	DeleteApiV4RunnersIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4RunnersIdResponse, error)

	// GetApiV4RunnersIdWithResponse request
	GetApiV4RunnersIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4RunnersIdResponse, error)

	// PutApiV4RunnersIdWithBodyWithResponse request with any body
	PutApiV4RunnersIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4RunnersIdResponse, error)

	PutApiV4RunnersIdWithResponse(ctx context.Context, id int32, body PutApiV4RunnersIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4RunnersIdResponse, error)

	// GetApiV4RunnersIdJobsWithResponse request
	GetApiV4RunnersIdJobsWithResponse(ctx context.Context, id int32, params *GetApiV4RunnersIdJobsParams, reqEditors ...RequestEditorFn) (*GetApiV4RunnersIdJobsResponse, error)

	// GetApiV4RunnersIdManagersWithResponse request
	GetApiV4RunnersIdManagersWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4RunnersIdManagersResponse, error)

	// PostApiV4RunnersIdResetAuthenticationTokenWithResponse request
	PostApiV4RunnersIdResetAuthenticationTokenWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4RunnersIdResetAuthenticationTokenResponse, error)

	// PostApiV4SlackTriggerWithBodyWithResponse request with any body
	PostApiV4SlackTriggerWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4SlackTriggerResponse, error)

	PostApiV4SlackTriggerWithResponse(ctx context.Context, body PostApiV4SlackTriggerJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4SlackTriggerResponse, error)

	// GetApiV4SnippetsWithResponse request
	GetApiV4SnippetsWithResponse(ctx context.Context, params *GetApiV4SnippetsParams, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsResponse, error)

	// PostApiV4SnippetsWithBodyWithResponse request with any body
	PostApiV4SnippetsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4SnippetsResponse, error)

	PostApiV4SnippetsWithResponse(ctx context.Context, body PostApiV4SnippetsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4SnippetsResponse, error)

	// GetApiV4SnippetsAllWithResponse request
	GetApiV4SnippetsAllWithResponse(ctx context.Context, params *GetApiV4SnippetsAllParams, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsAllResponse, error)

	// GetApiV4SnippetsPublicWithResponse request
	GetApiV4SnippetsPublicWithResponse(ctx context.Context, params *GetApiV4SnippetsPublicParams, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsPublicResponse, error)

	// DeleteApiV4SnippetsIdWithResponse request
	DeleteApiV4SnippetsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4SnippetsIdResponse, error)

	// GetApiV4SnippetsIdWithResponse request
	GetApiV4SnippetsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsIdResponse, error)

	// PutApiV4SnippetsIdWithBodyWithResponse request with any body
	PutApiV4SnippetsIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4SnippetsIdResponse, error)

	PutApiV4SnippetsIdWithResponse(ctx context.Context, id int32, body PutApiV4SnippetsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4SnippetsIdResponse, error)

	// GetApiV4SnippetsIdFilesRefFilePathRawWithResponse request
	GetApiV4SnippetsIdFilesRefFilePathRawWithResponse(ctx context.Context, id int32, ref string, filePath string, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsIdFilesRefFilePathRawResponse, error)

	// GetApiV4SnippetsIdRawWithResponse request
	GetApiV4SnippetsIdRawWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsIdRawResponse, error)

	// GetApiV4SnippetsIdUserAgentDetailWithResponse request
	GetApiV4SnippetsIdUserAgentDetailWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4SnippetsIdUserAgentDetailResponse, error)

	// PutApiV4SuggestionsBatchApplyWithBodyWithResponse request with any body
	PutApiV4SuggestionsBatchApplyWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4SuggestionsBatchApplyResponse, error)

	PutApiV4SuggestionsBatchApplyWithResponse(ctx context.Context, body PutApiV4SuggestionsBatchApplyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4SuggestionsBatchApplyResponse, error)

	// PutApiV4SuggestionsIdApplyWithBodyWithResponse request with any body
	PutApiV4SuggestionsIdApplyWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4SuggestionsIdApplyResponse, error)

	PutApiV4SuggestionsIdApplyWithResponse(ctx context.Context, id int32, body PutApiV4SuggestionsIdApplyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4SuggestionsIdApplyResponse, error)

	// GetApiV4TopicsWithResponse request
	GetApiV4TopicsWithResponse(ctx context.Context, params *GetApiV4TopicsParams, reqEditors ...RequestEditorFn) (*GetApiV4TopicsResponse, error)

	// PostApiV4TopicsWithBodyWithResponse request with any body
	PostApiV4TopicsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4TopicsResponse, error)

	PostApiV4TopicsWithResponse(ctx context.Context, body PostApiV4TopicsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4TopicsResponse, error)

	// PostApiV4TopicsMergeWithBodyWithResponse request with any body
	PostApiV4TopicsMergeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4TopicsMergeResponse, error)

	PostApiV4TopicsMergeWithResponse(ctx context.Context, body PostApiV4TopicsMergeJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4TopicsMergeResponse, error)

	// DeleteApiV4TopicsIdWithResponse request
	DeleteApiV4TopicsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*DeleteApiV4TopicsIdResponse, error)

	// GetApiV4TopicsIdWithResponse request
	GetApiV4TopicsIdWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*GetApiV4TopicsIdResponse, error)

	// PutApiV4TopicsIdWithBodyWithResponse request with any body
	PutApiV4TopicsIdWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4TopicsIdResponse, error)

	PutApiV4TopicsIdWithResponse(ctx context.Context, id int32, body PutApiV4TopicsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4TopicsIdResponse, error)

	// PostApiV4UsageDataIncrementCounterWithBodyWithResponse request with any body
	PostApiV4UsageDataIncrementCounterWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataIncrementCounterResponse, error)

	PostApiV4UsageDataIncrementCounterWithResponse(ctx context.Context, body PostApiV4UsageDataIncrementCounterJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataIncrementCounterResponse, error)

	// PostApiV4UsageDataIncrementUniqueUsersWithBodyWithResponse request with any body
	PostApiV4UsageDataIncrementUniqueUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataIncrementUniqueUsersResponse, error)

	PostApiV4UsageDataIncrementUniqueUsersWithResponse(ctx context.Context, body PostApiV4UsageDataIncrementUniqueUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataIncrementUniqueUsersResponse, error)

	// GetApiV4UsageDataMetricDefinitionsWithResponse request
	GetApiV4UsageDataMetricDefinitionsWithResponse(ctx context.Context, params *GetApiV4UsageDataMetricDefinitionsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsageDataMetricDefinitionsResponse, error)

	// GetApiV4UsageDataNonSqlMetricsWithResponse request
	GetApiV4UsageDataNonSqlMetricsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4UsageDataNonSqlMetricsResponse, error)

	// GetApiV4UsageDataQueriesWithResponse request
	GetApiV4UsageDataQueriesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4UsageDataQueriesResponse, error)

	// GetApiV4UsageDataServicePingWithResponse request
	GetApiV4UsageDataServicePingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4UsageDataServicePingResponse, error)

	// PostApiV4UsageDataTrackEventWithBodyWithResponse request with any body
	PostApiV4UsageDataTrackEventWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataTrackEventResponse, error)

	PostApiV4UsageDataTrackEventWithResponse(ctx context.Context, body PostApiV4UsageDataTrackEventJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataTrackEventResponse, error)

	// PostApiV4UsageDataTrackEventsWithBodyWithResponse request with any body
	PostApiV4UsageDataTrackEventsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataTrackEventsResponse, error)

	PostApiV4UsageDataTrackEventsWithResponse(ctx context.Context, body PostApiV4UsageDataTrackEventsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UsageDataTrackEventsResponse, error)

	// PostApiV4UserRunnersWithBodyWithResponse request with any body
	PostApiV4UserRunnersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4UserRunnersResponse, error)

	PostApiV4UserRunnersWithResponse(ctx context.Context, body PostApiV4UserRunnersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4UserRunnersResponse, error)

	// GetApiV4UserCountsWithResponse request
	GetApiV4UserCountsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4UserCountsResponse, error)

	// GetApiV4UsersIdEventsWithResponse request
	GetApiV4UsersIdEventsWithResponse(ctx context.Context, id string, params *GetApiV4UsersIdEventsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsersIdEventsResponse, error)

	// GetApiV4UsersUserIdContributedProjectsWithResponse request
	GetApiV4UsersUserIdContributedProjectsWithResponse(ctx context.Context, userId string, params *GetApiV4UsersUserIdContributedProjectsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsersUserIdContributedProjectsResponse, error)

	// GetApiV4UsersUserIdProjectsWithResponse request
	GetApiV4UsersUserIdProjectsWithResponse(ctx context.Context, userId string, params *GetApiV4UsersUserIdProjectsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsersUserIdProjectsResponse, error)

	// GetApiV4UsersUserIdStarredProjectsWithResponse request
	GetApiV4UsersUserIdStarredProjectsWithResponse(ctx context.Context, userId string, params *GetApiV4UsersUserIdStarredProjectsParams, reqEditors ...RequestEditorFn) (*GetApiV4UsersUserIdStarredProjectsResponse, error)

	// GetApiV4VersionWithResponse request
	GetApiV4VersionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4VersionResponse, error)

	// GetApiV4WebCommitsPublicKeyWithResponse request
	GetApiV4WebCommitsPublicKeyWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4WebCommitsPublicKeyResponse, error)
}

func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}
func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}
