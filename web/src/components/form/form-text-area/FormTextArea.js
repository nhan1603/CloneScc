import PropTypes from 'prop-types'
import classNames from 'classnames'
import { useController } from 'react-hook-form'

const FormTextArea = ({
  className,
  error,
  required,
  htmlFor,
  label,
  grid,
  name,
  ...otherProps
}) => {
  const { field, fieldState } = useController({ name })
  const value = field.value
  const fieldError = fieldState.error?.message

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
        <textarea id={htmlFor} value={value} name={name} {...field} {...otherProps} />
        {fieldError && <div className="error-block">{fieldError}</div>}
      </div>
    </div>
  )
}
export default FormTextArea

FormTextArea.defaultProps = {
  className: ''
}

FormTextArea.propTypes = {
  name: PropTypes.string.isRequired,
  className: PropTypes.string,
  onChange: PropTypes.func,
  error: PropTypes.string,
  required: PropTypes.bool,
  htmlFor: PropTypes.string,
  label: PropTypes.string,
  grid: PropTypes.bool
}
