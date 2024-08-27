import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import ErrorBoundary from './ErrorBoundary'
import { useRouteError } from 'react-router-dom'

jest.mock('react-router-dom', () => {
  return {
    ...jest.requireActual('react-router-dom'),
    useRouteError: jest.fn()
  }
})

describe('ErrorBoundary Component', () => {
  test('render ErrorBoundary with default value', () => {
    const { getByText } = render(
      <BrowserRouter>
        <ErrorBoundary />
      </BrowserRouter>
    )
    expect(getByText('Oops! Something went wrong')).toBeTruthy()
  })

  test('render ErrorBoundary with Error', () => {
    useRouteError.mockReturnValue(false)
    expect(() => render(ErrorBoundary(false))).toThrow('false')
  })
})
