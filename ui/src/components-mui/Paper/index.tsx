import { Paper as MUIPaper, PaperProps } from '@material-ui/core'

import React from 'react'

const Paper = ({ children, ...rest }: PaperProps) => {
  return (
    <MUIPaper {...rest} variant="outlined">
      {children}
    </MUIPaper>
  )
}

export default Paper
