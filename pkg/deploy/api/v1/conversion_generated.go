// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	deploy_api "github.com/openshift/origin/pkg/deploy/api"
	api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1_CustomDeploymentStrategyParams_To_api_CustomDeploymentStrategyParams,
		Convert_api_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams,
		Convert_v1_DeploymentCause_To_api_DeploymentCause,
		Convert_api_DeploymentCause_To_v1_DeploymentCause,
		Convert_v1_DeploymentCauseImageTrigger_To_api_DeploymentCauseImageTrigger,
		Convert_api_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger,
		Convert_v1_DeploymentConfig_To_api_DeploymentConfig,
		Convert_api_DeploymentConfig_To_v1_DeploymentConfig,
		Convert_v1_DeploymentConfigList_To_api_DeploymentConfigList,
		Convert_api_DeploymentConfigList_To_v1_DeploymentConfigList,
		Convert_v1_DeploymentConfigRollback_To_api_DeploymentConfigRollback,
		Convert_api_DeploymentConfigRollback_To_v1_DeploymentConfigRollback,
		Convert_v1_DeploymentConfigRollbackSpec_To_api_DeploymentConfigRollbackSpec,
		Convert_api_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec,
		Convert_v1_DeploymentConfigSpec_To_api_DeploymentConfigSpec,
		Convert_api_DeploymentConfigSpec_To_v1_DeploymentConfigSpec,
		Convert_v1_DeploymentConfigStatus_To_api_DeploymentConfigStatus,
		Convert_api_DeploymentConfigStatus_To_v1_DeploymentConfigStatus,
		Convert_v1_DeploymentDetails_To_api_DeploymentDetails,
		Convert_api_DeploymentDetails_To_v1_DeploymentDetails,
		Convert_v1_DeploymentLog_To_api_DeploymentLog,
		Convert_api_DeploymentLog_To_v1_DeploymentLog,
		Convert_v1_DeploymentLogOptions_To_api_DeploymentLogOptions,
		Convert_api_DeploymentLogOptions_To_v1_DeploymentLogOptions,
		Convert_v1_DeploymentStrategy_To_api_DeploymentStrategy,
		Convert_api_DeploymentStrategy_To_v1_DeploymentStrategy,
		Convert_v1_DeploymentTriggerImageChangeParams_To_api_DeploymentTriggerImageChangeParams,
		Convert_api_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams,
		Convert_v1_DeploymentTriggerPolicy_To_api_DeploymentTriggerPolicy,
		Convert_api_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy,
		Convert_v1_ExecNewPodHook_To_api_ExecNewPodHook,
		Convert_api_ExecNewPodHook_To_v1_ExecNewPodHook,
		Convert_v1_LifecycleHook_To_api_LifecycleHook,
		Convert_api_LifecycleHook_To_v1_LifecycleHook,
		Convert_v1_RecreateDeploymentStrategyParams_To_api_RecreateDeploymentStrategyParams,
		Convert_api_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams,
		Convert_v1_RollingDeploymentStrategyParams_To_api_RollingDeploymentStrategyParams,
		Convert_api_RollingDeploymentStrategyParams_To_v1_RollingDeploymentStrategyParams,
		Convert_v1_TagImageHook_To_api_TagImageHook,
		Convert_api_TagImageHook_To_v1_TagImageHook,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1_CustomDeploymentStrategyParams_To_api_CustomDeploymentStrategyParams(in *CustomDeploymentStrategyParams, out *deploy_api.CustomDeploymentStrategyParams, s conversion.Scope) error {
	out.Image = in.Image
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]api.EnvVar, len(*in))
		for i := range *in {
			if err := api_v1.Convert_v1_EnvVar_To_api_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Environment = nil
	}
	out.Command = in.Command
	return nil
}

func Convert_v1_CustomDeploymentStrategyParams_To_api_CustomDeploymentStrategyParams(in *CustomDeploymentStrategyParams, out *deploy_api.CustomDeploymentStrategyParams, s conversion.Scope) error {
	return autoConvert_v1_CustomDeploymentStrategyParams_To_api_CustomDeploymentStrategyParams(in, out, s)
}

func autoConvert_api_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(in *deploy_api.CustomDeploymentStrategyParams, out *CustomDeploymentStrategyParams, s conversion.Scope) error {
	out.Image = in.Image
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]api_v1.EnvVar, len(*in))
		for i := range *in {
			if err := api_v1.Convert_api_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Environment = nil
	}
	out.Command = in.Command
	return nil
}

func Convert_api_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(in *deploy_api.CustomDeploymentStrategyParams, out *CustomDeploymentStrategyParams, s conversion.Scope) error {
	return autoConvert_api_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(in, out, s)
}

func autoConvert_v1_DeploymentCause_To_api_DeploymentCause(in *DeploymentCause, out *deploy_api.DeploymentCause, s conversion.Scope) error {
	out.Type = deploy_api.DeploymentTriggerType(in.Type)
	if in.ImageTrigger != nil {
		in, out := &in.ImageTrigger, &out.ImageTrigger
		*out = new(deploy_api.DeploymentCauseImageTrigger)
		if err := Convert_v1_DeploymentCauseImageTrigger_To_api_DeploymentCauseImageTrigger(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageTrigger = nil
	}
	return nil
}

func Convert_v1_DeploymentCause_To_api_DeploymentCause(in *DeploymentCause, out *deploy_api.DeploymentCause, s conversion.Scope) error {
	return autoConvert_v1_DeploymentCause_To_api_DeploymentCause(in, out, s)
}

func autoConvert_api_DeploymentCause_To_v1_DeploymentCause(in *deploy_api.DeploymentCause, out *DeploymentCause, s conversion.Scope) error {
	out.Type = DeploymentTriggerType(in.Type)
	if in.ImageTrigger != nil {
		in, out := &in.ImageTrigger, &out.ImageTrigger
		*out = new(DeploymentCauseImageTrigger)
		if err := Convert_api_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageTrigger = nil
	}
	return nil
}

func Convert_api_DeploymentCause_To_v1_DeploymentCause(in *deploy_api.DeploymentCause, out *DeploymentCause, s conversion.Scope) error {
	return autoConvert_api_DeploymentCause_To_v1_DeploymentCause(in, out, s)
}

func autoConvert_v1_DeploymentCauseImageTrigger_To_api_DeploymentCauseImageTrigger(in *DeploymentCauseImageTrigger, out *deploy_api.DeploymentCauseImageTrigger, s conversion.Scope) error {
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_DeploymentCauseImageTrigger_To_api_DeploymentCauseImageTrigger(in *DeploymentCauseImageTrigger, out *deploy_api.DeploymentCauseImageTrigger, s conversion.Scope) error {
	return autoConvert_v1_DeploymentCauseImageTrigger_To_api_DeploymentCauseImageTrigger(in, out, s)
}

func autoConvert_api_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(in *deploy_api.DeploymentCauseImageTrigger, out *DeploymentCauseImageTrigger, s conversion.Scope) error {
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(in *deploy_api.DeploymentCauseImageTrigger, out *DeploymentCauseImageTrigger, s conversion.Scope) error {
	return autoConvert_api_DeploymentCauseImageTrigger_To_v1_DeploymentCauseImageTrigger(in, out, s)
}

func autoConvert_v1_DeploymentConfig_To_api_DeploymentConfig(in *DeploymentConfig, out *deploy_api.DeploymentConfig, s conversion.Scope) error {
	SetDefaults_DeploymentConfig(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_DeploymentConfigSpec_To_api_DeploymentConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_DeploymentConfigStatus_To_api_DeploymentConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_DeploymentConfig_To_api_DeploymentConfig(in *DeploymentConfig, out *deploy_api.DeploymentConfig, s conversion.Scope) error {
	return autoConvert_v1_DeploymentConfig_To_api_DeploymentConfig(in, out, s)
}

func autoConvert_api_DeploymentConfig_To_v1_DeploymentConfig(in *deploy_api.DeploymentConfig, out *DeploymentConfig, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_api_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_api_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_DeploymentConfig_To_v1_DeploymentConfig(in *deploy_api.DeploymentConfig, out *DeploymentConfig, s conversion.Scope) error {
	return autoConvert_api_DeploymentConfig_To_v1_DeploymentConfig(in, out, s)
}

func autoConvert_v1_DeploymentConfigList_To_api_DeploymentConfigList(in *DeploymentConfigList, out *deploy_api.DeploymentConfigList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]deploy_api.DeploymentConfig, len(*in))
		for i := range *in {
			if err := Convert_v1_DeploymentConfig_To_api_DeploymentConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_DeploymentConfigList_To_api_DeploymentConfigList(in *DeploymentConfigList, out *deploy_api.DeploymentConfigList, s conversion.Scope) error {
	return autoConvert_v1_DeploymentConfigList_To_api_DeploymentConfigList(in, out, s)
}

func autoConvert_api_DeploymentConfigList_To_v1_DeploymentConfigList(in *deploy_api.DeploymentConfigList, out *DeploymentConfigList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentConfig, len(*in))
		for i := range *in {
			if err := Convert_api_DeploymentConfig_To_v1_DeploymentConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_DeploymentConfigList_To_v1_DeploymentConfigList(in *deploy_api.DeploymentConfigList, out *DeploymentConfigList, s conversion.Scope) error {
	return autoConvert_api_DeploymentConfigList_To_v1_DeploymentConfigList(in, out, s)
}

func autoConvert_v1_DeploymentConfigRollback_To_api_DeploymentConfigRollback(in *DeploymentConfigRollback, out *deploy_api.DeploymentConfigRollback, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Name = in.Name
	out.UpdatedAnnotations = in.UpdatedAnnotations
	if err := Convert_v1_DeploymentConfigRollbackSpec_To_api_DeploymentConfigRollbackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_DeploymentConfigRollback_To_api_DeploymentConfigRollback(in *DeploymentConfigRollback, out *deploy_api.DeploymentConfigRollback, s conversion.Scope) error {
	return autoConvert_v1_DeploymentConfigRollback_To_api_DeploymentConfigRollback(in, out, s)
}

func autoConvert_api_DeploymentConfigRollback_To_v1_DeploymentConfigRollback(in *deploy_api.DeploymentConfigRollback, out *DeploymentConfigRollback, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Name = in.Name
	out.UpdatedAnnotations = in.UpdatedAnnotations
	if err := Convert_api_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_DeploymentConfigRollback_To_v1_DeploymentConfigRollback(in *deploy_api.DeploymentConfigRollback, out *DeploymentConfigRollback, s conversion.Scope) error {
	return autoConvert_api_DeploymentConfigRollback_To_v1_DeploymentConfigRollback(in, out, s)
}

func autoConvert_v1_DeploymentConfigRollbackSpec_To_api_DeploymentConfigRollbackSpec(in *DeploymentConfigRollbackSpec, out *deploy_api.DeploymentConfigRollbackSpec, s conversion.Scope) error {
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.Revision = in.Revision
	out.IncludeTriggers = in.IncludeTriggers
	out.IncludeTemplate = in.IncludeTemplate
	out.IncludeReplicationMeta = in.IncludeReplicationMeta
	out.IncludeStrategy = in.IncludeStrategy
	return nil
}

func Convert_v1_DeploymentConfigRollbackSpec_To_api_DeploymentConfigRollbackSpec(in *DeploymentConfigRollbackSpec, out *deploy_api.DeploymentConfigRollbackSpec, s conversion.Scope) error {
	return autoConvert_v1_DeploymentConfigRollbackSpec_To_api_DeploymentConfigRollbackSpec(in, out, s)
}

func autoConvert_api_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(in *deploy_api.DeploymentConfigRollbackSpec, out *DeploymentConfigRollbackSpec, s conversion.Scope) error {
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.Revision = in.Revision
	out.IncludeTriggers = in.IncludeTriggers
	out.IncludeTemplate = in.IncludeTemplate
	out.IncludeReplicationMeta = in.IncludeReplicationMeta
	out.IncludeStrategy = in.IncludeStrategy
	return nil
}

func Convert_api_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(in *deploy_api.DeploymentConfigRollbackSpec, out *DeploymentConfigRollbackSpec, s conversion.Scope) error {
	return autoConvert_api_DeploymentConfigRollbackSpec_To_v1_DeploymentConfigRollbackSpec(in, out, s)
}

func autoConvert_v1_DeploymentConfigSpec_To_api_DeploymentConfigSpec(in *DeploymentConfigSpec, out *deploy_api.DeploymentConfigSpec, s conversion.Scope) error {
	SetDefaults_DeploymentConfigSpec(in)
	if err := Convert_v1_DeploymentStrategy_To_api_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]deploy_api.DeploymentTriggerPolicy, len(*in))
		for i := range *in {
			if err := Convert_v1_DeploymentTriggerPolicy_To_api_DeploymentTriggerPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.Replicas = in.Replicas
	out.Test = in.Test
	out.Paused = in.Paused
	out.Selector = in.Selector
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(api.PodTemplateSpec)
		if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Template = nil
	}
	return nil
}

func Convert_v1_DeploymentConfigSpec_To_api_DeploymentConfigSpec(in *DeploymentConfigSpec, out *deploy_api.DeploymentConfigSpec, s conversion.Scope) error {
	return autoConvert_v1_DeploymentConfigSpec_To_api_DeploymentConfigSpec(in, out, s)
}

func autoConvert_api_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(in *deploy_api.DeploymentConfigSpec, out *DeploymentConfigSpec, s conversion.Scope) error {
	if err := Convert_api_DeploymentStrategy_To_v1_DeploymentStrategy(&in.Strategy, &out.Strategy, s); err != nil {
		return err
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]DeploymentTriggerPolicy, len(*in))
		for i := range *in {
			if err := Convert_api_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.Replicas = in.Replicas
	out.Test = in.Test
	out.Paused = in.Paused
	out.Selector = in.Selector
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(api_v1.PodTemplateSpec)
		if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Template = nil
	}
	return nil
}

func Convert_api_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(in *deploy_api.DeploymentConfigSpec, out *DeploymentConfigSpec, s conversion.Scope) error {
	return autoConvert_api_DeploymentConfigSpec_To_v1_DeploymentConfigSpec(in, out, s)
}

func autoConvert_v1_DeploymentConfigStatus_To_api_DeploymentConfigStatus(in *DeploymentConfigStatus, out *deploy_api.DeploymentConfigStatus, s conversion.Scope) error {
	out.LatestVersion = in.LatestVersion
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(deploy_api.DeploymentDetails)
		if err := Convert_v1_DeploymentDetails_To_api_DeploymentDetails(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Details = nil
	}
	return nil
}

func Convert_v1_DeploymentConfigStatus_To_api_DeploymentConfigStatus(in *DeploymentConfigStatus, out *deploy_api.DeploymentConfigStatus, s conversion.Scope) error {
	return autoConvert_v1_DeploymentConfigStatus_To_api_DeploymentConfigStatus(in, out, s)
}

func autoConvert_api_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(in *deploy_api.DeploymentConfigStatus, out *DeploymentConfigStatus, s conversion.Scope) error {
	out.LatestVersion = in.LatestVersion
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(DeploymentDetails)
		if err := Convert_api_DeploymentDetails_To_v1_DeploymentDetails(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Details = nil
	}
	return nil
}

func Convert_api_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(in *deploy_api.DeploymentConfigStatus, out *DeploymentConfigStatus, s conversion.Scope) error {
	return autoConvert_api_DeploymentConfigStatus_To_v1_DeploymentConfigStatus(in, out, s)
}

func autoConvert_v1_DeploymentDetails_To_api_DeploymentDetails(in *DeploymentDetails, out *deploy_api.DeploymentDetails, s conversion.Scope) error {
	out.Message = in.Message
	if in.Causes != nil {
		in, out := &in.Causes, &out.Causes
		*out = make([]deploy_api.DeploymentCause, len(*in))
		for i := range *in {
			if err := Convert_v1_DeploymentCause_To_api_DeploymentCause(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Causes = nil
	}
	return nil
}

func Convert_v1_DeploymentDetails_To_api_DeploymentDetails(in *DeploymentDetails, out *deploy_api.DeploymentDetails, s conversion.Scope) error {
	return autoConvert_v1_DeploymentDetails_To_api_DeploymentDetails(in, out, s)
}

func autoConvert_api_DeploymentDetails_To_v1_DeploymentDetails(in *deploy_api.DeploymentDetails, out *DeploymentDetails, s conversion.Scope) error {
	out.Message = in.Message
	if in.Causes != nil {
		in, out := &in.Causes, &out.Causes
		*out = make([]DeploymentCause, len(*in))
		for i := range *in {
			if err := Convert_api_DeploymentCause_To_v1_DeploymentCause(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Causes = nil
	}
	return nil
}

func Convert_api_DeploymentDetails_To_v1_DeploymentDetails(in *deploy_api.DeploymentDetails, out *DeploymentDetails, s conversion.Scope) error {
	return autoConvert_api_DeploymentDetails_To_v1_DeploymentDetails(in, out, s)
}

func autoConvert_v1_DeploymentLog_To_api_DeploymentLog(in *DeploymentLog, out *deploy_api.DeploymentLog, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_DeploymentLog_To_api_DeploymentLog(in *DeploymentLog, out *deploy_api.DeploymentLog, s conversion.Scope) error {
	return autoConvert_v1_DeploymentLog_To_api_DeploymentLog(in, out, s)
}

func autoConvert_api_DeploymentLog_To_v1_DeploymentLog(in *deploy_api.DeploymentLog, out *DeploymentLog, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_DeploymentLog_To_v1_DeploymentLog(in *deploy_api.DeploymentLog, out *DeploymentLog, s conversion.Scope) error {
	return autoConvert_api_DeploymentLog_To_v1_DeploymentLog(in, out, s)
}

func autoConvert_v1_DeploymentLogOptions_To_api_DeploymentLogOptions(in *DeploymentLogOptions, out *deploy_api.DeploymentLogOptions, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	out.SinceSeconds = in.SinceSeconds
	out.SinceTime = in.SinceTime
	out.Timestamps = in.Timestamps
	out.TailLines = in.TailLines
	out.LimitBytes = in.LimitBytes
	out.NoWait = in.NoWait
	out.Version = in.Version
	return nil
}

func Convert_v1_DeploymentLogOptions_To_api_DeploymentLogOptions(in *DeploymentLogOptions, out *deploy_api.DeploymentLogOptions, s conversion.Scope) error {
	return autoConvert_v1_DeploymentLogOptions_To_api_DeploymentLogOptions(in, out, s)
}

func autoConvert_api_DeploymentLogOptions_To_v1_DeploymentLogOptions(in *deploy_api.DeploymentLogOptions, out *DeploymentLogOptions, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	out.SinceSeconds = in.SinceSeconds
	out.SinceTime = in.SinceTime
	out.Timestamps = in.Timestamps
	out.TailLines = in.TailLines
	out.LimitBytes = in.LimitBytes
	out.NoWait = in.NoWait
	out.Version = in.Version
	return nil
}

func Convert_api_DeploymentLogOptions_To_v1_DeploymentLogOptions(in *deploy_api.DeploymentLogOptions, out *DeploymentLogOptions, s conversion.Scope) error {
	return autoConvert_api_DeploymentLogOptions_To_v1_DeploymentLogOptions(in, out, s)
}

func autoConvert_v1_DeploymentStrategy_To_api_DeploymentStrategy(in *DeploymentStrategy, out *deploy_api.DeploymentStrategy, s conversion.Scope) error {
	SetDefaults_DeploymentStrategy(in)
	out.Type = deploy_api.DeploymentStrategyType(in.Type)
	if in.CustomParams != nil {
		in, out := &in.CustomParams, &out.CustomParams
		*out = new(deploy_api.CustomDeploymentStrategyParams)
		if err := Convert_v1_CustomDeploymentStrategyParams_To_api_CustomDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CustomParams = nil
	}
	if in.RecreateParams != nil {
		in, out := &in.RecreateParams, &out.RecreateParams
		*out = new(deploy_api.RecreateDeploymentStrategyParams)
		if err := Convert_v1_RecreateDeploymentStrategyParams_To_api_RecreateDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RecreateParams = nil
	}
	if in.RollingParams != nil {
		in, out := &in.RollingParams, &out.RollingParams
		*out = new(deploy_api.RollingDeploymentStrategyParams)
		if err := Convert_v1_RollingDeploymentStrategyParams_To_api_RollingDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RollingParams = nil
	}
	if err := api_v1.Convert_v1_ResourceRequirements_To_api_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	return nil
}

func Convert_v1_DeploymentStrategy_To_api_DeploymentStrategy(in *DeploymentStrategy, out *deploy_api.DeploymentStrategy, s conversion.Scope) error {
	return autoConvert_v1_DeploymentStrategy_To_api_DeploymentStrategy(in, out, s)
}

func autoConvert_api_DeploymentStrategy_To_v1_DeploymentStrategy(in *deploy_api.DeploymentStrategy, out *DeploymentStrategy, s conversion.Scope) error {
	out.Type = DeploymentStrategyType(in.Type)
	if in.RecreateParams != nil {
		in, out := &in.RecreateParams, &out.RecreateParams
		*out = new(RecreateDeploymentStrategyParams)
		if err := Convert_api_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RecreateParams = nil
	}
	if in.RollingParams != nil {
		in, out := &in.RollingParams, &out.RollingParams
		*out = new(RollingDeploymentStrategyParams)
		if err := Convert_api_RollingDeploymentStrategyParams_To_v1_RollingDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RollingParams = nil
	}
	if in.CustomParams != nil {
		in, out := &in.CustomParams, &out.CustomParams
		*out = new(CustomDeploymentStrategyParams)
		if err := Convert_api_CustomDeploymentStrategyParams_To_v1_CustomDeploymentStrategyParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.CustomParams = nil
	}
	if err := api_v1.Convert_api_ResourceRequirements_To_v1_ResourceRequirements(&in.Resources, &out.Resources, s); err != nil {
		return err
	}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	return nil
}

func Convert_api_DeploymentStrategy_To_v1_DeploymentStrategy(in *deploy_api.DeploymentStrategy, out *DeploymentStrategy, s conversion.Scope) error {
	return autoConvert_api_DeploymentStrategy_To_v1_DeploymentStrategy(in, out, s)
}

func autoConvert_v1_DeploymentTriggerImageChangeParams_To_api_DeploymentTriggerImageChangeParams(in *DeploymentTriggerImageChangeParams, out *deploy_api.DeploymentTriggerImageChangeParams, s conversion.Scope) error {
	out.Automatic = in.Automatic
	out.ContainerNames = in.ContainerNames
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.LastTriggeredImage = in.LastTriggeredImage
	return nil
}

func autoConvert_api_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams(in *deploy_api.DeploymentTriggerImageChangeParams, out *DeploymentTriggerImageChangeParams, s conversion.Scope) error {
	out.Automatic = in.Automatic
	out.ContainerNames = in.ContainerNames
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.From, &out.From, s); err != nil {
		return err
	}
	out.LastTriggeredImage = in.LastTriggeredImage
	return nil
}

func autoConvert_v1_DeploymentTriggerPolicy_To_api_DeploymentTriggerPolicy(in *DeploymentTriggerPolicy, out *deploy_api.DeploymentTriggerPolicy, s conversion.Scope) error {
	out.Type = deploy_api.DeploymentTriggerType(in.Type)
	if in.ImageChangeParams != nil {
		in, out := &in.ImageChangeParams, &out.ImageChangeParams
		*out = new(deploy_api.DeploymentTriggerImageChangeParams)
		if err := Convert_v1_DeploymentTriggerImageChangeParams_To_api_DeploymentTriggerImageChangeParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageChangeParams = nil
	}
	return nil
}

func Convert_v1_DeploymentTriggerPolicy_To_api_DeploymentTriggerPolicy(in *DeploymentTriggerPolicy, out *deploy_api.DeploymentTriggerPolicy, s conversion.Scope) error {
	return autoConvert_v1_DeploymentTriggerPolicy_To_api_DeploymentTriggerPolicy(in, out, s)
}

func autoConvert_api_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(in *deploy_api.DeploymentTriggerPolicy, out *DeploymentTriggerPolicy, s conversion.Scope) error {
	out.Type = DeploymentTriggerType(in.Type)
	if in.ImageChangeParams != nil {
		in, out := &in.ImageChangeParams, &out.ImageChangeParams
		*out = new(DeploymentTriggerImageChangeParams)
		if err := Convert_api_DeploymentTriggerImageChangeParams_To_v1_DeploymentTriggerImageChangeParams(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ImageChangeParams = nil
	}
	return nil
}

func Convert_api_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(in *deploy_api.DeploymentTriggerPolicy, out *DeploymentTriggerPolicy, s conversion.Scope) error {
	return autoConvert_api_DeploymentTriggerPolicy_To_v1_DeploymentTriggerPolicy(in, out, s)
}

func autoConvert_v1_ExecNewPodHook_To_api_ExecNewPodHook(in *ExecNewPodHook, out *deploy_api.ExecNewPodHook, s conversion.Scope) error {
	out.Command = in.Command
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]api.EnvVar, len(*in))
		for i := range *in {
			if err := api_v1.Convert_v1_EnvVar_To_api_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ContainerName = in.ContainerName
	out.Volumes = in.Volumes
	return nil
}

func Convert_v1_ExecNewPodHook_To_api_ExecNewPodHook(in *ExecNewPodHook, out *deploy_api.ExecNewPodHook, s conversion.Scope) error {
	return autoConvert_v1_ExecNewPodHook_To_api_ExecNewPodHook(in, out, s)
}

func autoConvert_api_ExecNewPodHook_To_v1_ExecNewPodHook(in *deploy_api.ExecNewPodHook, out *ExecNewPodHook, s conversion.Scope) error {
	out.Command = in.Command
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]api_v1.EnvVar, len(*in))
		for i := range *in {
			if err := api_v1.Convert_api_EnvVar_To_v1_EnvVar(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ContainerName = in.ContainerName
	out.Volumes = in.Volumes
	return nil
}

func Convert_api_ExecNewPodHook_To_v1_ExecNewPodHook(in *deploy_api.ExecNewPodHook, out *ExecNewPodHook, s conversion.Scope) error {
	return autoConvert_api_ExecNewPodHook_To_v1_ExecNewPodHook(in, out, s)
}

func autoConvert_v1_LifecycleHook_To_api_LifecycleHook(in *LifecycleHook, out *deploy_api.LifecycleHook, s conversion.Scope) error {
	out.FailurePolicy = deploy_api.LifecycleHookFailurePolicy(in.FailurePolicy)
	if in.ExecNewPod != nil {
		in, out := &in.ExecNewPod, &out.ExecNewPod
		*out = new(deploy_api.ExecNewPodHook)
		if err := Convert_v1_ExecNewPodHook_To_api_ExecNewPodHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ExecNewPod = nil
	}
	if in.TagImages != nil {
		in, out := &in.TagImages, &out.TagImages
		*out = make([]deploy_api.TagImageHook, len(*in))
		for i := range *in {
			if err := Convert_v1_TagImageHook_To_api_TagImageHook(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.TagImages = nil
	}
	return nil
}

func Convert_v1_LifecycleHook_To_api_LifecycleHook(in *LifecycleHook, out *deploy_api.LifecycleHook, s conversion.Scope) error {
	return autoConvert_v1_LifecycleHook_To_api_LifecycleHook(in, out, s)
}

func autoConvert_api_LifecycleHook_To_v1_LifecycleHook(in *deploy_api.LifecycleHook, out *LifecycleHook, s conversion.Scope) error {
	out.FailurePolicy = LifecycleHookFailurePolicy(in.FailurePolicy)
	if in.ExecNewPod != nil {
		in, out := &in.ExecNewPod, &out.ExecNewPod
		*out = new(ExecNewPodHook)
		if err := Convert_api_ExecNewPodHook_To_v1_ExecNewPodHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ExecNewPod = nil
	}
	if in.TagImages != nil {
		in, out := &in.TagImages, &out.TagImages
		*out = make([]TagImageHook, len(*in))
		for i := range *in {
			if err := Convert_api_TagImageHook_To_v1_TagImageHook(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.TagImages = nil
	}
	return nil
}

func Convert_api_LifecycleHook_To_v1_LifecycleHook(in *deploy_api.LifecycleHook, out *LifecycleHook, s conversion.Scope) error {
	return autoConvert_api_LifecycleHook_To_v1_LifecycleHook(in, out, s)
}

func autoConvert_v1_RecreateDeploymentStrategyParams_To_api_RecreateDeploymentStrategyParams(in *RecreateDeploymentStrategyParams, out *deploy_api.RecreateDeploymentStrategyParams, s conversion.Scope) error {
	SetDefaults_RecreateDeploymentStrategyParams(in)
	out.TimeoutSeconds = in.TimeoutSeconds
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(deploy_api.LifecycleHook)
		if err := Convert_v1_LifecycleHook_To_api_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Mid != nil {
		in, out := &in.Mid, &out.Mid
		*out = new(deploy_api.LifecycleHook)
		if err := Convert_v1_LifecycleHook_To_api_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Mid = nil
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(deploy_api.LifecycleHook)
		if err := Convert_v1_LifecycleHook_To_api_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}

func Convert_v1_RecreateDeploymentStrategyParams_To_api_RecreateDeploymentStrategyParams(in *RecreateDeploymentStrategyParams, out *deploy_api.RecreateDeploymentStrategyParams, s conversion.Scope) error {
	return autoConvert_v1_RecreateDeploymentStrategyParams_To_api_RecreateDeploymentStrategyParams(in, out, s)
}

func autoConvert_api_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(in *deploy_api.RecreateDeploymentStrategyParams, out *RecreateDeploymentStrategyParams, s conversion.Scope) error {
	out.TimeoutSeconds = in.TimeoutSeconds
	if in.Pre != nil {
		in, out := &in.Pre, &out.Pre
		*out = new(LifecycleHook)
		if err := Convert_api_LifecycleHook_To_v1_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Mid != nil {
		in, out := &in.Mid, &out.Mid
		*out = new(LifecycleHook)
		if err := Convert_api_LifecycleHook_To_v1_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Mid = nil
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(LifecycleHook)
		if err := Convert_api_LifecycleHook_To_v1_LifecycleHook(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}

func Convert_api_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(in *deploy_api.RecreateDeploymentStrategyParams, out *RecreateDeploymentStrategyParams, s conversion.Scope) error {
	return autoConvert_api_RecreateDeploymentStrategyParams_To_v1_RecreateDeploymentStrategyParams(in, out, s)
}

func autoConvert_v1_TagImageHook_To_api_TagImageHook(in *TagImageHook, out *deploy_api.TagImageHook, s conversion.Scope) error {
	out.ContainerName = in.ContainerName
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_TagImageHook_To_api_TagImageHook(in *TagImageHook, out *deploy_api.TagImageHook, s conversion.Scope) error {
	return autoConvert_v1_TagImageHook_To_api_TagImageHook(in, out, s)
}

func autoConvert_api_TagImageHook_To_v1_TagImageHook(in *deploy_api.TagImageHook, out *TagImageHook, s conversion.Scope) error {
	out.ContainerName = in.ContainerName
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.To, &out.To, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_TagImageHook_To_v1_TagImageHook(in *deploy_api.TagImageHook, out *TagImageHook, s conversion.Scope) error {
	return autoConvert_api_TagImageHook_To_v1_TagImageHook(in, out, s)
}
