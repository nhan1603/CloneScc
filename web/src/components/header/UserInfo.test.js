import { render } from '@testing-library/react'
import { screen } from '@testing-library/dom'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import UserInfo from './UserInfo'

jest.mock('hooks/usePremiseParam', () => {
  return () => ({
    premiseId: '123',
    withPremiseParam: () => {
      return () => {
        return '/hello'
      }
    }
  })
})

describe('UserInfo Component', () => {
  test('render UserInfo with default value', () => {
    const { getByText } = render(
      <BrowserRouter>
        <UserInfo />
      </BrowserRouter>
    )
    const userInfo = screen.getAllByTestId('user-info')[0]
    userEvent.click(userInfo)
    expect(getByText('Logout')).toBeTruthy()

    // Logout
    const logout = screen.getAllByTestId('logout')[0]
    userEvent.click(logout)
  })
})
