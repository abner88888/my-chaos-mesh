/**
 * This file was auto-generated by @ui/openapi.
 * Do not make direct changes to the file.
 */

const shared = [
  {
    field: 'text',
    label: 'class',
    value: '',
    helperText: 'Optional.  Java class',
  },
  {
    field: 'label',
    label: 'containerNames',
    value: [],
    helperText:
      'Optional. ContainerNames indicates list of the name of affected container. If not set, the first container will be injected',
  },
  {
    field: 'text',
    label: 'cpuCount',
    value: 0,
    helperText: 'Optional.  the CPU core number need to use, only set it when action is stress',
  },
  {
    field: 'text',
    label: 'exception',
    value: '',
    helperText: 'Optional.  the exception which needs to throw for action `exception`',
  },
  {
    field: 'text',
    label: 'latency',
    value: 0,
    helperText: "Optional.  the latency duration for action \\'latency\\', unit ms",
  },
  {
    field: 'text',
    label: 'memType',
    value: '',
    helperText:
      "Optional.  the memory type need to locate, only set it when action is stress, the value can be \\'stack\\' or \\'heap\\'",
  },
  {
    field: 'text',
    label: 'method',
    value: '',
    helperText: 'Optional.  the method in Java class',
  },
  {
    field: 'text',
    label: 'name',
    value: '',
    helperText: "Optional.  byteman rule name, should be unique, and will use JVMChaos\\' name if not set",
  },
  {
    field: 'text',
    label: 'port',
    value: 0,
    helperText: 'Optional.  the port of agent server, default 9277',
  },
  {
    field: 'text',
    label: 'ruleData',
    value: '',
    helperText: 'Optional.',
  },
]

export default {
  latency: shared,
  return: shared,
  exception: shared,
  stress: shared,
  gc: shared,
  ruleData: shared,
}
