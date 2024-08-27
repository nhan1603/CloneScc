import { render } from '@testing-library/react'
import InputItem from './InputItem'

describe('Input Item tests', () => {
  test('render with label', () => {
    const { getByText } = render(<InputItem label="Test label" />)
    expect(getByText('Test label')).toBeInTheDocument()
  })

  test('should render with error', () => {
    const { getByText } = render(<InputItem error="input error" />)
    expect(getByText('input error')).toBeInTheDocument()
  })
})
