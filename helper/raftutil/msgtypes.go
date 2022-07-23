// Code generated by go generate; DO NOT EDIT.
package raftutil

import "github.com/hashicorp/nomad/nomad/structs"

var msgTypeNames = map[structs.MessageType]string{
	structs.NodeRegisterRequestType:                      "NodeRegisterRequestType",
	structs.NodeDeregisterRequestType:                    "NodeDeregisterRequestType",
	structs.NodeUpdateStatusRequestType:                  "NodeUpdateStatusRequestType",
	structs.NodeUpdateDrainRequestType:                   "NodeUpdateDrainRequestType",
	structs.JobRegisterRequestType:                       "JobRegisterRequestType",
	structs.JobDeregisterRequestType:                     "JobDeregisterRequestType",
	structs.EvalUpdateRequestType:                        "EvalUpdateRequestType",
	structs.EvalDeleteRequestType:                        "EvalDeleteRequestType",
	structs.AllocUpdateRequestType:                       "AllocUpdateRequestType",
	structs.AllocClientUpdateRequestType:                 "AllocClientUpdateRequestType",
	structs.ReconcileJobSummariesRequestType:             "ReconcileJobSummariesRequestType",
	structs.VaultAccessorRegisterRequestType:             "VaultAccessorRegisterRequestType",
	structs.VaultAccessorDeregisterRequestType:           "VaultAccessorDeregisterRequestType",
	structs.ApplyPlanResultsRequestType:                  "ApplyPlanResultsRequestType",
	structs.DeploymentStatusUpdateRequestType:            "DeploymentStatusUpdateRequestType",
	structs.DeploymentPromoteRequestType:                 "DeploymentPromoteRequestType",
	structs.DeploymentAllocHealthRequestType:             "DeploymentAllocHealthRequestType",
	structs.DeploymentDeleteRequestType:                  "DeploymentDeleteRequestType",
	structs.JobStabilityRequestType:                      "JobStabilityRequestType",
	structs.ACLPolicyUpsertRequestType:                   "ACLPolicyUpsertRequestType",
	structs.ACLPolicyDeleteRequestType:                   "ACLPolicyDeleteRequestType",
	structs.ACLTokenUpsertRequestType:                    "ACLTokenUpsertRequestType",
	structs.ACLTokenDeleteRequestType:                    "ACLTokenDeleteRequestType",
	structs.ACLTokenBootstrapRequestType:                 "ACLTokenBootstrapRequestType",
	structs.AutopilotRequestType:                         "AutopilotRequestType",
	structs.UpsertNodeEventsType:                         "UpsertNodeEventsType",
	structs.JobBatchDeregisterRequestType:                "JobBatchDeregisterRequestType",
	structs.AllocUpdateDesiredTransitionRequestType:      "AllocUpdateDesiredTransitionRequestType",
	structs.NodeUpdateEligibilityRequestType:             "NodeUpdateEligibilityRequestType",
	structs.BatchNodeUpdateDrainRequestType:              "BatchNodeUpdateDrainRequestType",
	structs.SchedulerConfigRequestType:                   "SchedulerConfigRequestType",
	structs.NodeBatchDeregisterRequestType:               "NodeBatchDeregisterRequestType",
	structs.ClusterMetadataRequestType:                   "ClusterMetadataRequestType",
	structs.ServiceIdentityAccessorRegisterRequestType:   "ServiceIdentityAccessorRegisterRequestType",
	structs.ServiceIdentityAccessorDeregisterRequestType: "ServiceIdentityAccessorDeregisterRequestType",
	structs.CSIVolumeRegisterRequestType:                 "CSIVolumeRegisterRequestType",
	structs.CSIVolumeDeregisterRequestType:               "CSIVolumeDeregisterRequestType",
	structs.CSIVolumeClaimRequestType:                    "CSIVolumeClaimRequestType",
	structs.ScalingEventRegisterRequestType:              "ScalingEventRegisterRequestType",
	structs.CSIVolumeClaimBatchRequestType:               "CSIVolumeClaimBatchRequestType",
	structs.CSIPluginDeleteRequestType:                   "CSIPluginDeleteRequestType",
	structs.EventSinkUpsertRequestType:                   "EventSinkUpsertRequestType",
	structs.EventSinkDeleteRequestType:                   "EventSinkDeleteRequestType",
	structs.BatchEventSinkUpdateProgressType:             "BatchEventSinkUpdateProgressType",
	structs.OneTimeTokenUpsertRequestType:                "OneTimeTokenUpsertRequestType",
	structs.OneTimeTokenDeleteRequestType:                "OneTimeTokenDeleteRequestType",
	structs.OneTimeTokenExpireRequestType:                "OneTimeTokenExpireRequestType",
	structs.ServiceRegistrationUpsertRequestType:         "ServiceRegistrationUpsertRequestType",
	structs.ServiceRegistrationDeleteByIDRequestType:     "ServiceRegistrationDeleteByIDRequestType",
	structs.ServiceRegistrationDeleteByNodeIDRequestType: "ServiceRegistrationDeleteByNodeIDRequestType",
	structs.SVApplyStateRequestType:                      "SVApplyStateRequestType",
	structs.RootKeyMetaUpsertRequestType:                 "RootKeyMetaUpsertRequestType",
	structs.RootKeyMetaDeleteRequestType:                 "RootKeyMetaDeleteRequestType",
	structs.NamespaceUpsertRequestType:                   "NamespaceUpsertRequestType",
	structs.NamespaceDeleteRequestType:                   "NamespaceDeleteRequestType",
}
