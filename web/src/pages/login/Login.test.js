import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import Login from './Login'
import * as auth from 'data/auth'

describe('Login Component', () => {
  test('render Login successfully', () => {
    const { getByText } = render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    )
    auth.getToken = () => true
    expect(getByText('Welcome to Security Command Center')).toBeTruthy()
  })

  test('render Login: Redirect to Dashboard when user has token', () => {
    const { queryByText } = render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    )
    auth.getToken = () => false
    const welcomeTitle = queryByText('Welcome to Security Command Center')
    expect(welcomeTitle).not.toBeInTheDocument()
  })
})
