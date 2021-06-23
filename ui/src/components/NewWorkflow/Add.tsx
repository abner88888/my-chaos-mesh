import { Box, Grid, MenuItem, StepLabel, Typography } from '@material-ui/core'
import { Form, Formik, FormikHelpers } from 'formik'
import MultiNode, { MultiNodeHandles } from './MultiNode'
import NewExperimentNext, { NewExperimentHandles } from 'components/NewExperimentNext'
import { SelectField, Submit, TextField } from 'components/FormField'
import { TemplateExperiment, setTemplate } from 'slices/workflows'
import { resetNewExperiment, setExternalExperiment } from 'slices/experiments'
import { useRef, useState } from 'react'
import { validateDeadline, validateName } from 'lib/formikhelpers'

import AddCircleIcon from '@material-ui/icons/AddCircle'
import Paper from 'components-mui/Paper'
import PaperTop from 'components-mui/PaperTop'
import Space from 'components-mui/Space'
import Suspend from './Suspend'
import T from 'components/T'
import _snakecase from 'lodash.snakecase'
import { makeStyles } from '@material-ui/styles'
import { setAlert } from 'slices/globalStatus'
import { useIntl } from 'react-intl'
import { useStoreDispatch } from 'store'

const useStyles = makeStyles((theme) => ({
  field: {
    width: 180,
    marginTop: 0,
    [theme.breakpoints.up('sm')]: {
      marginBottom: 0,
    },
    '& .MuiInputBase-input': {
      padding: 8,
    },
    '& .MuiInputLabel-root, fieldset': {
      fontSize: theme.typography.body2.fontSize,
      lineHeight: 0.875,
    },
  },
}))

const types = ['single', 'serial', 'parallel', 'suspend']

const Add = () => {
  const classes = useStyles()
  const intl = useIntl()

  const dispatch = useStoreDispatch()

  const [showNum, setShowNum] = useState(false)
  const [num, setNum] = useState(1)
  const [otherTypes, setOtherTypes] = useState(false)
  const [experiments, setExperiments] = useState<TemplateExperiment[]>([])
  const formRef = useRef<any>()
  const newERef = useRef<NewExperimentHandles>(null)
  const multiNodeRef = useRef<MultiNodeHandles>(null)

  const resetNoSingle = () => {
    setShowNum(false)
    setExperiments([])
    multiNodeRef.current?.setCurrent(0)
  }

  const onValidate = ({ type, num: newNum }: { type: string; num: number }) => {
    if (type !== 'suspend') {
      setOtherTypes(false)
    }

    const prevType = formRef.current.values.type

    if (prevType !== 'single' && type === 'single') {
      resetNoSingle()

      return
    }

    if (type === 'serial' || type === 'parallel') {
      setShowNum(true)

      // Delete extra experiments
      if (num > newNum) {
        setExperiments(experiments.slice(0, -1))
      }

      setNum(newNum)

      return
    }

    if (type === 'suspend') {
      if (prevType === 'serial' || prevType === 'parallel') {
        resetNoSingle()
      }

      setOtherTypes(true)
    }
  }

  const onSubmit = (experiment: any) => {
    const type = formRef.current.values.type

    if (type === 'single') {
      dispatch(
        setTemplate({
          type,
          name: experiment.basic.name,
          experiments: [experiment],
        })
      )
    } else {
      const current = multiNodeRef.current!.current

      multiNodeRef.current!.setCurrent(current + 1)

      // Edit the node that has been submitted before
      if (current < experiments.length) {
        const es = experiments

        es[current] = experiment

        setExperiments(es)

        dispatch(
          setAlert({
            type: 'success',
            message: T('confirm.success.update', intl),
          })
        )
      } else {
        setExperiments([...experiments, experiment])
      }
    }

    newERef.current?.setPanel('existing')
    dispatch(resetNewExperiment())
  }

  const submitNoSingleNode = (_: any, { resetForm }: FormikHelpers<any>) => {
    const { type, name, deadline } = formRef.current.values

    dispatch(
      setTemplate({
        type,
        name,
        deadline,
        experiments,
      })
    )

    resetNoSingle()
    resetForm()
  }

  const setCurrentCallback = (index: number) => {
    if (index > experiments.length) {
      dispatch(
        setAlert({
          type: 'warning',
          message: T('newW.messages.m1', intl),
        })
      )

      return false
    }

    if (index < experiments.length) {
      const e = experiments[index]

      const kind = e.target.kind

      dispatch(
        setExternalExperiment({
          kindAction: [kind, e.target[_snakecase(kind)].action ?? ''],
          target: e.target,
          basic: e.basic,
        })
      )

      newERef.current?.setPanel('initial')
    }

    return true
  }

  return (
    <>
      <Formik
        innerRef={formRef}
        initialValues={{ type: 'single', num: 2, name: '', deadline: '' }}
        onSubmit={submitNoSingleNode}
        validate={onValidate}
        validateOnBlur={false}
      >
        {({ values, errors, touched }) => (
          <Form>
            <StepLabel icon={<AddCircleIcon color="primary" />}>
              <Space direction="row">
                <SelectField className={classes.field} name="type" label={T('newW.node.choose')}>
                  {types.map((d) => (
                    <MenuItem key={d} value={d}>
                      <Typography variant="body2">{T(`newW.node.${d}`)}</Typography>
                    </MenuItem>
                  ))}
                </SelectField>
                {showNum && (
                  <TextField
                    className={classes.field}
                    type="number"
                    name="num"
                    label={T('newW.node.number')}
                    inputProps={{ min: 1 }}
                  />
                )}
              </Space>
            </StepLabel>

            {showNum && (
              <Box my={3} ml={8}>
                <Paper>
                  <PaperTop title={T(`newW.${values.type}Title`)} boxProps={{ mb: 3 }} />
                  <Grid container spacing={3}>
                    <Grid item xs={6}>
                      <TextField
                        name="name"
                        label={T('common.name')}
                        validate={validateName(T('newW.nameValidation', intl))}
                        helperText={errors.name && touched.name ? errors.name : T('newW.node.nameHelper')}
                        error={errors.name && touched.name ? true : false}
                      />
                    </Grid>
                    <Grid item xs={6}>
                      <TextField
                        name="deadline"
                        label={T('newW.node.deadline')}
                        validate={validateDeadline(T('newW.node.deadlineValidation', intl))}
                        helperText={
                          errors.deadline && touched.deadline ? errors.deadline : T('newW.node.deadlineHelper')
                        }
                        error={errors.deadline && touched.deadline ? true : false}
                      />
                    </Grid>
                  </Grid>
                  <Box display="flex" justifyContent="space-between" alignItems="center" mt={6}>
                    <MultiNode ref={multiNodeRef} count={num} setCurrentCallback={setCurrentCallback} />
                    <Submit mt={0} disabled={experiments.length !== num} />
                  </Box>
                </Paper>
              </Box>
            )}
          </Form>
        )}
      </Formik>
      <Box ml={8}>
        <Box display={otherTypes ? 'none' : 'initial'}>
          <NewExperimentNext ref={newERef} initPanel="existing" onSubmit={onSubmit} inWorkflow={true} />
        </Box>
        {otherTypes && (
          <Box mt={3}>
            <Suspend />
          </Box>
        )}
      </Box>
    </>
  )
}

export default Add
