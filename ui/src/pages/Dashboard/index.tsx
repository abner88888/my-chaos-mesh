import { Box, Grid, Grow, Typography } from '@material-ui/core'
import React, { useEffect, useState } from 'react'

import { Event } from 'api/events.type'
import Paper from 'components-mui/Paper'
import PaperTop from 'components-mui/PaperTop'
import Predefined from './Predefined'
import Recent from './Recent'
import T from 'components/T'
import Timeline from 'components/Timeline'
import TotalExperiments from './TotalExperiments'
import TotalState from './TotalState'
import Welcome from './Welcome'
import api from 'api'
import { makeStyles } from '@material-ui/core/styles'

const useStyles = makeStyles((theme) => ({
  container: {
    height: 300,
    margin: theme.spacing(3),
  },
  totalExperiments: {
    display: 'flex',
    alignItems: 'center',
    [theme.breakpoints.down('sm')]: {
      alignItems: 'unset',
      overflowY: 'scroll',
    },
  },
}))

export default function Dashboard() {
  const classes = useStyles()

  const [events, setEvents] = useState<Event[]>([])

  const fetchEvents = () => {
    api.events
      .dryEvents()
      .then(({ data }) => setEvents(data))
      .catch(console.error)
  }

  useEffect(fetchEvents, [])

  return (
    <>
      <Grow in={true} style={{ transformOrigin: '0 0 0' }}>
        <Grid container spacing={3}>
          <Grid item md={12} lg={3}>
            <Welcome />
          </Grid>

          <Grid item md={12} lg={6}>
            <Paper>
              <PaperTop title={T('dashboard.totalExperiments')} />

              <Box className={classes.totalExperiments} height={300} m={3}>
                <TotalExperiments />
              </Box>
            </Paper>
          </Grid>

          <Grid item xs={12} md={12} lg={3}>
            <Paper style={{ position: 'relative' }}>
              <PaperTop title={T('dashboard.totalState')} />

              <TotalState className={classes.container} />
            </Paper>
          </Grid>

          <Grid container item xs={12} md={12} lg={9}>
            <Grid item xs={12}>
              <Paper style={{ position: 'relative' }}>
                <PaperTop title={T('common.timeline')} />

                <Timeline events={events} className={classes.container} />
              </Paper>
            </Grid>

            <Grid item xs={12}>
              <Box mt={3}>
                <Paper>
                  <PaperTop title={T('dashboard.predefined')} />

                  <Box height={150} mx={3}>
                    <Typography>{T('dashboard.predefinedDesc')}</Typography>
                    <Box pt={6}>
                      <Predefined />
                    </Box>
                  </Box>
                </Paper>
              </Box>
            </Grid>
          </Grid>

          <Grid item xs={12} md={12} lg={3}>
            <Paper style={{ position: 'relative', height: '100%' }}>
              <PaperTop title={T('dashboard.recent')} />

              <Recent events={events.slice(-6)} />
            </Paper>
          </Grid>
        </Grid>
      </Grow>
    </>
  )
}
