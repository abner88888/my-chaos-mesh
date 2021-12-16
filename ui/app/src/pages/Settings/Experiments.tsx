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
import { Checkbox, FormControl, FormControlLabel, FormGroup, FormHelperText } from '@mui/material'
import { setDebugMode, setEnableKubeSystemNS } from 'slices/settings'
import { useStoreDispatch, useStoreSelector } from 'store'

import PaperTop from '@ui/mui-extends/esm/PaperTop'
import T from 'components/T'

const Experiments = () => {
  const { settings } = useStoreSelector((state) => state)
  const { debugMode, enableKubeSystemNS } = settings
  const dispatch = useStoreDispatch()

  const handleChangeDebugMode = (e: React.ChangeEvent<HTMLInputElement>) => dispatch(setDebugMode(e.target.checked))
  const handleChangeEnableKubeSystemNS = (e: React.ChangeEvent<HTMLInputElement>) =>
    dispatch(setEnableKubeSystemNS(e.target.checked))

  return (
    <>
      <PaperTop title={T('experiments.title')} divider />
      <FormGroup>
        <FormControl>
          <FormControlLabel
            control={<Checkbox color="primary" checked={debugMode} onChange={handleChangeDebugMode} />}
            label={T('settings.debugMode.title')}
          />
          <FormHelperText>{T('settings.debugMode.choose')}</FormHelperText>
        </FormControl>

        <FormControl>
          <FormControlLabel
            control={
              <Checkbox color="primary" checked={enableKubeSystemNS} onChange={handleChangeEnableKubeSystemNS} />
            }
            label={T('settings.enableKubeSystemNS.title')}
          />
          <FormHelperText>{T('settings.enableKubeSystemNS.choose')}</FormHelperText>
        </FormControl>
      </FormGroup>
    </>
  )
}

export default Experiments
