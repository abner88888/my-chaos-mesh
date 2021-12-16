/*
 * Copyright 2021 Chaos Mesh Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import { Alert, Box, Container, CssBaseline, Paper, Portal, Snackbar, useMediaQuery, useTheme } from '@mui/material'
import { Navigate, Route, Routes } from 'react-router-dom'
import { drawerCloseWidth, drawerWidth } from './Sidebar'
import { setAlertOpen, setConfig, setConfirmOpen, setNameSpace, setTokenName, setTokens } from 'slices/globalStatus'
import { useEffect, useMemo, useState } from 'react'
import { useStoreDispatch, useStoreSelector } from 'store'

import ConfirmDialog from '@ui/mui-extends/esm/ConfirmDialog'
import Cookies from 'js-cookie'
import { IntlProvider } from 'react-intl'
import LS from 'lib/localStorage'
import Loading from '@ui/mui-extends/esm/Loading'
import Navbar from './Navbar'
import Sidebar from './Sidebar'
import { TokenFormValues } from 'components/Token'
import api from 'api'
import flat from 'flat'
import insertCommonStyle from 'lib/d3/insertCommonStyle'
import loadable from '@loadable/component'
import { makeStyles } from '@mui/styles'
import messages from 'i18n/messages'
import routes from 'routes'
import { setNavigationBreadcrumbs } from 'slices/navigation'
import { useLocation } from 'react-router-dom'

const Auth = loadable(() => import('./Auth'))

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    marginLeft: drawerCloseWidth,
    width: `calc(100% - ${drawerCloseWidth})`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    [theme.breakpoints.down('sm')]: {
      minWidth: theme.breakpoints.values.md,
    },
  },
  rootShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth})`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  main: {
    display: 'flex',
    flexDirection: 'column',
    minHeight: '100vh',
    zIndex: 1,
  },
  switchContent: {
    display: 'flex',
    flex: 1,
  },
}))

const TopContainer = () => {
  const theme = useTheme()
  const classes = useStyles()

  const { pathname } = useLocation()

  const { settings, globalStatus, navigation } = useStoreSelector((state) => state)
  const { lang } = settings
  const { alert, alertOpen, confirm, confirmOpen } = globalStatus
  const { breadcrumbs } = navigation

  const intlMessages = useMemo<Record<string, string>>(() => flat(messages[lang]), [lang])

  const dispatch = useStoreDispatch()
  const handleSnackClose = () => dispatch(setAlertOpen(false))
  const handleConfirmClose = () => dispatch(setConfirmOpen(false))

  // Sidebar related
  const miniSidebar = LS.get('mini-sidebar') === 'y'
  const [openDrawer, setOpenDrawer] = useState(!miniSidebar)
  const handleDrawerToggle = () => {
    setOpenDrawer(!openDrawer)
    LS.set('mini-sidebar', openDrawer ? 'y' : 'n')
  }

  const [loading, setLoading] = useState(true)
  const [authOpen, setAuthOpen] = useState(false)

  useEffect(() => {
    /**
     * Set authorization (RBAC token / GCP) for API use.
     *
     */
    function setAuth() {
      // GCP
      const accessToken = Cookies.get('access_token')
      const expiry = Cookies.get('expiry')

      if (accessToken && expiry) {
        const token = {
          accessToken,
          expiry,
        }

        api.auth.token(token as any)
        dispatch(setTokenName('gcp'))

        return
      }

      const token = LS.get('token')
      const tokenName = LS.get('token-name')
      const globalNamespace = LS.get('global-namespace')

      if (token && tokenName) {
        const tokens: TokenFormValues[] = JSON.parse(token)

        api.auth.token(tokens.find(({ name }) => name === tokenName)!.token)
        dispatch(setTokens(tokens))
        dispatch(setTokenName(tokenName))
      } else {
        setAuthOpen(true)
      }

      if (globalNamespace) {
        api.auth.namespace(globalNamespace)
        dispatch(setNameSpace(globalNamespace))
      }
    }

    /**
     * Render different components according to server configuration.
     *
     */
    function fetchServerConfig() {
      api.common
        .config()
        .then(({ data }) => {
          if (data.security_mode) {
            setAuth()
          }

          dispatch(setConfig(data))
        })
        .finally(() => setLoading(false))
    }

    fetchServerConfig()
    insertCommonStyle()
  }, [dispatch])

  useEffect(() => {
    dispatch(setNavigationBreadcrumbs(pathname))
  }, [dispatch, pathname])

  const isTabletScreen = useMediaQuery(theme.breakpoints.down('md'))
  useEffect(() => {
    if (isTabletScreen) {
      setOpenDrawer(false)
    }
  }, [isTabletScreen])

  return (
    <IntlProvider messages={intlMessages} locale={lang} defaultLocale="en">
      <CssBaseline />

      <Box className={openDrawer ? classes.rootShift : classes.root}>
        <Sidebar open={openDrawer} />
        <Paper className={classes.main} component="main" elevation={0}>
          <Box className={classes.switchContent}>
            <Container maxWidth="xl" sx={{ position: 'relative' }}>
              <Navbar openDrawer={openDrawer} handleDrawerToggle={handleDrawerToggle} breadcrumbs={breadcrumbs} />

              {loading ? (
                <Loading />
              ) : (
                <Routes>
                  <Route path="/" element={<Navigate replace to="/dashboard" />} />
                  {!authOpen && routes.map((route) => <Route key={route.path as string} {...route} />)}
                </Routes>
              )}
            </Container>
          </Box>
        </Paper>
      </Box>

      <Auth open={authOpen} setOpen={setAuthOpen} />

      <Portal>
        <Snackbar
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'center',
          }}
          autoHideDuration={6000}
          open={alertOpen}
          onClose={handleSnackClose}
        >
          <Alert severity={alert.type} onClose={handleSnackClose}>
            {alert.message}
          </Alert>
        </Snackbar>
      </Portal>

      <Portal>
        <ConfirmDialog
          open={confirmOpen}
          close={handleConfirmClose}
          title={confirm.title}
          description={confirm.description}
          onConfirm={confirm.handle}
        />
      </Portal>
    </IntlProvider>
  )
}

export default TopContainer
