import { Box, Grid, Grow, Typography } from '@material-ui/core'

import Experiments from './Experiments'
import Other from './Other'
import Paper from 'components-mui/Paper'
import PaperTop from 'components-mui/PaperTop'
import React from 'react'
import T from 'components/T'
import Token from 'components/Token'
import TokensTable from './TokensTable'
import logo from 'images/logo.svg'
import logoWhite from 'images/logo-white.svg'
import { useStoreSelector } from 'store'

const Settings = () => {
  const state = useStoreSelector((state) => state)
  const { securityMode, version } = state.globalStatus
  const { theme } = state.settings

  return (
    <Grow in={true} style={{ transformOrigin: '0 0 0' }}>
      <Grid container>
        <Grid item sm={12} md={8}>
          <Paper>
            {securityMode && (
              <>
                <PaperTop title={T('settings.addToken.title')} />
                <Token />
                <Box mb={6} />
                <TokensTable />
                <Box mt={12} />
              </>
            )}

            <PaperTop title={T('experiments.title')} />
            <Experiments />

            <Box mt={12} />

            <PaperTop title={T('common.other')} />
            <Other />

            <Box mt={12} />

            <PaperTop title={T('common.version')} />
            <Box my={3}>
              <img style={{ height: 36 }} src={theme === 'light' ? logo : logoWhite} alt="Chaos Mesh" />
              <Box mt={1.5}>
                <Typography variant="body2" color="textSecondary">
                  Git Version: {version}
                </Typography>
              </Box>
            </Box>
          </Paper>
        </Grid>
      </Grid>
    </Grow>
  )
}

export default Settings
