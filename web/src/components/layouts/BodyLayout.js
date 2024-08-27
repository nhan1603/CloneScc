import './BodyLayout.scss'
import PropTypes from 'prop-types'

const BodyLayout = ({ children, leftContent, rightContent }) => {
  return (
    <div className="body-layout">
      <div className="body-layout__header">
        <div className="body-layout__header--left">{leftContent}</div>
        {rightContent && <div className="body-layout__header--right">{rightContent}</div>}
        <div></div>
      </div>

      <div className="body-content"> {children}</div>
    </div>
  )
}
export default BodyLayout

BodyLayout.propTypes = {
  children: PropTypes.node,
  leftContent: PropTypes.node,
  rightContent: PropTypes.node
}
