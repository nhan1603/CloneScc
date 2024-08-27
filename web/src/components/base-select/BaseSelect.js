import Select from 'react-select'
const BaseSelect = (props) => {
  return (
    <Select {...props} className="advanced-select-container" classNamePrefix="advanced-select" />
  )
}

export default BaseSelect
