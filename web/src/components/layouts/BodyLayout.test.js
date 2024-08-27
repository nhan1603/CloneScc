import { render } from '@testing-library/react'
import BodyLayout from './BodyLayout'

describe('BodyLayout Component', () => {
  test('render BodyLayout component successfully', () => {
    const { getByText } = render(
      <BodyLayout children='Hello' leftContent='left Content' rightContent='right Content' />
    )
    expect(getByText('Hello')).toBeTruthy()
    expect(getByText('left Content')).toBeTruthy()
    expect(getByText('right Content')).toBeTruthy()
  })
})
