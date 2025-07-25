/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package constant

// annotations defined by KubeBlocks
const (
	// CRDAPIVersionAnnotationKey indicates the CRD API version of the object.
	CRDAPIVersionAnnotationKey = "kubeblocks.io/crd-api-version"

	KubeBlocksGenerationKey              = "kubeblocks.io/generation"
	ClusterSnapshotAnnotationKey         = "kubeblocks.io/cluster-snapshot"          // ClusterSnapshotAnnotationKey saves the snapshot of cluster.
	EncryptedSystemAccountsAnnotationKey = "kubeblocks.io/encrypted-system-accounts" // EncryptedSystemAccountsAnnotationKey saves the encrypted system accounts.
	OpsRequestAnnotationKey              = "kubeblocks.io/ops-request"               // OpsRequestAnnotationKey OpsRequest annotation key in Cluster
	ReconcileAnnotationKey               = "kubeblocks.io/reconcile"                 // ReconcileAnnotationKey Notify k8s object to reconcile
	RestartAnnotationKey                 = "kubeblocks.io/restart"                   // RestartAnnotationKey the annotation which notices the StatefulSet/DeploySet to restart
	RestoreFromBackupAnnotationKey       = "kubeblocks.io/restore-from-backup"
	RestoreDoneAnnotationKey             = "kubeblocks.io/restore-done"
	BackupSourceTargetAnnotationKey      = "kubeblocks.io/backup-source-target" // RestoreFromBackupAnnotationKey specifies the component to recover from the backup.

	KBAppClusterUIDKey                   = "apps.kubeblocks.io/cluster-uid"
	BackupPolicyTemplateAnnotationKey    = "apps.kubeblocks.io/backup-policy-template"
	PVLastClaimPolicyAnnotationKey       = "apps.kubeblocks.io/pv-last-claim-policy"
	LastRoleSnapshotVersionAnnotationKey = "apps.kubeblocks.io/last-role-snapshot-version"
	ComponentScaleInAnnotationKey        = "apps.kubeblocks.io/component-scale-in" // ComponentScaleInAnnotationKey specifies whether the component is scaled in

	// SkipPreTerminateAnnotationKey specifies to skip the pre-terminate action for a component.
	SkipPreTerminateAnnotationKey = "apps.kubeblocks.io/skip-pre-terminate"

	// SkipImmutableCheckAnnotationKey specifies to skip the mutation check for the object.
	// The mutation check is only applied to the fields that are declared as immutable.
	SkipImmutableCheckAnnotationKey = "apps.kubeblocks.io/skip-immutable-check"

	// NodeSelectorOnceAnnotationKey adds nodeSelector in podSpec for one pod exactly once
	NodeSelectorOnceAnnotationKey = "workloads.kubeblocks.io/node-selector-once"

	PVCNamePrefixAnnotationKey = "apps.kubeblocks.io/pvc-name-prefix"
)

const (
	// SkipBaseBackupRestoreInPitrAnnotationKey is an experimental api to unify full and continuous restore job.
	// It is set on the actionset CR.
	// If this annotaion is set to "true", then only one job will be created during restore.
	SkipBaseBackupRestoreInPitrAnnotationKey = "dataprotection.kubeblocks.io/skip-base-backup-restore-in-pitr"
)

// annotations for multi-cluster
const (
	KBAppMultiClusterPlacementKey   = "apps.kubeblocks.io/multi-cluster-placement"
	MultiClusterServicePlacementKey = "apps.kubeblocks.io/multi-cluster-service-placement"
)

func InheritedAnnotations() []string {
	return []string{
		RestoreFromBackupAnnotationKey,
		BackupSourceTargetAnnotationKey,
		HostNetworkAnnotationKey,
		FeatureReconciliationInCompactModeAnnotationKey,
		KBAppMultiClusterPlacementKey,
	}
}
