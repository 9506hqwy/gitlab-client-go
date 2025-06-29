package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/oapi-codegen/runtime"
)

type PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONBody struct {
	Data struct {
		GlId        string `json:"gl_id"`
		PrimaryRepo string `json:"primary_repo"`
	} `json:"data"`
	SecretToken string `json:"secret_token"`
}
type PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONBody struct {
	Data struct {
		// GlId GitLab identifier of user that initiated the clone/pull
		GlId string `json:"gl_id"`

		// PrimaryRepo Primary repository to clone/pull
		PrimaryRepo string `json:"primary_repo"`
	} `json:"data"`

	// SecretToken Secret token to authenticate by gitlab shell
	SecretToken string `json:"secret_token"`
}
type PostApiV4GeoProxyGitSshReceivePackJSONBody struct {
	Data struct {
		GlId        string `json:"gl_id"`
		PrimaryRepo string `json:"primary_repo"`
	} `json:"data"`

	// Output Output from git-receive-pack
	Output      string `json:"output"`
	SecretToken string `json:"secret_token"`
}
type PostApiV4GeoProxyGitSshUploadPackJSONBody struct {
	Data struct {
		GlId        string `json:"gl_id"`
		PrimaryRepo string `json:"primary_repo"`
	} `json:"data"`

	// Output Output from git-upload-pack
	Output      string `json:"output"`
	SecretToken string `json:"secret_token"`
}
type PostApiV4GeoStatusJSONBody struct {
	Data *struct {
		// CursorLastEventDate Cursor last event date
		CursorLastEventDate *time.Time `json:"cursor_last_event_date,omitempty"`

		// CursorLastEventId Cursor last event ID
		CursorLastEventId *int32 `json:"cursor_last_event_id,omitempty"`

		// DbReplicationLagSeconds DB replication lag in seconds
		DbReplicationLagSeconds *int32 `json:"db_replication_lag_seconds,omitempty"`

		// GeoNodeId Geo Node ID to look up its status
		GeoNodeId int32 `json:"geo_node_id"`

		// LastEventDate Last event date
		LastEventDate *time.Time `json:"last_event_date,omitempty"`

		// LastEventId Last event ID
		LastEventId *int32 `json:"last_event_id,omitempty"`

		// LastSuccessfulStatusCheckAt Last successful status check date
		LastSuccessfulStatusCheckAt *time.Time `json:"last_successful_status_check_at,omitempty"`

		// ReplicationSlotsCount Replication slots count
		ReplicationSlotsCount *int32 `json:"replication_slots_count,omitempty"`

		// ReplicationSlotsMaxRetainedWalBytes Maximum number of bytes retained in the WAL on the primary
		ReplicationSlotsMaxRetainedWalBytes *int32 `json:"replication_slots_max_retained_wal_bytes,omitempty"`

		// ReplicationSlotsUsedCount Replication slots used count
		ReplicationSlotsUsedCount *int32 `json:"replication_slots_used_count,omitempty"`

		// Revision Gitlab revision
		Revision *string `json:"revision,omitempty"`
		Status   *struct {
			// CiSecureFilesChecksumFailedCount CI secure files checksum failed count
			CiSecureFilesChecksumFailedCount *int32 `json:"ci_secure_files_checksum_failed_count,omitempty"`

			// CiSecureFilesChecksumTotalCount CI secure files checksum total count
			CiSecureFilesChecksumTotalCount *int32 `json:"ci_secure_files_checksum_total_count,omitempty"`

			// CiSecureFilesChecksummedCount CI secure files checksummed count
			CiSecureFilesChecksummedCount *int32 `json:"ci_secure_files_checksummed_count,omitempty"`

			// CiSecureFilesCount CI secure files count
			CiSecureFilesCount *int32 `json:"ci_secure_files_count,omitempty"`

			// CiSecureFilesFailedCount CI secure files failed count
			CiSecureFilesFailedCount *int32 `json:"ci_secure_files_failed_count,omitempty"`

			// CiSecureFilesRegistryCount CI secure files registry count
			CiSecureFilesRegistryCount *int32 `json:"ci_secure_files_registry_count,omitempty"`

			// CiSecureFilesSyncedCount CI secure files synced count
			CiSecureFilesSyncedCount *int32 `json:"ci_secure_files_synced_count,omitempty"`

			// CiSecureFilesVerificationFailedCount CI secure files verification failed count
			CiSecureFilesVerificationFailedCount *int32 `json:"ci_secure_files_verification_failed_count,omitempty"`

			// CiSecureFilesVerificationTotalCount CI secure files verification total count
			CiSecureFilesVerificationTotalCount *int32 `json:"ci_secure_files_verification_total_count,omitempty"`

			// CiSecureFilesVerifiedCount CI secure files verified count
			CiSecureFilesVerifiedCount *int32 `json:"ci_secure_files_verified_count,omitempty"`

			// ContainerRepositoriesChecksumFailedCount Container repositories checksum failed count
			ContainerRepositoriesChecksumFailedCount *int32 `json:"container_repositories_checksum_failed_count,omitempty"`

			// ContainerRepositoriesChecksumTotalCount Container repositories checksum total count
			ContainerRepositoriesChecksumTotalCount *int32 `json:"container_repositories_checksum_total_count,omitempty"`

			// ContainerRepositoriesChecksummedCount Container repositories checksummed count
			ContainerRepositoriesChecksummedCount *int32 `json:"container_repositories_checksummed_count,omitempty"`

			// ContainerRepositoriesCount Container repositories count
			ContainerRepositoriesCount *int32 `json:"container_repositories_count,omitempty"`

			// ContainerRepositoriesFailedCount Container repositories failed count
			ContainerRepositoriesFailedCount *int32 `json:"container_repositories_failed_count,omitempty"`

			// ContainerRepositoriesRegistryCount Container repositories registry count
			ContainerRepositoriesRegistryCount *int32 `json:"container_repositories_registry_count,omitempty"`

			// ContainerRepositoriesReplicationEnabled Container repositories replication enabled
			ContainerRepositoriesReplicationEnabled *bool `json:"container_repositories_replication_enabled,omitempty"`

			// ContainerRepositoriesSyncedCount Container repositories synced count
			ContainerRepositoriesSyncedCount *int32 `json:"container_repositories_synced_count,omitempty"`

			// ContainerRepositoriesVerificationFailedCount Container repositories verification failed count
			ContainerRepositoriesVerificationFailedCount *int32 `json:"container_repositories_verification_failed_count,omitempty"`

			// ContainerRepositoriesVerificationTotalCount Container repositories verification total count
			ContainerRepositoriesVerificationTotalCount *int32 `json:"container_repositories_verification_total_count,omitempty"`

			// ContainerRepositoriesVerifiedCount Container repositories verified count
			ContainerRepositoriesVerifiedCount *int32 `json:"container_repositories_verified_count,omitempty"`

			// GitFetchEventCountWeekly Git fetch event count weekly
			GitFetchEventCountWeekly *int32 `json:"git_fetch_event_count_weekly,omitempty"`

			// GitPushEventCountWeekly Git push event count weekly
			GitPushEventCountWeekly *int32 `json:"git_push_event_count_weekly,omitempty"`

			// GroupWikiRepositoriesChecksumFailedCount Group wiki repositories checksum failed count
			GroupWikiRepositoriesChecksumFailedCount *int32 `json:"group_wiki_repositories_checksum_failed_count,omitempty"`

			// GroupWikiRepositoriesChecksumTotalCount Group wiki repositories checksum total count
			GroupWikiRepositoriesChecksumTotalCount *int32 `json:"group_wiki_repositories_checksum_total_count,omitempty"`

			// GroupWikiRepositoriesChecksummedCount Group wiki repositories checksummed count
			GroupWikiRepositoriesChecksummedCount *int32 `json:"group_wiki_repositories_checksummed_count,omitempty"`

			// GroupWikiRepositoriesCount Group wiki repositories count
			GroupWikiRepositoriesCount *int32 `json:"group_wiki_repositories_count,omitempty"`

			// GroupWikiRepositoriesFailedCount Group wiki repositories failed count
			GroupWikiRepositoriesFailedCount *int32 `json:"group_wiki_repositories_failed_count,omitempty"`

			// GroupWikiRepositoriesRegistryCount Group wiki repositories registry count
			GroupWikiRepositoriesRegistryCount *int32 `json:"group_wiki_repositories_registry_count,omitempty"`

			// GroupWikiRepositoriesSyncedCount Group wiki repositories synced count
			GroupWikiRepositoriesSyncedCount *int32 `json:"group_wiki_repositories_synced_count,omitempty"`

			// GroupWikiRepositoriesVerificationFailedCount Group wiki repositories verification failed count
			GroupWikiRepositoriesVerificationFailedCount *int32 `json:"group_wiki_repositories_verification_failed_count,omitempty"`

			// GroupWikiRepositoriesVerificationTotalCount Group wiki repositories verification total count
			GroupWikiRepositoriesVerificationTotalCount *int32 `json:"group_wiki_repositories_verification_total_count,omitempty"`

			// GroupWikiRepositoriesVerifiedCount Group wiki repositories verified count
			GroupWikiRepositoriesVerifiedCount *int32 `json:"group_wiki_repositories_verified_count,omitempty"`

			// JobArtifactsChecksumFailedCount Job artifacts checksum failed count
			JobArtifactsChecksumFailedCount *int32 `json:"job_artifacts_checksum_failed_count,omitempty"`

			// JobArtifactsChecksumTotalCount Job artifacts checksum total count
			JobArtifactsChecksumTotalCount *int32 `json:"job_artifacts_checksum_total_count,omitempty"`

			// JobArtifactsChecksummedCount Job artifacts checksummed count
			JobArtifactsChecksummedCount *int32 `json:"job_artifacts_checksummed_count,omitempty"`

			// JobArtifactsCount Job artifacts count
			JobArtifactsCount *int32 `json:"job_artifacts_count,omitempty"`

			// JobArtifactsFailedCount Job artifacts failed count
			JobArtifactsFailedCount *int32 `json:"job_artifacts_failed_count,omitempty"`

			// JobArtifactsRegistryCount Job artifacts registry count
			JobArtifactsRegistryCount *int32 `json:"job_artifacts_registry_count,omitempty"`

			// JobArtifactsSyncedCount Job artifacts synced count
			JobArtifactsSyncedCount *int32 `json:"job_artifacts_synced_count,omitempty"`

			// JobArtifactsVerificationFailedCount Job artifacts verification failed count
			JobArtifactsVerificationFailedCount *int32 `json:"job_artifacts_verification_failed_count,omitempty"`

			// JobArtifactsVerificationTotalCount Job artifacts verification total count
			JobArtifactsVerificationTotalCount *int32 `json:"job_artifacts_verification_total_count,omitempty"`

			// JobArtifactsVerifiedCount Job artifacts verified count
			JobArtifactsVerifiedCount *int32 `json:"job_artifacts_verified_count,omitempty"`

			// LfsObjectsChecksumFailedCount LFS objects checksum failed count
			LfsObjectsChecksumFailedCount *int32 `json:"lfs_objects_checksum_failed_count,omitempty"`

			// LfsObjectsChecksumTotalCount LFS objects checksum total count
			LfsObjectsChecksumTotalCount *int32 `json:"lfs_objects_checksum_total_count,omitempty"`

			// LfsObjectsChecksummedCount LFS objects checksummed count
			LfsObjectsChecksummedCount *int32 `json:"lfs_objects_checksummed_count,omitempty"`

			// LfsObjectsCount LFS objects count
			LfsObjectsCount *int32 `json:"lfs_objects_count,omitempty"`

			// LfsObjectsFailedCount LFS objects failed count
			LfsObjectsFailedCount *int32 `json:"lfs_objects_failed_count,omitempty"`

			// LfsObjectsRegistryCount LFS objects registry count
			LfsObjectsRegistryCount *int32 `json:"lfs_objects_registry_count,omitempty"`

			// LfsObjectsSyncedCount LFS objects synced count
			LfsObjectsSyncedCount *int32 `json:"lfs_objects_synced_count,omitempty"`

			// LfsObjectsVerificationFailedCount LFS objects verification failed count
			LfsObjectsVerificationFailedCount *int32 `json:"lfs_objects_verification_failed_count,omitempty"`

			// LfsObjectsVerificationTotalCount LFS objects verification total count
			LfsObjectsVerificationTotalCount *int32 `json:"lfs_objects_verification_total_count,omitempty"`

			// LfsObjectsVerifiedCount LFS objects verified count
			LfsObjectsVerifiedCount *int32 `json:"lfs_objects_verified_count,omitempty"`

			// MergeRequestDiffsChecksumFailedCount Merge request diffs checksum failed count
			MergeRequestDiffsChecksumFailedCount *int32 `json:"merge_request_diffs_checksum_failed_count,omitempty"`

			// MergeRequestDiffsChecksumTotalCount Merge request diffs checksum total count
			MergeRequestDiffsChecksumTotalCount *int32 `json:"merge_request_diffs_checksum_total_count,omitempty"`

			// MergeRequestDiffsChecksummedCount Merge request diffs checksummed count
			MergeRequestDiffsChecksummedCount *int32 `json:"merge_request_diffs_checksummed_count,omitempty"`

			// MergeRequestDiffsCount Merge request diffs count
			MergeRequestDiffsCount *int32 `json:"merge_request_diffs_count,omitempty"`

			// MergeRequestDiffsFailedCount Merge request diffs failed count
			MergeRequestDiffsFailedCount *int32 `json:"merge_request_diffs_failed_count,omitempty"`

			// MergeRequestDiffsRegistryCount Merge request diffs registry count
			MergeRequestDiffsRegistryCount *int32 `json:"merge_request_diffs_registry_count,omitempty"`

			// MergeRequestDiffsSyncedCount Merge request diffs synced count
			MergeRequestDiffsSyncedCount *int32 `json:"merge_request_diffs_synced_count,omitempty"`

			// MergeRequestDiffsVerificationFailedCount Merge request diffs verified count
			MergeRequestDiffsVerificationFailedCount *int32 `json:"merge_request_diffs_verification_failed_count,omitempty"`

			// MergeRequestDiffsVerificationTotalCount Merge request diffs verification total count
			MergeRequestDiffsVerificationTotalCount *int32 `json:"merge_request_diffs_verification_total_count,omitempty"`

			// MergeRequestDiffsVerifiedCount Merge request diffs verified count
			MergeRequestDiffsVerifiedCount *int32 `json:"merge_request_diffs_verified_count,omitempty"`

			// PackageFilesChecksumFailedCount Packages files checksum failed count
			PackageFilesChecksumFailedCount *int32 `json:"package_files_checksum_failed_count,omitempty"`

			// PackageFilesChecksumTotalCount Packages files checksum total count
			PackageFilesChecksumTotalCount *int32 `json:"package_files_checksum_total_count,omitempty"`

			// PackageFilesChecksummedCount Packages files checksummed count
			PackageFilesChecksummedCount *int32 `json:"package_files_checksummed_count,omitempty"`

			// PackageFilesCount Packages files count
			PackageFilesCount *int32 `json:"package_files_count,omitempty"`

			// PackageFilesFailedCount Packages files failed count
			PackageFilesFailedCount *int32 `json:"package_files_failed_count,omitempty"`

			// PackageFilesRegistryCount Packages files registry count
			PackageFilesRegistryCount *int32 `json:"package_files_registry_count,omitempty"`

			// PackageFilesSyncedCount Packages files synced count
			PackageFilesSyncedCount *int32 `json:"package_files_synced_count,omitempty"`

			// PackageFilesVerificationFailedCount Packages files verification failed count
			PackageFilesVerificationFailedCount *int32 `json:"package_files_verification_failed_count,omitempty"`

			// PackageFilesVerificationTotalCount Packages files verification total count
			PackageFilesVerificationTotalCount *int32 `json:"package_files_verification_total_count,omitempty"`

			// PackageFilesVerifiedCount Packages files verified count
			PackageFilesVerifiedCount *int32 `json:"package_files_verified_count,omitempty"`

			// PagesDeploymentsChecksumFailedCount Pages deployments checksum failed count
			PagesDeploymentsChecksumFailedCount *int32 `json:"pages_deployments_checksum_failed_count,omitempty"`

			// PagesDeploymentsChecksumTotalCount Pages deployments checksum total count
			PagesDeploymentsChecksumTotalCount *int32 `json:"pages_deployments_checksum_total_count,omitempty"`

			// PagesDeploymentsChecksummedCount Pages deployments checksummed count
			PagesDeploymentsChecksummedCount *int32 `json:"pages_deployments_checksummed_count,omitempty"`

			// PagesDeploymentsCount Pages deployments count
			PagesDeploymentsCount *int32 `json:"pages_deployments_count,omitempty"`

			// PagesDeploymentsFailedCount Pages deployments failed count
			PagesDeploymentsFailedCount *int32 `json:"pages_deployments_failed_count,omitempty"`

			// PagesDeploymentsRegistryCount Pages deployments registry count
			PagesDeploymentsRegistryCount *int32 `json:"pages_deployments_registry_count,omitempty"`

			// PagesDeploymentsSyncedCount Pages deployments synced count
			PagesDeploymentsSyncedCount *int32 `json:"pages_deployments_synced_count,omitempty"`

			// PagesDeploymentsVerificationFailedCount Pages deployments verification failed count
			PagesDeploymentsVerificationFailedCount *int32 `json:"pages_deployments_verification_failed_count,omitempty"`

			// PagesDeploymentsVerificationTotalCount Pages deployments verification total count
			PagesDeploymentsVerificationTotalCount *int32 `json:"pages_deployments_verification_total_count,omitempty"`

			// PagesDeploymentsVerifiedCount Pages deployments verified count
			PagesDeploymentsVerifiedCount *int32 `json:"pages_deployments_verified_count,omitempty"`

			// PipelineArtifactsChecksumFailedCount Pipeline artifacts checksum failed count
			PipelineArtifactsChecksumFailedCount *int32 `json:"pipeline_artifacts_checksum_failed_count,omitempty"`

			// PipelineArtifactsChecksumTotalCount Pipeline artifacts checksum total count
			PipelineArtifactsChecksumTotalCount *int32 `json:"pipeline_artifacts_checksum_total_count,omitempty"`

			// PipelineArtifactsChecksummedCount Pipeline artifacts checksummed count
			PipelineArtifactsChecksummedCount *int32 `json:"pipeline_artifacts_checksummed_count,omitempty"`

			// PipelineArtifactsCount Pipeline artifacts count
			PipelineArtifactsCount *int32 `json:"pipeline_artifacts_count,omitempty"`

			// PipelineArtifactsFailedCount Pipeline artifacts failed count
			PipelineArtifactsFailedCount *int32 `json:"pipeline_artifacts_failed_count,omitempty"`

			// PipelineArtifactsRegistryCount Pipeline artifacts registry count
			PipelineArtifactsRegistryCount *int32 `json:"pipeline_artifacts_registry_count,omitempty"`

			// PipelineArtifactsSyncedCount Pipeline artifacts synced count
			PipelineArtifactsSyncedCount *int32 `json:"pipeline_artifacts_synced_count,omitempty"`

			// PipelineArtifactsVerificationFailedCount Pipeline artifacts verification failed count
			PipelineArtifactsVerificationFailedCount *int32 `json:"pipeline_artifacts_verification_failed_count,omitempty"`

			// PipelineArtifactsVerificationTotalCount Pipeline artifacts verification total count
			PipelineArtifactsVerificationTotalCount *int32 `json:"pipeline_artifacts_verification_total_count,omitempty"`

			// PipelineArtifactsVerifiedCount Pipeline artifacts verified count
			PipelineArtifactsVerifiedCount *int32 `json:"pipeline_artifacts_verified_count,omitempty"`

			// ProjectsCount Projects count
			ProjectsCount *int32 `json:"projects_count,omitempty"`

			// ProxyLocalRequestsEventCountWeekly Proxy local requests event count weekly
			ProxyLocalRequestsEventCountWeekly *int32 `json:"proxy_local_requests_event_count_weekly,omitempty"`

			// ProxyRemoteRequestsEventCountWeekly Proxy remote requests event count weekly
			ProxyRemoteRequestsEventCountWeekly *int32 `json:"proxy_remote_requests_event_count_weekly,omitempty"`

			// SnippetRepositoriesChecksumFailedCount Snippet repositories checksum failed count
			SnippetRepositoriesChecksumFailedCount *int32 `json:"snippet_repositories_checksum_failed_count,omitempty"`

			// SnippetRepositoriesChecksumTotalCount Snippet repositories checksum total count
			SnippetRepositoriesChecksumTotalCount *int32 `json:"snippet_repositories_checksum_total_count,omitempty"`

			// SnippetRepositoriesChecksummedCount Snippet repositories checksummed count
			SnippetRepositoriesChecksummedCount *int32 `json:"snippet_repositories_checksummed_count,omitempty"`

			// SnippetRepositoriesCount Snippet repositories count
			SnippetRepositoriesCount *int32 `json:"snippet_repositories_count,omitempty"`

			// SnippetRepositoriesFailedCount Snippet repositories failed count
			SnippetRepositoriesFailedCount *int32 `json:"snippet_repositories_failed_count,omitempty"`

			// SnippetRepositoriesRegistryCount Snippet repositories registry count
			SnippetRepositoriesRegistryCount *int32 `json:"snippet_repositories_registry_count,omitempty"`

			// SnippetRepositoriesSyncedCount Snippet repositories synced count
			SnippetRepositoriesSyncedCount *int32 `json:"snippet_repositories_synced_count,omitempty"`

			// SnippetRepositoriesVerificationFailedCount Snippet repositories verification failed count
			SnippetRepositoriesVerificationFailedCount *int32 `json:"snippet_repositories_verification_failed_count,omitempty"`

			// SnippetRepositoriesVerificationTotalCount Snippet repositories verification total count
			SnippetRepositoriesVerificationTotalCount *int32 `json:"snippet_repositories_verification_total_count,omitempty"`

			// SnippetRepositoriesVerifiedCount Snippet repositories verified count
			SnippetRepositoriesVerifiedCount *int32 `json:"snippet_repositories_verified_count,omitempty"`

			// TerraformStateVersionsChecksumFailedCount Terraform state versions checksum failed count
			TerraformStateVersionsChecksumFailedCount *int32 `json:"terraform_state_versions_checksum_failed_count,omitempty"`

			// TerraformStateVersionsChecksumTotalCount Terraform state versions checksum total count
			TerraformStateVersionsChecksumTotalCount *int32 `json:"terraform_state_versions_checksum_total_count,omitempty"`

			// TerraformStateVersionsChecksummedCount Terraform state versions checksummed count
			TerraformStateVersionsChecksummedCount *int32 `json:"terraform_state_versions_checksummed_count,omitempty"`

			// TerraformStateVersionsCount Terraform state versions count
			TerraformStateVersionsCount *int32 `json:"terraform_state_versions_count,omitempty"`

			// TerraformStateVersionsFailedCount Terraform state versions failed count
			TerraformStateVersionsFailedCount *int32 `json:"terraform_state_versions_failed_count,omitempty"`

			// TerraformStateVersionsRegistryCount Terraform state versions registry count
			TerraformStateVersionsRegistryCount *int32 `json:"terraform_state_versions_registry_count,omitempty"`

			// TerraformStateVersionsSyncedCount Terraform state versions synced count
			TerraformStateVersionsSyncedCount *int32 `json:"terraform_state_versions_synced_count,omitempty"`

			// TerraformStateVersionsVerificationFailedCount Terraform state versions verification failed count
			TerraformStateVersionsVerificationFailedCount *int32 `json:"terraform_state_versions_verification_failed_count,omitempty"`

			// TerraformStateVersionsVerificationTotalCount Terraform state versions verification total count
			TerraformStateVersionsVerificationTotalCount *int32 `json:"terraform_state_versions_verification_total_count,omitempty"`

			// TerraformStateVersionsVerifiedCount Terraform state versions verified count
			TerraformStateVersionsVerifiedCount *int32 `json:"terraform_state_versions_verified_count,omitempty"`

			// UploadsChecksumFailedCount Uploads checksum failed count
			UploadsChecksumFailedCount *int32 `json:"uploads_checksum_failed_count,omitempty"`

			// UploadsChecksumTotalCount Uploads checksum total count
			UploadsChecksumTotalCount *int32 `json:"uploads_checksum_total_count,omitempty"`

			// UploadsChecksummedCount Uploads checksummed count
			UploadsChecksummedCount *int32 `json:"uploads_checksummed_count,omitempty"`

			// UploadsCount Uploads count
			UploadsCount *int32 `json:"uploads_count,omitempty"`

			// UploadsFailedCount Uploads failed count
			UploadsFailedCount *int32 `json:"uploads_failed_count,omitempty"`

			// UploadsRegistryCount Uploads registry count
			UploadsRegistryCount *int32 `json:"uploads_registry_count,omitempty"`

			// UploadsSyncedCount Uploads synced count
			UploadsSyncedCount *int32 `json:"uploads_synced_count,omitempty"`

			// UploadsVerificationFailedCount Uploads verification failed count
			UploadsVerificationFailedCount *int32 `json:"uploads_verification_failed_count,omitempty"`

			// UploadsVerificationTotalCount Uploads verification total count
			UploadsVerificationTotalCount *int32 `json:"uploads_verification_total_count,omitempty"`

			// UploadsVerifiedCount Uploads verified count
			UploadsVerifiedCount *int32 `json:"uploads_verified_count,omitempty"`
		} `json:"status,omitempty"`

		// StatusMessage Status message
		StatusMessage *string `json:"status_message,omitempty"`

		// Version Gitlab version
		Version *string `json:"version,omitempty"`
	} `json:"data,omitempty"`
}
type PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONRequestBody PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONBody
type PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONRequestBody PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONBody
type PostApiV4GeoProxyGitSshReceivePackJSONRequestBody PostApiV4GeoProxyGitSshReceivePackJSONBody
type PostApiV4GeoProxyGitSshUploadPackJSONRequestBody PostApiV4GeoProxyGitSshUploadPackJSONBody
type PostApiV4GeoStatusJSONRequestBody PostApiV4GeoStatusJSONBody
type PostApiV4GeoNodeProxyIdGraphqlResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4GeoProxyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4GeoProxyGitSshReceivePackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4GeoProxyGitSshUploadPackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		PipelineRefs *[]string `json:"pipeline_refs,omitempty"`
	}
}
type GetApiV4GeoRetrieveReplicableNameReplicableIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type PostApiV4GeoStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Links *struct {
			Node *string `json:"node,omitempty"`
			Self *string `json:"self,omitempty"`
		} `json:"_links,omitempty"`
		CiSecureFilesChecksumFailedCount                    *string `json:"ci_secure_files_checksum_failed_count,omitempty"`
		CiSecureFilesChecksumTotalCount                     *string `json:"ci_secure_files_checksum_total_count,omitempty"`
		CiSecureFilesChecksummedCount                       *string `json:"ci_secure_files_checksummed_count,omitempty"`
		CiSecureFilesCount                                  *string `json:"ci_secure_files_count,omitempty"`
		CiSecureFilesFailedCount                            *string `json:"ci_secure_files_failed_count,omitempty"`
		CiSecureFilesRegistryCount                          *string `json:"ci_secure_files_registry_count,omitempty"`
		CiSecureFilesSyncedCount                            *string `json:"ci_secure_files_synced_count,omitempty"`
		CiSecureFilesSyncedInPercentage                     *string `json:"ci_secure_files_synced_in_percentage,omitempty"`
		CiSecureFilesVerificationFailedCount                *string `json:"ci_secure_files_verification_failed_count,omitempty"`
		CiSecureFilesVerificationTotalCount                 *string `json:"ci_secure_files_verification_total_count,omitempty"`
		CiSecureFilesVerifiedCount                          *string `json:"ci_secure_files_verified_count,omitempty"`
		CiSecureFilesVerifiedInPercentage                   *string `json:"ci_secure_files_verified_in_percentage,omitempty"`
		ContainerRepositoriesChecksumFailedCount            *string `json:"container_repositories_checksum_failed_count,omitempty"`
		ContainerRepositoriesChecksumTotalCount             *string `json:"container_repositories_checksum_total_count,omitempty"`
		ContainerRepositoriesChecksummedCount               *string `json:"container_repositories_checksummed_count,omitempty"`
		ContainerRepositoriesCount                          *string `json:"container_repositories_count,omitempty"`
		ContainerRepositoriesFailedCount                    *string `json:"container_repositories_failed_count,omitempty"`
		ContainerRepositoriesRegistryCount                  *string `json:"container_repositories_registry_count,omitempty"`
		ContainerRepositoriesReplicationEnabled             *string `json:"container_repositories_replication_enabled,omitempty"`
		ContainerRepositoriesSyncedCount                    *string `json:"container_repositories_synced_count,omitempty"`
		ContainerRepositoriesSyncedInPercentage             *string `json:"container_repositories_synced_in_percentage,omitempty"`
		ContainerRepositoriesVerificationFailedCount        *string `json:"container_repositories_verification_failed_count,omitempty"`
		ContainerRepositoriesVerificationTotalCount         *string `json:"container_repositories_verification_total_count,omitempty"`
		ContainerRepositoriesVerifiedCount                  *string `json:"container_repositories_verified_count,omitempty"`
		ContainerRepositoriesVerifiedInPercentage           *string `json:"container_repositories_verified_in_percentage,omitempty"`
		CursorLastEventId                                   *string `json:"cursor_last_event_id,omitempty"`
		CursorLastEventTimestamp                            *string `json:"cursor_last_event_timestamp,omitempty"`
		DbReplicationLagSeconds                             *string `json:"db_replication_lag_seconds,omitempty"`
		DependencyProxyBlobsChecksumFailedCount             *string `json:"dependency_proxy_blobs_checksum_failed_count,omitempty"`
		DependencyProxyBlobsChecksumTotalCount              *string `json:"dependency_proxy_blobs_checksum_total_count,omitempty"`
		DependencyProxyBlobsChecksummedCount                *string `json:"dependency_proxy_blobs_checksummed_count,omitempty"`
		DependencyProxyBlobsCount                           *string `json:"dependency_proxy_blobs_count,omitempty"`
		DependencyProxyBlobsFailedCount                     *string `json:"dependency_proxy_blobs_failed_count,omitempty"`
		DependencyProxyBlobsRegistryCount                   *string `json:"dependency_proxy_blobs_registry_count,omitempty"`
		DependencyProxyBlobsSyncedCount                     *string `json:"dependency_proxy_blobs_synced_count,omitempty"`
		DependencyProxyBlobsSyncedInPercentage              *string `json:"dependency_proxy_blobs_synced_in_percentage,omitempty"`
		DependencyProxyBlobsVerificationFailedCount         *string `json:"dependency_proxy_blobs_verification_failed_count,omitempty"`
		DependencyProxyBlobsVerificationTotalCount          *string `json:"dependency_proxy_blobs_verification_total_count,omitempty"`
		DependencyProxyBlobsVerifiedCount                   *string `json:"dependency_proxy_blobs_verified_count,omitempty"`
		DependencyProxyBlobsVerifiedInPercentage            *string `json:"dependency_proxy_blobs_verified_in_percentage,omitempty"`
		DependencyProxyManifestsChecksumFailedCount         *string `json:"dependency_proxy_manifests_checksum_failed_count,omitempty"`
		DependencyProxyManifestsChecksumTotalCount          *string `json:"dependency_proxy_manifests_checksum_total_count,omitempty"`
		DependencyProxyManifestsChecksummedCount            *string `json:"dependency_proxy_manifests_checksummed_count,omitempty"`
		DependencyProxyManifestsCount                       *string `json:"dependency_proxy_manifests_count,omitempty"`
		DependencyProxyManifestsFailedCount                 *string `json:"dependency_proxy_manifests_failed_count,omitempty"`
		DependencyProxyManifestsRegistryCount               *string `json:"dependency_proxy_manifests_registry_count,omitempty"`
		DependencyProxyManifestsSyncedCount                 *string `json:"dependency_proxy_manifests_synced_count,omitempty"`
		DependencyProxyManifestsSyncedInPercentage          *string `json:"dependency_proxy_manifests_synced_in_percentage,omitempty"`
		DependencyProxyManifestsVerificationFailedCount     *string `json:"dependency_proxy_manifests_verification_failed_count,omitempty"`
		DependencyProxyManifestsVerificationTotalCount      *string `json:"dependency_proxy_manifests_verification_total_count,omitempty"`
		DependencyProxyManifestsVerifiedCount               *string `json:"dependency_proxy_manifests_verified_count,omitempty"`
		DependencyProxyManifestsVerifiedInPercentage        *string `json:"dependency_proxy_manifests_verified_in_percentage,omitempty"`
		DesignManagementRepositoriesChecksumFailedCount     *string `json:"design_management_repositories_checksum_failed_count,omitempty"`
		DesignManagementRepositoriesChecksumTotalCount      *string `json:"design_management_repositories_checksum_total_count,omitempty"`
		DesignManagementRepositoriesChecksummedCount        *string `json:"design_management_repositories_checksummed_count,omitempty"`
		DesignManagementRepositoriesCount                   *string `json:"design_management_repositories_count,omitempty"`
		DesignManagementRepositoriesFailedCount             *string `json:"design_management_repositories_failed_count,omitempty"`
		DesignManagementRepositoriesRegistryCount           *string `json:"design_management_repositories_registry_count,omitempty"`
		DesignManagementRepositoriesSyncedCount             *string `json:"design_management_repositories_synced_count,omitempty"`
		DesignManagementRepositoriesSyncedInPercentage      *string `json:"design_management_repositories_synced_in_percentage,omitempty"`
		DesignManagementRepositoriesVerificationFailedCount *string `json:"design_management_repositories_verification_failed_count,omitempty"`
		DesignManagementRepositoriesVerificationTotalCount  *string `json:"design_management_repositories_verification_total_count,omitempty"`
		DesignManagementRepositoriesVerifiedCount           *string `json:"design_management_repositories_verified_count,omitempty"`
		DesignManagementRepositoriesVerifiedInPercentage    *string `json:"design_management_repositories_verified_in_percentage,omitempty"`
		GeoNodeId                                           *string `json:"geo_node_id,omitempty"`
		GitFetchEventCountWeekly                            *string `json:"git_fetch_event_count_weekly,omitempty"`
		GitPushEventCountWeekly                             *string `json:"git_push_event_count_weekly,omitempty"`
		GroupWikiRepositoriesChecksumFailedCount            *string `json:"group_wiki_repositories_checksum_failed_count,omitempty"`
		GroupWikiRepositoriesChecksumTotalCount             *string `json:"group_wiki_repositories_checksum_total_count,omitempty"`
		GroupWikiRepositoriesChecksummedCount               *string `json:"group_wiki_repositories_checksummed_count,omitempty"`
		GroupWikiRepositoriesCount                          *string `json:"group_wiki_repositories_count,omitempty"`
		GroupWikiRepositoriesFailedCount                    *string `json:"group_wiki_repositories_failed_count,omitempty"`
		GroupWikiRepositoriesRegistryCount                  *string `json:"group_wiki_repositories_registry_count,omitempty"`
		GroupWikiRepositoriesSyncedCount                    *string `json:"group_wiki_repositories_synced_count,omitempty"`
		GroupWikiRepositoriesSyncedInPercentage             *string `json:"group_wiki_repositories_synced_in_percentage,omitempty"`
		GroupWikiRepositoriesVerificationFailedCount        *string `json:"group_wiki_repositories_verification_failed_count,omitempty"`
		GroupWikiRepositoriesVerificationTotalCount         *string `json:"group_wiki_repositories_verification_total_count,omitempty"`
		GroupWikiRepositoriesVerifiedCount                  *string `json:"group_wiki_repositories_verified_count,omitempty"`
		GroupWikiRepositoriesVerifiedInPercentage           *string `json:"group_wiki_repositories_verified_in_percentage,omitempty"`
		Health                                              *string `json:"health,omitempty"`
		HealthStatus                                        *string `json:"health_status,omitempty"`
		Healthy                                             *string `json:"healthy,omitempty"`
		JobArtifactsChecksumFailedCount                     *string `json:"job_artifacts_checksum_failed_count,omitempty"`
		JobArtifactsChecksumTotalCount                      *string `json:"job_artifacts_checksum_total_count,omitempty"`
		JobArtifactsChecksummedCount                        *string `json:"job_artifacts_checksummed_count,omitempty"`
		JobArtifactsCount                                   *string `json:"job_artifacts_count,omitempty"`
		JobArtifactsFailedCount                             *string `json:"job_artifacts_failed_count,omitempty"`
		JobArtifactsRegistryCount                           *string `json:"job_artifacts_registry_count,omitempty"`
		JobArtifactsSyncedCount                             *string `json:"job_artifacts_synced_count,omitempty"`
		JobArtifactsSyncedInPercentage                      *string `json:"job_artifacts_synced_in_percentage,omitempty"`
		JobArtifactsVerificationFailedCount                 *string `json:"job_artifacts_verification_failed_count,omitempty"`
		JobArtifactsVerificationTotalCount                  *string `json:"job_artifacts_verification_total_count,omitempty"`
		JobArtifactsVerifiedCount                           *string `json:"job_artifacts_verified_count,omitempty"`
		JobArtifactsVerifiedInPercentage                    *string `json:"job_artifacts_verified_in_percentage,omitempty"`
		LastEventId                                         *string `json:"last_event_id,omitempty"`
		LastEventTimestamp                                  *string `json:"last_event_timestamp,omitempty"`
		LastSuccessfulStatusCheckTimestamp                  *string `json:"last_successful_status_check_timestamp,omitempty"`
		LfsObjectsChecksumFailedCount                       *string `json:"lfs_objects_checksum_failed_count,omitempty"`
		LfsObjectsChecksumTotalCount                        *string `json:"lfs_objects_checksum_total_count,omitempty"`
		LfsObjectsChecksummedCount                          *string `json:"lfs_objects_checksummed_count,omitempty"`
		LfsObjectsCount                                     *string `json:"lfs_objects_count,omitempty"`
		LfsObjectsFailedCount                               *string `json:"lfs_objects_failed_count,omitempty"`
		LfsObjectsRegistryCount                             *string `json:"lfs_objects_registry_count,omitempty"`
		LfsObjectsSyncedCount                               *string `json:"lfs_objects_synced_count,omitempty"`
		LfsObjectsSyncedInPercentage                        *string `json:"lfs_objects_synced_in_percentage,omitempty"`
		LfsObjectsVerificationFailedCount                   *string `json:"lfs_objects_verification_failed_count,omitempty"`
		LfsObjectsVerificationTotalCount                    *string `json:"lfs_objects_verification_total_count,omitempty"`
		LfsObjectsVerifiedCount                             *string `json:"lfs_objects_verified_count,omitempty"`
		LfsObjectsVerifiedInPercentage                      *string `json:"lfs_objects_verified_in_percentage,omitempty"`
		MergeRequestDiffsChecksumFailedCount                *string `json:"merge_request_diffs_checksum_failed_count,omitempty"`
		MergeRequestDiffsChecksumTotalCount                 *string `json:"merge_request_diffs_checksum_total_count,omitempty"`
		MergeRequestDiffsChecksummedCount                   *string `json:"merge_request_diffs_checksummed_count,omitempty"`
		MergeRequestDiffsCount                              *string `json:"merge_request_diffs_count,omitempty"`
		MergeRequestDiffsFailedCount                        *string `json:"merge_request_diffs_failed_count,omitempty"`
		MergeRequestDiffsRegistryCount                      *string `json:"merge_request_diffs_registry_count,omitempty"`
		MergeRequestDiffsSyncedCount                        *string `json:"merge_request_diffs_synced_count,omitempty"`
		MergeRequestDiffsSyncedInPercentage                 *string `json:"merge_request_diffs_synced_in_percentage,omitempty"`
		MergeRequestDiffsVerificationFailedCount            *string `json:"merge_request_diffs_verification_failed_count,omitempty"`
		MergeRequestDiffsVerificationTotalCount             *string `json:"merge_request_diffs_verification_total_count,omitempty"`
		MergeRequestDiffsVerifiedCount                      *string `json:"merge_request_diffs_verified_count,omitempty"`
		MergeRequestDiffsVerifiedInPercentage               *string `json:"merge_request_diffs_verified_in_percentage,omitempty"`
		MissingOauthApplication                             *string `json:"missing_oauth_application,omitempty"`
		Namespaces                                          *struct {
			AvatarUrl *string `json:"avatar_url,omitempty"`
			FullPath  *string `json:"full_path,omitempty"`
			Id        *int32  `json:"id,omitempty"`
			Kind      *string `json:"kind,omitempty"`
			Name      *string `json:"name,omitempty"`
			ParentId  *int32  `json:"parent_id,omitempty"`
			Path      *string `json:"path,omitempty"`
			WebUrl    *string `json:"web_url,omitempty"`
		} `json:"namespaces,omitempty"`
		PackageFilesChecksumFailedCount                *string `json:"package_files_checksum_failed_count,omitempty"`
		PackageFilesChecksumTotalCount                 *string `json:"package_files_checksum_total_count,omitempty"`
		PackageFilesChecksummedCount                   *string `json:"package_files_checksummed_count,omitempty"`
		PackageFilesCount                              *string `json:"package_files_count,omitempty"`
		PackageFilesFailedCount                        *string `json:"package_files_failed_count,omitempty"`
		PackageFilesRegistryCount                      *string `json:"package_files_registry_count,omitempty"`
		PackageFilesSyncedCount                        *string `json:"package_files_synced_count,omitempty"`
		PackageFilesSyncedInPercentage                 *string `json:"package_files_synced_in_percentage,omitempty"`
		PackageFilesVerificationFailedCount            *string `json:"package_files_verification_failed_count,omitempty"`
		PackageFilesVerificationTotalCount             *string `json:"package_files_verification_total_count,omitempty"`
		PackageFilesVerifiedCount                      *string `json:"package_files_verified_count,omitempty"`
		PackageFilesVerifiedInPercentage               *string `json:"package_files_verified_in_percentage,omitempty"`
		PagesDeploymentsChecksumFailedCount            *string `json:"pages_deployments_checksum_failed_count,omitempty"`
		PagesDeploymentsChecksumTotalCount             *string `json:"pages_deployments_checksum_total_count,omitempty"`
		PagesDeploymentsChecksummedCount               *string `json:"pages_deployments_checksummed_count,omitempty"`
		PagesDeploymentsCount                          *string `json:"pages_deployments_count,omitempty"`
		PagesDeploymentsFailedCount                    *string `json:"pages_deployments_failed_count,omitempty"`
		PagesDeploymentsRegistryCount                  *string `json:"pages_deployments_registry_count,omitempty"`
		PagesDeploymentsSyncedCount                    *string `json:"pages_deployments_synced_count,omitempty"`
		PagesDeploymentsSyncedInPercentage             *string `json:"pages_deployments_synced_in_percentage,omitempty"`
		PagesDeploymentsVerificationFailedCount        *string `json:"pages_deployments_verification_failed_count,omitempty"`
		PagesDeploymentsVerificationTotalCount         *string `json:"pages_deployments_verification_total_count,omitempty"`
		PagesDeploymentsVerifiedCount                  *string `json:"pages_deployments_verified_count,omitempty"`
		PagesDeploymentsVerifiedInPercentage           *string `json:"pages_deployments_verified_in_percentage,omitempty"`
		PipelineArtifactsChecksumFailedCount           *string `json:"pipeline_artifacts_checksum_failed_count,omitempty"`
		PipelineArtifactsChecksumTotalCount            *string `json:"pipeline_artifacts_checksum_total_count,omitempty"`
		PipelineArtifactsChecksummedCount              *string `json:"pipeline_artifacts_checksummed_count,omitempty"`
		PipelineArtifactsCount                         *string `json:"pipeline_artifacts_count,omitempty"`
		PipelineArtifactsFailedCount                   *string `json:"pipeline_artifacts_failed_count,omitempty"`
		PipelineArtifactsRegistryCount                 *string `json:"pipeline_artifacts_registry_count,omitempty"`
		PipelineArtifactsSyncedCount                   *string `json:"pipeline_artifacts_synced_count,omitempty"`
		PipelineArtifactsSyncedInPercentage            *string `json:"pipeline_artifacts_synced_in_percentage,omitempty"`
		PipelineArtifactsVerificationFailedCount       *string `json:"pipeline_artifacts_verification_failed_count,omitempty"`
		PipelineArtifactsVerificationTotalCount        *string `json:"pipeline_artifacts_verification_total_count,omitempty"`
		PipelineArtifactsVerifiedCount                 *string `json:"pipeline_artifacts_verified_count,omitempty"`
		PipelineArtifactsVerifiedInPercentage          *string `json:"pipeline_artifacts_verified_in_percentage,omitempty"`
		ProjectRepositoriesChecksumFailedCount         *string `json:"project_repositories_checksum_failed_count,omitempty"`
		ProjectRepositoriesChecksumTotalCount          *string `json:"project_repositories_checksum_total_count,omitempty"`
		ProjectRepositoriesChecksummedCount            *string `json:"project_repositories_checksummed_count,omitempty"`
		ProjectRepositoriesCount                       *string `json:"project_repositories_count,omitempty"`
		ProjectRepositoriesFailedCount                 *string `json:"project_repositories_failed_count,omitempty"`
		ProjectRepositoriesRegistryCount               *string `json:"project_repositories_registry_count,omitempty"`
		ProjectRepositoriesSyncedCount                 *string `json:"project_repositories_synced_count,omitempty"`
		ProjectRepositoriesSyncedInPercentage          *string `json:"project_repositories_synced_in_percentage,omitempty"`
		ProjectRepositoriesVerificationFailedCount     *string `json:"project_repositories_verification_failed_count,omitempty"`
		ProjectRepositoriesVerificationTotalCount      *string `json:"project_repositories_verification_total_count,omitempty"`
		ProjectRepositoriesVerifiedCount               *string `json:"project_repositories_verified_count,omitempty"`
		ProjectRepositoriesVerifiedInPercentage        *string `json:"project_repositories_verified_in_percentage,omitempty"`
		ProjectWikiRepositoriesChecksumFailedCount     *string `json:"project_wiki_repositories_checksum_failed_count,omitempty"`
		ProjectWikiRepositoriesChecksumTotalCount      *string `json:"project_wiki_repositories_checksum_total_count,omitempty"`
		ProjectWikiRepositoriesChecksummedCount        *string `json:"project_wiki_repositories_checksummed_count,omitempty"`
		ProjectWikiRepositoriesCount                   *string `json:"project_wiki_repositories_count,omitempty"`
		ProjectWikiRepositoriesFailedCount             *string `json:"project_wiki_repositories_failed_count,omitempty"`
		ProjectWikiRepositoriesRegistryCount           *string `json:"project_wiki_repositories_registry_count,omitempty"`
		ProjectWikiRepositoriesSyncedCount             *string `json:"project_wiki_repositories_synced_count,omitempty"`
		ProjectWikiRepositoriesSyncedInPercentage      *string `json:"project_wiki_repositories_synced_in_percentage,omitempty"`
		ProjectWikiRepositoriesVerificationFailedCount *string `json:"project_wiki_repositories_verification_failed_count,omitempty"`
		ProjectWikiRepositoriesVerificationTotalCount  *string `json:"project_wiki_repositories_verification_total_count,omitempty"`
		ProjectWikiRepositoriesVerifiedCount           *string `json:"project_wiki_repositories_verified_count,omitempty"`
		ProjectWikiRepositoriesVerifiedInPercentage    *string `json:"project_wiki_repositories_verified_in_percentage,omitempty"`
		ProjectsCount                                  *string `json:"projects_count,omitempty"`
		ProxyLocalRequestsEventCountWeekly             *string `json:"proxy_local_requests_event_count_weekly,omitempty"`
		ProxyRemoteRequestsEventCountWeekly            *string `json:"proxy_remote_requests_event_count_weekly,omitempty"`
		ReplicationSlotsCount                          *string `json:"replication_slots_count,omitempty"`
		ReplicationSlotsMaxRetainedWalBytes            *string `json:"replication_slots_max_retained_wal_bytes,omitempty"`
		ReplicationSlotsUsedCount                      *string `json:"replication_slots_used_count,omitempty"`
		ReplicationSlotsUsedInPercentage               *string `json:"replication_slots_used_in_percentage,omitempty"`
		RepositoriesCheckedCount                       *string `json:"repositories_checked_count,omitempty"`
		RepositoriesCheckedFailedCount                 *string `json:"repositories_checked_failed_count,omitempty"`
		RepositoriesCheckedInPercentage                *string `json:"repositories_checked_in_percentage,omitempty"`
		RepositoriesCount                              *string `json:"repositories_count,omitempty"`
		Revision                                       *string `json:"revision,omitempty"`
		SelectiveSyncType                              *string `json:"selective_sync_type,omitempty"`
		SnippetRepositoriesChecksumFailedCount         *string `json:"snippet_repositories_checksum_failed_count,omitempty"`
		SnippetRepositoriesChecksumTotalCount          *string `json:"snippet_repositories_checksum_total_count,omitempty"`
		SnippetRepositoriesChecksummedCount            *string `json:"snippet_repositories_checksummed_count,omitempty"`
		SnippetRepositoriesCount                       *string `json:"snippet_repositories_count,omitempty"`
		SnippetRepositoriesFailedCount                 *string `json:"snippet_repositories_failed_count,omitempty"`
		SnippetRepositoriesRegistryCount               *string `json:"snippet_repositories_registry_count,omitempty"`
		SnippetRepositoriesSyncedCount                 *string `json:"snippet_repositories_synced_count,omitempty"`
		SnippetRepositoriesSyncedInPercentage          *string `json:"snippet_repositories_synced_in_percentage,omitempty"`
		SnippetRepositoriesVerificationFailedCount     *string `json:"snippet_repositories_verification_failed_count,omitempty"`
		SnippetRepositoriesVerificationTotalCount      *string `json:"snippet_repositories_verification_total_count,omitempty"`
		SnippetRepositoriesVerifiedCount               *string `json:"snippet_repositories_verified_count,omitempty"`
		SnippetRepositoriesVerifiedInPercentage        *string `json:"snippet_repositories_verified_in_percentage,omitempty"`
		StorageShards                                  *struct {
			Name *string `json:"name,omitempty"`
		} `json:"storage_shards,omitempty"`
		StorageShardsMatch                            *string `json:"storage_shards_match,omitempty"`
		TerraformStateVersionsChecksumFailedCount     *string `json:"terraform_state_versions_checksum_failed_count,omitempty"`
		TerraformStateVersionsChecksumTotalCount      *string `json:"terraform_state_versions_checksum_total_count,omitempty"`
		TerraformStateVersionsChecksummedCount        *string `json:"terraform_state_versions_checksummed_count,omitempty"`
		TerraformStateVersionsCount                   *string `json:"terraform_state_versions_count,omitempty"`
		TerraformStateVersionsFailedCount             *string `json:"terraform_state_versions_failed_count,omitempty"`
		TerraformStateVersionsRegistryCount           *string `json:"terraform_state_versions_registry_count,omitempty"`
		TerraformStateVersionsSyncedCount             *string `json:"terraform_state_versions_synced_count,omitempty"`
		TerraformStateVersionsSyncedInPercentage      *string `json:"terraform_state_versions_synced_in_percentage,omitempty"`
		TerraformStateVersionsVerificationFailedCount *string `json:"terraform_state_versions_verification_failed_count,omitempty"`
		TerraformStateVersionsVerificationTotalCount  *string `json:"terraform_state_versions_verification_total_count,omitempty"`
		TerraformStateVersionsVerifiedCount           *string `json:"terraform_state_versions_verified_count,omitempty"`
		TerraformStateVersionsVerifiedInPercentage    *string `json:"terraform_state_versions_verified_in_percentage,omitempty"`
		UpdatedAt                                     *string `json:"updated_at,omitempty"`
		UploadsChecksumFailedCount                    *string `json:"uploads_checksum_failed_count,omitempty"`
		UploadsChecksumTotalCount                     *string `json:"uploads_checksum_total_count,omitempty"`
		UploadsChecksummedCount                       *string `json:"uploads_checksummed_count,omitempty"`
		UploadsCount                                  *string `json:"uploads_count,omitempty"`
		UploadsFailedCount                            *string `json:"uploads_failed_count,omitempty"`
		UploadsRegistryCount                          *string `json:"uploads_registry_count,omitempty"`
		UploadsSyncedCount                            *string `json:"uploads_synced_count,omitempty"`
		UploadsSyncedInPercentage                     *string `json:"uploads_synced_in_percentage,omitempty"`
		UploadsVerificationFailedCount                *string `json:"uploads_verification_failed_count,omitempty"`
		UploadsVerificationTotalCount                 *string `json:"uploads_verification_total_count,omitempty"`
		UploadsVerifiedCount                          *string `json:"uploads_verified_count,omitempty"`
		UploadsVerifiedInPercentage                   *string `json:"uploads_verified_in_percentage,omitempty"`
		Version                                       *string `json:"version,omitempty"`
	}
}

func (c *Client) PostApiV4GeoNodeProxyIdGraphql(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoNodeProxyIdGraphqlRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4GeoProxy(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4GeoProxyRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoProxyGitSshInfoRefsReceivePackWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoProxyGitSshInfoRefsReceivePackRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoProxyGitSshInfoRefsReceivePack(ctx context.Context, body PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoProxyGitSshInfoRefsReceivePackRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoProxyGitSshInfoRefsUploadPackWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoProxyGitSshInfoRefsUploadPackRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoProxyGitSshInfoRefsUploadPack(ctx context.Context, body PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoProxyGitSshInfoRefsUploadPackRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoProxyGitSshReceivePackWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoProxyGitSshReceivePackRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoProxyGitSshReceivePack(ctx context.Context, body PostApiV4GeoProxyGitSshReceivePackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoProxyGitSshReceivePackRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoProxyGitSshUploadPackWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoProxyGitSshUploadPackRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoProxyGitSshUploadPack(ctx context.Context, body PostApiV4GeoProxyGitSshUploadPackJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoProxyGitSshUploadPackRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4GeoRepositoriesGlRepositoryPipelineRefs(ctx context.Context, glRepository string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4GeoRepositoriesGlRepositoryPipelineRefsRequest(c.Server, glRepository)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4GeoRetrieveReplicableNameReplicableId(ctx context.Context, replicableName string, replicableId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4GeoRetrieveReplicableNameReplicableIdRequest(c.Server, replicableName, replicableId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoStatusWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoStatusRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4GeoStatus(ctx context.Context, body PostApiV4GeoStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4GeoStatusRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewPostApiV4GeoNodeProxyIdGraphqlRequest(server string, id int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/node_proxy/%s/graphql", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4GeoProxyRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/proxy")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPostApiV4GeoProxyGitSshInfoRefsReceivePackRequest(server string, body PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4GeoProxyGitSshInfoRefsReceivePackRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4GeoProxyGitSshInfoRefsReceivePackRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/proxy_git_ssh/info_refs_receive_pack")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewPostApiV4GeoProxyGitSshInfoRefsUploadPackRequest(server string, body PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4GeoProxyGitSshInfoRefsUploadPackRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4GeoProxyGitSshInfoRefsUploadPackRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/proxy_git_ssh/info_refs_upload_pack")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewPostApiV4GeoProxyGitSshReceivePackRequest(server string, body PostApiV4GeoProxyGitSshReceivePackJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4GeoProxyGitSshReceivePackRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4GeoProxyGitSshReceivePackRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/proxy_git_ssh/receive_pack")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewPostApiV4GeoProxyGitSshUploadPackRequest(server string, body PostApiV4GeoProxyGitSshUploadPackJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4GeoProxyGitSshUploadPackRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4GeoProxyGitSshUploadPackRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/proxy_git_ssh/upload_pack")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewGetApiV4GeoRepositoriesGlRepositoryPipelineRefsRequest(server string, glRepository string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "gl_repository", runtime.ParamLocationPath, glRepository)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/repositories/%s/pipeline_refs", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4GeoRetrieveReplicableNameReplicableIdRequest(server string, replicableName string, replicableId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "replicable_name", runtime.ParamLocationPath, replicableName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "replicable_id", runtime.ParamLocationPath, replicableId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/retrieve/%s/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPostApiV4GeoStatusRequest(server string, body PostApiV4GeoStatusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4GeoStatusRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4GeoStatusRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/geo/status")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func (r PostApiV4GeoNodeProxyIdGraphqlResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4GeoNodeProxyIdGraphqlResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4GeoProxyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4GeoProxyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4GeoProxyGitSshReceivePackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4GeoProxyGitSshReceivePackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4GeoProxyGitSshUploadPackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4GeoProxyGitSshUploadPackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4GeoRetrieveReplicableNameReplicableIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4GeoRetrieveReplicableNameReplicableIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4GeoStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4GeoStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) PostApiV4GeoNodeProxyIdGraphqlWithResponse(ctx context.Context, id int32, reqEditors ...RequestEditorFn) (*PostApiV4GeoNodeProxyIdGraphqlResponse, error) {
	rsp, err := c.PostApiV4GeoNodeProxyIdGraphql(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoNodeProxyIdGraphqlResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4GeoProxyWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4GeoProxyResponse, error) {
	rsp, err := c.GetApiV4GeoProxy(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4GeoProxyResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoProxyGitSshInfoRefsReceivePackWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse, error) {
	rsp, err := c.PostApiV4GeoProxyGitSshInfoRefsReceivePackWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoProxyGitSshInfoRefsReceivePackResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoProxyGitSshInfoRefsReceivePackWithResponse(ctx context.Context, body PostApiV4GeoProxyGitSshInfoRefsReceivePackJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse, error) {
	rsp, err := c.PostApiV4GeoProxyGitSshInfoRefsReceivePack(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoProxyGitSshInfoRefsReceivePackResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoProxyGitSshInfoRefsUploadPackWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse, error) {
	rsp, err := c.PostApiV4GeoProxyGitSshInfoRefsUploadPackWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoProxyGitSshInfoRefsUploadPackResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoProxyGitSshInfoRefsUploadPackWithResponse(ctx context.Context, body PostApiV4GeoProxyGitSshInfoRefsUploadPackJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse, error) {
	rsp, err := c.PostApiV4GeoProxyGitSshInfoRefsUploadPack(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoProxyGitSshInfoRefsUploadPackResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoProxyGitSshReceivePackWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshReceivePackResponse, error) {
	rsp, err := c.PostApiV4GeoProxyGitSshReceivePackWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoProxyGitSshReceivePackResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoProxyGitSshReceivePackWithResponse(ctx context.Context, body PostApiV4GeoProxyGitSshReceivePackJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshReceivePackResponse, error) {
	rsp, err := c.PostApiV4GeoProxyGitSshReceivePack(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoProxyGitSshReceivePackResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoProxyGitSshUploadPackWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshUploadPackResponse, error) {
	rsp, err := c.PostApiV4GeoProxyGitSshUploadPackWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoProxyGitSshUploadPackResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoProxyGitSshUploadPackWithResponse(ctx context.Context, body PostApiV4GeoProxyGitSshUploadPackJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoProxyGitSshUploadPackResponse, error) {
	rsp, err := c.PostApiV4GeoProxyGitSshUploadPack(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoProxyGitSshUploadPackResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4GeoRepositoriesGlRepositoryPipelineRefsWithResponse(ctx context.Context, glRepository string, reqEditors ...RequestEditorFn) (*GetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse, error) {
	rsp, err := c.GetApiV4GeoRepositoriesGlRepositoryPipelineRefs(ctx, glRepository, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4GeoRetrieveReplicableNameReplicableIdWithResponse(ctx context.Context, replicableName string, replicableId int32, reqEditors ...RequestEditorFn) (*GetApiV4GeoRetrieveReplicableNameReplicableIdResponse, error) {
	rsp, err := c.GetApiV4GeoRetrieveReplicableNameReplicableId(ctx, replicableName, replicableId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4GeoRetrieveReplicableNameReplicableIdResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoStatusWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4GeoStatusResponse, error) {
	rsp, err := c.PostApiV4GeoStatusWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoStatusResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4GeoStatusWithResponse(ctx context.Context, body PostApiV4GeoStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4GeoStatusResponse, error) {
	rsp, err := c.PostApiV4GeoStatus(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4GeoStatusResponse(rsp)
}
func ParsePostApiV4GeoNodeProxyIdGraphqlResponse(rsp *http.Response) (*PostApiV4GeoNodeProxyIdGraphqlResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4GeoNodeProxyIdGraphqlResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4GeoProxyResponse(rsp *http.Response) (*GetApiV4GeoProxyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4GeoProxyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4GeoProxyGitSshInfoRefsReceivePackResponse(rsp *http.Response) (*PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4GeoProxyGitSshInfoRefsReceivePackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4GeoProxyGitSshInfoRefsUploadPackResponse(rsp *http.Response) (*PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4GeoProxyGitSshInfoRefsUploadPackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4GeoProxyGitSshReceivePackResponse(rsp *http.Response) (*PostApiV4GeoProxyGitSshReceivePackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4GeoProxyGitSshReceivePackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4GeoProxyGitSshUploadPackResponse(rsp *http.Response) (*PostApiV4GeoProxyGitSshUploadPackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4GeoProxyGitSshUploadPackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse(rsp *http.Response) (*GetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4GeoRepositoriesGlRepositoryPipelineRefsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			PipelineRefs *[]string `json:"pipeline_refs,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4GeoRetrieveReplicableNameReplicableIdResponse(rsp *http.Response) (*GetApiV4GeoRetrieveReplicableNameReplicableIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4GeoRetrieveReplicableNameReplicableIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParsePostApiV4GeoStatusResponse(rsp *http.Response) (*PostApiV4GeoStatusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4GeoStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Links *struct {
				Node *string `json:"node,omitempty"`
				Self *string `json:"self,omitempty"`
			} `json:"_links,omitempty"`
			CiSecureFilesChecksumFailedCount                    *string `json:"ci_secure_files_checksum_failed_count,omitempty"`
			CiSecureFilesChecksumTotalCount                     *string `json:"ci_secure_files_checksum_total_count,omitempty"`
			CiSecureFilesChecksummedCount                       *string `json:"ci_secure_files_checksummed_count,omitempty"`
			CiSecureFilesCount                                  *string `json:"ci_secure_files_count,omitempty"`
			CiSecureFilesFailedCount                            *string `json:"ci_secure_files_failed_count,omitempty"`
			CiSecureFilesRegistryCount                          *string `json:"ci_secure_files_registry_count,omitempty"`
			CiSecureFilesSyncedCount                            *string `json:"ci_secure_files_synced_count,omitempty"`
			CiSecureFilesSyncedInPercentage                     *string `json:"ci_secure_files_synced_in_percentage,omitempty"`
			CiSecureFilesVerificationFailedCount                *string `json:"ci_secure_files_verification_failed_count,omitempty"`
			CiSecureFilesVerificationTotalCount                 *string `json:"ci_secure_files_verification_total_count,omitempty"`
			CiSecureFilesVerifiedCount                          *string `json:"ci_secure_files_verified_count,omitempty"`
			CiSecureFilesVerifiedInPercentage                   *string `json:"ci_secure_files_verified_in_percentage,omitempty"`
			ContainerRepositoriesChecksumFailedCount            *string `json:"container_repositories_checksum_failed_count,omitempty"`
			ContainerRepositoriesChecksumTotalCount             *string `json:"container_repositories_checksum_total_count,omitempty"`
			ContainerRepositoriesChecksummedCount               *string `json:"container_repositories_checksummed_count,omitempty"`
			ContainerRepositoriesCount                          *string `json:"container_repositories_count,omitempty"`
			ContainerRepositoriesFailedCount                    *string `json:"container_repositories_failed_count,omitempty"`
			ContainerRepositoriesRegistryCount                  *string `json:"container_repositories_registry_count,omitempty"`
			ContainerRepositoriesReplicationEnabled             *string `json:"container_repositories_replication_enabled,omitempty"`
			ContainerRepositoriesSyncedCount                    *string `json:"container_repositories_synced_count,omitempty"`
			ContainerRepositoriesSyncedInPercentage             *string `json:"container_repositories_synced_in_percentage,omitempty"`
			ContainerRepositoriesVerificationFailedCount        *string `json:"container_repositories_verification_failed_count,omitempty"`
			ContainerRepositoriesVerificationTotalCount         *string `json:"container_repositories_verification_total_count,omitempty"`
			ContainerRepositoriesVerifiedCount                  *string `json:"container_repositories_verified_count,omitempty"`
			ContainerRepositoriesVerifiedInPercentage           *string `json:"container_repositories_verified_in_percentage,omitempty"`
			CursorLastEventId                                   *string `json:"cursor_last_event_id,omitempty"`
			CursorLastEventTimestamp                            *string `json:"cursor_last_event_timestamp,omitempty"`
			DbReplicationLagSeconds                             *string `json:"db_replication_lag_seconds,omitempty"`
			DependencyProxyBlobsChecksumFailedCount             *string `json:"dependency_proxy_blobs_checksum_failed_count,omitempty"`
			DependencyProxyBlobsChecksumTotalCount              *string `json:"dependency_proxy_blobs_checksum_total_count,omitempty"`
			DependencyProxyBlobsChecksummedCount                *string `json:"dependency_proxy_blobs_checksummed_count,omitempty"`
			DependencyProxyBlobsCount                           *string `json:"dependency_proxy_blobs_count,omitempty"`
			DependencyProxyBlobsFailedCount                     *string `json:"dependency_proxy_blobs_failed_count,omitempty"`
			DependencyProxyBlobsRegistryCount                   *string `json:"dependency_proxy_blobs_registry_count,omitempty"`
			DependencyProxyBlobsSyncedCount                     *string `json:"dependency_proxy_blobs_synced_count,omitempty"`
			DependencyProxyBlobsSyncedInPercentage              *string `json:"dependency_proxy_blobs_synced_in_percentage,omitempty"`
			DependencyProxyBlobsVerificationFailedCount         *string `json:"dependency_proxy_blobs_verification_failed_count,omitempty"`
			DependencyProxyBlobsVerificationTotalCount          *string `json:"dependency_proxy_blobs_verification_total_count,omitempty"`
			DependencyProxyBlobsVerifiedCount                   *string `json:"dependency_proxy_blobs_verified_count,omitempty"`
			DependencyProxyBlobsVerifiedInPercentage            *string `json:"dependency_proxy_blobs_verified_in_percentage,omitempty"`
			DependencyProxyManifestsChecksumFailedCount         *string `json:"dependency_proxy_manifests_checksum_failed_count,omitempty"`
			DependencyProxyManifestsChecksumTotalCount          *string `json:"dependency_proxy_manifests_checksum_total_count,omitempty"`
			DependencyProxyManifestsChecksummedCount            *string `json:"dependency_proxy_manifests_checksummed_count,omitempty"`
			DependencyProxyManifestsCount                       *string `json:"dependency_proxy_manifests_count,omitempty"`
			DependencyProxyManifestsFailedCount                 *string `json:"dependency_proxy_manifests_failed_count,omitempty"`
			DependencyProxyManifestsRegistryCount               *string `json:"dependency_proxy_manifests_registry_count,omitempty"`
			DependencyProxyManifestsSyncedCount                 *string `json:"dependency_proxy_manifests_synced_count,omitempty"`
			DependencyProxyManifestsSyncedInPercentage          *string `json:"dependency_proxy_manifests_synced_in_percentage,omitempty"`
			DependencyProxyManifestsVerificationFailedCount     *string `json:"dependency_proxy_manifests_verification_failed_count,omitempty"`
			DependencyProxyManifestsVerificationTotalCount      *string `json:"dependency_proxy_manifests_verification_total_count,omitempty"`
			DependencyProxyManifestsVerifiedCount               *string `json:"dependency_proxy_manifests_verified_count,omitempty"`
			DependencyProxyManifestsVerifiedInPercentage        *string `json:"dependency_proxy_manifests_verified_in_percentage,omitempty"`
			DesignManagementRepositoriesChecksumFailedCount     *string `json:"design_management_repositories_checksum_failed_count,omitempty"`
			DesignManagementRepositoriesChecksumTotalCount      *string `json:"design_management_repositories_checksum_total_count,omitempty"`
			DesignManagementRepositoriesChecksummedCount        *string `json:"design_management_repositories_checksummed_count,omitempty"`
			DesignManagementRepositoriesCount                   *string `json:"design_management_repositories_count,omitempty"`
			DesignManagementRepositoriesFailedCount             *string `json:"design_management_repositories_failed_count,omitempty"`
			DesignManagementRepositoriesRegistryCount           *string `json:"design_management_repositories_registry_count,omitempty"`
			DesignManagementRepositoriesSyncedCount             *string `json:"design_management_repositories_synced_count,omitempty"`
			DesignManagementRepositoriesSyncedInPercentage      *string `json:"design_management_repositories_synced_in_percentage,omitempty"`
			DesignManagementRepositoriesVerificationFailedCount *string `json:"design_management_repositories_verification_failed_count,omitempty"`
			DesignManagementRepositoriesVerificationTotalCount  *string `json:"design_management_repositories_verification_total_count,omitempty"`
			DesignManagementRepositoriesVerifiedCount           *string `json:"design_management_repositories_verified_count,omitempty"`
			DesignManagementRepositoriesVerifiedInPercentage    *string `json:"design_management_repositories_verified_in_percentage,omitempty"`
			GeoNodeId                                           *string `json:"geo_node_id,omitempty"`
			GitFetchEventCountWeekly                            *string `json:"git_fetch_event_count_weekly,omitempty"`
			GitPushEventCountWeekly                             *string `json:"git_push_event_count_weekly,omitempty"`
			GroupWikiRepositoriesChecksumFailedCount            *string `json:"group_wiki_repositories_checksum_failed_count,omitempty"`
			GroupWikiRepositoriesChecksumTotalCount             *string `json:"group_wiki_repositories_checksum_total_count,omitempty"`
			GroupWikiRepositoriesChecksummedCount               *string `json:"group_wiki_repositories_checksummed_count,omitempty"`
			GroupWikiRepositoriesCount                          *string `json:"group_wiki_repositories_count,omitempty"`
			GroupWikiRepositoriesFailedCount                    *string `json:"group_wiki_repositories_failed_count,omitempty"`
			GroupWikiRepositoriesRegistryCount                  *string `json:"group_wiki_repositories_registry_count,omitempty"`
			GroupWikiRepositoriesSyncedCount                    *string `json:"group_wiki_repositories_synced_count,omitempty"`
			GroupWikiRepositoriesSyncedInPercentage             *string `json:"group_wiki_repositories_synced_in_percentage,omitempty"`
			GroupWikiRepositoriesVerificationFailedCount        *string `json:"group_wiki_repositories_verification_failed_count,omitempty"`
			GroupWikiRepositoriesVerificationTotalCount         *string `json:"group_wiki_repositories_verification_total_count,omitempty"`
			GroupWikiRepositoriesVerifiedCount                  *string `json:"group_wiki_repositories_verified_count,omitempty"`
			GroupWikiRepositoriesVerifiedInPercentage           *string `json:"group_wiki_repositories_verified_in_percentage,omitempty"`
			Health                                              *string `json:"health,omitempty"`
			HealthStatus                                        *string `json:"health_status,omitempty"`
			Healthy                                             *string `json:"healthy,omitempty"`
			JobArtifactsChecksumFailedCount                     *string `json:"job_artifacts_checksum_failed_count,omitempty"`
			JobArtifactsChecksumTotalCount                      *string `json:"job_artifacts_checksum_total_count,omitempty"`
			JobArtifactsChecksummedCount                        *string `json:"job_artifacts_checksummed_count,omitempty"`
			JobArtifactsCount                                   *string `json:"job_artifacts_count,omitempty"`
			JobArtifactsFailedCount                             *string `json:"job_artifacts_failed_count,omitempty"`
			JobArtifactsRegistryCount                           *string `json:"job_artifacts_registry_count,omitempty"`
			JobArtifactsSyncedCount                             *string `json:"job_artifacts_synced_count,omitempty"`
			JobArtifactsSyncedInPercentage                      *string `json:"job_artifacts_synced_in_percentage,omitempty"`
			JobArtifactsVerificationFailedCount                 *string `json:"job_artifacts_verification_failed_count,omitempty"`
			JobArtifactsVerificationTotalCount                  *string `json:"job_artifacts_verification_total_count,omitempty"`
			JobArtifactsVerifiedCount                           *string `json:"job_artifacts_verified_count,omitempty"`
			JobArtifactsVerifiedInPercentage                    *string `json:"job_artifacts_verified_in_percentage,omitempty"`
			LastEventId                                         *string `json:"last_event_id,omitempty"`
			LastEventTimestamp                                  *string `json:"last_event_timestamp,omitempty"`
			LastSuccessfulStatusCheckTimestamp                  *string `json:"last_successful_status_check_timestamp,omitempty"`
			LfsObjectsChecksumFailedCount                       *string `json:"lfs_objects_checksum_failed_count,omitempty"`
			LfsObjectsChecksumTotalCount                        *string `json:"lfs_objects_checksum_total_count,omitempty"`
			LfsObjectsChecksummedCount                          *string `json:"lfs_objects_checksummed_count,omitempty"`
			LfsObjectsCount                                     *string `json:"lfs_objects_count,omitempty"`
			LfsObjectsFailedCount                               *string `json:"lfs_objects_failed_count,omitempty"`
			LfsObjectsRegistryCount                             *string `json:"lfs_objects_registry_count,omitempty"`
			LfsObjectsSyncedCount                               *string `json:"lfs_objects_synced_count,omitempty"`
			LfsObjectsSyncedInPercentage                        *string `json:"lfs_objects_synced_in_percentage,omitempty"`
			LfsObjectsVerificationFailedCount                   *string `json:"lfs_objects_verification_failed_count,omitempty"`
			LfsObjectsVerificationTotalCount                    *string `json:"lfs_objects_verification_total_count,omitempty"`
			LfsObjectsVerifiedCount                             *string `json:"lfs_objects_verified_count,omitempty"`
			LfsObjectsVerifiedInPercentage                      *string `json:"lfs_objects_verified_in_percentage,omitempty"`
			MergeRequestDiffsChecksumFailedCount                *string `json:"merge_request_diffs_checksum_failed_count,omitempty"`
			MergeRequestDiffsChecksumTotalCount                 *string `json:"merge_request_diffs_checksum_total_count,omitempty"`
			MergeRequestDiffsChecksummedCount                   *string `json:"merge_request_diffs_checksummed_count,omitempty"`
			MergeRequestDiffsCount                              *string `json:"merge_request_diffs_count,omitempty"`
			MergeRequestDiffsFailedCount                        *string `json:"merge_request_diffs_failed_count,omitempty"`
			MergeRequestDiffsRegistryCount                      *string `json:"merge_request_diffs_registry_count,omitempty"`
			MergeRequestDiffsSyncedCount                        *string `json:"merge_request_diffs_synced_count,omitempty"`
			MergeRequestDiffsSyncedInPercentage                 *string `json:"merge_request_diffs_synced_in_percentage,omitempty"`
			MergeRequestDiffsVerificationFailedCount            *string `json:"merge_request_diffs_verification_failed_count,omitempty"`
			MergeRequestDiffsVerificationTotalCount             *string `json:"merge_request_diffs_verification_total_count,omitempty"`
			MergeRequestDiffsVerifiedCount                      *string `json:"merge_request_diffs_verified_count,omitempty"`
			MergeRequestDiffsVerifiedInPercentage               *string `json:"merge_request_diffs_verified_in_percentage,omitempty"`
			MissingOauthApplication                             *string `json:"missing_oauth_application,omitempty"`
			Namespaces                                          *struct {
				AvatarUrl *string `json:"avatar_url,omitempty"`
				FullPath  *string `json:"full_path,omitempty"`
				Id        *int32  `json:"id,omitempty"`
				Kind      *string `json:"kind,omitempty"`
				Name      *string `json:"name,omitempty"`
				ParentId  *int32  `json:"parent_id,omitempty"`
				Path      *string `json:"path,omitempty"`
				WebUrl    *string `json:"web_url,omitempty"`
			} `json:"namespaces,omitempty"`
			PackageFilesChecksumFailedCount                *string `json:"package_files_checksum_failed_count,omitempty"`
			PackageFilesChecksumTotalCount                 *string `json:"package_files_checksum_total_count,omitempty"`
			PackageFilesChecksummedCount                   *string `json:"package_files_checksummed_count,omitempty"`
			PackageFilesCount                              *string `json:"package_files_count,omitempty"`
			PackageFilesFailedCount                        *string `json:"package_files_failed_count,omitempty"`
			PackageFilesRegistryCount                      *string `json:"package_files_registry_count,omitempty"`
			PackageFilesSyncedCount                        *string `json:"package_files_synced_count,omitempty"`
			PackageFilesSyncedInPercentage                 *string `json:"package_files_synced_in_percentage,omitempty"`
			PackageFilesVerificationFailedCount            *string `json:"package_files_verification_failed_count,omitempty"`
			PackageFilesVerificationTotalCount             *string `json:"package_files_verification_total_count,omitempty"`
			PackageFilesVerifiedCount                      *string `json:"package_files_verified_count,omitempty"`
			PackageFilesVerifiedInPercentage               *string `json:"package_files_verified_in_percentage,omitempty"`
			PagesDeploymentsChecksumFailedCount            *string `json:"pages_deployments_checksum_failed_count,omitempty"`
			PagesDeploymentsChecksumTotalCount             *string `json:"pages_deployments_checksum_total_count,omitempty"`
			PagesDeploymentsChecksummedCount               *string `json:"pages_deployments_checksummed_count,omitempty"`
			PagesDeploymentsCount                          *string `json:"pages_deployments_count,omitempty"`
			PagesDeploymentsFailedCount                    *string `json:"pages_deployments_failed_count,omitempty"`
			PagesDeploymentsRegistryCount                  *string `json:"pages_deployments_registry_count,omitempty"`
			PagesDeploymentsSyncedCount                    *string `json:"pages_deployments_synced_count,omitempty"`
			PagesDeploymentsSyncedInPercentage             *string `json:"pages_deployments_synced_in_percentage,omitempty"`
			PagesDeploymentsVerificationFailedCount        *string `json:"pages_deployments_verification_failed_count,omitempty"`
			PagesDeploymentsVerificationTotalCount         *string `json:"pages_deployments_verification_total_count,omitempty"`
			PagesDeploymentsVerifiedCount                  *string `json:"pages_deployments_verified_count,omitempty"`
			PagesDeploymentsVerifiedInPercentage           *string `json:"pages_deployments_verified_in_percentage,omitempty"`
			PipelineArtifactsChecksumFailedCount           *string `json:"pipeline_artifacts_checksum_failed_count,omitempty"`
			PipelineArtifactsChecksumTotalCount            *string `json:"pipeline_artifacts_checksum_total_count,omitempty"`
			PipelineArtifactsChecksummedCount              *string `json:"pipeline_artifacts_checksummed_count,omitempty"`
			PipelineArtifactsCount                         *string `json:"pipeline_artifacts_count,omitempty"`
			PipelineArtifactsFailedCount                   *string `json:"pipeline_artifacts_failed_count,omitempty"`
			PipelineArtifactsRegistryCount                 *string `json:"pipeline_artifacts_registry_count,omitempty"`
			PipelineArtifactsSyncedCount                   *string `json:"pipeline_artifacts_synced_count,omitempty"`
			PipelineArtifactsSyncedInPercentage            *string `json:"pipeline_artifacts_synced_in_percentage,omitempty"`
			PipelineArtifactsVerificationFailedCount       *string `json:"pipeline_artifacts_verification_failed_count,omitempty"`
			PipelineArtifactsVerificationTotalCount        *string `json:"pipeline_artifacts_verification_total_count,omitempty"`
			PipelineArtifactsVerifiedCount                 *string `json:"pipeline_artifacts_verified_count,omitempty"`
			PipelineArtifactsVerifiedInPercentage          *string `json:"pipeline_artifacts_verified_in_percentage,omitempty"`
			ProjectRepositoriesChecksumFailedCount         *string `json:"project_repositories_checksum_failed_count,omitempty"`
			ProjectRepositoriesChecksumTotalCount          *string `json:"project_repositories_checksum_total_count,omitempty"`
			ProjectRepositoriesChecksummedCount            *string `json:"project_repositories_checksummed_count,omitempty"`
			ProjectRepositoriesCount                       *string `json:"project_repositories_count,omitempty"`
			ProjectRepositoriesFailedCount                 *string `json:"project_repositories_failed_count,omitempty"`
			ProjectRepositoriesRegistryCount               *string `json:"project_repositories_registry_count,omitempty"`
			ProjectRepositoriesSyncedCount                 *string `json:"project_repositories_synced_count,omitempty"`
			ProjectRepositoriesSyncedInPercentage          *string `json:"project_repositories_synced_in_percentage,omitempty"`
			ProjectRepositoriesVerificationFailedCount     *string `json:"project_repositories_verification_failed_count,omitempty"`
			ProjectRepositoriesVerificationTotalCount      *string `json:"project_repositories_verification_total_count,omitempty"`
			ProjectRepositoriesVerifiedCount               *string `json:"project_repositories_verified_count,omitempty"`
			ProjectRepositoriesVerifiedInPercentage        *string `json:"project_repositories_verified_in_percentage,omitempty"`
			ProjectWikiRepositoriesChecksumFailedCount     *string `json:"project_wiki_repositories_checksum_failed_count,omitempty"`
			ProjectWikiRepositoriesChecksumTotalCount      *string `json:"project_wiki_repositories_checksum_total_count,omitempty"`
			ProjectWikiRepositoriesChecksummedCount        *string `json:"project_wiki_repositories_checksummed_count,omitempty"`
			ProjectWikiRepositoriesCount                   *string `json:"project_wiki_repositories_count,omitempty"`
			ProjectWikiRepositoriesFailedCount             *string `json:"project_wiki_repositories_failed_count,omitempty"`
			ProjectWikiRepositoriesRegistryCount           *string `json:"project_wiki_repositories_registry_count,omitempty"`
			ProjectWikiRepositoriesSyncedCount             *string `json:"project_wiki_repositories_synced_count,omitempty"`
			ProjectWikiRepositoriesSyncedInPercentage      *string `json:"project_wiki_repositories_synced_in_percentage,omitempty"`
			ProjectWikiRepositoriesVerificationFailedCount *string `json:"project_wiki_repositories_verification_failed_count,omitempty"`
			ProjectWikiRepositoriesVerificationTotalCount  *string `json:"project_wiki_repositories_verification_total_count,omitempty"`
			ProjectWikiRepositoriesVerifiedCount           *string `json:"project_wiki_repositories_verified_count,omitempty"`
			ProjectWikiRepositoriesVerifiedInPercentage    *string `json:"project_wiki_repositories_verified_in_percentage,omitempty"`
			ProjectsCount                                  *string `json:"projects_count,omitempty"`
			ProxyLocalRequestsEventCountWeekly             *string `json:"proxy_local_requests_event_count_weekly,omitempty"`
			ProxyRemoteRequestsEventCountWeekly            *string `json:"proxy_remote_requests_event_count_weekly,omitempty"`
			ReplicationSlotsCount                          *string `json:"replication_slots_count,omitempty"`
			ReplicationSlotsMaxRetainedWalBytes            *string `json:"replication_slots_max_retained_wal_bytes,omitempty"`
			ReplicationSlotsUsedCount                      *string `json:"replication_slots_used_count,omitempty"`
			ReplicationSlotsUsedInPercentage               *string `json:"replication_slots_used_in_percentage,omitempty"`
			RepositoriesCheckedCount                       *string `json:"repositories_checked_count,omitempty"`
			RepositoriesCheckedFailedCount                 *string `json:"repositories_checked_failed_count,omitempty"`
			RepositoriesCheckedInPercentage                *string `json:"repositories_checked_in_percentage,omitempty"`
			RepositoriesCount                              *string `json:"repositories_count,omitempty"`
			Revision                                       *string `json:"revision,omitempty"`
			SelectiveSyncType                              *string `json:"selective_sync_type,omitempty"`
			SnippetRepositoriesChecksumFailedCount         *string `json:"snippet_repositories_checksum_failed_count,omitempty"`
			SnippetRepositoriesChecksumTotalCount          *string `json:"snippet_repositories_checksum_total_count,omitempty"`
			SnippetRepositoriesChecksummedCount            *string `json:"snippet_repositories_checksummed_count,omitempty"`
			SnippetRepositoriesCount                       *string `json:"snippet_repositories_count,omitempty"`
			SnippetRepositoriesFailedCount                 *string `json:"snippet_repositories_failed_count,omitempty"`
			SnippetRepositoriesRegistryCount               *string `json:"snippet_repositories_registry_count,omitempty"`
			SnippetRepositoriesSyncedCount                 *string `json:"snippet_repositories_synced_count,omitempty"`
			SnippetRepositoriesSyncedInPercentage          *string `json:"snippet_repositories_synced_in_percentage,omitempty"`
			SnippetRepositoriesVerificationFailedCount     *string `json:"snippet_repositories_verification_failed_count,omitempty"`
			SnippetRepositoriesVerificationTotalCount      *string `json:"snippet_repositories_verification_total_count,omitempty"`
			SnippetRepositoriesVerifiedCount               *string `json:"snippet_repositories_verified_count,omitempty"`
			SnippetRepositoriesVerifiedInPercentage        *string `json:"snippet_repositories_verified_in_percentage,omitempty"`
			StorageShards                                  *struct {
				Name *string `json:"name,omitempty"`
			} `json:"storage_shards,omitempty"`
			StorageShardsMatch                            *string `json:"storage_shards_match,omitempty"`
			TerraformStateVersionsChecksumFailedCount     *string `json:"terraform_state_versions_checksum_failed_count,omitempty"`
			TerraformStateVersionsChecksumTotalCount      *string `json:"terraform_state_versions_checksum_total_count,omitempty"`
			TerraformStateVersionsChecksummedCount        *string `json:"terraform_state_versions_checksummed_count,omitempty"`
			TerraformStateVersionsCount                   *string `json:"terraform_state_versions_count,omitempty"`
			TerraformStateVersionsFailedCount             *string `json:"terraform_state_versions_failed_count,omitempty"`
			TerraformStateVersionsRegistryCount           *string `json:"terraform_state_versions_registry_count,omitempty"`
			TerraformStateVersionsSyncedCount             *string `json:"terraform_state_versions_synced_count,omitempty"`
			TerraformStateVersionsSyncedInPercentage      *string `json:"terraform_state_versions_synced_in_percentage,omitempty"`
			TerraformStateVersionsVerificationFailedCount *string `json:"terraform_state_versions_verification_failed_count,omitempty"`
			TerraformStateVersionsVerificationTotalCount  *string `json:"terraform_state_versions_verification_total_count,omitempty"`
			TerraformStateVersionsVerifiedCount           *string `json:"terraform_state_versions_verified_count,omitempty"`
			TerraformStateVersionsVerifiedInPercentage    *string `json:"terraform_state_versions_verified_in_percentage,omitempty"`
			UpdatedAt                                     *string `json:"updated_at,omitempty"`
			UploadsChecksumFailedCount                    *string `json:"uploads_checksum_failed_count,omitempty"`
			UploadsChecksumTotalCount                     *string `json:"uploads_checksum_total_count,omitempty"`
			UploadsChecksummedCount                       *string `json:"uploads_checksummed_count,omitempty"`
			UploadsCount                                  *string `json:"uploads_count,omitempty"`
			UploadsFailedCount                            *string `json:"uploads_failed_count,omitempty"`
			UploadsRegistryCount                          *string `json:"uploads_registry_count,omitempty"`
			UploadsSyncedCount                            *string `json:"uploads_synced_count,omitempty"`
			UploadsSyncedInPercentage                     *string `json:"uploads_synced_in_percentage,omitempty"`
			UploadsVerificationFailedCount                *string `json:"uploads_verification_failed_count,omitempty"`
			UploadsVerificationTotalCount                 *string `json:"uploads_verification_total_count,omitempty"`
			UploadsVerifiedCount                          *string `json:"uploads_verified_count,omitempty"`
			UploadsVerifiedInPercentage                   *string `json:"uploads_verified_in_percentage,omitempty"`
			Version                                       *string `json:"version,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
