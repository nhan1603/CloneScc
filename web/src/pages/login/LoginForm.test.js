import { act, render } from '@testing-library/react'
import { screen } from '@testing-library/dom'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import LoginForm from './LoginForm'

jest.mock('data/api', () => {
  return () => ({
    post: () => {
      return {
        data: {
          token: 'token'
        }
      }
    }
  })
})

describe('LoginForm Component', () => {
  test('render LoginForm successfully', async () => {
    const { getByText } = render(
      <BrowserRouter>
        <LoginForm />
      </BrowserRouter>
    )
    await act(() => {
      /* fire events that update state */
      const usernameInput = screen.getAllByTestId('username')[0]
      userEvent.type(usernameInput, "admin")
      const passwordInput = screen.getAllByTestId('password')[0]
      userEvent.type(passwordInput, "admin")
      const submitLogin = screen.getAllByTestId('submit-login')[0]
      userEvent.click(submitLogin)
    });
    expect(getByText('Welcome to Security Command Center')).toBeTruthy()
  })
})
