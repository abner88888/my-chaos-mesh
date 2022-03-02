/**
 * This file was auto-generated by @ui/openapi.
 * Do not make direct changes to the file.
 */

export const actions = ['latency', 'fault', 'attrOverride', 'mistake'],
  data = [
    {
      field: 'ref',
      label: 'attr',
      children: [
        {
          field: 'ref',
          label: 'atime',
          children: [
            {
              field: 'number',
              label: 'nsec',
              value: 0,
              helperText: '',
            },
            {
              field: 'number',
              label: 'sec',
              value: 0,
              helperText: '',
            },
          ],
        },
        {
          field: 'number',
          label: 'blocks',
          value: 0,
          helperText: 'Optional.',
        },
        {
          field: 'ref',
          label: 'ctime',
          children: [
            {
              field: 'number',
              label: 'nsec',
              value: 0,
              helperText: '',
            },
            {
              field: 'number',
              label: 'sec',
              value: 0,
              helperText: '',
            },
          ],
        },
        {
          field: 'number',
          label: 'gid',
          value: 0,
          helperText: 'Optional.',
        },
        {
          field: 'number',
          label: 'ino',
          value: 0,
          helperText: 'Optional.',
        },
        {
          field: 'text',
          label: 'kind',
          value: '',
          helperText: 'Optional.',
        },
        {
          field: 'ref',
          label: 'mtime',
          children: [
            {
              field: 'number',
              label: 'nsec',
              value: 0,
              helperText: '',
            },
            {
              field: 'number',
              label: 'sec',
              value: 0,
              helperText: '',
            },
          ],
        },
        {
          field: 'number',
          label: 'nlink',
          value: 0,
          helperText: 'Optional.',
        },
        {
          field: 'number',
          label: 'perm',
          value: 0,
          helperText: 'Optional.',
        },
        {
          field: 'number',
          label: 'rdev',
          value: 0,
          helperText: 'Optional.',
        },
        {
          field: 'number',
          label: 'size',
          value: 0,
          helperText: 'Optional.',
        },
        {
          field: 'number',
          label: 'uid',
          value: 0,
          helperText: 'Optional.',
        },
      ],
      when: "action=='attrOverride'",
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
      label: 'delay',
      value: '',
      helperText:
        'Optional. Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as \\"300ms\\". Valid time units are \\"ns\\", \\"us\\" (or \\"\u00B5s\\"), \\"ms\\", \\"s\\", \\"m\\", \\"h\\".',
      when: "action=='latency'",
    },
    {
      field: 'number',
      label: 'errno',
      value: 0,
      helperText:
        'Optional. Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html',
      when: "action=='fault'",
    },
    {
      field: 'label',
      label: 'methods',
      value: [],
      helperText: 'Optional. Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.',
    },
    {
      field: 'ref',
      label: 'mistake',
      children: [
        {
          field: 'text',
          label: 'filling',
          value: '',
          helperText: 'Optional. Filling determines what is filled in the miskate data.',
        },
        {
          field: 'number',
          label: 'maxLength',
          value: 0,
          helperText: 'Optional. Max length of each wrong data segment in bytes',
        },
        {
          field: 'number',
          label: 'maxOccurrences',
          value: 0,
          helperText: 'Optional. There will be [1, MaxOccurrences] segments of wrong data.',
        },
      ],
      when: "action=='mistake'",
    },
    {
      field: 'text',
      label: 'path',
      value: '',
      helperText: 'Optional. Path defines the path of files for injecting I/O chaos action.',
    },
    {
      field: 'number',
      label: 'percent',
      value: 0,
      helperText:
        'Optional. Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.',
    },
    {
      field: 'text',
      label: 'volumePath',
      value: '',
      helperText: 'VolumePath represents the mount path of injected volume',
    },
  ]
