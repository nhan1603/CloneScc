import { render } from '@testing-library/react'
import CustomSelect from './CustomSelect'

describe('CustomSelect Component', () => {
  test('render CustomSelect with default value', () => {
    const { getByText } = render(
      <CustomSelect
        valueContainerLabel='valueContainerLabel'
        subOption='Sub Option'
        value='value'
        options={['value']}
      />
    )
    expect(getByText('Sub Option')).toBeTruthy()
  })

  test('render CustomSelect with empty value and subOption', () => {
    const { getByText } = render(
      <CustomSelect
        valueContainerLabel='Value Label'
        subOption={null}
        value={null}
        options={['value']}
      />
    )
    expect(getByText('Value Label')).toBeTruthy()
  })
})
