import CommonLayout from 'components/layouts'
import BodyLayout from 'components/layouts/BodyLayout'
import classNames from 'classnames'
import Images from 'images'
import './RequestManagementDetails.scss'
import Button from 'components/button'
import CreateRequestModal from 'components/modal/CreateRequestModal'
import { useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import useGet from 'hooks/useGet'
import { REQUESTS_URL } from 'utils/contants'
import Loading from 'components/loading'
import { formatDateTime } from 'utils/datetime'
import usePremiseParam from 'hooks/usePremiseParam'
import { addErrorToast, addSuccessToast } from 'components/toastify/Toast'
import { getFloor, checkIsVideoOfUrl } from 'utils'
import Zoom from 'react-medium-image-zoom'

const renderRequestInformation = (infoData, title, isFullWidth = false) => {
  return (
    <div className="requests-details__block">
      <div className="requests-details__title">{title}</div>
      <div className="requests-details__content">
        {Array.isArray(infoData)
          ? infoData.map((item) => {
              return (
                <div
                  className={classNames('requests-details__field', isFullWidth && 'is-fullwidth')}
                  key={item.label}>
                  <div className="requests-details__name">{item.label}</div>
                  {Array.isArray(item.value) ? (
                    <div className="requests-details__images">
                      {item.value.map((image, index) => {
                        const isVideo = checkIsVideoOfUrl(image.FileExtension)
                        return isVideo ? (
                          <iframe
                            key={index}
                            width="100%"
                            height="400"
                            src={image.URL}
                            frameBorder="0"
                            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                            allowFullScreen
                            title="Embedded youtube"
                          />
                        ) : (
                          <Zoom>
                            <img key={index} src={image.URL} alt={image.FileName} />
                          </Zoom>
                        )
                      })}
                    </div>
                  ) : (
                    <div className="requests-details__info">{item.value}</div>
                  )}
                </div>
              )
            })
          : infoData}
      </div>
    </div>
  )
}

const RequestManagementDetails = () => {
  const [isOpen, setIsOpen] = useState(false)
  const { requestId = '' } = useParams()
  const { withPremiseParam } = usePremiseParam()
  const { response, isLoading, error } = useGet(`${REQUESTS_URL}/${requestId}`)

  if (error) throw error

  const data = response?.data

  return (
    <CommonLayout>
      <BodyLayout
        leftContent={
          <Link className="request-header" to={withPremiseParam('/requests')}>
            <img src={Images.backIcon} alt="backIcon" />
            <span> Request Details</span>
          </Link>
        }
        rightContent={
          <Button buttonType="solid" onClick={() => setIsOpen(true)}>
            Create Request
          </Button>
        }>
        {isLoading && <Loading />}
        {data && (
          <div className="requests-details">
            {/* TODO: Will enhance this later */}
            {renderRequestInformation(getAlertData(data), 'Alert Details')}
            {renderRequestInformation(getRequestData(data), 'Request')}
            {renderRequestInformation(getRespondData(data), 'Response', true)}
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
          id: data?.alertDetail?.id,
          date: formatDateTime(data?.alertDetail?.incidentAt),
          type: data?.alertDetail?.type,
          premise: {
            name: data?.alertDetail?.premiseName,
            location: data?.alertDetail?.premiseLocation
          },
          cameraId: `${data?.alertDetail?.cctvDevice} - ${getFloor(
            data?.alertDetail?.cctvDeviceFloor
          )}`
        }}
      />
    </CommonLayout>
  )
}

const getAlertData = (data = {}) => {
  return [
    {
      label: 'Premises',
      value: data.alertDetail?.premiseName
    },
    {
      label: 'Alert Type',
      value: data.alertDetail?.type
    },
    {
      label: 'Camera ID',
      value: `${data.alertDetail?.cctvDevice} - ${getFloor(data?.alertDetail?.cctvDeviceFloor)}`
    },
    {
      label: 'date',
      value: formatDateTime(data.alertDetail?.incidentAt)
    }
  ]
}

const getRequestData = (data = {}) => [
  {
    label: 'Author',
    value: data.author
  },
  {
    label: 'Date',
    value: formatDateTime(data.startTime)
  },
  {
    label: 'Assignee',
    value: data.assignee
  },
  {
    label: 'Message',
    value: data.message
  }
]

const getRespondData = (data = {}) => [
  {
    label: 'Date',
    value: formatDateTime(data.respond?.verifiedAt)
  },
  {
    label: 'User',
    value: data.respond?.user
  },
  {
    label: 'Message',
    value: data.respond?.message
  },
  {
    label: 'Media',
    value: data.respond?.mediaData
  }
]

export default RequestManagementDetails
