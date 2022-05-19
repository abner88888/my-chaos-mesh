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
import { SvgIcon, SvgIconProps } from '@mui/material'

function Experiment(props: SvgIconProps) {
  return (
    <SvgIcon {...props} viewBox="0 0 512 512">
      <path d="m444.609 404.627-108.171-211.322v-163.305h14.707c8.284 0 15-6.716 15-15s-6.716-15-15-15h-190.29c-8.284 0-15 6.716-15 15s6.716 15 15 15h14.707v163.305l-108.171 211.323c-11.799 23.05-10.777 50 2.732 72.092 13.509 22.091 37.033 35.28 62.928 35.28h245.898c25.895 0 49.419-13.189 62.928-35.28 13.51-22.092 14.531-49.042 2.732-72.093zm-240.694-200.871c1.083-2.116 1.647-4.458 1.647-6.835v-166.921h100.875v166.921c0 2.376.564 4.719 1.647 6.835l41.738 81.539h-187.645zm175.034 278.244h-245.898c-15.362 0-29.319-7.825-37.334-20.931-4.509-7.373-6.673-15.657-6.458-23.929.12-4.595.974-9.187 2.568-13.616.638-1.772 54.993-108.229 54.993-108.229h218.36s53.57 104.72 53.948 105.592c2.269 5.229 3.47 10.738 3.613 16.253.215 8.272-1.949 16.556-6.458 23.929-8.014 13.106-21.971 20.931-37.334 20.931z" />
    </SvgIcon>
  )
}

export default Experiment
