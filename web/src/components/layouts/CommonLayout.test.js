import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import CommonLayout from './CommonLayout'

jest.mock('react-router-dom', () => {
  return {
    ...jest.requireActual('react-router-dom'),
    useLocation: () => {
      return ({
        pathname: '/'
      })
    }
  }
})

describe('CommonLayout Component', () => {
  test('render CommonLayout component successfully', () => {
    const { getByText } = render(
      <BrowserRouter>
        <CommonLayout children='Hello' />
      </BrowserRouter>
    )
    expect(getByText('Hello')).toBeTruthy()
  })
})
