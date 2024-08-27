import PropTypes from 'prop-types'
import classNames from 'classnames'
import BaseSelect from 'components/base-select'
import { useController } from 'react-hook-form'

const FormBaseSelect = ({
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
  const fieldError = fieldState.error?.value?.message

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
        <BaseSelect
          id={htmlFor}
          value={value}
          onChange={(value) => {
            field.onChange(value)
          }}
          {...otherProps}
        />
        {fieldError && <div className="error-block">{fieldError}</div>}
      </div>
    </div>
  )
}
export default FormBaseSelect

FormBaseSelect.defaultProps = {
  className: ''
}

FormBaseSelect.propTypes = {
  name: PropTypes.string.isRequired,
  className: PropTypes.string,
  onChange: PropTypes.func,
  error: PropTypes.string,
  required: PropTypes.bool,
  htmlFor: PropTypes.string,
  label: PropTypes.string,
  grid: PropTypes.bool
}
