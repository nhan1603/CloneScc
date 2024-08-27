import { render } from '@testing-library/react'
import Indicator from './Indicator'

describe('Indicator tests', () => {
  test('render Indicator with valid status', () => {
    const { getByText, container } = render(<Indicator status="NEW" />)
    expect(container.getElementsByClassName('indicator new')).toBeTruthy()
    expect(getByText('New')).toBeInTheDocument()
  })

  test('render Indicator with invalid status', () => {
    const { getByText, container } = render(<Indicator status="NEW123" />)
    expect(container.getElementsByClassName('indicator new123')).toBeTruthy()
    expect(getByText('-')).toBeInTheDocument()
  })

  test('render Indicator with valid status and has background', () => {
    const { getByText, container } = render(<Indicator status="RESOLVED" hasBackGround />)
    expect(container.getElementsByClassName('indicator resolved')).toBeTruthy()
    expect(container.getElementsByClassName('resolved--has-bg')).toBeTruthy()
    expect(getByText('RESOLVED')).toBeInTheDocument()
  })
})
