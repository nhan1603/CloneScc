import { render } from '@testing-library/react'
import BaseSelect from './BaseSelect'

describe('Base Select tests', () => {
  it('should display valid dropdown', () => {
    const options = [
      { value: 1, label: 'Label 1' },
      { value: 2, label: 'Label 2' }
    ]
    const { getByText, container } = render(
      <BaseSelect value={{ value: 2, label: 'Label 2' }} options={options} />
    )
    expect(container.getElementsByClassName('advanced-select-container')).toBeTruthy()
    expect(container.getElementsByClassName('advanced-select')).toBeTruthy()
    expect(getByText('Label 2')).toBeInTheDocument()
  })
})
