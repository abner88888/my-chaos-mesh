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
import NewWorkflow from 'components/NewWorkflow'
import Archives from 'pages/Archives'
import Archive from 'pages/Archives/Single'
import Dashboard from 'pages/Dashboard'
import Events from 'pages/Events'
import Experiments from 'pages/Experiments'
import NewExperiment from 'pages/Experiments/New'
import Experiment from 'pages/Experiments/Single'
import Schedules from 'pages/Schedules'
import NewSchedule from 'pages/Schedules/New'
import Schedule from 'pages/Schedules/Single'
import Settings from 'pages/Settings'
import Workflows from 'pages/Workflows'
import Workflow from 'pages/Workflows/Single'
import { RouteProps } from 'react-router'

const routes: RouteProps[] = [
  {
    element: <Dashboard />,
    path: '/dashboard',
  },
  {
    element: <NewWorkflow />,
    path: '/workflows/new',
  },
  {
    element: <Workflows />,
    path: '/workflows',
  },
  {
    element: <Workflow />,
    path: '/workflows/:uuid',
  },
  {
    element: <NewSchedule />,
    path: '/schedules/new',
  },
  {
    element: <Schedules />,
    path: '/schedules',
  },
  {
    element: <Schedule />,
    path: '/schedules/:uuid',
  },
  {
    element: <NewExperiment />,
    path: '/experiments/new',
  },
  {
    element: <Experiments />,
    path: '/experiments',
  },
  {
    element: <Experiment />,
    path: '/experiments/:uuid',
  },
  {
    element: <Events />,
    path: '/events',
  },
  {
    element: <Archives />,
    path: '/archives',
  },
  {
    element: <Archive />,
    path: '/archives/:uuid',
  },
  {
    element: <Settings />,
    path: '/settings',
  },
]

export default routes
