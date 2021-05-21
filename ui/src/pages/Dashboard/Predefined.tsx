import { Box, Button, Card, Modal, Typography } from '@material-ui/core'
import { PreDefinedValue, getDB } from 'lib/idb'
import React, { useEffect, useRef, useState } from 'react'
import { parseSubmit, yamlToExperiment } from 'lib/formikhelpers'
import { useStoreDispatch, useStoreSelector } from 'store'

import { Ace } from 'ace-builds'
import Paper from 'components-mui/Paper'
import PaperTop from 'components-mui/PaperTop'
import Space from 'components-mui/Space'
import T from 'components/T'
import YAML from 'components/YAML'
import api from 'api'
import clsx from 'clsx'
import { iconByKind } from 'lib/byKind'
import loadable from '@loadable/component'
import { makeStyles } from '@material-ui/core/styles'
import { setAlert } from 'slices/globalStatus'
import { useIntl } from 'react-intl'
import yaml from 'js-yaml'

const YAMLEditor = loadable(() => import('components/YAMLEditor'))

const useStyles = makeStyles((theme) => ({
  container: {
    display: 'flex',
    height: 88,
    overflowX: 'scroll',
  },
  card: {
    flex: '0 0 240px',
    cursor: 'pointer',
    '&:hover': {
      background: theme.palette.action.hover,
    },
  },
  addCard: {
    width: 210,
  },
  editorPaperWrapper: {
    position: 'absolute',
    top: '50%',
    left: '50%',
    width: '50vw',
    height: '90vh',
    transform: 'translate(-50%, -50%)',
    [theme.breakpoints.down('sm')]: {
      width: '90vw',
    },
  },
}))

const Predefined = () => {
  const classes = useStyles()

  const intl = useIntl()

  const { theme } = useStoreSelector((state) => state.settings)
  const dispatch = useStoreDispatch()

  const idb = useRef(getDB())

  const [yamlEditor, setYAMLEditor] = useState<Ace.Editor>()
  const [editorOpen, seteditorOpen] = useState(false)
  const [experiment, setExperiment] = useState<PreDefinedValue>()
  const [experiments, setExperiments] = useState<PreDefinedValue[]>([])

  async function getExperiments() {
    setExperiments(await (await idb.current).getAll('predefined'))
  }

  useEffect(() => {
    getExperiments()
  }, [])

  const saveExperiment = async (y: any) => {
    const db = await idb.current

    await db.put('predefined', {
      name: y.metadata.name,
      kind: y.kind,
      yaml: y,
    })

    getExperiments()
  }

  const onModalOpen = (exp: PreDefinedValue) => () => {
    seteditorOpen(true)
    setExperiment(exp)
  }
  const onModalClose = () => seteditorOpen(false)

  const handleApplyExperiment = () => {
    const { basic, target } = yamlToExperiment(yaml.load(yamlEditor!.getValue()))
    const parsedValues = parseSubmit({
      ...basic,
      target,
    })

    if (process.env.NODE_ENV === 'development') {
      console.debug('Debug parsedValues:', parsedValues)
    }

    api.experiments
      .newExperiment(parsedValues)
      .then(() => {
        seteditorOpen(false)
        dispatch(
          setAlert({
            type: 'success',
            message: intl.formatMessage({ id: 'confirm.createSuccessfully' }),
          })
        )
      })
      .catch(console.error)
  }

  const handleDeleteExperiment = async () => {
    const db = await idb.current

    await db.delete('predefined', experiment!.name)

    getExperiments()
    seteditorOpen(false)
    dispatch(
      setAlert({
        type: 'success',
        message: intl.formatMessage({ id: 'confirm.deleteSuccessfully' }),
      })
    )
  }

  return (
    <>
      <Space className={classes.container}>
        <YAML
          callback={saveExperiment}
          buttonProps={{ className: clsx(classes.card, classes.addCard, 'predefined-upload') }}
        />
        {experiments.map((d) => (
          <Card key={d.name} className={classes.card} variant="outlined" onClick={onModalOpen(d)}>
            <Box display="flex" justifyContent="center" alignItems="center" height="100%">
              <Box display="flex" justifyContent="center" flex={1}>
                {iconByKind(d.kind)}
              </Box>
              <Box display="flex" justifyContent="center" flex={2} px={1.5}>
                <Typography>{d.name}</Typography>
              </Box>
            </Box>
          </Card>
        ))}
      </Space>
      <Modal open={editorOpen} onClose={onModalClose}>
        <div>
          <Paper className={classes.editorPaperWrapper} padding={0}>
            {experiment && (
              <Box display="flex" flexDirection="column" height="100%">
                <Box px={3} pt={3}>
                  <PaperTop title={experiment.name}>
                    <Space>
                      <Button color="secondary" size="small" onClick={handleDeleteExperiment}>
                        {T('common.delete')}
                      </Button>
                      <Button variant="contained" color="primary" size="small" onClick={handleApplyExperiment}>
                        {T('common.submit')}
                      </Button>
                    </Space>
                  </PaperTop>
                </Box>
                <Box flex={1}>
                  <YAMLEditor theme={theme} data={yaml.dump(experiment.yaml)} mountEditor={setYAMLEditor} />
                </Box>
              </Box>
            )}
          </Paper>
        </div>
      </Modal>
    </>
  )
}

export default Predefined
