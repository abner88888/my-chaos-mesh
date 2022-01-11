/**
 * This file was auto-generated by @ui/openapi.
 * Do not make direct changes to the file.
 */

const shared = [
  {
    field: 'label',
    label: 'containerNames',
    value: [],
    helperText:
      'Optional. ContainerNames indicates list of the name of affected container. If not set, the first container will be injected',
  },
  {
    field: 'text',
    label: 'gracePeriod',
    value: 0,
    helperText:
      'Optional. GracePeriod is used in pod-kill action. It represents the duration in seconds before the pod should be deleted. Value must be non-negative integer. The default value is zero that indicates delete immediately.  +kubebuilder:validation:Minimum=0',
  },
]

export default {
  'pod-kill': shared,
  'pod-failure': shared,
  'container-kill': shared,
}
