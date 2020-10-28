import { Box, Card, Divider, GridList, GridListTile, Typography, useMediaQuery, useTheme } from '@material-ui/core'
import React, { useEffect, useState } from 'react'
import { RootState, useStoreDispatch } from 'store'
import { createStyles, makeStyles } from '@material-ui/core/styles'
import { setStep1, setTarget as setTargetToStore } from 'slices/experiments'
import targetData, { Category, Kind } from './data/target'

import CheckCircleOutlineIcon from '@material-ui/icons/CheckCircleOutline'
import Paper from 'components-mui/Paper'
import PaperTop from 'components/PaperTop'
import RadioButtonCheckedOutlinedIcon from '@material-ui/icons/RadioButtonCheckedOutlined'
import RadioButtonUncheckedOutlinedIcon from '@material-ui/icons/RadioButtonUncheckedOutlined'
import Stress from './form/Stress'
import T from 'components/T'
import TargetGenerated from './form/TargetGenerated'
import UndoIcon from '@material-ui/icons/Undo'
import _snakecase from 'lodash.snakecase'
import clsx from 'clsx'
import { useSelector } from 'react-redux'

const useStyles = makeStyles((theme) => {
  const cardActive = {
    color: theme.palette.primary.main,
    borderColor: theme.palette.primary.main,
  }

  return createStyles({
    gridList: {
      flexWrap: 'nowrap',
      transform: 'translateZ(0)',
    },
    card: {
      cursor: 'pointer',
      '&:hover': cardActive,
    },
    cardActive,
    submit: {
      borderColor: theme.palette.success.main,
    },
    submitIcon: {
      color: theme.palette.success.main,
    },
    asButton: {
      cursor: 'pointer',
    },
  })
})

const submitDirectly = ['pod-failure', 'pod-kill']

const Step1 = () => {
  const theme = useTheme()
  const isDesktopScreen = useMediaQuery(theme.breakpoints.down('md'))
  const classes = useStyles()

  const { kind: _kind, action: _action, step1 } = useSelector((state: RootState) => state.experiments)
  const dispatch = useStoreDispatch()

  const [kind, setKind] = useState<Kind | ''>(_kind)
  const [action, setAction] = useState<Omit<Category, 'name'> | undefined>(_action)

  useEffect(() => {
    setKind(_kind)
    setAction(_action)
  }, [_kind, _action])

  const handleSelectTarget = (key: Kind) => () => {
    setKind(key)
    setAction(undefined)
  }
  const handleSelectAction = (d: Omit<Category, 'name'>) => () => {
    if (submitDirectly.includes(d.key)) {
      handleSubmitStep1(d.spec)
    }

    setAction(d)
  }

  const handleSubmitStep1 = (values: Record<string, any>) => {
    const result = {
      kind,
      [_snakecase(kind)]: values,
    }

    if (process.env.NODE_ENV === 'development') {
      console.debug('Debug handleSubmitStep1', result)
    }

    dispatch(setTargetToStore(result))
    dispatch(setStep1(true))
  }

  const handleUndo = () => dispatch(setStep1(false))

  return (
    <Paper className={step1 ? classes.submit : ''}>
      <PaperTop
        title={
          <Box display="flex">
            {step1 && (
              <Box display="flex" alignItems="center" mr={3}>
                <CheckCircleOutlineIcon className={classes.submitIcon} />
              </Box>
            )}
            {T('newE.titleStep1')}
          </Box>
        }
      >
        {step1 && (
          <Box display="flex" alignItems="center">
            <UndoIcon className={classes.asButton} onClick={handleUndo} />
          </Box>
        )}
      </PaperTop>
      <Box hidden={step1}>
        <Box p={6} overflow="hidden">
          <GridList className={classes.gridList} cols={isDesktopScreen ? 1.5 : 4.5} spacing={9} cellHeight="auto">
            {Object.entries(targetData).map(([key, t]) => (
              <GridListTile key={key}>
                <Card
                  className={clsx(classes.card, kind === key ? classes.cardActive : '')}
                  variant="outlined"
                  onClick={handleSelectTarget(key as Kind)}
                >
                  <Box display="flex" justifyContent="center" alignItems="center" height="100px">
                    <Box display="flex" justifyContent="center" alignItems="center" flex={1}>
                      {t.icon}
                    </Box>
                    <Box
                      display="flex"
                      justifyContent="center"
                      alignItems="center"
                      flex={2}
                      px={1.5}
                      textAlign="center"
                    >
                      <Typography variant="button">{t.name}</Typography>
                    </Box>
                  </Box>
                </Card>
              </GridListTile>
            ))}
          </GridList>
        </Box>
        {kind && (
          <>
            <Divider />
            <Box p={6} overflow="hidden">
              {targetData[kind].categories ? (
                <GridList className={classes.gridList} cols={isDesktopScreen ? 2.5 : 5.5} spacing={9} cellHeight="auto">
                  {targetData[kind].categories!.map((d: any) => (
                    <GridListTile key={d.key}>
                      <Card
                        className={clsx(classes.card, action?.key === d.key ? classes.cardActive : '')}
                        variant="outlined"
                        onClick={handleSelectAction(d)}
                      >
                        <Box display="flex" justifyContent="center" alignItems="center" height="50px">
                          <Box display="flex" justifyContent="center" alignItems="center" flex={1}>
                            {action?.key === d.key ? (
                              <RadioButtonCheckedOutlinedIcon />
                            ) : (
                              <RadioButtonUncheckedOutlinedIcon />
                            )}
                          </Box>
                          <Box
                            display="flex"
                            justifyContent="center"
                            alignItems="center"
                            flex={2}
                            px={1.5}
                            textAlign="center"
                          >
                            <Typography variant="button">{d.name}</Typography>
                          </Box>
                        </Box>
                      </Card>
                    </GridListTile>
                  ))}
                </GridList>
              ) : kind === 'TimeChaos' ? (
                <TargetGenerated data={targetData[kind].spec!} onSubmit={handleSubmitStep1} />
              ) : kind === 'StressChaos' ? (
                <Stress onSubmit={handleSubmitStep1} />
              ) : null}
            </Box>
          </>
        )}
        {action && !submitDirectly.includes(action.key) && (
          <>
            <Divider />
            <Box p={6}>
              {/* Force re-render when spec changed */}
              <TargetGenerated key={JSON.stringify(action.spec)} data={action.spec} onSubmit={handleSubmitStep1} />
            </Box>
          </>
        )}
      </Box>
    </Paper>
  )
}

export default Step1
