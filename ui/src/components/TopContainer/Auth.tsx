import { Box, Button, Link, Typography } from '@material-ui/core'
import React, { useState } from 'react'
import { useHistory, useLocation } from 'react-router-dom'

import ConfirmDialog from 'components-mui/ConfirmDialog'
import RBACGenerator from 'components/RBACGenerator'
import T from 'components/T'
import Token from 'components/Token'

interface AuthProps {
  open: boolean
  setOpen: (open: boolean) => void
}

const Auth: React.FC<AuthProps> = ({ open, setOpen }) => {
  const history = useHistory()
  const { pathname } = useLocation()

  const [genOpen, setGenOpen] = useState(false)

  const handleSubmitCallback = () => {
    setOpen(false)

    history.replace('/authed')
    setTimeout(() => history.replace(pathname))
  }

  const openGenerator = () => setGenOpen(true)
  const closeGenerator = () => setGenOpen(false)

  return (
    <ConfirmDialog
      open={open}
      setOpen={setOpen}
      title={T('settings.addToken.prompt')}
      dialogProps={{
        disableBackdropClick: true,
        disableEscapeKeyDown: true,
        PaperProps: {
          style: { width: 512 },
        },
      }}
    >
      <Box mb={3}>
        <Typography variant="body2" color="textSecondary">
          {T('settings.addToken.prompt2')}{' '}
          <Link style={{ cursor: 'pointer' }} onClick={openGenerator}>
            {T('settings.addToken.prompt3')}
          </Link>
        </Typography>
      </Box>
      <Token onSubmitCallback={handleSubmitCallback} />
      {genOpen && (
        <ConfirmDialog
          open={genOpen}
          setOpen={setGenOpen}
          title={T('settings.addToken.generator')}
          dialogProps={{
            PaperProps: {
              style: { width: 750, maxWidth: 'unset' }, // max-width: 600
            },
          }}
        >
          <RBACGenerator />
          <Box textAlign="right">
            <Button onClick={closeGenerator}>{T('common.close')}</Button>
          </Box>
        </ConfirmDialog>
      )}
    </ConfirmDialog>
  )
}

export default Auth
