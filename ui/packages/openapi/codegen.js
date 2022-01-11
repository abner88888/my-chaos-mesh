/*
 * Copyright 2022 Chaos Mesh Authors.
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

import { genForms } from './index.js'
import { hideBin } from 'yargs/helpers'
import yargs from 'yargs'

const argv = yargs(hideBin(process.argv))
  .command('formik', 'convert CRDs to TypeScript forms with @openapitools/openapi-generator-cli generated')
  .alias('help', 'h')
  .version(false)
  .wrap(120).argv

// eslint-disable-next-line default-case
switch (argv._[0]) {
  case 'formik':
    runFormik()

    break
}

/**
 * Internal function to convert CRDs to TypeScript forms.
 *
 */
function runFormik() {
  genForms('../../app/src/openapi/api.ts')
}
