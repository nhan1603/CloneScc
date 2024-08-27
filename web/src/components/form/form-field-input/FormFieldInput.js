import InputItem from 'components/input-item'
import { useController } from 'react-hook-form'

/* eslint-disable react/prop-types */
const FormFieldInput = ({ label, name, ...props }) => {
  const { field, fieldState } = useController({ name })
  // don't need passing ref
  // eslint-disable-next-line no-unused-vars
  const { ref, ...fieldProps } = field
  const value = field.value
  const fieldError = fieldState.error?.message

  return (
    <InputItem label={label} error={fieldError} {...fieldProps} {...props} value={value || ''} />
  )
}

export default FormFieldInput
