import Button from 'components/button'
import CommonLayout from 'components/layouts'
import BodyLayout from 'components/layouts/BodyLayout'
import './Alerts.scss'
import Table from 'components/table'
import { CustomSelectTitle } from 'components/custom-select'
import CreateRequestModal from 'components/modal/CreateRequestModal'
import { useState } from 'react'
import useGet from 'hooks/useGet'
import { formatDateTime } from 'utils/datetime'
import { GET_ALERTS_URL, LIMIT } from 'utils/contants'
import usePremiseParam from 'hooks/usePremiseParam'
import { addErrorToast, addSuccessToast } from 'components/toastify/Toast'
import { Link } from 'react-router-dom'
import { getFloor } from 'utils'
import PaginationSummary from 'components/pagination-summary'

const Alerts = () => {
  const [selected, setSelected] = useState(null)
  const [page, setPage] = useState(1)
  const { premiseId, withPremiseParam } = usePremiseParam()

  const [premiseSelected, setPremiseSelected] = useState(premiseId)

  const premiseIdParam = {}
  if (premiseSelected && premiseSelected !== 'all') {
    premiseIdParam.premiseID = premiseSelected
  }
  const { response, isLoading, error } = useGet(GET_ALERTS_URL, {
    limit: LIMIT,
    page,
    ...premiseIdParam
  })

  if (error) {
    throw error
  }

  const data =
    response?.data?.items?.map((item) => ({
      ...item,
      date: item.incidentAt,
      premise: {
        name: item.premiseName,
        location: item.premiseLocation
      },
      cameraId: `${item.cctvDevice} -  ${getFloor(item.cctvDeviceFloor)}`
    })) ?? []

  const totalCount = response?.data?.pagination?.totalCount ?? 0
  const pageCount = Math.ceil(totalCount / LIMIT)
  const columns = [
    {
      Header: 'Type',
      id: 'type',
      accessor: 'type'
    },
    {
      Header: 'Premise',
      id: 'premise',
      accessor: 'premise',
      minWidth: 180,
      Cell: (props) => {
        // eslint-disable-next-line react/prop-types
        const premise = props.original?.premise
        return (
          <div>
            <div>{premise.name}</div>
            <div>{premise.location}</div>
          </div>
        )
      }
    },
    {
      Header: 'CameraID',
      id: 'cameraId',
      accessor: 'cameraId'
    },
    {
      Header: 'Date',
      id: 'date',
      accessor: 'date',
      // eslint-disable-next-line react/prop-types
      Cell: (props) => formatDateTime(props.value)
    },
    {
      Header: '',
      id: 'action',
      accessor: 'action',
      alignRight: true,
      minWidth: 180,
      Cell: (props) => (
        <div>
          {/* eslint-disable-next-line react/prop-types */}
          <Link to={withPremiseParam(`/alerts/${props.original?.id}`)}>
            <Button buttonType="outlined">View</Button>
          </Link>{' '}
          <Button
            buttonType="solid"
            // eslint-disable-next-line react/prop-types
            onClick={() => setSelected(props.original)}
            data-testid="create-request">
            Create Request
          </Button>
        </div>
      )
    }
  ]
  return (
    <CommonLayout>
      <BodyLayout
        leftContent={
          <CustomSelectTitle
            handlePremisesChange={(data) => {
              // reset page when premise changed
              setPage(1)
              setPremiseSelected(data.value)
            }}
          />
        }>
        <div className="alerts-container">
          <Table columns={columns} loading={isLoading} sortable={false} data={data} />
          {data.length > 0 && (
            <PaginationSummary
              page={page}
              data={data}
              totalCount={totalCount}
              pageCount={pageCount}
              onPageChange={({ selected }) => {
                setPage(selected + 1)
              }}
            />
          )}
        </div>
      </BodyLayout>
      <CreateRequestModal
        isOpen={selected !== null}
        onClose={() => {
          setSelected(null)
        }}
        onError={() => {
          addErrorToast('Send Request failed!')
        }}
        onSuccess={() => {
          addSuccessToast('Send Request Successfully!')
          setSelected(null)
        }}
        alertDetails={{
          id: selected?.id,
          date: formatDateTime(selected?.date),
          type: selected?.type,
          premise: {
            name: selected?.premiseName,
            location: selected?.premiseLocation
          },
          cameraId: selected?.cameraId
        }}
      />
    </CommonLayout>
  )
}

export default Alerts
