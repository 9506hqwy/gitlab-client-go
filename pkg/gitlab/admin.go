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

// Defines values for GetApiV4AdminBatchedBackgroundMigrationsParamsDatabase.
const (
	GetApiV4AdminBatchedBackgroundMigrationsParamsDatabaseCi        GetApiV4AdminBatchedBackgroundMigrationsParamsDatabase = "ci"
	GetApiV4AdminBatchedBackgroundMigrationsParamsDatabaseEmbedding GetApiV4AdminBatchedBackgroundMigrationsParamsDatabase = "embedding"
	GetApiV4AdminBatchedBackgroundMigrationsParamsDatabaseGeo       GetApiV4AdminBatchedBackgroundMigrationsParamsDatabase = "geo"
	GetApiV4AdminBatchedBackgroundMigrationsParamsDatabaseMain      GetApiV4AdminBatchedBackgroundMigrationsParamsDatabase = "main"
	GetApiV4AdminBatchedBackgroundMigrationsParamsDatabaseSec       GetApiV4AdminBatchedBackgroundMigrationsParamsDatabase = "sec"
)

// Defines values for GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabase.
const (
	GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabaseCi        GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabase = "ci"
	GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabaseEmbedding GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabase = "embedding"
	GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabaseGeo       GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabase = "geo"
	GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabaseMain      GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabase = "main"
	GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabaseSec       GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabase = "sec"
)

// Defines values for PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabase.
const (
	PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabaseCi        PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabase = "ci"
	PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabaseEmbedding PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabase = "embedding"
	PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabaseGeo       PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabase = "geo"
	PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabaseMain      PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabase = "main"
	PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabaseSec       PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabase = "sec"
)

// Defines values for PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabase.
const (
	PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabaseCi        PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabase = "ci"
	PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabaseEmbedding PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabase = "embedding"
	PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabaseGeo       PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabase = "geo"
	PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabaseMain      PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabase = "main"
	PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabaseSec       PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabase = "sec"
)

// Defines values for PostApiV4AdminCiVariablesJSONBodyVariableType.
const (
	PostApiV4AdminCiVariablesJSONBodyVariableTypeEnvVar PostApiV4AdminCiVariablesJSONBodyVariableType = "env_var"
	PostApiV4AdminCiVariablesJSONBodyVariableTypeFile   PostApiV4AdminCiVariablesJSONBodyVariableType = "file"
)

// Defines values for PutApiV4AdminCiVariablesKeyJSONBodyVariableType.
const (
	PutApiV4AdminCiVariablesKeyJSONBodyVariableTypeEnvVar PutApiV4AdminCiVariablesKeyJSONBodyVariableType = "env_var"
	PutApiV4AdminCiVariablesKeyJSONBodyVariableTypeFile   PutApiV4AdminCiVariablesKeyJSONBodyVariableType = "file"
)

// Defines values for PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationType.
const (
	PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationTypeAbac                 PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationType = "abac"
	PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationTypeRbac                 PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationType = "rbac"
	PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationTypeUnknownAuthorization PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationType = "unknown_authorization"
)

// Defines values for GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseName.
const (
	GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseNameCi   GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseName = "ci"
	GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseNameMain GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseName = "main"
)

// Defines values for PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabase.
const (
	PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabaseCi        PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabase = "ci"
	PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabaseEmbedding PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabase = "embedding"
	PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabaseGeo       PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabase = "geo"
	PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabaseMain      PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabase = "main"
	PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabaseSec       PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabase = "sec"
)

type GetApiV4AdminBatchedBackgroundMigrationsParams struct {
	// Database The name of the database, the default `main`
	Database *GetApiV4AdminBatchedBackgroundMigrationsParamsDatabase `form:"database,omitempty" json:"database,omitempty"`

	// JobClassName Filter migrations by job class name.
	JobClassName *string `form:"job_class_name,omitempty" json:"job_class_name,omitempty"`
}
type GetApiV4AdminBatchedBackgroundMigrationsParamsDatabase string
type GetApiV4AdminBatchedBackgroundMigrationsIdParams struct {
	// Database The name of the database
	Database *GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabase `form:"database,omitempty" json:"database,omitempty"`
}
type GetApiV4AdminBatchedBackgroundMigrationsIdParamsDatabase string
type PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBody struct {
	// Database The name of the database
	Database *PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabase `json:"database,omitempty"`
}
type PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBodyDatabase string
type PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBody struct {
	// Database The name of the database
	Database *PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabase `json:"database,omitempty"`
}
type PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBodyDatabase string
type GetApiV4AdminCiVariablesParams struct {
	// Page Current page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// PerPage Number of items per page
	PerPage *int32 `form:"per_page,omitempty" json:"per_page,omitempty"`
}
type PostApiV4AdminCiVariablesJSONBody struct {
	// Description The description of the variable
	Description *string `json:"description,omitempty"`

	// Key The key of the variable. Max 255 characters
	Key string `json:"key"`

	// Masked Whether the variable is masked
	Masked *bool `json:"masked,omitempty"`

	// Protected Whether the variable is protected
	Protected *bool `json:"protected,omitempty"`

	// Raw Whether the variable will be expanded
	Raw *bool `json:"raw,omitempty"`

	// Value The value of a variable
	Value string `json:"value"`

	// VariableType The type of a variable. Available types are: env_var (default) and file
	VariableType *PostApiV4AdminCiVariablesJSONBodyVariableType `json:"variable_type,omitempty"`
}
type PostApiV4AdminCiVariablesJSONBodyVariableType string
type PutApiV4AdminCiVariablesKeyJSONBody struct {
	// Description The description of the variable
	Description *string `json:"description,omitempty"`

	// Masked Whether the variable is masked
	Masked *bool `json:"masked,omitempty"`

	// Protected Whether the variable is protected
	Protected *bool `json:"protected,omitempty"`

	// Raw Whether the variable will be expanded
	Raw *bool `json:"raw,omitempty"`

	// Value The value of a variable
	Value *string `json:"value,omitempty"`

	// VariableType The type of a variable. Available types are: env_var (default) and file
	VariableType *PutApiV4AdminCiVariablesKeyJSONBodyVariableType `json:"variable_type,omitempty"`
}
type PutApiV4AdminCiVariablesKeyJSONBodyVariableType string
type PostApiV4AdminClustersAddJSONBody struct {
	// Domain Cluster base domain
	Domain *string `json:"domain,omitempty"`

	// Enabled Determines if cluster is active or not, defaults to true
	Enabled *bool `json:"enabled,omitempty"`

	// EnvironmentScope The associated environment to the cluster
	EnvironmentScope *string `json:"environment_scope,omitempty"`

	// Managed Determines if GitLab will manage namespaces and service accounts for this cluster, defaults to true
	Managed *bool `json:"managed,omitempty"`

	// ManagementProjectId The ID of the management project
	ManagementProjectId *int32 `json:"management_project_id,omitempty"`

	// Name Cluster name
	Name string `json:"name"`

	// NamespacePerEnvironment Deploy each environment to a separate Kubernetes namespace
	NamespacePerEnvironment *bool `json:"namespace_per_environment,omitempty"`

	// PlatformKubernetesAttributes Platform Kubernetes data
	PlatformKubernetesAttributes struct {
		// ApiUrl URL to access the Kubernetes API
		ApiUrl string `json:"api_url"`

		// AuthorizationType Cluster authorization type, defaults to RBAC
		AuthorizationType *PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationType `json:"authorization_type,omitempty"`

		// CaCert TLS certificate (needed if API is using a self-signed TLS certificate)
		CaCert *string `json:"ca_cert,omitempty"`

		// Namespace Unique namespace related to Project
		Namespace *string `json:"namespace,omitempty"`

		// Token Token to authenticate against Kubernetes
		Token string `json:"token"`
	} `json:"platform_kubernetes_attributes"`
}
type PostApiV4AdminClustersAddJSONBodyPlatformKubernetesAttributesAuthorizationType string
type PutApiV4AdminClustersClusterIdJSONBody struct {
	// Domain Cluster base domain
	Domain *string `json:"domain,omitempty"`

	// Enabled Enable or disable Gitlab's connection to your Kubernetes cluster
	Enabled *bool `json:"enabled,omitempty"`

	// EnvironmentScope The associated environment to the cluster
	EnvironmentScope *string `json:"environment_scope,omitempty"`

	// Managed Determines if GitLab will manage namespaces and service accounts for this cluster
	Managed *bool `json:"managed,omitempty"`

	// ManagementProjectId The ID of the management project
	ManagementProjectId *int32 `json:"management_project_id,omitempty"`

	// Name Cluster name
	Name *string `json:"name,omitempty"`

	// NamespacePerEnvironment Deploy each environment to a separate Kubernetes namespace
	NamespacePerEnvironment *bool `json:"namespace_per_environment,omitempty"`

	// PlatformKubernetesAttributes Platform Kubernetes data
	PlatformKubernetesAttributes *struct {
		// ApiUrl URL to access the Kubernetes API
		ApiUrl *string `json:"api_url,omitempty"`

		// CaCert TLS certificate (needed if API is using a self-signed TLS certificate)
		CaCert *string `json:"ca_cert,omitempty"`

		// Namespace Unique namespace related to Project
		Namespace *string `json:"namespace,omitempty"`

		// Token Token to authenticate against Kubernetes
		Token *string `json:"token,omitempty"`
	} `json:"platform_kubernetes_attributes,omitempty"`
}
type GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseName string
type PostApiV4AdminMigrationsTimestampMarkJSONBody struct {
	// Database The name of the database
	Database *PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabase `json:"database,omitempty"`
}
type PostApiV4AdminMigrationsTimestampMarkJSONBodyDatabase string
type PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONRequestBody PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONBody
type PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONRequestBody PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONBody
type PostApiV4AdminCiVariablesJSONRequestBody PostApiV4AdminCiVariablesJSONBody
type PutApiV4AdminCiVariablesKeyJSONRequestBody PutApiV4AdminCiVariablesKeyJSONBody
type PostApiV4AdminClustersAddJSONRequestBody PostApiV4AdminClustersAddJSONBody
type PutApiV4AdminClustersClusterIdJSONRequestBody PutApiV4AdminClustersClusterIdJSONBody
type PostApiV4AdminMigrationsTimestampMarkJSONRequestBody PostApiV4AdminMigrationsTimestampMarkJSONBody
type GetApiV4AdminBatchedBackgroundMigrationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		ColumnName   *string    `json:"column_name,omitempty"`
		CreatedAt    *time.Time `json:"created_at,omitempty"`
		Id           *string    `json:"id,omitempty"`
		JobClassName *string    `json:"job_class_name,omitempty"`
		Progress     *float32   `json:"progress,omitempty"`
		Status       *string    `json:"status,omitempty"`
		TableName    *string    `json:"table_name,omitempty"`
	}
}
type GetApiV4AdminBatchedBackgroundMigrationsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		ColumnName   *string    `json:"column_name,omitempty"`
		CreatedAt    *time.Time `json:"created_at,omitempty"`
		Id           *string    `json:"id,omitempty"`
		JobClassName *string    `json:"job_class_name,omitempty"`
		Progress     *float32   `json:"progress,omitempty"`
		Status       *string    `json:"status,omitempty"`
		TableName    *string    `json:"table_name,omitempty"`
	}
}
type PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		ColumnName   *string    `json:"column_name,omitempty"`
		CreatedAt    *time.Time `json:"created_at,omitempty"`
		Id           *string    `json:"id,omitempty"`
		JobClassName *string    `json:"job_class_name,omitempty"`
		Progress     *float32   `json:"progress,omitempty"`
		Status       *string    `json:"status,omitempty"`
		TableName    *string    `json:"table_name,omitempty"`
	}
}
type PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		ColumnName   *string    `json:"column_name,omitempty"`
		CreatedAt    *time.Time `json:"created_at,omitempty"`
		Id           *string    `json:"id,omitempty"`
		JobClassName *string    `json:"job_class_name,omitempty"`
		Progress     *float32   `json:"progress,omitempty"`
		Status       *string    `json:"status,omitempty"`
		TableName    *string    `json:"table_name,omitempty"`
	}
}
type GetApiV4AdminCiVariablesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Description      *string `json:"description,omitempty"`
		EnvironmentScope *string `json:"environment_scope,omitempty"`
		Hidden           *bool   `json:"hidden,omitempty"`
		Key              *string `json:"key,omitempty"`
		Masked           *bool   `json:"masked,omitempty"`
		Protected        *bool   `json:"protected,omitempty"`
		Raw              *bool   `json:"raw,omitempty"`
		Value            *string `json:"value,omitempty"`
		VariableType     *string `json:"variable_type,omitempty"`
	}
}
type PostApiV4AdminCiVariablesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		Description      *string `json:"description,omitempty"`
		EnvironmentScope *string `json:"environment_scope,omitempty"`
		Hidden           *bool   `json:"hidden,omitempty"`
		Key              *string `json:"key,omitempty"`
		Masked           *bool   `json:"masked,omitempty"`
		Protected        *bool   `json:"protected,omitempty"`
		Raw              *bool   `json:"raw,omitempty"`
		Value            *string `json:"value,omitempty"`
		VariableType     *string `json:"variable_type,omitempty"`
	}
}
type DeleteApiV4AdminCiVariablesKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}
type GetApiV4AdminCiVariablesKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Description      *string `json:"description,omitempty"`
		EnvironmentScope *string `json:"environment_scope,omitempty"`
		Hidden           *bool   `json:"hidden,omitempty"`
		Key              *string `json:"key,omitempty"`
		Masked           *bool   `json:"masked,omitempty"`
		Protected        *bool   `json:"protected,omitempty"`
		Raw              *bool   `json:"raw,omitempty"`
		Value            *string `json:"value,omitempty"`
		VariableType     *string `json:"variable_type,omitempty"`
	}
}
type PutApiV4AdminCiVariablesKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Description      *string `json:"description,omitempty"`
		EnvironmentScope *string `json:"environment_scope,omitempty"`
		Hidden           *bool   `json:"hidden,omitempty"`
		Key              *string `json:"key,omitempty"`
		Masked           *bool   `json:"masked,omitempty"`
		Protected        *bool   `json:"protected,omitempty"`
		Raw              *bool   `json:"raw,omitempty"`
		Value            *string `json:"value,omitempty"`
		VariableType     *string `json:"variable_type,omitempty"`
	}
}
type GetApiV4AdminClustersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]struct {
		ClusterType       *string `json:"cluster_type,omitempty"`
		CreatedAt         *string `json:"created_at,omitempty"`
		Domain            *string `json:"domain,omitempty"`
		Enabled           *string `json:"enabled,omitempty"`
		EnvironmentScope  *string `json:"environment_scope,omitempty"`
		Id                *string `json:"id,omitempty"`
		Managed           *string `json:"managed,omitempty"`
		ManagementProject *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"management_project,omitempty"`
		Name                    *string `json:"name,omitempty"`
		NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
		PlatformKubernetes      *struct {
			ApiUrl            *string `json:"api_url,omitempty"`
			AuthorizationType *string `json:"authorization_type,omitempty"`
			CaCert            *string `json:"ca_cert,omitempty"`
			Namespace         *string `json:"namespace,omitempty"`
		} `json:"platform_kubernetes,omitempty"`
		PlatformType *string `json:"platform_type,omitempty"`
		ProviderGcp  *struct {
			ClusterId    *string `json:"cluster_id,omitempty"`
			Endpoint     *string `json:"endpoint,omitempty"`
			GcpProjectId *string `json:"gcp_project_id,omitempty"`
			MachineType  *string `json:"machine_type,omitempty"`
			NumNodes     *string `json:"num_nodes,omitempty"`
			StatusName   *string `json:"status_name,omitempty"`
			Zone         *string `json:"zone,omitempty"`
		} `json:"provider_gcp,omitempty"`
		ProviderType *string `json:"provider_type,omitempty"`

		// User API_Entities_UserBasic model
		User *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"user,omitempty"`
	}
}
type PostApiV4AdminClustersAddResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		ClusterType       *string `json:"cluster_type,omitempty"`
		CreatedAt         *string `json:"created_at,omitempty"`
		Domain            *string `json:"domain,omitempty"`
		Enabled           *string `json:"enabled,omitempty"`
		EnvironmentScope  *string `json:"environment_scope,omitempty"`
		Id                *string `json:"id,omitempty"`
		Managed           *string `json:"managed,omitempty"`
		ManagementProject *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"management_project,omitempty"`
		Name                    *string `json:"name,omitempty"`
		NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
		PlatformKubernetes      *struct {
			ApiUrl            *string `json:"api_url,omitempty"`
			AuthorizationType *string `json:"authorization_type,omitempty"`
			CaCert            *string `json:"ca_cert,omitempty"`
			Namespace         *string `json:"namespace,omitempty"`
		} `json:"platform_kubernetes,omitempty"`
		PlatformType *string `json:"platform_type,omitempty"`
		ProviderGcp  *struct {
			ClusterId    *string `json:"cluster_id,omitempty"`
			Endpoint     *string `json:"endpoint,omitempty"`
			GcpProjectId *string `json:"gcp_project_id,omitempty"`
			MachineType  *string `json:"machine_type,omitempty"`
			NumNodes     *string `json:"num_nodes,omitempty"`
			StatusName   *string `json:"status_name,omitempty"`
			Zone         *string `json:"zone,omitempty"`
		} `json:"provider_gcp,omitempty"`
		ProviderType *string `json:"provider_type,omitempty"`

		// User API_Entities_UserBasic model
		User *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"user,omitempty"`
	}
}
type DeleteApiV4AdminClustersClusterIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON204      *struct {
		ClusterType       *string `json:"cluster_type,omitempty"`
		CreatedAt         *string `json:"created_at,omitempty"`
		Domain            *string `json:"domain,omitempty"`
		Enabled           *string `json:"enabled,omitempty"`
		EnvironmentScope  *string `json:"environment_scope,omitempty"`
		Id                *string `json:"id,omitempty"`
		Managed           *string `json:"managed,omitempty"`
		ManagementProject *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"management_project,omitempty"`
		Name                    *string `json:"name,omitempty"`
		NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
		PlatformKubernetes      *struct {
			ApiUrl            *string `json:"api_url,omitempty"`
			AuthorizationType *string `json:"authorization_type,omitempty"`
			CaCert            *string `json:"ca_cert,omitempty"`
			Namespace         *string `json:"namespace,omitempty"`
		} `json:"platform_kubernetes,omitempty"`
		PlatformType *string `json:"platform_type,omitempty"`
		ProviderGcp  *struct {
			ClusterId    *string `json:"cluster_id,omitempty"`
			Endpoint     *string `json:"endpoint,omitempty"`
			GcpProjectId *string `json:"gcp_project_id,omitempty"`
			MachineType  *string `json:"machine_type,omitempty"`
			NumNodes     *string `json:"num_nodes,omitempty"`
			StatusName   *string `json:"status_name,omitempty"`
			Zone         *string `json:"zone,omitempty"`
		} `json:"provider_gcp,omitempty"`
		ProviderType *string `json:"provider_type,omitempty"`

		// User API_Entities_UserBasic model
		User *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"user,omitempty"`
	}
}
type GetApiV4AdminClustersClusterIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		ClusterType       *string `json:"cluster_type,omitempty"`
		CreatedAt         *string `json:"created_at,omitempty"`
		Domain            *string `json:"domain,omitempty"`
		Enabled           *string `json:"enabled,omitempty"`
		EnvironmentScope  *string `json:"environment_scope,omitempty"`
		Id                *string `json:"id,omitempty"`
		Managed           *string `json:"managed,omitempty"`
		ManagementProject *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"management_project,omitempty"`
		Name                    *string `json:"name,omitempty"`
		NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
		PlatformKubernetes      *struct {
			ApiUrl            *string `json:"api_url,omitempty"`
			AuthorizationType *string `json:"authorization_type,omitempty"`
			CaCert            *string `json:"ca_cert,omitempty"`
			Namespace         *string `json:"namespace,omitempty"`
		} `json:"platform_kubernetes,omitempty"`
		PlatformType *string `json:"platform_type,omitempty"`
		ProviderGcp  *struct {
			ClusterId    *string `json:"cluster_id,omitempty"`
			Endpoint     *string `json:"endpoint,omitempty"`
			GcpProjectId *string `json:"gcp_project_id,omitempty"`
			MachineType  *string `json:"machine_type,omitempty"`
			NumNodes     *string `json:"num_nodes,omitempty"`
			StatusName   *string `json:"status_name,omitempty"`
			Zone         *string `json:"zone,omitempty"`
		} `json:"provider_gcp,omitempty"`
		ProviderType *string `json:"provider_type,omitempty"`

		// User API_Entities_UserBasic model
		User *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"user,omitempty"`
	}
}
type PutApiV4AdminClustersClusterIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		ClusterType       *string `json:"cluster_type,omitempty"`
		CreatedAt         *string `json:"created_at,omitempty"`
		Domain            *string `json:"domain,omitempty"`
		Enabled           *string `json:"enabled,omitempty"`
		EnvironmentScope  *string `json:"environment_scope,omitempty"`
		Id                *string `json:"id,omitempty"`
		Managed           *string `json:"managed,omitempty"`
		ManagementProject *struct {
			CreatedAt         *time.Time `json:"created_at,omitempty"`
			Description       *string    `json:"description,omitempty"`
			Id                *int32     `json:"id,omitempty"`
			Name              *string    `json:"name,omitempty"`
			NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
			Path              *string    `json:"path,omitempty"`
			PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
		} `json:"management_project,omitempty"`
		Name                    *string `json:"name,omitempty"`
		NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
		PlatformKubernetes      *struct {
			ApiUrl            *string `json:"api_url,omitempty"`
			AuthorizationType *string `json:"authorization_type,omitempty"`
			CaCert            *string `json:"ca_cert,omitempty"`
			Namespace         *string `json:"namespace,omitempty"`
		} `json:"platform_kubernetes,omitempty"`
		PlatformType *string `json:"platform_type,omitempty"`
		ProviderGcp  *struct {
			ClusterId    *string `json:"cluster_id,omitempty"`
			Endpoint     *string `json:"endpoint,omitempty"`
			GcpProjectId *string `json:"gcp_project_id,omitempty"`
			MachineType  *string `json:"machine_type,omitempty"`
			NumNodes     *string `json:"num_nodes,omitempty"`
			StatusName   *string `json:"status_name,omitempty"`
			Zone         *string `json:"zone,omitempty"`
		} `json:"provider_gcp,omitempty"`
		ProviderType *string `json:"provider_type,omitempty"`

		// User API_Entities_UserBasic model
		User *struct {
			AvatarPath       *string `json:"avatar_path,omitempty"`
			AvatarUrl        *string `json:"avatar_url,omitempty"`
			CustomAttributes *[]struct {
				Key   *string `json:"key,omitempty"`
				Value *string `json:"value,omitempty"`
			} `json:"custom_attributes,omitempty"`
			Id          *int32  `json:"id,omitempty"`
			Locked      *bool   `json:"locked,omitempty"`
			Name        *string `json:"name,omitempty"`
			PublicEmail *string `json:"public_email,omitempty"`
			State       *string `json:"state,omitempty"`
			Username    *string `json:"username,omitempty"`
			WebUrl      *string `json:"web_url,omitempty"`
		} `json:"user,omitempty"`
	}
}
type GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		FeatureCategories *[]string `json:"feature_categories,omitempty"`
		TableName         *string   `json:"table_name,omitempty"`
	}
}
type PostApiV4AdminMigrationsTimestampMarkResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

func (c *Client) GetApiV4AdminBatchedBackgroundMigrations(ctx context.Context, params *GetApiV4AdminBatchedBackgroundMigrationsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4AdminBatchedBackgroundMigrationsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4AdminBatchedBackgroundMigrationsId(ctx context.Context, id int32, params *GetApiV4AdminBatchedBackgroundMigrationsIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4AdminBatchedBackgroundMigrationsIdRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4AdminBatchedBackgroundMigrationsIdPauseRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4AdminBatchedBackgroundMigrationsIdPause(ctx context.Context, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4AdminBatchedBackgroundMigrationsIdPauseRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithBody(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4AdminBatchedBackgroundMigrationsIdResumeRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4AdminBatchedBackgroundMigrationsIdResume(ctx context.Context, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4AdminBatchedBackgroundMigrationsIdResumeRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4AdminCiVariables(ctx context.Context, params *GetApiV4AdminCiVariablesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4AdminCiVariablesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4AdminCiVariablesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4AdminCiVariablesRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4AdminCiVariables(ctx context.Context, body PostApiV4AdminCiVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4AdminCiVariablesRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4AdminCiVariablesKey(ctx context.Context, key string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4AdminCiVariablesKeyRequest(c.Server, key)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4AdminCiVariablesKey(ctx context.Context, key string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4AdminCiVariablesKeyRequest(c.Server, key)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4AdminCiVariablesKeyWithBody(ctx context.Context, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4AdminCiVariablesKeyRequestWithBody(c.Server, key, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4AdminCiVariablesKey(ctx context.Context, key string, body PutApiV4AdminCiVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4AdminCiVariablesKeyRequest(c.Server, key, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4AdminClusters(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4AdminClustersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4AdminClustersAddWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4AdminClustersAddRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4AdminClustersAdd(ctx context.Context, body PostApiV4AdminClustersAddJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4AdminClustersAddRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) DeleteApiV4AdminClustersClusterId(ctx context.Context, clusterId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV4AdminClustersClusterIdRequest(c.Server, clusterId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4AdminClustersClusterId(ctx context.Context, clusterId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4AdminClustersClusterIdRequest(c.Server, clusterId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4AdminClustersClusterIdWithBody(ctx context.Context, clusterId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4AdminClustersClusterIdRequestWithBody(c.Server, clusterId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PutApiV4AdminClustersClusterId(ctx context.Context, clusterId int32, body PutApiV4AdminClustersClusterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV4AdminClustersClusterIdRequest(c.Server, clusterId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName(ctx context.Context, databaseName GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseName, tableName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameRequest(c.Server, databaseName, tableName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4AdminMigrationsTimestampMarkWithBody(ctx context.Context, timestamp int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4AdminMigrationsTimestampMarkRequestWithBody(c.Server, timestamp, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func (c *Client) PostApiV4AdminMigrationsTimestampMark(ctx context.Context, timestamp int32, body PostApiV4AdminMigrationsTimestampMarkJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV4AdminMigrationsTimestampMarkRequest(c.Server, timestamp, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4AdminBatchedBackgroundMigrationsRequest(server string, params *GetApiV4AdminBatchedBackgroundMigrationsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/batched_background_migrations")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Database != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "database", runtime.ParamLocationQuery, *params.Database); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.JobClassName != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "job_class_name", runtime.ParamLocationQuery, *params.JobClassName); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4AdminBatchedBackgroundMigrationsIdRequest(server string, id int32, params *GetApiV4AdminBatchedBackgroundMigrationsIdParams) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/admin/batched_background_migrations/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Database != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "database", runtime.ParamLocationQuery, *params.Database); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPutApiV4AdminBatchedBackgroundMigrationsIdPauseRequest(server string, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4AdminBatchedBackgroundMigrationsIdPauseRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPutApiV4AdminBatchedBackgroundMigrationsIdPauseRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/admin/batched_background_migrations/%s/pause", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewPutApiV4AdminBatchedBackgroundMigrationsIdResumeRequest(server string, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4AdminBatchedBackgroundMigrationsIdResumeRequestWithBody(server, id, "application/json", bodyReader)
}
func NewPutApiV4AdminBatchedBackgroundMigrationsIdResumeRequestWithBody(server string, id int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/api/v4/admin/batched_background_migrations/%s/resume", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewGetApiV4AdminCiVariablesRequest(server string, params *GetApiV4AdminCiVariablesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/ci/variables")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Page != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.PerPage != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "per_page", runtime.ParamLocationQuery, *params.PerPage); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewPostApiV4AdminCiVariablesRequest(server string, body PostApiV4AdminCiVariablesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4AdminCiVariablesRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4AdminCiVariablesRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/ci/variables")
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
func NewDeleteApiV4AdminCiVariablesKeyRequest(server string, key string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "key", runtime.ParamLocationPath, key)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/ci/variables/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4AdminCiVariablesKeyRequest(server string, key string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "key", runtime.ParamLocationPath, key)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/ci/variables/%s", pathParam0)
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
func NewPutApiV4AdminCiVariablesKeyRequest(server string, key string, body PutApiV4AdminCiVariablesKeyJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4AdminCiVariablesKeyRequestWithBody(server, key, "application/json", bodyReader)
}
func NewPutApiV4AdminCiVariablesKeyRequestWithBody(server string, key string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "key", runtime.ParamLocationPath, key)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/ci/variables/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewGetApiV4AdminClustersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/clusters")
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
func NewPostApiV4AdminClustersAddRequest(server string, body PostApiV4AdminClustersAddJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4AdminClustersAddRequestWithBody(server, "application/json", bodyReader)
}
func NewPostApiV4AdminClustersAddRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/clusters/add")
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
func NewDeleteApiV4AdminClustersClusterIdRequest(server string, clusterId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "cluster_id", runtime.ParamLocationPath, clusterId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/clusters/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func NewGetApiV4AdminClustersClusterIdRequest(server string, clusterId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "cluster_id", runtime.ParamLocationPath, clusterId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/clusters/%s", pathParam0)
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
func NewPutApiV4AdminClustersClusterIdRequest(server string, clusterId int32, body PutApiV4AdminClustersClusterIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV4AdminClustersClusterIdRequestWithBody(server, clusterId, "application/json", bodyReader)
}
func NewPutApiV4AdminClustersClusterIdRequestWithBody(server string, clusterId int32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "cluster_id", runtime.ParamLocationPath, clusterId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/clusters/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}
func NewGetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameRequest(server string, databaseName GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseName, tableName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "database_name", runtime.ParamLocationPath, databaseName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "table_name", runtime.ParamLocationPath, tableName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/databases/%s/dictionary/tables/%s", pathParam0, pathParam1)
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
func NewPostApiV4AdminMigrationsTimestampMarkRequest(server string, timestamp int32, body PostApiV4AdminMigrationsTimestampMarkJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV4AdminMigrationsTimestampMarkRequestWithBody(server, timestamp, "application/json", bodyReader)
}
func NewPostApiV4AdminMigrationsTimestampMarkRequestWithBody(server string, timestamp int32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "timestamp", runtime.ParamLocationPath, timestamp)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/admin/migrations/%s/mark", pathParam0)
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
func (r GetApiV4AdminBatchedBackgroundMigrationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4AdminBatchedBackgroundMigrationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4AdminBatchedBackgroundMigrationsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4AdminBatchedBackgroundMigrationsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4AdminCiVariablesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4AdminCiVariablesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4AdminCiVariablesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4AdminCiVariablesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4AdminCiVariablesKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4AdminCiVariablesKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4AdminCiVariablesKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4AdminCiVariablesKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4AdminCiVariablesKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4AdminCiVariablesKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4AdminClustersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4AdminClustersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4AdminClustersAddResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4AdminClustersAddResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r DeleteApiV4AdminClustersClusterIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r DeleteApiV4AdminClustersClusterIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4AdminClustersClusterIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4AdminClustersClusterIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PutApiV4AdminClustersClusterIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PutApiV4AdminClustersClusterIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (r PostApiV4AdminMigrationsTimestampMarkResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r PostApiV4AdminMigrationsTimestampMarkResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4AdminBatchedBackgroundMigrationsWithResponse(ctx context.Context, params *GetApiV4AdminBatchedBackgroundMigrationsParams, reqEditors ...RequestEditorFn) (*GetApiV4AdminBatchedBackgroundMigrationsResponse, error) {
	rsp, err := c.GetApiV4AdminBatchedBackgroundMigrations(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4AdminBatchedBackgroundMigrationsResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4AdminBatchedBackgroundMigrationsIdWithResponse(ctx context.Context, id int32, params *GetApiV4AdminBatchedBackgroundMigrationsIdParams, reqEditors ...RequestEditorFn) (*GetApiV4AdminBatchedBackgroundMigrationsIdResponse, error) {
	rsp, err := c.GetApiV4AdminBatchedBackgroundMigrationsId(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4AdminBatchedBackgroundMigrationsIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse, error) {
	rsp, err := c.PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4AdminBatchedBackgroundMigrationsIdPauseWithResponse(ctx context.Context, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdPauseJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse, error) {
	rsp, err := c.PutApiV4AdminBatchedBackgroundMigrationsIdPause(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithBodyWithResponse(ctx context.Context, id int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse, error) {
	rsp, err := c.PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4AdminBatchedBackgroundMigrationsIdResumeWithResponse(ctx context.Context, id int32, body PutApiV4AdminBatchedBackgroundMigrationsIdResumeJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse, error) {
	rsp, err := c.PutApiV4AdminBatchedBackgroundMigrationsIdResume(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4AdminCiVariablesWithResponse(ctx context.Context, params *GetApiV4AdminCiVariablesParams, reqEditors ...RequestEditorFn) (*GetApiV4AdminCiVariablesResponse, error) {
	rsp, err := c.GetApiV4AdminCiVariables(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4AdminCiVariablesResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4AdminCiVariablesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4AdminCiVariablesResponse, error) {
	rsp, err := c.PostApiV4AdminCiVariablesWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4AdminCiVariablesResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4AdminCiVariablesWithResponse(ctx context.Context, body PostApiV4AdminCiVariablesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4AdminCiVariablesResponse, error) {
	rsp, err := c.PostApiV4AdminCiVariables(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4AdminCiVariablesResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4AdminCiVariablesKeyWithResponse(ctx context.Context, key string, reqEditors ...RequestEditorFn) (*DeleteApiV4AdminCiVariablesKeyResponse, error) {
	rsp, err := c.DeleteApiV4AdminCiVariablesKey(ctx, key, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4AdminCiVariablesKeyResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4AdminCiVariablesKeyWithResponse(ctx context.Context, key string, reqEditors ...RequestEditorFn) (*GetApiV4AdminCiVariablesKeyResponse, error) {
	rsp, err := c.GetApiV4AdminCiVariablesKey(ctx, key, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4AdminCiVariablesKeyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4AdminCiVariablesKeyWithBodyWithResponse(ctx context.Context, key string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4AdminCiVariablesKeyResponse, error) {
	rsp, err := c.PutApiV4AdminCiVariablesKeyWithBody(ctx, key, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4AdminCiVariablesKeyResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4AdminCiVariablesKeyWithResponse(ctx context.Context, key string, body PutApiV4AdminCiVariablesKeyJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4AdminCiVariablesKeyResponse, error) {
	rsp, err := c.PutApiV4AdminCiVariablesKey(ctx, key, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4AdminCiVariablesKeyResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4AdminClustersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4AdminClustersResponse, error) {
	rsp, err := c.GetApiV4AdminClusters(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4AdminClustersResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4AdminClustersAddWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4AdminClustersAddResponse, error) {
	rsp, err := c.PostApiV4AdminClustersAddWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4AdminClustersAddResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4AdminClustersAddWithResponse(ctx context.Context, body PostApiV4AdminClustersAddJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4AdminClustersAddResponse, error) {
	rsp, err := c.PostApiV4AdminClustersAdd(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4AdminClustersAddResponse(rsp)
}
func (c *ClientWithResponses) DeleteApiV4AdminClustersClusterIdWithResponse(ctx context.Context, clusterId int32, reqEditors ...RequestEditorFn) (*DeleteApiV4AdminClustersClusterIdResponse, error) {
	rsp, err := c.DeleteApiV4AdminClustersClusterId(ctx, clusterId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV4AdminClustersClusterIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4AdminClustersClusterIdWithResponse(ctx context.Context, clusterId int32, reqEditors ...RequestEditorFn) (*GetApiV4AdminClustersClusterIdResponse, error) {
	rsp, err := c.GetApiV4AdminClustersClusterId(ctx, clusterId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4AdminClustersClusterIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4AdminClustersClusterIdWithBodyWithResponse(ctx context.Context, clusterId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV4AdminClustersClusterIdResponse, error) {
	rsp, err := c.PutApiV4AdminClustersClusterIdWithBody(ctx, clusterId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4AdminClustersClusterIdResponse(rsp)
}
func (c *ClientWithResponses) PutApiV4AdminClustersClusterIdWithResponse(ctx context.Context, clusterId int32, body PutApiV4AdminClustersClusterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV4AdminClustersClusterIdResponse, error) {
	rsp, err := c.PutApiV4AdminClustersClusterId(ctx, clusterId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV4AdminClustersClusterIdResponse(rsp)
}
func (c *ClientWithResponses) GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameWithResponse(ctx context.Context, databaseName GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameParamsDatabaseName, tableName string, reqEditors ...RequestEditorFn) (*GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse, error) {
	rsp, err := c.GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableName(ctx, databaseName, tableName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4AdminMigrationsTimestampMarkWithBodyWithResponse(ctx context.Context, timestamp int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV4AdminMigrationsTimestampMarkResponse, error) {
	rsp, err := c.PostApiV4AdminMigrationsTimestampMarkWithBody(ctx, timestamp, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4AdminMigrationsTimestampMarkResponse(rsp)
}
func (c *ClientWithResponses) PostApiV4AdminMigrationsTimestampMarkWithResponse(ctx context.Context, timestamp int32, body PostApiV4AdminMigrationsTimestampMarkJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV4AdminMigrationsTimestampMarkResponse, error) {
	rsp, err := c.PostApiV4AdminMigrationsTimestampMark(ctx, timestamp, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV4AdminMigrationsTimestampMarkResponse(rsp)
}
func ParseGetApiV4AdminBatchedBackgroundMigrationsResponse(rsp *http.Response) (*GetApiV4AdminBatchedBackgroundMigrationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4AdminBatchedBackgroundMigrationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			ColumnName   *string    `json:"column_name,omitempty"`
			CreatedAt    *time.Time `json:"created_at,omitempty"`
			Id           *string    `json:"id,omitempty"`
			JobClassName *string    `json:"job_class_name,omitempty"`
			Progress     *float32   `json:"progress,omitempty"`
			Status       *string    `json:"status,omitempty"`
			TableName    *string    `json:"table_name,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4AdminBatchedBackgroundMigrationsIdResponse(rsp *http.Response) (*GetApiV4AdminBatchedBackgroundMigrationsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4AdminBatchedBackgroundMigrationsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			ColumnName   *string    `json:"column_name,omitempty"`
			CreatedAt    *time.Time `json:"created_at,omitempty"`
			Id           *string    `json:"id,omitempty"`
			JobClassName *string    `json:"job_class_name,omitempty"`
			Progress     *float32   `json:"progress,omitempty"`
			Status       *string    `json:"status,omitempty"`
			TableName    *string    `json:"table_name,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse(rsp *http.Response) (*PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4AdminBatchedBackgroundMigrationsIdPauseResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			ColumnName   *string    `json:"column_name,omitempty"`
			CreatedAt    *time.Time `json:"created_at,omitempty"`
			Id           *string    `json:"id,omitempty"`
			JobClassName *string    `json:"job_class_name,omitempty"`
			Progress     *float32   `json:"progress,omitempty"`
			Status       *string    `json:"status,omitempty"`
			TableName    *string    `json:"table_name,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse(rsp *http.Response) (*PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4AdminBatchedBackgroundMigrationsIdResumeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			ColumnName   *string    `json:"column_name,omitempty"`
			CreatedAt    *time.Time `json:"created_at,omitempty"`
			Id           *string    `json:"id,omitempty"`
			JobClassName *string    `json:"job_class_name,omitempty"`
			Progress     *float32   `json:"progress,omitempty"`
			Status       *string    `json:"status,omitempty"`
			TableName    *string    `json:"table_name,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4AdminCiVariablesResponse(rsp *http.Response) (*GetApiV4AdminCiVariablesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4AdminCiVariablesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Description      *string `json:"description,omitempty"`
			EnvironmentScope *string `json:"environment_scope,omitempty"`
			Hidden           *bool   `json:"hidden,omitempty"`
			Key              *string `json:"key,omitempty"`
			Masked           *bool   `json:"masked,omitempty"`
			Protected        *bool   `json:"protected,omitempty"`
			Raw              *bool   `json:"raw,omitempty"`
			Value            *string `json:"value,omitempty"`
			VariableType     *string `json:"variable_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4AdminCiVariablesResponse(rsp *http.Response) (*PostApiV4AdminCiVariablesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4AdminCiVariablesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			Description      *string `json:"description,omitempty"`
			EnvironmentScope *string `json:"environment_scope,omitempty"`
			Hidden           *bool   `json:"hidden,omitempty"`
			Key              *string `json:"key,omitempty"`
			Masked           *bool   `json:"masked,omitempty"`
			Protected        *bool   `json:"protected,omitempty"`
			Raw              *bool   `json:"raw,omitempty"`
			Value            *string `json:"value,omitempty"`
			VariableType     *string `json:"variable_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4AdminCiVariablesKeyResponse(rsp *http.Response) (*DeleteApiV4AdminCiVariablesKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4AdminCiVariablesKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
func ParseGetApiV4AdminCiVariablesKeyResponse(rsp *http.Response) (*GetApiV4AdminCiVariablesKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4AdminCiVariablesKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Description      *string `json:"description,omitempty"`
			EnvironmentScope *string `json:"environment_scope,omitempty"`
			Hidden           *bool   `json:"hidden,omitempty"`
			Key              *string `json:"key,omitempty"`
			Masked           *bool   `json:"masked,omitempty"`
			Protected        *bool   `json:"protected,omitempty"`
			Raw              *bool   `json:"raw,omitempty"`
			Value            *string `json:"value,omitempty"`
			VariableType     *string `json:"variable_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4AdminCiVariablesKeyResponse(rsp *http.Response) (*PutApiV4AdminCiVariablesKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4AdminCiVariablesKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Description      *string `json:"description,omitempty"`
			EnvironmentScope *string `json:"environment_scope,omitempty"`
			Hidden           *bool   `json:"hidden,omitempty"`
			Key              *string `json:"key,omitempty"`
			Masked           *bool   `json:"masked,omitempty"`
			Protected        *bool   `json:"protected,omitempty"`
			Raw              *bool   `json:"raw,omitempty"`
			Value            *string `json:"value,omitempty"`
			VariableType     *string `json:"variable_type,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4AdminClustersResponse(rsp *http.Response) (*GetApiV4AdminClustersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4AdminClustersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []struct {
			ClusterType       *string `json:"cluster_type,omitempty"`
			CreatedAt         *string `json:"created_at,omitempty"`
			Domain            *string `json:"domain,omitempty"`
			Enabled           *string `json:"enabled,omitempty"`
			EnvironmentScope  *string `json:"environment_scope,omitempty"`
			Id                *string `json:"id,omitempty"`
			Managed           *string `json:"managed,omitempty"`
			ManagementProject *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"management_project,omitempty"`
			Name                    *string `json:"name,omitempty"`
			NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
			PlatformKubernetes      *struct {
				ApiUrl            *string `json:"api_url,omitempty"`
				AuthorizationType *string `json:"authorization_type,omitempty"`
				CaCert            *string `json:"ca_cert,omitempty"`
				Namespace         *string `json:"namespace,omitempty"`
			} `json:"platform_kubernetes,omitempty"`
			PlatformType *string `json:"platform_type,omitempty"`
			ProviderGcp  *struct {
				ClusterId    *string `json:"cluster_id,omitempty"`
				Endpoint     *string `json:"endpoint,omitempty"`
				GcpProjectId *string `json:"gcp_project_id,omitempty"`
				MachineType  *string `json:"machine_type,omitempty"`
				NumNodes     *string `json:"num_nodes,omitempty"`
				StatusName   *string `json:"status_name,omitempty"`
				Zone         *string `json:"zone,omitempty"`
			} `json:"provider_gcp,omitempty"`
			ProviderType *string `json:"provider_type,omitempty"`

			// User API_Entities_UserBasic model
			User *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4AdminClustersAddResponse(rsp *http.Response) (*PostApiV4AdminClustersAddResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4AdminClustersAddResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			ClusterType       *string `json:"cluster_type,omitempty"`
			CreatedAt         *string `json:"created_at,omitempty"`
			Domain            *string `json:"domain,omitempty"`
			Enabled           *string `json:"enabled,omitempty"`
			EnvironmentScope  *string `json:"environment_scope,omitempty"`
			Id                *string `json:"id,omitempty"`
			Managed           *string `json:"managed,omitempty"`
			ManagementProject *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"management_project,omitempty"`
			Name                    *string `json:"name,omitempty"`
			NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
			PlatformKubernetes      *struct {
				ApiUrl            *string `json:"api_url,omitempty"`
				AuthorizationType *string `json:"authorization_type,omitempty"`
				CaCert            *string `json:"ca_cert,omitempty"`
				Namespace         *string `json:"namespace,omitempty"`
			} `json:"platform_kubernetes,omitempty"`
			PlatformType *string `json:"platform_type,omitempty"`
			ProviderGcp  *struct {
				ClusterId    *string `json:"cluster_id,omitempty"`
				Endpoint     *string `json:"endpoint,omitempty"`
				GcpProjectId *string `json:"gcp_project_id,omitempty"`
				MachineType  *string `json:"machine_type,omitempty"`
				NumNodes     *string `json:"num_nodes,omitempty"`
				StatusName   *string `json:"status_name,omitempty"`
				Zone         *string `json:"zone,omitempty"`
			} `json:"provider_gcp,omitempty"`
			ProviderType *string `json:"provider_type,omitempty"`

			// User API_Entities_UserBasic model
			User *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}
func ParseDeleteApiV4AdminClustersClusterIdResponse(rsp *http.Response) (*DeleteApiV4AdminClustersClusterIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV4AdminClustersClusterIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 204:
		var dest struct {
			ClusterType       *string `json:"cluster_type,omitempty"`
			CreatedAt         *string `json:"created_at,omitempty"`
			Domain            *string `json:"domain,omitempty"`
			Enabled           *string `json:"enabled,omitempty"`
			EnvironmentScope  *string `json:"environment_scope,omitempty"`
			Id                *string `json:"id,omitempty"`
			Managed           *string `json:"managed,omitempty"`
			ManagementProject *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"management_project,omitempty"`
			Name                    *string `json:"name,omitempty"`
			NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
			PlatformKubernetes      *struct {
				ApiUrl            *string `json:"api_url,omitempty"`
				AuthorizationType *string `json:"authorization_type,omitempty"`
				CaCert            *string `json:"ca_cert,omitempty"`
				Namespace         *string `json:"namespace,omitempty"`
			} `json:"platform_kubernetes,omitempty"`
			PlatformType *string `json:"platform_type,omitempty"`
			ProviderGcp  *struct {
				ClusterId    *string `json:"cluster_id,omitempty"`
				Endpoint     *string `json:"endpoint,omitempty"`
				GcpProjectId *string `json:"gcp_project_id,omitempty"`
				MachineType  *string `json:"machine_type,omitempty"`
				NumNodes     *string `json:"num_nodes,omitempty"`
				StatusName   *string `json:"status_name,omitempty"`
				Zone         *string `json:"zone,omitempty"`
			} `json:"provider_gcp,omitempty"`
			ProviderType *string `json:"provider_type,omitempty"`

			// User API_Entities_UserBasic model
			User *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON204 = &dest

	}

	return response, nil
}
func ParseGetApiV4AdminClustersClusterIdResponse(rsp *http.Response) (*GetApiV4AdminClustersClusterIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4AdminClustersClusterIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			ClusterType       *string `json:"cluster_type,omitempty"`
			CreatedAt         *string `json:"created_at,omitempty"`
			Domain            *string `json:"domain,omitempty"`
			Enabled           *string `json:"enabled,omitempty"`
			EnvironmentScope  *string `json:"environment_scope,omitempty"`
			Id                *string `json:"id,omitempty"`
			Managed           *string `json:"managed,omitempty"`
			ManagementProject *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"management_project,omitempty"`
			Name                    *string `json:"name,omitempty"`
			NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
			PlatformKubernetes      *struct {
				ApiUrl            *string `json:"api_url,omitempty"`
				AuthorizationType *string `json:"authorization_type,omitempty"`
				CaCert            *string `json:"ca_cert,omitempty"`
				Namespace         *string `json:"namespace,omitempty"`
			} `json:"platform_kubernetes,omitempty"`
			PlatformType *string `json:"platform_type,omitempty"`
			ProviderGcp  *struct {
				ClusterId    *string `json:"cluster_id,omitempty"`
				Endpoint     *string `json:"endpoint,omitempty"`
				GcpProjectId *string `json:"gcp_project_id,omitempty"`
				MachineType  *string `json:"machine_type,omitempty"`
				NumNodes     *string `json:"num_nodes,omitempty"`
				StatusName   *string `json:"status_name,omitempty"`
				Zone         *string `json:"zone,omitempty"`
			} `json:"provider_gcp,omitempty"`
			ProviderType *string `json:"provider_type,omitempty"`

			// User API_Entities_UserBasic model
			User *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePutApiV4AdminClustersClusterIdResponse(rsp *http.Response) (*PutApiV4AdminClustersClusterIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV4AdminClustersClusterIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			ClusterType       *string `json:"cluster_type,omitempty"`
			CreatedAt         *string `json:"created_at,omitempty"`
			Domain            *string `json:"domain,omitempty"`
			Enabled           *string `json:"enabled,omitempty"`
			EnvironmentScope  *string `json:"environment_scope,omitempty"`
			Id                *string `json:"id,omitempty"`
			Managed           *string `json:"managed,omitempty"`
			ManagementProject *struct {
				CreatedAt         *time.Time `json:"created_at,omitempty"`
				Description       *string    `json:"description,omitempty"`
				Id                *int32     `json:"id,omitempty"`
				Name              *string    `json:"name,omitempty"`
				NameWithNamespace *string    `json:"name_with_namespace,omitempty"`
				Path              *string    `json:"path,omitempty"`
				PathWithNamespace *string    `json:"path_with_namespace,omitempty"`
			} `json:"management_project,omitempty"`
			Name                    *string `json:"name,omitempty"`
			NamespacePerEnvironment *string `json:"namespace_per_environment,omitempty"`
			PlatformKubernetes      *struct {
				ApiUrl            *string `json:"api_url,omitempty"`
				AuthorizationType *string `json:"authorization_type,omitempty"`
				CaCert            *string `json:"ca_cert,omitempty"`
				Namespace         *string `json:"namespace,omitempty"`
			} `json:"platform_kubernetes,omitempty"`
			PlatformType *string `json:"platform_type,omitempty"`
			ProviderGcp  *struct {
				ClusterId    *string `json:"cluster_id,omitempty"`
				Endpoint     *string `json:"endpoint,omitempty"`
				GcpProjectId *string `json:"gcp_project_id,omitempty"`
				MachineType  *string `json:"machine_type,omitempty"`
				NumNodes     *string `json:"num_nodes,omitempty"`
				StatusName   *string `json:"status_name,omitempty"`
				Zone         *string `json:"zone,omitempty"`
			} `json:"provider_gcp,omitempty"`
			ProviderType *string `json:"provider_type,omitempty"`

			// User API_Entities_UserBasic model
			User *struct {
				AvatarPath       *string `json:"avatar_path,omitempty"`
				AvatarUrl        *string `json:"avatar_url,omitempty"`
				CustomAttributes *[]struct {
					Key   *string `json:"key,omitempty"`
					Value *string `json:"value,omitempty"`
				} `json:"custom_attributes,omitempty"`
				Id          *int32  `json:"id,omitempty"`
				Locked      *bool   `json:"locked,omitempty"`
				Name        *string `json:"name,omitempty"`
				PublicEmail *string `json:"public_email,omitempty"`
				State       *string `json:"state,omitempty"`
				Username    *string `json:"username,omitempty"`
				WebUrl      *string `json:"web_url,omitempty"`
			} `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParseGetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse(rsp *http.Response) (*GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4AdminDatabasesDatabaseNameDictionaryTablesTableNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			FeatureCategories *[]string `json:"feature_categories,omitempty"`
			TableName         *string   `json:"table_name,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
func ParsePostApiV4AdminMigrationsTimestampMarkResponse(rsp *http.Response) (*PostApiV4AdminMigrationsTimestampMarkResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV4AdminMigrationsTimestampMarkResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
