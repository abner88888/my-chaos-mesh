import { Box, Grid, Grow } from '@material-ui/core'
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
import clsx from 'clsx'
import { makeStyles } from '@material-ui/core/styles'

const useStyles = makeStyles((theme) => ({
  container: {
    position: 'relative',
    height: 300,
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

  useEffect(() => {
    fetchEvents()

    const id = setInterval(fetchEvents, 15000)

    return () => clearInterval(id)
  }, [])

  return (
    <>
      <Grow in={true} style={{ transformOrigin: '0 0 0' }}>
        <Grid container spacing={6}>
          <Grid item xs={12} md={12} lg={3}>
            <Welcome />
          </Grid>

          <Grid item md={12} lg={6}>
            <Paper>
              <PaperTop title={T('dashboard.totalExperiments')} />

              <Box className={clsx(classes.container, classes.totalExperiments)}>
                <TotalExperiments />
              </Box>
            </Paper>
          </Grid>

          <Grid item xs={12} md={12} lg={3}>
            <Paper>
              <PaperTop title={T('dashboard.totalState')} />

              <TotalState className={classes.container} />
            </Paper>
          </Grid>

          <Grid container item xs={12} md={12} lg={9}>
            <Grid item xs={12}>
              <Paper>
                <PaperTop title={T('common.timeline')} />

                <Timeline events={events} className={classes.container} />
              </Paper>
            </Grid>

            <Grid item xs={12}>
              <Box mt={6}>
                <Paper>
                  <PaperTop title={T('dashboard.predefined')} subtitle={T('dashboard.predefinedDesc')} />

                  <Predefined />
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
