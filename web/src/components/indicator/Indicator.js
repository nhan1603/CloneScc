import classNames from 'classnames'
import './Indicator.scss'
import PropTypes from 'prop-types'

export const STATUS = {
  REQUESTED: 'requested',
  RECEIVED: 'received',
  ATTENTION_REQUIRED: 'attention_required',
  RESOLVED: 'resolved',
  VERY_HIGH: 'very_high',
  NEW: 'new'
}

const STATUS_LABEL_MAPPING = {
  [STATUS.REQUESTED]: 'REQUESTED',
  [STATUS.RECEIVED]: 'RECEIVED',
  [STATUS.ATTENTION_REQUIRED]: 'ATTENTION REQUIRED',
  [STATUS.RESOLVED]: 'RESOLVED',
  [STATUS.VERY_HIGH]: 'Very High',
  [STATUS.NEW]: 'New'
}
const Indicator = ({ status, hasBackGround }) => {
  const lowerCaseStatus = status?.toLowerCase()
  const label = STATUS_LABEL_MAPPING[lowerCaseStatus]
  return (
    <div
      className={classNames('indicator', lowerCaseStatus, {
        [`${lowerCaseStatus}--has-bg`]: hasBackGround
      })}>
      {label || '-'}
    </div>
  )
}

export default Indicator

Indicator.propTypes = {
  status: PropTypes.string,
  hasBackGround: PropTypes.bool
}

Indicator.defaultProps = {
  hasBackGround: true
}
