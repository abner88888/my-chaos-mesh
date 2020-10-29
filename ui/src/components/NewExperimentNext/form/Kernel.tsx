import { Box, Button, Divider, InputAdornment, MenuItem, Typography } from '@material-ui/core'
import { Form, Formik } from 'formik'
import { LabelField, SelectField, TextField } from 'components/FormField'
import React, { useEffect, useState } from 'react'

import AddIcon from '@material-ui/icons/Add'
import Paper from 'components-mui/Paper'
import PaperTop from 'components/PaperTop'
import RemoveIcon from '@material-ui/icons/Remove'
import { RootState } from 'store'
import targetData from '../data/target'
import { useSelector } from 'react-redux'

interface KernelProps {
  onSubmit: (values: Record<string, any>) => void
}

const Kernel: React.FC<KernelProps> = ({ onSubmit }) => {
  const { target } = useSelector((state: RootState) => state.experiments)

  const initialValues = targetData.KernelChaos.spec!

  const [init, setInit] = useState(initialValues)

  useEffect(() => {
    setInit({
      fail_kern_request: {
        ...initialValues.fail_kern_request,
        ...target['kernel_chaos']?.fail_kern_request,
      },
    })
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [target])

  return (
    <Formik enableReinitialize initialValues={init} onSubmit={onSubmit}>
      {({ values, setFieldValue }) => {
        const callchain = (values.fail_kern_request as any).callchain

        const addFrame = () =>
          setFieldValue(
            'fail_kern_request.callchain',
            callchain.concat([
              {
                funcname: '',
                parameters: '',
                predicate: '',
              },
            ])
          )

        const removeFrame = (index: number) => () => {
          setFieldValue(
            'fail_kern_request.callchain',
            callchain.filter((_: any, i: number) => index !== i)
          )
        }

        return (
          <Form>
            <Paper>
              <PaperTop title="Callchain">
                <Button variant="outlined" color="primary" startIcon={<AddIcon />} onClick={addFrame}>
                  Add
                </Button>
              </PaperTop>
              <Box>
                {callchain.map((_: any, i: number) => (
                  <Box key={'frame' + i} p={3}>
                    <Box display="flex" justifyContent="space-between" alignItems="center" mb={3}>
                      <Typography variant="body2" gutterBottom>
                        Frame {i + 1}
                      </Typography>
                      <Button
                        variant="outlined"
                        size="small"
                        color="secondary"
                        startIcon={<RemoveIcon />}
                        onClick={removeFrame(i)}
                      >
                        Remove
                      </Button>
                    </Box>
                    <TextField
                      id={`fail_kern_request.callchain[${i}].funcname`}
                      name={`fail_kern_request.callchain[${i}].funcname`}
                      label="funcname"
                    />
                    <TextField
                      id={`fail_kern_request.callchain[${i}].parameters`}
                      name={`fail_kern_request.callchain[${i}].parameters`}
                      label="parameters"
                    />
                    <TextField
                      id={`fail_kern_request.callchain[${i}].predicate`}
                      name={`fail_kern_request.callchain[${i}].predicate`}
                      label="predicate"
                    />
                  </Box>
                ))}
              </Box>
            </Paper>
            <Box my={6}>
              <Divider />
            </Box>
            <SelectField
              id="fail_kern_request.failtype"
              name="fail_kern_request.failtype"
              label="Failtype"
              helperText="What to fail, can be set to 0 / 1 / 2"
            >
              {[0, 1, 2].map((option) => (
                <MenuItem key={option} value={option}>
                  {option}
                </MenuItem>
              ))}
            </SelectField>
            <LabelField
              id="fail_kern_request.headers"
              name="fail_kern_request.headers"
              label="Headers"
              helperText="Type string and end with a space to generate the appropriate kernel headers"
            />
            <TextField
              type="number"
              id="fail_kern_request.probability"
              name="fail_kern_request.probability"
              helperText="The fails with probability"
              InputProps={{
                endAdornment: <InputAdornment position="end">%</InputAdornment>,
              }}
            />
            <TextField
              type="number"
              id="fail_kern_request.times"
              name="fail_kern_request.times"
              helperText="The max times of failures"
            />
          </Form>
        )
      }}
    </Formik>
  )
}

export default Kernel
