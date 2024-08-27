import PropTypes from 'prop-types'
import classNames from 'classnames'

const InputItem = ({
  className,
  error,
  required,
  htmlFor,
  label,
  onChange,
  type,
  grid,
  ...otherProps
}) => {
  return (
    <div
      className={classNames(className, 'lm--formItem', {
        'is-error': error,
        'lm--formItem--inline': grid
      })}>
      {label && (
        <label
          className={classNames('form-label lm--formItem-label', {
            required,
            'lm--formItem-left': grid
          })}
          htmlFor={htmlFor}>
          {label}
        </label>
      )}
      <div
        className={classNames('input-container', {
          'lm--formItem-right': grid
        })}>
        <input id={htmlFor} onChange={onChange} type={type} {...otherProps} />
        {error && <div className="error-block">{error}</div>}
      </div>
    </div>
  )
}
export default InputItem

InputItem.defaultProps = {
  className: '',
  type: 'text'
}

InputItem.propTypes = {
  className: PropTypes.string,
  onChange: PropTypes.func,
  error: PropTypes.string,
  required: PropTypes.bool,
  htmlFor: PropTypes.string,
  label: PropTypes.string,
  type: PropTypes.string,
  grid: PropTypes.bool
}
