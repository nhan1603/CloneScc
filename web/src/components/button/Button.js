import classNames from 'classnames'
import './Button.scss'
import PropTypes from 'prop-types'

const Button = ({ children, buttonType, buttonSize, type, ...otherProps }) => {
  return (
    <button
      className={classNames('lm--button', `lm--button--${buttonType}`, `lm--button--${buttonSize}`)}
      type={type}
      {...otherProps}>
      {children}
    </button>
  )
}

export default Button

Button.defaultProps = {
  type: 'button'
}

Button.propTypes = {
  children: PropTypes.node,
  type: PropTypes.string,
  buttonType: PropTypes.oneOf(['solid', 'outlined', 'grey']),
  buttonSize: PropTypes.oneOf(['small'])
}
