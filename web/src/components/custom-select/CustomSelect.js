/* eslint-disable react/prop-types */
import classNames from 'classnames'
import React, { useCallback } from 'react'
import Select, { components } from 'react-select'
import './CustomSelect.scss'

const { Input, SingleValue, ValueContainer } = components

function CustomValueContainer({ valueContainerLabel, subOption, children, ...props }) {
  return (
    <ValueContainer {...props}>
      {subOption ? (
        <div className="custom-option">
          <div className="custom-option__title">{children}</div>
          <div className="custom-option__sub-content">{subOption}</div>
        </div>
      ) : (
        <div className="value_container">
          <div className="value_container_label">{valueContainerLabel}</div>
          <div className="value_container_children">{children}</div>
        </div>
      )}
    </ValueContainer>
  )
}

function CustomSingleValue({ ...props }) {
  return <SingleValue {...props} className="single_value" />
}

const CustomInput = ({ ...props }) => {
  const hasValue = props.value !== ''
  return <Input {...props} className={classNames(!hasValue && 'input_no_value')} />
}

function CustomSelect({ valueContainerLabel, subOption, value, ...props }) {
  let selectedOption
  if (value === null || typeof value === 'object') {
    selectedOption = value
  } else {
    selectedOption = props.options.find((o) => o.value === value)
  }
  const renderingValueContainer = useCallback(
    (props) => {
      return (
        <CustomValueContainer
          valueContainerLabel={valueContainerLabel}
          subOption={subOption}
          {...props}
        />
      )
    },
    [valueContainerLabel, subOption]
  )
  return (
    <Select
      {...props}
      value={selectedOption}
      placeholder=""
      className={classNames('advanced-select-container', subOption && 'sub-option')}
      classNamePrefix="advanced-select"
      components={{
        ValueContainer: renderingValueContainer,
        SingleValue: CustomSingleValue,
        Input: CustomInput
      }}
    />
  )
}

export default React.memo(CustomSelect)
