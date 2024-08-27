import CommonLayout from 'components/layouts'
import BodyLayout from 'components/layouts/BodyLayout'
import './RequestManagement.scss'
import Table from 'components/table'
import { CustomSelectTitle } from 'components/custom-select'
import { Link } from 'react-router-dom'
import Button from 'components/button'
import Indicator from 'components/indicator'
import { LIMIT, REQUESTS_URL } from 'utils/contants'
import usePremiseParam from 'hooks/usePremiseParam'
import useGet from 'hooks/useGet'
import { useState } from 'react'
import Loading from 'components/loading'
import { formatDateTime } from '../../utils/datetime'
import TableSummary from 'components/pagination-summary'

const RequestManagement = () => {
  const { premiseId, withPremiseParam, setPremiseParam } = usePremiseParam()
  const [page, setPage] = useState(1)

  const premiseParam = premiseId && premiseId !== 'all' ? { premiseID: premiseId } : {}
  const { response, isLoading, error } = useGet(REQUESTS_URL, {
    page,
    limit: LIMIT,
    ...premiseParam
  })

  if (error) {
    throw error
  }

  const data = response?.data?.items?.map((item) => ({
    ...item,
    premise: {
      name: item.premiseName,
      location: item.premiseLocation
    }
  }))

  const totalCount = response?.data?.pagination?.totalCount
  const pageCount = Math.ceil(totalCount / LIMIT)
  const columns = [
    {
      Header: 'Alert Type',
      id: 'alType',
      accessor: 'alertType',
      minWidth: 120
    },
    {
      Header: 'Premise',
      id: 'premise',
      accessor: 'premise',
      minWidth: 200,
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
      Header: 'Author',
      id: 'author',
      accessor: 'author'
    },
    { Header: 'Assignee', id: 'assignee', accessor: 'assignee' },
    {
      Header: 'Status',
      id: 'status',
      accessor: 'status',
      maxWidth: 100,
      Cell: (props) => (
        <div>
          {/* eslint-disable-next-line react/prop-types */}
          <Indicator status={props.original?.status} hasBackGround />
        </div>
      )
    },
    {
      Header: 'Created Date',
      id: 'startTime',
      minWidth: 130,
      accessor: 'startTime',
      // eslint-disable-next-line react/prop-types
      Cell: (props) => formatDateTime(props.original?.startTime)
    },
    {
      Header: 'Updated Date',
      id: 'verifiedAt',
      minWidth: 130,
      accessor: 'verifiedAt',
      // eslint-disable-next-line react/prop-types
      Cell: (props) => formatDateTime(props.original?.verifiedAt)
    },
    {
      Header: '',
      id: 'action',
      accessor: 'action',
      alignRight: true,
      Cell: (props) => {
        // eslint-disable-next-line react/prop-types
        const requestId = props.original?.id
        return (
          <Link to={withPremiseParam(`/requests/${requestId}`)}>
            <Button buttonType="outlined">View</Button>
          </Link>
        )
      }
    }
  ]
  return (
    <CommonLayout>
      <BodyLayout
        leftContent={
          <CustomSelectTitle
            handlePremisesChange={(data) => {
              setPage(1)
              setPremiseParam(data.value)
            }}
          />
        }>
        {isLoading && <Loading />}
        {!isLoading && (
          <div className="requests-container">
            <Table columns={columns} sortable={false} data={data} />
            {data?.length > 0 && (
              <TableSummary
                page={page}
                data={data}
                pageCount={pageCount}
                totalCount={totalCount}
                onPageChange={({ selected }) => {
                  setPage(selected + 1)
                }}
              />
            )}
          </div>
        )}
      </BodyLayout>
    </CommonLayout>
  )
}

export default RequestManagement
