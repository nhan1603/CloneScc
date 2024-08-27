import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import Header from './Header'

jest.mock('react-router-dom', () => {
  return {
    ...jest.requireActual('react-router-dom'),
    useLocation: () => {
      return {
        pathname: '/hello'
      }
    }
  }
})

jest.mock('hooks/usePremiseParam', () => {
  return () => ({
    withPremiseParam: () => {
      return () => {
        return '/hello'
      }
    }
  })
})

describe('Header Component', () => {
  test('render Header component successfully', () => {
    const { getByText } = render(
      <BrowserRouter>
        <Header />
      </BrowserRouter>
    )
    expect(getByText('Security Command Center')).toBeTruthy()
    expect(getByText('Dashboard')).toBeTruthy()
    expect(getByText('Alerts')).toBeTruthy()
  })
})
