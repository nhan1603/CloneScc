import { useEffect, useRef, useState } from 'react'
import './Notification.scss'
import useClickOutside from 'hooks/useClickOutside'
import Images from 'images'
import Button from 'components/button'
import { removeAlertMessages, getAlertMessages } from 'data/alertMessage'
import { formatDateTime } from 'utils/datetime'
import useWebSocket from 'react-use-websocket'
import { getFloor } from 'utils'
import { Link } from 'react-router-dom'
import usePremiseParam from 'hooks/usePremiseParam'

const Notification = () => {
  const { lastMessage } = useWebSocket(process.env.REACT_APP_SOCKET_API_ENDPOINT)
  const alertMessages = getAlertMessages()
  const [hasNotification, setHasNotification] = useState(alertMessages?.length > 0)
  const [isOpen, setIsOpen] = useState(false)
  const ref = useRef()
  const { withPremiseParam } = usePremiseParam()

  useClickOutside({
    ref,
    isOpen,
    handleClose: () => {
      setIsOpen(false)
    }
  })

  const handleMarkAll = () => {
    removeAlertMessages()
    setIsOpen(false)
    setHasNotification(false)
  }

  useEffect(() => {
    lastMessage && setHasNotification(!!lastMessage)
  }, [lastMessage])

  return (
    <div
      className="notification"
      data-testid="notification"
      ref={ref}
      onClick={() => {
        if (!isOpen) {
          setIsOpen(true)
        }
      }}>
      <img className="bell" src={Images.bell}></img>
      {hasNotification > 0 && <span className="counter"></span>}
      {isOpen && (
        <div id="notification-list-id" className="notification-list">
          <div className="notification-list-container">
            {!hasNotification && <div className="no-records">You have no new notifications</div>}
            {hasNotification &&
              alertMessages?.map((item, index) => {
                return (
                  <Link
                    className="item"
                    key={index}
                    target="_blank"
                    rel="noreferrer noopener"
                    to={withPremiseParam(item?.id ? `/alerts/${item.id}` : '/alerts/')}>
                    <div className="item__title">
                      <div>{item.cctv_name + ' - ' + getFloor(item.floor_number)}</div>
                    </div>
                    <div className="item__message">{item.description}</div>
                    <div className="item__date">{formatDateTime(item.incident_at)}</div>
                  </Link>
                )
              })}
          </div>
          {hasNotification && (
            <Button onClick={handleMarkAll} data-testid="mark-all" buttonType="solid">
              Mark all as read
            </Button>
          )}
        </div>
      )}
    </div>
  )
}

export default Notification
