import CommonLayout from 'components/layouts'
import './AlertDetails.scss'
import BodyLayout from 'components/layouts/BodyLayout'
import Button from 'components/button'
import { Link, useParams } from 'react-router-dom'
import Images from 'images'
import usePremiseParam from 'hooks/usePremiseParam'
import useGet from 'hooks/useGet'
import { GET_ALERTS_URL } from 'utils/contants'
import Loading from 'components/loading'
import { formatDateTime } from 'utils/datetime'
import CreateRequestModal from 'components/modal/CreateRequestModal'
import { useState } from 'react'
import { addErrorToast, addSuccessToast } from 'components/toastify/Toast'
import { getFloor } from 'utils'

const AlertDetails = () => {
  const { withPremiseParam } = usePremiseParam()
  const { alertId } = useParams()
  const [isOpen, setIsOpen] = useState(false)
  const { isLoading, error, response } = useGet(`${GET_ALERTS_URL}/${alertId}`)
  if (error) {
    throw error
  }

  const data = response?.data

  return (
    <CommonLayout>
      <BodyLayout
        leftContent={
          <Link className="request-header" to={withPremiseParam('/alerts')}>
            <img src={Images.backIcon} alt="backIcon" />
            <span>Alert Details</span>
          </Link>
        }
        rightContent={
          <Button
            buttonType="solid"
            data-testid="create-request-details"
            onClick={() => {
              setIsOpen(true)
            }}>
            Create Request
          </Button>
        }>
        {isLoading ? (
          <Loading />
        ) : (
          <div className="alert-details">
            <div className="information">
              <div className="title">Type</div>
              <div>{data?.type}</div>
              <div className="title">Date</div>
              <div>{formatDateTime(data?.incidentAt)}</div>
              <div className="title">Premise</div>
              <div>
                <div>{data?.premiseName}</div>
                <div>{data.premiseLocation}</div>
              </div>
              <div className="title">Camera</div>
              <div>
                {data?.cctvDevice} - {getFloor(data?.cctvDeviceFloor)}
              </div>
            </div>
          </div>
        )}
      </BodyLayout>

      <CreateRequestModal
        isOpen={isOpen}
        onClose={() => {
          setIsOpen(false)
        }}
        onError={() => {
          addErrorToast('Send Request failed!')
        }}
        onSuccess={() => {
          addSuccessToast('Send Request Successfully!')
          setIsOpen(false)
        }}
        alertDetails={{
          id: data?.id,
          date: formatDateTime(data?.incidentAt),
          type: data?.type,
          premise: {
            name: data?.premiseName,
            location: data?.premiseLocation
          },
          cameraId: `${data?.cctvDevice} - ${getFloor(data?.cctvDeviceFloor)}`
        }}
      />
    </CommonLayout>
  )
}

export default AlertDetails
