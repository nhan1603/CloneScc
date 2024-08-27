import { toast } from 'react-toastify'
import PropTypes from 'prop-types'
import './Toast.scss'
import Images from 'images'
import classNames from 'classnames'

const Toast = ({ title, message, date, link, status }) => {
  const renderStatus = () => {
    if (status === 'success') {
      return <img className="icon" src={Images.successIcon} />
    }
    if (status === 'error') {
      return <img className="icon" src={Images.errorIcon} />
    }
    if (status === 'info') {
      return <img className="icon" src={Images.infoIcon} />
    }
  }
  return (
    <div className="toast">
      <div
        className={classNames('toast--left', {
          [`toast--left--${status}`]: status
        })}>
        {renderStatus()}
      </div>
      <div className="toast--right">
        <div className="toast__title">{title}</div>
        <div className="toast__message">{message}</div>
        {date && <div className="toast__date">{date}</div>}
        <a href={link} target="_blank" rel="noreferrer noopener">
          Details &rarr;
        </a>
      </div>
    </div>
  )
}
Toast.propTypes = {
  title: PropTypes.string,
  message: PropTypes.string,
  link: PropTypes.string,
  status: PropTypes.string,
  date: PropTypes.string
}

export default Toast

export const addToast = ({ message, options, ...otherProps }) => {
  toast(() => <Toast message={message} {...otherProps} />, options)
}

export const addSuccessToast = (message, options) => {
  toast.success(message, { position: 'top-right', ...options })
}

export const addErrorToast = (message, options) => {
  toast.error(message, { position: 'top-right', ...options })
}
