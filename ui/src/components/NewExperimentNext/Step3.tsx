import { Box, Typography } from '@material-ui/core'
import { useStoreDispatch, useStoreSelector } from 'store'

import DoneAllIcon from '@material-ui/icons/DoneAll'
import { ExperimentKind } from 'components/NewExperiment/types'
import Paper from 'components-mui/Paper'
import PaperTop from 'components-mui/PaperTop'
import { Submit } from 'components/FormField'
import T from 'components/T'
import api from 'api'
import { parseSubmit } from 'lib/formikhelpers'
import { resetNewExperiment } from 'slices/experiments'
import { setAlert } from 'slices/globalStatus'
import { useHistory } from 'react-router-dom'
import { useIntl } from 'react-intl'

interface Step3Props {
  onSubmit?: (parsedValues: any) => void
  inSchedule?: boolean
}

const Step3: React.FC<Step3Props> = ({ onSubmit, inSchedule }) => {
  const history = useHistory()
  const intl = useIntl()

  const state = useStoreSelector((state) => state)
  const { step1, step2, kindAction, basic, spec } = state.experiments
  const { debugMode } = state.settings
  const dispatch = useStoreDispatch()

  const submitExperiment = () => {
    const parsedValues = parseSubmit(
      kindAction[0] as ExperimentKind,
      {
        ...basic,
        spec: {
          ...basic.spec,
          ...spec,
        },
      },
      { inSchedule }
    )

    if (process.env.NODE_ENV === 'development' || debugMode) {
      console.debug('Debug parsedValues:', parsedValues)
    }

    if (!debugMode) {
      if (onSubmit) {
        onSubmit(parsedValues)
      } else {
        api.experiments
          .newExperiment(parsedValues)
          .then(() => {
            dispatch(
              setAlert({
                type: 'success',
                message: T('confirm.success.create', intl),
              })
            )

            dispatch(resetNewExperiment())

            history.push('/experiments')
          })
          .catch(console.error)
      }
    }
  }

  return (
    <>
      {step1 && step2 && (
        <Paper>
          <PaperTop title={T('common.submit')} boxProps={{ mb: 6 }} />
          <Box textAlign="center">
            <DoneAllIcon fontSize="large" />
            <Typography>{T('newE.complete')}</Typography>
          </Box>
          <Submit onClick={submitExperiment} />
        </Paper>
      )}
    </>
  )
}

export default Step3
