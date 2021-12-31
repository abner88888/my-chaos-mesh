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

import { Box, BoxProps, useTheme } from '@mui/material'

import EmptyStreet from 'images/assets/undraw_empty_street.svg'
import EmptyStreetDark from 'images/assets/undraw_empty_street-dark.svg'

interface NotFoundProps extends BoxProps {
  illustrated?: boolean
}

const NotFound: React.FC<NotFoundProps> = ({ illustrated = false, children, ...rest }) => {
  const theme = useTheme()

  return (
    <Box
      sx={{
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: 'translate3d(-50%, -50%, 0)',
        ...rest,
      }}
    >
      {illustrated && (
        <Box mb={6}>
          <img
            style={{ width: '50%' }}
            src={theme.palette.mode === 'light' ? EmptyStreet : EmptyStreetDark}
            alt="Not found"
          />
        </Box>
      )}
      {children}
    </Box>
  )
}

export default NotFound
