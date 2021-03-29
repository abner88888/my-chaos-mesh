import { Button, IconButton, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from '@material-ui/core'
import ConfirmDialog, { ConfirmDialogHandles } from 'components-mui/ConfirmDialog'
import { useRef, useState } from 'react'

import DeleteOutlinedIcon from '@material-ui/icons/DeleteOutlined'
import Space from 'components-mui/Space'
import T from 'components/T'
import { Workflow } from 'api/workflows.type'
import api from 'api'
import { setAlert } from 'slices/globalStatus'
import { useHistory } from 'react-router-dom'
import { useIntl } from 'react-intl'
import { useStoreDispatch } from 'store'

const initialSelected = {
  namespace: '',
  name: '',
  title: '',
  description: '',
  action: '',
}

interface DataTableProps {
  data: Workflow[]
  fetchData: () => void
}

const DataTable: React.FC<DataTableProps> = ({ data, fetchData }) => {
  const history = useHistory()
  const intl = useIntl()

  const dispatch = useStoreDispatch()

  const [selected, setSelected] = useState(initialSelected)
  const confirmRef = useRef<ConfirmDialogHandles>(null)

  const handleJumpTo = (ns: string, name: string) => () => history.push(`/workflows/${ns}/${name}`)

  const handleSelect = (selected: typeof initialSelected) => (e: any) => {
    e.stopPropagation()

    setSelected(selected)

    confirmRef.current!.setOpen(true)
  }

  const handleAction = (action: string) => () => {
    let actionFunc: any

    switch (action) {
      case 'delete':
        actionFunc = api.workflows.del

        break
      default:
        actionFunc = null
    }

    confirmRef.current!.setOpen(false)

    const { namespace, name } = selected

    if (actionFunc) {
      actionFunc(namespace, name)
        .then(() => {
          dispatch(
            setAlert({
              type: 'success',
              message: intl.formatMessage({ id: `common.${action}Successfully` }),
            })
          )

          setTimeout(fetchData, 300)
        })
        .catch(console.error)
    }
  }

  return (
    <>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>{T('common.name')}</TableCell>
              <TableCell>{T('workflow.entry')}</TableCell>
              <TableCell>{T('workflow.partialTopology')}</TableCell>
              <TableCell>{T('workflow.time')}</TableCell>
              <TableCell>{T('workflow.state')}</TableCell>
              <TableCell>{T('workflow.created')}</TableCell>
              <TableCell>{T('common.operation')}</TableCell>
              <TableCell />
            </TableRow>
          </TableHead>
          <TableBody>
            {data.map((d) => {
              const key = `${d.namespace}/${d.name}`

              return (
                <TableRow key={key} hover>
                  <TableCell>{d.name}</TableCell>
                  <TableCell>{d.entry}</TableCell>
                  <TableCell></TableCell>
                  <TableCell></TableCell>
                  <TableCell></TableCell>
                  <TableCell></TableCell>
                  <TableCell>
                    <Space>
                      <IconButton
                        color="primary"
                        title={intl.formatMessage({ id: 'common.delete' })}
                        aria-label={intl.formatMessage({ id: 'common.delete' })}
                        component="span"
                        size="small"
                        onClick={handleSelect({
                          namespace: d.namespace,
                          name: d.name,
                          title: `${intl.formatMessage({ id: 'common.delete' })} ${d.name}`,
                          description: intl.formatMessage({ id: 'workflows.deleteDesc' }),
                          action: 'delete',
                        })}
                      >
                        <DeleteOutlinedIcon />
                      </IconButton>
                    </Space>
                  </TableCell>
                  <TableCell>
                    <Button variant="outlined" size="small" color="primary" onClick={handleJumpTo(d.namespace, d.name)}>
                      {T('common.detail')}
                    </Button>
                  </TableCell>
                </TableRow>
              )
            })}
          </TableBody>
        </Table>
      </TableContainer>
      <ConfirmDialog
        ref={confirmRef}
        title={selected.title}
        description={selected.description}
        onConfirm={handleAction(selected.action)}
      />
    </>
  )
}

export default DataTable
